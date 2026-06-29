# MatchedFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisId** | **int64** | Analysis the candidate&#39;s binary belongs to | 
**BinaryId** | **int64** | Binary the candidate belongs to | 
**BinaryName** | **string** | Binary name | 
**Confidence** | **float64** | Softmax-normalised confidence over the candidate pool | 
**Debug** | **bool** | Whether the candidate&#39;s name came from debug info | 
**FunctionId** | **int64** | Candidate function ID | 
**FunctionName** | **string** | Candidate function name | 
**FunctionVaddr** | **int64** | Candidate&#39;s virtual address inside its binary | 
**MangledName** | **string** | Mangled name when available | 
**Sha256Hash** | **string** | SHA-256 of the candidate&#39;s binary | 
**Similarity** | **float64** | Cosine similarity scaled to a percentage | 

## Methods

### NewMatchedFunction

`func NewMatchedFunction(analysisId int64, binaryId int64, binaryName string, confidence float64, debug bool, functionId int64, functionName string, functionVaddr int64, mangledName string, sha256Hash string, similarity float64, ) *MatchedFunction`

NewMatchedFunction instantiates a new MatchedFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchedFunctionWithDefaults

`func NewMatchedFunctionWithDefaults() *MatchedFunction`

NewMatchedFunctionWithDefaults instantiates a new MatchedFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisId

`func (o *MatchedFunction) GetAnalysisId() int64`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *MatchedFunction) GetAnalysisIdOk() (*int64, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *MatchedFunction) SetAnalysisId(v int64)`

SetAnalysisId sets AnalysisId field to given value.


### GetBinaryId

`func (o *MatchedFunction) GetBinaryId() int64`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *MatchedFunction) GetBinaryIdOk() (*int64, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *MatchedFunction) SetBinaryId(v int64)`

SetBinaryId sets BinaryId field to given value.


### GetBinaryName

`func (o *MatchedFunction) GetBinaryName() string`

GetBinaryName returns the BinaryName field if non-nil, zero value otherwise.

### GetBinaryNameOk

`func (o *MatchedFunction) GetBinaryNameOk() (*string, bool)`

GetBinaryNameOk returns a tuple with the BinaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryName

`func (o *MatchedFunction) SetBinaryName(v string)`

SetBinaryName sets BinaryName field to given value.


### GetConfidence

`func (o *MatchedFunction) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *MatchedFunction) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *MatchedFunction) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.


### GetDebug

`func (o *MatchedFunction) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *MatchedFunction) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *MatchedFunction) SetDebug(v bool)`

SetDebug sets Debug field to given value.


### GetFunctionId

`func (o *MatchedFunction) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *MatchedFunction) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *MatchedFunction) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetFunctionName

`func (o *MatchedFunction) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *MatchedFunction) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *MatchedFunction) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetFunctionVaddr

`func (o *MatchedFunction) GetFunctionVaddr() int64`

GetFunctionVaddr returns the FunctionVaddr field if non-nil, zero value otherwise.

### GetFunctionVaddrOk

`func (o *MatchedFunction) GetFunctionVaddrOk() (*int64, bool)`

GetFunctionVaddrOk returns a tuple with the FunctionVaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionVaddr

`func (o *MatchedFunction) SetFunctionVaddr(v int64)`

SetFunctionVaddr sets FunctionVaddr field to given value.


### GetMangledName

`func (o *MatchedFunction) GetMangledName() string`

GetMangledName returns the MangledName field if non-nil, zero value otherwise.

### GetMangledNameOk

`func (o *MatchedFunction) GetMangledNameOk() (*string, bool)`

GetMangledNameOk returns a tuple with the MangledName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMangledName

`func (o *MatchedFunction) SetMangledName(v string)`

SetMangledName sets MangledName field to given value.


### GetSha256Hash

`func (o *MatchedFunction) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *MatchedFunction) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *MatchedFunction) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.


### GetSimilarity

`func (o *MatchedFunction) GetSimilarity() float64`

GetSimilarity returns the Similarity field if non-nil, zero value otherwise.

### GetSimilarityOk

`func (o *MatchedFunction) GetSimilarityOk() (*float64, bool)`

GetSimilarityOk returns a tuple with the Similarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimilarity

`func (o *MatchedFunction) SetSimilarity(v float64)`

SetSimilarity sets Similarity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


