# AnalysisBasicInfoOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisScope** | **string** | PUBLIC, PRIVATE, or TEAM | 
**BaseAddress** | **NullableInt64** | Base address of the binary, null when unknown | 
**BinaryId** | **int64** | Binary ID | 
**BinaryName** | **string** | Binary filename | 
**BinarySize** | **int64** | Binary size in bytes | 
**BinaryUuid** | Pointer to **string** | UUID of the binary, omitted when not set | [optional] 
**Creation** | **time.Time** | When the binary was uploaded | 
**Debug** | **bool** | True when the binary was analysed with debug symbols | 
**FunctionCount** | **int64** | Number of functions in the binary | 
**IsAdvanced** | **bool** | True when the analysis was run in advanced mode | 
**IsOwner** | **bool** | True when the caller is the analysis owner | 
**IsSystem** | **bool** | True when the analysis is owned by a system user | 
**ModelId** | **int64** | Model ID | 
**ModelName** | **string** | Model used for analysis | 
**OwnerUsername** | **string** | Username of the analysis owner | 
**SequencerVersion** | Pointer to **string** | Sequencer version, omitted when not set | [optional] 
**Sha256Hash** | **string** | SHA-256 hash of the binary | 
**TeamId** | **int64** | Team ID of the analysis | 

## Methods

### NewAnalysisBasicInfoOutputBody

`func NewAnalysisBasicInfoOutputBody(analysisScope string, baseAddress NullableInt64, binaryId int64, binaryName string, binarySize int64, creation time.Time, debug bool, functionCount int64, isAdvanced bool, isOwner bool, isSystem bool, modelId int64, modelName string, ownerUsername string, sha256Hash string, teamId int64, ) *AnalysisBasicInfoOutputBody`

NewAnalysisBasicInfoOutputBody instantiates a new AnalysisBasicInfoOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisBasicInfoOutputBodyWithDefaults

`func NewAnalysisBasicInfoOutputBodyWithDefaults() *AnalysisBasicInfoOutputBody`

NewAnalysisBasicInfoOutputBodyWithDefaults instantiates a new AnalysisBasicInfoOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisScope

`func (o *AnalysisBasicInfoOutputBody) GetAnalysisScope() string`

GetAnalysisScope returns the AnalysisScope field if non-nil, zero value otherwise.

### GetAnalysisScopeOk

`func (o *AnalysisBasicInfoOutputBody) GetAnalysisScopeOk() (*string, bool)`

GetAnalysisScopeOk returns a tuple with the AnalysisScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisScope

`func (o *AnalysisBasicInfoOutputBody) SetAnalysisScope(v string)`

SetAnalysisScope sets AnalysisScope field to given value.


### GetBaseAddress

`func (o *AnalysisBasicInfoOutputBody) GetBaseAddress() int64`

GetBaseAddress returns the BaseAddress field if non-nil, zero value otherwise.

### GetBaseAddressOk

`func (o *AnalysisBasicInfoOutputBody) GetBaseAddressOk() (*int64, bool)`

GetBaseAddressOk returns a tuple with the BaseAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAddress

`func (o *AnalysisBasicInfoOutputBody) SetBaseAddress(v int64)`

SetBaseAddress sets BaseAddress field to given value.


### SetBaseAddressNil

`func (o *AnalysisBasicInfoOutputBody) SetBaseAddressNil(b bool)`

 SetBaseAddressNil sets the value for BaseAddress to be an explicit nil

### UnsetBaseAddress
`func (o *AnalysisBasicInfoOutputBody) UnsetBaseAddress()`

UnsetBaseAddress ensures that no value is present for BaseAddress, not even an explicit nil
### GetBinaryId

