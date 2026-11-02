# EntityUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]int64** |  | [optional] 
**Refs** | Pointer to [**[]EntityRef**](EntityRef.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewEntityUpdate

`func NewEntityUpdate() *EntityUpdate`

NewEntityUpdate instantiates a new EntityUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityUpdateWithDefaults

`func NewEntityUpdateWithDefaults() *EntityUpdate`

NewEntityUpdateWithDefaults instantiates a new EntityUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *EntityUpdate) GetIds() []int64`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *EntityUpdate) GetIdsOk() (*[]int64, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *EntityUpdate) SetIds(v []int64)`

SetIds sets Ids field to given value.

### HasIds

`func (o *EntityUpdate) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *EntityUpdate) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *EntityUpdate) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetRefs

`func (o *EntityUpdate) GetRefs() []EntityRef`

GetRefs returns the Refs field if non-nil, zero value otherwise.

### GetRefsOk

`func (o *EntityUpdate) GetRefsOk() (*[]EntityRef, bool)`

GetRefsOk returns a tuple with the Refs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefs

`func (o *EntityUpdate) SetRefs(v []EntityRef)`

SetRefs sets Refs field to given value.

### HasRefs

`func (o *EntityUpdate) HasRefs() bool`

HasRefs returns a boolean if a field has been set.

### SetRefsNil

`func (o *EntityUpdate) SetRefsNil(b bool)`

 SetRefsNil sets the value for Refs to be an explicit nil

### UnsetRefs
`func (o *EntityUpdate) UnsetRefs()`

UnsetRefs ensures that no value is present for Refs, not even an explicit nil
### GetType

`func (o *EntityUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntityUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntityUpdate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EntityUpdate) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


