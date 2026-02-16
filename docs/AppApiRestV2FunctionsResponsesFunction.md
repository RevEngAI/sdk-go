# AppApiRestV2FunctionsResponsesFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionId** | Pointer to **NullableInt32** |  | [optional] 
**FunctionVaddr** | **int64** | Function virtual address | 

## Methods

### NewAppApiRestV2FunctionsResponsesFunction

`func NewAppApiRestV2FunctionsResponsesFunction(functionVaddr int64, ) *AppApiRestV2FunctionsResponsesFunction`

NewAppApiRestV2FunctionsResponsesFunction instantiates a new AppApiRestV2FunctionsResponsesFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppApiRestV2FunctionsResponsesFunctionWithDefaults

`func NewAppApiRestV2FunctionsResponsesFunctionWithDefaults() *AppApiRestV2FunctionsResponsesFunction`

NewAppApiRestV2FunctionsResponsesFunctionWithDefaults instantiates a new AppApiRestV2FunctionsResponsesFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionId

`func (o *AppApiRestV2FunctionsResponsesFunction) GetFunctionId() int32`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *AppApiRestV2FunctionsResponsesFunction) GetFunctionIdOk() (*int32, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *AppApiRestV2FunctionsResponsesFunction) SetFunctionId(v int32)`

SetFunctionId sets FunctionId field to given value.

### HasFunctionId

`func (o *AppApiRestV2FunctionsResponsesFunction) HasFunctionId() bool`

HasFunctionId returns a boolean if a field has been set.

### SetFunctionIdNil

`func (o *AppApiRestV2FunctionsResponsesFunction) SetFunctionIdNil(b bool)`

 SetFunctionIdNil sets the value for FunctionId to be an explicit nil

### UnsetFunctionId
`func (o *AppApiRestV2FunctionsResponsesFunction) UnsetFunctionId()`

UnsetFunctionId ensures that no value is present for FunctionId, not even an explicit nil
### GetFunctionVaddr

`func (o *AppApiRestV2FunctionsResponsesFunction) GetFunctionVaddr() int64`

GetFunctionVaddr returns the FunctionVaddr field if non-nil, zero value otherwise.

### GetFunctionVaddrOk

`func (o *AppApiRestV2FunctionsResponsesFunction) GetFunctionVaddrOk() (*int64, bool)`

GetFunctionVaddrOk returns a tuple with the FunctionVaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionVaddr

`func (o *AppApiRestV2FunctionsResponsesFunction) SetFunctionVaddr(v int64)`

SetFunctionVaddr sets FunctionVaddr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


