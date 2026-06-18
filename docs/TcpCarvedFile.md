# TcpCarvedFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | **string** |  | 
**Filename** | Pointer to **string** |  | [optional] 
**IsPe** | **bool** |  | 
**MimeType** | Pointer to **string** |  | [optional] 
**Offset** | **int64** |  | 
**Sha256** | **string** |  | 
**Size** | **int64** |  | 
**YaraHits** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTcpCarvedFile

`func NewTcpCarvedFile(direction string, isPe bool, offset int64, sha256 string, size int64, ) *TcpCarvedFile`

NewTcpCarvedFile instantiates a new TcpCarvedFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTcpCarvedFileWithDefaults

`func NewTcpCarvedFileWithDefaults() *TcpCarvedFile`

NewTcpCarvedFileWithDefaults instantiates a new TcpCarvedFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *TcpCarvedFile) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *TcpCarvedFile) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *TcpCarvedFile) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetFilename

`func (o *TcpCarvedFile) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *TcpCarvedFile) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *TcpCarvedFile) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *TcpCarvedFile) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetIsPe

`func (o *TcpCarvedFile) GetIsPe() bool`

GetIsPe returns the IsPe field if non-nil, zero value otherwise.

### GetIsPeOk

`func (o *TcpCarvedFile) GetIsPeOk() (*bool, bool)`

GetIsPeOk returns a tuple with the IsPe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPe

`func (o *TcpCarvedFile) SetIsPe(v bool)`

SetIsPe sets IsPe field to given value.


### GetMimeType

`func (o *TcpCarvedFile) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *TcpCarvedFile) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *TcpCarvedFile) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *TcpCarvedFile) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetOffset

`func (o *TcpCarvedFile) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TcpCarvedFile) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TcpCarvedFile) SetOffset(v int64)`

SetOffset sets Offset field to given value.


### GetSha256

`func (o *TcpCarvedFile) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *TcpCarvedFile) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *TcpCarvedFile) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.


### GetSize

`func (o *TcpCarvedFile) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TcpCarvedFile) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TcpCarvedFile) SetSize(v int64)`

SetSize sets Size field to given value.


### GetYaraHits

`func (o *TcpCarvedFile) GetYaraHits() []string`

GetYaraHits returns the YaraHits field if non-nil, zero value otherwise.

### GetYaraHitsOk

`func (o *TcpCarvedFile) GetYaraHitsOk() (*[]string, bool)`

GetYaraHitsOk returns a tuple with the YaraHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaraHits

`func (o *TcpCarvedFile) SetYaraHits(v []string)`

SetYaraHits sets YaraHits field to given value.

### HasYaraHits

`func (o *TcpCarvedFile) HasYaraHits() bool`

HasYaraHits returns a boolean if a field has been set.

### SetYaraHitsNil

`func (o *TcpCarvedFile) SetYaraHitsNil(b bool)`

 SetYaraHitsNil sets the value for YaraHits to be an explicit nil

### UnsetYaraHits
`func (o *TcpCarvedFile) UnsetYaraHits()`

UnsetYaraHits ensures that no value is present for YaraHits, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


