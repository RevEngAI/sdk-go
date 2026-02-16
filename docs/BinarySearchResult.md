# BinarySearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinaryId** | **int32** | The binary ID | 
**BinaryName** | **string** | The name of the binary | 
**AnalysisId** | **int32** | The analysis ID | 
**Sha256Hash** | **string** | The SHA-256 hash of the binary | 
**Tags** | **[]string** |  | 
**CreatedAt** | **time.Time** | The creation date of the binary | 
**ModelId** | **int32** | The model ID of the binary | 
**ModelName** | **string** | The name of the model | 
**OwnedBy** | **string** | The owner of the binary | 

## Methods

### NewBinarySearchResult

`func NewBinarySearchResult(binaryId int32, binaryName string, analysisId int32, sha256Hash string, tags []string, createdAt time.Time, modelId int32, modelName string, ownedBy string, ) *BinarySearchResult`

NewBinarySearchResult instantiates a new BinarySearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBinarySearchResultWithDefaults

`func NewBinarySearchResultWithDefaults() *BinarySearchResult`

NewBinarySearchResultWithDefaults instantiates a new BinarySearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaryId

`func (o *BinarySearchResult) GetBinaryId() int32`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *BinarySearchResult) GetBinaryIdOk() (*int32, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *BinarySearchResult) SetBinaryId(v int32)`

SetBinaryId sets BinaryId field to given value.


### GetBinaryName

`func (o *BinarySearchResult) GetBinaryName() string`

GetBinaryName returns the BinaryName field if non-nil, zero value otherwise.

### GetBinaryNameOk

`func (o *BinarySearchResult) GetBinaryNameOk() (*string, bool)`

GetBinaryNameOk returns a tuple with the BinaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryName

`func (o *BinarySearchResult) SetBinaryName(v string)`

SetBinaryName sets BinaryName field to given value.


### GetAnalysisId

`func (o *BinarySearchResult) GetAnalysisId() int32`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *BinarySearchResult) GetAnalysisIdOk() (*int32, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *BinarySearchResult) SetAnalysisId(v int32)`

SetAnalysisId sets AnalysisId field to given value.


### GetSha256Hash

`func (o *BinarySearchResult) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *BinarySearchResult) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *BinarySearchResult) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.


### GetTags

`func (o *BinarySearchResult) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BinarySearchResult) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BinarySearchResult) SetTags(v []string)`

SetTags sets Tags field to given value.


### SetTagsNil

`func (o *BinarySearchResult) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *BinarySearchResult) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetCreatedAt

`func (o *BinarySearchResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BinarySearchResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BinarySearchResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModelId

`func (o *BinarySearchResult) GetModelId() int32`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *BinarySearchResult) GetModelIdOk() (*int32, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *BinarySearchResult) SetModelId(v int32)`

SetModelId sets ModelId field to given value.


### GetModelName

`func (o *BinarySearchResult) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *BinarySearchResult) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *BinarySearchResult) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetOwnedBy

`func (o *BinarySearchResult) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *BinarySearchResult) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *BinarySearchResult) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


