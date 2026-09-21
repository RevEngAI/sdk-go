# Meta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endianness** | Pointer to [**Endianness**](Endianness.md) |  | [optional] [default to ENDIANNESS_UNSPECIFIED]

## Methods

### NewMeta

`func NewMeta() *Meta`

NewMeta instantiates a new Meta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaWithDefaults

`func NewMetaWithDefaults() *Meta`

NewMetaWithDefaults instantiates a new Meta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndianness

`func (o *Meta) GetEndianness() Endianness`

GetEndianness returns the Endianness field if non-nil, zero value otherwise.

### GetEndiannessOk

`func (o *Meta) GetEndiannessOk() (*Endianness, bool)`

GetEndiannessOk returns a tuple with the Endianness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndianness

`func (o *Meta) SetEndianness(v Endianness)`

SetEndianness sets Endianness field to given value.

### HasEndianness

`func (o *Meta) HasEndianness() bool`

HasEndianness returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


