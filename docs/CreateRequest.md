# CreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisConfig** | Pointer to [**Config**](Config.md) |  | [optional] 
**AnalysisScope** | Pointer to **string** |  | [optional] [default to "PRIVATE"]
**AutoRunAgents** | Pointer to [**AutoRunAgents**](AutoRunAgents.md) |  | [optional] 
**BinaryConfig** | Pointer to [**BinaryConfig**](BinaryConfig.md) |  | [optional] 
**DebugHash** | Pointer to **string** |  | [optional] 
**Filename** | **string** |  | 
**Sha256Hash** | **string** |  | 
**Symbols** | Pointer to [**Symbols**](Symbols.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateRequest

`func NewCreateRequest(filename string, sha256Hash string, ) *CreateRequest`

NewCreateRequest instantiates a new CreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRequestWithDefaults

`func NewCreateRequestWithDefaults() *CreateRequest`

NewCreateRequestWithDefaults instantiates a new CreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisConfig

`func (o *CreateRequest) GetAnalysisConfig() Config`

GetAnalysisConfig returns the AnalysisConfig field if non-nil, zero value otherwise.

### GetAnalysisConfigOk

`func (o *CreateRequest) GetAnalysisConfigOk() (*Config, bool)`

GetAnalysisConfigOk returns a tuple with the AnalysisConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisConfig

`func (o *CreateRequest) SetAnalysisConfig(v Config)`

SetAnalysisConfig sets AnalysisConfig field to given value.

### HasAnalysisConfig

`func (o *CreateRequest) HasAnalysisConfig() bool`

HasAnalysisConfig returns a boolean if a field has been set.

### GetAnalysisScope

`func (o *CreateRequest) GetAnalysisScope() string`

GetAnalysisScope returns the AnalysisScope field if non-nil, zero value otherwise.

### GetAnalysisScopeOk

`func (o *CreateRequest) GetAnalysisScopeOk() (*string, bool)`

GetAnalysisScopeOk returns a tuple with the AnalysisScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisScope

`func (o *CreateRequest) SetAnalysisScope(v string)`

SetAnalysisScope sets AnalysisScope field to given value.

### HasAnalysisScope

`func (o *CreateRequest) HasAnalysisScope() bool`

HasAnalysisScope returns a boolean if a field has been set.

### GetAutoRunAgents

`func (o *CreateRequest) GetAutoRunAgents() AutoRunAgents`

GetAutoRunAgents returns the AutoRunAgents field if non-nil, zero value otherwise.

### GetAutoRunAgentsOk

`func (o *CreateRequest) GetAutoRunAgentsOk() (*AutoRunAgents, bool)`

GetAutoRunAgentsOk returns a tuple with the AutoRunAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRunAgents

`func (o *CreateRequest) SetAutoRunAgents(v AutoRunAgents)`

SetAutoRunAgents sets AutoRunAgents field to given value.

### HasAutoRunAgents

`func (o *CreateRequest) HasAutoRunAgents() bool`

HasAutoRunAgents returns a boolean if a field has been set.

### GetBinaryConfig

`func (o *CreateRequest) GetBinaryConfig() BinaryConfig`

GetBinaryConfig returns the BinaryConfig field if non-nil, zero value otherwise.

### GetBinaryConfigOk

`func (o *CreateRequest) GetBinaryConfigOk() (*BinaryConfig, bool)`

GetBinaryConfigOk returns a tuple with the BinaryConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryConfig

`func (o *CreateRequest) SetBinaryConfig(v BinaryConfig)`

SetBinaryConfig sets BinaryConfig field to given value.

### HasBinaryConfig

`func (o *CreateRequest) HasBinaryConfig() bool`

HasBinaryConfig returns a boolean if a field has been set.

### GetDebugHash

`func (o *CreateRequest) GetDebugHash() string`

GetDebugHash returns the DebugHash field if non-nil, zero value otherwise.

### GetDebugHashOk

`func (o *CreateRequest) GetDebugHashOk() (*string, bool)`

GetDebugHashOk returns a tuple with the DebugHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugHash

`func (o *CreateRequest) SetDebugHash(v string)`

SetDebugHash sets DebugHash field to given value.

### HasDebugHash

`func (o *CreateRequest) HasDebugHash() bool`

HasDebugHash returns a boolean if a field has been set.

### GetFilename

`func (o *CreateRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *CreateRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *CreateRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetSha256Hash

`func (o *CreateRequest) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *CreateRequest) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *CreateRequest) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.


### GetSymbols

`func (o *CreateRequest) GetSymbols() Symbols`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *CreateRequest) GetSymbolsOk() (*Symbols, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *CreateRequest) SetSymbols(v Symbols)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *CreateRequest) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.

### GetTags

`func (o *CreateRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CreateRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CreateRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


