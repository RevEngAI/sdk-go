# CollectionSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionId** | **int32** | The ID of the collection | 
**CollectionName** | **string** | The name of the collection | 
**Scope** | **string** | The scope of the collection | 
**LastUpdatedAt** | **time.Time** | The last update date of the collection | 
**CreatedAt** | **time.Time** | The creation date of the collection | 
**ModelId** | **int32** | The model ID of the binary | 
**ModelName** | **string** | The name of the model | 
**OwnedBy** | **string** | The owner of the collection | 
**Tags** | Pointer to **[]string** |  | [optional] 
**Size** | Pointer to **NullableInt32** |  | [optional] 
**Description** | **string** | The description of the collection | 
**TeamId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCollectionSearchResult

`func NewCollectionSearchResult(collectionId int32, collectionName string, scope string, lastUpdatedAt time.Time, createdAt time.Time, modelId int32, modelName string, ownedBy string, description string, ) *CollectionSearchResult`

NewCollectionSearchResult instantiates a new CollectionSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionSearchResultWithDefaults

`func NewCollectionSearchResultWithDefaults() *CollectionSearchResult`

NewCollectionSearchResultWithDefaults instantiates a new CollectionSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionId

`func (o *CollectionSearchResult) GetCollectionId() int32`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *CollectionSearchResult) GetCollectionIdOk() (*int32, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *CollectionSearchResult) SetCollectionId(v int32)`

SetCollectionId sets CollectionId field to given value.


### GetCollectionName

`func (o *CollectionSearchResult) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *CollectionSearchResult) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *CollectionSearchResult) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.


### GetScope

`func (o *CollectionSearchResult) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CollectionSearchResult) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CollectionSearchResult) SetScope(v string)`

SetScope sets Scope field to given value.


### GetLastUpdatedAt

`func (o *CollectionSearchResult) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *CollectionSearchResult) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *CollectionSearchResult) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.


### GetCreatedAt

`func (o *CollectionSearchResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CollectionSearchResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CollectionSearchResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModelId

`func (o *CollectionSearchResult) GetModelId() int32`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *CollectionSearchResult) GetModelIdOk() (*int32, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *CollectionSearchResult) SetModelId(v int32)`

SetModelId sets ModelId field to given value.


### GetModelName

`func (o *CollectionSearchResult) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *CollectionSearchResult) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *CollectionSearchResult) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetOwnedBy

`func (o *CollectionSearchResult) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *CollectionSearchResult) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *CollectionSearchResult) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.


### GetTags

`func (o *CollectionSearchResult) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CollectionSearchResult) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CollectionSearchResult) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CollectionSearchResult) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CollectionSearchResult) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CollectionSearchResult) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetSize

`func (o *CollectionSearchResult) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CollectionSearchResult) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CollectionSearchResult) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *CollectionSearchResult) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *CollectionSearchResult) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *CollectionSearchResult) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetDescription

`func (o *CollectionSearchResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CollectionSearchResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CollectionSearchResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTeamId

`func (o *CollectionSearchResult) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *CollectionSearchResult) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *CollectionSearchResult) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *CollectionSearchResult) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### SetTeamIdNil

`func (o *CollectionSearchResult) SetTeamIdNil(b bool)`

 SetTeamIdNil sets the value for TeamId to be an explicit nil

### UnsetTeamId
`func (o *CollectionSearchResult) UnsetTeamId()`

UnsetTeamId ensures that no value is present for TeamId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


