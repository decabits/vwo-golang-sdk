package utils

import (
	"math/rand"
	"time"
)
// GetKeyValue function
func GetKeyValue(segment {}interface) (string, string) {
	 // To be done
	return "",""	
}

//GetRandomNumber function
func GetRandomNumber() float32{
	return rand.Float32()
}

//GetCurrentUnixTimestamp function
func GetCurrentUnixTimestamp() string{
	return string(time.Now().Unix())
}