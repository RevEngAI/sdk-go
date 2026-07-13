# ArchiveContentEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Encrypted** | **bool** | Whether this entry is password-protected | 
**Path** | **string** | Path relative to the archive root | 
**Size** | **int64** | Uncompressed size in bytes | 

## Methods

### NewArchiveContentEntry

`func NewArchiveContentEntry(encrypted bool, path string, size int64, ) *ArchiveContentEntry`

NewArchiveContentEntry instantiates a new ArchiveContentEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchiveContentEntryWithDefaults

`func NewArchiveContentEntryWithDefaults() *ArchiveContentEntry`

NewArchiveContentEntryWithDefaults instantiates a new ArchiveContentEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncrypted

`func (o *ArchiveContentEntry) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *ArchiveContentEntry) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *ArchiveContentEntry) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.


### GetPath

`func (o *ArchiveContentEntry) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ArchiveContentEntry) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ArchiveContentEntry) SetPath(v string)`

SetPath sets Path field to given value.


### GetSize

`func (o *ArchiveContentEntry) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ArchiveContentEntry) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ArchiveContentEntry) SetSize(v int64)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


