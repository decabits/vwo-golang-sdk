package utils

import (
	"net/http"

	"github.com/decabits/vwo-golang-sdk/lib/schema"
)

//Get ...
func Get(url string) schema.Response {
	resp, err := http.Get(url)

	var response schema.Response
	response.Text = resp.Status
	response.StatusCode = resp.StatusCode
	return response
}

//GetWithParams ...
func GetWithParams(url string, params schema.Impression) schema.Response{

}
