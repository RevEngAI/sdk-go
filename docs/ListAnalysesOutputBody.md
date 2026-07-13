# ListAnalysesOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | Opaque cursor to fetch the next page; empty on the last page | [optional] 
**PageSize** | **int64** | Number of results in this page | 
**Results** | [**[]AnalysisRecordBody**](AnalysisRecordBody.md) | The page of matching analyses | 

## Methods

### NewListAnalysesOutputBody

`func NewListAnalysesOutputBody(pageSize int64, results []AnalysisRecordBody, ) *ListAnalysesOutputBody`

NewListAnalysesOutputBody instantiates a new ListAnalysesOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAnalysesOutputBodyWithDefaults

`func NewListAnalysesOutputBodyWithDefaults() *ListAnalysesOutputBody`

NewListAnalysesOutputBodyWithDefaults instantiates a new ListAnalysesOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *ListAnalysesOutputBody) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListAnalysesOutputBody) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListAnalysesOutputBody) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListAnalysesOutputBody) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetPageSize

`func (o *ListAnalysesOutputBody) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListAnalysesOutputBody) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListAnalysesOutputBody) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetResults

`func (o *ListAnalysesOutputBody) GetResults() []AnalysisRecordBody`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListAnalysesOutputBody) GetResultsOk() (*[]AnalysisRecordBody, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListAnalysesOutputBody) SetResults(v []AnalysisRecordBody)`

SetResults sets Results field to given value.


### SetResultsNil

`func (o *ListAnalysesOutputBody) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *ListAnalysesOutputBody) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


