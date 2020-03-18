package main

import(
	"fmt"
	validate "../../vwo-golang-sdk/vwo/validate"
)

func main(){
	fmt.Println(validate.IsValidNonZeroNumber(2)) //true
	fmt.Println(validate.IsValidNonZeroNumber(0)) // false
}