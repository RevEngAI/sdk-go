# DataTypeFunctionEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionId** | **int64** |  | 
**FunctionName** | **string** | Current name of the function. | 
**UsageKinds** | **[]string** | How the function uses the type, each listed once, in alphabetical order: signature (its return type or a parameter type), inline (its AI decompilation names the type in its code), member (its AI decompilation reads a struct or union member of this type, or it references a global variable of this type). | 

## Methods

### NewDataTypeFunctionEntry

`func NewDataTypeFunctionEntry(functionId int64, functionName string, usageKinds []string, ) *DataTypeFunctionEntry`

NewDataTypeFunctionEntry instantiates a new DataTypeFunctionEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTypeFunctionEntryWithDefaults

`func NewDataTypeFunctionEntryWithDefaults() *DataTypeFunctionEntry`

NewDataTypeFunctionEntryWithDefaults instantiates a new DataTypeFunctionEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionId

`func (o *DataTypeFunctionEntry) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *DataTypeFunctionEntry) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *DataTypeFunctionEntry) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetFunctionName

`func (o *DataTypeFunctionEntry) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *DataTypeFunctionEntry) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *DataTypeFunctionEntry) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetUsageKinds

`func (o *DataTypeFunctionEntry) GetUsageKinds() []string`

GetUsageKinds returns the UsageKinds field if non-nil, zero value otherwise.

### GetUsageKindsOk

`func (o *DataTypeFunctionEntry) GetUsageKindsOk() (*[]string, bool)`

GetUsageKindsOk returns a tuple with the UsageKinds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageKinds

`func (o *DataTypeFunctionEntry) SetUsageKinds(v []string)`

SetUsageKinds sets UsageKinds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


