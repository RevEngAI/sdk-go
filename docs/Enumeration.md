# Enumeration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastChange** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** | Name of the enumeration | 
**Members** | **map[string]int32** | Dictionary of enumeration members and their values | 
**ArtifactType** | Pointer to **string** | Type of artifact that the enumeration is associated with | [optional] 

## Methods

### NewEnumeration

`func NewEnumeration(name string, members map[string]int32, ) *Enumeration`

NewEnumeration instantiates a new Enumeration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnumerationWithDefaults

`func NewEnumerationWithDefaults() *Enumeration`

NewEnumerationWithDefaults instantiates a new Enumeration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastChange

`func (o *Enumeration) GetLastChange() string`

GetLastChange returns the LastChange field if non-nil, zero value otherwise.

### GetLastChangeOk

`func (o *Enumeration) GetLastChangeOk() (*string, bool)`

GetLastChangeOk returns a tuple with the LastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChange

`func (o *Enumeration) SetLastChange(v string)`

SetLastChange sets LastChange field to given value.

### HasLastChange

`func (o *Enumeration) HasLastChange() bool`

HasLastChange returns a boolean if a field has been set.

### SetLastChangeNil

`func (o *Enumeration) SetLastChangeNil(b bool)`

 SetLastChangeNil sets the value for LastChange to be an explicit nil

### UnsetLastChange
`func (o *Enumeration) UnsetLastChange()`

UnsetLastChange ensures that no value is present for LastChange, not even an explicit nil
### GetName

`func (o *Enumeration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Enumeration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Enumeration) SetName(v string)`

SetName sets Name field to given value.


### GetMembers

`func (o *Enumeration) GetMembers() map[string]int32`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Enumeration) GetMembersOk() (*map[string]int32, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Enumeration) SetMembers(v map[string]int32)`

SetMembers sets Members field to given value.


### GetArtifactType

`func (o *Enumeration) GetArtifactType() string`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *Enumeration) GetArtifactTypeOk() (*string, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *Enumeration) SetArtifactType(v string)`

SetArtifactType sets ArtifactType field to given value.

### HasArtifactType

`func (o *Enumeration) HasArtifactType() bool`

HasArtifactType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


