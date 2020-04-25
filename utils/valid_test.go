package utils

import (
	"io/ioutil"
	"testing"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/google/logger"
	"github.com/stretchr/testify/assert"
)

type CLog interface {
	CustomLog(a, b string)
}
type CLogS struct{}

func (c *CLogS) CustomLog(a, b string) {}

type WLog interface {
	CustomLogger(a, b string)
}
type WLogS struct{}

func (w *WLogS) CustomLogger(a, b string) {}

func TestValidateLogger(t *testing.T) {
	logs := logger.Init(constants.SDKName, true, false, ioutil.Discard)
	actual := ValidateLogger(logs)
	assert.True(t, actual, "google logger not validated")

	correctLog := &CLogS{}
	actual = ValidateLogger(correctLog)
	assert.True(t, actual)

	wrongLog := &WLogS{}
	actual = ValidateLogger(wrongLog)
	assert.False(t, actual)

	actual = ValidateLogger(nil)
	assert.False(t, actual)
}

type CUserStorage interface {
	Get(userID, campaignKey string) schema.UserData
	Set(userID, campaignKey, variationName string)
}
type CUserStorageData struct{}

func (us *CUserStorageData) Get(userID, campaignKey string) schema.UserData {
	return schema.UserData{}
}
func (us *CUserStorageData) Set(userID, campaignKey, variationName string)  {}

type WUserStorage interface {
	Getter(userID, campaignKey string) schema.UserData
	Setter(userID, campaignKey, variationName string)
}
type WUserStorageData struct{}

func (us *WUserStorageData) Getter(userID, campaignKey string) schema.UserData {
	return schema.UserData{}
}
func (us *WUserStorageData) Setter(userID, campaignKey, variationName string)  {}

func TestValidateStorage(t *testing.T) {
	actual := ValidateStorage(nil)
	assert.True(t, actual)

	correctStorage := &CUserStorageData{}
	actual = ValidateStorage(correctStorage)
	assert.True(t, actual)

	wrongStorage := &WUserStorageData{}
	actual = ValidateStorage(wrongStorage)
	assert.False(t, actual)
}

func TestValidateActivate(t *testing.T) {
	actual := ValidateActivate("", "")
	assert.False(t, actual)

	actual = ValidateActivate("campaignKey", "userID")
	assert.True(t, actual)
}

func TestValidateGetFeatureVariableValue(t *testing.T) {
	actual := ValidateGetFeatureVariableValue("", "", "")
	assert.False(t, actual)

	actual = ValidateGetFeatureVariableValue("campaignKey", "variableKey", "userID")
	assert.True(t, actual)
}

func TestValidateGetVariationName(t *testing.T) {
	actual := ValidateGetVariationName("", "")
	assert.False(t, actual)

	actual = ValidateGetVariationName("campaignKey", "userID")
	assert.True(t, actual)
}

func TestValidateIsFeatureEnabled(t *testing.T) {
	actual := ValidateIsFeatureEnabled("", "")
	assert.False(t, actual)

	actual = ValidateIsFeatureEnabled("campaignKey", "userID")
	assert.True(t, actual)
}

func TestValidatePush(t *testing.T) {
	actual := ValidatePush("", "", "")
	assert.False(t, actual)

	actual = ValidatePush("tagKey", "tagValue", "userID")
	assert.True(t, actual)
}

func TestValidateTrack(t *testing.T) {
	actual := ValidateTrack("", "", "")
	assert.False(t, actual)

	actual = ValidateTrack("campaignKey", "userID", "goalIdentifier")
	assert.True(t, actual)
}
