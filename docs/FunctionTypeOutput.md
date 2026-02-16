# FunctionTypeOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastChange** | Pointer to **NullableString** |  | [optional] 
**Addr** | **int32** | Memory address of the function | 
**Size** | **int32** | Size of the function in bytes | 
**Header** | [**FunctionHeader**](FunctionHeader.md) | Function header information | 
**StackVars** | Pointer to [**map[string]StackVariable**](StackVariable.md) |  | [optional] 
**Name** | **string** | Name of the function | 
**Type** | **string** | Return type of the function | 
**ArtifactType** | Pointer to **string** | Type of artifact that the structure is associated with | [optional] [default to "Function"]

## Methods

### NewFunctionTypeOutput

`func NewFunctionTypeOutput(addr int32, size int32, header FunctionHeader, name string, type_ string, ) *FunctionTypeOutput`

NewFunctionTypeOutput instantiates a new FunctionTypeOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionTypeOutputWithDefaults

`func NewFunctionTypeOutputWithDefaults() *FunctionTypeOutput`

NewFunctionTypeOutputWithDefaults instantiates a new FunctionTypeOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastChange

`func (o *FunctionTypeOutput) GetLastChange() string`

GetLastChange returns the LastChange field if non-nil, zero value otherwise.

### GetLastChangeOk

`func (o *FunctionTypeOutput) GetLastChangeOk() (*string, bool)`

GetLastChangeOk returns a tuple with the LastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChange

`func (o *FunctionTypeOutput) SetLastChange(v string)`

SetLastChange sets LastChange field to given value.

### HasLastChange

`func (o *FunctionTypeOutput) HasLastChange() bool`

HasLastChange returns a boolean if a field has been set.

### SetLastChangeNil

`func (o *FunctionTypeOutput) SetLastChangeNil(b bool)`

 SetLastChangeNil sets the value for LastChange to be an explicit nil

### UnsetLastChange
`func (o *FunctionTypeOutput) UnsetLastChange()`

UnsetLastChange ensures that no value is present for LastChange, not even an explicit nil
### GetAddr

`func (o *FunctionTypeOutput) GetAddr() int32`

GetAddr returns the Addr field if non-nil, zero value otherwise.

### GetAddrOk

`func (o *FunctionTypeOutput) GetAddrOk() (*int32, bool)`

GetAddrOk returns a tuple with the Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddr

`func (o *FunctionTypeOutput) SetAddr(v int32)`

SetAddr sets Addr field to given value.


### GetSize

`func (o *FunctionTypeOutput) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FunctionTypeOutput) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FunctionTypeOutput) SetSize(v int32)`

SetSize sets Size field to given value.


### GetHeader

`func (o *FunctionTypeOutput) GetHeader() FunctionHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *FunctionTypeOutput) GetHeaderOk() (*FunctionHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *FunctionTypeOutput) SetHeader(v FunctionHeader)`

SetHeader sets Header field to given value.


### GetStackVars

`func (o *FunctionTypeOutput) GetStackVars() map[string]StackVariable`

GetStackVars returns the StackVars field if non-nil, zero value otherwise.

### GetStackVarsOk

`func (o *FunctionTypeOutput) GetStackVarsOk() (*map[string]StackVariable, bool)`

GetStackVarsOk returns a tuple with the StackVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackVars

`func (o *FunctionTypeOutput) SetStackVars(v map[string]StackVariable)`

SetStackVars sets StackVars field to given value.

### HasStackVars

`func (o *FunctionTypeOutput) HasStackVars() bool`

HasStackVars returns a boolean if a field has been set.

### SetStackVarsNil

`func (o *FunctionTypeOutput) SetStackVarsNil(b bool)`

 SetStackVarsNil sets the value for StackVars to be an explicit nil

### UnsetStackVars
`func (o *FunctionTypeOutput) UnsetStackVars()`

UnsetStackVars ensures that no value is present for StackVars, not even an explicit nil
### GetName

`func (o *FunctionTypeOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionTypeOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionTypeOutput) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *FunctionTypeOutput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FunctionTypeOutput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FunctionTypeOutput) SetType(v string)`

SetType sets Type field to given value.


### GetArtifactType

`func (o *FunctionTypeOutput) GetArtifactType() string`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *FunctionTypeOutput) GetArtifactTypeOk() (*string, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *FunctionTypeOutput) SetArtifactType(v string)`

SetArtifactType sets ArtifactType field to given value.

### HasArtifactType

`func (o *FunctionTypeOutput) HasArtifactType() bool`

HasArtifactType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


