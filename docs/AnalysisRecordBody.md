# AnalysisRecordBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisId** | **int64** | Analysis ID | 
**AnalysisScope** | **string** | Scope of the analysis | 
**BaseAddress** | **int64** | Binary base address | 
**BinaryId** | **int64** | Binary ID | 
**BinaryName** | **string** | Binary filename | 
**BinarySize** | **int64** | Binary size in bytes | 
**Creation** | **time.Time** | When the analysis was created | 
**FunctionBoundariesHash** | **string** | Hash of the binary&#39;s provided function boundaries | 
**IsOwner** | **bool** | True when the caller owns the analysis | 
**ModelId** | **int64** | Model ID | 
**ModelName** | **string** | Model name | 
**Sha256Hash** | **string** | SHA-256 hash of the binary | 
**Status** | **string** | Analysis status | 
**Tags** | [**[]AnalysisTagBody**](AnalysisTagBody.md) | Tags associated with the binary | 
**Username** | **string** | Username of the analysis owner | 

## Methods

### NewAnalysisRecordBody

`func NewAnalysisRecordBody(analysisId int64, analysisScope string, baseAddress int64, binaryId int64, binaryName string, binarySize int64, creation time.Time, functionBoundariesHash string, isOwner bool, modelId int64, modelName string, sha256Hash string, status string, tags []AnalysisTagBody, username string, ) *AnalysisRecordBody`

NewAnalysisRecordBody instantiates a new AnalysisRecordBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisRecordBodyWithDefaults

`func NewAnalysisRecordBodyWithDefaults() *AnalysisRecordBody`

NewAnalysisRecordBodyWithDefaults instantiates a new AnalysisRecordBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisId

`func (o *AnalysisRecordBody) GetAnalysisId() int64`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *AnalysisRecordBody) GetAnalysisIdOk() (*int64, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *AnalysisRecordBody) SetAnalysisId(v int64)`

SetAnalysisId sets AnalysisId field to given value.


### GetAnalysisScope

`func (o *AnalysisRecordBody) GetAnalysisScope() string`

GetAnalysisScope returns the AnalysisScope field if non-nil, zero value otherwise.

### GetAnalysisScopeOk

`func (o *AnalysisRecordBody) GetAnalysisScopeOk() (*string, bool)`

GetAnalysisScopeOk returns a tuple with the AnalysisScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisScope

`func (o *AnalysisRecordBody) SetAnalysisScope(v string)`

SetAnalysisScope sets AnalysisScope field to given value.


### GetBaseAddress

`func (o *AnalysisRecordBody) GetBaseAddress() int64`

GetBaseAddress returns the BaseAddress field if non-nil, zero value otherwise.

### GetBaseAddressOk

`func (o *AnalysisRecordBody) GetBaseAddressOk() (*int64, bool)`

GetBaseAddressOk returns a tuple with the BaseAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAddress

`func (o *AnalysisRecordBody) SetBaseAddress(v int64)`

SetBaseAddress sets BaseAddress field to given value.


### GetBinaryId

`func (o *AnalysisRecordBody) GetBinaryId() int64`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *AnalysisRecordBody) GetBinaryIdOk() (*int64, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *AnalysisRecordBody) SetBinaryId(v int64)`

SetBinaryId sets BinaryId field to given value.


### GetBinaryName

`func (o *AnalysisRecordBody) GetBinaryName() string`

GetBinaryName returns the BinaryName field if non-nil, zero value otherwise.

### GetBinaryNameOk

`func (o *AnalysisRecordBody) GetBinaryNameOk() (*string, bool)`

GetBinaryNameOk returns a tuple with the BinaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryName

`func (o *AnalysisRecordBody) SetBinaryName(v string)`

SetBinaryName sets BinaryName field to given value.


### GetBinarySize

`func (o *AnalysisRecordBody) GetBinarySize() int64`

GetBinarySize returns the BinarySize field if non-nil, zero value otherwise.

### GetBinarySizeOk

`func (o *AnalysisRecordBody) GetBinarySizeOk() (*int64, bool)`

GetBinarySizeOk returns a tuple with the BinarySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinarySize

`func (o *AnalysisRecordBody) SetBinarySize(v int64)`

SetBinarySize sets BinarySize field to given value.


### GetCreation

`func (o *AnalysisRecordBody) GetCreation() time.Time`

GetCreation returns the Creation field if non-nil, zero value otherwise.

### GetCreationOk

`func (o *AnalysisRecordBody) GetCreationOk() (*time.Time, bool)`

GetCreationOk returns a tuple with the Creation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreation

`func (o *AnalysisRecordBody) SetCreation(v time.Time)`

SetCreation sets Creation field to given value.


### GetFunctionBoundariesHash

`func (o *AnalysisRecordBody) GetFunctionBoundariesHash() string`

GetFunctionBoundariesHash returns the FunctionBoundariesHash field if non-nil, zero value otherwise.

### GetFunctionBoundariesHashOk

`func (o *AnalysisRecordBody) GetFunctionBoundariesHashOk() (*string, bool)`

GetFunctionBoundariesHashOk returns a tuple with the FunctionBoundariesHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionBoundariesHash

`func (o *AnalysisRecordBody) SetFunctionBoundariesHash(v string)`

SetFunctionBoundariesHash sets FunctionBoundariesHash field to given value.


### GetIsOwner

`func (o *AnalysisRecordBody) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *AnalysisRecordBody) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *AnalysisRecordBody) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.


### GetModelId

`func (o *AnalysisRecordBody) GetModelId() int64`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *AnalysisRecordBody) GetModelIdOk() (*int64, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *AnalysisRecordBody) SetModelId(v int64)`

SetModelId sets ModelId field to given value.


### GetModelName

`func (o *AnalysisRecordBody) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *AnalysisRecordBody) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *AnalysisRecordBody) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetSha256Hash

`func (o *AnalysisRecordBody) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *AnalysisRecordBody) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *AnalysisRecordBody) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.


### GetStatus

`func (o *AnalysisRecordBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AnalysisRecordBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AnalysisRecordBody) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTags

`func (o *AnalysisRecordBody) GetTags() []AnalysisTagBody`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AnalysisRecordBody) GetTagsOk() (*[]AnalysisTagBody, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AnalysisRecordBody) SetTags(v []AnalysisTagBody)`

SetTags sets Tags field to given value.


### SetTagsNil

`func (o *AnalysisRecordBody) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AnalysisRecordBody) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetUsername

`func (o *AnalysisRecordBody) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AnalysisRecordBody) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AnalysisRecordBody) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


