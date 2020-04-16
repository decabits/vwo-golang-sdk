package core

import (
	"testing"

	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/stretchr/testify/assert"
)

func TestEvaluate(t *testing.T) {
	options := schema.Options{}

	operator := "and"
	res := []bool{true, true, true, true}
	actual := evaluate(operator, res, options)
	assert.True(t, actual, "Value expected to be true")

	res = []bool{true, false, true, true}
	actual = evaluate(operator, res, options)
	assert.False(t, actual, "Value expected to be true")

	operator = "or"
	res = []bool{false, false, false, false}
	actual = evaluate(operator, res, options)
	assert.False(t, actual, "Value expected to be true")

	res = []bool{false, false, false, true}
	actual = evaluate(operator, res, options)
	assert.True(t, actual, "Value expected to be true")

	operator = "not"
	res = []bool{}
	actual = evaluate(operator, res, options)
	assert.False(t, actual, "Value expected to be true")

}

func TestOperandUserParser(t *testing.T) {
	options := schema.Options{
		VWOUserID: "Robert",
	}
	operand := "Robert, Lizzie, Chris, Gimmy, Matt"
	value := operandUserParser(operand, options)
	assert.True(t, value, "User expected to be present")

	options = schema.Options{
		VWOUserID: "Robert",
	}
	operand = "Lizzie, Chris, Gimmy, Matt"
	value = operandUserParser(operand, options)
	assert.False(t, value, "User expected to be absent")
}

func TestExtractResult(t *testing.T) {
	operandType := 1
	operandValue := "tempval"
	tagValue := "TempVal"
	value := extractResult(operandType, operandValue, tagValue)
	assert.True(t, value, "Incorrect check")

	operandValue = "tempval"
	tagValue = "TempVool"
	value = extractResult(operandType, operandValue, tagValue)
	assert.False(t, value, "Incorrect check")

	operandType = 2
	operandValue = "tempval"
	tagValue = "pval"
	value = extractResult(operandType, operandValue, tagValue)
	assert.True(t, value, "Incorrect check")

	operandValue = "tempval"
	tagValue = "pVal"
	value = extractResult(operandType, operandValue, tagValue)
	assert.False(t, value, "Incorrect check")

	operandType = 3
	operandValue = "tempval"
	tagValue = "val"
	value = extractResult(operandType, operandValue, tagValue)
	assert.True(t, value, "Incorrect check")

	operandValue = "tempval"
	tagValue = "Val"
	value = extractResult(operandType, operandValue, tagValue)
	assert.False(t, value, "Incorrect check")

	operandType = 4
	operandValue = "tempval"
	tagValue = "tem"
	value = extractResult(operandType, operandValue, tagValue)
	assert.True(t, value, "Incorrect check")

	operandValue = "tempval"
	tagValue = "temv"
	value = extractResult(operandType, operandValue, tagValue)
	assert.False(t, value, "Incorrect check")

	operandType = 5
	operandValue = "tempval"
	tagValue = "mp"
	value = extractResult(operandType, operandValue, tagValue)
	assert.True(t, value, "Incorrect check")

	operandValue = "tempval"
	tagValue = "MP"
	value = extractResult(operandType, operandValue, tagValue)
	assert.False(t, value, "Incorrect check")

	operandType = 6
	operandValue = "val1"
	tagValue = "val1"
	value = extractResult(operandType, operandValue, tagValue)
	assert.True(t, value, "Incorrect check")

	operandValue = "val1"
	tagValue = "val"
	value = extractResult(operandType, operandValue, tagValue)
	assert.False(t, value, "Incorrect check")
}

func TestEvaluateCustomVariables(t *testing.T) {

}
