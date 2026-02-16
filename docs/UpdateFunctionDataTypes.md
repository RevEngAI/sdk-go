# UpdateFunctionDataTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTypesVersion** | **int32** | Version of the function data types, used to check this update is not overwriting a newer one | 
**DataTypes** | [**FunctionInfoInput**](FunctionInfoInput.md) | Function data types information to update | 

## Methods

### NewUpdateFunctionDataTypes

`func NewUpdateFunctionDataTypes(dataTypesVersion int32, dataTypes FunctionInfoInput, ) *UpdateFunctionDataTypes`

NewUpdateFunctionDataTypes instantiates a new UpdateFunctionDataTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFunctionDataTypesWithDefaults

`func NewUpdateFunctionDataTypesWithDefaults() *UpdateFunctionDataTypes`

NewUpdateFunctionDataTypesWithDefaults instantiates a new UpdateFunctionDataTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTypesVersion

`func (o *UpdateFunctionDataTypes) GetDataTypesVersion() int32`

GetDataTypesVersion returns the DataTypesVersion field if non-nil, zero value otherwise.

### GetDataTypesVersionOk

`func (o *UpdateFunctionDataTypes) GetDataTypesVersionOk() (*int32, bool)`

GetDataTypesVersionOk returns a tuple with the DataTypesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypesVersion

`func (o *UpdateFunctionDataTypes) SetDataTypesVersion(v int32)`

SetDataTypesVersion sets DataTypesVersion field to given value.


### GetDataTypes

`func (o *UpdateFunctionDataTypes) GetDataTypes() FunctionInfoInput`

GetDataTypes returns the DataTypes field if non-nil, zero value otherwise.

### GetDataTypesOk

`func (o *UpdateFunctionDataTypes) GetDataTypesOk() (*FunctionInfoInput, bool)`

GetDataTypesOk returns a tuple with the DataTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypes

`func (o *UpdateFunctionDataTypes) SetDataTypes(v FunctionInfoInput)`

SetDataTypes sets DataTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


