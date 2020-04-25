package constants

// constants for logger
const (
	InfoMessageNoWhitelistedVariation        = "No whitelisting variation found in campaign: %v"
	InfoMessageNoTargettedVariation          = "No targetted variation found : %v"
	InfoMessageGettingDataUserStorageService = "Getting data from UserStorageService for User ID: %v successful"
	InfoMessageNoSegmentsInVariation         = "For userId: %v of Campaign: %v, segment was missing, hence skipping segmentation %v"
	InfoMessageNoVariationAllocated          = "userID: %v of campaign_key: %v type: %v did not get any variation : %v"
	InfoMessageGotVariationForUser           = "User ID: %v for Campaign: %v got variationName: %v inside method: %v"
	InfoMessageSettingDataUserStorageService = "Setting data into UserStorageService for User ID: %v successful"
	InfoMessageVariationRageAllocation       = "Variation: %v with weight: %v got range as: ( %v - %v ))"
	

	DebugMessageUUIDForUser             = "Uuid generated for User ID: %v and accountId: %v is %v"
	DebugMessageImpressionForPush       = "impression built for pushing - %v "
	DebugMessageImpressionForTrackGoal  = "impression built for track goal - %v "
	DebugMessageImpressionForTrackUser  = "impression built for track user - %v "
	DebugMessageNoUserStorageServiceGet = "No UserStorageService to get stored data"
	DebugMessageSegmentationSkipped     = "Segment was missing, hence skipping segmentation for segments: %v and custom variables: %v"
	DebugMessageWhitelistingSkipped     = "variationDecider.go: For userId: %v of Campaign: %v, whitelisting was skipped"
	DebugMessageGettingStoredVariation  = "Got stored variation for User ID: %v of Campaign: %v as Variation: %v, found in UserStorageService"
	DebugMessagesSettingsFileProcessed  = "Settings file processed"
	DebugMessagesDevelopmentMode        = "Development mode is : %v"
	DebugMessageUserHashBucketValue     = "User ID: %v having hash: %v got bucketValue: %v"
	DebugMessageGotVariation            = "User ID: %v of Campaign: %v got variation: %v"

	ErrorMessageGoalNotFound                            = "Goal: %v not found"
	ErrorMessageVariationNotFound                       = "Variation : %v not found in campaign : %v "
	ErrorMessageURLNotFound                             = "URL not Found: %v"
	ErrorMessageResponceNotParsed                       = "Error parsing response for URL: %v"
	ErrorMessageCouldNotGetURL                          = "Failed get request for URL: %v"
	ErrorMessagesInvalidAccountID                       = "AccountId is required for fetching account settings. Aborting"
	ErrorMessagesInvalidSDKKey                          = "SDKKey is required for fetching account settings. Aborting"
	ErrorMessagesSettingsFileCorrupted                  = "Settings file is corrupted. Please contact VWO Support for help : %v"
	ErrorMessagesCannotReadSettingsFile                 = "Settings file could not be readed and processed. Please contact VWO Support for help : %v"
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

	InfoMessageNewVariation                  = "New Variation : %v for user"
	InfoMessageImpressionSuccess          = "Impression event was successfully received by VWO having keys: %v"
	InfoMessageInvalidVariationKey        = "Variation was not assigned to User ID: %v for Campaign: %v"
	InfoMessageUserRecievedVariableValue  = "Value for variable: %v of feature flag: %v is: %v for user: %v"
	InfoMessageFeatureEnabledForUser      = "Campaign: %v for user ID: %v is enabled"
	InfoMessageFeatureNotEnabledForUser   = "Campaign: %v for user ID: %v is not enabled"
	InfoMessageUserEligibilityForCampaign = "Is userID: %v part of campaign ? %v"
	InfoMessageSegmentationStatus         = "For userId: %v of Campaign: %v with variables: %v %v Whitelisting %v"
	InfoMessageUserGotNoVariation         = "User ID: %v for Campaign: %v did not allot any variation"
	InfoMessageVariationAllocated                = "User ID: %v of Campaign: %v got variation: %v"

)
/*
const (
	// Each Message is preceeded by ({file}):
	DebugMessagesCustomLoggerUsed       = "Custom logger used"
	DebugMessagesGettingStoredVariation = "Got stored variation for User ID:{userId} of Campaign:{campaignKey} as Variation:{variationName}, found in UserStorageService"
	DebugMessagesGotVariationForUser    = "User ID:{userId} for Campaign:{campaignKey} got variationName:{variationName} inside method:{method}"
	DebugMessagesIMPRESSIONForPush        = "impression built for pushing - {properties}"
	DebugMessagesIMPRESSIONForTrackGoal   = "impression built for track-goal - {properties}"
	DebugMessagesIMPRESSIONForTrackUser   = "impression built for track-user - {properties}"
	DebugMessagesNoStoredVariation       = "No stored variation for User ID:{userId} for Campaign:{campaignKey} found in UserStorageService"
	DebugMessagesNoUserStorageServiceGet = "No UserStorageService to get stored data"
	DebugMessagesNoUserStorageServiceSet = "No UserStorageService to set data"
	DebugMessagesSDKInitialized          = "SDK properly initialzed"
	DebugMessagesSegmentationSkipped     = "For userId:{userId} of Campaign:{campaignKey}, segment was missing, hence skipping segmentation {variation}"
	DebugMessagesSegmentationStatus      = "For userId:{userId} of Campaign:{campaignKey} with variables:{customVariables} {status} {segmentationType} {variation}"
	DebugMessagesSetDevelopmentMode      = "Development mode is ON"
	DebugMessagesSettingsFileProcessed   = "Settings file processed"
	DebugMessagesUserHashBucketValue      = "User ID:{userId} having hash:{hashValue} got bucketValue:{bucketValue}"
	DebugMessagesUserNotPartOfCampaign = "userId:{userId} for Campaign:{campaignKey} did not become part of campaign, method:{method}"
	DebugMessagesUUIDForUser              = "Uuid generated for User ID:{userId} and accountId:{accountId} is {desiredUuid}"
	DebugMessagesValidConfiguration       = "SDK configuration and account settings are valid"
	DebugMessagesVariationHashBucketValue = "User ID:{userId} for Campaign:{campaignKey} having percent traffic:{percentTraffic} got hash-value:{hashValue} and bucket value:{bucketValue}"
	DebugMessagesWhitelistingSkipped      = "For userId:{userId} of Campaign:{campaignKey}, whitelisting was skipped"

	ErrorMessagesAPIHasCorruptedSettingsFile            = "{api} API has corrupted settings-file. Please check or reach out to VWO support"
	ErrorMessagesActivateAPIMissingParams               = "activate API got bad parameters. It expects campaignKey(String) as first, userId(String) as second and options(optional Object) as third argument"
	ErrorMessagesGetFeatureVariableMissingParams        = "getFeatureVariableValue API got bad parameters. It expects campaignKey(String) as first, variableKey(String) as second, userId(String) as third, and options(optional Object) as fourth argument"
	ErrorMessagesGetVariationAPIMissingParams           = "getVariation API got bad parameters. It expects campaignKey(String) as first, userId(String) as second and options(optional Object) as third argument"
	ErrorMessagesImpressionFailed                       = "Impression event could not be sent to VWO - {endPoint}. Reason: {err}"
	ErrorMessagesIsFeatureEnabledAPIMissingParams       = "isFeatureEnabled API got bad parameters. It expects Campaign(String) as first, userId(String) as second and options(optional Object) as third argument"
	ErrorMessagesGetUserStorageServiceFailed            = "Getting data from UserStorageService failed for User ID:{userId}"
	ErrorMessagesSDKConfigCorrupted                     = "config passed to launch API is not a valid JSON object"
	ErrorMessagesPushInvalidPARAMS                      = "{method} API got bad parameters. It expects tagKey(String) as first, tagValue(String) as second and userId(String) as third argument"
	ErrorMessagesRegexCreationFailed                    = "Regex cound not be processed"
	ErrorMessagesSetUserStorageServiceFailed            = "Saving data into UserStorageService failed for User ID:{userId}"
	ErrorMessagesSegmentationError                      = "Error while segmenting the user:{userId} of Campaign:{campaignKey}{variation} with customVariables:{customVariables}. Error message: {err}"


	ErrorMessagesUnableToCastValueValue                 = "Unable to cast value:{variableValue} to type:{variableType}, returning null"
	ErrorMessagesvariableNotFound                       = "Variable:{variableKey} for User ID:{userId} is not found in settings-file. Returning null"

	InfoMessageGettingDataUserStorageService     = "Getting data from UserStorageService for User ID:{userId} successful"
    InfoMessageNoVariationAllocated               = "userID:{userID} of campaign_key:{campaign_key} type: {campaign_type} did not get any variation"
    InfoMessageGotVariationForUser                = "User_id:{userID} for campaign_key:{campaign_key} type: {campaign_type} got variation_name:{variation_name} inside method:{method}"
    InfoMessageSettingDataUserStorageService     = "Setting data into UserStorageService for User ID:{userId} successful"
	InfoMessageVariableNotUsedReturnDefaultValue = "Variable:{variableKey} is not used in variation:{variationName}. Returning default value"
)
*/
