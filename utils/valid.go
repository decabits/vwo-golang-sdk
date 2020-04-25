package utils

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
