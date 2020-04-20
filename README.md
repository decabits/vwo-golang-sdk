# VWO-Golang-SDK


This open source library allows you to A/B Test your Website at server-side.

## Requirements

* Works with Go 1.1x


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
// storage should be of type schema.UserStorage
VWO := vwo.Default("accountID", "SDKKey", storage)

// Get Settings
settingsFile = VWO.SettingsFile

// Activate API
// With Custom Variables
options = schema.Options{
        CustomVariables : { "a" : "x"},
    }
variationName = api.ActivateWithOptions(VWO, campaignKey, userID, options)

// Without Custom Variables
variationName = api.Activate(VWO, campaignKey, userID)


// GetVariation
// With Custom Variables
options = schema.Options{
        CustomVariables : { "a" : "x"},
    }
variationName = api.GetVariationName(VWO, campaignKey, userID, options)

//Without Custom Variables
options = {}
variationName = api.GetVariationName(VWO, campaignKey, userID, options)


// Track API
// With Custom Variables
options = schema.Options{
        CustomVariables : { "a" : "x"},
    }
isSuccessful = api.TrackWithOptions(VWO, campaignKey, userID, goalIdentifier, options)

// With Revenue Value
options = schema.Options{
        RevenueGoal => 10.23,
    }
isSuccessful = api.TrackWithOptions(VWO, campaignKey, userID, goalIdentifier, options)

// With both Custom Variables and Revenue Value
options = schema.Options{
        CustomVariables : { "a" : "x"},
        RevenueGoal : 10.23,
    }
isSuccessful = api.TrackWithOptions(VWO, campaignKey, userID, goalIdentifier, options)


// FeatureEnabled API
// With Custom Varibles
options = schema.Options{
        CustomVariables : { "a" : "x"},
    }
isSuccessful = api.IsFeatureEnabled(VWO, campaignKey, userID, options)

// Without Custom Variables
options = {}
isSuccessful = api.IsFeatureEnabled(VWO, campaignKey, userID, options)


// GetFeatureVariableValue API
// With Custom Variables
options = schema.Options{
        CustomVariables : { "a" : "x"},
    }
variableValue = api.GetFeatureVariableValue(VWO, campaignKey, variableKey, userID, options)

// Without Custom Variables
options = {}
variableValue = api.GetFeatureVariableValue(VWO, campaignKey, variableKey, userID, options)

// Push API
isSuccessful = api.Push(tagKey, tagValue, userID)
```

1. `accountID` - Account for which sdk needs to be initialized
1. `SDKKey` - SDK key for that account
1. `logger` - If you need to pass your own logger. Check documentation below
1. `UserStorage.new` - An object allowing `get` and `set` for maintaining user storage
1. `developmentMode` - on/off (true/false). Default - false
1. `settingsFile` - Settings file if already present during initialization. Its stringified JSON format.


**User Storage**




## Documentation

Refer [Official VWO Documentation](https://developers.vwo.com/reference#server-side-introduction)


## Running Unit Tests

```shell
./test.sh
```


## Authors


## Contributing

Please go through our [contributing guidelines](CONTRIBUTING.md)


## Code of Conduct

[Code of Conduct](CODE_OF_CONDUCT.md)


## License


