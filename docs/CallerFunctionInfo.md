# CallerFunctionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionId** | **int64** | Unique identifier of the function | 
**MatchedFunctionId** | **NullableInt32** |  | 
**DashboardUrl** | **NullableString** |  | 
**IsExternal** | Pointer to **bool** | Indicates if the function is external | [optional] [default to false]
**CallerName** | **string** | Name of the calling function | 
**CallerVaddr** | **string** | Virtual address of the calling function | 

## Methods

### NewCallerFunctionInfo

`func NewCallerFunctionInfo(functionId int64, matchedFunctionId NullableInt32, dashboardUrl NullableString, callerName string, callerVaddr string, ) *CallerFunctionInfo`

NewCallerFunctionInfo instantiates a new CallerFunctionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallerFunctionInfoWithDefaults

`func NewCallerFunctionInfoWithDefaults() *CallerFunctionInfo`

NewCallerFunctionInfoWithDefaults instantiates a new CallerFunctionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionId

`func (o *CallerFunctionInfo) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *CallerFunctionInfo) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *CallerFunctionInfo) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetMatchedFunctionId

`func (o *CallerFunctionInfo) GetMatchedFunctionId() int32`

GetMatchedFunctionId returns the MatchedFunctionId field if non-nil, zero value otherwise.

### GetMatchedFunctionIdOk

`func (o *CallerFunctionInfo) GetMatchedFunctionIdOk() (*int32, bool)`

GetMatchedFunctionIdOk returns a tuple with the MatchedFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedFunctionId

`func (o *CallerFunctionInfo) SetMatchedFunctionId(v int32)`

SetMatchedFunctionId sets MatchedFunctionId field to given value.


### SetMatchedFunctionIdNil

`func (o *CallerFunctionInfo) SetMatchedFunctionIdNil(b bool)`

 SetMatchedFunctionIdNil sets the value for MatchedFunctionId to be an explicit nil

### UnsetMatchedFunctionId
`func (o *CallerFunctionInfo) UnsetMatchedFunctionId()`

UnsetMatchedFunctionId ensures that no value is present for MatchedFunctionId, not even an explicit nil
### GetDashboardUrl

`func (o *CallerFunctionInfo) GetDashboardUrl() string`

GetDashboardUrl returns the DashboardUrl field if non-nil, zero value otherwise.

### GetDashboardUrlOk

`func (o *CallerFunctionInfo) GetDashboardUrlOk() (*string, bool)`

GetDashboardUrlOk returns a tuple with the DashboardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardUrl

`func (o *CallerFunctionInfo) SetDashboardUrl(v string)`

SetDashboardUrl sets DashboardUrl field to given value.


### SetDashboardUrlNil

`func (o *CallerFunctionInfo) SetDashboardUrlNil(b bool)`

 SetDashboardUrlNil sets the value for DashboardUrl to be an explicit nil

### UnsetDashboardUrl
`func (o *CallerFunctionInfo) UnsetDashboardUrl()`

UnsetDashboardUrl ensures that no value is present for DashboardUrl, not even an explicit nil
### GetIsExternal

`func (o *CallerFunctionInfo) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *CallerFunctionInfo) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *CallerFunctionInfo) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *CallerFunctionInfo) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.

### GetCallerName

`func (o *CallerFunctionInfo) GetCallerName() string`

GetCallerName returns the CallerName field if non-nil, zero value otherwise.

### GetCallerNameOk

`func (o *CallerFunctionInfo) GetCallerNameOk() (*string, bool)`

GetCallerNameOk returns a tuple with the CallerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerName

`func (o *CallerFunctionInfo) SetCallerName(v string)`

SetCallerName sets CallerName field to given value.


### GetCallerVaddr

`func (o *CallerFunctionInfo) GetCallerVaddr() string`

GetCallerVaddr returns the CallerVaddr field if non-nil, zero value otherwise.

### GetCallerVaddrOk

`func (o *CallerFunctionInfo) GetCallerVaddrOk() (*string, bool)`

GetCallerVaddrOk returns a tuple with the CallerVaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerVaddr

`func (o *CallerFunctionInfo) SetCallerVaddr(v string)`

SetCallerVaddr sets CallerVaddr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


