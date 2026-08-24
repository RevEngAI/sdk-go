# ListFunctionSignaturesOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTypes** | Pointer to [**[]AnalysisDataTypesGroup**](AnalysisDataTypesGroup.md) | The types the returned signatures name, grouped by analysis and ordered by analysis_id. Returned only when include_data_types is true. | [optional] 
**Items** | [**[]BatchFunctionSignatureEntry**](BatchFunctionSignatureEntry.md) | One entry per distinct requested function ID, in request order. A repeated ID yields one entry. | 

## Methods

### NewListFunctionSignaturesOutputBody

`func NewListFunctionSignaturesOutputBody(items []BatchFunctionSignatureEntry, ) *ListFunctionSignaturesOutputBody`

NewListFunctionSignaturesOutputBody instantiates a new ListFunctionSignaturesOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFunctionSignaturesOutputBodyWithDefaults

`func NewListFunctionSignaturesOutputBodyWithDefaults() *ListFunctionSignaturesOutputBody`

NewListFunctionSignaturesOutputBodyWithDefaults instantiates a new ListFunctionSignaturesOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTypes

`func (o *ListFunctionSignaturesOutputBody) GetDataTypes() []AnalysisDataTypesGroup`

GetDataTypes returns the DataTypes field if non-nil, zero value otherwise.

### GetDataTypesOk

`func (o *ListFunctionSignaturesOutputBody) GetDataTypesOk() (*[]AnalysisDataTypesGroup, bool)`

GetDataTypesOk returns a tuple with the DataTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypes

`func (o *ListFunctionSignaturesOutputBody) SetDataTypes(v []AnalysisDataTypesGroup)`

SetDataTypes sets DataTypes field to given value.

### HasDataTypes

`func (o *ListFunctionSignaturesOutputBody) HasDataTypes() bool`

HasDataTypes returns a boolean if a field has been set.

### SetDataTypesNil

`func (o *ListFunctionSignaturesOutputBody) SetDataTypesNil(b bool)`

 SetDataTypesNil sets the value for DataTypes to be an explicit nil

### UnsetDataTypes
`func (o *ListFunctionSignaturesOutputBody) UnsetDataTypes()`

UnsetDataTypes ensures that no value is present for DataTypes, not even an explicit nil
### GetItems

`func (o *ListFunctionSignaturesOutputBody) GetItems() []BatchFunctionSignatureEntry`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListFunctionSignaturesOutputBody) GetItemsOk() (*[]BatchFunctionSignatureEntry, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListFunctionSignaturesOutputBody) SetItems(v []BatchFunctionSignatureEntry)`

SetItems sets Items field to given value.


### SetItemsNil

`func (o *ListFunctionSignaturesOutputBody) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *ListFunctionSignaturesOutputBody) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


