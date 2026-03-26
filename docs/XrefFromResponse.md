# XrefFromResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsCall** | Pointer to **NullableBool** |  | [optional] 
**IsData** | Pointer to **NullableBool** |  | [optional] 
**IsScalar** | Pointer to **NullableBool** |  | [optional] 
**IsString** | Pointer to **NullableBool** |  | [optional] 
**OrigStrEncoding** | Pointer to **NullableString** |  | [optional] 
**RawData** | Pointer to **NullableString** |  | [optional] 
**Segment** | Pointer to [**NullableSegmentInfo**](SegmentInfo.md) |  | [optional] 
**Value** | **NullableString** |  | 
**XrefTo** | **NullableString** |  | 

## Methods

### NewXrefFromResponse

`func NewXrefFromResponse(value NullableString, xrefTo NullableString, ) *XrefFromResponse`

NewXrefFromResponse instantiates a new XrefFromResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXrefFromResponseWithDefaults

`func NewXrefFromResponseWithDefaults() *XrefFromResponse`

NewXrefFromResponseWithDefaults instantiates a new XrefFromResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsCall

`func (o *XrefFromResponse) GetIsCall() bool`

GetIsCall returns the IsCall field if non-nil, zero value otherwise.

### GetIsCallOk

`func (o *XrefFromResponse) GetIsCallOk() (*bool, bool)`

GetIsCallOk returns a tuple with the IsCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCall

`func (o *XrefFromResponse) SetIsCall(v bool)`

SetIsCall sets IsCall field to given value.

### HasIsCall

`func (o *XrefFromResponse) HasIsCall() bool`

HasIsCall returns a boolean if a field has been set.

### SetIsCallNil

`func (o *XrefFromResponse) SetIsCallNil(b bool)`

 SetIsCallNil sets the value for IsCall to be an explicit nil

### UnsetIsCall
`func (o *XrefFromResponse) UnsetIsCall()`

UnsetIsCall ensures that no value is present for IsCall, not even an explicit nil
### GetIsData

`func (o *XrefFromResponse) GetIsData() bool`

GetIsData returns the IsData field if non-nil, zero value otherwise.

### GetIsDataOk

`func (o *XrefFromResponse) GetIsDataOk() (*bool, bool)`

GetIsDataOk returns a tuple with the IsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsData

`func (o *XrefFromResponse) SetIsData(v bool)`

SetIsData sets IsData field to given value.

### HasIsData

`func (o *XrefFromResponse) HasIsData() bool`

HasIsData returns a boolean if a field has been set.

### SetIsDataNil

`func (o *XrefFromResponse) SetIsDataNil(b bool)`

 SetIsDataNil sets the value for IsData to be an explicit nil

### UnsetIsData
`func (o *XrefFromResponse) UnsetIsData()`

UnsetIsData ensures that no value is present for IsData, not even an explicit nil
### GetIsScalar

`func (o *XrefFromResponse) GetIsScalar() bool`

GetIsScalar returns the IsScalar field if non-nil, zero value otherwise.

### GetIsScalarOk

`func (o *XrefFromResponse) GetIsScalarOk() (*bool, bool)`

GetIsScalarOk returns a tuple with the IsScalar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScalar

`func (o *XrefFromResponse) SetIsScalar(v bool)`

SetIsScalar sets IsScalar field to given value.

### HasIsScalar

`func (o *XrefFromResponse) HasIsScalar() bool`

HasIsScalar returns a boolean if a field has been set.

### SetIsScalarNil

`func (o *XrefFromResponse) SetIsScalarNil(b bool)`

 SetIsScalarNil sets the value for IsScalar to be an explicit nil

### UnsetIsScalar
`func (o *XrefFromResponse) UnsetIsScalar()`

UnsetIsScalar ensures that no value is present for IsScalar, not even an explicit nil
### GetIsString

`func (o *XrefFromResponse) GetIsString() bool`

GetIsString returns the IsString field if non-nil, zero value otherwise.

