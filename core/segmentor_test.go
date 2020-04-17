package core

import (
	// "encoding/json"
	// "io/ioutil"
	"testing"

	"github.com/decabits/vwo-golang-sdk/schema"
	// "github.com/google/logger"
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

// type TestCase struct {
// 	DSL            map[string]interface{} `json:"dsl"`
// 	Expected       bool                   `json:"expectation"`
// 	CustomVariable map[string]interface{} `json:"custom_variable"`
// }

// func TestSegmentEvaluator(t *testing.T) {
// 	var testData map[string]map[string]TestCase
// 	data, err := ioutil.ReadFile("./testData/segmentTestCase.json")
// 	if err != nil {
// 		logger.Info("Error: " + err.Error())
// 	}

// 	if err = json.Unmarshal(data, &testData); err != nil {
// 		logger.Info("Error: " + err.Error())
// 	}

// 	for parent, v := range testData {
// 		for child, value := range v {
// 			options := schema.Options{
// 				CustomVariables: value.CustomVariable,
// 			}
// 			actual := SegmentEvaluator(value.DSL, options)
// 			expected := value.Expected
// 			assert.Equal(t, expected, actual, parent + " " + child)
// 		}
// 	}

// }