`func (o *AnalysisBasicInfoOutputBody) GetBinaryId() int64`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *AnalysisBasicInfoOutputBody) GetBinaryIdOk() (*int64, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *AnalysisBasicInfoOutputBody) SetBinaryId(v int64)`

SetBinaryId sets BinaryId field to given value.


### GetBinaryName

`func (o *AnalysisBasicInfoOutputBody) GetBinaryName() string`

GetBinaryName returns the BinaryName field if non-nil, zero value otherwise.

### GetBinaryNameOk

`func (o *AnalysisBasicInfoOutputBody) GetBinaryNameOk() (*string, bool)`

GetBinaryNameOk returns a tuple with the BinaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryName

`func (o *AnalysisBasicInfoOutputBody) SetBinaryName(v string)`

SetBinaryName sets BinaryName field to given value.


### GetBinarySize

`func (o *AnalysisBasicInfoOutputBody) GetBinarySize() int64`

GetBinarySize returns the BinarySize field if non-nil, zero value otherwise.

### GetBinarySizeOk

`func (o *AnalysisBasicInfoOutputBody) GetBinarySizeOk() (*int64, bool)`

GetBinarySizeOk returns a tuple with the BinarySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinarySize

`func (o *AnalysisBasicInfoOutputBody) SetBinarySize(v int64)`

SetBinarySize sets BinarySize field to given value.


### GetBinaryUuid

`func (o *AnalysisBasicInfoOutputBody) GetBinaryUuid() string`

GetBinaryUuid returns the BinaryUuid field if non-nil, zero value otherwise.

### GetBinaryUuidOk

`func (o *AnalysisBasicInfoOutputBody) GetBinaryUuidOk() (*string, bool)`

GetBinaryUuidOk returns a tuple with the BinaryUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryUuid

`func (o *AnalysisBasicInfoOutputBody) SetBinaryUuid(v string)`

SetBinaryUuid sets BinaryUuid field to given value.

### HasBinaryUuid

`func (o *AnalysisBasicInfoOutputBody) HasBinaryUuid() bool`

HasBinaryUuid returns a boolean if a field has been set.

### GetCreation

`func (o *AnalysisBasicInfoOutputBody) GetCreation() time.Time`

GetCreation returns the Creation field if non-nil, zero value otherwise.

### GetCreationOk

`func (o *AnalysisBasicInfoOutputBody) GetCreationOk() (*time.Time, bool)`

GetCreationOk returns a tuple with the Creation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreation

`func (o *AnalysisBasicInfoOutputBody) SetCreation(v time.Time)`

SetCreation sets Creation field to given value.


### GetDebug

`func (o *AnalysisBasicInfoOutputBody) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *AnalysisBasicInfoOutputBody) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *AnalysisBasicInfoOutputBody) SetDebug(v bool)`

SetDebug sets Debug field to given value.


### GetFunctionCount

`func (o *AnalysisBasicInfoOutputBody) GetFunctionCount() int64`

GetFunctionCount returns the FunctionCount field if non-nil, zero value otherwise.

### GetFunctionCountOk

`func (o *AnalysisBasicInfoOutputBody) GetFunctionCountOk() (*int64, bool)`

GetFunctionCountOk returns a tuple with the FunctionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionCount

`func (o *AnalysisBasicInfoOutputBody) SetFunctionCount(v int64)`

SetFunctionCount sets FunctionCount field to given value.


### GetIsAdvanced

`func (o *AnalysisBasicInfoOutputBody) GetIsAdvanced() bool`

GetIsAdvanced returns the IsAdvanced field if non-nil, zero value otherwise.

### GetIsAdvancedOk

`func (o *AnalysisBasicInfoOutputBody) GetIsAdvancedOk() (*bool, bool)`

GetIsAdvancedOk returns a tuple with the IsAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdvanced

`func (o *AnalysisBasicInfoOutputBody) SetIsAdvanced(v bool)`

SetIsAdvanced sets IsAdvanced field to given value.


### GetIsOwner

`func (o *AnalysisBasicInfoOutputBody) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *AnalysisBasicInfoOutputBody) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *AnalysisBasicInfoOutputBody) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.


### GetIsSystem

`func (o *AnalysisBasicInfoOutputBody) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *AnalysisBasicInfoOutputBody) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *AnalysisBasicInfoOutputBody) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.


### GetModelId

`func (o *AnalysisBasicInfoOutputBody) GetModelId() int64`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *AnalysisBasicInfoOutputBody) GetModelIdOk() (*int64, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *AnalysisBasicInfoOutputBody) SetModelId(v int64)`

SetModelId sets ModelId field to given value.


### GetModelName

`func (o *AnalysisBasicInfoOutputBody) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *AnalysisBasicInfoOutputBody) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *AnalysisBasicInfoOutputBody) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetOwnerUsername

`func (o *AnalysisBasicInfoOutputBody) GetOwnerUsername() string`

GetOwnerUsername returns the OwnerUsername field if non-nil, zero value otherwise.

### GetOwnerUsernameOk

`func (o *AnalysisBasicInfoOutputBody) GetOwnerUsernameOk() (*string, bool)`

GetOwnerUsernameOk returns a tuple with the OwnerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUsername

`func (o *AnalysisBasicInfoOutputBody) SetOwnerUsername(v string)`

SetOwnerUsername sets OwnerUsername field to given value.


### GetSequencerVersion

`func (o *AnalysisBasicInfoOutputBody) GetSequencerVersion() string`

GetSequencerVersion returns the SequencerVersion field if non-nil, zero value otherwise.

### GetSequencerVersionOk

`func (o *AnalysisBasicInfoOutputBody) GetSequencerVersionOk() (*string, bool)`

GetSequencerVersionOk returns a tuple with the SequencerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequencerVersion

`func (o *AnalysisBasicInfoOutputBody) SetSequencerVersion(v string)`

SetSequencerVersion sets SequencerVersion field to given value.

### HasSequencerVersion

`func (o *AnalysisBasicInfoOutputBody) HasSequencerVersion() bool`

HasSequencerVersion returns a boolean if a field has been set.

### GetSha256Hash

`func (o *AnalysisBasicInfoOutputBody) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *AnalysisBasicInfoOutputBody) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *AnalysisBasicInfoOutputBody) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.


### GetTeamId

`func (o *AnalysisBasicInfoOutputBody) GetTeamId() int64`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *AnalysisBasicInfoOutputBody) GetTeamIdOk() (*int64, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *AnalysisBasicInfoOutputBody) SetTeamId(v int64)`

SetTeamId sets TeamId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


