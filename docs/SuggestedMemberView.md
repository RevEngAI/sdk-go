# SuggestedMemberView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BitOffset** | **NullableInt64** | Bit offset within the containing word for a bitfield member. | 
**ByteOffset** | **NullableInt64** | Offset within the type. Null when no placement could be established. | 
**ByteSize** | **NullableInt64** | Width of the access in bytes. | 
**Confidence** | **NullableString** | Where suggested_type came from. declared and inferred were observed; width knows only the access width; model is the language model&#39;s proposal. | 
**Name** | **string** | Name the member renders as: a database or frozen name where one exists, else the suggested one. | 
**Origin** | **string** | Which function revealed this member: self, caller:&lt;function_id&gt; or callee:&lt;function_id&gt;. | 
**Packed** | **bool** | The member sits at an offset its own width does not divide. | 
**Placement** | **string** | observed means the offset was read off an access; guessed means the model proposed it; unplaced means the member has no offset. | 
**SuggestedType** | **NullableString** | Type expression for the member. | 
**Token** | Pointer to **string** | Placeholder this member renders as in the tokenised source. Absent for a member no access in this function revealed. | [optional] 

## Methods

### NewSuggestedMemberView

`func NewSuggestedMemberView(bitOffset NullableInt64, byteOffset NullableInt64, byteSize NullableInt64, confidence NullableString, name string, origin string, packed bool, placement string, suggestedType NullableString, ) *SuggestedMemberView`

NewSuggestedMemberView instantiates a new SuggestedMemberView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestedMemberViewWithDefaults

`func NewSuggestedMemberViewWithDefaults() *SuggestedMemberView`

NewSuggestedMemberViewWithDefaults instantiates a new SuggestedMemberView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBitOffset

`func (o *SuggestedMemberView) GetBitOffset() int64`

GetBitOffset returns the BitOffset field if non-nil, zero value otherwise.

### GetBitOffsetOk

`func (o *SuggestedMemberView) GetBitOffsetOk() (*int64, bool)`

GetBitOffsetOk returns a tuple with the BitOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitOffset

`func (o *SuggestedMemberView) SetBitOffset(v int64)`

SetBitOffset sets BitOffset field to given value.


### SetBitOffsetNil

`func (o *SuggestedMemberView) SetBitOffsetNil(b bool)`

 SetBitOffsetNil sets the value for BitOffset to be an explicit nil

### UnsetBitOffset
`func (o *SuggestedMemberView) UnsetBitOffset()`

UnsetBitOffset ensures that no value is present for BitOffset, not even an explicit nil
### GetByteOffset

`func (o *SuggestedMemberView) GetByteOffset() int64`

GetByteOffset returns the ByteOffset field if non-nil, zero value otherwise.

### GetByteOffsetOk

`func (o *SuggestedMemberView) GetByteOffsetOk() (*int64, bool)`

GetByteOffsetOk returns a tuple with the ByteOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteOffset

`func (o *SuggestedMemberView) SetByteOffset(v int64)`

SetByteOffset sets ByteOffset field to given value.


### SetByteOffsetNil

`func (o *SuggestedMemberView) SetByteOffsetNil(b bool)`

 SetByteOffsetNil sets the value for ByteOffset to be an explicit nil

### UnsetByteOffset
`func (o *SuggestedMemberView) UnsetByteOffset()`

UnsetByteOffset ensures that no value is present for ByteOffset, not even an explicit nil
### GetByteSize

`func (o *SuggestedMemberView) GetByteSize() int64`

GetByteSize returns the ByteSize field if non-nil, zero value otherwise.

### GetByteSizeOk

`func (o *SuggestedMemberView) GetByteSizeOk() (*int64, bool)`

GetByteSizeOk returns a tuple with the ByteSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteSize

`func (o *SuggestedMemberView) SetByteSize(v int64)`

SetByteSize sets ByteSize field to given value.


### SetByteSizeNil

`func (o *SuggestedMemberView) SetByteSizeNil(b bool)`

 SetByteSizeNil sets the value for ByteSize to be an explicit nil

### UnsetByteSize
`func (o *SuggestedMemberView) UnsetByteSize()`

UnsetByteSize ensures that no value is present for ByteSize, not even an explicit nil
### GetConfidence

`func (o *SuggestedMemberView) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *SuggestedMemberView) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *SuggestedMemberView) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.


### SetConfidenceNil

`func (o *SuggestedMemberView) SetConfidenceNil(b bool)`

 SetConfidenceNil sets the value for Confidence to be an explicit nil

### UnsetConfidence
`func (o *SuggestedMemberView) UnsetConfidence()`

UnsetConfidence ensures that no value is present for Confidence, not even an explicit nil
### GetName

`func (o *SuggestedMemberView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SuggestedMemberView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SuggestedMemberView) SetName(v string)`

SetName sets Name field to given value.


### GetOrigin

`func (o *SuggestedMemberView) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *SuggestedMemberView) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *SuggestedMemberView) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetPacked

`func (o *SuggestedMemberView) GetPacked() bool`

GetPacked returns the Packed field if non-nil, zero value otherwise.

### GetPackedOk

`func (o *SuggestedMemberView) GetPackedOk() (*bool, bool)`

GetPackedOk returns a tuple with the Packed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacked

`func (o *SuggestedMemberView) SetPacked(v bool)`

SetPacked sets Packed field to given value.


### GetPlacement

`func (o *SuggestedMemberView) GetPlacement() string`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *SuggestedMemberView) GetPlacementOk() (*string, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *SuggestedMemberView) SetPlacement(v string)`

SetPlacement sets Placement field to given value.


### GetSuggestedType

`func (o *SuggestedMemberView) GetSuggestedType() string`

GetSuggestedType returns the SuggestedType field if non-nil, zero value otherwise.

### GetSuggestedTypeOk

`func (o *SuggestedMemberView) GetSuggestedTypeOk() (*string, bool)`

GetSuggestedTypeOk returns a tuple with the SuggestedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedType

`func (o *SuggestedMemberView) SetSuggestedType(v string)`

SetSuggestedType sets SuggestedType field to given value.


### SetSuggestedTypeNil

`func (o *SuggestedMemberView) SetSuggestedTypeNil(b bool)`

 SetSuggestedTypeNil sets the value for SuggestedType to be an explicit nil

### UnsetSuggestedType
`func (o *SuggestedMemberView) UnsetSuggestedType()`

UnsetSuggestedType ensures that no value is present for SuggestedType, not even an explicit nil
### GetToken

`func (o *SuggestedMemberView) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SuggestedMemberView) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SuggestedMemberView) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SuggestedMemberView) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


