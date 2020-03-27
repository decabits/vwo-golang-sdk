package event

import (
	"log"
	
	"github.com/decabits/vwo-golang-sdk/lib/schema"
	"github.com/decabits/vwo-golang-sdk/lib/utils"
)

//Dispatch ...
func Dispatch(impression schema.Impression) bool{
	isDevelopmentMode := false
	URL := impression.URL
	var result bool
	if (isDevelopmentMode){
		result := true
	} else{
		response := utils.Get(URL) //impression to be passed 
		if (response.StatusCode == 200){
			result := true
		} else{
			result := false
		}
	}

	if (result == true) {
		log.Println("Impression Success")
		return true
	}
	
	log.Println("Impression Failed")
	return false


}