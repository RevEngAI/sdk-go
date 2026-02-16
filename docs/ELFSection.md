# ELFSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**VirtualAddress** | **int32** |  | 
**VirtualSize** | **int32** |  | 
**RawSize** | **int32** |  | 
**FileOffset** | **int32** |  | 
**Flags** | **string** |  | 
**FlagsRaw** | **int32** |  | 
**Entropy** | **float32** |  | 
**Alignment** | **int32** |  | 

## Methods

### NewELFSection

`func NewELFSection(name string, type_ string, virtualAddress int32, virtualSize int32, rawSize int32, fileOffset int32, flags string, flagsRaw int32, entropy float32, alignment int32, ) *ELFSection`

NewELFSection instantiates a new ELFSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewELFSectionWithDefaults

`func NewELFSectionWithDefaults() *ELFSection`

NewELFSectionWithDefaults instantiates a new ELFSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ELFSection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ELFSection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ELFSection) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ELFSection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ELFSection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ELFSection) SetType(v string)`

SetType sets Type field to given value.


### GetVirtualAddress

`func (o *ELFSection) GetVirtualAddress() int32`

GetVirtualAddress returns the VirtualAddress field if non-nil, zero value otherwise.

### GetVirtualAddressOk

`func (o *ELFSection) GetVirtualAddressOk() (*int32, bool)`

GetVirtualAddressOk returns a tuple with the VirtualAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAddress

`func (o *ELFSection) SetVirtualAddress(v int32)`

SetVirtualAddress sets VirtualAddress field to given value.


### GetVirtualSize

`func (o *ELFSection) GetVirtualSize() int32`

GetVirtualSize returns the VirtualSize field if non-nil, zero value otherwise.

### GetVirtualSizeOk

`func (o *ELFSection) GetVirtualSizeOk() (*int32, bool)`

GetVirtualSizeOk returns a tuple with the VirtualSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualSize

`func (o *ELFSection) SetVirtualSize(v int32)`

SetVirtualSize sets VirtualSize field to given value.


### GetRawSize

`func (o *ELFSection) GetRawSize() int32`

GetRawSize returns the RawSize field if non-nil, zero value otherwise.

### GetRawSizeOk

`func (o *ELFSection) GetRawSizeOk() (*int32, bool)`

GetRawSizeOk returns a tuple with the RawSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSize

`func (o *ELFSection) SetRawSize(v int32)`

SetRawSize sets RawSize field to given value.


### GetFileOffset

`func (o *ELFSection) GetFileOffset() int32`

GetFileOffset returns the FileOffset field if non-nil, zero value otherwise.

### GetFileOffsetOk

`func (o *ELFSection) GetFileOffsetOk() (*int32, bool)`

GetFileOffsetOk returns a tuple with the FileOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileOffset

`func (o *ELFSection) SetFileOffset(v int32)`

SetFileOffset sets FileOffset field to given value.


### GetFlags

`func (o *ELFSection) GetFlags() string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *ELFSection) GetFlagsOk() (*string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *ELFSection) SetFlags(v string)`

SetFlags sets Flags field to given value.


### GetFlagsRaw

`func (o *ELFSection) GetFlagsRaw() int32`

GetFlagsRaw returns the FlagsRaw field if non-nil, zero value otherwise.

### GetFlagsRawOk

`func (o *ELFSection) GetFlagsRawOk() (*int32, bool)`

GetFlagsRawOk returns a tuple with the FlagsRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagsRaw

`func (o *ELFSection) SetFlagsRaw(v int32)`

SetFlagsRaw sets FlagsRaw field to given value.


### GetEntropy

`func (o *ELFSection) GetEntropy() float32`

GetEntropy returns the Entropy field if non-nil, zero value otherwise.

### GetEntropyOk

`func (o *ELFSection) GetEntropyOk() (*float32, bool)`

GetEntropyOk returns a tuple with the Entropy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntropy

`func (o *ELFSection) SetEntropy(v float32)`

SetEntropy sets Entropy field to given value.


### GetAlignment

`func (o *ELFSection) GetAlignment() int32`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *ELFSection) GetAlignmentOk() (*int32, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *ELFSection) SetAlignment(v int32)`

SetAlignment sets Alignment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


