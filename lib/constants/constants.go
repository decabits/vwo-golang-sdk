package constants

//constants
const (
	MaxTrafficValue = 10000
	StatusRunning   = "RUNNING"
	SDKVersion      = "1.0.0" //check its defination in constants.py
	SDKName         = "vwo-golang-sdk"
	Platform        = "server"

	CampaignTypeVisualAB       = "VISUAL_AB"
	CampaignTypeFeatureTest    = "FEATURE_TEST"
	CampaignTypeFeatureRollout = "FEATURE_ROLLOUT"

	OperatorTypeAnd = "and"
	OperatorTypeOr  = "or"
	OperatorTypeNot = "not"

	OperandTypesCustomVariable = "custom_variable"
	OperandTypesUser           = "user"

	HTTPSProtocol = "https://"

	EndPointsBaseURL         = "dev.visualwebsiteoptimizer.com"
	EndPointsAccountSettings = "/server-side/settings"
	EndPointsTrackUser       = "/server-side/track-user"
	EndPointsTrackGoal       = "/server-side/track-goal"
	EndPointsPush            = "/server-side/push"
)
