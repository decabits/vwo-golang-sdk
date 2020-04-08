package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/google/logger"
)

// SettingsFile JSON temp
var SettingsFile = `{
	 "sdkKey": "7aeed7f67f5a0b0fbe476c1f086a7038",
	 "accountId": 89499,
	 "campaigns": [{
      "id": 283,
      "status": "RUNNING",
      "percentTraffic": 100,
      "goals": [
        {
          "identifier": "rev",
          "id": 2,
          "type": "REVENUE_TRACKING"
        },
        {
          "identifier": "custom",
          "id": 281,
          "type": "CUSTOM_GOAL"
        }
	  ],
	  "variables": [
        {
          "value": 30,
          "type": "integer",
          "key": "int2",
          "id": 1
        },
        {
          "value": 10.67,
          "type": "double",
          "key": "float2",
          "id": 2
        }
      ],
      "isForcedVariationEnabled": true,
      "key": "php2",
      "variations": [
        {
          "changes": {},
          "id": 1,
          "variables": [
            {
              "value": "abcd",
              "type": "string",
              "key": "string2",
              "id": 3
            },
            {
              "value": false,
              "type": "boolean",
              "key": "bool2",
              "id": 4
            }
          ],
          "isFeatureEnabled": false,
          "weight": 33.3333,
          "name": "Control"
        },
        {
          "segments": {
            "or": [
              {
                "custom_variable": {
                  "abcd": "regex(1)"
                }
              }
            ]
          },
          "changes": {},
          "id": 2,
          "variables": [
            {
              "value": 301,
              "type": "integer",
              "key": "int2",
              "id": 1
            },
            {
              "value": 10.671,
              "type": "double",
              "key": "float2",
              "id": 2
            },
            {
              "value": "abcde",
              "type": "string",
              "key": "string2",
              "id": 3
            },
            {
              "value": true,
              "type": "boolean",
              "key": "bool2",
              "id": 4
            }
          ],
          "isFeatureEnabled": true,
          "weight": 33.3333,
          "name": "Variation-1"
        },
        {
          "segments": {},
          "changes": {},
          "id": 3,
          "variables": [
            {
              "value": 302,
              "type": "integer",
              "key": "int2",
              "id": 1
            },
            {
              "value": 10.672,
              "type": "double",
              "key": "float2",
              "id": 2
            },
            {
              "value": "abcdef",
              "type": "string",
              "key": "string2",
              "id": 3
            },
            {
              "value": false,
              "type": "boolean",
              "key": "bool2",
              "id": 4
            }
          ],
          "isFeatureEnabled": true,
          "weight": 33.3333,
          "name": "Variation-2"
        }
      ],
      "type": "FEATURE_ROLLOUT"
    }]
 }`

// Campaign1 Feature Rollout Campaign
var Campaign1 = `{
      "id": 283,
      "status": "RUNNING",
      "percentTraffic": 100,
      "goals": [
        {
          "identifier": "rev",
          "id": 2,
          "type": "REVENUE_TRACKING"
        },
        {
          "identifier": "custom",
          "id": 281,
          "type": "CUSTOM_GOAL"
        }
	  ],
	  "variables": [
        {
          "value": 30,
          "type": "integer",
          "key": "int2",
          "id": 1
        },
        {
          "value": 10.67,
          "type": "double",
          "key": "float2",
          "id": 2
        }
      ],
      "isForcedVariationEnabled": true,
      "key": "php2",
      "variations": [
        {
          "changes": {},
          "id": 1,
          "variables": [
            {
              "value": "abcd",
              "type": "string",
              "key": "string2",
              "id": 3
            },
            {
              "value": false,
              "type": "boolean",
              "key": "bool2",
              "id": 4
            }
          ],
          "isFeatureEnabled": false,
          "weight": 33.3333,
          "name": "Control"
        },
        {
          "segments": {
            "or": [
              {
                "custom_variable": {
                  "abcd": "regex(1)"
                }
              }
            ]
          },
          "changes": {},
          "id": 2,
          "variables": [
            {
              "value": 301,
              "type": "integer",
              "key": "int2",
              "id": 1
            },
            {
              "value": 10.671,
              "type": "double",
              "key": "float2",
              "id": 2
            },
            {
              "value": "abcde",
              "type": "string",
              "key": "string2",
              "id": 3
            },
            {
              "value": true,
              "type": "boolean",
              "key": "bool2",
              "id": 4
            }
          ],
          "isFeatureEnabled": true,
          "weight": 33.3333,
          "name": "Variation-1"
        },
        {
          "segments": {},
          "changes": {},
          "id": 3,
          "variables": [
            {
              "value": 302,
              "type": "integer",
              "key": "int2",
              "id": 1
            },
            {
              "value": 10.672,
              "type": "double",
              "key": "float2",
              "id": 2
            },
            {
              "value": "abcdef",
              "type": "string",
              "key": "string2",
              "id": 3
            },
            {
              "value": false,
              "type": "boolean",
              "key": "bool2",
              "id": 4
            }
          ],
          "isFeatureEnabled": true,
          "weight": 33.3333,
          "name": "Variation-2"
        }
      ],
      "type": "FEATURE_ROLLOUT"
    }`

