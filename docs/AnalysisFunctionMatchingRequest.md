# AnalysisFunctionMatchingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinSimilarity** | Pointer to **float32** | Minimum similarity expected for a match as a percentage, default is 90 | [optional] [default to 90.0]
**Filters** | Pointer to [**NullableFunctionMatchingFilters**](FunctionMatchingFilters.md) |  | [optional] 
**ResultsPerFunction** | Pointer to **int32** | Maximum number of matches to return per function, default is 1, max is 10 | [optional] [default to 1]
**Page** | Pointer to **int32** | Page number for paginated results, default is 1 (first page) | [optional] [default to 1]
**PageSize** | Pointer to **int32** | Number of functions to return per page, default is 0 (all functions), max is 1000 | [optional] [default to 0]
**StatusOnly** | Pointer to **bool** | If set to true, only returns the status of the matching operation without the actual results | [optional] [default to false]
**NoCache** | Pointer to **bool** | If set to true, forces the system to bypass any cached results and perform a fresh computation | [optional] [default to false]
**UseCanonicalNames** | Pointer to **bool** | Whether to use canonical function names during function matching for confidence results, default is False | [optional] [default to false]

## Methods

### NewAnalysisFunctionMatchingRequest

`func NewAnalysisFunctionMatchingRequest() *AnalysisFunctionMatchingRequest`

NewAnalysisFunctionMatchingRequest instantiates a new AnalysisFunctionMatchingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisFunctionMatchingRequestWithDefaults

`func NewAnalysisFunctionMatchingRequestWithDefaults() *AnalysisFunctionMatchingRequest`

NewAnalysisFunctionMatchingRequestWithDefaults instantiates a new AnalysisFunctionMatchingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinSimilarity

`func (o *AnalysisFunctionMatchingRequest) GetMinSimilarity() float32`

GetMinSimilarity returns the MinSimilarity field if non-nil, zero value otherwise.

### GetMinSimilarityOk

`func (o *AnalysisFunctionMatchingRequest) GetMinSimilarityOk() (*float32, bool)`

GetMinSimilarityOk returns a tuple with the MinSimilarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSimilarity

`func (o *AnalysisFunctionMatchingRequest) SetMinSimilarity(v float32)`

SetMinSimilarity sets MinSimilarity field to given value.

### HasMinSimilarity

`func (o *AnalysisFunctionMatchingRequest) HasMinSimilarity() bool`

HasMinSimilarity returns a boolean if a field has been set.

### GetFilters

`func (o *AnalysisFunctionMatchingRequest) GetFilters() FunctionMatchingFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *AnalysisFunctionMatchingRequest) GetFiltersOk() (*FunctionMatchingFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *AnalysisFunctionMatchingRequest) SetFilters(v FunctionMatchingFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *AnalysisFunctionMatchingRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *AnalysisFunctionMatchingRequest) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *AnalysisFunctionMatchingRequest) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetResultsPerFunction

`func (o *AnalysisFunctionMatchingRequest) GetResultsPerFunction() int32`

GetResultsPerFunction returns the ResultsPerFunction field if non-nil, zero value otherwise.

### GetResultsPerFunctionOk

`func (o *AnalysisFunctionMatchingRequest) GetResultsPerFunctionOk() (*int32, bool)`

GetResultsPerFunctionOk returns a tuple with the ResultsPerFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsPerFunction

`func (o *AnalysisFunctionMatchingRequest) SetResultsPerFunction(v int32)`

SetResultsPerFunction sets ResultsPerFunction field to given value.

### HasResultsPerFunction

`func (o *AnalysisFunctionMatchingRequest) HasResultsPerFunction() bool`

HasResultsPerFunction returns a boolean if a field has been set.

### GetPage

`func (o *AnalysisFunctionMatchingRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *AnalysisFunctionMatchingRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *AnalysisFunctionMatchingRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *AnalysisFunctionMatchingRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *AnalysisFunctionMatchingRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *AnalysisFunctionMatchingRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *AnalysisFunctionMatchingRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *AnalysisFunctionMatchingRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetStatusOnly

`func (o *AnalysisFunctionMatchingRequest) GetStatusOnly() bool`

GetStatusOnly returns the StatusOnly field if non-nil, zero value otherwise.

### GetStatusOnlyOk

`func (o *AnalysisFunctionMatchingRequest) GetStatusOnlyOk() (*bool, bool)`

GetStatusOnlyOk returns a tuple with the StatusOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusOnly

`func (o *AnalysisFunctionMatchingRequest) SetStatusOnly(v bool)`

SetStatusOnly sets StatusOnly field to given value.

### HasStatusOnly

`func (o *AnalysisFunctionMatchingRequest) HasStatusOnly() bool`

HasStatusOnly returns a boolean if a field has been set.

### GetNoCache

`func (o *AnalysisFunctionMatchingRequest) GetNoCache() bool`

GetNoCache returns the NoCache field if non-nil, zero value otherwise.

### GetNoCacheOk

`func (o *AnalysisFunctionMatchingRequest) GetNoCacheOk() (*bool, bool)`

GetNoCacheOk returns a tuple with the NoCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoCache

`func (o *AnalysisFunctionMatchingRequest) SetNoCache(v bool)`

SetNoCache sets NoCache field to given value.

### HasNoCache

`func (o *AnalysisFunctionMatchingRequest) HasNoCache() bool`

HasNoCache returns a boolean if a field has been set.

### GetUseCanonicalNames

`func (o *AnalysisFunctionMatchingRequest) GetUseCanonicalNames() bool`

GetUseCanonicalNames returns the UseCanonicalNames field if non-nil, zero value otherwise.

### GetUseCanonicalNamesOk

`func (o *AnalysisFunctionMatchingRequest) GetUseCanonicalNamesOk() (*bool, bool)`

GetUseCanonicalNamesOk returns a tuple with the UseCanonicalNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCanonicalNames

`func (o *AnalysisFunctionMatchingRequest) SetUseCanonicalNames(v bool)`

SetUseCanonicalNames sets UseCanonicalNames field to given value.

### HasUseCanonicalNames

`func (o *AnalysisFunctionMatchingRequest) HasUseCanonicalNames() bool`

HasUseCanonicalNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


