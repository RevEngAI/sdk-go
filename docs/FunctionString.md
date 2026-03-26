# FunctionString

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to [**StringSource**](StringSource.md) | The source of the string | [optional] [default to STRINGSOURCE_SYSTEM]
**Vaddr** | **int32** | The vaddr of the string value | 
**Value** | **string** | The value of the string literal | 

## Methods

### NewFunctionString

`func NewFunctionString(vaddr int32, value string, ) *FunctionString`

NewFunctionString instantiates a new FunctionString object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionStringWithDefaults

`func NewFunctionStringWithDefaults() *FunctionString`

NewFunctionStringWithDefaults instantiates a new FunctionString object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *FunctionString) GetSource() StringSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FunctionString) GetSourceOk() (*StringSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FunctionString) SetSource(v StringSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *FunctionString) HasSource() bool`

HasSource returns a boolean if a field has been set.

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


