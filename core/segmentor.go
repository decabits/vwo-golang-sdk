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

// SegmentEvaluator function
func SegmentEvaluator(segments map[string]interface{}, options schema.Options) bool {
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

func evaluate(operator string, res []bool) bool {
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

func evaluateCustomVariables(custom map[string]interface{}, options schema.Options) bool {
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

func operandUserParser(operand string, options schema.Options) bool {
	users := strings.Split(operand, ",")
	for _, user := range users {
		if strings.TrimSpace(user) == options.CustomVariables["_vwo_user_id"] || strings.TrimSpace(user) == options.VariationTargetingVariables["_vwo_user_id"] {
			return true
		}
	}
	return false
}

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

// ProcessValues function
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

func matchWithRegex(operand, regex string) bool {
	result, err := regexp.MatchString(regex, operand)
	if err != nil {
		return false
	}
	return result
}

func extractOperandValue(operand, regex string) string {
	re := regexp.MustCompile(regex)
	submatchall := re.FindStringSubmatch(operand)
	if len(submatchall) > 0 {
		return submatchall[1]
	}
	return ""
}
