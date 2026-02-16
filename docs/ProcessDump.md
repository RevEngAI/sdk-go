# ProcessDump

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseAddress** | **int32** |  | 
**ActualFilename** | **string** |  | 
**FilenameFriendly** | **string** |  | 
**ExtendedMetadata** | [**ProcessDumpMetadata**](ProcessDumpMetadata.md) |  | 

## Methods

### NewProcessDump

`func NewProcessDump(baseAddress int32, actualFilename string, filenameFriendly string, extendedMetadata ProcessDumpMetadata, ) *ProcessDump`

NewProcessDump instantiates a new ProcessDump object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessDumpWithDefaults

`func NewProcessDumpWithDefaults() *ProcessDump`

NewProcessDumpWithDefaults instantiates a new ProcessDump object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseAddress

`func (o *ProcessDump) GetBaseAddress() int32`

GetBaseAddress returns the BaseAddress field if non-nil, zero value otherwise.

### GetBaseAddressOk

`func (o *ProcessDump) GetBaseAddressOk() (*int32, bool)`

GetBaseAddressOk returns a tuple with the BaseAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAddress

`func (o *ProcessDump) SetBaseAddress(v int32)`

SetBaseAddress sets BaseAddress field to given value.


### GetActualFilename

`func (o *ProcessDump) GetActualFilename() string`

GetActualFilename returns the ActualFilename field if non-nil, zero value otherwise.

### GetActualFilenameOk

`func (o *ProcessDump) GetActualFilenameOk() (*string, bool)`

GetActualFilenameOk returns a tuple with the ActualFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualFilename

`func (o *ProcessDump) SetActualFilename(v string)`

SetActualFilename sets ActualFilename field to given value.


### GetFilenameFriendly

`func (o *ProcessDump) GetFilenameFriendly() string`

GetFilenameFriendly returns the FilenameFriendly field if non-nil, zero value otherwise.

### GetFilenameFriendlyOk

`func (o *ProcessDump) GetFilenameFriendlyOk() (*string, bool)`

GetFilenameFriendlyOk returns a tuple with the FilenameFriendly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilenameFriendly

`func (o *ProcessDump) SetFilenameFriendly(v string)`

SetFilenameFriendly sets FilenameFriendly field to given value.


### GetExtendedMetadata

`func (o *ProcessDump) GetExtendedMetadata() ProcessDumpMetadata`

GetExtendedMetadata returns the ExtendedMetadata field if non-nil, zero value otherwise.

### GetExtendedMetadataOk

`func (o *ProcessDump) GetExtendedMetadataOk() (*ProcessDumpMetadata, bool)`

GetExtendedMetadataOk returns a tuple with the ExtendedMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedMetadata

`func (o *ProcessDump) SetExtendedMetadata(v ProcessDumpMetadata)`

SetExtendedMetadata sets ExtendedMetadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


