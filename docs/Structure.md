# Structure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastChange** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** | Name of the structure | 
**Size** | Pointer to **NullableInt32** |  | [optional] 
**Members** | [**map[string]StructureMember**](StructureMember.md) | Dictionary of structure members | 
**ArtifactType** | Pointer to **string** | Type of artifact that the structure is associated with | [optional] 

## Methods

### NewStructure

`func NewStructure(name string, members map[string]StructureMember, ) *Structure`

NewStructure instantiates a new Structure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructureWithDefaults

`func NewStructureWithDefaults() *Structure`

NewStructureWithDefaults instantiates a new Structure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastChange

`func (o *Structure) GetLastChange() string`

GetLastChange returns the LastChange field if non-nil, zero value otherwise.

### GetLastChangeOk

`func (o *Structure) GetLastChangeOk() (*string, bool)`

GetLastChangeOk returns a tuple with the LastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChange

`func (o *Structure) SetLastChange(v string)`

SetLastChange sets LastChange field to given value.

### HasLastChange

`func (o *Structure) HasLastChange() bool`

HasLastChange returns a boolean if a field has been set.

### SetLastChangeNil

`func (o *Structure) SetLastChangeNil(b bool)`

 SetLastChangeNil sets the value for LastChange to be an explicit nil

### UnsetLastChange
`func (o *Structure) UnsetLastChange()`

UnsetLastChange ensures that no value is present for LastChange, not even an explicit nil
### GetName

`func (o *Structure) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Structure) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Structure) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *Structure) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Structure) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Structure) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Structure) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *Structure) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *Structure) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetMembers

`func (o *Structure) GetMembers() map[string]StructureMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Structure) GetMembersOk() (*map[string]StructureMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Structure) SetMembers(v map[string]StructureMember)`

SetMembers sets Members field to given value.


### GetArtifactType

`func (o *Structure) GetArtifactType() string`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *Structure) GetArtifactTypeOk() (*string, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *Structure) SetArtifactType(v string)`

SetArtifactType sets ArtifactType field to given value.

### HasArtifactType

`func (o *Structure) HasArtifactType() bool`

HasArtifactType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


