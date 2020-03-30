package utils

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/decabits/vwo-golang-sdk/lib/schema"
)

//Get ...
func Get(url string) (schema.Response, error) {
	response, err := http.Get(url)
	if err != nil {
		return schema.Response{}, errors.New("URL not Found")
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return schema.Response{}, errors.New("Error parsing response")
	}
	resp := schema.Response{
		StatusCode: response.StatusCode,
		Text:       string(body),
	}
	return resp, nil
}

/*
//GetWithParams ...
func GetWithParams(url string, params schema.Impression) schema.Response{

}

*/
