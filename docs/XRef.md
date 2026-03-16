# XRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **NullableString** |  | 
**XrefTo** | **NullableString** |  | 
**IsScalar** | Pointer to **NullableBool** |  | [optional] 
**IsCall** | Pointer to **NullableBool** |  | [optional] 
**IsData** | Pointer to **NullableBool** |  | [optional] 
**IsString** | Pointer to **NullableBool** |  | [optional] 
**RawData** | Pointer to **Nullable*os.File** |  | [optional] 
**Segment** | Pointer to [**NullableSegmentInfo**](SegmentInfo.md) |  | [optional] 
**OrigStrEncoding** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewXRef

`func NewXRef(value NullableString, xrefTo NullableString, ) *XRef`

NewXRef instantiates a new XRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXRefWithDefaults

`func NewXRefWithDefaults() *XRef`

NewXRefWithDefaults instantiates a new XRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *XRef) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *XRef) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *XRef) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *XRef) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *XRef) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetXrefTo

`func (o *XRef) GetXrefTo() string`

GetXrefTo returns the XrefTo field if non-nil, zero value otherwise.

### GetXrefToOk

`func (o *XRef) GetXrefToOk() (*string, bool)`

GetXrefToOk returns a tuple with the XrefTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXrefTo

`func (o *XRef) SetXrefTo(v string)`

SetXrefTo sets XrefTo field to given value.


### SetXrefToNil

`func (o *XRef) SetXrefToNil(b bool)`

 SetXrefToNil sets the value for XrefTo to be an explicit nil

### UnsetXrefTo
`func (o *XRef) UnsetXrefTo()`

UnsetXrefTo ensures that no value is present for XrefTo, not even an explicit nil
### GetIsScalar

`func (o *XRef) GetIsScalar() bool`

GetIsScalar returns the IsScalar field if non-nil, zero value otherwise.

### GetIsScalarOk

`func (o *XRef) GetIsScalarOk() (*bool, bool)`

GetIsScalarOk returns a tuple with the IsScalar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScalar

`func (o *XRef) SetIsScalar(v bool)`

SetIsScalar sets IsScalar field to given value.

### HasIsScalar

`func (o *XRef) HasIsScalar() bool`

HasIsScalar returns a boolean if a field has been set.

### SetIsScalarNil

`func (o *XRef) SetIsScalarNil(b bool)`

 SetIsScalarNil sets the value for IsScalar to be an explicit nil

### UnsetIsScalar
`func (o *XRef) UnsetIsScalar()`

UnsetIsScalar ensures that no value is present for IsScalar, not even an explicit nil
### GetIsCall

`func (o *XRef) GetIsCall() bool`

GetIsCall returns the IsCall field if non-nil, zero value otherwise.

### GetIsCallOk

`func (o *XRef) GetIsCallOk() (*bool, bool)`

GetIsCallOk returns a tuple with the IsCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCall

`func (o *XRef) SetIsCall(v bool)`

SetIsCall sets IsCall field to given value.

### HasIsCall

`func (o *XRef) HasIsCall() bool`

HasIsCall returns a boolean if a field has been set.

### SetIsCallNil

`func (o *XRef) SetIsCallNil(b bool)`

 SetIsCallNil sets the value for IsCall to be an explicit nil

### UnsetIsCall
`func (o *XRef) UnsetIsCall()`

UnsetIsCall ensures that no value is present for IsCall, not even an explicit nil
### GetIsData

`func (o *XRef) GetIsData() bool`

GetIsData returns the IsData field if non-nil, zero value otherwise.

### GetIsDataOk

`func (o *XRef) GetIsDataOk() (*bool, bool)`

GetIsDataOk returns a tuple with the IsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsData

`func (o *XRef) SetIsData(v bool)`

SetIsData sets IsData field to given value.

### HasIsData

`func (o *XRef) HasIsData() bool`

HasIsData returns a boolean if a field has been set.

### SetIsDataNil

`func (o *XRef) SetIsDataNil(b bool)`

 SetIsDataNil sets the value for IsData to be an explicit nil

### UnsetIsData
`func (o *XRef) UnsetIsData()`

UnsetIsData ensures that no value is present for IsData, not even an explicit nil
### GetIsString

`func (o *XRef) GetIsString() bool`

GetIsString returns the IsString field if non-nil, zero value otherwise.

### GetIsStringOk

`func (o *XRef) GetIsStringOk() (*bool, bool)`

GetIsStringOk returns a tuple with the IsString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsString

`func (o *XRef) SetIsString(v bool)`

SetIsString sets IsString field to given value.

### HasIsString

`func (o *XRef) HasIsString() bool`

HasIsString returns a boolean if a field has been set.

### SetIsStringNil

`func (o *XRef) SetIsStringNil(b bool)`

 SetIsStringNil sets the value for IsString to be an explicit nil

### UnsetIsString
`func (o *XRef) UnsetIsString()`

UnsetIsString ensures that no value is present for IsString, not even an explicit nil
### GetRawData

`func (o *XRef) GetRawData() *os.File`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *XRef) GetRawDataOk() (**os.File, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *XRef) SetRawData(v *os.File)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *XRef) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### SetRawDataNil

`func (o *XRef) SetRawDataNil(b bool)`

 SetRawDataNil sets the value for RawData to be an explicit nil

### UnsetRawData
`func (o *XRef) UnsetRawData()`

UnsetRawData ensures that no value is present for RawData, not even an explicit nil
### GetSegment

`func (o *XRef) GetSegment() SegmentInfo`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *XRef) GetSegmentOk() (*SegmentInfo, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *XRef) SetSegment(v SegmentInfo)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *XRef) HasSegment() bool`

HasSegment returns a boolean if a field has been set.

### SetSegmentNil

`func (o *XRef) SetSegmentNil(b bool)`

 SetSegmentNil sets the value for Segment to be an explicit nil

### UnsetSegment
`func (o *XRef) UnsetSegment()`

UnsetSegment ensures that no value is present for Segment, not even an explicit nil
### GetOrigStrEncoding

`func (o *XRef) GetOrigStrEncoding() string`

GetOrigStrEncoding returns the OrigStrEncoding field if non-nil, zero value otherwise.

### GetOrigStrEncodingOk

`func (o *XRef) GetOrigStrEncodingOk() (*string, bool)`

GetOrigStrEncodingOk returns a tuple with the OrigStrEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigStrEncoding

`func (o *XRef) SetOrigStrEncoding(v string)`

SetOrigStrEncoding sets OrigStrEncoding field to given value.

### HasOrigStrEncoding

`func (o *XRef) HasOrigStrEncoding() bool`

HasOrigStrEncoding returns a boolean if a field has been set.

### SetOrigStrEncodingNil

`func (o *XRef) SetOrigStrEncodingNil(b bool)`

 SetOrigStrEncodingNil sets the value for OrigStrEncoding to be an explicit nil

### UnsetOrigStrEncoding
`func (o *XRef) UnsetOrigStrEncoding()`

UnsetOrigStrEncoding ensures that no value is present for OrigStrEncoding, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


