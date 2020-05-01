/*
   Copyright 2019-2020 Wingify Software Pvt. Ltd.

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

package constants

// constants for logger
const (
	DebugMessageNoCustomLoggerFound      = "No custom logger found, using pre-defined logger "
	DebugMessageCustomLoggerFound        = "Custom logger found"
	DebugMessagesSDKInitialized          = "SDK properly initialized"
	DebugMessageNoWhitelistedVariation   = "No whitelisting variation found in campaign: %v"
	DebugMessageNoTargettedVariation     = "No targetted variation found : %v"
	DebugMessageUUIDForUser              = "Uuid generated for User ID: %v and accountId: %v is %v"
	DebugMessageImpressionForPush        = "impression built for pushing - %v "
	DebugMessageImpressionForTrackGoal   = "impression built for track goal - %v "
	DebugMessageImpressionForTrackUser   = "impression built for track user - %v "
	DebugMessageNoUserStorageServiceGet  = "No UserStorageService to get stored data"
	DebugMessageSegmentationSkipped      = "Segment was missing, hence skipping segmentation for segments: %v and custom variables: %v"
	DebugMessageWhitelistingSkipped      = "variationDecider.go: For userId: %v of Campaign: %v, whitelisting was skipped"
	DebugMessageGettingStoredVariation   = "Got stored variation for User ID: %v of Campaign: %v as Variation: %v, found in UserStorageService"
	DebugMessagesSettingsFileProcessed   = "Settings file processed"
	DebugMessagesDevelopmentMode         = "Development mode is : %v "
	DebugMessageUserHashBucketValue      = "User ID: %v having hash: %v got bucketValue: %v"
	DebugMessageGotVariation             = "User ID: %v of Campaign: %v got variation: %v"
	DebugMessageNoSegmentsInVariation    = "For userId: %v of Campaign: %v, segment was missing, hence skipping segmentation %v"
	DebugMessageNoVariationAllocated     = "userID: %v of campaign_key: %v type: %v did not get any variation : %v"
	DebugMessagesNoUserStorageServiceSet = "No UserStorageService to set data"
	DebugMessagesValidConfiguration      = "SDK configuration and account settings are valid"
	DebugMessagesNoStoredVariation       = "No stored variation for User ID: %v for Campaign: %v found in UserStorageService"

	ErrorMessageGoalNotFound                            = "Goal: %v not found"
	ErrorMessageVariationNotFound                       = "Variation : %v not found in campaign : %v "
	ErrorMessageURLNotFound                             = "URL not Found: %v"
	ErrorMessageResponceNotParsed                       = "Error parsing response for URL: %v"
	ErrorMessageCouldNotGetURL                          = "Failed get request for URL: %v"
	ErrorMessagesInvalidAccountID                       = "AccountId is required for fetching account settings. Aborting"
	ErrorMessagesInvalidSDKKey                          = "SDKKey is required for fetching account settings. Aborting"
	ErrorMessagesSettingsFileCorrupted                  = "Settings file is corrupted. Please contact VWO Support for help : %v"
	ErrorMessagesCannotReadSettingsFile                 = "Settings file could not be read and processed. Please contact VWO Support for help : %v"
	ErrorMessagesInvalidSettingsFile                    = "Settings-file fetched is not proper : %v"
	ErrorMessagesCampaignNotRunning                     = "API used: %v - Campaign: %v is not RUNNING. Please verify from VWO App"
	ErrorMessagesInvalidAPI                             = "%v API is not valid for Campaign: %v of type: %v for User ID: %v"
	ErrorMessageCampaignNotFound                        = "Campaign key: %v not found"
	ErrorMessagesvariableNotFound                       = "Variable: %v for User ID: %vis not found in settings-file. Returning nil"
	ErrorMessagesTagKeyLengthExceeded                   = "Length of tagKey: %v for userID: %v can not be greater than 255"
	ErrorMessagesTagValueLengthExceeded                 = "Length of value: %v of tagKey: %v for userID: %v can not be greater than 255"
	ErrorMessagesTrackAPIGoalNotFound                   = "Goal: %v not found for Campaign: %v and userId: %v"
	ErrorMessagesTrackAPIRevenueNotPassedForRevenueGoal = "Revenue value should be passed for revenue goal: %v for Campaign: %v and userId: %v"
	ErrorMessageNoVariationForBucketValue               = "No variation found for bucket value: %v"
	ErrorMessageNoVariationInCampaign                   = "No variations in campaign: %v"
	ErrorMessageNoVariationAlloted                      = "userID: %v of campaign_key: %v type: %v did not get any variation"
	ErrorMessagesImpressionFailed                       = "Impression event could not be sent to VWO Reason: %v"
	ErrorMessagesActivateAPIMissingParams               = "activate API got bad parameters. It expects campaignKey(String) as first, userId(String) as second and options(Optional) as third argument"
	ErrorMessagesGetFeatureVariableMissingParams        = "getFeatureVariableValue API got bad parameters. It expects campaignKey(String) as first, variableKey(String) as second, userId(String) as third, and options as fourth argument"
	ErrorMessagesGetVariationAPIMissingParams           = "getVariation API got bad parameters. It expects campaignKey(String) as first, userId(String) as second and options(Optional) as third argument"
	ErrorMessagesIsFeatureEnabledAPIMissingParams       = "isFeatureEnabled API got bad parameters. It expects Campaign(String) as first, userId(String) as second and options(Optional) as third argument"
	ErrorMessagesPushAPIMissingParams                   = "push API got bad parameters. It expects tagKey(String) as first, tagKey(String) as second and userID(String) as third argument"
	ErrorMessagesTrackAPIMissingParams                  = "track API got bad parameters. It expects campaignKey(String) as first, userId(String) as second, goalIdentifier(String) as third argument and options(Optional) as fourth parameter"
	ErrorMessagesGetUserStorageServiceFailed            = "Getting data from UserStorageService failed for User ID:{userId}"
	ErrorMessageInvalidLoggerStorage                    = "Invalid storage object/Logger given. Refer documentation on how to pass custom storage."
	ErrorMessageCannotProcessSettingsFile               = "Error processing settings file err: "

	InfoMessageNewVariation                  = "New Variation : %v for user"
	InfoMessageImpressionSuccess             = "Impression event was successfully received by VWO having keys: %v"
	InfoMessageInvalidVariationKey           = "Variation was not assigned to User ID: %v for Campaign: %v"
	InfoMessageUserRecievedVariableValue     = "Value for variable: %v of feature flag: %v is: %v for user: %v"
	InfoMessageFeatureEnabledForUser         = "Campaign: %v for user ID: %v is enabled"
	InfoMessageFeatureNotEnabledForUser      = "Campaign: %v for user ID: %v is not enabled"
	InfoMessageUserEligibilityForCampaign    = "Is userID: %v part of campaign ? %v"
	InfoMessageSegmentationStatus            = "For userId: %v of Campaign: %v with variables: %v %v Whitelisting %v"
	InfoMessageUserGotNoVariation            = "User ID: %v for Campaign: %v did not allot any variation"
	InfoMessageVariationAllocated            = "User ID: %v of Campaign: %v got variation: %v"
	InfoMessageVariationRageAllocation       = "Variation: %v with weight: %v got range as: ( %v - %v ))"
	InfoMessageSettingDataUserStorageService = "Setting data into UserStorageService for User ID: %v successful"
)