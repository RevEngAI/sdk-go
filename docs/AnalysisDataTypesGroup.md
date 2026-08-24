# AnalysisDataTypesGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisId** | **int64** |  | 
**Items** | [**[]DataTypeEntry**](DataTypeEntry.md) | The analysis&#39;s types the returned signatures name, ordered by data_type_id. | 

## Methods

### NewAnalysisDataTypesGroup

`func NewAnalysisDataTypesGroup(analysisId int64, items []DataTypeEntry, ) *AnalysisDataTypesGroup`

NewAnalysisDataTypesGroup instantiates a new AnalysisDataTypesGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisDataTypesGroupWithDefaults

`func NewAnalysisDataTypesGroupWithDefaults() *AnalysisDataTypesGroup`

NewAnalysisDataTypesGroupWithDefaults instantiates a new AnalysisDataTypesGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisId

`func (o *AnalysisDataTypesGroup) GetAnalysisId() int64`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *AnalysisDataTypesGroup) GetAnalysisIdOk() (*int64, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *AnalysisDataTypesGroup) SetAnalysisId(v int64)`

SetAnalysisId sets AnalysisId field to given value.


### GetItems

`func (o *AnalysisDataTypesGroup) GetItems() []DataTypeEntry`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AnalysisDataTypesGroup) GetItemsOk() (*[]DataTypeEntry, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AnalysisDataTypesGroup) SetItems(v []DataTypeEntry)`

SetItems sets Items field to given value.


### SetItemsNil

`func (o *AnalysisDataTypesGroup) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *AnalysisDataTypesGroup) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


