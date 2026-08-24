# RenderedToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTypeId** | Pointer to **int64** | Data type this token names, or for a field the type that declares it. Scoped by the response&#39;s analysis_id. | [optional] 
**FunctionId** | Pointer to **int64** | Function this token calls or names. Absent when the address reaches no function of this binary. | [optional] 
**ImportedFunctionId** | Pointer to **int64** | Imported function this token calls. Set instead of function_id for an external call. | [optional] 
**Kind** | **string** | What the token names. | 
**Vaddr** | Pointer to **int64** | Virtual address the token resolves to. Absent for a token with no address. | [optional] 
**Value** | **string** | Name the token resolves to. | 

## Methods

### NewRenderedToken

`func NewRenderedToken(kind string, value string, ) *RenderedToken`

NewRenderedToken instantiates a new RenderedToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenderedTokenWithDefaults

`func NewRenderedTokenWithDefaults() *RenderedToken`

NewRenderedTokenWithDefaults instantiates a new RenderedToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTypeId

`func (o *RenderedToken) GetDataTypeId() int64`

GetDataTypeId returns the DataTypeId field if non-nil, zero value otherwise.

### GetDataTypeIdOk

`func (o *RenderedToken) GetDataTypeIdOk() (*int64, bool)`

GetDataTypeIdOk returns a tuple with the DataTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypeId

`func (o *RenderedToken) SetDataTypeId(v int64)`

SetDataTypeId sets DataTypeId field to given value.

### HasDataTypeId

`func (o *RenderedToken) HasDataTypeId() bool`

HasDataTypeId returns a boolean if a field has been set.

### GetFunctionId

`func (o *RenderedToken) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *RenderedToken) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *RenderedToken) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.

### HasFunctionId

`func (o *RenderedToken) HasFunctionId() bool`

HasFunctionId returns a boolean if a field has been set.

### GetImportedFunctionId

`func (o *RenderedToken) GetImportedFunctionId() int64`

GetImportedFunctionId returns the ImportedFunctionId field if non-nil, zero value otherwise.

### GetImportedFunctionIdOk

`func (o *RenderedToken) GetImportedFunctionIdOk() (*int64, bool)`

GetImportedFunctionIdOk returns a tuple with the ImportedFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedFunctionId

`func (o *RenderedToken) SetImportedFunctionId(v int64)`

SetImportedFunctionId sets ImportedFunctionId field to given value.

### HasImportedFunctionId

`func (o *RenderedToken) HasImportedFunctionId() bool`

HasImportedFunctionId returns a boolean if a field has been set.

### GetKind

`func (o *RenderedToken) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RenderedToken) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RenderedToken) SetKind(v string)`

SetKind sets Kind field to given value.


### GetVaddr

`func (o *RenderedToken) GetVaddr() int64`

GetVaddr returns the Vaddr field if non-nil, zero value otherwise.

### GetVaddrOk

`func (o *RenderedToken) GetVaddrOk() (*int64, bool)`

GetVaddrOk returns a tuple with the Vaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaddr

`func (o *RenderedToken) SetVaddr(v int64)`

SetVaddr sets Vaddr field to given value.

### HasVaddr

`func (o *RenderedToken) HasVaddr() bool`

HasVaddr returns a boolean if a field has been set.

### GetValue

`func (o *RenderedToken) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RenderedToken) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RenderedToken) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


