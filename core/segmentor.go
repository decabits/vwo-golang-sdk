package core

import (
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

// SegmentEvaluator function evaluates segments to get the keys and values and perform appropriate functions
func SegmentEvaluator(segments map[string]interface{}, options schema.Options) bool {
	/*
		Args:
			segments: segments from campaign or variation
			options: options object containing CustomVariables, VariationTargertting variables and Revenue Goal
		
		Returns: 
			bool: if the options falls in the segments criteria
	*/

	operator, subSegments := utils.GetKeyValue(segments)

	if operator == constants.OperatorTypeNot {
		return SegmentEvaluator(subSegments.(map[string]interface{}), options) == false
	} else if operator == constants.OperatorTypeAnd {
		var res []bool
		for _, v := range subSegments.([]interface{}) {
			res = append(res, SegmentEvaluator(v.(map[string]interface{}), options))
		}
		return evaluate(operator, res)
	} else if operator == constants.OperatorTypeOr {
		var res []bool
		for _, v := range subSegments.([]interface{}) {
			res = append(res, SegmentEvaluator(v.(map[string]interface{}), options))
		}
		return evaluate(operator, res)
	} else if operator == constants.OperandTypesCustomVariable {
		return evaluateCustomVariables(subSegments.(map[string]interface{}), options)
	} else if operator == constants.OperandTypesUser {
		return operandUserParser(subSegments.(string), options)
	}
	return true
}

// evaluate function checks the res array, if operator is or then performs or on all elements else and
func evaluate(operator string, res []bool) bool {
	/*
		Args:
			operator: AND or OR operator
			res: array of bool values
		
		Returns: 
			bool: final computed value of OR or AND
	*/

	if operator == constants.OperatorTypeAnd {
		for _, v := range res {
			if v == false {
				return false
			}
		}
		return true
	} else if operator == constants.OperatorTypeOr {
		for _, v := range res {
			if v == true {
				return true
			}
		}
		return false
	}
	return false
}

//evaluateCustomVariables function processes the custom variables in the segments
func evaluateCustomVariables(custom map[string]interface{}, options schema.Options) bool {
	/*
		Args:
			segments: segments from campaign or variation
			options: options object containing CustomVariables, VariationTargertting variables and Revenue Goal
		
		Returns: 
			bool: if the options falls in the segments criteria
	*/

	operandKey, operand := utils.GetKeyValue(custom)
	_, okCustomVar := options.CustomVariables[operandKey]
	_, okVariationTar := options.VariationTargetingVariables[operandKey]
	if !okCustomVar && !okVariationTar {
		return false
	}
	var tag interface{}
	if okCustomVar {
		tag = options.CustomVariables[operandKey]
	} else if okVariationTar {
		tag = options.VariationTargetingVariables[operandKey]
	}

	operandType, operandValue := preProcessOperandValue(operand)
	tagValue := processCustomVariablesValue(tag)
	processedValues, tagValue := processValues(operandValue, tagValue)
	return extractResult(operandType, processedValues, tagValue)
}

// extractResult function compares the operand value and tag value on the basis of operand type
func extractResult(operandType int, operandValue, tagValue string) bool {
	result := false
	switch operandType {
	case constants.LowerValue:
		if tagValue != "" {
			result = strings.ToLower(operandValue) == strings.ToLower(tagValue)
		}
	case constants.StartingEndingStarValue:
		if tagValue != "" {
			result = strings.Index(tagValue, operandValue) > -1
		}
	case constants.StartingStarValue:
		if tagValue != "" {
			result = strings.HasSuffix(tagValue, operandValue)
		}
	case constants.EndingStarValue:
		if tagValue != "" {
			result = strings.HasPrefix(tagValue, operandValue)
		}
	case constants.RegexValue:
		result = matchWithRegex(tagValue, operandValue)
	default:
		result = tagValue == operandValue
	}
	return result
}

//operandUserParser function checks if the VWO user lies in the list of users in the segments
func operandUserParser(operand string, options schema.Options) bool {
	/*
		Args:
			operand: list of users
			options: options object containing CustomVariables, VariationTargertting variables and Revenue Goal

		Returns:
			bool: true if user in list, else false
	*/

	users := strings.Split(operand, ",")
	for _, user := range users {
		if strings.TrimSpace(user) == options.CustomVariables["_vwo_user_id"] || strings.TrimSpace(user) == options.VariationTargetingVariables["_vwo_user_id"] {
			return true
		}
	}
	return false
}

// processCustomVariablesValue function converts interface value of customVariables to string
func processCustomVariablesValue(value interface{}) string {
	switch value.(type) {
	// handle cases
	case bool:
		return strconv.FormatBool(value.(bool))
	case float64:
		return strconv.FormatFloat(value.(float64), 'f', -1, 64)
	case nil:
		return ""
	}
	return value.(string)
}

// preProcessOperandValue
func preProcessOperandValue(operand interface{}) (operandType int, operandValue string) {
	if matchWithRegex(operand.(string), constants.LowerMatch) {
		operandType = constants.LowerValue
		operandValue = extractOperandValue(operand.(string), constants.LowerMatch)
	} else if matchWithRegex(operand.(string), constants.WildcardMatch) {
		operandValue = extractOperandValue(operand.(string), constants.WildcardMatch)
		startingStar := matchWithRegex(operandValue, constants.StartingStar)
		endingStar := matchWithRegex(operandValue, constants.EndingStar)
		// In case of wildcard, the operand type is further divided into contains, startswith and endswith
		if startingStar && endingStar {
			operandType = constants.StartingEndingStarValue
		} else if startingStar {
			operandType = constants.StartingStarValue
		} else if endingStar {
			operandType = constants.EndingStarValue
		}
		remStartStar := regexp.MustCompile(constants.StartingStar)
		remEndingStar := regexp.MustCompile(constants.EndingStar)
		operandValue = remStartStar.ReplaceAllString(operandValue, "")
		operandValue = remEndingStar.ReplaceAllString(operandValue, "")
	} else if matchWithRegex(operand.(string), constants.RegexMatch) {
		operandType = constants.RegexValue
		operandValue = extractOperandValue(operand.(string), constants.RegexMatch)
	} else {
		operandType = constants.EqualValue
		operandValue = operand.(string)
	}
	return
}

func processValues(operandValue string, tagValue interface{}) (newProcessedOperandValue string, newProcessedTagValue string) {
	processedOperandValue, err := strconv.ParseFloat(operandValue, 64)
	if err != nil {
		return operandValue, tagValue.(string)
	}
	processedTagValue, err := strconv.ParseFloat(tagValue.(string), 64)
	if err != nil {
		return operandValue, tagValue.(string)
	}
	// now we have surity that both are numbers
	// now we can convert them independently to int type if they
	// are int rather than floats
	if processedOperandValue == math.Floor(processedOperandValue) {
		newProcessedOperandValue = strconv.Itoa(int(processedOperandValue))
	} else {
		newProcessedOperandValue = strconv.FormatFloat(processedOperandValue, 'f', -1, 64)
		trailZero := regexp.MustCompile("0*$")
		newProcessedOperandValue = trailZero.ReplaceAllString(newProcessedOperandValue, "")
	}
	if processedTagValue == math.Floor(processedTagValue) {
		newProcessedTagValue = strconv.Itoa(int(processedTagValue))
	} else {
		newProcessedTagValue = strconv.FormatFloat(processedTagValue, 'f', -1, 64)
	}

	// convert it back to string and return
	return
}

// matchWithRegex function reports whether the string s contains any match of the regular expression pattern
func matchWithRegex(operand, regex string) bool {
	result, err := regexp.MatchString(regex, operand)
	if err != nil {
		return false
	}
	return result
}

//extractOperandValue function a string holding the text of the leftmost match of the regular expression in s and the matches, if any, of its subexpressions, as defined by the 'Submatch' description in the package comment. A return value of nil indicates no match.
func extractOperandValue(operand, regex string) string {
	re := regexp.MustCompile(regex)
	submatchall := re.FindStringSubmatch(operand)
	if len(submatchall) > 0 {
		return submatchall[1]
	}
	return ""
}
