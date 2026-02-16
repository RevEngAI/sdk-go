# CollectionListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionName** | **string** | The name of the collection | 
**Description** | **string** | The description of the collection | 
**CollectionScope** | **string** | The scope of the collection | 
**CollectionOwner** | **string** | The owner of the collection | 
**OfficialCollection** | **bool** | Whether the collection is maintained by RevEng.AI | 
**CollectionTags** | Pointer to **[]string** | The tags of the collection | [optional] [default to {}]
**CollectionSize** | **int32** | The size of the collection | 
**CollectionId** | **int32** | The ID of the collection | 
**Creation** | **time.Time** | The current status of analysis | 
**ModelName** | **string** | The model being used for the collection | 
**TeamId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCollectionListItem

`func NewCollectionListItem(collectionName string, description string, collectionScope string, collectionOwner string, officialCollection bool, collectionSize int32, collectionId int32, creation time.Time, modelName string, ) *CollectionListItem`

NewCollectionListItem instantiates a new CollectionListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionListItemWithDefaults

`func NewCollectionListItemWithDefaults() *CollectionListItem`

NewCollectionListItemWithDefaults instantiates a new CollectionListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionName

`func (o *CollectionListItem) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *CollectionListItem) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *CollectionListItem) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.


### GetDescription

`func (o *CollectionListItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CollectionListItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CollectionListItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCollectionScope

`func (o *CollectionListItem) GetCollectionScope() string`

GetCollectionScope returns the CollectionScope field if non-nil, zero value otherwise.

### GetCollectionScopeOk

`func (o *CollectionListItem) GetCollectionScopeOk() (*string, bool)`

GetCollectionScopeOk returns a tuple with the CollectionScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionScope

`func (o *CollectionListItem) SetCollectionScope(v string)`

SetCollectionScope sets CollectionScope field to given value.


### GetCollectionOwner

`func (o *CollectionListItem) GetCollectionOwner() string`

GetCollectionOwner returns the CollectionOwner field if non-nil, zero value otherwise.

### GetCollectionOwnerOk

`func (o *CollectionListItem) GetCollectionOwnerOk() (*string, bool)`

GetCollectionOwnerOk returns a tuple with the CollectionOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionOwner

`func (o *CollectionListItem) SetCollectionOwner(v string)`

SetCollectionOwner sets CollectionOwner field to given value.


### GetOfficialCollection

`func (o *CollectionListItem) GetOfficialCollection() bool`

GetOfficialCollection returns the OfficialCollection field if non-nil, zero value otherwise.

### GetOfficialCollectionOk

`func (o *CollectionListItem) GetOfficialCollectionOk() (*bool, bool)`

GetOfficialCollectionOk returns a tuple with the OfficialCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialCollection

`func (o *CollectionListItem) SetOfficialCollection(v bool)`

SetOfficialCollection sets OfficialCollection field to given value.


### GetCollectionTags

`func (o *CollectionListItem) GetCollectionTags() []string`

GetCollectionTags returns the CollectionTags field if non-nil, zero value otherwise.

### GetCollectionTagsOk

`func (o *CollectionListItem) GetCollectionTagsOk() (*[]string, bool)`

GetCollectionTagsOk returns a tuple with the CollectionTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionTags

`func (o *CollectionListItem) SetCollectionTags(v []string)`

SetCollectionTags sets CollectionTags field to given value.

### HasCollectionTags

`func (o *CollectionListItem) HasCollectionTags() bool`

HasCollectionTags returns a boolean if a field has been set.

### GetCollectionSize

`func (o *CollectionListItem) GetCollectionSize() int32`

GetCollectionSize returns the CollectionSize field if non-nil, zero value otherwise.

### GetCollectionSizeOk

`func (o *CollectionListItem) GetCollectionSizeOk() (*int32, bool)`

GetCollectionSizeOk returns a tuple with the CollectionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionSize

`func (o *CollectionListItem) SetCollectionSize(v int32)`

SetCollectionSize sets CollectionSize field to given value.


### GetCollectionId

`func (o *CollectionListItem) GetCollectionId() int32`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *CollectionListItem) GetCollectionIdOk() (*int32, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *CollectionListItem) SetCollectionId(v int32)`

SetCollectionId sets CollectionId field to given value.


### GetCreation

`func (o *CollectionListItem) GetCreation() time.Time`

GetCreation returns the Creation field if non-nil, zero value otherwise.

### GetCreationOk

`func (o *CollectionListItem) GetCreationOk() (*time.Time, bool)`

GetCreationOk returns a tuple with the Creation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreation

`func (o *CollectionListItem) SetCreation(v time.Time)`

SetCreation sets Creation field to given value.


### GetModelName

`func (o *CollectionListItem) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *CollectionListItem) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *CollectionListItem) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetTeamId

`func (o *CollectionListItem) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *CollectionListItem) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *CollectionListItem) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *CollectionListItem) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### SetTeamIdNil

`func (o *CollectionListItem) SetTeamIdNil(b bool)`

 SetTeamIdNil sets the value for TeamId to be an explicit nil

### UnsetTeamId
`func (o *CollectionListItem) UnsetTeamId()`

UnsetTeamId ensures that no value is present for TeamId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


