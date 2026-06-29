# FunctionStackVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addr** | **int64** |  | 
**LastChange** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Offset** | **int64** |  | 
**Scope** | Pointer to **string** |  | [optional] 
**Size** | **int64** |  | 
**Type** | **string** |  | 

## Methods

### NewFunctionStackVariable

`func NewFunctionStackVariable(addr int64, name string, offset int64, size int64, type_ string, ) *FunctionStackVariable`

NewFunctionStackVariable instantiates a new FunctionStackVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionStackVariableWithDefaults

`func NewFunctionStackVariableWithDefaults() *FunctionStackVariable`

NewFunctionStackVariableWithDefaults instantiates a new FunctionStackVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddr

`func (o *FunctionStackVariable) GetAddr() int64`

GetAddr returns the Addr field if non-nil, zero value otherwise.

### GetAddrOk

`func (o *FunctionStackVariable) GetAddrOk() (*int64, bool)`

GetAddrOk returns a tuple with the Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddr

`func (o *FunctionStackVariable) SetAddr(v int64)`

SetAddr sets Addr field to given value.


### GetLastChange

`func (o *FunctionStackVariable) GetLastChange() string`

GetLastChange returns the LastChange field if non-nil, zero value otherwise.

### GetLastChangeOk

`func (o *FunctionStackVariable) GetLastChangeOk() (*string, bool)`

GetLastChangeOk returns a tuple with the LastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChange

`func (o *FunctionStackVariable) SetLastChange(v string)`

SetLastChange sets LastChange field to given value.

### HasLastChange

`func (o *FunctionStackVariable) HasLastChange() bool`

HasLastChange returns a boolean if a field has been set.

### GetName

`func (o *FunctionStackVariable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionStackVariable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionStackVariable) SetName(v string)`

SetName sets Name field to given value.


### GetOffset

`func (o *FunctionStackVariable) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *FunctionStackVariable) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *FunctionStackVariable) SetOffset(v int64)`

SetOffset sets Offset field to given value.


### GetScope

`func (o *FunctionStackVariable) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *FunctionStackVariable) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *FunctionStackVariable) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *FunctionStackVariable) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSize

`func (o *FunctionStackVariable) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FunctionStackVariable) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FunctionStackVariable) SetSize(v int64)`

SetSize sets Size field to given value.


### GetType

`func (o *FunctionStackVariable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FunctionStackVariable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FunctionStackVariable) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


