# ExtractedFileEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileHash** | Pointer to **string** |  | [optional] 
**FileSize** | **int64** |  | 
**FileType** | Pointer to **string** |  | [optional] 
**Filename** | **string** |  | 
**IsPe** | Pointer to **bool** |  | [optional] 
**MimeType** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**SeqNum** | **int64** |  | 
**Sha256** | Pointer to **string** |  | [optional] 
**ZipFilename** | **string** |  | 

## Methods

### NewExtractedFileEntry

`func NewExtractedFileEntry(fileSize int64, filename string, seqNum int64, zipFilename string, ) *ExtractedFileEntry`

NewExtractedFileEntry instantiates a new ExtractedFileEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtractedFileEntryWithDefaults

`func NewExtractedFileEntryWithDefaults() *ExtractedFileEntry`

NewExtractedFileEntryWithDefaults instantiates a new ExtractedFileEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileHash

`func (o *ExtractedFileEntry) GetFileHash() string`

GetFileHash returns the FileHash field if non-nil, zero value otherwise.

### GetFileHashOk

`func (o *ExtractedFileEntry) GetFileHashOk() (*string, bool)`

GetFileHashOk returns a tuple with the FileHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileHash

`func (o *ExtractedFileEntry) SetFileHash(v string)`

SetFileHash sets FileHash field to given value.

### HasFileHash

`func (o *ExtractedFileEntry) HasFileHash() bool`

HasFileHash returns a boolean if a field has been set.

### GetFileSize

`func (o *ExtractedFileEntry) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *ExtractedFileEntry) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *ExtractedFileEntry) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.


### GetFileType

`func (o *ExtractedFileEntry) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *ExtractedFileEntry) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *ExtractedFileEntry) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *ExtractedFileEntry) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetFilename

`func (o *ExtractedFileEntry) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ExtractedFileEntry) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ExtractedFileEntry) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetIsPe

`func (o *ExtractedFileEntry) GetIsPe() bool`

GetIsPe returns the IsPe field if non-nil, zero value otherwise.

### GetIsPeOk

`func (o *ExtractedFileEntry) GetIsPeOk() (*bool, bool)`

GetIsPeOk returns a tuple with the IsPe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPe

`func (o *ExtractedFileEntry) SetIsPe(v bool)`

SetIsPe sets IsPe field to given value.

### HasIsPe

`func (o *ExtractedFileEntry) HasIsPe() bool`

HasIsPe returns a boolean if a field has been set.

### GetMimeType

`func (o *ExtractedFileEntry) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *ExtractedFileEntry) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *ExtractedFileEntry) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *ExtractedFileEntry) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetReason

`func (o *ExtractedFileEntry) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ExtractedFileEntry) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ExtractedFileEntry) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ExtractedFileEntry) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSeqNum

`func (o *ExtractedFileEntry) GetSeqNum() int64`

GetSeqNum returns the SeqNum field if non-nil, zero value otherwise.

### GetSeqNumOk

`func (o *ExtractedFileEntry) GetSeqNumOk() (*int64, bool)`

GetSeqNumOk returns a tuple with the SeqNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqNum

`func (o *ExtractedFileEntry) SetSeqNum(v int64)`

SetSeqNum sets SeqNum field to given value.


### GetSha256

`func (o *ExtractedFileEntry) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *ExtractedFileEntry) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *ExtractedFileEntry) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *ExtractedFileEntry) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetZipFilename

`func (o *ExtractedFileEntry) GetZipFilename() string`

GetZipFilename returns the ZipFilename field if non-nil, zero value otherwise.

### GetZipFilenameOk

`func (o *ExtractedFileEntry) GetZipFilenameOk() (*string, bool)`

GetZipFilenameOk returns a tuple with the ZipFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipFilename

`func (o *ExtractedFileEntry) SetZipFilename(v string)`

SetZipFilename sets ZipFilename field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


