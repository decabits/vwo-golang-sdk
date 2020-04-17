package utils

import (
	"github.com/decabits/vwo-golang-sdk/schema"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	storage := &UserStorageData{}
	storage.Set("abc", "cde", "efg")
}

func TestGet(t *testing.T) {
	storage := &UserStorageData{}
	actual := storage.Get("Gimmy", "php1")
	expected := schema.UserData{
		UserID: "Gimmy",
		CampaignKey: "php1",
		VariationName: "Variation-2",
	}
	assert.Equal(t, expected, actual, "Non-matching data")

	actual = storage.Get("Jimmy", "php1")
	expected = schema.UserData{
		UserID: "Jimmy",
		CampaignKey: "php1",
		VariationName: "Variation-2",
	}
	assert.Empty(t, actual, "Non-matching data")

}

func TestExist(t *testing.T) {
	storage := &UserStorageData{}
	assert.True(t, storage.Exist(), "Expected true")
}
