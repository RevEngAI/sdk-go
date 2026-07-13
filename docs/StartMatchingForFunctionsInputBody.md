# StartMatchingForFunctionsInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**MatchFilters**](MatchFilters.md) | Narrow the candidate pool. | [optional] 
**FunctionIds** | **[]int64** | Source function IDs to match against the rest of the corpus. | 
**MinSimilarity** | Pointer to **float64** | Similarity floor as a percentage. Defaults to 90. | [optional] 
**NoCache** | Pointer to **bool** | By default a completed matching run for the same request is reused (response status&#x3D;COMPLETED, no new run). Set true to force a fresh run. | [optional] 
**ResultsPerFunction** | Pointer to **int64** | Max matches returned per source function. Defaults to 1. | [optional] 
**UseCanonicalNames** | Pointer to **bool** | Collapse near-duplicate candidate names into canonical buckets and return per-name confidences (the response &#39;confidences&#39; array). Adds a canonicalisation step; defaults to false. | [optional] 

## Methods

### NewStartMatchingForFunctionsInputBody

`func NewStartMatchingForFunctionsInputBody(functionIds []int64, ) *StartMatchingForFunctionsInputBody`

NewStartMatchingForFunctionsInputBody instantiates a new StartMatchingForFunctionsInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartMatchingForFunctionsInputBodyWithDefaults

`func NewStartMatchingForFunctionsInputBodyWithDefaults() *StartMatchingForFunctionsInputBody`

NewStartMatchingForFunctionsInputBodyWithDefaults instantiates a new StartMatchingForFunctionsInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *StartMatchingForFunctionsInputBody) GetFilters() MatchFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *StartMatchingForFunctionsInputBody) GetFiltersOk() (*MatchFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *StartMatchingForFunctionsInputBody) SetFilters(v MatchFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *StartMatchingForFunctionsInputBody) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetFunctionIds

`func (o *StartMatchingForFunctionsInputBody) GetFunctionIds() []int64`

GetFunctionIds returns the FunctionIds field if non-nil, zero value otherwise.

### GetFunctionIdsOk

`func (o *StartMatchingForFunctionsInputBody) GetFunctionIdsOk() (*[]int64, bool)`

GetFunctionIdsOk returns a tuple with the FunctionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionIds

`func (o *StartMatchingForFunctionsInputBody) SetFunctionIds(v []int64)`

SetFunctionIds sets FunctionIds field to given value.


### SetFunctionIdsNil

`func (o *StartMatchingForFunctionsInputBody) SetFunctionIdsNil(b bool)`

 SetFunctionIdsNil sets the value for FunctionIds to be an explicit nil

### UnsetFunctionIds
`func (o *StartMatchingForFunctionsInputBody) UnsetFunctionIds()`

UnsetFunctionIds ensures that no value is present for FunctionIds, not even an explicit nil
### GetMinSimilarity

`func (o *StartMatchingForFunctionsInputBody) GetMinSimilarity() float64`

GetMinSimilarity returns the MinSimilarity field if non-nil, zero value otherwise.

### GetMinSimilarityOk

`func (o *StartMatchingForFunctionsInputBody) GetMinSimilarityOk() (*float64, bool)`

GetMinSimilarityOk returns a tuple with the MinSimilarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSimilarity

`func (o *StartMatchingForFunctionsInputBody) SetMinSimilarity(v float64)`

SetMinSimilarity sets MinSimilarity field to given value.

### HasMinSimilarity

`func (o *StartMatchingForFunctionsInputBody) HasMinSimilarity() bool`

HasMinSimilarity returns a boolean if a field has been set.

### GetNoCache

`func (o *StartMatchingForFunctionsInputBody) GetNoCache() bool`

GetNoCache returns the NoCache field if non-nil, zero value otherwise.

### GetNoCacheOk

`func (o *StartMatchingForFunctionsInputBody) GetNoCacheOk() (*bool, bool)`

GetNoCacheOk returns a tuple with the NoCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoCache

`func (o *StartMatchingForFunctionsInputBody) SetNoCache(v bool)`

SetNoCache sets NoCache field to given value.

### HasNoCache

`func (o *StartMatchingForFunctionsInputBody) HasNoCache() bool`

HasNoCache returns a boolean if a field has been set.

### GetResultsPerFunction

`func (o *StartMatchingForFunctionsInputBody) GetResultsPerFunction() int64`

GetResultsPerFunction returns the ResultsPerFunction field if non-nil, zero value otherwise.

### GetResultsPerFunctionOk

`func (o *StartMatchingForFunctionsInputBody) GetResultsPerFunctionOk() (*int64, bool)`

GetResultsPerFunctionOk returns a tuple with the ResultsPerFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsPerFunction

`func (o *StartMatchingForFunctionsInputBody) SetResultsPerFunction(v int64)`

SetResultsPerFunction sets ResultsPerFunction field to given value.

### HasResultsPerFunction

`func (o *StartMatchingForFunctionsInputBody) HasResultsPerFunction() bool`

HasResultsPerFunction returns a boolean if a field has been set.

### GetUseCanonicalNames

`func (o *StartMatchingForFunctionsInputBody) GetUseCanonicalNames() bool`

GetUseCanonicalNames returns the UseCanonicalNames field if non-nil, zero value otherwise.

### GetUseCanonicalNamesOk

`func (o *StartMatchingForFunctionsInputBody) GetUseCanonicalNamesOk() (*bool, bool)`

GetUseCanonicalNamesOk returns a tuple with the UseCanonicalNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCanonicalNames

`func (o *StartMatchingForFunctionsInputBody) SetUseCanonicalNames(v bool)`

SetUseCanonicalNames sets UseCanonicalNames field to given value.

### HasUseCanonicalNames

`func (o *StartMatchingForFunctionsInputBody) HasUseCanonicalNames() bool`

HasUseCanonicalNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


