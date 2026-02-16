# CalleeFunctionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionId** | **int64** | Unique identifier of the function | 
**MatchedFunctionId** | **NullableInt32** |  | 
**DashboardUrl** | **NullableString** |  | 
**IsExternal** | Pointer to **bool** | Indicates if the function is external | [optional] [default to false]
**CalleeName** | **string** | Name of the called function | 
**CalleeVaddr** | **string** | Virtual address of the called function | 

## Methods

### NewCalleeFunctionInfo

`func NewCalleeFunctionInfo(functionId int64, matchedFunctionId NullableInt32, dashboardUrl NullableString, calleeName string, calleeVaddr string, ) *CalleeFunctionInfo`

NewCalleeFunctionInfo instantiates a new CalleeFunctionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalleeFunctionInfoWithDefaults

`func NewCalleeFunctionInfoWithDefaults() *CalleeFunctionInfo`

NewCalleeFunctionInfoWithDefaults instantiates a new CalleeFunctionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionId

`func (o *CalleeFunctionInfo) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *CalleeFunctionInfo) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *CalleeFunctionInfo) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetMatchedFunctionId

`func (o *CalleeFunctionInfo) GetMatchedFunctionId() int32`

GetMatchedFunctionId returns the MatchedFunctionId field if non-nil, zero value otherwise.

### GetMatchedFunctionIdOk

`func (o *CalleeFunctionInfo) GetMatchedFunctionIdOk() (*int32, bool)`

GetMatchedFunctionIdOk returns a tuple with the MatchedFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedFunctionId

`func (o *CalleeFunctionInfo) SetMatchedFunctionId(v int32)`

SetMatchedFunctionId sets MatchedFunctionId field to given value.


### SetMatchedFunctionIdNil

`func (o *CalleeFunctionInfo) SetMatchedFunctionIdNil(b bool)`

 SetMatchedFunctionIdNil sets the value for MatchedFunctionId to be an explicit nil

### UnsetMatchedFunctionId
`func (o *CalleeFunctionInfo) UnsetMatchedFunctionId()`

UnsetMatchedFunctionId ensures that no value is present for MatchedFunctionId, not even an explicit nil
### GetDashboardUrl

`func (o *CalleeFunctionInfo) GetDashboardUrl() string`

GetDashboardUrl returns the DashboardUrl field if non-nil, zero value otherwise.

### GetDashboardUrlOk

`func (o *CalleeFunctionInfo) GetDashboardUrlOk() (*string, bool)`

GetDashboardUrlOk returns a tuple with the DashboardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardUrl

`func (o *CalleeFunctionInfo) SetDashboardUrl(v string)`

SetDashboardUrl sets DashboardUrl field to given value.


### SetDashboardUrlNil

`func (o *CalleeFunctionInfo) SetDashboardUrlNil(b bool)`

 SetDashboardUrlNil sets the value for DashboardUrl to be an explicit nil

### UnsetDashboardUrl
`func (o *CalleeFunctionInfo) UnsetDashboardUrl()`

UnsetDashboardUrl ensures that no value is present for DashboardUrl, not even an explicit nil
### GetIsExternal

`func (o *CalleeFunctionInfo) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *CalleeFunctionInfo) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *CalleeFunctionInfo) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *CalleeFunctionInfo) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.

### GetCalleeName

`func (o *CalleeFunctionInfo) GetCalleeName() string`

GetCalleeName returns the CalleeName field if non-nil, zero value otherwise.

### GetCalleeNameOk

`func (o *CalleeFunctionInfo) GetCalleeNameOk() (*string, bool)`

GetCalleeNameOk returns a tuple with the CalleeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalleeName

`func (o *CalleeFunctionInfo) SetCalleeName(v string)`

SetCalleeName sets CalleeName field to given value.


### GetCalleeVaddr

`func (o *CalleeFunctionInfo) GetCalleeVaddr() string`

GetCalleeVaddr returns the CalleeVaddr field if non-nil, zero value otherwise.

### GetCalleeVaddrOk

`func (o *CalleeFunctionInfo) GetCalleeVaddrOk() (*string, bool)`

GetCalleeVaddrOk returns a tuple with the CalleeVaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalleeVaddr

`func (o *CalleeFunctionInfo) SetCalleeVaddr(v string)`

SetCalleeVaddr sets CalleeVaddr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


