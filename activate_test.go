package vwo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestActivate(t *testing.T) {
	assertOutput := assert.New(t)

	vwoInstance := VWOInstance{}
	err := vwoInstance.getInstance("./testdata/testdata.json")
	assertOutput.Nil(err, "error fetching instance")

	userID := ""
	campaignKey := ""
	value := vwoInstance.Activate(campaignKey, userID, nil)
	assertOutput.Empty(value, "Invalid params")

	userID = "Varun"
	campaignKey = "notPresent"
	value = vwoInstance.Activate(campaignKey, userID, nil)
	assertOutput.Empty(value, "Campaign does not exist")

	userID = "Varun"
	campaignKey = "phpab1"
	value = vwoInstance.Activate(campaignKey, userID, nil)
	assertOutput.Empty(value, "Campaign Not running")

	userID = "Liza"
	campaignKey = "php1"
	value = vwoInstance.Activate(campaignKey, userID, nil)
	assertOutput.Empty(value, "Campaign Not Valid")

	userID = "Liza"
	campaignKey = "phpab2"
	value = vwoInstance.Activate(campaignKey, userID, nil)
	assertOutput.Empty(value, "No Variation in Campaign")

	userID = "Liza"
	campaignKey = "phpab3"
	actual := vwoInstance.Activate(campaignKey, userID, nil)
	assertOutput.NotEmpty(actual, "Variation should be found")
}
