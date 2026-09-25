# XrefSegmentBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **int64** | End address of the segment, inclusive | 
**Exec** | **bool** | True when the segment is executable | 
**Kind** | **string** | Coarse classification of the segment | 
**Name** | **string** | Segment name | 
**Read** | **bool** | True when the segment is readable | 
**Start** | **int64** | Start address of the segment | 
**Write** | **bool** | True when the segment is writable | 

## Methods

### NewXrefSegmentBody

`func NewXrefSegmentBody(end int64, exec bool, kind string, name string, read bool, start int64, write bool, ) *XrefSegmentBody`

NewXrefSegmentBody instantiates a new XrefSegmentBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXrefSegmentBodyWithDefaults

`func NewXrefSegmentBodyWithDefaults() *XrefSegmentBody`

NewXrefSegmentBodyWithDefaults instantiates a new XrefSegmentBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *XrefSegmentBody) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *XrefSegmentBody) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *XrefSegmentBody) SetEnd(v int64)`

SetEnd sets End field to given value.


### GetExec

`func (o *XrefSegmentBody) GetExec() bool`

GetExec returns the Exec field if non-nil, zero value otherwise.

### GetExecOk

`func (o *XrefSegmentBody) GetExecOk() (*bool, bool)`

GetExecOk returns a tuple with the Exec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExec

`func (o *XrefSegmentBody) SetExec(v bool)`

SetExec sets Exec field to given value.


### GetKind

`func (o *XrefSegmentBody) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *XrefSegmentBody) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *XrefSegmentBody) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *XrefSegmentBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *XrefSegmentBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *XrefSegmentBody) SetName(v string)`

SetName sets Name field to given value.


### GetRead

`func (o *XrefSegmentBody) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *XrefSegmentBody) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *XrefSegmentBody) SetRead(v bool)`

SetRead sets Read field to given value.


### GetStart

`func (o *XrefSegmentBody) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *XrefSegmentBody) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *XrefSegmentBody) SetStart(v int64)`

SetStart sets Start field to given value.


### GetWrite

`func (o *XrefSegmentBody) GetWrite() bool`

GetWrite returns the Write field if non-nil, zero value otherwise.

### GetWriteOk

`func (o *XrefSegmentBody) GetWriteOk() (*bool, bool)`

GetWriteOk returns a tuple with the Write field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrite

`func (o *XrefSegmentBody) SetWrite(v bool)`

SetWrite sets Write field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


