package schema

// VwoInstance struct utils
type VwoInstance struct {
	// Logger    tils.Logger
	SettingsFile      SettingsFile
	UserStorage       UserStorage
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
	ID      string      `json:"id"`
	Name    string      `json:"name"`
	Changes interface{} `json:"changes"`
	Weight  float64     `json:"weight"`

	Segments         map[string]interface{} `json:"segments"`
	Variables        []Variable             `json:"variables"`
	IsFeatureEnabled bool                   `json:"isFeatureEnabled"`

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
	CustomVariables             []int
	VariationTargetingVariables []int
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
	Random       float32
	Sdk          string
	SdkV         string
	Ap           string
	SID          string
	U            string
	AccountID    int
	UID          string
	URL          string
	GoalID       int
	ExperimentID int
	Combination  string
	R            int
	ED           []byte
}

// Response struct
type Response struct {
	Text       string
	StatusCode int
}

// VariationResponse struct
type VariationResponse struct {
	Variation     Variation
	VariationID   string
	VariationName string
}

// UserStorage struct
type UserStorage interface {
	Get(userID, campaignKey string) (UserData, error)
	Set(userStorageData interface{}) bool
	Exist() bool
}
