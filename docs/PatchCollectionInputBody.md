# PatchCollectionInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionName** | Pointer to **string** | New collection name. Omit, null, or empty string to keep existing. | [optional] 
**CollectionScope** | Pointer to **string** | New scope (PUBLIC, PRIVATE, PROTECTED, TEAM). Omit or send null to keep existing. Empty string returns 422 UNPROCESSABLE ENTITY. | [optional] 
**Description** | Pointer to **string** | New description. Omit, null, or empty string to keep existing. | [optional] 

## Methods

### NewPatchCollectionInputBody

`func NewPatchCollectionInputBody() *PatchCollectionInputBody`

NewPatchCollectionInputBody instantiates a new PatchCollectionInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchCollectionInputBodyWithDefaults

`func NewPatchCollectionInputBodyWithDefaults() *PatchCollectionInputBody`

NewPatchCollectionInputBodyWithDefaults instantiates a new PatchCollectionInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionName

`func (o *PatchCollectionInputBody) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *PatchCollectionInputBody) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *PatchCollectionInputBody) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.

### HasCollectionName

`func (o *PatchCollectionInputBody) HasCollectionName() bool`

HasCollectionName returns a boolean if a field has been set.

### GetCollectionScope

`func (o *PatchCollectionInputBody) GetCollectionScope() string`

GetCollectionScope returns the CollectionScope field if non-nil, zero value otherwise.

### GetCollectionScopeOk

`func (o *PatchCollectionInputBody) GetCollectionScopeOk() (*string, bool)`

GetCollectionScopeOk returns a tuple with the CollectionScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionScope

`func (o *PatchCollectionInputBody) SetCollectionScope(v string)`

SetCollectionScope sets CollectionScope field to given value.

### HasCollectionScope

`func (o *PatchCollectionInputBody) HasCollectionScope() bool`

HasCollectionScope returns a boolean if a field has been set.

### GetDescription

`func (o *PatchCollectionInputBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchCollectionInputBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchCollectionInputBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchCollectionInputBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


