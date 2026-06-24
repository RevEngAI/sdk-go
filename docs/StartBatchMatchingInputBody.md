# StartBatchMatchingInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinaryIds** | **[]int64** | Binary IDs to match the analysis against, one workflow per binary. | 
**DebugTypes** | Pointer to **[]string** | Restrict matches to candidates with these debug source types. Defaults to [\&quot;SYSTEM\&quot;]. | [optional] 
**MinSimilarity** | Pointer to **float64** | Similarity floor as a percentage. Defaults to 90. | [optional] 
**ResultsPerFunction** | Pointer to **int64** | Max matches returned per source function. Defaults to 1. | [optional] 

## Methods

### NewStartBatchMatchingInputBody

`func NewStartBatchMatchingInputBody(binaryIds []int64, ) *StartBatchMatchingInputBody`

NewStartBatchMatchingInputBody instantiates a new StartBatchMatchingInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartBatchMatchingInputBodyWithDefaults

`func NewStartBatchMatchingInputBodyWithDefaults() *StartBatchMatchingInputBody`

NewStartBatchMatchingInputBodyWithDefaults instantiates a new StartBatchMatchingInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaryIds

`func (o *StartBatchMatchingInputBody) GetBinaryIds() []int64`

GetBinaryIds returns the BinaryIds field if non-nil, zero value otherwise.

### GetBinaryIdsOk

`func (o *StartBatchMatchingInputBody) GetBinaryIdsOk() (*[]int64, bool)`

GetBinaryIdsOk returns a tuple with the BinaryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryIds

`func (o *StartBatchMatchingInputBody) SetBinaryIds(v []int64)`

SetBinaryIds sets BinaryIds field to given value.


### SetBinaryIdsNil

`func (o *StartBatchMatchingInputBody) SetBinaryIdsNil(b bool)`

 SetBinaryIdsNil sets the value for BinaryIds to be an explicit nil

### UnsetBinaryIds
`func (o *StartBatchMatchingInputBody) UnsetBinaryIds()`

UnsetBinaryIds ensures that no value is present for BinaryIds, not even an explicit nil
### GetDebugTypes

`func (o *StartBatchMatchingInputBody) GetDebugTypes() []*string`

GetDebugTypes returns the DebugTypes field if non-nil, zero value otherwise.

### GetDebugTypesOk

`func (o *StartBatchMatchingInputBody) GetDebugTypesOk() (*[]*string, bool)`

GetDebugTypesOk returns a tuple with the DebugTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugTypes

`func (o *StartBatchMatchingInputBody) SetDebugTypes(v []*string)`

SetDebugTypes sets DebugTypes field to given value.

### HasDebugTypes

`func (o *StartBatchMatchingInputBody) HasDebugTypes() bool`

HasDebugTypes returns a boolean if a field has been set.

### SetDebugTypesNil

`func (o *StartBatchMatchingInputBody) SetDebugTypesNil(b bool)`

 SetDebugTypesNil sets the value for DebugTypes to be an explicit nil

### UnsetDebugTypes
`func (o *StartBatchMatchingInputBody) UnsetDebugTypes()`

UnsetDebugTypes ensures that no value is present for DebugTypes, not even an explicit nil
### GetMinSimilarity

`func (o *StartBatchMatchingInputBody) GetMinSimilarity() float64`

GetMinSimilarity returns the MinSimilarity field if non-nil, zero value otherwise.

### GetMinSimilarityOk

`func (o *StartBatchMatchingInputBody) GetMinSimilarityOk() (*float64, bool)`

GetMinSimilarityOk returns a tuple with the MinSimilarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSimilarity

`func (o *StartBatchMatchingInputBody) SetMinSimilarity(v float64)`

SetMinSimilarity sets MinSimilarity field to given value.

### HasMinSimilarity

`func (o *StartBatchMatchingInputBody) HasMinSimilarity() bool`

HasMinSimilarity returns a boolean if a field has been set.

### GetResultsPerFunction

`func (o *StartBatchMatchingInputBody) GetResultsPerFunction() int64`

GetResultsPerFunction returns the ResultsPerFunction field if non-nil, zero value otherwise.

### GetResultsPerFunctionOk

`func (o *StartBatchMatchingInputBody) GetResultsPerFunctionOk() (*int64, bool)`

GetResultsPerFunctionOk returns a tuple with the ResultsPerFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsPerFunction

`func (o *StartBatchMatchingInputBody) SetResultsPerFunction(v int64)`

SetResultsPerFunction sets ResultsPerFunction field to given value.

### HasResultsPerFunction

`func (o *StartBatchMatchingInputBody) HasResultsPerFunction() bool`

HasResultsPerFunction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


