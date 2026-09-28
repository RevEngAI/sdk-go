# GetFunctionMapsOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionMap** | **map[string]int64** | Function ID (as a string key) to virtual address, for every function in the analysis&#39;s binary. | 
**InverseFunctionMap** | **map[string]int64** | Virtual address (as a string key) to function ID — the inverse of function_map. | 
**NameMap** | **map[string]string** | Virtual address (as a string key) to mangled function name. Empty string for a function with no mangled name. | 

## Methods

### NewGetFunctionMapsOutputBody

`func NewGetFunctionMapsOutputBody(functionMap map[string]int64, inverseFunctionMap map[string]int64, nameMap map[string]string, ) *GetFunctionMapsOutputBody`

NewGetFunctionMapsOutputBody instantiates a new GetFunctionMapsOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFunctionMapsOutputBodyWithDefaults

`func NewGetFunctionMapsOutputBodyWithDefaults() *GetFunctionMapsOutputBody`

NewGetFunctionMapsOutputBodyWithDefaults instantiates a new GetFunctionMapsOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionMap

`func (o *GetFunctionMapsOutputBody) GetFunctionMap() map[string]int64`

GetFunctionMap returns the FunctionMap field if non-nil, zero value otherwise.

### GetFunctionMapOk

`func (o *GetFunctionMapsOutputBody) GetFunctionMapOk() (*map[string]int64, bool)`

GetFunctionMapOk returns a tuple with the FunctionMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionMap

`func (o *GetFunctionMapsOutputBody) SetFunctionMap(v map[string]int64)`

SetFunctionMap sets FunctionMap field to given value.


### GetInverseFunctionMap

`func (o *GetFunctionMapsOutputBody) GetInverseFunctionMap() map[string]int64`

GetInverseFunctionMap returns the InverseFunctionMap field if non-nil, zero value otherwise.

### GetInverseFunctionMapOk

`func (o *GetFunctionMapsOutputBody) GetInverseFunctionMapOk() (*map[string]int64, bool)`

GetInverseFunctionMapOk returns a tuple with the InverseFunctionMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInverseFunctionMap

`func (o *GetFunctionMapsOutputBody) SetInverseFunctionMap(v map[string]int64)`

SetInverseFunctionMap sets InverseFunctionMap field to given value.


### GetNameMap

`func (o *GetFunctionMapsOutputBody) GetNameMap() map[string]string`

GetNameMap returns the NameMap field if non-nil, zero value otherwise.

### GetNameMapOk

`func (o *GetFunctionMapsOutputBody) GetNameMapOk() (*map[string]string, bool)`

GetNameMapOk returns a tuple with the NameMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameMap

`func (o *GetFunctionMapsOutputBody) SetNameMap(v map[string]string)`

SetNameMap sets NameMap field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


