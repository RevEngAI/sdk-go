# FunctionString

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | The value of the string literal | 
**Vaddr** | **int32** | The vaddr of the string value | 

## Methods

### NewFunctionString

`func NewFunctionString(value string, vaddr int32, ) *FunctionString`

NewFunctionString instantiates a new FunctionString object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionStringWithDefaults

`func NewFunctionStringWithDefaults() *FunctionString`

NewFunctionStringWithDefaults instantiates a new FunctionString object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *FunctionString) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FunctionString) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FunctionString) SetValue(v string)`

SetValue sets Value field to given value.


### GetVaddr

`func (o *FunctionString) GetVaddr() int32`

GetVaddr returns the Vaddr field if non-nil, zero value otherwise.

### GetVaddrOk

`func (o *FunctionString) GetVaddrOk() (*int32, bool)`

GetVaddrOk returns a tuple with the Vaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaddr

`func (o *FunctionString) SetVaddr(v int32)`

SetVaddr sets Vaddr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


