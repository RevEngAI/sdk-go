# TriggerDynamicExecutionInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandLineArgs** | Pointer to **string** | Command-line arguments passed to the sample when the sandbox launches it | [optional] 
**StartMethod** | Pointer to **string** | How the sandbox launches the sample. Defaults to the sandbox&#39;s standard behaviour when omitted. | [optional] 
**Timeout** | Pointer to **int64** | Maximum sandbox execution time in seconds | [optional] [default to 120]

## Methods

### NewTriggerDynamicExecutionInputBody

`func NewTriggerDynamicExecutionInputBody() *TriggerDynamicExecutionInputBody`

NewTriggerDynamicExecutionInputBody instantiates a new TriggerDynamicExecutionInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerDynamicExecutionInputBodyWithDefaults

`func NewTriggerDynamicExecutionInputBodyWithDefaults() *TriggerDynamicExecutionInputBody`

NewTriggerDynamicExecutionInputBodyWithDefaults instantiates a new TriggerDynamicExecutionInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandLineArgs

`func (o *TriggerDynamicExecutionInputBody) GetCommandLineArgs() string`

GetCommandLineArgs returns the CommandLineArgs field if non-nil, zero value otherwise.

### GetCommandLineArgsOk

`func (o *TriggerDynamicExecutionInputBody) GetCommandLineArgsOk() (*string, bool)`

GetCommandLineArgsOk returns a tuple with the CommandLineArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineArgs

`func (o *TriggerDynamicExecutionInputBody) SetCommandLineArgs(v string)`

SetCommandLineArgs sets CommandLineArgs field to given value.

### HasCommandLineArgs

`func (o *TriggerDynamicExecutionInputBody) HasCommandLineArgs() bool`

HasCommandLineArgs returns a boolean if a field has been set.

### GetStartMethod

`func (o *TriggerDynamicExecutionInputBody) GetStartMethod() string`

GetStartMethod returns the StartMethod field if non-nil, zero value otherwise.

### GetStartMethodOk

`func (o *TriggerDynamicExecutionInputBody) GetStartMethodOk() (*string, bool)`

GetStartMethodOk returns a tuple with the StartMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMethod

`func (o *TriggerDynamicExecutionInputBody) SetStartMethod(v string)`

SetStartMethod sets StartMethod field to given value.

### HasStartMethod

`func (o *TriggerDynamicExecutionInputBody) HasStartMethod() bool`

HasStartMethod returns a boolean if a field has been set.

### GetTimeout

`func (o *TriggerDynamicExecutionInputBody) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *TriggerDynamicExecutionInputBody) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *TriggerDynamicExecutionInputBody) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *TriggerDynamicExecutionInputBody) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


