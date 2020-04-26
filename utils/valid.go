package utils

import (
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/google/logger"
)

// ParseOptions - parses custom options
func ParseOptions(option interface{}) (options schema.Options) {
	if option == nil {
		return
	}
	optionMap, okMap := option.(map[string]interface{})
	if okMap {
		customVariables, okCustomVariables := optionMap["customVariables"]
		if okCustomVariables {
			options.CustomVariables = customVariables.(map[string]interface{})
		}
		variationTargetingVariables, okVariationTargetingVariables := optionMap["variationTargetingVariables"]
		if okVariationTargetingVariables {
			options.VariationTargetingVariables = variationTargetingVariables.(map[string]interface{})
		}
		revenueGoal, okRevenueGoal := optionMap["revenueGoal"]
		if okRevenueGoal {
			options.RevenueGoal = revenueGoal.(int)
		}
	}
	return
}

// ValidateLogger - validates Custom logger
func ValidateLogger(logs interface{}) bool {
	_, ok := logs.(interface {
		CustomLog(a, b string)
	})
	if !ok {
		_, ok = logs.(*logger.Logger)
	}
	return ok
}

// ValidateStorage - validates Custom Storage
func ValidateStorage(storage interface{}) bool {
	_, okGet := storage.(interface {
		Get(a, b string) schema.UserData
	})
	_, okSet := storage.(interface{ Set(a, b, c string) })
	if (okGet && okSet) || storage == nil {
		return true
	}
	return false
}

// ValidateActivate - validates Activate API parameters
func ValidateActivate(campaignKey, userID string) bool {
	if campaignKey == "" || userID == "" {
		return false
	}
	return true
}

// ValidateGetFeatureVariableValue - validates GetFeatureVariableValue API parameters
func ValidateGetFeatureVariableValue(campaignKey, variableKey, userID string) bool {
	if campaignKey == "" || userID == "" || variableKey == "" {
		return false
	}
	return true
}

// ValidateGetVariationName - validates GetVariationName API parameters
func ValidateGetVariationName(campaignKey, userID string) bool {
	if campaignKey == "" || userID == "" {
		return false
	}
	return true
}

// ValidateIsFeatureEnabled - validates IsFeatureEnabled API parameters
func ValidateIsFeatureEnabled(campaignKey, userID string) bool {
	if campaignKey == "" || userID == "" {
		return false
	}
	return true
}

// ValidatePush - validates Push API parameters
func ValidatePush(tagKey, tagValue, userID string) bool {
	if tagKey == "" || tagValue == "" || userID == "" {
		return false
	}
	return true
}

// ValidateTrack - validates Track API parameters
func ValidateTrack(campaignKey, userID, goalIdentifier string) bool {
	if campaignKey == "" || userID == "" || goalIdentifier == "" {
		return false
	}
	return true
}
