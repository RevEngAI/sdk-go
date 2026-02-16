# FunctionMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionMap** | **map[string]int32** | Mapping of remote function ids to local function addresses | 
**InverseFunctionMap** | **map[string]int32** | Mapping of local function addresses to remote function ids | 
**NameMap** | **map[string]string** | Mapping of local function addresses to mangled names | 

## Methods

### NewFunctionMapping

`func NewFunctionMapping(functionMap map[string]int32, inverseFunctionMap map[string]int32, nameMap map[string]string, ) *FunctionMapping`

NewFunctionMapping instantiates a new FunctionMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionMappingWithDefaults

`func NewFunctionMappingWithDefaults() *FunctionMapping`

NewFunctionMappingWithDefaults instantiates a new FunctionMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionMap

`func (o *FunctionMapping) GetFunctionMap() map[string]int32`

GetFunctionMap returns the FunctionMap field if non-nil, zero value otherwise.

### GetFunctionMapOk

`func (o *FunctionMapping) GetFunctionMapOk() (*map[string]int32, bool)`

GetFunctionMapOk returns a tuple with the FunctionMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionMap

`func (o *FunctionMapping) SetFunctionMap(v map[string]int32)`

SetFunctionMap sets FunctionMap field to given value.


### GetInverseFunctionMap

`func (o *FunctionMapping) GetInverseFunctionMap() map[string]int32`

GetInverseFunctionMap returns the InverseFunctionMap field if non-nil, zero value otherwise.

### GetInverseFunctionMapOk

`func (o *FunctionMapping) GetInverseFunctionMapOk() (*map[string]int32, bool)`

GetInverseFunctionMapOk returns a tuple with the InverseFunctionMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInverseFunctionMap

`func (o *FunctionMapping) SetInverseFunctionMap(v map[string]int32)`

SetInverseFunctionMap sets InverseFunctionMap field to given value.


### GetNameMap

`func (o *FunctionMapping) GetNameMap() map[string]string`

GetNameMap returns the NameMap field if non-nil, zero value otherwise.

### GetNameMapOk

`func (o *FunctionMapping) GetNameMapOk() (*map[string]string, bool)`

GetNameMapOk returns a tuple with the NameMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameMap

`func (o *FunctionMapping) SetNameMap(v map[string]string)`

SetNameMap sets NameMap field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


