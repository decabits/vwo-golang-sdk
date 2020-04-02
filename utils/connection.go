package utils

import (
	"errors"
	"io/ioutil"
	"net/http"
)

// GetRequest function to do a get call
func GetRequest(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", errors.New("URL not Found" + err.Error())
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", errors.New("Error parsing response")
	}
	if response.StatusCode != 200 {
		return "", errors.New("Failed get request")
	}
	return string(body), nil
}
