# FunctionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FuncTypes** | Pointer to [**NullableFunctionType**](FunctionType.md) |  | [optional] 
**FuncDeps** | [**[]FuncDepsInner**](FuncDepsInner.md) | List of function dependencies | 

## Methods

### NewFunctionInfo

`func NewFunctionInfo(funcDeps []FuncDepsInner, ) *FunctionInfo`

NewFunctionInfo instantiates a new FunctionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionInfoWithDefaults

`func NewFunctionInfoWithDefaults() *FunctionInfo`

NewFunctionInfoWithDefaults instantiates a new FunctionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFuncTypes

`func (o *FunctionInfo) GetFuncTypes() FunctionType`

GetFuncTypes returns the FuncTypes field if non-nil, zero value otherwise.

### GetFuncTypesOk

`func (o *FunctionInfo) GetFuncTypesOk() (*FunctionType, bool)`

GetFuncTypesOk returns a tuple with the FuncTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncTypes

`func (o *FunctionInfo) SetFuncTypes(v FunctionType)`

SetFuncTypes sets FuncTypes field to given value.

### HasFuncTypes

`func (o *FunctionInfo) HasFuncTypes() bool`

HasFuncTypes returns a boolean if a field has been set.

### SetFuncTypesNil

`func (o *FunctionInfo) SetFuncTypesNil(b bool)`

 SetFuncTypesNil sets the value for FuncTypes to be an explicit nil

### UnsetFuncTypes
`func (o *FunctionInfo) UnsetFuncTypes()`

UnsetFuncTypes ensures that no value is present for FuncTypes, not even an explicit nil
### GetFuncDeps

`func (o *FunctionInfo) GetFuncDeps() []FuncDepsInner`

GetFuncDeps returns the FuncDeps field if non-nil, zero value otherwise.

### GetFuncDepsOk

`func (o *FunctionInfo) GetFuncDepsOk() (*[]FuncDepsInner, bool)`

GetFuncDepsOk returns a tuple with the FuncDeps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncDeps

`func (o *FunctionInfo) SetFuncDeps(v []FuncDepsInner)`

SetFuncDeps sets FuncDeps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


