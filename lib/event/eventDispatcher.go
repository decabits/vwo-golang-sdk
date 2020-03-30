package event

import (
	"log"
	
	"github.com/decabits/vwo-golang-sdk/lib/schema"
	"github.com/decabits/vwo-golang-sdk/lib/utils"
)

//Dispatch ...
func Dispatch(impression schema.Impression) bool {
	URL := impression.URL
	var result bool
	response := utils.Get(URL) //impression to be passed with header
	if response.StatusCode == 200 {
		result = true
	} else {
		result = false
	}

	if result == true {
		log.Println("Impression Success")
		return true
	}

	log.Println("Impression Failed")
	return false
}