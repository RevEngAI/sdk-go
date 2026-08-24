# UpdateAnalysisDataTypesInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTypes** | [**[]UpdateDataTypeEntry**](UpdateDataTypeEntry.md) | The replacements. Every data_type_id must belong to the analysis, and none may repeat. | 

## Methods

### NewUpdateAnalysisDataTypesInputBody

`func NewUpdateAnalysisDataTypesInputBody(dataTypes []UpdateDataTypeEntry, ) *UpdateAnalysisDataTypesInputBody`

NewUpdateAnalysisDataTypesInputBody instantiates a new UpdateAnalysisDataTypesInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAnalysisDataTypesInputBodyWithDefaults

`func NewUpdateAnalysisDataTypesInputBodyWithDefaults() *UpdateAnalysisDataTypesInputBody`

NewUpdateAnalysisDataTypesInputBodyWithDefaults instantiates a new UpdateAnalysisDataTypesInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTypes

`func (o *UpdateAnalysisDataTypesInputBody) GetDataTypes() []UpdateDataTypeEntry`

GetDataTypes returns the DataTypes field if non-nil, zero value otherwise.

### GetDataTypesOk

`func (o *UpdateAnalysisDataTypesInputBody) GetDataTypesOk() (*[]UpdateDataTypeEntry, bool)`

GetDataTypesOk returns a tuple with the DataTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypes

`func (o *UpdateAnalysisDataTypesInputBody) SetDataTypes(v []UpdateDataTypeEntry)`

SetDataTypes sets DataTypes field to given value.


### SetDataTypesNil

`func (o *UpdateAnalysisDataTypesInputBody) SetDataTypesNil(b bool)`

 SetDataTypesNil sets the value for DataTypes to be an explicit nil

### UnsetDataTypes
`func (o *UpdateAnalysisDataTypesInputBody) UnsetDataTypes()`

UnsetDataTypes ensures that no value is present for DataTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


