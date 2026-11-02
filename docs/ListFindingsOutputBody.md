# ListFindingsOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNextPage** | **bool** |  | 
**Results** | [**[]FindingBody**](FindingBody.md) |  | 
**TotalCount** | **int64** | Findings matching the filters across all pages. | 

## Methods

### NewListFindingsOutputBody

`func NewListFindingsOutputBody(hasNextPage bool, results []FindingBody, totalCount int64, ) *ListFindingsOutputBody`

NewListFindingsOutputBody instantiates a new ListFindingsOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFindingsOutputBodyWithDefaults

`func NewListFindingsOutputBodyWithDefaults() *ListFindingsOutputBody`

NewListFindingsOutputBodyWithDefaults instantiates a new ListFindingsOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNextPage

`func (o *ListFindingsOutputBody) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *ListFindingsOutputBody) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *ListFindingsOutputBody) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.


### GetResults

`func (o *ListFindingsOutputBody) GetResults() []FindingBody`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListFindingsOutputBody) GetResultsOk() (*[]FindingBody, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListFindingsOutputBody) SetResults(v []FindingBody)`

SetResults sets Results field to given value.


### GetTotalCount

`func (o *ListFindingsOutputBody) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListFindingsOutputBody) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListFindingsOutputBody) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


