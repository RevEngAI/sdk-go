# V2FunctionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FuncTypes** | Pointer to [**NullableV2FunctionType**](V2FunctionType.md) |  | [optional] 
**FuncDeps** | [**[]FuncDepsInner**](FuncDepsInner.md) | List of function dependencies | 

## Methods

### NewV2FunctionInfo

`func NewV2FunctionInfo(funcDeps []FuncDepsInner, ) *V2FunctionInfo`

NewV2FunctionInfo instantiates a new V2FunctionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2FunctionInfoWithDefaults

`func NewV2FunctionInfoWithDefaults() *V2FunctionInfo`

NewV2FunctionInfoWithDefaults instantiates a new V2FunctionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFuncTypes

`func (o *V2FunctionInfo) GetFuncTypes() V2FunctionType`

GetFuncTypes returns the FuncTypes field if non-nil, zero value otherwise.

### GetFuncTypesOk

`func (o *V2FunctionInfo) GetFuncTypesOk() (*V2FunctionType, bool)`

GetFuncTypesOk returns a tuple with the FuncTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncTypes

`func (o *V2FunctionInfo) SetFuncTypes(v V2FunctionType)`

SetFuncTypes sets FuncTypes field to given value.

### HasFuncTypes

`func (o *V2FunctionInfo) HasFuncTypes() bool`

HasFuncTypes returns a boolean if a field has been set.

### SetFuncTypesNil

`func (o *V2FunctionInfo) SetFuncTypesNil(b bool)`

 SetFuncTypesNil sets the value for FuncTypes to be an explicit nil

### UnsetFuncTypes
`func (o *V2FunctionInfo) UnsetFuncTypes()`

UnsetFuncTypes ensures that no value is present for FuncTypes, not even an explicit nil
### GetFuncDeps

`func (o *V2FunctionInfo) GetFuncDeps() []FuncDepsInner`

GetFuncDeps returns the FuncDeps field if non-nil, zero value otherwise.

### GetFuncDepsOk

`func (o *V2FunctionInfo) GetFuncDepsOk() (*[]FuncDepsInner, bool)`

GetFuncDepsOk returns a tuple with the FuncDeps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncDeps

`func (o *V2FunctionInfo) SetFuncDeps(v []FuncDepsInner)`

SetFuncDeps sets FuncDeps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


