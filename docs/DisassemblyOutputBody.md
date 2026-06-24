# DisassemblyOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasicBlocks** | Pointer to **interface{}** |  | [optional] 
**FunctionId** | **int64** |  | 
**LocalVariables** | Pointer to **interface{}** |  | [optional] 
**Params** | Pointer to **interface{}** |  | [optional] 
**ReturnType** | Pointer to **NullableString** |  | [optional] 
**Returns** | **bool** |  | 

## Methods

### NewDisassemblyOutputBody

`func NewDisassemblyOutputBody(functionId int64, returns bool, ) *DisassemblyOutputBody`

NewDisassemblyOutputBody instantiates a new DisassemblyOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisassemblyOutputBodyWithDefaults

`func NewDisassemblyOutputBodyWithDefaults() *DisassemblyOutputBody`

NewDisassemblyOutputBodyWithDefaults instantiates a new DisassemblyOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasicBlocks

`func (o *DisassemblyOutputBody) GetBasicBlocks() interface{}`

GetBasicBlocks returns the BasicBlocks field if non-nil, zero value otherwise.

### GetBasicBlocksOk

`func (o *DisassemblyOutputBody) GetBasicBlocksOk() (*interface{}, bool)`

GetBasicBlocksOk returns a tuple with the BasicBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicBlocks

`func (o *DisassemblyOutputBody) SetBasicBlocks(v interface{})`

SetBasicBlocks sets BasicBlocks field to given value.

### HasBasicBlocks

`func (o *DisassemblyOutputBody) HasBasicBlocks() bool`

HasBasicBlocks returns a boolean if a field has been set.

### SetBasicBlocksNil

`func (o *DisassemblyOutputBody) SetBasicBlocksNil(b bool)`

 SetBasicBlocksNil sets the value for BasicBlocks to be an explicit nil

### UnsetBasicBlocks
`func (o *DisassemblyOutputBody) UnsetBasicBlocks()`

UnsetBasicBlocks ensures that no value is present for BasicBlocks, not even an explicit nil
### GetFunctionId

`func (o *DisassemblyOutputBody) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *DisassemblyOutputBody) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *DisassemblyOutputBody) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetLocalVariables

`func (o *DisassemblyOutputBody) GetLocalVariables() interface{}`

GetLocalVariables returns the LocalVariables field if non-nil, zero value otherwise.

### GetLocalVariablesOk

`func (o *DisassemblyOutputBody) GetLocalVariablesOk() (*interface{}, bool)`

GetLocalVariablesOk returns a tuple with the LocalVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVariables

`func (o *DisassemblyOutputBody) SetLocalVariables(v interface{})`

SetLocalVariables sets LocalVariables field to given value.

### HasLocalVariables

`func (o *DisassemblyOutputBody) HasLocalVariables() bool`

HasLocalVariables returns a boolean if a field has been set.

### SetLocalVariablesNil

`func (o *DisassemblyOutputBody) SetLocalVariablesNil(b bool)`

 SetLocalVariablesNil sets the value for LocalVariables to be an explicit nil

### UnsetLocalVariables
`func (o *DisassemblyOutputBody) UnsetLocalVariables()`

UnsetLocalVariables ensures that no value is present for LocalVariables, not even an explicit nil
### GetParams

`func (o *DisassemblyOutputBody) GetParams() interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *DisassemblyOutputBody) GetParamsOk() (*interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *DisassemblyOutputBody) SetParams(v interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *DisassemblyOutputBody) HasParams() bool`

HasParams returns a boolean if a field has been set.

### SetParamsNil

`func (o *DisassemblyOutputBody) SetParamsNil(b bool)`

 SetParamsNil sets the value for Params to be an explicit nil

### UnsetParams
`func (o *DisassemblyOutputBody) UnsetParams()`

UnsetParams ensures that no value is present for Params, not even an explicit nil
### GetReturnType

`func (o *DisassemblyOutputBody) GetReturnType() string`

GetReturnType returns the ReturnType field if non-nil, zero value otherwise.

### GetReturnTypeOk

`func (o *DisassemblyOutputBody) GetReturnTypeOk() (*string, bool)`

GetReturnTypeOk returns a tuple with the ReturnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnType

`func (o *DisassemblyOutputBody) SetReturnType(v string)`

SetReturnType sets ReturnType field to given value.

### HasReturnType

`func (o *DisassemblyOutputBody) HasReturnType() bool`

HasReturnType returns a boolean if a field has been set.

### SetReturnTypeNil

`func (o *DisassemblyOutputBody) SetReturnTypeNil(b bool)`

 SetReturnTypeNil sets the value for ReturnType to be an explicit nil

### UnsetReturnType
`func (o *DisassemblyOutputBody) UnsetReturnType()`

UnsetReturnType ensures that no value is present for ReturnType, not even an explicit nil
### GetReturns

`func (o *DisassemblyOutputBody) GetReturns() bool`

GetReturns returns the Returns field if non-nil, zero value otherwise.

### GetReturnsOk

`func (o *DisassemblyOutputBody) GetReturnsOk() (*bool, bool)`

GetReturnsOk returns a tuple with the Returns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturns

`func (o *DisassemblyOutputBody) SetReturns(v bool)`

SetReturns sets Returns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


