# AnalysisDataTypesOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTypes** | [**[]DataTypeEntry**](DataTypeEntry.md) | The stored types, ordered by data_type_id. | 

## Methods

### NewAnalysisDataTypesOutputBody

`func NewAnalysisDataTypesOutputBody(dataTypes []DataTypeEntry, ) *AnalysisDataTypesOutputBody`

NewAnalysisDataTypesOutputBody instantiates a new AnalysisDataTypesOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisDataTypesOutputBodyWithDefaults

`func NewAnalysisDataTypesOutputBodyWithDefaults() *AnalysisDataTypesOutputBody`

NewAnalysisDataTypesOutputBodyWithDefaults instantiates a new AnalysisDataTypesOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTypes

`func (o *AnalysisDataTypesOutputBody) GetDataTypes() []DataTypeEntry`

GetDataTypes returns the DataTypes field if non-nil, zero value otherwise.

### GetDataTypesOk

`func (o *AnalysisDataTypesOutputBody) GetDataTypesOk() (*[]DataTypeEntry, bool)`

GetDataTypesOk returns a tuple with the DataTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypes

`func (o *AnalysisDataTypesOutputBody) SetDataTypes(v []DataTypeEntry)`

SetDataTypes sets DataTypes field to given value.


### SetDataTypesNil

`func (o *AnalysisDataTypesOutputBody) SetDataTypesNil(b bool)`

 SetDataTypesNil sets the value for DataTypes to be an explicit nil

### UnsetDataTypes
`func (o *AnalysisDataTypesOutputBody) UnsetDataTypes()`

UnsetDataTypes ensures that no value is present for DataTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


