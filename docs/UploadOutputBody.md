# UploadOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisRequirements** | [**[]AnalysisRequirement**](AnalysisRequirement.md) | Ways to unblock POST /v3/analyses for this file if it cannot be statically analysed as-is; empty if no requirement applies. | 
**CanExtract** | **bool** | Whether the firmware/extraction flow can accept this file. | 
**CanSandbox** | **bool** | Whether the file can be extracted and run in the Windows sandbox. | 
**FileType** | **string** | The kind of file that was uploaded. | 
**Filename** | **string** | The filename as given by the caller. | 
**IsArchive** | **bool** | Whether the detected format is a container/compression archive. | 
**Mime** | **string** | The MIME type detected from the file&#39;s contents, independent of upload_file_type. | 
**Sha256Hash** | **string** | SHA-256 hash of the uploaded file; the storage key for every subsequent reference to it. | 

## Methods

### NewUploadOutputBody

`func NewUploadOutputBody(analysisRequirements []AnalysisRequirement, canExtract bool, canSandbox bool, fileType string, filename string, isArchive bool, mime string, sha256Hash string, ) *UploadOutputBody`

NewUploadOutputBody instantiates a new UploadOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadOutputBodyWithDefaults

`func NewUploadOutputBodyWithDefaults() *UploadOutputBody`

NewUploadOutputBodyWithDefaults instantiates a new UploadOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisRequirements

`func (o *UploadOutputBody) GetAnalysisRequirements() []AnalysisRequirement`

GetAnalysisRequirements returns the AnalysisRequirements field if non-nil, zero value otherwise.

### GetAnalysisRequirementsOk

`func (o *UploadOutputBody) GetAnalysisRequirementsOk() (*[]AnalysisRequirement, bool)`

GetAnalysisRequirementsOk returns a tuple with the AnalysisRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisRequirements

`func (o *UploadOutputBody) SetAnalysisRequirements(v []AnalysisRequirement)`

SetAnalysisRequirements sets AnalysisRequirements field to given value.


### SetAnalysisRequirementsNil

`func (o *UploadOutputBody) SetAnalysisRequirementsNil(b bool)`

 SetAnalysisRequirementsNil sets the value for AnalysisRequirements to be an explicit nil

### UnsetAnalysisRequirements
`func (o *UploadOutputBody) UnsetAnalysisRequirements()`

UnsetAnalysisRequirements ensures that no value is present for AnalysisRequirements, not even an explicit nil
### GetCanExtract

`func (o *UploadOutputBody) GetCanExtract() bool`

GetCanExtract returns the CanExtract field if non-nil, zero value otherwise.

### GetCanExtractOk

`func (o *UploadOutputBody) GetCanExtractOk() (*bool, bool)`

GetCanExtractOk returns a tuple with the CanExtract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanExtract

`func (o *UploadOutputBody) SetCanExtract(v bool)`

SetCanExtract sets CanExtract field to given value.


### GetCanSandbox

`func (o *UploadOutputBody) GetCanSandbox() bool`

GetCanSandbox returns the CanSandbox field if non-nil, zero value otherwise.

### GetCanSandboxOk

`func (o *UploadOutputBody) GetCanSandboxOk() (*bool, bool)`

GetCanSandboxOk returns a tuple with the CanSandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSandbox

`func (o *UploadOutputBody) SetCanSandbox(v bool)`

SetCanSandbox sets CanSandbox field to given value.


### GetFileType

`func (o *UploadOutputBody) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *UploadOutputBody) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *UploadOutputBody) SetFileType(v string)`

SetFileType sets FileType field to given value.


### GetFilename

`func (o *UploadOutputBody) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *UploadOutputBody) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *UploadOutputBody) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetIsArchive

`func (o *UploadOutputBody) GetIsArchive() bool`

GetIsArchive returns the IsArchive field if non-nil, zero value otherwise.

### GetIsArchiveOk

`func (o *UploadOutputBody) GetIsArchiveOk() (*bool, bool)`

GetIsArchiveOk returns a tuple with the IsArchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchive

`func (o *UploadOutputBody) SetIsArchive(v bool)`

SetIsArchive sets IsArchive field to given value.


### GetMime

`func (o *UploadOutputBody) GetMime() string`

GetMime returns the Mime field if non-nil, zero value otherwise.

### GetMimeOk

`func (o *UploadOutputBody) GetMimeOk() (*string, bool)`

GetMimeOk returns a tuple with the Mime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMime

`func (o *UploadOutputBody) SetMime(v string)`

SetMime sets Mime field to given value.


### GetSha256Hash

`func (o *UploadOutputBody) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *UploadOutputBody) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *UploadOutputBody) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


