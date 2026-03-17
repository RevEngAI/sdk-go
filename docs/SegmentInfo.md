# SegmentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] [default to ""]
**R** | Pointer to **NullableBool** |  | [optional] 
**W** | Pointer to **NullableBool** |  | [optional] 
**X** | Pointer to **NullableBool** |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] [default to 0]
**End** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewSegmentInfo

`func NewSegmentInfo() *SegmentInfo`

NewSegmentInfo instantiates a new SegmentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentInfoWithDefaults

`func NewSegmentInfoWithDefaults() *SegmentInfo`

NewSegmentInfoWithDefaults instantiates a new SegmentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SegmentInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SegmentInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SegmentInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SegmentInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetR

`func (o *SegmentInfo) GetR() bool`

GetR returns the R field if non-nil, zero value otherwise.

### GetROk

`func (o *SegmentInfo) GetROk() (*bool, bool)`

GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR

`func (o *SegmentInfo) SetR(v bool)`

SetR sets R field to given value.

### HasR

`func (o *SegmentInfo) HasR() bool`

HasR returns a boolean if a field has been set.

### SetRNil

`func (o *SegmentInfo) SetRNil(b bool)`

 SetRNil sets the value for R to be an explicit nil

### UnsetR
`func (o *SegmentInfo) UnsetR()`

UnsetR ensures that no value is present for R, not even an explicit nil
### GetW

`func (o *SegmentInfo) GetW() bool`

GetW returns the W field if non-nil, zero value otherwise.

### GetWOk

`func (o *SegmentInfo) GetWOk() (*bool, bool)`

GetWOk returns a tuple with the W field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetW

`func (o *SegmentInfo) SetW(v bool)`

SetW sets W field to given value.

### HasW

`func (o *SegmentInfo) HasW() bool`

HasW returns a boolean if a field has been set.

### SetWNil

`func (o *SegmentInfo) SetWNil(b bool)`

 SetWNil sets the value for W to be an explicit nil

### UnsetW
`func (o *SegmentInfo) UnsetW()`

UnsetW ensures that no value is present for W, not even an explicit nil
### GetX

`func (o *SegmentInfo) GetX() bool`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *SegmentInfo) GetXOk() (*bool, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *SegmentInfo) SetX(v bool)`

SetX sets X field to given value.

### HasX

`func (o *SegmentInfo) HasX() bool`

HasX returns a boolean if a field has been set.

### SetXNil

`func (o *SegmentInfo) SetXNil(b bool)`

 SetXNil sets the value for X to be an explicit nil

### UnsetX
`func (o *SegmentInfo) UnsetX()`

UnsetX ensures that no value is present for X, not even an explicit nil
### GetStart

`func (o *SegmentInfo) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SegmentInfo) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SegmentInfo) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *SegmentInfo) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *SegmentInfo) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SegmentInfo) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SegmentInfo) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *SegmentInfo) HasEnd() bool`

HasEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


