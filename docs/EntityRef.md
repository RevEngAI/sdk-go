# EntityRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Vaddr** | Pointer to **int64** |  | [optional] 

## Methods

### NewEntityRef

`func NewEntityRef(id int64, ) *EntityRef`

NewEntityRef instantiates a new EntityRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityRefWithDefaults

`func NewEntityRefWithDefaults() *EntityRef`

NewEntityRefWithDefaults instantiates a new EntityRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntityRef) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityRef) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityRef) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *EntityRef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityRef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityRef) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntityRef) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVaddr

`func (o *EntityRef) GetVaddr() int64`

GetVaddr returns the Vaddr field if non-nil, zero value otherwise.

### GetVaddrOk

`func (o *EntityRef) GetVaddrOk() (*int64, bool)`

GetVaddrOk returns a tuple with the Vaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaddr

`func (o *EntityRef) SetVaddr(v int64)`

SetVaddr sets Vaddr field to given value.

### HasVaddr

`func (o *EntityRef) HasVaddr() bool`

HasVaddr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


