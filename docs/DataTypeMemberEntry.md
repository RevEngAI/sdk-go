# DataTypeMemberEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BitOffset** | Pointer to **int64** | Bit offset from the start of the containing type, for bitfields. | [optional] 
**BitSize** | Pointer to **int64** | Width in bits, for bitfields. | [optional] 
**DataTypeId** | Pointer to **int64** | The member&#39;s type. | [optional] 
**IsBitfield** | **bool** | Whether this member is a bitfield. | 
**Name** | Pointer to **string** | Member name, absent for unnamed padding. | [optional] 
**Offset** | **int64** | Byte offset from the start of the containing type. | 
**Size** | **int64** | Member size in bytes. | 

## Methods

### NewDataTypeMemberEntry

`func NewDataTypeMemberEntry(isBitfield bool, offset int64, size int64, ) *DataTypeMemberEntry`

NewDataTypeMemberEntry instantiates a new DataTypeMemberEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTypeMemberEntryWithDefaults

`func NewDataTypeMemberEntryWithDefaults() *DataTypeMemberEntry`

NewDataTypeMemberEntryWithDefaults instantiates a new DataTypeMemberEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBitOffset

`func (o *DataTypeMemberEntry) GetBitOffset() int64`

GetBitOffset returns the BitOffset field if non-nil, zero value otherwise.

### GetBitOffsetOk

`func (o *DataTypeMemberEntry) GetBitOffsetOk() (*int64, bool)`

GetBitOffsetOk returns a tuple with the BitOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitOffset

`func (o *DataTypeMemberEntry) SetBitOffset(v int64)`

SetBitOffset sets BitOffset field to given value.

### HasBitOffset

`func (o *DataTypeMemberEntry) HasBitOffset() bool`

HasBitOffset returns a boolean if a field has been set.

### GetBitSize

`func (o *DataTypeMemberEntry) GetBitSize() int64`

GetBitSize returns the BitSize field if non-nil, zero value otherwise.

### GetBitSizeOk

`func (o *DataTypeMemberEntry) GetBitSizeOk() (*int64, bool)`

GetBitSizeOk returns a tuple with the BitSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitSize

`func (o *DataTypeMemberEntry) SetBitSize(v int64)`

SetBitSize sets BitSize field to given value.

### HasBitSize

`func (o *DataTypeMemberEntry) HasBitSize() bool`

HasBitSize returns a boolean if a field has been set.

### GetDataTypeId

`func (o *DataTypeMemberEntry) GetDataTypeId() int64`

GetDataTypeId returns the DataTypeId field if non-nil, zero value otherwise.

### GetDataTypeIdOk

`func (o *DataTypeMemberEntry) GetDataTypeIdOk() (*int64, bool)`

GetDataTypeIdOk returns a tuple with the DataTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypeId

`func (o *DataTypeMemberEntry) SetDataTypeId(v int64)`

SetDataTypeId sets DataTypeId field to given value.

### HasDataTypeId

`func (o *DataTypeMemberEntry) HasDataTypeId() bool`

HasDataTypeId returns a boolean if a field has been set.

### GetIsBitfield

`func (o *DataTypeMemberEntry) GetIsBitfield() bool`

GetIsBitfield returns the IsBitfield field if non-nil, zero value otherwise.

### GetIsBitfieldOk

`func (o *DataTypeMemberEntry) GetIsBitfieldOk() (*bool, bool)`

GetIsBitfieldOk returns a tuple with the IsBitfield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBitfield

`func (o *DataTypeMemberEntry) SetIsBitfield(v bool)`

SetIsBitfield sets IsBitfield field to given value.


### GetName

`func (o *DataTypeMemberEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataTypeMemberEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataTypeMemberEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataTypeMemberEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOffset

`func (o *DataTypeMemberEntry) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *DataTypeMemberEntry) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *DataTypeMemberEntry) SetOffset(v int64)`

SetOffset sets Offset field to given value.


### GetSize

`func (o *DataTypeMemberEntry) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DataTypeMemberEntry) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DataTypeMemberEntry) SetSize(v int64)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


