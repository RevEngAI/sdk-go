# Artifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to **string** |  | [optional] 
**DumpAddr** | Pointer to **string** |  | [optional] 
**DumpPid** | Pointer to **int64** |  | [optional] 
**FileType** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**IsPe** | **bool** |  | 
**MimeType** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**NetworkSource** | Pointer to **string** |  | [optional] 
**OriginalFilename** | Pointer to **string** |  | [optional] 
**Path** | **string** |  | 
**ProcessSeqid** | Pointer to **int64** |  | [optional] 
**Reason** | **string** |  | 
**ResponseStatus** | Pointer to **int64** |  | [optional] 
**Sha256** | Pointer to **string** |  | [optional] 
**Size** | **int64** |  | 
**Source** | **string** |  | 
**Uri** | Pointer to **string** |  | [optional] 
**WasMapped** | Pointer to **bool** |  | [optional] 
**YaraHits** | Pointer to **[]string** |  | [optional] 

## Methods

### NewArtifact

`func NewArtifact(isPe bool, name string, path string, reason string, size int64, source string, ) *Artifact`

NewArtifact instantiates a new Artifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactWithDefaults

`func NewArtifactWithDefaults() *Artifact`

NewArtifactWithDefaults instantiates a new Artifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *Artifact) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *Artifact) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *Artifact) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *Artifact) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetDumpAddr

`func (o *Artifact) GetDumpAddr() string`

GetDumpAddr returns the DumpAddr field if non-nil, zero value otherwise.

### GetDumpAddrOk

`func (o *Artifact) GetDumpAddrOk() (*string, bool)`

GetDumpAddrOk returns a tuple with the DumpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDumpAddr

`func (o *Artifact) SetDumpAddr(v string)`

SetDumpAddr sets DumpAddr field to given value.

### HasDumpAddr

`func (o *Artifact) HasDumpAddr() bool`

HasDumpAddr returns a boolean if a field has been set.

### GetDumpPid

`func (o *Artifact) GetDumpPid() int64`

GetDumpPid returns the DumpPid field if non-nil, zero value otherwise.

### GetDumpPidOk

`func (o *Artifact) GetDumpPidOk() (*int64, bool)`

GetDumpPidOk returns a tuple with the DumpPid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDumpPid

`func (o *Artifact) SetDumpPid(v int64)`

SetDumpPid sets DumpPid field to given value.

### HasDumpPid

`func (o *Artifact) HasDumpPid() bool`

HasDumpPid returns a boolean if a field has been set.

### GetFileType

`func (o *Artifact) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *Artifact) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *Artifact) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *Artifact) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetHost

`func (o *Artifact) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Artifact) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Artifact) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *Artifact) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetIsPe

`func (o *Artifact) GetIsPe() bool`

GetIsPe returns the IsPe field if non-nil, zero value otherwise.

### GetIsPeOk

`func (o *Artifact) GetIsPeOk() (*bool, bool)`

GetIsPeOk returns a tuple with the IsPe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPe

`func (o *Artifact) SetIsPe(v bool)`

SetIsPe sets IsPe field to given value.


### GetMimeType

`func (o *Artifact) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *Artifact) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *Artifact) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *Artifact) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetName

`func (o *Artifact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Artifact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Artifact) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkSource

`func (o *Artifact) GetNetworkSource() string`

GetNetworkSource returns the NetworkSource field if non-nil, zero value otherwise.

### GetNetworkSourceOk

`func (o *Artifact) GetNetworkSourceOk() (*string, bool)`

GetNetworkSourceOk returns a tuple with the NetworkSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSource

`func (o *Artifact) SetNetworkSource(v string)`

SetNetworkSource sets NetworkSource field to given value.

### HasNetworkSource

`func (o *Artifact) HasNetworkSource() bool`

HasNetworkSource returns a boolean if a field has been set.

### GetOriginalFilename

`func (o *Artifact) GetOriginalFilename() string`

