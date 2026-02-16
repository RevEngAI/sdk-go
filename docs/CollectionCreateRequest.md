# CollectionCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionName** | **string** |  | 
**Description** | **string** |  | 
**CollectionScope** | Pointer to [**CollectionScope**](CollectionScope.md) |  | [optional] [default to COLLECTIONSCOPE_PRIVATE]
**Tags** | Pointer to **[]string** |  | [optional] 
**Binaries** | Pointer to **[]int32** |  | [optional] 
**ModelId** | **int32** |  | 

## Methods

### NewCollectionCreateRequest

`func NewCollectionCreateRequest(collectionName string, description string, modelId int32, ) *CollectionCreateRequest`

NewCollectionCreateRequest instantiates a new CollectionCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionCreateRequestWithDefaults

`func NewCollectionCreateRequestWithDefaults() *CollectionCreateRequest`

NewCollectionCreateRequestWithDefaults instantiates a new CollectionCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionName

`func (o *CollectionCreateRequest) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *CollectionCreateRequest) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *CollectionCreateRequest) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.


### GetDescription

`func (o *CollectionCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CollectionCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CollectionCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCollectionScope

`func (o *CollectionCreateRequest) GetCollectionScope() CollectionScope`

GetCollectionScope returns the CollectionScope field if non-nil, zero value otherwise.

### GetCollectionScopeOk

`func (o *CollectionCreateRequest) GetCollectionScopeOk() (*CollectionScope, bool)`

GetCollectionScopeOk returns a tuple with the CollectionScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionScope

`func (o *CollectionCreateRequest) SetCollectionScope(v CollectionScope)`

SetCollectionScope sets CollectionScope field to given value.

### HasCollectionScope

`func (o *CollectionCreateRequest) HasCollectionScope() bool`

HasCollectionScope returns a boolean if a field has been set.

### GetTags

`func (o *CollectionCreateRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CollectionCreateRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CollectionCreateRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CollectionCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CollectionCreateRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CollectionCreateRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetBinaries

`func (o *CollectionCreateRequest) GetBinaries() []int32`

GetBinaries returns the Binaries field if non-nil, zero value otherwise.

### GetBinariesOk

`func (o *CollectionCreateRequest) GetBinariesOk() (*[]int32, bool)`

GetBinariesOk returns a tuple with the Binaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaries

`func (o *CollectionCreateRequest) SetBinaries(v []int32)`

SetBinaries sets Binaries field to given value.

### HasBinaries

`func (o *CollectionCreateRequest) HasBinaries() bool`

HasBinaries returns a boolean if a field has been set.

### SetBinariesNil

`func (o *CollectionCreateRequest) SetBinariesNil(b bool)`

 SetBinariesNil sets the value for Binaries to be an explicit nil

### UnsetBinaries
`func (o *CollectionCreateRequest) UnsetBinaries()`

UnsetBinaries ensures that no value is present for Binaries, not even an explicit nil
### GetModelId

`func (o *CollectionCreateRequest) GetModelId() int32`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *CollectionCreateRequest) GetModelIdOk() (*int32, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *CollectionCreateRequest) SetModelId(v int32)`

SetModelId sets ModelId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


