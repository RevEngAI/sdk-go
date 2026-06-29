# FunctionType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addr** | **int64** |  | 
**ArtifactType** | Pointer to **string** |  | [optional] 
**Header** | [**FunctionHeader**](FunctionHeader.md) |  | 
**LastChange** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Scope** | Pointer to **string** |  | [optional] 
**Size** | **int64** |  | 
**StackVars** | Pointer to [**map[string]FunctionStackVariable**](FunctionStackVariable.md) | Stack variables keyed by offset hex. | [optional] 
**Type** | **string** |  | 

## Methods

### NewFunctionType

`func NewFunctionType(addr int64, header FunctionHeader, name string, size int64, type_ string, ) *FunctionType`

NewFunctionType instantiates a new FunctionType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionTypeWithDefaults

`func NewFunctionTypeWithDefaults() *FunctionType`

NewFunctionTypeWithDefaults instantiates a new FunctionType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddr

`func (o *FunctionType) GetAddr() int64`

GetAddr returns the Addr field if non-nil, zero value otherwise.

### GetAddrOk

`func (o *FunctionType) GetAddrOk() (*int64, bool)`

GetAddrOk returns a tuple with the Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddr

`func (o *FunctionType) SetAddr(v int64)`

SetAddr sets Addr field to given value.


### GetArtifactType

`func (o *FunctionType) GetArtifactType() string`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *FunctionType) GetArtifactTypeOk() (*string, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *FunctionType) SetArtifactType(v string)`

SetArtifactType sets ArtifactType field to given value.

### HasArtifactType

`func (o *FunctionType) HasArtifactType() bool`

HasArtifactType returns a boolean if a field has been set.

### GetHeader

`func (o *FunctionType) GetHeader() FunctionHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *FunctionType) GetHeaderOk() (*FunctionHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *FunctionType) SetHeader(v FunctionHeader)`

SetHeader sets Header field to given value.


### GetLastChange

`func (o *FunctionType) GetLastChange() string`

GetLastChange returns the LastChange field if non-nil, zero value otherwise.

### GetLastChangeOk

`func (o *FunctionType) GetLastChangeOk() (*string, bool)`

GetLastChangeOk returns a tuple with the LastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChange

`func (o *FunctionType) SetLastChange(v string)`

SetLastChange sets LastChange field to given value.

### HasLastChange

`func (o *FunctionType) HasLastChange() bool`

HasLastChange returns a boolean if a field has been set.

### GetName

`func (o *FunctionType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionType) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *FunctionType) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *FunctionType) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *FunctionType) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *FunctionType) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSize

`func (o *FunctionType) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FunctionType) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FunctionType) SetSize(v int64)`

SetSize sets Size field to given value.


### GetStackVars

`func (o *FunctionType) GetStackVars() map[string]FunctionStackVariable`

GetStackVars returns the StackVars field if non-nil, zero value otherwise.

### GetStackVarsOk

`func (o *FunctionType) GetStackVarsOk() (*map[string]FunctionStackVariable, bool)`

GetStackVarsOk returns a tuple with the StackVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackVars

`func (o *FunctionType) SetStackVars(v map[string]FunctionStackVariable)`

SetStackVars sets StackVars field to given value.

### HasStackVars

`func (o *FunctionType) HasStackVars() bool`

HasStackVars returns a boolean if a field has been set.

### GetType

`func (o *FunctionType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FunctionType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FunctionType) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


