# BinarySearchResultBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisId** | **int64** |  | 
**BinaryId** | **int64** |  | 
**BinaryName** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**ModelId** | **int64** |  | 
**ModelName** | **string** |  | 
**OwnedBy** | **string** |  | 
**Sha256Hash** | **string** |  | 
**Tags** | **[]string** |  | 

## Methods

### NewBinarySearchResultBody

`func NewBinarySearchResultBody(analysisId int64, binaryId int64, binaryName string, createdAt time.Time, modelId int64, modelName string, ownedBy string, sha256Hash string, tags []string, ) *BinarySearchResultBody`

NewBinarySearchResultBody instantiates a new BinarySearchResultBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBinarySearchResultBodyWithDefaults

`func NewBinarySearchResultBodyWithDefaults() *BinarySearchResultBody`

NewBinarySearchResultBodyWithDefaults instantiates a new BinarySearchResultBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisId

`func (o *BinarySearchResultBody) GetAnalysisId() int64`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *BinarySearchResultBody) GetAnalysisIdOk() (*int64, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *BinarySearchResultBody) SetAnalysisId(v int64)`

SetAnalysisId sets AnalysisId field to given value.


### GetBinaryId

`func (o *BinarySearchResultBody) GetBinaryId() int64`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *BinarySearchResultBody) GetBinaryIdOk() (*int64, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *BinarySearchResultBody) SetBinaryId(v int64)`

SetBinaryId sets BinaryId field to given value.


### GetBinaryName

`func (o *BinarySearchResultBody) GetBinaryName() string`

GetBinaryName returns the BinaryName field if non-nil, zero value otherwise.

### GetBinaryNameOk

`func (o *BinarySearchResultBody) GetBinaryNameOk() (*string, bool)`

GetBinaryNameOk returns a tuple with the BinaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryName

`func (o *BinarySearchResultBody) SetBinaryName(v string)`

SetBinaryName sets BinaryName field to given value.


### GetCreatedAt

`func (o *BinarySearchResultBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BinarySearchResultBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BinarySearchResultBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModelId

`func (o *BinarySearchResultBody) GetModelId() int64`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *BinarySearchResultBody) GetModelIdOk() (*int64, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *BinarySearchResultBody) SetModelId(v int64)`

SetModelId sets ModelId field to given value.


### GetModelName

`func (o *BinarySearchResultBody) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *BinarySearchResultBody) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *BinarySearchResultBody) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetOwnedBy

`func (o *BinarySearchResultBody) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *BinarySearchResultBody) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *BinarySearchResultBody) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.


### GetSha256Hash

`func (o *BinarySearchResultBody) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *BinarySearchResultBody) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *BinarySearchResultBody) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.


### GetTags

`func (o *BinarySearchResultBody) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BinarySearchResultBody) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BinarySearchResultBody) SetTags(v []string)`

SetTags sets Tags field to given value.


### SetTagsNil

`func (o *BinarySearchResultBody) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *BinarySearchResultBody) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


