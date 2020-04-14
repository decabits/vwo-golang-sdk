package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessCustomVariablesValue(t *testing.T) {
	var value interface{}
	value = true
	actual := ProcessCustomVariablesValue(value)
	expected := "true"
	assert.Equal(t, expected, actual, "String mismatch")

	value = "Value"
	actual = ProcessCustomVariablesValue(value)
	expected = "Value"
	assert.Equal(t, expected, actual, "String mismatch")
}

func TestMatchWithRegex(t *testing.T) {
	operand := "str"
	regex := `^str`
	value := matchWithRegex(operand, regex)
	assert.True(t, value, "TRUE")

	operand = "str"
	regex = `^string`
	value = matchWithRegex(operand, regex)
	assert.False(t, value, "False")
}

func TestExtractOperandValue(t *testing.T) {
	
}
