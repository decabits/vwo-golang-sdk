package api

import (
	"testing"

	"github.com/decabits/vwo-golang-sdk/schema"

	"github.com/decabits/vwo-golang-sdk/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetVariationName(t *testing.T) {
	vwoInstance := utils.GetInstance("../settingsFiles/settings1.json")
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
	expected := "Variation-1"
	assert.Equal(t, expected, actual, "Variation Name does not match")
}
