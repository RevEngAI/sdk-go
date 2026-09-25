# XrefIntoBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsCall** | **bool** | True when the xref is a call instruction | 
**IsData** | **bool** | True when the xref targets data rather than code | 
**IsScalar** | **bool** | True when the xref is a scalar constant | 
**IsString** | **bool** | True when the xref targets a string | 
**OrigStrEncoding** | Pointer to **string** | String encoding, set only when is_string is true | [optional] 
**RawData** | Pointer to **string** | Raw bytes at the xref target, when captured | [optional] 
**Segment** | Pointer to [**XrefSegmentBody**](XrefSegmentBody.md) | Memory segment the xref target sits in | [optional] 
**Value** | Pointer to **string** | The xref&#39;s resolved value, when the sequencer could determine one | [optional] 
**XrefFrom** | **int64** | Address the reference originates from | 

## Methods

### NewXrefIntoBody

`func NewXrefIntoBody(isCall bool, isData bool, isScalar bool, isString bool, xrefFrom int64, ) *XrefIntoBody`

NewXrefIntoBody instantiates a new XrefIntoBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXrefIntoBodyWithDefaults

`func NewXrefIntoBodyWithDefaults() *XrefIntoBody`

NewXrefIntoBodyWithDefaults instantiates a new XrefIntoBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsCall

`func (o *XrefIntoBody) GetIsCall() bool`

GetIsCall returns the IsCall field if non-nil, zero value otherwise.

### GetIsCallOk

`func (o *XrefIntoBody) GetIsCallOk() (*bool, bool)`

GetIsCallOk returns a tuple with the IsCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCall

`func (o *XrefIntoBody) SetIsCall(v bool)`

SetIsCall sets IsCall field to given value.


### GetIsData

`func (o *XrefIntoBody) GetIsData() bool`

GetIsData returns the IsData field if non-nil, zero value otherwise.

### GetIsDataOk

`func (o *XrefIntoBody) GetIsDataOk() (*bool, bool)`

GetIsDataOk returns a tuple with the IsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsData

`func (o *XrefIntoBody) SetIsData(v bool)`

SetIsData sets IsData field to given value.


### GetIsScalar

`func (o *XrefIntoBody) GetIsScalar() bool`

GetIsScalar returns the IsScalar field if non-nil, zero value otherwise.

### GetIsScalarOk

`func (o *XrefIntoBody) GetIsScalarOk() (*bool, bool)`

GetIsScalarOk returns a tuple with the IsScalar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScalar

`func (o *XrefIntoBody) SetIsScalar(v bool)`

SetIsScalar sets IsScalar field to given value.


### GetIsString

`func (o *XrefIntoBody) GetIsString() bool`

GetIsString returns the IsString field if non-nil, zero value otherwise.

### GetIsStringOk

`func (o *XrefIntoBody) GetIsStringOk() (*bool, bool)`

GetIsStringOk returns a tuple with the IsString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsString

`func (o *XrefIntoBody) SetIsString(v bool)`

SetIsString sets IsString field to given value.


### GetOrigStrEncoding

`func (o *XrefIntoBody) GetOrigStrEncoding() string`

GetOrigStrEncoding returns the OrigStrEncoding field if non-nil, zero value otherwise.

### GetOrigStrEncodingOk

`func (o *XrefIntoBody) GetOrigStrEncodingOk() (*string, bool)`

GetOrigStrEncodingOk returns a tuple with the OrigStrEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigStrEncoding

`func (o *XrefIntoBody) SetOrigStrEncoding(v string)`

SetOrigStrEncoding sets OrigStrEncoding field to given value.

### HasOrigStrEncoding

`func (o *XrefIntoBody) HasOrigStrEncoding() bool`

HasOrigStrEncoding returns a boolean if a field has been set.

### GetRawData

`func (o *XrefIntoBody) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *XrefIntoBody) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *XrefIntoBody) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *XrefIntoBody) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### GetSegment

`func (o *XrefIntoBody) GetSegment() XrefSegmentBody`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *XrefIntoBody) GetSegmentOk() (*XrefSegmentBody, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *XrefIntoBody) SetSegment(v XrefSegmentBody)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *XrefIntoBody) HasSegment() bool`

HasSegment returns a boolean if a field has been set.

### GetValue

`func (o *XrefIntoBody) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *XrefIntoBody) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *XrefIntoBody) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *XrefIntoBody) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetXrefFrom

`func (o *XrefIntoBody) GetXrefFrom() int64`

GetXrefFrom returns the XrefFrom field if non-nil, zero value otherwise.

### GetXrefFromOk

`func (o *XrefIntoBody) GetXrefFromOk() (*int64, bool)`

GetXrefFromOk returns a tuple with the XrefFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXrefFrom

`func (o *XrefIntoBody) SetXrefFrom(v int64)`

SetXrefFrom sets XrefFrom field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


