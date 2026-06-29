# FunctionDependency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addr** | Pointer to **int64** | Memory address (GlobalVariable). | [optional] 
**ArtifactType** | Pointer to **string** |  | [optional] 
**LastChange** | Pointer to **string** |  | [optional] 
**Members** | Pointer to **interface{}** |  | [optional] 
**Name** | **string** |  | 
**Scope** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int64** | Total byte size (Struct, GlobalVariable). | [optional] 
**Type** | Pointer to **string** | Underlying type (TypeDefinition, GlobalVariable). | [optional] 

## Methods

### NewFunctionDependency

`func NewFunctionDependency(name string, ) *FunctionDependency`

NewFunctionDependency instantiates a new FunctionDependency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionDependencyWithDefaults

`func NewFunctionDependencyWithDefaults() *FunctionDependency`

NewFunctionDependencyWithDefaults instantiates a new FunctionDependency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddr

`func (o *FunctionDependency) GetAddr() int64`

GetAddr returns the Addr field if non-nil, zero value otherwise.

### GetAddrOk

`func (o *FunctionDependency) GetAddrOk() (*int64, bool)`

GetAddrOk returns a tuple with the Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddr

`func (o *FunctionDependency) SetAddr(v int64)`

SetAddr sets Addr field to given value.

### HasAddr

`func (o *FunctionDependency) HasAddr() bool`

HasAddr returns a boolean if a field has been set.

### GetArtifactType

`func (o *FunctionDependency) GetArtifactType() string`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *FunctionDependency) GetArtifactTypeOk() (*string, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *FunctionDependency) SetArtifactType(v string)`

SetArtifactType sets ArtifactType field to given value.

### HasArtifactType

`func (o *FunctionDependency) HasArtifactType() bool`

HasArtifactType returns a boolean if a field has been set.

### GetLastChange

`func (o *FunctionDependency) GetLastChange() string`

GetLastChange returns the LastChange field if non-nil, zero value otherwise.

### GetLastChangeOk

`func (o *FunctionDependency) GetLastChangeOk() (*string, bool)`

GetLastChangeOk returns a tuple with the LastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChange

`func (o *FunctionDependency) SetLastChange(v string)`

SetLastChange sets LastChange field to given value.

### HasLastChange

`func (o *FunctionDependency) HasLastChange() bool`

HasLastChange returns a boolean if a field has been set.

### GetMembers

`func (o *FunctionDependency) GetMembers() interface{}`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *FunctionDependency) GetMembersOk() (*interface{}, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *FunctionDependency) SetMembers(v interface{})`

SetMembers sets Members field to given value.

### HasMembers

`func (o *FunctionDependency) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembersNil

`func (o *FunctionDependency) SetMembersNil(b bool)`

 SetMembersNil sets the value for Members to be an explicit nil

### UnsetMembers
`func (o *FunctionDependency) UnsetMembers()`

UnsetMembers ensures that no value is present for Members, not even an explicit nil
### GetName

`func (o *FunctionDependency) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionDependency) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionDependency) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *FunctionDependency) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *FunctionDependency) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *FunctionDependency) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *FunctionDependency) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSize

`func (o *FunctionDependency) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FunctionDependency) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FunctionDependency) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *FunctionDependency) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *FunctionDependency) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FunctionDependency) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FunctionDependency) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FunctionDependency) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


