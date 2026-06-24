# StartMatchingForAnalysisInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**MatchFilters**](MatchFilters.md) | Narrow the candidate pool. | [optional] 
**MinSimilarity** | Pointer to **float64** | Similarity floor as a percentage. Defaults to 90. | [optional] 
**ResultsPerFunction** | Pointer to **int64** | Max matches returned per source function. Defaults to 1. | [optional] 

## Methods

### NewStartMatchingForAnalysisInputBody

`func NewStartMatchingForAnalysisInputBody() *StartMatchingForAnalysisInputBody`

NewStartMatchingForAnalysisInputBody instantiates a new StartMatchingForAnalysisInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartMatchingForAnalysisInputBodyWithDefaults

`func NewStartMatchingForAnalysisInputBodyWithDefaults() *StartMatchingForAnalysisInputBody`

NewStartMatchingForAnalysisInputBodyWithDefaults instantiates a new StartMatchingForAnalysisInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *StartMatchingForAnalysisInputBody) GetFilters() MatchFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *StartMatchingForAnalysisInputBody) GetFiltersOk() (*MatchFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *StartMatchingForAnalysisInputBody) SetFilters(v MatchFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *StartMatchingForAnalysisInputBody) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetMinSimilarity

`func (o *StartMatchingForAnalysisInputBody) GetMinSimilarity() float64`

GetMinSimilarity returns the MinSimilarity field if non-nil, zero value otherwise.

### GetMinSimilarityOk

`func (o *StartMatchingForAnalysisInputBody) GetMinSimilarityOk() (*float64, bool)`

GetMinSimilarityOk returns a tuple with the MinSimilarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSimilarity

`func (o *StartMatchingForAnalysisInputBody) SetMinSimilarity(v float64)`

SetMinSimilarity sets MinSimilarity field to given value.

### HasMinSimilarity

`func (o *StartMatchingForAnalysisInputBody) HasMinSimilarity() bool`

HasMinSimilarity returns a boolean if a field has been set.

### GetResultsPerFunction

`func (o *StartMatchingForAnalysisInputBody) GetResultsPerFunction() int64`

GetResultsPerFunction returns the ResultsPerFunction field if non-nil, zero value otherwise.

### GetResultsPerFunctionOk

`func (o *StartMatchingForAnalysisInputBody) GetResultsPerFunctionOk() (*int64, bool)`

GetResultsPerFunctionOk returns a tuple with the ResultsPerFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsPerFunction

`func (o *StartMatchingForAnalysisInputBody) SetResultsPerFunction(v int64)`

SetResultsPerFunction sets ResultsPerFunction field to given value.

### HasResultsPerFunction

`func (o *StartMatchingForAnalysisInputBody) HasResultsPerFunction() bool`

HasResultsPerFunction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


