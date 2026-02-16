# StackVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastChange** | Pointer to **NullableString** |  | [optional] 
**Offset** | **int32** | Offset of the stack variable | 
**Name** | **string** | Name of the stack variable | 
**Type** | **string** | Data type of the stack variable | 
**Size** | **int32** | Size of the stack variable in bytes | 
**Addr** | **int32** | Memory address of the stack variable | 

## Methods

### NewStackVariable

`func NewStackVariable(offset int32, name string, type_ string, size int32, addr int32, ) *StackVariable`

NewStackVariable instantiates a new StackVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackVariableWithDefaults

`func NewStackVariableWithDefaults() *StackVariable`

NewStackVariableWithDefaults instantiates a new StackVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastChange

`func (o *StackVariable) GetLastChange() string`

GetLastChange returns the LastChange field if non-nil, zero value otherwise.

### GetLastChangeOk

`func (o *StackVariable) GetLastChangeOk() (*string, bool)`

GetLastChangeOk returns a tuple with the LastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChange

`func (o *StackVariable) SetLastChange(v string)`

SetLastChange sets LastChange field to given value.

### HasLastChange

`func (o *StackVariable) HasLastChange() bool`

HasLastChange returns a boolean if a field has been set.

### SetLastChangeNil

`func (o *StackVariable) SetLastChangeNil(b bool)`

 SetLastChangeNil sets the value for LastChange to be an explicit nil

### UnsetLastChange
`func (o *StackVariable) UnsetLastChange()`

UnsetLastChange ensures that no value is present for LastChange, not even an explicit nil
### GetOffset

`func (o *StackVariable) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *StackVariable) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *StackVariable) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetName

`func (o *StackVariable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StackVariable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StackVariable) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *StackVariable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StackVariable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StackVariable) SetType(v string)`

SetType sets Type field to given value.


### GetSize

`func (o *StackVariable) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StackVariable) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StackVariable) SetSize(v int32)`

SetSize sets Size field to given value.


### GetAddr

`func (o *StackVariable) GetAddr() int32`

GetAddr returns the Addr field if non-nil, zero value otherwise.

### GetAddrOk

`func (o *StackVariable) GetAddrOk() (*int32, bool)`

GetAddrOk returns a tuple with the Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddr

`func (o *StackVariable) SetAddr(v int32)`

SetAddr sets Addr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


