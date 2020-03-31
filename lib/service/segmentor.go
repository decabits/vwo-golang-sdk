package service

import (
	"github.com/decabits/vwo-golang-sdk/lib/constants"
	"github.com/decabits/vwo-golang-sdk/lib/schema"
	"github.com/decabits/vwo-golang-sdk/lib/utils"
)

type segmentFunc interface {
}

type segmentData struct {
}

// Evaluate function
func Evaluate(segments []schema.Segment, customVariables []int) bool {
	operator, subSegments := utils.GetKeyValue(segments)
	if operator == constants.OperatorTypeNot {
		value := Evaluate(subSegments, customVariables)
		return !value
	} else if operator == constants.OperatorTypeAnd {
		for _, v := range subSegments {
			value := Evaluate(v, customVariables)
			if value == false {
				return false
			}
		}
		return true
	} else if operator == constants.OperatorTypeOr {
		for _, v := range subSegments {
			value := Evaluate(v, customVariables)
			if value == true {
				return true
			}
		}
		return false
	} else if operator == constants.OperandTypesCustomVariable {
		for _, v := range subSegments {
			value := evaluateCustomVariable(operand, customVariables)
			if value == true {
				return true
			}
		}
		return false
	} else {
		for _, v := range subSegments {
			value := evaluateCustomVariable(operand, customVariables)
			if value == true {
				return true
			}
		}
		return false
	}
}

func evaluateCustomVariable(operand []schema.Segment, customVariables []int) bool {
	operandKey, operand := utils.GetKeyValue(operand)
	//customVariables[operandKey] := utils.ProcessCustomVariablesValue(customVariables[operandKey])
	//operandType, operandValue := utils.ProcessOperandValue(operand)
	return true
}
