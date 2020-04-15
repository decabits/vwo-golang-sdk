package api

import (
	"testing"

	"github.com/decabits/vwo-golang-sdk/utils"
	"github.com/stretchr/testify/assert"
)

func TestActivate(t *testing.T) {
	vwoInstance := utils.GetInstance("../settingsFile.json")

	userID := "Varun"
	campaignKey := "notPresent"
	value := Activate(vwoInstance, campaignKey, userID)
	assert.Empty(t, value, "Campaign does not exist")

	userID = "Varun"
	campaignKey = "phpab5"
	value = Activate(vwoInstance, campaignKey, userID)
	assert.Empty(t, value, "Campaign Not running")

	userID = "Liza"
	campaignKey = "phpab4"
	value = Activate(vwoInstance, campaignKey, userID)
	assert.Empty(t, value, "Campaign Not Valid")

	userID = "Liza"
	campaignKey = "phpab2"
	actual := Activate(vwoInstance, campaignKey, userID)
	expected := "Variation-1"
	assert.Equal(t, expected, actual, "Variation Name does not match")

	userID = "Liza"
	campaignKey = "php1"
	value = Activate(vwoInstance, campaignKey, userID)
	assert.Empty(t, value, "Variation Not found")
}
