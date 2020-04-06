package core

import (
	"strings"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

// SegmentEvaluator function
func SegmentEvaluator(segments map[string]interface{}, options schema.Options) bool {
	operator, subSegments := utils.GetKeyValue(segments)

	if operator == constants.OperatorTypeNot {
		return !SegmentEvaluator(subSegments.(map[string]interface{}), options)
	} else if operator == constants.OperatorTypeAnd {
		var res []bool
		for _, v := range subSegments.([]interface{}) {
			res = append(res, SegmentEvaluator(v.(map[string]interface{}), options))
		}
		return evaluate(operator, res, options)
	} else if operator == constants.OperatorTypeOr {
		var res []bool
		for _, v := range subSegments.([]interface{}) {
			res = append(res, SegmentEvaluator(v.(map[string]interface{}), options))
		}
		return evaluate(operator, res, options)
	} else if operator == constants.OperandTypesCustomVariable {
		return evaluateCustomVariables(subSegments.(map[string]interface{}), options)
	} else if operator == constants.OperandTypesUser {
		return operandUserParser(subSegments.(string), options)
	}
	return true
}

func evaluate(operator string, res []bool, options schema.Options) bool {
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
	_, ok := options.CustomVariables["key"]
	if !ok {
		return false
	}
	tagValue := options.CustomVariables[operandKey]
	processedTagValue := utils.ProcessCustomVariablesValue(tagValue)
	operandType, operandValue := utils.PreProcessOperandValue(operand)
	processedValues, tagValue := utils.ProcessValues(operandValue, processedTagValue)
	return extractResult(operandType, processedValues, processedTagValue)
}

func extractResult(operandType int, operandValue, tagValue string) bool {
	var result bool
	switch operandType {
	case constants.LowerValue:
		if tagValue != "" {
			result = strings.ToLower(operandValue) == strings.ToLower(tagValue)
		}
	case constants.StartingEndingStarValue:
		if tagValue != "" {
			result = strings.Index(operandValue, tagValue) > -1
		}
	case constants.StartingStarValue:
		if tagValue != "" {
			result = strings.HasSuffix(operandValue, tagValue)
		}
	case constants.EndingStarValue:
		if tagValue != "" {
			result = strings.HasPrefix(operandValue, tagValue)
		}
	case constants.RegexValue:
		result = strings.Contains(operandValue, tagValue)
	default:
		result = tagValue == operandValue
	}
	return result
}

func operandUserParser(operand string, options schema.Options) bool {
	users := strings.Split(operand, ",")
	for _, user := range users {
		if strings.Trim(user, " ") == options.VWOUserID {
			return true
		}
	}
	return false
}
