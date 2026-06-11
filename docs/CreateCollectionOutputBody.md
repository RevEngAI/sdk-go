# CreateCollectionOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Binaries** | Pointer to [**[]Binary**](Binary.md) |  | [optional] 
**CollectionId** | **int64** |  | 
**CollectionName** | **string** |  | 
**CollectionScope** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Description** | **string** |  | 
**ModelId** | **int64** |  | 
**Tags** | Pointer to **[]string** |  | [optional] 
**TeamId** | **int64** |  | 
**UpdatedAt** | **time.Time** |  | 
**UserId** | **int64** |  | 

## Methods

### NewCreateCollectionOutputBody

`func NewCreateCollectionOutputBody(collectionId int64, collectionName string, collectionScope string, createdAt time.Time, description string, modelId int64, teamId int64, updatedAt time.Time, userId int64, ) *CreateCollectionOutputBody`

NewCreateCollectionOutputBody instantiates a new CreateCollectionOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCollectionOutputBodyWithDefaults

`func NewCreateCollectionOutputBodyWithDefaults() *CreateCollectionOutputBody`

NewCreateCollectionOutputBodyWithDefaults instantiates a new CreateCollectionOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaries

`func (o *CreateCollectionOutputBody) GetBinaries() []Binary`

GetBinaries returns the Binaries field if non-nil, zero value otherwise.

### GetBinariesOk

`func (o *CreateCollectionOutputBody) GetBinariesOk() (*[]Binary, bool)`

GetBinariesOk returns a tuple with the Binaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaries

`func (o *CreateCollectionOutputBody) SetBinaries(v []Binary)`

SetBinaries sets Binaries field to given value.

### HasBinaries

`func (o *CreateCollectionOutputBody) HasBinaries() bool`

HasBinaries returns a boolean if a field has been set.

### SetBinariesNil

`func (o *CreateCollectionOutputBody) SetBinariesNil(b bool)`

 SetBinariesNil sets the value for Binaries to be an explicit nil

### UnsetBinaries
`func (o *CreateCollectionOutputBody) UnsetBinaries()`

UnsetBinaries ensures that no value is present for Binaries, not even an explicit nil
### GetCollectionId

`func (o *CreateCollectionOutputBody) GetCollectionId() int64`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *CreateCollectionOutputBody) GetCollectionIdOk() (*int64, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *CreateCollectionOutputBody) SetCollectionId(v int64)`

SetCollectionId sets CollectionId field to given value.


### GetCollectionName

`func (o *CreateCollectionOutputBody) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *CreateCollectionOutputBody) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *CreateCollectionOutputBody) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.


### GetCollectionScope

`func (o *CreateCollectionOutputBody) GetCollectionScope() string`

GetCollectionScope returns the CollectionScope field if non-nil, zero value otherwise.

### GetCollectionScopeOk

`func (o *CreateCollectionOutputBody) GetCollectionScopeOk() (*string, bool)`

GetCollectionScopeOk returns a tuple with the CollectionScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionScope

`func (o *CreateCollectionOutputBody) SetCollectionScope(v string)`

SetCollectionScope sets CollectionScope field to given value.


### GetCreatedAt

`func (o *CreateCollectionOutputBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateCollectionOutputBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateCollectionOutputBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *CreateCollectionOutputBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCollectionOutputBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCollectionOutputBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetModelId

`func (o *CreateCollectionOutputBody) GetModelId() int64`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *CreateCollectionOutputBody) GetModelIdOk() (*int64, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *CreateCollectionOutputBody) SetModelId(v int64)`

SetModelId sets ModelId field to given value.


### GetTags

`func (o *CreateCollectionOutputBody) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateCollectionOutputBody) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateCollectionOutputBody) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateCollectionOutputBody) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CreateCollectionOutputBody) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CreateCollectionOutputBody) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTeamId

`func (o *CreateCollectionOutputBody) GetTeamId() int64`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *CreateCollectionOutputBody) GetTeamIdOk() (*int64, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *CreateCollectionOutputBody) SetTeamId(v int64)`

SetTeamId sets TeamId field to given value.


### GetUpdatedAt

`func (o *CreateCollectionOutputBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateCollectionOutputBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateCollectionOutputBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *CreateCollectionOutputBody) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateCollectionOutputBody) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateCollectionOutputBody) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


