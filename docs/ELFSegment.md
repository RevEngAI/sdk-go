# ELFSegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**VirtualAddress** | **int32** |  | 
**VirtualSize** | **int32** |  | 
**PhysicalAddress** | **int32** |  | 
**PhysicalSize** | **int32** |  | 
**FileOffset** | **int32** |  | 
**Flags** | **string** |  | 
**FlagsRaw** | **int32** |  | 
**Alignment** | **int32** |  | 

## Methods

### NewELFSegment

`func NewELFSegment(type_ string, virtualAddress int32, virtualSize int32, physicalAddress int32, physicalSize int32, fileOffset int32, flags string, flagsRaw int32, alignment int32, ) *ELFSegment`

NewELFSegment instantiates a new ELFSegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewELFSegmentWithDefaults

`func NewELFSegmentWithDefaults() *ELFSegment`

NewELFSegmentWithDefaults instantiates a new ELFSegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ELFSegment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ELFSegment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ELFSegment) SetType(v string)`

SetType sets Type field to given value.


### GetVirtualAddress

`func (o *ELFSegment) GetVirtualAddress() int32`

GetVirtualAddress returns the VirtualAddress field if non-nil, zero value otherwise.

### GetVirtualAddressOk

`func (o *ELFSegment) GetVirtualAddressOk() (*int32, bool)`

GetVirtualAddressOk returns a tuple with the VirtualAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAddress

`func (o *ELFSegment) SetVirtualAddress(v int32)`

SetVirtualAddress sets VirtualAddress field to given value.


### GetVirtualSize

`func (o *ELFSegment) GetVirtualSize() int32`

GetVirtualSize returns the VirtualSize field if non-nil, zero value otherwise.

### GetVirtualSizeOk

`func (o *ELFSegment) GetVirtualSizeOk() (*int32, bool)`

GetVirtualSizeOk returns a tuple with the VirtualSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualSize

`func (o *ELFSegment) SetVirtualSize(v int32)`

SetVirtualSize sets VirtualSize field to given value.


### GetPhysicalAddress

`func (o *ELFSegment) GetPhysicalAddress() int32`

GetPhysicalAddress returns the PhysicalAddress field if non-nil, zero value otherwise.

### GetPhysicalAddressOk

`func (o *ELFSegment) GetPhysicalAddressOk() (*int32, bool)`

GetPhysicalAddressOk returns a tuple with the PhysicalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalAddress

`func (o *ELFSegment) SetPhysicalAddress(v int32)`

SetPhysicalAddress sets PhysicalAddress field to given value.


### GetPhysicalSize

`func (o *ELFSegment) GetPhysicalSize() int32`

GetPhysicalSize returns the PhysicalSize field if non-nil, zero value otherwise.

### GetPhysicalSizeOk

`func (o *ELFSegment) GetPhysicalSizeOk() (*int32, bool)`

GetPhysicalSizeOk returns a tuple with the PhysicalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalSize

`func (o *ELFSegment) SetPhysicalSize(v int32)`

SetPhysicalSize sets PhysicalSize field to given value.


### GetFileOffset

`func (o *ELFSegment) GetFileOffset() int32`

GetFileOffset returns the FileOffset field if non-nil, zero value otherwise.

### GetFileOffsetOk

`func (o *ELFSegment) GetFileOffsetOk() (*int32, bool)`

GetFileOffsetOk returns a tuple with the FileOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileOffset

`func (o *ELFSegment) SetFileOffset(v int32)`

SetFileOffset sets FileOffset field to given value.


### GetFlags

`func (o *ELFSegment) GetFlags() string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *ELFSegment) GetFlagsOk() (*string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *ELFSegment) SetFlags(v string)`

SetFlags sets Flags field to given value.


### GetFlagsRaw

`func (o *ELFSegment) GetFlagsRaw() int32`

GetFlagsRaw returns the FlagsRaw field if non-nil, zero value otherwise.

### GetFlagsRawOk

`func (o *ELFSegment) GetFlagsRawOk() (*int32, bool)`

GetFlagsRawOk returns a tuple with the FlagsRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagsRaw

`func (o *ELFSegment) SetFlagsRaw(v int32)`

SetFlagsRaw sets FlagsRaw field to given value.


### GetAlignment

`func (o *ELFSegment) GetAlignment() int32`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *ELFSegment) GetAlignmentOk() (*int32, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *ELFSegment) SetAlignment(v int32)`

SetAlignment sets Alignment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


