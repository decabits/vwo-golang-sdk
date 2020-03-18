package vwo

import(
	"fmt"
	"log"
	"net/http"
	"io"
	// "constants"
	// "vwo" "vwo-golang-sdk" 
	// constants "vwo/constants"
	// singleton "../vwo/segments"
)

func Test() {
	// fmt.Println(log.error())
	fmt.Println(constants.ConstantPlatform)
}

type VWO struct{

}

func (self *VWO)  __init__(settings_file, logger, user_storage, is_development_mode, *args, **kwargs )
{
	
	/*
		Initializes the services required by the VWO APIs.

        Args:
            settings_file: JSON string representing the project.
            logger(object): Optional component which provides a log method
                to log messages. By default everything would be logged.
            user_storage(object): Optional component which provides
                methods to store and manage user data.
            is_development_mode(bool): To specify whether the request
				to our server should be sent or not.
	*/


}