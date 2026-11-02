# ListSpendOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNextPage** | **bool** |  | 
**Results** | [**[]SpendBody**](SpendBody.md) |  | 
**TotalCount** | **int64** | Spend records across all pages. | 

## Methods

### NewListSpendOutputBody

`func NewListSpendOutputBody(hasNextPage bool, results []SpendBody, totalCount int64, ) *ListSpendOutputBody`

NewListSpendOutputBody instantiates a new ListSpendOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSpendOutputBodyWithDefaults

`func NewListSpendOutputBodyWithDefaults() *ListSpendOutputBody`

NewListSpendOutputBodyWithDefaults instantiates a new ListSpendOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNextPage

`func (o *ListSpendOutputBody) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *ListSpendOutputBody) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *ListSpendOutputBody) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.


### GetResults

`func (o *ListSpendOutputBody) GetResults() []SpendBody`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListSpendOutputBody) GetResultsOk() (*[]SpendBody, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListSpendOutputBody) SetResults(v []SpendBody)`

SetResults sets Results field to given value.


### GetTotalCount

`func (o *ListSpendOutputBody) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListSpendOutputBody) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListSpendOutputBody) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


