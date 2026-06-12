# GetCollectionOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Binaries** | Pointer to [**[]Binary**](Binary.md) |  | [optional] 
**CollectionId** | **int64** |  | 
**CollectionName** | **string** |  | 
**CollectionScope** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Description** | **string** |  | 
**HasNextPage** | Pointer to **bool** |  | [optional] 
**ModelId** | **int64** |  | 
**PageNumber** | Pointer to **int64** |  | [optional] 
**PageSize** | Pointer to **int64** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TeamId** | **int64** |  | 
**UpdatedAt** | **time.Time** |  | 
**UserId** | **int64** |  | 

## Methods

### NewGetCollectionOutputBody

`func NewGetCollectionOutputBody(collectionId int64, collectionName string, collectionScope string, createdAt time.Time, description string, modelId int64, teamId int64, updatedAt time.Time, userId int64, ) *GetCollectionOutputBody`

NewGetCollectionOutputBody instantiates a new GetCollectionOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCollectionOutputBodyWithDefaults

`func NewGetCollectionOutputBodyWithDefaults() *GetCollectionOutputBody`

NewGetCollectionOutputBodyWithDefaults instantiates a new GetCollectionOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaries

`func (o *GetCollectionOutputBody) GetBinaries() []Binary`

GetBinaries returns the Binaries field if non-nil, zero value otherwise.

### GetBinariesOk

`func (o *GetCollectionOutputBody) GetBinariesOk() (*[]Binary, bool)`

GetBinariesOk returns a tuple with the Binaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaries

`func (o *GetCollectionOutputBody) SetBinaries(v []Binary)`

SetBinaries sets Binaries field to given value.

### HasBinaries

`func (o *GetCollectionOutputBody) HasBinaries() bool`

HasBinaries returns a boolean if a field has been set.

### SetBinariesNil

`func (o *GetCollectionOutputBody) SetBinariesNil(b bool)`

 SetBinariesNil sets the value for Binaries to be an explicit nil

### UnsetBinaries
`func (o *GetCollectionOutputBody) UnsetBinaries()`

UnsetBinaries ensures that no value is present for Binaries, not even an explicit nil
### GetCollectionId

`func (o *GetCollectionOutputBody) GetCollectionId() int64`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *GetCollectionOutputBody) GetCollectionIdOk() (*int64, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *GetCollectionOutputBody) SetCollectionId(v int64)`

SetCollectionId sets CollectionId field to given value.


### GetCollectionName

`func (o *GetCollectionOutputBody) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *GetCollectionOutputBody) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *GetCollectionOutputBody) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.


### GetCollectionScope

`func (o *GetCollectionOutputBody) GetCollectionScope() string`

GetCollectionScope returns the CollectionScope field if non-nil, zero value otherwise.

### GetCollectionScopeOk

`func (o *GetCollectionOutputBody) GetCollectionScopeOk() (*string, bool)`

GetCollectionScopeOk returns a tuple with the CollectionScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionScope

`func (o *GetCollectionOutputBody) SetCollectionScope(v string)`

SetCollectionScope sets CollectionScope field to given value.


### GetCreatedAt

`func (o *GetCollectionOutputBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetCollectionOutputBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetCollectionOutputBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *GetCollectionOutputBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetCollectionOutputBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetCollectionOutputBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHasNextPage

`func (o *GetCollectionOutputBody) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *GetCollectionOutputBody) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *GetCollectionOutputBody) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.

### HasHasNextPage

`func (o *GetCollectionOutputBody) HasHasNextPage() bool`

HasHasNextPage returns a boolean if a field has been set.

### GetModelId

`func (o *GetCollectionOutputBody) GetModelId() int64`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *GetCollectionOutputBody) GetModelIdOk() (*int64, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *GetCollectionOutputBody) SetModelId(v int64)`

SetModelId sets ModelId field to given value.


### GetPageNumber

`func (o *GetCollectionOutputBody) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *GetCollectionOutputBody) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *GetCollectionOutputBody) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *GetCollectionOutputBody) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *GetCollectionOutputBody) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetCollectionOutputBody) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetCollectionOutputBody) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *GetCollectionOutputBody) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTags

`func (o *GetCollectionOutputBody) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetCollectionOutputBody) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetCollectionOutputBody) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetCollectionOutputBody) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *GetCollectionOutputBody) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *GetCollectionOutputBody) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTeamId

`func (o *GetCollectionOutputBody) GetTeamId() int64`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *GetCollectionOutputBody) GetTeamIdOk() (*int64, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *GetCollectionOutputBody) SetTeamId(v int64)`

SetTeamId sets TeamId field to given value.


### GetUpdatedAt

`func (o *GetCollectionOutputBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetCollectionOutputBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetCollectionOutputBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *GetCollectionOutputBody) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetCollectionOutputBody) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetCollectionOutputBody) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


