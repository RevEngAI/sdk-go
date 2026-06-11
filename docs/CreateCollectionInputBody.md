# CreateCollectionInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Binaries** | Pointer to **[]int64** | Optional binary IDs to link to the collection. | [optional] 
**CollectionName** | **string** | Collection name. | 
**CollectionScope** | **string** | Visibility scope. | [default to "PRIVATE"]
**Description** | **string** | Collection description. | 
**ModelId** | **int64** | Model ID the collection is associated with. | 
**Tags** | Pointer to **[]string** | Optional tags to attach to the collection. | [optional] 

## Methods

### NewCreateCollectionInputBody

`func NewCreateCollectionInputBody(collectionName string, collectionScope string, description string, modelId int64, ) *CreateCollectionInputBody`

NewCreateCollectionInputBody instantiates a new CreateCollectionInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCollectionInputBodyWithDefaults

`func NewCreateCollectionInputBodyWithDefaults() *CreateCollectionInputBody`

NewCreateCollectionInputBodyWithDefaults instantiates a new CreateCollectionInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaries

`func (o *CreateCollectionInputBody) GetBinaries() []int64`

GetBinaries returns the Binaries field if non-nil, zero value otherwise.

### GetBinariesOk

`func (o *CreateCollectionInputBody) GetBinariesOk() (*[]int64, bool)`

GetBinariesOk returns a tuple with the Binaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaries

`func (o *CreateCollectionInputBody) SetBinaries(v []int64)`

SetBinaries sets Binaries field to given value.

### HasBinaries

`func (o *CreateCollectionInputBody) HasBinaries() bool`

HasBinaries returns a boolean if a field has been set.

### SetBinariesNil

`func (o *CreateCollectionInputBody) SetBinariesNil(b bool)`

 SetBinariesNil sets the value for Binaries to be an explicit nil

### UnsetBinaries
`func (o *CreateCollectionInputBody) UnsetBinaries()`

UnsetBinaries ensures that no value is present for Binaries, not even an explicit nil
### GetCollectionName

`func (o *CreateCollectionInputBody) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *CreateCollectionInputBody) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *CreateCollectionInputBody) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.


### GetCollectionScope

`func (o *CreateCollectionInputBody) GetCollectionScope() string`

GetCollectionScope returns the CollectionScope field if non-nil, zero value otherwise.

### GetCollectionScopeOk

`func (o *CreateCollectionInputBody) GetCollectionScopeOk() (*string, bool)`

GetCollectionScopeOk returns a tuple with the CollectionScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionScope

`func (o *CreateCollectionInputBody) SetCollectionScope(v string)`

SetCollectionScope sets CollectionScope field to given value.


### GetDescription

`func (o *CreateCollectionInputBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCollectionInputBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCollectionInputBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetModelId

`func (o *CreateCollectionInputBody) GetModelId() int64`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *CreateCollectionInputBody) GetModelIdOk() (*int64, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *CreateCollectionInputBody) SetModelId(v int64)`

SetModelId sets ModelId field to given value.


### GetTags

`func (o *CreateCollectionInputBody) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateCollectionInputBody) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateCollectionInputBody) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateCollectionInputBody) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CreateCollectionInputBody) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CreateCollectionInputBody) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


