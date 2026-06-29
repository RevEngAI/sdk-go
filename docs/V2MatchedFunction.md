# V2MatchedFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionId** | **int64** | Unique identifier of the matched function | 
**BinaryId** | **int32** |  | 
**FunctionName** | **string** |  | 
**FunctionVaddr** | **int64** |  | 
**MangledName** | **string** |  | 
**Debug** | **bool** |  | 
**BinaryName** | **string** |  | 
**Sha256Hash** | **string** |  | 
**AnalysisId** | **int32** |  | 
**Similarity** | Pointer to **NullableFloat32** |  | [optional] 
**Confidence** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewV2MatchedFunction

`func NewV2MatchedFunction(functionId int64, binaryId int32, functionName string, functionVaddr int64, mangledName string, debug bool, binaryName string, sha256Hash string, analysisId int32, ) *V2MatchedFunction`

NewV2MatchedFunction instantiates a new V2MatchedFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MatchedFunctionWithDefaults

`func NewV2MatchedFunctionWithDefaults() *V2MatchedFunction`

NewV2MatchedFunctionWithDefaults instantiates a new V2MatchedFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionId

`func (o *V2MatchedFunction) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *V2MatchedFunction) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *V2MatchedFunction) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetBinaryId

`func (o *V2MatchedFunction) GetBinaryId() int32`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *V2MatchedFunction) GetBinaryIdOk() (*int32, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *V2MatchedFunction) SetBinaryId(v int32)`

SetBinaryId sets BinaryId field to given value.


### GetFunctionName

`func (o *V2MatchedFunction) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *V2MatchedFunction) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *V2MatchedFunction) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetFunctionVaddr

`func (o *V2MatchedFunction) GetFunctionVaddr() int64`

GetFunctionVaddr returns the FunctionVaddr field if non-nil, zero value otherwise.

### GetFunctionVaddrOk

`func (o *V2MatchedFunction) GetFunctionVaddrOk() (*int64, bool)`

GetFunctionVaddrOk returns a tuple with the FunctionVaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionVaddr

`func (o *V2MatchedFunction) SetFunctionVaddr(v int64)`

SetFunctionVaddr sets FunctionVaddr field to given value.


### GetMangledName

`func (o *V2MatchedFunction) GetMangledName() string`

GetMangledName returns the MangledName field if non-nil, zero value otherwise.

### GetMangledNameOk

`func (o *V2MatchedFunction) GetMangledNameOk() (*string, bool)`

GetMangledNameOk returns a tuple with the MangledName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMangledName

`func (o *V2MatchedFunction) SetMangledName(v string)`

SetMangledName sets MangledName field to given value.


### GetDebug

`func (o *V2MatchedFunction) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *V2MatchedFunction) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *V2MatchedFunction) SetDebug(v bool)`

SetDebug sets Debug field to given value.


### GetBinaryName

`func (o *V2MatchedFunction) GetBinaryName() string`

GetBinaryName returns the BinaryName field if non-nil, zero value otherwise.

### GetBinaryNameOk

`func (o *V2MatchedFunction) GetBinaryNameOk() (*string, bool)`

GetBinaryNameOk returns a tuple with the BinaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryName

`func (o *V2MatchedFunction) SetBinaryName(v string)`

SetBinaryName sets BinaryName field to given value.


### GetSha256Hash

`func (o *V2MatchedFunction) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *V2MatchedFunction) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *V2MatchedFunction) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.


### GetAnalysisId

`func (o *V2MatchedFunction) GetAnalysisId() int32`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *V2MatchedFunction) GetAnalysisIdOk() (*int32, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *V2MatchedFunction) SetAnalysisId(v int32)`

SetAnalysisId sets AnalysisId field to given value.


### GetSimilarity

`func (o *V2MatchedFunction) GetSimilarity() float32`

GetSimilarity returns the Similarity field if non-nil, zero value otherwise.

### GetSimilarityOk

`func (o *V2MatchedFunction) GetSimilarityOk() (*float32, bool)`

GetSimilarityOk returns a tuple with the Similarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimilarity

`func (o *V2MatchedFunction) SetSimilarity(v float32)`

SetSimilarity sets Similarity field to given value.

### HasSimilarity

`func (o *V2MatchedFunction) HasSimilarity() bool`

HasSimilarity returns a boolean if a field has been set.

### SetSimilarityNil

`func (o *V2MatchedFunction) SetSimilarityNil(b bool)`

 SetSimilarityNil sets the value for Similarity to be an explicit nil

### UnsetSimilarity
`func (o *V2MatchedFunction) UnsetSimilarity()`

UnsetSimilarity ensures that no value is present for Similarity, not even an explicit nil
### GetConfidence

`func (o *V2MatchedFunction) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *V2MatchedFunction) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *V2MatchedFunction) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *V2MatchedFunction) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### SetConfidenceNil

`func (o *V2MatchedFunction) SetConfidenceNil(b bool)`

 SetConfidenceNil sets the value for Confidence to be an explicit nil

### UnsetConfidence
`func (o *V2MatchedFunction) UnsetConfidence()`

UnsetConfidence ensures that no value is present for Confidence, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


