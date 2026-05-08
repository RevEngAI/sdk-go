# MemdumpEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**DumpReason** | **string** |  | 
**FileType** | Pointer to **string** |  | [optional] 
**Filename** | **string** |  | 
**Index** | **int64** |  | 
**IsPe** | Pointer to **bool** |  | [optional] 
**Method** | **string** |  | 
**MimeType** | Pointer to **string** |  | [optional] 
**Sha256** | Pointer to **string** |  | [optional] 
**Size** | **int64** |  | 
**TargetAddr** | Pointer to **string** |  | [optional] 
**TargetProcess** | Pointer to **int64** |  | [optional] 

## Methods

### NewMemdumpEntry

`func NewMemdumpEntry(address string, dumpReason string, filename string, index int64, method string, size int64, ) *MemdumpEntry`

NewMemdumpEntry instantiates a new MemdumpEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemdumpEntryWithDefaults

`func NewMemdumpEntryWithDefaults() *MemdumpEntry`

NewMemdumpEntryWithDefaults instantiates a new MemdumpEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *MemdumpEntry) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MemdumpEntry) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MemdumpEntry) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetDumpReason

`func (o *MemdumpEntry) GetDumpReason() string`

GetDumpReason returns the DumpReason field if non-nil, zero value otherwise.

### GetDumpReasonOk

`func (o *MemdumpEntry) GetDumpReasonOk() (*string, bool)`

GetDumpReasonOk returns a tuple with the DumpReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDumpReason

`func (o *MemdumpEntry) SetDumpReason(v string)`

SetDumpReason sets DumpReason field to given value.


### GetFileType

`func (o *MemdumpEntry) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *MemdumpEntry) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *MemdumpEntry) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *MemdumpEntry) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetFilename

`func (o *MemdumpEntry) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *MemdumpEntry) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *MemdumpEntry) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetIndex

`func (o *MemdumpEntry) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *MemdumpEntry) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *MemdumpEntry) SetIndex(v int64)`

SetIndex sets Index field to given value.


### GetIsPe

`func (o *MemdumpEntry) GetIsPe() bool`

GetIsPe returns the IsPe field if non-nil, zero value otherwise.

### GetIsPeOk

`func (o *MemdumpEntry) GetIsPeOk() (*bool, bool)`

GetIsPeOk returns a tuple with the IsPe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPe

`func (o *MemdumpEntry) SetIsPe(v bool)`

SetIsPe sets IsPe field to given value.

### HasIsPe

`func (o *MemdumpEntry) HasIsPe() bool`

HasIsPe returns a boolean if a field has been set.

### GetMethod

`func (o *MemdumpEntry) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *MemdumpEntry) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *MemdumpEntry) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetMimeType

`func (o *MemdumpEntry) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *MemdumpEntry) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *MemdumpEntry) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *MemdumpEntry) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetSha256

`func (o *MemdumpEntry) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *MemdumpEntry) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *MemdumpEntry) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *MemdumpEntry) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetSize

`func (o *MemdumpEntry) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MemdumpEntry) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MemdumpEntry) SetSize(v int64)`

SetSize sets Size field to given value.


### GetTargetAddr

`func (o *MemdumpEntry) GetTargetAddr() string`

GetTargetAddr returns the TargetAddr field if non-nil, zero value otherwise.

### GetTargetAddrOk

`func (o *MemdumpEntry) GetTargetAddrOk() (*string, bool)`

GetTargetAddrOk returns a tuple with the TargetAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAddr

`func (o *MemdumpEntry) SetTargetAddr(v string)`

SetTargetAddr sets TargetAddr field to given value.

### HasTargetAddr

`func (o *MemdumpEntry) HasTargetAddr() bool`

HasTargetAddr returns a boolean if a field has been set.

### GetTargetProcess

`func (o *MemdumpEntry) GetTargetProcess() int64`

GetTargetProcess returns the TargetProcess field if non-nil, zero value otherwise.

### GetTargetProcessOk

`func (o *MemdumpEntry) GetTargetProcessOk() (*int64, bool)`

GetTargetProcessOk returns a tuple with the TargetProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetProcess

`func (o *MemdumpEntry) SetTargetProcess(v int64)`

SetTargetProcess sets TargetProcess field to given value.

### HasTargetProcess

`func (o *MemdumpEntry) HasTargetProcess() bool`

HasTargetProcess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