// Campaign2 Feature Test Campaign
var Campaign2 = `{
      "id": 283,
      "status": "RUNNING",
      "percentTraffic": 100,
      "goals": [
        {
          "identifier": "rev",
          "id": 2,
          "type": "REVENUE_TRACKING"
        },
        {
          "identifier": "custom",
          "id": 281,
          "type": "CUSTOM_GOAL"
        }
	  ],
	  "variables": [
        {
          "value": 30,
          "type": "integer",
          "key": "int2",
          "id": 1
        },
        {
          "value": 10.67,
          "type": "double",
          "key": "float2",
          "id": 2
        }
      ],
      "isForcedVariationEnabled": true,
      "key": "php2",
      "variations": [
        {
          "changes": {},
          "id": 1,
          "variables": [
            {
              "value": "abcd",
              "type": "string",
              "key": "string2",
              "id": 3
            },
            {
              "value": false,
              "type": "boolean",
              "key": "bool2",
              "id": 4
            }
          ],
          "isFeatureEnabled": false,
          "weight": 33.3333,
          "name": "Control"
        },
        {
          "segments": {
            "or": [
              {
                "custom_variable": {
                  "abcd": "regex(1)"
                }
              }
            ]
          },
          "changes": {},
          "id": 2,
          "variables": [
            {
              "value": 301,
              "type": "integer",
              "key": "int2",
              "id": 1
            },
            {
              "value": 10.671,
              "type": "double",
              "key": "float2",
              "id": 2
            },
            {
              "value": "abcde",
              "type": "string",
              "key": "string2",
              "id": 3
            },
            {
              "value": true,
              "type": "boolean",
              "key": "bool2",
              "id": 4
            }
          ],
          "isFeatureEnabled": true,
          "weight": 33.3333,
          "name": "Variation-1"
        },
        {
          "segments": {},
          "changes": {},
          "id": 3,
          "variables": [
            {
              "value": 302,
              "type": "integer",
              "key": "int2",
              "id": 1
            },
            {
              "value": 10.672,
              "type": "double",
              "key": "float2",
              "id": 2
            },
            {
              "value": "abcdef",
              "type": "string",
              "key": "string2",
              "id": 3
            },
            {
              "value": false,
              "type": "boolean",
              "key": "bool2",
              "id": 4
            }
          ],
          "isFeatureEnabled": false,
          "weight": 33.3333,
          "name": "Variation-2"
        }
      ],
      "type": "FEATURE_TEST"
	}`

// GetTempInstance function gets a temporary vwo instance for testing
func GetTempInstance() schema.VwoInstance {
	var settingsFile schema.SettingsFile
	if err := json.Unmarshal([]byte(SettingsFile), &settingsFile); err != nil {
		fmt.Println(err)
	}

	logger := logger.Init(constants.SDKName, true, false, ioutil.Discard)
	defer logger.Close()

	vwoInstance := schema.VwoInstance{
		SettingsFile: settingsFile,
		Logger:       logger,
	}

	return vwoInstance
}
