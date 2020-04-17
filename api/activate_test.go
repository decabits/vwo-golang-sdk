package api

import (
	"github.com/decabits/vwo-golang-sdk/schema"
	"testing"

	"github.com/decabits/vwo-golang-sdk/utils"
	"github.com/stretchr/testify/assert"
)

func TestActivate(t *testing.T) {
	vwoInstance := utils.GetInstance("../settingsFiles/settings1.json")

	userID := "Varun"
	campaignKey := "notPresent"
	value := Activate(vwoInstance, campaignKey, userID)
	assert.Empty(t, value, "Campaign does not exist")

	userID = "Varun"
	campaignKey = "phpab1"
	value = Activate(vwoInstance, campaignKey, userID)
	assert.Empty(t, value, "Campaign Not running")

	userID = "Liza"
	campaignKey = "php1"
	value = Activate(vwoInstance, campaignKey, userID)
	assert.Empty(t, value, "Campaign Not Valid")

	userID = "Liza"
	campaignKey = "phpab2"
	actual := Activate(vwoInstance, campaignKey, userID)
	assert.Empty(t, actual, "No Variation in Campaign")

	userID = "Liza"
	campaignKey = "phpab3"
	actual = Activate(vwoInstance, campaignKey, userID)
	expected := vwoInstance.SettingsFile.Campaigns[3].Variations[1].Name
	assert.Equal(t, expected, actual, "Variation should be found")

	userID = "Gimmy"
	campaignKey = "phpab3"
	options := schema.Options{}
	actual = ActivateWithOptions(vwoInstance, campaignKey, userID, options)
	expected = vwoInstance.SettingsFile.Campaigns[3].Variations[2].Name
	assert.Equal(t, expected, actual, "Variation Not found with options")
}
