package constants

const (
	MaxTrafficPercent = 100
	MaxTrafficValue   = 10000
	StatusRunning     = "RUNNING"
	SDKVersion        = "1.0.0" //check its defination in constants.py
	SDKName           = "vwo-golang-sdk"
	Platform          = "server"
	SeedValue         = 1

	CampaignTypeVisualAB       = "VISUAL_AB"
	CampaignTypeFeatureTest    = "FEATURE_TEST"
	CampaignTypeFeatureRollout = "FEATURE_ROLLOUT"

	GoalTypeRevenue = "REVENUE_TRACKING"
	GoalTypeCustom  = "CUSTOM_GOAL"

	PushAPITagValueLength = 255
	PushAPITagKeyLength   = 255

	OperatorTypeAnd = "and"
	OperatorTypeOr  = "or"
	OperatorTypeNot = "not"

	OperandTypesCustomVariable = "custom_variable"
	OperandTypesUser           = "user"

	HTTPSProtocol            = "https://"
	EndPointsBaseURL         = "dev.visualwebsiteoptimizer.com"
	EndPointsAccountSettings = "/server-side/settings"
	EndPointsTrackUser       = "/server-side/track-user"
	EndPointsTrackGoal       = "/server-side/track-goal"
	EndPointsPush            = "/server-side/push"

	BaseURL         = "dev.visualwebsiteoptimizer.com"
	AccountSettings = "/server-side/settings"
	TrackUser       = "/server-side/track-user"
	TrackGoal       = "/server-side/track-goal"
	Push            = "/server-side/push"

	Boolean = "boolean"
	Double  = "double"
	Integer = "integer"
	String  = "string"

	LowerMatch    = `^lower\((.*)\)`
	WildcardMatch = `^wildcard\((.*)\)`
	RegexMatch    = `^regex\((.*)\)`
	StartingStar  = `^\*`
	EndingStar    = `\*$`

	LowerValue              = 1
	StartingEndingStarValue = 2
	StartingStarValue       = 3
	EndingStarValue         = 4
	RegexValue              = 5
	EqualValue              = 6
)
