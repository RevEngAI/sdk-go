# ListAnalysisDataTypesOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]DataTypeEntry**](DataTypeEntry.md) |  | 
**TotalCount** | **int64** | Total types matching the filters, ignoring pagination. | 

## Methods

### NewListAnalysisDataTypesOutputBody

`func NewListAnalysisDataTypesOutputBody(items []DataTypeEntry, totalCount int64, ) *ListAnalysisDataTypesOutputBody`

NewListAnalysisDataTypesOutputBody instantiates a new ListAnalysisDataTypesOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAnalysisDataTypesOutputBodyWithDefaults

`func NewListAnalysisDataTypesOutputBodyWithDefaults() *ListAnalysisDataTypesOutputBody`

NewListAnalysisDataTypesOutputBodyWithDefaults instantiates a new ListAnalysisDataTypesOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListAnalysisDataTypesOutputBody) GetItems() []DataTypeEntry`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListAnalysisDataTypesOutputBody) GetItemsOk() (*[]DataTypeEntry, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListAnalysisDataTypesOutputBody) SetItems(v []DataTypeEntry)`

SetItems sets Items field to given value.


### SetItemsNil

`func (o *ListAnalysisDataTypesOutputBody) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *ListAnalysisDataTypesOutputBody) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetTotalCount

`func (o *ListAnalysisDataTypesOutputBody) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListAnalysisDataTypesOutputBody) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListAnalysisDataTypesOutputBody) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


