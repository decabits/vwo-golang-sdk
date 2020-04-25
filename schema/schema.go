package schema

import (
	"github.com/google/logger"
)

// VwoInstance struct utils
type VwoInstance struct {
	SettingsFile      SettingsFile
	UserStorage       UserStorage
	Logger            *logger.Logger
	IsDevelopmentMode bool
}

// SettingsFile struct
type SettingsFile struct {
	SDKKey    string     `json:"sdkKey"`
	Campaigns []Campaign `json:"campaigns"`
	AccountID int        `json:"accountId"`
}

// Campaign struct
type Campaign struct {
	ID                int                    `json:"id"`
	Segments          map[string]interface{} `json:"segments"`
	Status            string                 `json:"status"`
	PercentTraffic    int                    `json:"percentTraffic"`
	Goals             []Goal                 `json:"goals"`
	Variations        []Variation            `json:"variations"`
	Variables         []Variable             `json:"variables"`
	IsForcedVariation bool                   `json:"isForcedVariationEnabled"`
	Key               string                 `json:"key"`
	Type              string                 `json:"type"`
}

// Goal struct
type Goal struct {
	Identifier string `json:"identifier"`
	ID         int    `json:"id"`
	Type       string `json:"type"`
}

// Variation struct
type Variation struct {
	ID       int                    `json:"id"`
	Name     string                 `json:"name"`
	Changes  interface{}            `json:"changes"`
	Weight   float64                `json:"weight"`
	Segments map[string]interface{} `json:"segments"`

	Variables        []Variable `json:"variables"`
	IsFeatureEnabled bool       `json:"isFeatureEnabled"`

	StartVariationAllocation int
	EndVariationAllocation   int
}

// Variable struct
type Variable struct {
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
	Key   string      `json:"key"`
	ID    int         `json:"id"`
}

// Options struct
type Options struct {
	CustomVariables             map[string]interface{} `json:"custom_variables"`
	VariationTargetingVariables map[string]interface{} `json:"variation_targeting_variables"`
	RevenueGoal                 int
}

// UserData  struct
type UserData struct {
	UserID        string
	CampaignKey   string
	VariationName string
}

// VariationAllocationRange struct
type VariationAllocationRange struct {
	StartRange int
	EndRange   int
}

// Impression struct
type Impression struct {
	AccountID    int     `json:"account_id"`
	UID          string  `json:"uId"`
	Random       float32 `json:"random"`
	SID          string  `json:"sId"`
	U            string  `json:"u"`
	Sdk          string  `json:"sdk"`
	SdkV         string  `json:"sdk-v"`
	Ap           string  `json:"ap"`
	URL          string  `json:"url"`
	ExperimentID int     `json:"experiment_id"`
	Combination  int     `json:"combination"`
	ED           string  `json:"ed"`
	GoalID       int     `json:"goal_id"`
	R            int     `json:"r"`
}

// Response struct
type Response struct {
	Text       string
	StatusCode int
}

// UserStorage interface
type UserStorage interface {
	Get(userID, campaignKey string) UserData
	Set(string, string, string)
	Exist() bool
}
