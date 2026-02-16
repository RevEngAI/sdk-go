# FunctionsDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionId** | **int64** | Function id | 
**FunctionName** | **string** |  | 
**FunctionNameMangled** | **string** |  | 
**FunctionVaddr** | **int64** |  | 
**FunctionSize** | **int32** |  | 
**AnalysisId** | **int32** |  | 
**BinaryId** | **int32** |  | 
**BinaryName** | **string** |  | 
**Sha256Hash** | **string** |  | 
**DebugHash** | **NullableString** |  | 
**Debug** | **bool** |  | 
**Embedding3d** | Pointer to **[]float32** |  | [optional] 
**Embedding1d** | Pointer to **[]float32** |  | [optional] 

## Methods

### NewFunctionsDetailResponse

`func NewFunctionsDetailResponse(functionId int64, functionName string, functionNameMangled string, functionVaddr int64, functionSize int32, analysisId int32, binaryId int32, binaryName string, sha256Hash string, debugHash NullableString, debug bool, ) *FunctionsDetailResponse`

NewFunctionsDetailResponse instantiates a new FunctionsDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionsDetailResponseWithDefaults

`func NewFunctionsDetailResponseWithDefaults() *FunctionsDetailResponse`

NewFunctionsDetailResponseWithDefaults instantiates a new FunctionsDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionId

`func (o *FunctionsDetailResponse) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *FunctionsDetailResponse) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *FunctionsDetailResponse) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetFunctionName

`func (o *FunctionsDetailResponse) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *FunctionsDetailResponse) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *FunctionsDetailResponse) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetFunctionNameMangled

`func (o *FunctionsDetailResponse) GetFunctionNameMangled() string`

GetFunctionNameMangled returns the FunctionNameMangled field if non-nil, zero value otherwise.

### GetFunctionNameMangledOk

`func (o *FunctionsDetailResponse) GetFunctionNameMangledOk() (*string, bool)`

GetFunctionNameMangledOk returns a tuple with the FunctionNameMangled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionNameMangled

`func (o *FunctionsDetailResponse) SetFunctionNameMangled(v string)`

SetFunctionNameMangled sets FunctionNameMangled field to given value.


### GetFunctionVaddr

`func (o *FunctionsDetailResponse) GetFunctionVaddr() int64`

GetFunctionVaddr returns the FunctionVaddr field if non-nil, zero value otherwise.

### GetFunctionVaddrOk

`func (o *FunctionsDetailResponse) GetFunctionVaddrOk() (*int64, bool)`

GetFunctionVaddrOk returns a tuple with the FunctionVaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionVaddr

`func (o *FunctionsDetailResponse) SetFunctionVaddr(v int64)`

SetFunctionVaddr sets FunctionVaddr field to given value.


### GetFunctionSize

`func (o *FunctionsDetailResponse) GetFunctionSize() int32`

GetFunctionSize returns the FunctionSize field if non-nil, zero value otherwise.

### GetFunctionSizeOk

`func (o *FunctionsDetailResponse) GetFunctionSizeOk() (*int32, bool)`

GetFunctionSizeOk returns a tuple with the FunctionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionSize

`func (o *FunctionsDetailResponse) SetFunctionSize(v int32)`

SetFunctionSize sets FunctionSize field to given value.


### GetAnalysisId

`func (o *FunctionsDetailResponse) GetAnalysisId() int32`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *FunctionsDetailResponse) GetAnalysisIdOk() (*int32, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *FunctionsDetailResponse) SetAnalysisId(v int32)`

SetAnalysisId sets AnalysisId field to given value.


### GetBinaryId

`func (o *FunctionsDetailResponse) GetBinaryId() int32`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *FunctionsDetailResponse) GetBinaryIdOk() (*int32, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *FunctionsDetailResponse) SetBinaryId(v int32)`

SetBinaryId sets BinaryId field to given value.


### GetBinaryName

