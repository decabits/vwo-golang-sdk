package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
	
}