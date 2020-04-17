package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"

)

func TestGetRequest(t *testing.T) {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	content, err := GetRequest(url)
	assert.Nil(t, err, "Could not make the Get Request")
	assert.NotEmpty(t, content, "Recieved no content")

	url = "https.com"
	content, err = GetRequest(url)
	assert.NotNil(t, err, "Could not make the Get Request")
	assert.Empty(t, content, "Recieved no content")

	url = "https://github.com/vaibhavsethia/Portfolio---Me"
	content, err = GetRequest(url)
	assert.NotNil(t, err, "Could not make the Get Request")
	assert.Empty(t, content, "Recieved no content")
}

