# ListAnalysisFunctionsDataTypesOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]DataTypesEntry**](DataTypesEntry.md) |  | 
**TotalCount** | **int64** | Total functions in the analysis, ignoring pagination. | 

## Methods

### NewListAnalysisFunctionsDataTypesOutputBody

`func NewListAnalysisFunctionsDataTypesOutputBody(items []DataTypesEntry, totalCount int64, ) *ListAnalysisFunctionsDataTypesOutputBody`

NewListAnalysisFunctionsDataTypesOutputBody instantiates a new ListAnalysisFunctionsDataTypesOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAnalysisFunctionsDataTypesOutputBodyWithDefaults

`func NewListAnalysisFunctionsDataTypesOutputBodyWithDefaults() *ListAnalysisFunctionsDataTypesOutputBody`

NewListAnalysisFunctionsDataTypesOutputBodyWithDefaults instantiates a new ListAnalysisFunctionsDataTypesOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListAnalysisFunctionsDataTypesOutputBody) GetItems() []DataTypesEntry`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListAnalysisFunctionsDataTypesOutputBody) GetItemsOk() (*[]DataTypesEntry, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListAnalysisFunctionsDataTypesOutputBody) SetItems(v []DataTypesEntry)`

SetItems sets Items field to given value.


### SetItemsNil

`func (o *ListAnalysisFunctionsDataTypesOutputBody) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *ListAnalysisFunctionsDataTypesOutputBody) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetTotalCount

`func (o *ListAnalysisFunctionsDataTypesOutputBody) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListAnalysisFunctionsDataTypesOutputBody) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListAnalysisFunctionsDataTypesOutputBody) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


