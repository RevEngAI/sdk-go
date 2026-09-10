# ExtractedBinary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | **string** | Original filename inside the upload | 
**Sha256Hash** | **string** | Content hash of the recovered binary | 
**Size** | **int64** | File size (in bytes) | 

## Methods

### NewExtractedBinary

`func NewExtractedBinary(filename string, sha256Hash string, size int64, ) *ExtractedBinary`

NewExtractedBinary instantiates a new ExtractedBinary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtractedBinaryWithDefaults

`func NewExtractedBinaryWithDefaults() *ExtractedBinary`

NewExtractedBinaryWithDefaults instantiates a new ExtractedBinary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *ExtractedBinary) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ExtractedBinary) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ExtractedBinary) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetSha256Hash

`func (o *ExtractedBinary) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *ExtractedBinary) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *ExtractedBinary) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.


### GetSize

`func (o *ExtractedBinary) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ExtractedBinary) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ExtractedBinary) SetSize(v int64)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


