# AutoUnstripRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinSimilarity** | Pointer to **float32** | Minimum similarity expected for a match as a percentage, default is 90 | [optional] [default to 90.0]
**Apply** | Pointer to **bool** | Whether to apply the matched function names to the target binary, default is False | [optional] [default to false]
**ConfidenceThreshold** | Pointer to **float32** | Confidence threshold for applying function names as a percentage, default is 90 | [optional] [default to 90.0]
**MinGroupSize** | Pointer to **int32** | Minimum number of matching functions required to consider for a match, default is 10 | [optional] [default to 10]
**StatusOnly** | Pointer to **bool** | If set to true, only returns the status of the auto-unstrip operation without the actual results | [optional] [default to false]
**NoCache** | Pointer to **bool** | If set to true, forces the system to bypass any cached results and perform a fresh computation | [optional] [default to false]
**UseCanonicalNames** | Pointer to **bool** | Whether to use canonical function names during matching for auto-unstrip, default is False | [optional] [default to false]

## Methods

### NewAutoUnstripRequest

`func NewAutoUnstripRequest() *AutoUnstripRequest`

NewAutoUnstripRequest instantiates a new AutoUnstripRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoUnstripRequestWithDefaults

`func NewAutoUnstripRequestWithDefaults() *AutoUnstripRequest`

NewAutoUnstripRequestWithDefaults instantiates a new AutoUnstripRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinSimilarity

`func (o *AutoUnstripRequest) GetMinSimilarity() float32`

GetMinSimilarity returns the MinSimilarity field if non-nil, zero value otherwise.

### GetMinSimilarityOk

`func (o *AutoUnstripRequest) GetMinSimilarityOk() (*float32, bool)`

GetMinSimilarityOk returns a tuple with the MinSimilarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSimilarity

`func (o *AutoUnstripRequest) SetMinSimilarity(v float32)`

SetMinSimilarity sets MinSimilarity field to given value.

### HasMinSimilarity

`func (o *AutoUnstripRequest) HasMinSimilarity() bool`

HasMinSimilarity returns a boolean if a field has been set.

### GetApply

`func (o *AutoUnstripRequest) GetApply() bool`

GetApply returns the Apply field if non-nil, zero value otherwise.

### GetApplyOk

`func (o *AutoUnstripRequest) GetApplyOk() (*bool, bool)`

GetApplyOk returns a tuple with the Apply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApply

`func (o *AutoUnstripRequest) SetApply(v bool)`

SetApply sets Apply field to given value.

### HasApply

`func (o *AutoUnstripRequest) HasApply() bool`

HasApply returns a boolean if a field has been set.

### GetConfidenceThreshold

`func (o *AutoUnstripRequest) GetConfidenceThreshold() float32`

GetConfidenceThreshold returns the ConfidenceThreshold field if non-nil, zero value otherwise.

### GetConfidenceThresholdOk

`func (o *AutoUnstripRequest) GetConfidenceThresholdOk() (*float32, bool)`

GetConfidenceThresholdOk returns a tuple with the ConfidenceThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidenceThreshold

`func (o *AutoUnstripRequest) SetConfidenceThreshold(v float32)`

SetConfidenceThreshold sets ConfidenceThreshold field to given value.

### HasConfidenceThreshold

`func (o *AutoUnstripRequest) HasConfidenceThreshold() bool`

HasConfidenceThreshold returns a boolean if a field has been set.

### GetMinGroupSize

`func (o *AutoUnstripRequest) GetMinGroupSize() int32`

GetMinGroupSize returns the MinGroupSize field if non-nil, zero value otherwise.

### GetMinGroupSizeOk

`func (o *AutoUnstripRequest) GetMinGroupSizeOk() (*int32, bool)`

GetMinGroupSizeOk returns a tuple with the MinGroupSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinGroupSize

`func (o *AutoUnstripRequest) SetMinGroupSize(v int32)`

SetMinGroupSize sets MinGroupSize field to given value.

### HasMinGroupSize

`func (o *AutoUnstripRequest) HasMinGroupSize() bool`

HasMinGroupSize returns a boolean if a field has been set.

### GetStatusOnly

`func (o *AutoUnstripRequest) GetStatusOnly() bool`

GetStatusOnly returns the StatusOnly field if non-nil, zero value otherwise.

### GetStatusOnlyOk

`func (o *AutoUnstripRequest) GetStatusOnlyOk() (*bool, bool)`

GetStatusOnlyOk returns a tuple with the StatusOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusOnly

`func (o *AutoUnstripRequest) SetStatusOnly(v bool)`

SetStatusOnly sets StatusOnly field to given value.

### HasStatusOnly

`func (o *AutoUnstripRequest) HasStatusOnly() bool`

HasStatusOnly returns a boolean if a field has been set.

### GetNoCache

`func (o *AutoUnstripRequest) GetNoCache() bool`

GetNoCache returns the NoCache field if non-nil, zero value otherwise.

### GetNoCacheOk

`func (o *AutoUnstripRequest) GetNoCacheOk() (*bool, bool)`

GetNoCacheOk returns a tuple with the NoCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoCache

`func (o *AutoUnstripRequest) SetNoCache(v bool)`

SetNoCache sets NoCache field to given value.

### HasNoCache

`func (o *AutoUnstripRequest) HasNoCache() bool`

HasNoCache returns a boolean if a field has been set.

### GetUseCanonicalNames

`func (o *AutoUnstripRequest) GetUseCanonicalNames() bool`

GetUseCanonicalNames returns the UseCanonicalNames field if non-nil, zero value otherwise.

### GetUseCanonicalNamesOk

`func (o *AutoUnstripRequest) GetUseCanonicalNamesOk() (*bool, bool)`

GetUseCanonicalNamesOk returns a tuple with the UseCanonicalNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCanonicalNames

`func (o *AutoUnstripRequest) SetUseCanonicalNames(v bool)`

SetUseCanonicalNames sets UseCanonicalNames field to given value.

### HasUseCanonicalNames

`func (o *AutoUnstripRequest) HasUseCanonicalNames() bool`

HasUseCanonicalNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


