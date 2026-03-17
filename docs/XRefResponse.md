# XRefResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **NullableString** |  | 
**XrefTo** | **NullableString** |  | 
**IsScalar** | Pointer to **NullableBool** |  | [optional] 
**IsCall** | Pointer to **NullableBool** |  | [optional] 
**IsData** | Pointer to **NullableBool** |  | [optional] 
**IsString** | Pointer to **NullableBool** |  | [optional] 
**RawData** | Pointer to **NullableString** |  | [optional] 
**Segment** | Pointer to [**NullableSegmentInfo**](SegmentInfo.md) |  | [optional] 
**OrigStrEncoding** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewXRefResponse

`func NewXRefResponse(value NullableString, xrefTo NullableString, ) *XRefResponse`

NewXRefResponse instantiates a new XRefResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXRefResponseWithDefaults

`func NewXRefResponseWithDefaults() *XRefResponse`

NewXRefResponseWithDefaults instantiates a new XRefResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *XRefResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *XRefResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *XRefResponse) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *XRefResponse) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *XRefResponse) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetXrefTo

`func (o *XRefResponse) GetXrefTo() string`

GetXrefTo returns the XrefTo field if non-nil, zero value otherwise.

### GetXrefToOk

`func (o *XRefResponse) GetXrefToOk() (*string, bool)`

GetXrefToOk returns a tuple with the XrefTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXrefTo

`func (o *XRefResponse) SetXrefTo(v string)`

SetXrefTo sets XrefTo field to given value.


### SetXrefToNil

`func (o *XRefResponse) SetXrefToNil(b bool)`

 SetXrefToNil sets the value for XrefTo to be an explicit nil

### UnsetXrefTo
`func (o *XRefResponse) UnsetXrefTo()`

UnsetXrefTo ensures that no value is present for XrefTo, not even an explicit nil
### GetIsScalar

`func (o *XRefResponse) GetIsScalar() bool`

GetIsScalar returns the IsScalar field if non-nil, zero value otherwise.

### GetIsScalarOk

`func (o *XRefResponse) GetIsScalarOk() (*bool, bool)`

GetIsScalarOk returns a tuple with the IsScalar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScalar

`func (o *XRefResponse) SetIsScalar(v bool)`

SetIsScalar sets IsScalar field to given value.

### HasIsScalar

`func (o *XRefResponse) HasIsScalar() bool`

HasIsScalar returns a boolean if a field has been set.

### SetIsScalarNil

`func (o *XRefResponse) SetIsScalarNil(b bool)`

 SetIsScalarNil sets the value for IsScalar to be an explicit nil

### UnsetIsScalar
`func (o *XRefResponse) UnsetIsScalar()`

UnsetIsScalar ensures that no value is present for IsScalar, not even an explicit nil
### GetIsCall

`func (o *XRefResponse) GetIsCall() bool`

GetIsCall returns the IsCall field if non-nil, zero value otherwise.

### GetIsCallOk

`func (o *XRefResponse) GetIsCallOk() (*bool, bool)`

GetIsCallOk returns a tuple with the IsCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCall

`func (o *XRefResponse) SetIsCall(v bool)`

SetIsCall sets IsCall field to given value.

### HasIsCall

`func (o *XRefResponse) HasIsCall() bool`

HasIsCall returns a boolean if a field has been set.

### SetIsCallNil

`func (o *XRefResponse) SetIsCallNil(b bool)`

 SetIsCallNil sets the value for IsCall to be an explicit nil

### UnsetIsCall
`func (o *XRefResponse) UnsetIsCall()`

UnsetIsCall ensures that no value is present for IsCall, not even an explicit nil
### GetIsData

`func (o *XRefResponse) GetIsData() bool`

GetIsData returns the IsData field if non-nil, zero value otherwise.

### GetIsDataOk

`func (o *XRefResponse) GetIsDataOk() (*bool, bool)`

GetIsDataOk returns a tuple with the IsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsData

`func (o *XRefResponse) SetIsData(v bool)`

SetIsData sets IsData field to given value.

### HasIsData

`func (o *XRefResponse) HasIsData() bool`

HasIsData returns a boolean if a field has been set.

### SetIsDataNil

`func (o *XRefResponse) SetIsDataNil(b bool)`

 SetIsDataNil sets the value for IsData to be an explicit nil

### UnsetIsData
`func (o *XRefResponse) UnsetIsData()`

UnsetIsData ensures that no value is present for IsData, not even an explicit nil
### GetIsString

`func (o *XRefResponse) GetIsString() bool`

GetIsString returns the IsString field if non-nil, zero value otherwise.

### GetIsStringOk

`func (o *XRefResponse) GetIsStringOk() (*bool, bool)`

GetIsStringOk returns a tuple with the IsString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsString

`func (o *XRefResponse) SetIsString(v bool)`

SetIsString sets IsString field to given value.

### HasIsString

`func (o *XRefResponse) HasIsString() bool`

HasIsString returns a boolean if a field has been set.

### SetIsStringNil

`func (o *XRefResponse) SetIsStringNil(b bool)`

 SetIsStringNil sets the value for IsString to be an explicit nil

### UnsetIsString
`func (o *XRefResponse) UnsetIsString()`

UnsetIsString ensures that no value is present for IsString, not even an explicit nil
### GetRawData

`func (o *XRefResponse) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *XRefResponse) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *XRefResponse) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *XRefResponse) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### SetRawDataNil

`func (o *XRefResponse) SetRawDataNil(b bool)`

 SetRawDataNil sets the value for RawData to be an explicit nil

### UnsetRawData
`func (o *XRefResponse) UnsetRawData()`

UnsetRawData ensures that no value is present for RawData, not even an explicit nil
### GetSegment

`func (o *XRefResponse) GetSegment() SegmentInfo`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *XRefResponse) GetSegmentOk() (*SegmentInfo, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *XRefResponse) SetSegment(v SegmentInfo)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *XRefResponse) HasSegment() bool`

HasSegment returns a boolean if a field has been set.

### SetSegmentNil

`func (o *XRefResponse) SetSegmentNil(b bool)`

 SetSegmentNil sets the value for Segment to be an explicit nil

### UnsetSegment
`func (o *XRefResponse) UnsetSegment()`

UnsetSegment ensures that no value is present for Segment, not even an explicit nil
### GetOrigStrEncoding

`func (o *XRefResponse) GetOrigStrEncoding() string`

GetOrigStrEncoding returns the OrigStrEncoding field if non-nil, zero value otherwise.

### GetOrigStrEncodingOk

`func (o *XRefResponse) GetOrigStrEncodingOk() (*string, bool)`

GetOrigStrEncodingOk returns a tuple with the OrigStrEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigStrEncoding

`func (o *XRefResponse) SetOrigStrEncoding(v string)`

SetOrigStrEncoding sets OrigStrEncoding field to given value.

### HasOrigStrEncoding

`func (o *XRefResponse) HasOrigStrEncoding() bool`

HasOrigStrEncoding returns a boolean if a field has been set.

### SetOrigStrEncodingNil

`func (o *XRefResponse) SetOrigStrEncodingNil(b bool)`

 SetOrigStrEncodingNil sets the value for OrigStrEncoding to be an explicit nil

### UnsetOrigStrEncoding
`func (o *XRefResponse) UnsetOrigStrEncoding()`

UnsetOrigStrEncoding ensures that no value is present for OrigStrEncoding, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


