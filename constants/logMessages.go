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
	//Debug Messages
	DebugMessageDevelopmentMode                 = "[%v] Development mode is : %v"
	DebugMessageSDKInitialized                  = "[%v] SDK properly initialized"
	DebugMessageNoCustomLoggerFound             = "[%v] No custom logger found, using pre-defined google logger "
	DebugMessageCustomLoggerUsed                = "[%v] Custom logger used"
	DebugMessageSegmentationSkipped             = "[%v] For User ID:%v of CampaignKey: %v segments are missing, hence skipping segmentation"
	DebugMessageNoUserStorageServiceGet         = "[%v] No UserStorageService to get stored data"
	DebugMessageGettingStoredVariation          = "[%v] Got stored variation for User ID: %v of Campaign: %v as Variation: %v, found in UserStorageService"
	DebugMessageNoStoredVariation               = "[%v] No stored variation for User ID: %v for Campaign: %v found in UserStorageService"
	DebugMessageWhitelistingSkipped             = "[%v] For User ID: %v of Campaign: %v, whitelisting was skipped"
	DebugMessageNoUserStorageServiceSet         = "[%v] No UserStorageService to set data"
	DebugMessageImpressionForPush               = "[%v] impression built for pushing - %v "
	DebugMessageImpressionForTrackGoal          = "[%v] impression built for track goal - %v "
	DebugMessageImpressionForTrackUser          = "[%v] impression built for track user - %v "
	DebugMessageGotVariationForUser             = "[%v] User ID: %v of Campaign: %v campaignType: %v got variation: %v inside method: %v"
	DebugMessageUUIDForUser                     = "[%v] Uuid generated for User ID: %v and accountId: %v is %v"
	DebugMessageUserHashBucketValue             = "[%v] User ID: %v having hash: %v got bucketValue: %v"
	DebugMessageVariationHashBucketValue        = "[%v] User ID: %v for CampaignKey:%v having percent traffic:%v got bucket value:%v"
	DebugMessageUserNotPartOfCampaign           = "[%v] User ID:%v for CampaignKey:%v type: %v did not become part of campaign method:%v"
	DebugMessageSegmentationSkippedForVariation = "[%v] For User ID:%v of CampaignKey: %v segments are missing, hence skipping segmentation for variation %v"
	DebugMessageSegmentationStatusForVariation  = "[%v] For User ID: %v of Campaign: %v with variables: %v %v %v for variation %v"

	/*Extras*/
	DebugMessageNoWhitelistedVariation = "[%v] No whitelisting variation found in campaign: %v"
	DebugMessageNoTargettedVariation   = "[%v] No targetted variation found : %v"
	DebugMessageCustomLoggerFound      = "[%v] Custom logger found"
	DebugMessageValidConfiguration     = "[%v] SDK configuration and account settings are valid"
	DebugMessageSettingsFileProcessed  = "[%v] Settings file processed"
	DebugMessageNoSegmentsInVariation  = "[%v] For User ID: %v of Campaign: %v, segment was missing, hence skipping segmentation %v"

	//Error Messages
	ErrorMessageCustomLoggerMisconfigured              = "[%v] Custom logger is provided but seems to have misconfigured. Please check the API Docs. Using default logger."
	ErrorMessageSettingsFileCorrupted                  = "[%v] Settings file is corrupted. Please contact VWO Support for help : %v"
	ErrorMessageImpressionFailed                       = "[%v] Impression event could not be sent to VWO endpoint: %v"
	ErrorMessageSetUserStorageServiceFailed            = "[%v] Error while saving data into UserStorage for User ID:%v."
	ErrorMessageActivateAPIMissingParams               = "[%v] activate API got bad parameters. It expects campaignKey(String) as first, User ID(String) as second and options(Optional) as third argument"
	ErrorMessagePushAPIMissingParams                   = "[%v] push API got bad parameters. It expects tagKey(String) as first, tagKey(String) as second and User ID(String) as third argument"
	ErrorMessageTrackAPIMissingParams                  = "[%v] track API got bad parameters. It expects campaignKey(String) as first, User ID(String) as second, goalIdentifier(String) as third argument and options(Optional) as fourth parameter"
	ErrorMessageGetFeatureVariableMissingParams        = "[%v] getFeatureVariableValue API got bad parameters. It expects campaignKey(String) as first, variableKey(String) as second, User ID(String) as third, and options as fourth argument"
	ErrorMessageGetVariationAPIMissingParams           = "[%v] getVariation API got bad parameters. It expects campaignKey(String) as first, User ID(String) as second and options(Optional) as third argument"
	ErrorMessageIsFeatureEnabledAPIMissingParams       = "[%v] isFeatureEnabled API got bad parameters. It expects Campaign(String) as first, User ID(String) as second and options(Optional) as third argument"
	ErrorMessageGetUserStorageServiceFailed            = "[%v] Getting data from UserStorageService failed for User ID: %v"
	ErrorMessageInvalidAPI                             = "[%v] API is not valid for Campaign: %v of type: %v for User ID: %v"
	ErrorMessageTrackAPIGoalNotFound                   = "[%v] Goal: %v not found for Campaign: %v and User ID: %v"
	ErrorMessageTrackAPIRevenueNotPassedForRevenueGoal = "[%v] Revenue value should be passed for revenue goal: %v for Campaign: %v and User ID: %v"
	ErrorMessageTagKeyLengthExceeded                   = "[%v] Length of tagKey: %v for User ID: %v can not be greater than 255"
	ErrorMessageTagValueLengthExceeded                 = "[%v] Length of value: %v of tagKey: %v for User ID: %v can not be greater than 255"
	ErrorMessageCampaignNotRunning                     = "[%v] API Campaign: %v is not RUNNING. Please verify from VWO App"

	/*Extras*/
	ErrorMessageNoVariationAlloted        = "[%v] User ID: %v of CampaignKey: %v type: %v did not get any variation "
	ErrorMessageCampaignNotFound          = "[%v] Campaign key: %v not found"
	ErrorMessageGoalNotFound              = "[%v] Goal: %v not found"
	ErrorMessageNoVariationForBucketValue = "[%v] No variation found for user ID %v in campaignKey: %v having bucket value: %v"
	ErrorMessageInvalidLoggerStorage      = "[%v] Invalid storage object/Logger given. Refer documentation on how to pass custom storage."
	ErrorMessageInvalidAccountID          = "[%v] AccountId is required for fetching account settings. Aborting"
	ErrorMessageInvalidSDKKey             = "[%v] SDKKey is required for fetching account settings. Aborting"
	ErrorMessageInvalidSettingsFile       = "[%v] Settings-file fetched is not proper : %v"
	ErrorMessageCannotProcessSettingsFile = "[%v] Error processing settings file err : %v"
	ErrorMessageNoVariationInCampaign     = "[%v] No variations in campaign: %v"
	ErrorMessageVariationNotFound         = "[%v] Variation : %v not found in campaign : %v "
	ErrorMessageURLNotFound               = "[%v] URL not Found: %v"
	ErrorMessageResponseNotParsed         = "[%v] Error parsing response for URL: %v"
	ErrorMessageCouldNotGetURL            = "[%v] Failed get request for URL: %v"
	ErrorMessageCannotReadSettingsFile    = "[%v] Settings file could not be read and processed. Please contact VWO Support for help : %v"
	ErrorMessageVariableNotFound          = "[%v] Variable: %v for User ID: %v for campaign %v of type %v"

	//Info Messages
	InfoMessageGettingDataUserStorageService  = "[%v] Getting data into UserStorageService for User ID: %v successful"
	InfoMessageGotStoredVariation             = "[%v] Got stored variation:%v of CampaignKey:%v for User ID:%v from UserStorage"
	InfoMessageSettingDataUserStorageService  = "[%v] Setting data into UserStorageService for User ID: %v successful"
	InfoMessageInvalidVariationKey            = "[%v] Variation was not assigned to User ID: %v for Campaign: %v"
	InfoMessageVariationRangeAllocation       = "[%v] Variation: %v with weight: %v got range as: ( %v - %v ))"
	InfoMessageFeatureEnabledForUser          = "[%v] Campaign: %v for user ID: %v is enabled"
	InfoMessageFeatureNotEnabledForUser       = "[%v] Campaign: %v for user ID: %v is not enabled"
	InfoMessageUserEligibilityForCampaign     = "[%v] Is User ID: %v part of campaign ? %v"
	InfoMessageSegmentationStatus             = "[%v] For User ID: %v of Campaign: %v with variables: %v %v %v"
	InfoMessageImpressionSuccess              = "[%v] Impression event - %v was successfully received by VWO having keys: %v"
	InfoMessageForcedvariationAllocated       = "[%v] User ID:%v of CampaignKey:%v type: %v got forced-variation: %v"
	InfoMessageVariationAllocated             = "[%v] User ID: %v of Campaign: %v got variation: %v"
	InfoMessageUserGotNoVariation             = "[%v] User ID: %v for Campaign: %v did not allot any variation : %v"
	InfoMessageMainKeysForPushAPI             = "[%v] Having main keys: AccountID:%v User ID:%v u:%v and tags:%v"
	InfoMessageMainKeysForImpression          = "[%v] Having main keys: AccountID:%v User ID:%v campaignId:%v and VariationID:%v"
	InfoMessageSegmentationStatusForVariation = "[%v] For User ID: %v of Campaign: %v with variables: %v %v %v for variation %v"
	InfoMessageGotVariationForUser            = "[%v] User ID:%v for CampaignKey:%v type: %v got variation_name:%v"

	/*Extras*/
	InfoMessageUserRecievedVariableValue = "[%v] Value for variable: %v of feature flag: %v is: %v for user: %v"
)
