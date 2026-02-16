# CollectionUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionName** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**CollectionScope** | Pointer to [**NullableCollectionScope**](CollectionScope.md) |  | [optional] 

## Methods

### NewCollectionUpdateRequest

`func NewCollectionUpdateRequest() *CollectionUpdateRequest`

NewCollectionUpdateRequest instantiates a new CollectionUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionUpdateRequestWithDefaults

`func NewCollectionUpdateRequestWithDefaults() *CollectionUpdateRequest`

NewCollectionUpdateRequestWithDefaults instantiates a new CollectionUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionName

`func (o *CollectionUpdateRequest) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *CollectionUpdateRequest) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *CollectionUpdateRequest) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.

### HasCollectionName

`func (o *CollectionUpdateRequest) HasCollectionName() bool`

HasCollectionName returns a boolean if a field has been set.

### SetCollectionNameNil

`func (o *CollectionUpdateRequest) SetCollectionNameNil(b bool)`

 SetCollectionNameNil sets the value for CollectionName to be an explicit nil

### UnsetCollectionName
`func (o *CollectionUpdateRequest) UnsetCollectionName()`

UnsetCollectionName ensures that no value is present for CollectionName, not even an explicit nil
### GetDescription

`func (o *CollectionUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CollectionUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CollectionUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CollectionUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CollectionUpdateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CollectionUpdateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCollectionScope

`func (o *CollectionUpdateRequest) GetCollectionScope() CollectionScope`

GetCollectionScope returns the CollectionScope field if non-nil, zero value otherwise.

### GetCollectionScopeOk

`func (o *CollectionUpdateRequest) GetCollectionScopeOk() (*CollectionScope, bool)`

GetCollectionScopeOk returns a tuple with the CollectionScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionScope

`func (o *CollectionUpdateRequest) SetCollectionScope(v CollectionScope)`

SetCollectionScope sets CollectionScope field to given value.

### HasCollectionScope

`func (o *CollectionUpdateRequest) HasCollectionScope() bool`

HasCollectionScope returns a boolean if a field has been set.

### SetCollectionScopeNil

`func (o *CollectionUpdateRequest) SetCollectionScopeNil(b bool)`

 SetCollectionScopeNil sets the value for CollectionScope to be an explicit nil

### UnsetCollectionScope
`func (o *CollectionUpdateRequest) UnsetCollectionScope()`

UnsetCollectionScope ensures that no value is present for CollectionScope, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


