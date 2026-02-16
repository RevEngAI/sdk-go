# FunctionDataTypesStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionId** | **int64** | Function id | 
**Completed** | **bool** | Whether the service has completed data types generation | 
**Status** | **string** | The current status of the data types service | 

## Methods

### NewFunctionDataTypesStatus

`func NewFunctionDataTypesStatus(functionId int64, completed bool, status string, ) *FunctionDataTypesStatus`

NewFunctionDataTypesStatus instantiates a new FunctionDataTypesStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionDataTypesStatusWithDefaults

`func NewFunctionDataTypesStatusWithDefaults() *FunctionDataTypesStatus`

NewFunctionDataTypesStatusWithDefaults instantiates a new FunctionDataTypesStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionId

`func (o *FunctionDataTypesStatus) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *FunctionDataTypesStatus) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *FunctionDataTypesStatus) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetCompleted

`func (o *FunctionDataTypesStatus) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *FunctionDataTypesStatus) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *FunctionDataTypesStatus) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.


### GetStatus

`func (o *FunctionDataTypesStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FunctionDataTypesStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FunctionDataTypesStatus) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


