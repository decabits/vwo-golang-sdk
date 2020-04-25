package vwo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVariationName(t *testing.T) {
	vwoInstance := VWOInstance{}
	vwoInstance.getInstance("./testData/testdata.json")

	userID := "Varun"
	campaignKey := "notPresent"
	value := vwoInstance.GetVariationName(campaignKey, userID)
	assert.Empty(t, value, "Campaign does not exist")

	userID = "Varun"
	campaignKey = "phpab1"
	value = vwoInstance.GetVariationName(campaignKey, userID)
	assert.Empty(t, value, "Campaign Not running")

	userID = "Liza"
	campaignKey = "php1"
	value = vwoInstance.GetVariationName(campaignKey, userID)
	assert.Empty(t, value, "Campaign Not Valid")

	userID = "Liza"
	campaignKey = "phpab2"
	value = vwoInstance.GetVariationName(campaignKey, userID)
	assert.Empty(t, value, "Variation Not found")

	userID = "Liza"
	campaignKey = "phpab3"
	actual := vwoInstance.GetVariationName(campaignKey, userID)
	assert.NotEmpty(t, actual, "Variation Name does not match")
}
