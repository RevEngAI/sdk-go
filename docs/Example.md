# Example

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisId** | **int64** | Analysis ID | 
**AnalysisScope** | **string** | Scope of the analysis (always PUBLIC for examples) | 
**BinaryId** | **int64** | Binary ID | 
**BinaryName** | **string** | Binary filename | 
**BinarySize** | **int64** | Binary size in bytes | 
**Creation** | **time.Time** | When the analysis was created | 
**IsOwner** | **bool** | True when the caller owns the analysis | 
**ModelId** | **int64** | Model ID | 
**Sha256Hash** | **string** | SHA-256 hash of the binary | 
**Status** | **string** | Analysis status | 
**Tags** | **[]string** | Tags associated with this binary | 
**Username** | **string** | Username of the analysis owner | 

## Methods

### NewExample

`func NewExample(analysisId int64, analysisScope string, binaryId int64, binaryName string, binarySize int64, creation time.Time, isOwner bool, modelId int64, sha256Hash string, status string, tags []string, username string, ) *Example`

NewExample instantiates a new Example object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExampleWithDefaults

`func NewExampleWithDefaults() *Example`

NewExampleWithDefaults instantiates a new Example object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisId

`func (o *Example) GetAnalysisId() int64`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *Example) GetAnalysisIdOk() (*int64, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *Example) SetAnalysisId(v int64)`

SetAnalysisId sets AnalysisId field to given value.


### GetAnalysisScope

`func (o *Example) GetAnalysisScope() string`

GetAnalysisScope returns the AnalysisScope field if non-nil, zero value otherwise.

### GetAnalysisScopeOk

`func (o *Example) GetAnalysisScopeOk() (*string, bool)`

GetAnalysisScopeOk returns a tuple with the AnalysisScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisScope

`func (o *Example) SetAnalysisScope(v string)`

SetAnalysisScope sets AnalysisScope field to given value.


### GetBinaryId

`func (o *Example) GetBinaryId() int64`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *Example) GetBinaryIdOk() (*int64, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *Example) SetBinaryId(v int64)`

SetBinaryId sets BinaryId field to given value.


### GetBinaryName

`func (o *Example) GetBinaryName() string`

GetBinaryName returns the BinaryName field if non-nil, zero value otherwise.

### GetBinaryNameOk

`func (o *Example) GetBinaryNameOk() (*string, bool)`

GetBinaryNameOk returns a tuple with the BinaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryName

`func (o *Example) SetBinaryName(v string)`

SetBinaryName sets BinaryName field to given value.


### GetBinarySize

`func (o *Example) GetBinarySize() int64`

GetBinarySize returns the BinarySize field if non-nil, zero value otherwise.

### GetBinarySizeOk

`func (o *Example) GetBinarySizeOk() (*int64, bool)`

GetBinarySizeOk returns a tuple with the BinarySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinarySize

`func (o *Example) SetBinarySize(v int64)`

SetBinarySize sets BinarySize field to given value.


### GetCreation

`func (o *Example) GetCreation() time.Time`

GetCreation returns the Creation field if non-nil, zero value otherwise.

### GetCreationOk

`func (o *Example) GetCreationOk() (*time.Time, bool)`

GetCreationOk returns a tuple with the Creation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreation

`func (o *Example) SetCreation(v time.Time)`

SetCreation sets Creation field to given value.


### GetIsOwner

`func (o *Example) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *Example) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *Example) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.


### GetModelId

`func (o *Example) GetModelId() int64`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *Example) GetModelIdOk() (*int64, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *Example) SetModelId(v int64)`

SetModelId sets ModelId field to given value.


### GetSha256Hash

`func (o *Example) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *Example) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *Example) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.


### GetStatus

`func (o *Example) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Example) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Example) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTags

`func (o *Example) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Example) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Example) SetTags(v []string)`

SetTags sets Tags field to given value.


### SetTagsNil

`func (o *Example) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *Example) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetUsername

`func (o *Example) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Example) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Example) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


