# CreateAnalysisDataTypesInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTypes** | [**[]CreateDataTypeEntry**](CreateDataTypeEntry.md) | The types to create. Each namespace, name and kind must be new to the analysis. | 

## Methods

### NewCreateAnalysisDataTypesInputBody

`func NewCreateAnalysisDataTypesInputBody(dataTypes []CreateDataTypeEntry, ) *CreateAnalysisDataTypesInputBody`

NewCreateAnalysisDataTypesInputBody instantiates a new CreateAnalysisDataTypesInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAnalysisDataTypesInputBodyWithDefaults

`func NewCreateAnalysisDataTypesInputBodyWithDefaults() *CreateAnalysisDataTypesInputBody`

NewCreateAnalysisDataTypesInputBodyWithDefaults instantiates a new CreateAnalysisDataTypesInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTypes

`func (o *CreateAnalysisDataTypesInputBody) GetDataTypes() []CreateDataTypeEntry`

GetDataTypes returns the DataTypes field if non-nil, zero value otherwise.

### GetDataTypesOk

`func (o *CreateAnalysisDataTypesInputBody) GetDataTypesOk() (*[]CreateDataTypeEntry, bool)`

GetDataTypesOk returns a tuple with the DataTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypes

`func (o *CreateAnalysisDataTypesInputBody) SetDataTypes(v []CreateDataTypeEntry)`

SetDataTypes sets DataTypes field to given value.


### SetDataTypesNil

`func (o *CreateAnalysisDataTypesInputBody) SetDataTypesNil(b bool)`

 SetDataTypesNil sets the value for DataTypes to be an explicit nil

### UnsetDataTypes
`func (o *CreateAnalysisDataTypesInputBody) UnsetDataTypes()`

UnsetDataTypes ensures that no value is present for DataTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


