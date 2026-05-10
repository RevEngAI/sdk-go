# StartupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arguments** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**Pid** | Pointer to **int64** |  | [optional] 
**Process** | Pointer to **int64** |  | [optional] 
**ProcessName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewStartupInfo

`func NewStartupInfo() *StartupInfo`

NewStartupInfo instantiates a new StartupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartupInfoWithDefaults

`func NewStartupInfoWithDefaults() *StartupInfo`

NewStartupInfoWithDefaults instantiates a new StartupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArguments

`func (o *StartupInfo) GetArguments() string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *StartupInfo) GetArgumentsOk() (*string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *StartupInfo) SetArguments(v string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *StartupInfo) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetError

`func (o *StartupInfo) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StartupInfo) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StartupInfo) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *StartupInfo) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorCode

`func (o *StartupInfo) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *StartupInfo) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *StartupInfo) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *StartupInfo) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetPid

`func (o *StartupInfo) GetPid() int64`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *StartupInfo) GetPidOk() (*int64, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *StartupInfo) SetPid(v int64)`

SetPid sets Pid field to given value.

### HasPid

`func (o *StartupInfo) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetProcess

`func (o *StartupInfo) GetProcess() int64`

GetProcess returns the Process field if non-nil, zero value otherwise.

### GetProcessOk

`func (o *StartupInfo) GetProcessOk() (*int64, bool)`

GetProcessOk returns a tuple with the Process field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcess

`func (o *StartupInfo) SetProcess(v int64)`

SetProcess sets Process field to given value.

### HasProcess

`func (o *StartupInfo) HasProcess() bool`

HasProcess returns a boolean if a field has been set.

### GetProcessName

`func (o *StartupInfo) GetProcessName() string`

GetProcessName returns the ProcessName field if non-nil, zero value otherwise.

### GetProcessNameOk

`func (o *StartupInfo) GetProcessNameOk() (*string, bool)`

GetProcessNameOk returns a tuple with the ProcessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessName

`func (o *StartupInfo) SetProcessName(v string)`

SetProcessName sets ProcessName field to given value.

### HasProcessName

`func (o *StartupInfo) HasProcessName() bool`

HasProcessName returns a boolean if a field has been set.

### GetStatus

`func (o *StartupInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StartupInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StartupInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StartupInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


