# CollectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionId** | **int32** | Collection ID | 
**CollectionName** | **string** | Collection name | 
**Description** | **string** | Collection description | 
**ModelId** | **int32** | Collection model ID | 
**UserId** | **int32** | Collection user ID | 
**TeamId** | Pointer to **NullableInt32** |  | [optional] 
**CollectionScope** | [**CollectionScope**](CollectionScope.md) | Collection public status | 
**CreatedAt** | **time.Time** | Collection creation date | 
**UpdatedAt** | **time.Time** | Collection last update date | 
**Tags** | Pointer to **[]string** |  | [optional] 
**Binaries** | Pointer to [**[]CollectionResponseBinariesInner**](CollectionResponseBinariesInner.md) |  | [optional] 

## Methods

### NewCollectionResponse

`func NewCollectionResponse(collectionId int32, collectionName string, description string, modelId int32, userId int32, collectionScope CollectionScope, createdAt time.Time, updatedAt time.Time, ) *CollectionResponse`

NewCollectionResponse instantiates a new CollectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseWithDefaults

`func NewCollectionResponseWithDefaults() *CollectionResponse`

NewCollectionResponseWithDefaults instantiates a new CollectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionId

`func (o *CollectionResponse) GetCollectionId() int32`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *CollectionResponse) GetCollectionIdOk() (*int32, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *CollectionResponse) SetCollectionId(v int32)`

SetCollectionId sets CollectionId field to given value.


### GetCollectionName

`func (o *CollectionResponse) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *CollectionResponse) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *CollectionResponse) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.


### GetDescription

`func (o *CollectionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CollectionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CollectionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetModelId

`func (o *CollectionResponse) GetModelId() int32`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *CollectionResponse) GetModelIdOk() (*int32, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *CollectionResponse) SetModelId(v int32)`

SetModelId sets ModelId field to given value.


### GetUserId

`func (o *CollectionResponse) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CollectionResponse) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CollectionResponse) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetTeamId

`func (o *CollectionResponse) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *CollectionResponse) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *CollectionResponse) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *CollectionResponse) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### SetTeamIdNil

`func (o *CollectionResponse) SetTeamIdNil(b bool)`

 SetTeamIdNil sets the value for TeamId to be an explicit nil

### UnsetTeamId
`func (o *CollectionResponse) UnsetTeamId()`

UnsetTeamId ensures that no value is present for TeamId, not even an explicit nil
### GetCollectionScope

`func (o *CollectionResponse) GetCollectionScope() CollectionScope`

GetCollectionScope returns the CollectionScope field if non-nil, zero value otherwise.

### GetCollectionScopeOk

`func (o *CollectionResponse) GetCollectionScopeOk() (*CollectionScope, bool)`

GetCollectionScopeOk returns a tuple with the CollectionScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionScope

`func (o *CollectionResponse) SetCollectionScope(v CollectionScope)`

SetCollectionScope sets CollectionScope field to given value.


### GetCreatedAt

`func (o *CollectionResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CollectionResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CollectionResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CollectionResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CollectionResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CollectionResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetTags

`func (o *CollectionResponse) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CollectionResponse) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CollectionResponse) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CollectionResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CollectionResponse) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CollectionResponse) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetBinaries

`func (o *CollectionResponse) GetBinaries() []CollectionResponseBinariesInner`

GetBinaries returns the Binaries field if non-nil, zero value otherwise.

### GetBinariesOk

`func (o *CollectionResponse) GetBinariesOk() (*[]CollectionResponseBinariesInner, bool)`

GetBinariesOk returns a tuple with the Binaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaries

`func (o *CollectionResponse) SetBinaries(v []CollectionResponseBinariesInner)`

SetBinaries sets Binaries field to given value.

### HasBinaries

`func (o *CollectionResponse) HasBinaries() bool`

HasBinaries returns a boolean if a field has been set.

### SetBinariesNil

`func (o *CollectionResponse) SetBinariesNil(b bool)`

 SetBinariesNil sets the value for Binaries to be an explicit nil

### UnsetBinaries
`func (o *CollectionResponse) UnsetBinaries()`

UnsetBinaries ensures that no value is present for Binaries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


