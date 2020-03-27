package schema

// "github.com/Piyushhbhutoria/vwo-go-sdk/lib/utils"

// VwoInstance struct utils
type VwoInstance struct {
	// Logger       *utils.Logger
	SettingsFile SettingsFile
}

// SettingsFile struct
type SettingsFile struct {
	Campaigns []Campaign `json:"campaign"`
	AccountID string     `json:"accountId"`
	Version   int        `json:"version"`
}

// Campaign struct
type Campaign struct {
	Goals             []Goal      `json:"goals"`
	Variations        []Variation `json:"variations"`
	ID                int         `json:"id"`
	PercentTraffic    int         `json:"percentTraffic"`
	Key               string      `json:"key"`
	Status            string      `json:"status"`
	Type              string      `json:"type"`
	IsForcedVariation bool
	Segments          []Segment
}

// Segment struct
type Segment struct {
}

// Goal struct
type Goal struct {
	Identifier string `json:"identifier"`
	ID         int    `json:"id"`
	Type       string `json:"type"`
}

// Variation struct
type Variation struct {
	Name                     string `json:"name"`
	ID                       string `json:"id"`
	Weight                   int    `json:"weight"`
	StartVariationAllocation int
	EndVariationAllocation   int
}

// Options struct
type Options struct {
	CustomVariables             []int
	VariationTargetingVariables []int
	RevenueGoal                 int
}

//UserStorageData struct
type UserStorageData struct {
	UserID        string
	CampaignKey   string
	VariationName string
}

// VariationAllocationRange struct
type VariationAllocationRange struct {
	StartRange int
	EndRange   int
}

//Impression struct
type Impression struct {
	Random       float32
	Sdk          string
	SdkV         string
	Ap           string
	SID          int
	U            string
	AccountID    string
	UID          string
	URL          string
	GoalID       int
	ExperimentID string
	Combination  string
	R            int
	//Ed
}

//Response struct
type Response struct {
	Text       string
	StatusCode int
}
