package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"

)

func TestGetRequest(t *testing.T) {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	content, error := GetRequest(url)
	assert.Empty(t, error, "Could not make the Get Request")
	assert.NotEmpty(t, content, "Recieved no content")
}