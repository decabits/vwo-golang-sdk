/*
   Copyright 2020 Wingify Software Pvt. Ltd.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package utils

import (
	"github.com/decabits/vwo-golang-sdk/pkg/logger"
	"github.com/decabits/vwo-golang-sdk/pkg/schema"
)

// ParseOptions - parses custom options
func ParseOptions(option interface{}) (options schema.Options) {
	if option == nil {
		options.CustomVariables = make(map[string]interface{})
		options.VariationTargetingVariables = make(map[string]interface{})
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
		revenueValue, okRevenueValue := optionMap["revenueValue"]
		if okRevenueValue {
			options.RevenueValue = revenueValue
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
	if storage == nil {
		return true
	}
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
