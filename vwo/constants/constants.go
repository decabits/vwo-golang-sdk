package vwo

// import(
// 	"os"
// )

const(
	ConstantAPIVersion = 1
	ConstantPlatform = "server"
	ConstantSeedValue = 1
	ConstantMaxTrafficPercent = 100
	ConstantMaxTrafficValue = 10000
	ConstantStatusRunning = "RUNNING"
	//rubocop:disable Style/ExpandPathArguments
	// ConstantLibraryPath = os.path.normpath(os.getcwd() + os.sep + os.pardir)
	//rubocop:enable Style/ExpandPathArguments
	ConstatntsHTTPProtocol = "http://"
	ConstantHTTPSProtocol = "https://"
	ConstantURLNamespace = ""
	ConstantSDKVersion = "1.0.0"
	ConstantSDKName = "vwo-golang-sdk"
	EndpointsBaseURL = "ev.visualwebsiteoptimizer.com"
	EndpointsAccountSettings = "/server-side/settings"
	EndpointsTrackUser = "/server-side/track-user"
	EndpointsTrackGoal = "/server-side/track-goal"
	EndpointsPush = "/server-side/push"
	EventsTrackUser = "track-user"
	EventsTrackGoal = "track-goal"
	DataTypesNumber = "number"
    DataTypesString = "string"
    DataTypesFunction = "function"
    DataTypesBoolean = "boolean"
	GoalTypesRevenue = "REVENUE_TRACKING"
	GoalTypesCustom = "CUSTOM_GOAL"
	VariableTypesInteger = "integer"
    VariableTypesString = "string"
    VariableTypesDouble = "double"
	VariableTypesBoolean = "boolean"
	// GoVariableTypes = {
	// 	'string' => [string],
	// 	'integer' => [int32],
	// 	'double' => [float32],
	// 	'boolean' => [bool]
	//   }
	APIMethodsActivate = "activate"
	APIMethodsGetVariationName = "get_variation_name"
	APIMethodsTrack = "track"
	APIMethodsisFeatureenabled = "is_feature_enabled"
	APIMethodsgetFeatureVariableValue = "get_feature_variable_value"
	APIMethodsPush = "push"
	PushAPITagValueLength = 255
	PushAPITagKeyLength = 255
	CampaignTypesVisualAB = "VISUAL_AB"
	CampaignTypesFeatureTest = "FEATURE_TEST"
	CampaignTypesFeatureRollout = "FEATURE_ROLLOUT"
)

