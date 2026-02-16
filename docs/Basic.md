# Basic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinaryId** | **int32** | The ID of the binary | 
**BinaryName** | **string** | The name of the binary uploaded | 
**BinarySize** | **int32** | The size of the binary uploaded (bytes) | 
**Creation** | **time.Time** | When the binary was uploaded | 
**Sha256Hash** | **string** | The hash of the binary uploaded | 
**ModelName** | **string** | The model name used for analysis | 
**ModelId** | **int32** | The model ID used for analysis | 
**OwnerUsername** | **string** | The name of the owner of the binary | 
**IsSystem** | **bool** | Whether the analysis is a system analysis | 
**AnalysisScope** | **string** | The scope of the analysis | 
**IsOwner** | **bool** | Whether the current user is the owner | 
**Debug** | **bool** | Whether the current analysis was analysed with debug symbols | 
**FunctionCount** | **int32** | The number of functions in the binary | 
**IsAdvanced** | **bool** | Whether the analysis was advanced | 
**BaseAddress** | **NullableInt32** |  | 
**BinaryUuid** | Pointer to **NullableString** |  | [optional] 
**SequencerVersion** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBasic

`func NewBasic(binaryId int32, binaryName string, binarySize int32, creation time.Time, sha256Hash string, modelName string, modelId int32, ownerUsername string, isSystem bool, analysisScope string, isOwner bool, debug bool, functionCount int32, isAdvanced bool, baseAddress NullableInt32, ) *Basic`

NewBasic instantiates a new Basic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicWithDefaults

`func NewBasicWithDefaults() *Basic`

NewBasicWithDefaults instantiates a new Basic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaryId

`func (o *Basic) GetBinaryId() int32`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *Basic) GetBinaryIdOk() (*int32, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *Basic) SetBinaryId(v int32)`

SetBinaryId sets BinaryId field to given value.


### GetBinaryName

`func (o *Basic) GetBinaryName() string`

GetBinaryName returns the BinaryName field if non-nil, zero value otherwise.

### GetBinaryNameOk

`func (o *Basic) GetBinaryNameOk() (*string, bool)`

GetBinaryNameOk returns a tuple with the BinaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryName

`func (o *Basic) SetBinaryName(v string)`

SetBinaryName sets BinaryName field to given value.


### GetBinarySize

`func (o *Basic) GetBinarySize() int32`

GetBinarySize returns the BinarySize field if non-nil, zero value otherwise.

### GetBinarySizeOk

`func (o *Basic) GetBinarySizeOk() (*int32, bool)`

GetBinarySizeOk returns a tuple with the BinarySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinarySize

`func (o *Basic) SetBinarySize(v int32)`

SetBinarySize sets BinarySize field to given value.


### GetCreation

`func (o *Basic) GetCreation() time.Time`

GetCreation returns the Creation field if non-nil, zero value otherwise.

### GetCreationOk

`func (o *Basic) GetCreationOk() (*time.Time, bool)`

GetCreationOk returns a tuple with the Creation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreation

`func (o *Basic) SetCreation(v time.Time)`

SetCreation sets Creation field to given value.


### GetSha256Hash

`func (o *Basic) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *Basic) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *Basic) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.


### GetModelName

`func (o *Basic) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *Basic) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *Basic) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetModelId

`func (o *Basic) GetModelId() int32`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *Basic) GetModelIdOk() (*int32, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *Basic) SetModelId(v int32)`

SetModelId sets ModelId field to given value.


### GetOwnerUsername

`func (o *Basic) GetOwnerUsername() string`

GetOwnerUsername returns the OwnerUsername field if non-nil, zero value otherwise.

### GetOwnerUsernameOk

`func (o *Basic) GetOwnerUsernameOk() (*string, bool)`

GetOwnerUsernameOk returns a tuple with the OwnerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUsername

`func (o *Basic) SetOwnerUsername(v string)`

SetOwnerUsername sets OwnerUsername field to given value.


### GetIsSystem

