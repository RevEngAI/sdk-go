# PcapBodyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | Pointer to **string** |  | [optional] 
**IsPe** | **bool** |  | 
**MimeType** | Pointer to **string** |  | [optional] 
**Preview** | Pointer to **string** |  | [optional] 
**Sha256** | Pointer to **string** |  | [optional] 
**Size** | **int64** |  | 
**YaraHits** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPcapBodyInfo

`func NewPcapBodyInfo(isPe bool, size int64, ) *PcapBodyInfo`

NewPcapBodyInfo instantiates a new PcapBodyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPcapBodyInfoWithDefaults

`func NewPcapBodyInfoWithDefaults() *PcapBodyInfo`

NewPcapBodyInfoWithDefaults instantiates a new PcapBodyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *PcapBodyInfo) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *PcapBodyInfo) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *PcapBodyInfo) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *PcapBodyInfo) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetIsPe

`func (o *PcapBodyInfo) GetIsPe() bool`

GetIsPe returns the IsPe field if non-nil, zero value otherwise.

### GetIsPeOk

`func (o *PcapBodyInfo) GetIsPeOk() (*bool, bool)`

GetIsPeOk returns a tuple with the IsPe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPe

`func (o *PcapBodyInfo) SetIsPe(v bool)`

SetIsPe sets IsPe field to given value.


### GetMimeType

`func (o *PcapBodyInfo) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *PcapBodyInfo) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *PcapBodyInfo) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *PcapBodyInfo) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetPreview

`func (o *PcapBodyInfo) GetPreview() string`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *PcapBodyInfo) GetPreviewOk() (*string, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *PcapBodyInfo) SetPreview(v string)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *PcapBodyInfo) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### GetSha256

`func (o *PcapBodyInfo) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *PcapBodyInfo) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *PcapBodyInfo) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *PcapBodyInfo) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetSize

`func (o *PcapBodyInfo) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PcapBodyInfo) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PcapBodyInfo) SetSize(v int64)`

SetSize sets Size field to given value.


### GetYaraHits

`func (o *PcapBodyInfo) GetYaraHits() []string`

GetYaraHits returns the YaraHits field if non-nil, zero value otherwise.

### GetYaraHitsOk

`func (o *PcapBodyInfo) GetYaraHitsOk() (*[]string, bool)`

GetYaraHitsOk returns a tuple with the YaraHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaraHits

`func (o *PcapBodyInfo) SetYaraHits(v []string)`

SetYaraHits sets YaraHits field to given value.

### HasYaraHits

`func (o *PcapBodyInfo) HasYaraHits() bool`

HasYaraHits returns a boolean if a field has been set.

### SetYaraHitsNil

`func (o *PcapBodyInfo) SetYaraHitsNil(b bool)`

 SetYaraHitsNil sets the value for YaraHits to be an explicit nil

### UnsetYaraHits
`func (o *PcapBodyInfo) UnsetYaraHits()`

UnsetYaraHits ensures that no value is present for YaraHits, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


