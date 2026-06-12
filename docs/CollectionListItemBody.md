# CollectionListItemBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionId** | **int64** |  | 
**CollectionName** | **string** |  | 
**CollectionOwner** | **string** |  | 
**CollectionScope** | **string** |  | 
**CollectionSize** | **int64** |  | 
**CollectionTags** | **[]string** |  | 
**Creation** | **time.Time** |  | 
**Description** | **string** |  | 
**ModelName** | **string** |  | 
**OfficialCollection** | **bool** |  | 
**TeamId** | **int64** |  | 

## Methods

### NewCollectionListItemBody

`func NewCollectionListItemBody(collectionId int64, collectionName string, collectionOwner string, collectionScope string, collectionSize int64, collectionTags []string, creation time.Time, description string, modelName string, officialCollection bool, teamId int64, ) *CollectionListItemBody`

NewCollectionListItemBody instantiates a new CollectionListItemBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionListItemBodyWithDefaults

`func NewCollectionListItemBodyWithDefaults() *CollectionListItemBody`

NewCollectionListItemBodyWithDefaults instantiates a new CollectionListItemBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionId

`func (o *CollectionListItemBody) GetCollectionId() int64`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *CollectionListItemBody) GetCollectionIdOk() (*int64, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *CollectionListItemBody) SetCollectionId(v int64)`

SetCollectionId sets CollectionId field to given value.


### GetCollectionName

`func (o *CollectionListItemBody) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *CollectionListItemBody) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *CollectionListItemBody) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.


### GetCollectionOwner

`func (o *CollectionListItemBody) GetCollectionOwner() string`

GetCollectionOwner returns the CollectionOwner field if non-nil, zero value otherwise.

### GetCollectionOwnerOk

`func (o *CollectionListItemBody) GetCollectionOwnerOk() (*string, bool)`

GetCollectionOwnerOk returns a tuple with the CollectionOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionOwner

`func (o *CollectionListItemBody) SetCollectionOwner(v string)`

SetCollectionOwner sets CollectionOwner field to given value.


### GetCollectionScope

`func (o *CollectionListItemBody) GetCollectionScope() string`

GetCollectionScope returns the CollectionScope field if non-nil, zero value otherwise.

### GetCollectionScopeOk

`func (o *CollectionListItemBody) GetCollectionScopeOk() (*string, bool)`

GetCollectionScopeOk returns a tuple with the CollectionScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionScope

`func (o *CollectionListItemBody) SetCollectionScope(v string)`

SetCollectionScope sets CollectionScope field to given value.


### GetCollectionSize

`func (o *CollectionListItemBody) GetCollectionSize() int64`

GetCollectionSize returns the CollectionSize field if non-nil, zero value otherwise.

### GetCollectionSizeOk

`func (o *CollectionListItemBody) GetCollectionSizeOk() (*int64, bool)`

GetCollectionSizeOk returns a tuple with the CollectionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionSize

`func (o *CollectionListItemBody) SetCollectionSize(v int64)`

SetCollectionSize sets CollectionSize field to given value.


### GetCollectionTags

`func (o *CollectionListItemBody) GetCollectionTags() []string`

GetCollectionTags returns the CollectionTags field if non-nil, zero value otherwise.

### GetCollectionTagsOk

`func (o *CollectionListItemBody) GetCollectionTagsOk() (*[]string, bool)`

GetCollectionTagsOk returns a tuple with the CollectionTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionTags

`func (o *CollectionListItemBody) SetCollectionTags(v []string)`

SetCollectionTags sets CollectionTags field to given value.


### SetCollectionTagsNil

`func (o *CollectionListItemBody) SetCollectionTagsNil(b bool)`

 SetCollectionTagsNil sets the value for CollectionTags to be an explicit nil

### UnsetCollectionTags
`func (o *CollectionListItemBody) UnsetCollectionTags()`

UnsetCollectionTags ensures that no value is present for CollectionTags, not even an explicit nil
### GetCreation

`func (o *CollectionListItemBody) GetCreation() time.Time`

GetCreation returns the Creation field if non-nil, zero value otherwise.

### GetCreationOk

`func (o *CollectionListItemBody) GetCreationOk() (*time.Time, bool)`

GetCreationOk returns a tuple with the Creation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreation

`func (o *CollectionListItemBody) SetCreation(v time.Time)`

SetCreation sets Creation field to given value.


### GetDescription

`func (o *CollectionListItemBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CollectionListItemBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CollectionListItemBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetModelName

`func (o *CollectionListItemBody) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *CollectionListItemBody) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *CollectionListItemBody) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetOfficialCollection

`func (o *CollectionListItemBody) GetOfficialCollection() bool`

GetOfficialCollection returns the OfficialCollection field if non-nil, zero value otherwise.

### GetOfficialCollectionOk

`func (o *CollectionListItemBody) GetOfficialCollectionOk() (*bool, bool)`

GetOfficialCollectionOk returns a tuple with the OfficialCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialCollection

`func (o *CollectionListItemBody) SetOfficialCollection(v bool)`

SetOfficialCollection sets OfficialCollection field to given value.


### GetTeamId

`func (o *CollectionListItemBody) GetTeamId() int64`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *CollectionListItemBody) GetTeamIdOk() (*int64, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *CollectionListItemBody) SetTeamId(v int64)`

SetTeamId sets TeamId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


