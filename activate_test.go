package vwo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestActivate(t *testing.T) {
	vwoInstance := VWOInstance{}
	vwoInstance.getInstance("./testData/testdata.json")

	userID := "Varun"
	campaignKey := "notPresent"
	value := vwoInstance.Activate(campaignKey, userID)
	assert.Empty(t, value, "Campaign does not exist")

	userID = "Varun"
	campaignKey = "phpab1"
	value = vwoInstance.Activate(campaignKey, userID)
	assert.Empty(t, value, "Campaign Not running")

	userID = "Liza"
	campaignKey = "php1"
	value = vwoInstance.Activate(campaignKey, userID)
	assert.Empty(t, value, "Campaign Not Valid")

	userID = "Liza"
	campaignKey = "phpab2"
	value = vwoInstance.Activate(campaignKey, userID)
	assert.Empty(t, value, "No Variation in Campaign")

	userID = "Liza"
	campaignKey = "phpab3"
	actual := vwoInstance.Activate(campaignKey, userID)
	assert.NotEmpty(t, actual, "Variation should be found")
}
