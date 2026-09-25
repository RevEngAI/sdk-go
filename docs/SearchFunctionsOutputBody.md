# SearchFunctionsOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNextPage** | **bool** |  | 
**PageNumber** | **int64** |  | 
**PageSize** | **int64** |  | 
**Results** | [**[]FunctionSearchResultBody**](FunctionSearchResultBody.md) |  | 

## Methods

### NewSearchFunctionsOutputBody

`func NewSearchFunctionsOutputBody(hasNextPage bool, pageNumber int64, pageSize int64, results []FunctionSearchResultBody, ) *SearchFunctionsOutputBody`

NewSearchFunctionsOutputBody instantiates a new SearchFunctionsOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchFunctionsOutputBodyWithDefaults

`func NewSearchFunctionsOutputBodyWithDefaults() *SearchFunctionsOutputBody`

NewSearchFunctionsOutputBodyWithDefaults instantiates a new SearchFunctionsOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNextPage

`func (o *SearchFunctionsOutputBody) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *SearchFunctionsOutputBody) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *SearchFunctionsOutputBody) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.


### GetPageNumber

`func (o *SearchFunctionsOutputBody) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *SearchFunctionsOutputBody) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *SearchFunctionsOutputBody) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *SearchFunctionsOutputBody) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *SearchFunctionsOutputBody) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *SearchFunctionsOutputBody) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetResults

`func (o *SearchFunctionsOutputBody) GetResults() []FunctionSearchResultBody`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SearchFunctionsOutputBody) GetResultsOk() (*[]FunctionSearchResultBody, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SearchFunctionsOutputBody) SetResults(v []FunctionSearchResultBody)`

SetResults sets Results field to given value.


### SetResultsNil

`func (o *SearchFunctionsOutputBody) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *SearchFunctionsOutputBody) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


