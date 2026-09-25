# SearchBinariesOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNextPage** | **bool** |  | 
**PageNumber** | **int64** |  | 
**PageSize** | **int64** |  | 
**Results** | [**[]BinarySearchResultBody**](BinarySearchResultBody.md) |  | 

## Methods

### NewSearchBinariesOutputBody

`func NewSearchBinariesOutputBody(hasNextPage bool, pageNumber int64, pageSize int64, results []BinarySearchResultBody, ) *SearchBinariesOutputBody`

NewSearchBinariesOutputBody instantiates a new SearchBinariesOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchBinariesOutputBodyWithDefaults

`func NewSearchBinariesOutputBodyWithDefaults() *SearchBinariesOutputBody`

NewSearchBinariesOutputBodyWithDefaults instantiates a new SearchBinariesOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNextPage

`func (o *SearchBinariesOutputBody) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *SearchBinariesOutputBody) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *SearchBinariesOutputBody) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.


### GetPageNumber

`func (o *SearchBinariesOutputBody) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *SearchBinariesOutputBody) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *SearchBinariesOutputBody) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *SearchBinariesOutputBody) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *SearchBinariesOutputBody) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *SearchBinariesOutputBody) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetResults

`func (o *SearchBinariesOutputBody) GetResults() []BinarySearchResultBody`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SearchBinariesOutputBody) GetResultsOk() (*[]BinarySearchResultBody, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SearchBinariesOutputBody) SetResults(v []BinarySearchResultBody)`

SetResults sets Results field to given value.


### SetResultsNil

`func (o *SearchBinariesOutputBody) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *SearchBinariesOutputBody) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


