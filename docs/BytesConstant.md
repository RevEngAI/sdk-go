# BytesConstant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** |  | [optional] [default to "bytes"]
**Value** | **string** | Hexadecimal byte pairs representing one exact, contiguous binary value. | 
**Display** | Pointer to [**Display**](Display.md) |  | [optional] [default to DISPLAY_HEX]

## Methods

### NewBytesConstant

`func NewBytesConstant(value string, ) *BytesConstant`

NewBytesConstant instantiates a new BytesConstant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBytesConstantWithDefaults

`func NewBytesConstantWithDefaults() *BytesConstant`

NewBytesConstantWithDefaults instantiates a new BytesConstant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *BytesConstant) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *BytesConstant) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *BytesConstant) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *BytesConstant) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetValue

`func (o *BytesConstant) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BytesConstant) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BytesConstant) SetValue(v string)`

SetValue sets Value field to given value.


### GetDisplay

`func (o *BytesConstant) GetDisplay() Display`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BytesConstant) GetDisplayOk() (*Display, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BytesConstant) SetDisplay(v Display)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *BytesConstant) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


