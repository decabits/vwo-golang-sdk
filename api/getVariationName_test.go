package api

import (
	"testing"

	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/stretchr/testify/assert"
)

func TestGetVariationName(t *testing.T) {
	vwoInstance := GetInstance("./testData/settings1.json")
	options := schema.Options{}

	userID := "Varun"
	campaignKey := "notPresent"
	value := GetVariationName(vwoInstance, campaignKey, userID, options)
	assert.Empty(t, value, "Campaign does not exist")

	userID = "Varun"
	campaignKey = "phpab1"
	value = GetVariationName(vwoInstance, campaignKey, userID, options)
	assert.Empty(t, value, "Campaign Not running")

	userID = "Liza"
	campaignKey = "php1"
	value = GetVariationName(vwoInstance, campaignKey, userID, options)
	assert.Empty(t, value, "Campaign Not Valid")

	userID = "Liza"
	campaignKey = "phpab2"
	value = GetVariationName(vwoInstance, campaignKey, userID, options)
	assert.Empty(t, value, "Variation Not found")

	userID = "Liza"
	campaignKey = "phpab3"
	actual := GetVariationName(vwoInstance, campaignKey, userID, options)
	assert.NotEmpty(t, actual, "Variation Name does not match")
}
