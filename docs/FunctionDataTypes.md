# FunctionDataTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completed** | **bool** | Whether the service has completed data types generation | 
**Status** | **string** | The current status of the data types service | 
**DataTypes** | Pointer to [**NullableFunctionInfoOutput**](FunctionInfoOutput.md) |  | [optional] 
**DataTypesVersion** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewFunctionDataTypes

`func NewFunctionDataTypes(completed bool, status string, ) *FunctionDataTypes`

NewFunctionDataTypes instantiates a new FunctionDataTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionDataTypesWithDefaults

`func NewFunctionDataTypesWithDefaults() *FunctionDataTypes`

NewFunctionDataTypesWithDefaults instantiates a new FunctionDataTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleted

`func (o *FunctionDataTypes) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *FunctionDataTypes) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *FunctionDataTypes) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.


### GetStatus

`func (o *FunctionDataTypes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FunctionDataTypes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FunctionDataTypes) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDataTypes

`func (o *FunctionDataTypes) GetDataTypes() FunctionInfoOutput`

GetDataTypes returns the DataTypes field if non-nil, zero value otherwise.

### GetDataTypesOk

`func (o *FunctionDataTypes) GetDataTypesOk() (*FunctionInfoOutput, bool)`

GetDataTypesOk returns a tuple with the DataTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypes

`func (o *FunctionDataTypes) SetDataTypes(v FunctionInfoOutput)`

SetDataTypes sets DataTypes field to given value.

### HasDataTypes

`func (o *FunctionDataTypes) HasDataTypes() bool`

HasDataTypes returns a boolean if a field has been set.

### SetDataTypesNil

`func (o *FunctionDataTypes) SetDataTypesNil(b bool)`

 SetDataTypesNil sets the value for DataTypes to be an explicit nil

### UnsetDataTypes
`func (o *FunctionDataTypes) UnsetDataTypes()`

UnsetDataTypes ensures that no value is present for DataTypes, not even an explicit nil
### GetDataTypesVersion

`func (o *FunctionDataTypes) GetDataTypesVersion() int32`

GetDataTypesVersion returns the DataTypesVersion field if non-nil, zero value otherwise.

### GetDataTypesVersionOk

`func (o *FunctionDataTypes) GetDataTypesVersionOk() (*int32, bool)`

GetDataTypesVersionOk returns a tuple with the DataTypesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypesVersion

`func (o *FunctionDataTypes) SetDataTypesVersion(v int32)`

SetDataTypesVersion sets DataTypesVersion field to given value.

### HasDataTypesVersion

`func (o *FunctionDataTypes) HasDataTypesVersion() bool`

HasDataTypesVersion returns a boolean if a field has been set.

### SetDataTypesVersionNil

`func (o *FunctionDataTypes) SetDataTypesVersionNil(b bool)`

 SetDataTypesVersionNil sets the value for DataTypesVersion to be an explicit nil

### UnsetDataTypesVersion
`func (o *FunctionDataTypes) UnsetDataTypesVersion()`

UnsetDataTypesVersion ensures that no value is present for DataTypesVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


