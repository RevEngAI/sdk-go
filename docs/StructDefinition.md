# StructDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | [**[]DataTypeMemberEntry**](DataTypeMemberEntry.md) | The type&#39;s fields, in offset order. | 

## Methods

### NewStructDefinition

`func NewStructDefinition(members []DataTypeMemberEntry, ) *StructDefinition`

NewStructDefinition instantiates a new StructDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructDefinitionWithDefaults

`func NewStructDefinitionWithDefaults() *StructDefinition`

NewStructDefinitionWithDefaults instantiates a new StructDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *StructDefinition) GetMembers() []DataTypeMemberEntry`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *StructDefinition) GetMembersOk() (*[]DataTypeMemberEntry, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *StructDefinition) SetMembers(v []DataTypeMemberEntry)`

SetMembers sets Members field to given value.


### SetMembersNil

`func (o *StructDefinition) SetMembersNil(b bool)`

 SetMembersNil sets the value for Members to be an explicit nil

### UnsetMembers
`func (o *StructDefinition) UnsetMembers()`

UnsetMembers ensures that no value is present for Members, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


