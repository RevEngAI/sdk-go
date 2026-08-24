# FunctionTypeDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | [**[]DataTypeFunctionParameterEntry**](DataTypeFunctionParameterEntry.md) | The parameters, in argument order. | 
**ReturnDataTypeId** | Pointer to **int64** | The return type. | [optional] 

## Methods

### NewFunctionTypeDefinition

`func NewFunctionTypeDefinition(parameters []DataTypeFunctionParameterEntry, ) *FunctionTypeDefinition`

NewFunctionTypeDefinition instantiates a new FunctionTypeDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionTypeDefinitionWithDefaults

`func NewFunctionTypeDefinitionWithDefaults() *FunctionTypeDefinition`

NewFunctionTypeDefinitionWithDefaults instantiates a new FunctionTypeDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *FunctionTypeDefinition) GetParameters() []DataTypeFunctionParameterEntry`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *FunctionTypeDefinition) GetParametersOk() (*[]DataTypeFunctionParameterEntry, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *FunctionTypeDefinition) SetParameters(v []DataTypeFunctionParameterEntry)`

SetParameters sets Parameters field to given value.


### SetParametersNil

`func (o *FunctionTypeDefinition) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *FunctionTypeDefinition) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetReturnDataTypeId

`func (o *FunctionTypeDefinition) GetReturnDataTypeId() int64`

GetReturnDataTypeId returns the ReturnDataTypeId field if non-nil, zero value otherwise.

### GetReturnDataTypeIdOk

`func (o *FunctionTypeDefinition) GetReturnDataTypeIdOk() (*int64, bool)`

GetReturnDataTypeIdOk returns a tuple with the ReturnDataTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnDataTypeId

`func (o *FunctionTypeDefinition) SetReturnDataTypeId(v int64)`

SetReturnDataTypeId sets ReturnDataTypeId field to given value.

### HasReturnDataTypeId

`func (o *FunctionTypeDefinition) HasReturnDataTypeId() bool`

HasReturnDataTypeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


