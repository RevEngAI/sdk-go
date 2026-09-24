# ImportDynamicExecutionFileOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanExtract** | **bool** | Whether the firmware/extraction flow can accept this file. | 
**CanSandbox** | **bool** | Whether the file can be extracted and run in the Windows sandbox. | 
**IsArchive** | **bool** | Whether the detected format is a container/compression archive. | 
**Mime** | **string** | The MIME type detected from the file&#39;s contents. | 
**Sha256Hash** | **string** | SHA-256 hash the file is stored under; the storage key for every subsequent reference to it. | 
**Size** | **int64** | Size of the file in bytes. | 

## Methods

### NewImportDynamicExecutionFileOutputBody

`func NewImportDynamicExecutionFileOutputBody(canExtract bool, canSandbox bool, isArchive bool, mime string, sha256Hash string, size int64, ) *ImportDynamicExecutionFileOutputBody`

NewImportDynamicExecutionFileOutputBody instantiates a new ImportDynamicExecutionFileOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportDynamicExecutionFileOutputBodyWithDefaults

`func NewImportDynamicExecutionFileOutputBodyWithDefaults() *ImportDynamicExecutionFileOutputBody`

NewImportDynamicExecutionFileOutputBodyWithDefaults instantiates a new ImportDynamicExecutionFileOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanExtract

`func (o *ImportDynamicExecutionFileOutputBody) GetCanExtract() bool`

GetCanExtract returns the CanExtract field if non-nil, zero value otherwise.

### GetCanExtractOk

`func (o *ImportDynamicExecutionFileOutputBody) GetCanExtractOk() (*bool, bool)`

GetCanExtractOk returns a tuple with the CanExtract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanExtract

`func (o *ImportDynamicExecutionFileOutputBody) SetCanExtract(v bool)`

SetCanExtract sets CanExtract field to given value.


### GetCanSandbox

`func (o *ImportDynamicExecutionFileOutputBody) GetCanSandbox() bool`

GetCanSandbox returns the CanSandbox field if non-nil, zero value otherwise.

### GetCanSandboxOk

`func (o *ImportDynamicExecutionFileOutputBody) GetCanSandboxOk() (*bool, bool)`

GetCanSandboxOk returns a tuple with the CanSandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSandbox

`func (o *ImportDynamicExecutionFileOutputBody) SetCanSandbox(v bool)`

SetCanSandbox sets CanSandbox field to given value.


### GetIsArchive

`func (o *ImportDynamicExecutionFileOutputBody) GetIsArchive() bool`

GetIsArchive returns the IsArchive field if non-nil, zero value otherwise.

### GetIsArchiveOk

`func (o *ImportDynamicExecutionFileOutputBody) GetIsArchiveOk() (*bool, bool)`

GetIsArchiveOk returns a tuple with the IsArchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchive

`func (o *ImportDynamicExecutionFileOutputBody) SetIsArchive(v bool)`

SetIsArchive sets IsArchive field to given value.


### GetMime

`func (o *ImportDynamicExecutionFileOutputBody) GetMime() string`

GetMime returns the Mime field if non-nil, zero value otherwise.

### GetMimeOk

`func (o *ImportDynamicExecutionFileOutputBody) GetMimeOk() (*string, bool)`

GetMimeOk returns a tuple with the Mime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMime

`func (o *ImportDynamicExecutionFileOutputBody) SetMime(v string)`

SetMime sets Mime field to given value.


### GetSha256Hash

`func (o *ImportDynamicExecutionFileOutputBody) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *ImportDynamicExecutionFileOutputBody) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *ImportDynamicExecutionFileOutputBody) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.


### GetSize

`func (o *ImportDynamicExecutionFileOutputBody) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ImportDynamicExecutionFileOutputBody) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ImportDynamicExecutionFileOutputBody) SetSize(v int64)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


