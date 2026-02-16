# TagSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]TagSearchResult**](TagSearchResult.md) | The results of the search | 

## Methods

### NewTagSearchResponse

`func NewTagSearchResponse(results []TagSearchResult, ) *TagSearchResponse`

NewTagSearchResponse instantiates a new TagSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagSearchResponseWithDefaults

`func NewTagSearchResponseWithDefaults() *TagSearchResponse`

NewTagSearchResponseWithDefaults instantiates a new TagSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *TagSearchResponse) GetResults() []TagSearchResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *TagSearchResponse) GetResultsOk() (*[]TagSearchResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *TagSearchResponse) SetResults(v []TagSearchResult)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


