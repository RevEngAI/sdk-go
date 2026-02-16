# FunctionTypeInput

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

### NewFunctionTypeInput

`func NewFunctionTypeInput(addr int32, size int32, header FunctionHeader, name string, type_ string, ) *FunctionTypeInput`

NewFunctionTypeInput instantiates a new FunctionTypeInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionTypeInputWithDefaults

`func NewFunctionTypeInputWithDefaults() *FunctionTypeInput`

NewFunctionTypeInputWithDefaults instantiates a new FunctionTypeInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastChange

`func (o *FunctionTypeInput) GetLastChange() string`

GetLastChange returns the LastChange field if non-nil, zero value otherwise.

### GetLastChangeOk

`func (o *FunctionTypeInput) GetLastChangeOk() (*string, bool)`

GetLastChangeOk returns a tuple with the LastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChange

`func (o *FunctionTypeInput) SetLastChange(v string)`

SetLastChange sets LastChange field to given value.

### HasLastChange

`func (o *FunctionTypeInput) HasLastChange() bool`

HasLastChange returns a boolean if a field has been set.

### SetLastChangeNil

`func (o *FunctionTypeInput) SetLastChangeNil(b bool)`

 SetLastChangeNil sets the value for LastChange to be an explicit nil

### UnsetLastChange
`func (o *FunctionTypeInput) UnsetLastChange()`

UnsetLastChange ensures that no value is present for LastChange, not even an explicit nil
### GetAddr

`func (o *FunctionTypeInput) GetAddr() int32`

GetAddr returns the Addr field if non-nil, zero value otherwise.

### GetAddrOk

`func (o *FunctionTypeInput) GetAddrOk() (*int32, bool)`

GetAddrOk returns a tuple with the Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddr

`func (o *FunctionTypeInput) SetAddr(v int32)`

SetAddr sets Addr field to given value.


### GetSize

`func (o *FunctionTypeInput) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FunctionTypeInput) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FunctionTypeInput) SetSize(v int32)`

SetSize sets Size field to given value.


### GetHeader

`func (o *FunctionTypeInput) GetHeader() FunctionHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *FunctionTypeInput) GetHeaderOk() (*FunctionHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *FunctionTypeInput) SetHeader(v FunctionHeader)`

SetHeader sets Header field to given value.


### GetStackVars

`func (o *FunctionTypeInput) GetStackVars() map[string]StackVariable`

GetStackVars returns the StackVars field if non-nil, zero value otherwise.

### GetStackVarsOk

`func (o *FunctionTypeInput) GetStackVarsOk() (*map[string]StackVariable, bool)`

GetStackVarsOk returns a tuple with the StackVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackVars

`func (o *FunctionTypeInput) SetStackVars(v map[string]StackVariable)`

SetStackVars sets StackVars field to given value.

### HasStackVars

`func (o *FunctionTypeInput) HasStackVars() bool`

HasStackVars returns a boolean if a field has been set.

### SetStackVarsNil

`func (o *FunctionTypeInput) SetStackVarsNil(b bool)`

 SetStackVarsNil sets the value for StackVars to be an explicit nil

### UnsetStackVars
`func (o *FunctionTypeInput) UnsetStackVars()`

UnsetStackVars ensures that no value is present for StackVars, not even an explicit nil
### GetName

`func (o *FunctionTypeInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionTypeInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionTypeInput) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *FunctionTypeInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FunctionTypeInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FunctionTypeInput) SetType(v string)`

SetType sets Type field to given value.


### GetArtifactType

`func (o *FunctionTypeInput) GetArtifactType() string`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *FunctionTypeInput) GetArtifactTypeOk() (*string, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *FunctionTypeInput) SetArtifactType(v string)`

SetArtifactType sets ArtifactType field to given value.

### HasArtifactType

`func (o *FunctionTypeInput) HasArtifactType() bool`

HasArtifactType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


