# NameSourceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The source (process) the function name came from | 
**FunctionId** | Pointer to **NullableInt32** |  | [optional] 
**BinaryId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewNameSourceType

`func NewNameSourceType(type_ string, ) *NameSourceType`

NewNameSourceType instantiates a new NameSourceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNameSourceTypeWithDefaults

`func NewNameSourceTypeWithDefaults() *NameSourceType`

NewNameSourceTypeWithDefaults instantiates a new NameSourceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NameSourceType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NameSourceType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NameSourceType) SetType(v string)`

SetType sets Type field to given value.


### GetFunctionId

`func (o *NameSourceType) GetFunctionId() int32`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *NameSourceType) GetFunctionIdOk() (*int32, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *NameSourceType) SetFunctionId(v int32)`

SetFunctionId sets FunctionId field to given value.

### HasFunctionId

`func (o *NameSourceType) HasFunctionId() bool`

HasFunctionId returns a boolean if a field has been set.

### SetFunctionIdNil

`func (o *NameSourceType) SetFunctionIdNil(b bool)`

 SetFunctionIdNil sets the value for FunctionId to be an explicit nil

### UnsetFunctionId
`func (o *NameSourceType) UnsetFunctionId()`

UnsetFunctionId ensures that no value is present for FunctionId, not even an explicit nil
### GetBinaryId

`func (o *NameSourceType) GetBinaryId() int32`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *NameSourceType) GetBinaryIdOk() (*int32, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *NameSourceType) SetBinaryId(v int32)`

SetBinaryId sets BinaryId field to given value.

### HasBinaryId

`func (o *NameSourceType) HasBinaryId() bool`

HasBinaryId returns a boolean if a field has been set.

### SetBinaryIdNil

`func (o *NameSourceType) SetBinaryIdNil(b bool)`

 SetBinaryIdNil sets the value for BinaryId to be an explicit nil

### UnsetBinaryId
`func (o *NameSourceType) UnsetBinaryId()`

UnsetBinaryId ensures that no value is present for BinaryId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


