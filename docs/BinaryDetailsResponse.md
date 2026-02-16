# BinaryDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | **string** | The architecture of the binary | 
**Bits** | **int32** | The size of the binary in bits | 
**Crc32** | **string** |  | 
**Class** | **string** |  | 
**Entropy** | **float32** |  | 
**FileSize** | **int32** |  | 
**Language** | **string** |  | 
**Md5** | **string** |  | 
**Machine** | **string** |  | 
**Os** | **string** | OS target of the binary | 
**Sha1** | **string** | SHA1 hash of the binary | 
**Sha256** | **string** | SHA256 hash of the binary | 
**Ssdeep** | **string** |  | 
**Static** | **bool** |  | 
**Stripped** | **bool** |  | 
**SubSys** | **string** |  | 
**Tlsh** | **string** |  | 
**Type** | **string** |  | 
**Debug** | **bool** |  | 
**FirstSeen** | **time.Time** |  | 

## Methods

### NewBinaryDetailsResponse

`func NewBinaryDetailsResponse(arch string, bits int32, crc32 string, class string, entropy float32, fileSize int32, language string, md5 string, machine string, os string, sha1 string, sha256 string, ssdeep string, static bool, stripped bool, subSys string, tlsh string, type_ string, debug bool, firstSeen time.Time, ) *BinaryDetailsResponse`

NewBinaryDetailsResponse instantiates a new BinaryDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBinaryDetailsResponseWithDefaults

`func NewBinaryDetailsResponseWithDefaults() *BinaryDetailsResponse`

NewBinaryDetailsResponseWithDefaults instantiates a new BinaryDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *BinaryDetailsResponse) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *BinaryDetailsResponse) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *BinaryDetailsResponse) SetArch(v string)`

SetArch sets Arch field to given value.


### GetBits

`func (o *BinaryDetailsResponse) GetBits() int32`

GetBits returns the Bits field if non-nil, zero value otherwise.

### GetBitsOk

`func (o *BinaryDetailsResponse) GetBitsOk() (*int32, bool)`

GetBitsOk returns a tuple with the Bits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBits

`func (o *BinaryDetailsResponse) SetBits(v int32)`

SetBits sets Bits field to given value.


### GetCrc32

`func (o *BinaryDetailsResponse) GetCrc32() string`

GetCrc32 returns the Crc32 field if non-nil, zero value otherwise.

### GetCrc32Ok

`func (o *BinaryDetailsResponse) GetCrc32Ok() (*string, bool)`

GetCrc32Ok returns a tuple with the Crc32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrc32

`func (o *BinaryDetailsResponse) SetCrc32(v string)`

SetCrc32 sets Crc32 field to given value.


### GetClass

`func (o *BinaryDetailsResponse) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *BinaryDetailsResponse) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *BinaryDetailsResponse) SetClass(v string)`

SetClass sets Class field to given value.


### GetEntropy

`func (o *BinaryDetailsResponse) GetEntropy() float32`

GetEntropy returns the Entropy field if non-nil, zero value otherwise.

### GetEntropyOk

`func (o *BinaryDetailsResponse) GetEntropyOk() (*float32, bool)`

GetEntropyOk returns a tuple with the Entropy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntropy

`func (o *BinaryDetailsResponse) SetEntropy(v float32)`

SetEntropy sets Entropy field to given value.


### GetFileSize

`func (o *BinaryDetailsResponse) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *BinaryDetailsResponse) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *BinaryDetailsResponse) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.


### GetLanguage

`func (o *BinaryDetailsResponse) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *BinaryDetailsResponse) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *BinaryDetailsResponse) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetMd5

`func (o *BinaryDetailsResponse) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *BinaryDetailsResponse) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *BinaryDetailsResponse) SetMd5(v string)`

SetMd5 sets Md5 field to given value.


### GetMachine

`func (o *BinaryDetailsResponse) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *BinaryDetailsResponse) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *BinaryDetailsResponse) SetMachine(v string)`

SetMachine sets Machine field to given value.


### GetOs

`func (o *BinaryDetailsResponse) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *BinaryDetailsResponse) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *BinaryDetailsResponse) SetOs(v string)`

SetOs sets Os field to given value.


### GetSha1

`func (o *BinaryDetailsResponse) GetSha1() string`

GetSha1 returns the Sha1 field if non-nil, zero value otherwise.

### GetSha1Ok

`func (o *BinaryDetailsResponse) GetSha1Ok() (*string, bool)`

GetSha1Ok returns a tuple with the Sha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1

`func (o *BinaryDetailsResponse) SetSha1(v string)`

SetSha1 sets Sha1 field to given value.


### GetSha256

`func (o *BinaryDetailsResponse) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *BinaryDetailsResponse) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *BinaryDetailsResponse) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.


### GetSsdeep

`func (o *BinaryDetailsResponse) GetSsdeep() string`

GetSsdeep returns the Ssdeep field if non-nil, zero value otherwise.

### GetSsdeepOk

`func (o *BinaryDetailsResponse) GetSsdeepOk() (*string, bool)`

GetSsdeepOk returns a tuple with the Ssdeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsdeep

`func (o *BinaryDetailsResponse) SetSsdeep(v string)`

SetSsdeep sets Ssdeep field to given value.


### GetStatic

`func (o *BinaryDetailsResponse) GetStatic() bool`

GetStatic returns the Static field if non-nil, zero value otherwise.

### GetStaticOk

`func (o *BinaryDetailsResponse) GetStaticOk() (*bool, bool)`

GetStaticOk returns a tuple with the Static field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic

`func (o *BinaryDetailsResponse) SetStatic(v bool)`

SetStatic sets Static field to given value.


### GetStripped

`func (o *BinaryDetailsResponse) GetStripped() bool`

GetStripped returns the Stripped field if non-nil, zero value otherwise.

### GetStrippedOk

`func (o *BinaryDetailsResponse) GetStrippedOk() (*bool, bool)`

GetStrippedOk returns a tuple with the Stripped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripped

`func (o *BinaryDetailsResponse) SetStripped(v bool)`

SetStripped sets Stripped field to given value.


### GetSubSys

`func (o *BinaryDetailsResponse) GetSubSys() string`

GetSubSys returns the SubSys field if non-nil, zero value otherwise.

### GetSubSysOk

`func (o *BinaryDetailsResponse) GetSubSysOk() (*string, bool)`

GetSubSysOk returns a tuple with the SubSys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubSys

`func (o *BinaryDetailsResponse) SetSubSys(v string)`

SetSubSys sets SubSys field to given value.


### GetTlsh

`func (o *BinaryDetailsResponse) GetTlsh() string`

GetTlsh returns the Tlsh field if non-nil, zero value otherwise.

### GetTlshOk

`func (o *BinaryDetailsResponse) GetTlshOk() (*string, bool)`

GetTlshOk returns a tuple with the Tlsh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsh

`func (o *BinaryDetailsResponse) SetTlsh(v string)`

SetTlsh sets Tlsh field to given value.


### GetType

`func (o *BinaryDetailsResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BinaryDetailsResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BinaryDetailsResponse) SetType(v string)`

SetType sets Type field to given value.


### GetDebug

`func (o *BinaryDetailsResponse) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *BinaryDetailsResponse) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *BinaryDetailsResponse) SetDebug(v bool)`

SetDebug sets Debug field to given value.


### GetFirstSeen

`func (o *BinaryDetailsResponse) GetFirstSeen() time.Time`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *BinaryDetailsResponse) GetFirstSeenOk() (*time.Time, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *BinaryDetailsResponse) SetFirstSeen(v time.Time)`

SetFirstSeen sets FirstSeen field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


