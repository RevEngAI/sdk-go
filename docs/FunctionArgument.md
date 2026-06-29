# FunctionArgument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastChange** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Offset** | **int64** |  | 
**Scope** | Pointer to **string** |  | [optional] 
**Size** | **int64** |  | 
**Type** | **string** |  | 

## Methods

### NewFunctionArgument

`func NewFunctionArgument(name string, offset int64, size int64, type_ string, ) *FunctionArgument`

NewFunctionArgument instantiates a new FunctionArgument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionArgumentWithDefaults

`func NewFunctionArgumentWithDefaults() *FunctionArgument`

NewFunctionArgumentWithDefaults instantiates a new FunctionArgument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastChange

`func (o *FunctionArgument) GetLastChange() string`

GetLastChange returns the LastChange field if non-nil, zero value otherwise.

### GetLastChangeOk

`func (o *FunctionArgument) GetLastChangeOk() (*string, bool)`

GetLastChangeOk returns a tuple with the LastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChange

`func (o *FunctionArgument) SetLastChange(v string)`

SetLastChange sets LastChange field to given value.

### HasLastChange

`func (o *FunctionArgument) HasLastChange() bool`

HasLastChange returns a boolean if a field has been set.

### GetName

`func (o *FunctionArgument) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionArgument) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionArgument) SetName(v string)`

SetName sets Name field to given value.


### GetOffset

`func (o *FunctionArgument) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *FunctionArgument) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *FunctionArgument) SetOffset(v int64)`

SetOffset sets Offset field to given value.


### GetScope

`func (o *FunctionArgument) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *FunctionArgument) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *FunctionArgument) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *FunctionArgument) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSize

`func (o *FunctionArgument) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FunctionArgument) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FunctionArgument) SetSize(v int64)`

SetSize sets Size field to given value.


### GetType

`func (o *FunctionArgument) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FunctionArgument) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FunctionArgument) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


