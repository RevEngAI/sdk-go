# FuncDepsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastChange** | Pointer to **string** |  | [optional] 
**Name** | **string** | Name of the global variable | 
**Size** | **int32** | Size of the global variable in bytes | 
**Members** | **map[string]int32** | Dictionary of enumeration members and their values | 
**ArtifactType** | Pointer to **string** | Type of artifact that the global variable is associated with | [optional] 
**Type** | **string** | Data type of the global variable | 
**Addr** | **int32** | Memory address of the global variable | 

## Methods

### NewFuncDepsInner

`func NewFuncDepsInner(name string, size int32, members map[string]int32, type_ string, addr int32, ) *FuncDepsInner`

NewFuncDepsInner instantiates a new FuncDepsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFuncDepsInnerWithDefaults

`func NewFuncDepsInnerWithDefaults() *FuncDepsInner`

NewFuncDepsInnerWithDefaults instantiates a new FuncDepsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastChange

`func (o *FuncDepsInner) GetLastChange() string`

GetLastChange returns the LastChange field if non-nil, zero value otherwise.

### GetLastChangeOk

`func (o *FuncDepsInner) GetLastChangeOk() (*string, bool)`

GetLastChangeOk returns a tuple with the LastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChange

`func (o *FuncDepsInner) SetLastChange(v string)`

SetLastChange sets LastChange field to given value.

### HasLastChange

`func (o *FuncDepsInner) HasLastChange() bool`

HasLastChange returns a boolean if a field has been set.

### GetName

`func (o *FuncDepsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FuncDepsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FuncDepsInner) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *FuncDepsInner) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FuncDepsInner) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FuncDepsInner) SetSize(v int32)`

SetSize sets Size field to given value.


### GetMembers

`func (o *FuncDepsInner) GetMembers() map[string]int32`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *FuncDepsInner) GetMembersOk() (*map[string]int32, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *FuncDepsInner) SetMembers(v map[string]int32)`

SetMembers sets Members field to given value.


### GetArtifactType

`func (o *FuncDepsInner) GetArtifactType() string`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *FuncDepsInner) GetArtifactTypeOk() (*string, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *FuncDepsInner) SetArtifactType(v string)`

SetArtifactType sets ArtifactType field to given value.

### HasArtifactType

`func (o *FuncDepsInner) HasArtifactType() bool`

HasArtifactType returns a boolean if a field has been set.

### GetType

`func (o *FuncDepsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FuncDepsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FuncDepsInner) SetType(v string)`

SetType sets Type field to given value.


### GetAddr

`func (o *FuncDepsInner) GetAddr() int32`

GetAddr returns the Addr field if non-nil, zero value otherwise.

### GetAddrOk

`func (o *FuncDepsInner) GetAddrOk() (*int32, bool)`

GetAddrOk returns a tuple with the Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddr

`func (o *FuncDepsInner) SetAddr(v int32)`

SetAddr sets Addr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


