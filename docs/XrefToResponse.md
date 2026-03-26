# XrefToResponse

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
**XrefFrom** | **NullableString** |  | 

## Methods

### NewXrefToResponse

`func NewXrefToResponse(value NullableString, xrefFrom NullableString, ) *XrefToResponse`

NewXrefToResponse instantiates a new XrefToResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXrefToResponseWithDefaults

`func NewXrefToResponseWithDefaults() *XrefToResponse`

NewXrefToResponseWithDefaults instantiates a new XrefToResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsCall

`func (o *XrefToResponse) GetIsCall() bool`

GetIsCall returns the IsCall field if non-nil, zero value otherwise.

### GetIsCallOk

`func (o *XrefToResponse) GetIsCallOk() (*bool, bool)`

GetIsCallOk returns a tuple with the IsCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCall

`func (o *XrefToResponse) SetIsCall(v bool)`

SetIsCall sets IsCall field to given value.

### HasIsCall

`func (o *XrefToResponse) HasIsCall() bool`

HasIsCall returns a boolean if a field has been set.

### SetIsCallNil

`func (o *XrefToResponse) SetIsCallNil(b bool)`

 SetIsCallNil sets the value for IsCall to be an explicit nil

### UnsetIsCall
`func (o *XrefToResponse) UnsetIsCall()`

UnsetIsCall ensures that no value is present for IsCall, not even an explicit nil
### GetIsData

`func (o *XrefToResponse) GetIsData() bool`

GetIsData returns the IsData field if non-nil, zero value otherwise.

### GetIsDataOk

`func (o *XrefToResponse) GetIsDataOk() (*bool, bool)`

GetIsDataOk returns a tuple with the IsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsData

`func (o *XrefToResponse) SetIsData(v bool)`

SetIsData sets IsData field to given value.

### HasIsData

`func (o *XrefToResponse) HasIsData() bool`

HasIsData returns a boolean if a field has been set.

### SetIsDataNil

`func (o *XrefToResponse) SetIsDataNil(b bool)`

 SetIsDataNil sets the value for IsData to be an explicit nil

### UnsetIsData
`func (o *XrefToResponse) UnsetIsData()`

UnsetIsData ensures that no value is present for IsData, not even an explicit nil
### GetIsScalar

`func (o *XrefToResponse) GetIsScalar() bool`

GetIsScalar returns the IsScalar field if non-nil, zero value otherwise.

### GetIsScalarOk

`func (o *XrefToResponse) GetIsScalarOk() (*bool, bool)`

GetIsScalarOk returns a tuple with the IsScalar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScalar

`func (o *XrefToResponse) SetIsScalar(v bool)`

SetIsScalar sets IsScalar field to given value.

### HasIsScalar

`func (o *XrefToResponse) HasIsScalar() bool`

HasIsScalar returns a boolean if a field has been set.

### SetIsScalarNil

`func (o *XrefToResponse) SetIsScalarNil(b bool)`

 SetIsScalarNil sets the value for IsScalar to be an explicit nil

### UnsetIsScalar
`func (o *XrefToResponse) UnsetIsScalar()`

UnsetIsScalar ensures that no value is present for IsScalar, not even an explicit nil
### GetIsString

`func (o *XrefToResponse) GetIsString() bool`

GetIsString returns the IsString field if non-nil, zero value otherwise.

### GetIsStringOk

`func (o *XrefToResponse) GetIsStringOk() (*bool, bool)`

GetIsStringOk returns a tuple with the IsString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsString

`func (o *XrefToResponse) SetIsString(v bool)`

SetIsString sets IsString field to given value.

### HasIsString

`func (o *XrefToResponse) HasIsString() bool`

HasIsString returns a boolean if a field has been set.

### SetIsStringNil

`func (o *XrefToResponse) SetIsStringNil(b bool)`

 SetIsStringNil sets the value for IsString to be an explicit nil

### UnsetIsString
`func (o *XrefToResponse) UnsetIsString()`

UnsetIsString ensures that no value is present for IsString, not even an explicit nil
### GetOrigStrEncoding

`func (o *XrefToResponse) GetOrigStrEncoding() string`

GetOrigStrEncoding returns the OrigStrEncoding field if non-nil, zero value otherwise.

### GetOrigStrEncodingOk

`func (o *XrefToResponse) GetOrigStrEncodingOk() (*string, bool)`

GetOrigStrEncodingOk returns a tuple with the OrigStrEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigStrEncoding

`func (o *XrefToResponse) SetOrigStrEncoding(v string)`

SetOrigStrEncoding sets OrigStrEncoding field to given value.

### HasOrigStrEncoding

`func (o *XrefToResponse) HasOrigStrEncoding() bool`

HasOrigStrEncoding returns a boolean if a field has been set.

### SetOrigStrEncodingNil

`func (o *XrefToResponse) SetOrigStrEncodingNil(b bool)`

 SetOrigStrEncodingNil sets the value for OrigStrEncoding to be an explicit nil

### UnsetOrigStrEncoding
`func (o *XrefToResponse) UnsetOrigStrEncoding()`

UnsetOrigStrEncoding ensures that no value is present for OrigStrEncoding, not even an explicit nil
### GetRawData

`func (o *XrefToResponse) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *XrefToResponse) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *XrefToResponse) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *XrefToResponse) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### SetRawDataNil

`func (o *XrefToResponse) SetRawDataNil(b bool)`

 SetRawDataNil sets the value for RawData to be an explicit nil

### UnsetRawData
`func (o *XrefToResponse) UnsetRawData()`

UnsetRawData ensures that no value is present for RawData, not even an explicit nil
### GetSegment

`func (o *XrefToResponse) GetSegment() SegmentInfo`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *XrefToResponse) GetSegmentOk() (*SegmentInfo, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *XrefToResponse) SetSegment(v SegmentInfo)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *XrefToResponse) HasSegment() bool`

HasSegment returns a boolean if a field has been set.

### SetSegmentNil

`func (o *XrefToResponse) SetSegmentNil(b bool)`

 SetSegmentNil sets the value for Segment to be an explicit nil

### UnsetSegment
`func (o *XrefToResponse) UnsetSegment()`

UnsetSegment ensures that no value is present for Segment, not even an explicit nil
### GetValue

`func (o *XrefToResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *XrefToResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *XrefToResponse) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *XrefToResponse) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *XrefToResponse) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetXrefFrom

`func (o *XrefToResponse) GetXrefFrom() string`

GetXrefFrom returns the XrefFrom field if non-nil, zero value otherwise.

### GetXrefFromOk

`func (o *XrefToResponse) GetXrefFromOk() (*string, bool)`

GetXrefFromOk returns a tuple with the XrefFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXrefFrom

`func (o *XrefToResponse) SetXrefFrom(v string)`

SetXrefFrom sets XrefFrom field to given value.


### SetXrefFromNil

`func (o *XrefToResponse) SetXrefFromNil(b bool)`

 SetXrefFromNil sets the value for XrefFrom to be an explicit nil

### UnsetXrefFrom
`func (o *XrefToResponse) UnsetXrefFrom()`

UnsetXrefFrom ensures that no value is present for XrefFrom, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