GetOriginalFilename returns the OriginalFilename field if non-nil, zero value otherwise.

### GetOriginalFilenameOk

`func (o *Artifact) GetOriginalFilenameOk() (*string, bool)`

GetOriginalFilenameOk returns a tuple with the OriginalFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFilename

`func (o *Artifact) SetOriginalFilename(v string)`

SetOriginalFilename sets OriginalFilename field to given value.

### HasOriginalFilename

`func (o *Artifact) HasOriginalFilename() bool`

HasOriginalFilename returns a boolean if a field has been set.

### GetPath

`func (o *Artifact) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Artifact) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Artifact) SetPath(v string)`

SetPath sets Path field to given value.


### GetProcessSeqid

`func (o *Artifact) GetProcessSeqid() int64`

GetProcessSeqid returns the ProcessSeqid field if non-nil, zero value otherwise.

### GetProcessSeqidOk

`func (o *Artifact) GetProcessSeqidOk() (*int64, bool)`

GetProcessSeqidOk returns a tuple with the ProcessSeqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessSeqid

`func (o *Artifact) SetProcessSeqid(v int64)`

SetProcessSeqid sets ProcessSeqid field to given value.

### HasProcessSeqid

`func (o *Artifact) HasProcessSeqid() bool`

HasProcessSeqid returns a boolean if a field has been set.

### GetReason

`func (o *Artifact) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Artifact) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Artifact) SetReason(v string)`

SetReason sets Reason field to given value.


### GetResponseStatus

`func (o *Artifact) GetResponseStatus() int64`

GetResponseStatus returns the ResponseStatus field if non-nil, zero value otherwise.

### GetResponseStatusOk

`func (o *Artifact) GetResponseStatusOk() (*int64, bool)`

GetResponseStatusOk returns a tuple with the ResponseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatus

`func (o *Artifact) SetResponseStatus(v int64)`

SetResponseStatus sets ResponseStatus field to given value.

### HasResponseStatus

`func (o *Artifact) HasResponseStatus() bool`

HasResponseStatus returns a boolean if a field has been set.

### GetSha256

`func (o *Artifact) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *Artifact) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *Artifact) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *Artifact) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetSize

`func (o *Artifact) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Artifact) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Artifact) SetSize(v int64)`

SetSize sets Size field to given value.


### GetSource

`func (o *Artifact) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Artifact) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Artifact) SetSource(v string)`

SetSource sets Source field to given value.


### GetUri

`func (o *Artifact) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Artifact) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Artifact) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Artifact) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetWasMapped

`func (o *Artifact) GetWasMapped() bool`

GetWasMapped returns the WasMapped field if non-nil, zero value otherwise.

### GetWasMappedOk

`func (o *Artifact) GetWasMappedOk() (*bool, bool)`

GetWasMappedOk returns a tuple with the WasMapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasMapped

`func (o *Artifact) SetWasMapped(v bool)`

SetWasMapped sets WasMapped field to given value.

### HasWasMapped

`func (o *Artifact) HasWasMapped() bool`

HasWasMapped returns a boolean if a field has been set.

### GetYaraHits

`func (o *Artifact) GetYaraHits() []string`

GetYaraHits returns the YaraHits field if non-nil, zero value otherwise.

### GetYaraHitsOk

`func (o *Artifact) GetYaraHitsOk() (*[]string, bool)`

GetYaraHitsOk returns a tuple with the YaraHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaraHits

`func (o *Artifact) SetYaraHits(v []string)`

SetYaraHits sets YaraHits field to given value.

### HasYaraHits

`func (o *Artifact) HasYaraHits() bool`

HasYaraHits returns a boolean if a field has been set.

### SetYaraHitsNil

`func (o *Artifact) SetYaraHitsNil(b bool)`

 SetYaraHitsNil sets the value for YaraHits to be an explicit nil

### UnsetYaraHits
`func (o *Artifact) UnsetYaraHits()`

UnsetYaraHits ensures that no value is present for YaraHits, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