`func (o *Basic) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *Basic) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *Basic) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.


### GetAnalysisScope

`func (o *Basic) GetAnalysisScope() string`

GetAnalysisScope returns the AnalysisScope field if non-nil, zero value otherwise.

### GetAnalysisScopeOk

`func (o *Basic) GetAnalysisScopeOk() (*string, bool)`

GetAnalysisScopeOk returns a tuple with the AnalysisScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisScope

`func (o *Basic) SetAnalysisScope(v string)`

SetAnalysisScope sets AnalysisScope field to given value.


### GetIsOwner

`func (o *Basic) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *Basic) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *Basic) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.


### GetDebug

`func (o *Basic) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *Basic) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *Basic) SetDebug(v bool)`

SetDebug sets Debug field to given value.


### GetFunctionCount

`func (o *Basic) GetFunctionCount() int32`

GetFunctionCount returns the FunctionCount field if non-nil, zero value otherwise.

### GetFunctionCountOk

`func (o *Basic) GetFunctionCountOk() (*int32, bool)`

GetFunctionCountOk returns a tuple with the FunctionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionCount

`func (o *Basic) SetFunctionCount(v int32)`

SetFunctionCount sets FunctionCount field to given value.


### GetIsAdvanced

`func (o *Basic) GetIsAdvanced() bool`

GetIsAdvanced returns the IsAdvanced field if non-nil, zero value otherwise.

### GetIsAdvancedOk

`func (o *Basic) GetIsAdvancedOk() (*bool, bool)`

GetIsAdvancedOk returns a tuple with the IsAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdvanced

`func (o *Basic) SetIsAdvanced(v bool)`

SetIsAdvanced sets IsAdvanced field to given value.


### GetBaseAddress

`func (o *Basic) GetBaseAddress() int32`

GetBaseAddress returns the BaseAddress field if non-nil, zero value otherwise.

### GetBaseAddressOk

`func (o *Basic) GetBaseAddressOk() (*int32, bool)`

GetBaseAddressOk returns a tuple with the BaseAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAddress

`func (o *Basic) SetBaseAddress(v int32)`

SetBaseAddress sets BaseAddress field to given value.


### SetBaseAddressNil

`func (o *Basic) SetBaseAddressNil(b bool)`

 SetBaseAddressNil sets the value for BaseAddress to be an explicit nil

### UnsetBaseAddress
`func (o *Basic) UnsetBaseAddress()`

UnsetBaseAddress ensures that no value is present for BaseAddress, not even an explicit nil
### GetBinaryUuid

`func (o *Basic) GetBinaryUuid() string`

GetBinaryUuid returns the BinaryUuid field if non-nil, zero value otherwise.

### GetBinaryUuidOk

`func (o *Basic) GetBinaryUuidOk() (*string, bool)`

GetBinaryUuidOk returns a tuple with the BinaryUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryUuid

`func (o *Basic) SetBinaryUuid(v string)`

SetBinaryUuid sets BinaryUuid field to given value.

### HasBinaryUuid

`func (o *Basic) HasBinaryUuid() bool`

HasBinaryUuid returns a boolean if a field has been set.

### SetBinaryUuidNil

`func (o *Basic) SetBinaryUuidNil(b bool)`

 SetBinaryUuidNil sets the value for BinaryUuid to be an explicit nil

### UnsetBinaryUuid
`func (o *Basic) UnsetBinaryUuid()`

UnsetBinaryUuid ensures that no value is present for BinaryUuid, not even an explicit nil
### GetSequencerVersion

`func (o *Basic) GetSequencerVersion() string`

GetSequencerVersion returns the SequencerVersion field if non-nil, zero value otherwise.

### GetSequencerVersionOk

`func (o *Basic) GetSequencerVersionOk() (*string, bool)`

GetSequencerVersionOk returns a tuple with the SequencerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequencerVersion

`func (o *Basic) SetSequencerVersion(v string)`

SetSequencerVersion sets SequencerVersion field to given value.

### HasSequencerVersion

`func (o *Basic) HasSequencerVersion() bool`

HasSequencerVersion returns a boolean if a field has been set.

### SetSequencerVersionNil

`func (o *Basic) SetSequencerVersionNil(b bool)`

 SetSequencerVersionNil sets the value for SequencerVersion to be an explicit nil

### UnsetSequencerVersion
`func (o *Basic) UnsetSequencerVersion()`

UnsetSequencerVersion ensures that no value is present for SequencerVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


