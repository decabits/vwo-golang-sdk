package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/decabits/vwo-golang-sdk/constants"
)

// GetRequest function to do a get call
func GetRequest(url string) (string, error) {
	/*
		Args:
			url: URL needed

		Return:
			string: stringified content recieved
			error: error encountered while Get rewuest, nil if no error
	*/
	response, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf(constants.ErrorMessageURLNotFound, err.Error())
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf(constants.ErrorMessageResponceNotParsed, url)
	}
	if response.StatusCode != 200 {
		return "", fmt.Errorf(constants.ErrorMessageCouldNotGetURL, url)
	}
	return string(body), nil
}
