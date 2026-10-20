# ListFunctionGlobalsOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Decompiled** | **bool** | Whether the function has a successful AI decompilation. Its globals are recorded from that decompilation, so before one exists the list is empty however many globals the function uses. | 
**Globals** | [**[]FunctionGlobalVariable**](FunctionGlobalVariable.md) |  | 

## Methods

### NewListFunctionGlobalsOutputBody

`func NewListFunctionGlobalsOutputBody(decompiled bool, globals []FunctionGlobalVariable, ) *ListFunctionGlobalsOutputBody`

NewListFunctionGlobalsOutputBody instantiates a new ListFunctionGlobalsOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFunctionGlobalsOutputBodyWithDefaults

`func NewListFunctionGlobalsOutputBodyWithDefaults() *ListFunctionGlobalsOutputBody`

NewListFunctionGlobalsOutputBodyWithDefaults instantiates a new ListFunctionGlobalsOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecompiled

`func (o *ListFunctionGlobalsOutputBody) GetDecompiled() bool`

GetDecompiled returns the Decompiled field if non-nil, zero value otherwise.

### GetDecompiledOk

`func (o *ListFunctionGlobalsOutputBody) GetDecompiledOk() (*bool, bool)`

GetDecompiledOk returns a tuple with the Decompiled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecompiled

`func (o *ListFunctionGlobalsOutputBody) SetDecompiled(v bool)`

SetDecompiled sets Decompiled field to given value.


### GetGlobals

`func (o *ListFunctionGlobalsOutputBody) GetGlobals() []FunctionGlobalVariable`

GetGlobals returns the Globals field if non-nil, zero value otherwise.

### GetGlobalsOk

`func (o *ListFunctionGlobalsOutputBody) GetGlobalsOk() (*[]FunctionGlobalVariable, bool)`

GetGlobalsOk returns a tuple with the Globals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobals

`func (o *ListFunctionGlobalsOutputBody) SetGlobals(v []FunctionGlobalVariable)`

SetGlobals sets Globals field to given value.


### SetGlobalsNil

`func (o *ListFunctionGlobalsOutputBody) SetGlobalsNil(b bool)`

 SetGlobalsNil sets the value for Globals to be an explicit nil

### UnsetGlobals
`func (o *ListFunctionGlobalsOutputBody) UnsetGlobals()`

UnsetGlobals ensures that no value is present for Globals, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


