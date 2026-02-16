# AnalysisCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | **string** | The name of the file | 
**Sha256Hash** | **string** | The name of the file | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | List of community tags to assign to an analysis | [optional] [default to {}]
**AnalysisScope** | Pointer to [**AnalysisScope**](AnalysisScope.md) | The scope of the analysis determines who can access it | [optional] [default to ANALYSISSCOPE_PRIVATE]
**Symbols** | Pointer to [**NullableSymbols**](Symbols.md) |  | [optional] 
**DebugHash** | Pointer to **NullableString** |  | [optional] 
**AnalysisConfig** | Pointer to [**AnalysisConfig**](AnalysisConfig.md) | The analysis config enables the configuration of optional analysis stages | [optional] 
**BinaryConfig** | Pointer to [**BinaryConfig**](BinaryConfig.md) | The binary config can override automatically determined values such as ISA, Platform, File Format, etc | [optional] 

## Methods

### NewAnalysisCreateRequest

`func NewAnalysisCreateRequest(filename string, sha256Hash string, ) *AnalysisCreateRequest`

NewAnalysisCreateRequest instantiates a new AnalysisCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisCreateRequestWithDefaults

`func NewAnalysisCreateRequestWithDefaults() *AnalysisCreateRequest`

NewAnalysisCreateRequestWithDefaults instantiates a new AnalysisCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *AnalysisCreateRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *AnalysisCreateRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *AnalysisCreateRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetSha256Hash

`func (o *AnalysisCreateRequest) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *AnalysisCreateRequest) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *AnalysisCreateRequest) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.


### GetTags

`func (o *AnalysisCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AnalysisCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AnalysisCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AnalysisCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAnalysisScope

`func (o *AnalysisCreateRequest) GetAnalysisScope() AnalysisScope`

GetAnalysisScope returns the AnalysisScope field if non-nil, zero value otherwise.

### GetAnalysisScopeOk

`func (o *AnalysisCreateRequest) GetAnalysisScopeOk() (*AnalysisScope, bool)`

GetAnalysisScopeOk returns a tuple with the AnalysisScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisScope

`func (o *AnalysisCreateRequest) SetAnalysisScope(v AnalysisScope)`

SetAnalysisScope sets AnalysisScope field to given value.

### HasAnalysisScope

`func (o *AnalysisCreateRequest) HasAnalysisScope() bool`

HasAnalysisScope returns a boolean if a field has been set.

### GetSymbols

`func (o *AnalysisCreateRequest) GetSymbols() Symbols`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *AnalysisCreateRequest) GetSymbolsOk() (*Symbols, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *AnalysisCreateRequest) SetSymbols(v Symbols)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *AnalysisCreateRequest) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.

### SetSymbolsNil

`func (o *AnalysisCreateRequest) SetSymbolsNil(b bool)`

 SetSymbolsNil sets the value for Symbols to be an explicit nil

### UnsetSymbols
`func (o *AnalysisCreateRequest) UnsetSymbols()`

UnsetSymbols ensures that no value is present for Symbols, not even an explicit nil
### GetDebugHash

`func (o *AnalysisCreateRequest) GetDebugHash() string`

GetDebugHash returns the DebugHash field if non-nil, zero value otherwise.

### GetDebugHashOk

`func (o *AnalysisCreateRequest) GetDebugHashOk() (*string, bool)`

GetDebugHashOk returns a tuple with the DebugHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugHash

`func (o *AnalysisCreateRequest) SetDebugHash(v string)`

SetDebugHash sets DebugHash field to given value.

### HasDebugHash

`func (o *AnalysisCreateRequest) HasDebugHash() bool`

HasDebugHash returns a boolean if a field has been set.

### SetDebugHashNil

`func (o *AnalysisCreateRequest) SetDebugHashNil(b bool)`

 SetDebugHashNil sets the value for DebugHash to be an explicit nil

### UnsetDebugHash
`func (o *AnalysisCreateRequest) UnsetDebugHash()`

UnsetDebugHash ensures that no value is present for DebugHash, not even an explicit nil
### GetAnalysisConfig

`func (o *AnalysisCreateRequest) GetAnalysisConfig() AnalysisConfig`

GetAnalysisConfig returns the AnalysisConfig field if non-nil, zero value otherwise.

### GetAnalysisConfigOk

`func (o *AnalysisCreateRequest) GetAnalysisConfigOk() (*AnalysisConfig, bool)`

GetAnalysisConfigOk returns a tuple with the AnalysisConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisConfig

`func (o *AnalysisCreateRequest) SetAnalysisConfig(v AnalysisConfig)`

SetAnalysisConfig sets AnalysisConfig field to given value.

### HasAnalysisConfig

`func (o *AnalysisCreateRequest) HasAnalysisConfig() bool`

HasAnalysisConfig returns a boolean if a field has been set.

### GetBinaryConfig

`func (o *AnalysisCreateRequest) GetBinaryConfig() BinaryConfig`

GetBinaryConfig returns the BinaryConfig field if non-nil, zero value otherwise.

### GetBinaryConfigOk

`func (o *AnalysisCreateRequest) GetBinaryConfigOk() (*BinaryConfig, bool)`

GetBinaryConfigOk returns a tuple with the BinaryConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryConfig

`func (o *AnalysisCreateRequest) SetBinaryConfig(v BinaryConfig)`

SetBinaryConfig sets BinaryConfig field to given value.

### HasBinaryConfig

`func (o *AnalysisCreateRequest) HasBinaryConfig() bool`

HasBinaryConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


