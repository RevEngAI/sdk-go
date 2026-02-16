# FunctionInfoOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FuncTypes** | Pointer to [**NullableFunctionTypeOutput**](FunctionTypeOutput.md) |  | [optional] 
**FuncDeps** | [**[]FuncDepsInner**](FuncDepsInner.md) | List of function dependencies | 

## Methods

### NewFunctionInfoOutput

`func NewFunctionInfoOutput(funcDeps []FuncDepsInner, ) *FunctionInfoOutput`

NewFunctionInfoOutput instantiates a new FunctionInfoOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionInfoOutputWithDefaults

`func NewFunctionInfoOutputWithDefaults() *FunctionInfoOutput`

NewFunctionInfoOutputWithDefaults instantiates a new FunctionInfoOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFuncTypes

`func (o *FunctionInfoOutput) GetFuncTypes() FunctionTypeOutput`

GetFuncTypes returns the FuncTypes field if non-nil, zero value otherwise.

### GetFuncTypesOk

`func (o *FunctionInfoOutput) GetFuncTypesOk() (*FunctionTypeOutput, bool)`

GetFuncTypesOk returns a tuple with the FuncTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncTypes

`func (o *FunctionInfoOutput) SetFuncTypes(v FunctionTypeOutput)`

SetFuncTypes sets FuncTypes field to given value.

### HasFuncTypes

`func (o *FunctionInfoOutput) HasFuncTypes() bool`

HasFuncTypes returns a boolean if a field has been set.

### SetFuncTypesNil

`func (o *FunctionInfoOutput) SetFuncTypesNil(b bool)`

 SetFuncTypesNil sets the value for FuncTypes to be an explicit nil

### UnsetFuncTypes
`func (o *FunctionInfoOutput) UnsetFuncTypes()`

UnsetFuncTypes ensures that no value is present for FuncTypes, not even an explicit nil
### GetFuncDeps

`func (o *FunctionInfoOutput) GetFuncDeps() []FuncDepsInner`

GetFuncDeps returns the FuncDeps field if non-nil, zero value otherwise.

### GetFuncDepsOk

`func (o *FunctionInfoOutput) GetFuncDepsOk() (*[]FuncDepsInner, bool)`

GetFuncDepsOk returns a tuple with the FuncDeps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncDeps

`func (o *FunctionInfoOutput) SetFuncDeps(v []FuncDepsInner)`

SetFuncDeps sets FuncDeps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


