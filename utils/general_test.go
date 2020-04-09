package utils

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/stretchr/testify/assert"
)

var Segments1 = `{
            "or": [
              {
				"custom_variable": {"abcd": "regex(1)"}
              }
			]
		  }`

var Segments2 = `{
	"and": [
              {
                "user": "Varun"
              }
            ]
		  }`

func TestCheckCampaignType(t *testing.T) {
	campaign := schema.Campaign{
		ID:   123,
		Key:  "DemoTest",
		Type: constants.CampaignTypeFeatureRollout,
	}
	campaignType := constants.CampaignTypeFeatureRollout
	assert.True(t, CheckCampaignType(campaign, campaignType), "Campaign should match")

	campaignType = constants.CampaignTypeFeatureTest
	assert.False(t, CheckCampaignType(campaign, campaignType), "Campaign should not match")
}

func TestGetKeyValue(t *testing.T) {
	var segments1 map[string]interface{}
	if err := json.Unmarshal([]byte(Segments1), &segments1); err != nil {
		fmt.Println(err)
	}
	var segments2 map[string]interface{}
	if err := json.Unmarshal([]byte(Segments2), &segments2); err != nil {
		fmt.Println(err)
	}

	actualKey, actualValue := GetKeyValue(segments1)
	expectedKey := "or"
	assert.Equal(t, expectedKey, actualKey, "Expected and Actual Keys should be same")
	//assert.Equal(t, expectedValue, actualValue, "Expected and Actual Values should be same")
	var Temp []interface{}
	assert.IsType(t, Temp, actualValue, "Type Mismatch")

	actualKey, actualValue = GetKeyValue(segments2)
	expectedKey = "and"
	assert.Equal(t, expectedKey, actualKey, "Expected and Actual Keys should be same")
	//assert.Equal(t, expectedValue, actualValue, "Expected and Actual Values should be same")
	assert.IsType(t, Temp, actualValue, "Type Mismatch")
}
