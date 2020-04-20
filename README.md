# VWO-Golang-SDK

This open source library allows you to A/B Test your Website at server-side.

## Requirements

* Works with Go 1.11 +


## Installation

```go
go get "github.com/decabits/vwo-golang-sdk"
```


## Basic usage

**Importing and Instantiation**

```go
import (
	vwo "github.com/decabits/vwo-golang-sdk"
    "github.com/decabits/vwo-golang-sdk/schema"
    "github.com/decabits/vwo-golang-sdk/api"
    "github.com/decabits/vwo-golang-sdk/service"
)

// Initialize client
VWO := vwo.Default(config.GetString("accountID"), config.GetString("SDKKey"), storage)

// // Initialize client with all parameters(explained in next section)
// vwo_client_instance = VWO.new(account_id, sdk_key, custom_logger, UserStorage.new, true, settings_file)

// Get Settings
settingsFile = service.SettingsFileManager.GetSettingsFile()

// Get Instance
// path of the required settings file is passed as an arguement 
vwoInstance = api.GetInstance(path)

// Activate API
// With Custom Variables
options = schema.Options{
        CustomVariables : { "a" = "x"},
    }
variationName = api.ActivateWithOptions(vwoInstance, campaignKey, userID, options)

// Without Custom Variables
variationName = api.Activate(vwoInstance, campaignKey, userID)

// GetVariation
// With Custom Variables
options = schema.Options{
        CustomVariables : { "a" = "x"},
    }
variationName = api.GetVariationName(vwoInstance, campaignKey, userID, options)

//Without Custom Variables
options = {}
variationName = api.GetVariationName(vwoInstance, campaignKey, userID, options)

// Track API
// With Custom Variables
options = schema.Options{
        CustomVariables : { "a" = "x"},
    }
isSuccessful = api.TrackWithOptions(vwoInstance, campaignKey, userID, goalIdentifier, options)

// With Revenue Value
options = schema.Options{
        RevenueGoal => 10.23,
    }
isSuccessful = api.TrackWithOptions(vwoInstance, campaignKey, userID, goalIdentifier, options)

// With both Custom Variables and Revenue Value
options = schema.Options{
        CustomVariables : { "a" = "x"},
        RevenueGoal : 10.23,
    }
isSuccessful = api.TrackWithOptions(vwoInstance, campaignKey, userID, goalIdentifier, options)

// FeatureEnabled API
// With Custom Varibles
options = schema.Options{
        CustomVariables : { "a" = "x"},
    }
isSuccessful = api.IsFeatureEnabled(vwoInstance, campaignKey, userID, options)

// Without Custom Variables
options = {}
isSuccessful = api.IsFeatureEnabled(vwoInstance, campaignKey, userID, options)

// GetFeatureVariableValue API
// With Custom Variables
options = schema.Options{
        CustomVariables : { "a" = "x"}
    }
variableValue = api.GetFeatureVariableValue(vwoInstance, campaignKey, variableKey, userID, options)

// Without Custom Variables
options = {}
variableValue = api.GetFeatureVariableValue(vwoInstance, campaignKey, variableKey, userID, options)

// Push API
isSuccessful = api.Push(tagKey, tagValue, userID)
```

1. `accountID` - Account for which sdk needs to be initialized
1. `SDKKey` - SDK key for that account
1. `logger` - If you need to pass your own logger. Check documentation below
1. `UserStorage.new` - An object allowing `get` and `set` for maintaining user storage
1. `developmentMode` - on/off (true/false). Default - false
1. `settingsFile` - Settings file if already present during initialization. Its stringified JSON format.


**API usage**

**User Defined Logger**

## Documentation

Refer [Official VWO Documentation](https://developers.vwo.com/reference#server-side-introduction)


## Setting up Local development environment


## Running Unit Tests

```shell
./test.sh
```

## Third-party Resources and Credits


## Authors


## Changelog


## Contributing

Please go through our [contributing guidelines](https://github.com/decabits/vwo-golang-sdk/blob/master/CONTRIBUTING.md)


## Code of Conduct

[Code of Conduct](https://github.com/decabits/vwo-golang-sdk/blob/master/CODE_OF_CONDUCT.md)


## License


