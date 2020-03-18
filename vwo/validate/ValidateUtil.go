package validate

import(
	// "sys"
	// "encoding/json"
	// "jsonschema"
	"reflect"
)

var services = map[string] []string{
	"logger": []string{"log"},
	"eventDispatcher": []string{"dispatch"},
	"userStorage": []string{"get","set"},
} 

/*
Issue: data type for file is needed (and some functions)
func IsValidSettingsFile(settingsFile){
	
}
*/

func IsValidService(service, serviceName string) bool{
	/*
		Checks whether the service passed by the user
			contains the necessary methods or not

		Args:
			service (classobj): User defined class instance
			service_name (string): Name of the service

		Returns:
			bool: Whether the class instance provided is valid or not
		*/
	serviceAttribute :=services[serviceName]
	/*
	Issue: cannot compare for Nullness
	if serviceAttribute == nil{
		return false
	}
	*/
	for i:=0;i<len(serviceAttribute);i++{
		if serviceAttribute[i] == service{
			return true
		}
	}
	return false
}

func IsValidLogLevel(level string) bool{
	stringLevels := [8] string{
		"CRITICAL",
		"FATAL",
		"ERROR",
		"WARN",
		"WARNING",
		"INFO",
		"DEBUG",
		"NOTEST",
	}
	for i :=0; i<8; i++{
		if stringLevels[i] == level{
			return true
		}
	}
	return false
}

/*
Issue: No dictionary data type
func IsValidDict(val []string) bool{
	if reflect.TypeOf(val) == dict{
		return true
	}
	return false
}
*/

/*
Issue: Need to pass the type of the arguement
func IsValidValue(val) bool{
	if val != nil{
		return bool(val)
	}
	return false
}
*/

func IsValidNonZeroNumber(val int) bool{
	var temp int
	if reflect.TypeOf(val) == reflect.TypeOf(temp) && val>0{
		return IsValidNumber(val)
	}
	return  false
}


func IsValidNumber(val int) bool{
	var temp int
	if reflect.TypeOf(val) == reflect.TypeOf(temp){
		return true
	}
	return false
}

func IsValidString(val string) bool{
	var temp string
	if reflect.TypeOf(val) == reflect.TypeOf(temp){
		return true
	}
	return false
}

/*
Issue: Need to pass the type of the arguement
func IsValidBasicDataType(val) bool{
	if reflect.TypeOf(val) == reflect.TypeOf(string) || 
	reflect.TypeOf(val) == reflect.TypeOf(int) || 
	reflect.TypeOf(val) == reflect.TypeOf(float32) || 
	reflect.TypeOf(val) == reflect.TypeOf(float64) ||
	reflect.TypeOf(val) == reflect.TypeOf(bool) {
		return true
	}
	return false
}
*/