`func (o *FunctionsDetailResponse) GetBinaryName() string`

GetBinaryName returns the BinaryName field if non-nil, zero value otherwise.

### GetBinaryNameOk

`func (o *FunctionsDetailResponse) GetBinaryNameOk() (*string, bool)`

GetBinaryNameOk returns a tuple with the BinaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryName

`func (o *FunctionsDetailResponse) SetBinaryName(v string)`

SetBinaryName sets BinaryName field to given value.


### GetSha256Hash

`func (o *FunctionsDetailResponse) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *FunctionsDetailResponse) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *FunctionsDetailResponse) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.


### GetDebugHash

`func (o *FunctionsDetailResponse) GetDebugHash() string`

GetDebugHash returns the DebugHash field if non-nil, zero value otherwise.

### GetDebugHashOk

`func (o *FunctionsDetailResponse) GetDebugHashOk() (*string, bool)`

GetDebugHashOk returns a tuple with the DebugHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugHash

`func (o *FunctionsDetailResponse) SetDebugHash(v string)`

SetDebugHash sets DebugHash field to given value.


### SetDebugHashNil

`func (o *FunctionsDetailResponse) SetDebugHashNil(b bool)`

 SetDebugHashNil sets the value for DebugHash to be an explicit nil

### UnsetDebugHash
`func (o *FunctionsDetailResponse) UnsetDebugHash()`

UnsetDebugHash ensures that no value is present for DebugHash, not even an explicit nil
### GetDebug

`func (o *FunctionsDetailResponse) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *FunctionsDetailResponse) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *FunctionsDetailResponse) SetDebug(v bool)`

SetDebug sets Debug field to given value.


### GetEmbedding3d

`func (o *FunctionsDetailResponse) GetEmbedding3d() []float32`

GetEmbedding3d returns the Embedding3d field if non-nil, zero value otherwise.

### GetEmbedding3dOk

`func (o *FunctionsDetailResponse) GetEmbedding3dOk() (*[]float32, bool)`

GetEmbedding3dOk returns a tuple with the Embedding3d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedding3d

`func (o *FunctionsDetailResponse) SetEmbedding3d(v []float32)`

SetEmbedding3d sets Embedding3d field to given value.

### HasEmbedding3d

`func (o *FunctionsDetailResponse) HasEmbedding3d() bool`

HasEmbedding3d returns a boolean if a field has been set.

### SetEmbedding3dNil

`func (o *FunctionsDetailResponse) SetEmbedding3dNil(b bool)`

 SetEmbedding3dNil sets the value for Embedding3d to be an explicit nil

### UnsetEmbedding3d
`func (o *FunctionsDetailResponse) UnsetEmbedding3d()`

UnsetEmbedding3d ensures that no value is present for Embedding3d, not even an explicit nil
### GetEmbedding1d

`func (o *FunctionsDetailResponse) GetEmbedding1d() []float32`

GetEmbedding1d returns the Embedding1d field if non-nil, zero value otherwise.

### GetEmbedding1dOk

`func (o *FunctionsDetailResponse) GetEmbedding1dOk() (*[]float32, bool)`

GetEmbedding1dOk returns a tuple with the Embedding1d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedding1d

`func (o *FunctionsDetailResponse) SetEmbedding1d(v []float32)`

SetEmbedding1d sets Embedding1d field to given value.

### HasEmbedding1d

`func (o *FunctionsDetailResponse) HasEmbedding1d() bool`

HasEmbedding1d returns a boolean if a field has been set.

### SetEmbedding1dNil

`func (o *FunctionsDetailResponse) SetEmbedding1dNil(b bool)`

 SetEmbedding1dNil sets the value for Embedding1d to be an explicit nil

### UnsetEmbedding1d
`func (o *FunctionsDetailResponse) UnsetEmbedding1d()`

UnsetEmbedding1d ensures that no value is present for Embedding1d, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