### GetIsStringOk

`func (o *XrefFromResponse) GetIsStringOk() (*bool, bool)`

GetIsStringOk returns a tuple with the IsString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsString

`func (o *XrefFromResponse) SetIsString(v bool)`

SetIsString sets IsString field to given value.

### HasIsString

`func (o *XrefFromResponse) HasIsString() bool`

HasIsString returns a boolean if a field has been set.

### SetIsStringNil

`func (o *XrefFromResponse) SetIsStringNil(b bool)`

 SetIsStringNil sets the value for IsString to be an explicit nil

### UnsetIsString
`func (o *XrefFromResponse) UnsetIsString()`

UnsetIsString ensures that no value is present for IsString, not even an explicit nil
### GetOrigStrEncoding

`func (o *XrefFromResponse) GetOrigStrEncoding() string`

GetOrigStrEncoding returns the OrigStrEncoding field if non-nil, zero value otherwise.

### GetOrigStrEncodingOk

`func (o *XrefFromResponse) GetOrigStrEncodingOk() (*string, bool)`

GetOrigStrEncodingOk returns a tuple with the OrigStrEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigStrEncoding

`func (o *XrefFromResponse) SetOrigStrEncoding(v string)`

SetOrigStrEncoding sets OrigStrEncoding field to given value.

### HasOrigStrEncoding

`func (o *XrefFromResponse) HasOrigStrEncoding() bool`

HasOrigStrEncoding returns a boolean if a field has been set.

### SetOrigStrEncodingNil

`func (o *XrefFromResponse) SetOrigStrEncodingNil(b bool)`

 SetOrigStrEncodingNil sets the value for OrigStrEncoding to be an explicit nil

### UnsetOrigStrEncoding
`func (o *XrefFromResponse) UnsetOrigStrEncoding()`

UnsetOrigStrEncoding ensures that no value is present for OrigStrEncoding, not even an explicit nil
### GetRawData

`func (o *XrefFromResponse) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *XrefFromResponse) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *XrefFromResponse) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *XrefFromResponse) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### SetRawDataNil

`func (o *XrefFromResponse) SetRawDataNil(b bool)`

 SetRawDataNil sets the value for RawData to be an explicit nil

### UnsetRawData
`func (o *XrefFromResponse) UnsetRawData()`

UnsetRawData ensures that no value is present for RawData, not even an explicit nil
### GetSegment

`func (o *XrefFromResponse) GetSegment() SegmentInfo`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *XrefFromResponse) GetSegmentOk() (*SegmentInfo, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *XrefFromResponse) SetSegment(v SegmentInfo)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *XrefFromResponse) HasSegment() bool`

HasSegment returns a boolean if a field has been set.

### SetSegmentNil

`func (o *XrefFromResponse) SetSegmentNil(b bool)`

 SetSegmentNil sets the value for Segment to be an explicit nil

### UnsetSegment
`func (o *XrefFromResponse) UnsetSegment()`

UnsetSegment ensures that no value is present for Segment, not even an explicit nil
### GetValue

`func (o *XrefFromResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *XrefFromResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *XrefFromResponse) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *XrefFromResponse) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *XrefFromResponse) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetXrefTo

`func (o *XrefFromResponse) GetXrefTo() string`

GetXrefTo returns the XrefTo field if non-nil, zero value otherwise.

### GetXrefToOk

`func (o *XrefFromResponse) GetXrefToOk() (*string, bool)`

GetXrefToOk returns a tuple with the XrefTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXrefTo

`func (o *XrefFromResponse) SetXrefTo(v string)`

SetXrefTo sets XrefTo field to given value.


### SetXrefToNil

`func (o *XrefFromResponse) SetXrefToNil(b bool)`

 SetXrefToNil sets the value for XrefTo to be an explicit nil

### UnsetXrefTo
`func (o *XrefFromResponse) UnsetXrefTo()`

UnsetXrefTo ensures that no value is present for XrefTo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


