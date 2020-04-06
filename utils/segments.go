package utils

import (
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/decabits/vwo-golang-sdk/constants"
)

// ProcessCustomVariablesValue function
func ProcessCustomVariablesValue(value interface{}) string {
	switch value.(type) {
	// handle cases
	case bool:
		return strconv.FormatBool(value.(bool))
	}
	return value.(string)
}

// PreProcessOperandValue function
func PreProcessOperandValue(operand interface{}) (operandType int, operandValue string) {
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
		operandValue = strings.Replace(strings.Replace(operandValue, constants.StartingStar, "", -1), constants.EndingStar, "", -1)
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
func ProcessValues(operandValue string, tagValue interface{}) (string, string) {
	processedOperandValue, err := strconv.ParseFloat(operandValue, 10)
	if err != nil {
		return operandValue, tagValue.(string)
	}
	processedTagValue, err := strconv.ParseFloat(tagValue.(string), 10)
	if err != nil {
		return operandValue, tagValue.(string)
	}
	// now we have surity that both are numbers
	// now we can convert them independently to int type if they
	// are int rather than floats
	var newProcessedOperandValue, newProcessedTagValue int
	if processedOperandValue == math.Floor(processedOperandValue) {
		newProcessedOperandValue = int(processedOperandValue)
	}
	if processedTagValue == math.Floor(processedTagValue) {
		newProcessedTagValue = int(processedTagValue)
	}
	// convert it back to string and return
	return strconv.Itoa(newProcessedOperandValue), strconv.Itoa(newProcessedTagValue)
}

func matchWithRegex(operand, regex string) bool {
	var re = regexp.MustCompile(regex)
	if len(re.FindStringIndex(operand)) > 0 {
		return true
	}
	return false
}

func extractOperandValue(operand, regex string) string {
	re := regexp.MustCompile(regex)
	submatchall := re.FindStringSubmatch(operand)
	if len(submatchall) > 0 {
		return submatchall[1]
	}
	return ""
}
