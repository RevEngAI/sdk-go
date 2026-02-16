# GlobalVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastChange** | Pointer to **NullableString** |  | [optional] 
**Addr** | **int32** | Memory address of the global variable | 
**Name** | **string** | Name of the global variable | 
**Type** | **string** | Data type of the global variable | 
**Size** | **int32** | Size of the global variable in bytes | 
**ArtifactType** | Pointer to **string** | Type of artifact that the global variable is associated with | [optional] 

## Methods

### NewGlobalVariable

`func NewGlobalVariable(addr int32, name string, type_ string, size int32, ) *GlobalVariable`

NewGlobalVariable instantiates a new GlobalVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalVariableWithDefaults

`func NewGlobalVariableWithDefaults() *GlobalVariable`

NewGlobalVariableWithDefaults instantiates a new GlobalVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastChange

`func (o *GlobalVariable) GetLastChange() string`

GetLastChange returns the LastChange field if non-nil, zero value otherwise.

### GetLastChangeOk

`func (o *GlobalVariable) GetLastChangeOk() (*string, bool)`

GetLastChangeOk returns a tuple with the LastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChange

`func (o *GlobalVariable) SetLastChange(v string)`

SetLastChange sets LastChange field to given value.

### HasLastChange

`func (o *GlobalVariable) HasLastChange() bool`

HasLastChange returns a boolean if a field has been set.

### SetLastChangeNil

`func (o *GlobalVariable) SetLastChangeNil(b bool)`

 SetLastChangeNil sets the value for LastChange to be an explicit nil

### UnsetLastChange
`func (o *GlobalVariable) UnsetLastChange()`

UnsetLastChange ensures that no value is present for LastChange, not even an explicit nil
### GetAddr

`func (o *GlobalVariable) GetAddr() int32`

GetAddr returns the Addr field if non-nil, zero value otherwise.

### GetAddrOk

`func (o *GlobalVariable) GetAddrOk() (*int32, bool)`

GetAddrOk returns a tuple with the Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddr

`func (o *GlobalVariable) SetAddr(v int32)`

SetAddr sets Addr field to given value.


### GetName

`func (o *GlobalVariable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlobalVariable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlobalVariable) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *GlobalVariable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GlobalVariable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GlobalVariable) SetType(v string)`

SetType sets Type field to given value.


### GetSize

`func (o *GlobalVariable) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GlobalVariable) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GlobalVariable) SetSize(v int32)`

SetSize sets Size field to given value.


### GetArtifactType

`func (o *GlobalVariable) GetArtifactType() string`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *GlobalVariable) GetArtifactTypeOk() (*string, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *GlobalVariable) SetArtifactType(v string)`

SetArtifactType sets ArtifactType field to given value.

### HasArtifactType

`func (o *GlobalVariable) HasArtifactType() bool`

HasArtifactType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


