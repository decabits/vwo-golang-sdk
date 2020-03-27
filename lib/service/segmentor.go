package service

import (
	"github.com/decabits/vwo-golang-sdk/lib/constants"
	"github.com/decabits/vwo-golang-sdk/lib/schema"
)

// Evaluate function
func Evaluate(segments []schema.Segment, customVariables schema.Options) bool {
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

	} else if operator == constants.OperandTypesUser {

	}
}

/*
func evaluateCustomVariable(operand, customVariables schema.Options) bool {
	//TO BE COMPLETED
	operandKey, operand := utils.GetKeyValue(operand)
	customVariables[operandKey] = utils.process_custom_variables_value(customVariables[operandKey])
	operandType, operandValue := utils.process_operand_value(operand)
}
*/