# FunctionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FuncDeps** | [**[]FunctionDependency**](FunctionDependency.md) |  | 
**FuncTypes** | Pointer to [**FunctionType**](FunctionType.md) |  | [optional] 

## Methods

### NewFunctionInfo

`func NewFunctionInfo(funcDeps []FunctionDependency, ) *FunctionInfo`

NewFunctionInfo instantiates a new FunctionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionInfoWithDefaults

`func NewFunctionInfoWithDefaults() *FunctionInfo`

NewFunctionInfoWithDefaults instantiates a new FunctionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFuncDeps

`func (o *FunctionInfo) GetFuncDeps() []FunctionDependency`

GetFuncDeps returns the FuncDeps field if non-nil, zero value otherwise.

### GetFuncDepsOk

`func (o *FunctionInfo) GetFuncDepsOk() (*[]FunctionDependency, bool)`

GetFuncDepsOk returns a tuple with the FuncDeps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncDeps

`func (o *FunctionInfo) SetFuncDeps(v []FunctionDependency)`

SetFuncDeps sets FuncDeps field to given value.


### SetFuncDepsNil

`func (o *FunctionInfo) SetFuncDepsNil(b bool)`

 SetFuncDepsNil sets the value for FuncDeps to be an explicit nil

### UnsetFuncDeps
`func (o *FunctionInfo) UnsetFuncDeps()`

UnsetFuncDeps ensures that no value is present for FuncDeps, not even an explicit nil
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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


