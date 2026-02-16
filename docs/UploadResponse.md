# UploadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sha256Hash** | **string** |  | 
**FileType** | [**UploadFileType**](UploadFileType.md) |  | 
**Filename** | **string** |  | 

## Methods

### NewUploadResponse

`func NewUploadResponse(sha256Hash string, fileType UploadFileType, filename string, ) *UploadResponse`

NewUploadResponse instantiates a new UploadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadResponseWithDefaults

`func NewUploadResponseWithDefaults() *UploadResponse`

NewUploadResponseWithDefaults instantiates a new UploadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSha256Hash

`func (o *UploadResponse) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *UploadResponse) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *UploadResponse) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.


### GetFileType

`func (o *UploadResponse) GetFileType() UploadFileType`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *UploadResponse) GetFileTypeOk() (*UploadFileType, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *UploadResponse) SetFileType(v UploadFileType)`

SetFileType sets FileType field to given value.


### GetFilename

`func (o *UploadResponse) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *UploadResponse) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *UploadResponse) SetFilename(v string)`

SetFilename sets Filename field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


