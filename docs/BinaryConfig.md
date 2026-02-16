# BinaryConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Isa** | Pointer to [**NullableISA**](ISA.md) |  | [optional] 
**Platform** | Pointer to [**NullablePlatform**](Platform.md) |  | [optional] 
**FileFormat** | Pointer to [**NullableFileFormat**](FileFormat.md) |  | [optional] 

## Methods

### NewBinaryConfig

`func NewBinaryConfig() *BinaryConfig`

NewBinaryConfig instantiates a new BinaryConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBinaryConfigWithDefaults

`func NewBinaryConfigWithDefaults() *BinaryConfig`

NewBinaryConfigWithDefaults instantiates a new BinaryConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsa

`func (o *BinaryConfig) GetIsa() ISA`

GetIsa returns the Isa field if non-nil, zero value otherwise.

### GetIsaOk

`func (o *BinaryConfig) GetIsaOk() (*ISA, bool)`

GetIsaOk returns a tuple with the Isa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsa

`func (o *BinaryConfig) SetIsa(v ISA)`

SetIsa sets Isa field to given value.

### HasIsa

`func (o *BinaryConfig) HasIsa() bool`

HasIsa returns a boolean if a field has been set.

### SetIsaNil

`func (o *BinaryConfig) SetIsaNil(b bool)`

 SetIsaNil sets the value for Isa to be an explicit nil

### UnsetIsa
`func (o *BinaryConfig) UnsetIsa()`

UnsetIsa ensures that no value is present for Isa, not even an explicit nil
### GetPlatform

`func (o *BinaryConfig) GetPlatform() Platform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *BinaryConfig) GetPlatformOk() (*Platform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *BinaryConfig) SetPlatform(v Platform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *BinaryConfig) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *BinaryConfig) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *BinaryConfig) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetFileFormat

`func (o *BinaryConfig) GetFileFormat() FileFormat`

GetFileFormat returns the FileFormat field if non-nil, zero value otherwise.

### GetFileFormatOk

`func (o *BinaryConfig) GetFileFormatOk() (*FileFormat, bool)`

GetFileFormatOk returns a tuple with the FileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFormat

`func (o *BinaryConfig) SetFileFormat(v FileFormat)`

SetFileFormat sets FileFormat field to given value.

### HasFileFormat

`func (o *BinaryConfig) HasFileFormat() bool`

HasFileFormat returns a boolean if a field has been set.

### SetFileFormatNil

`func (o *BinaryConfig) SetFileFormatNil(b bool)`

 SetFileFormatNil sets the value for FileFormat to be an explicit nil

### UnsetFileFormat
`func (o *BinaryConfig) UnsetFileFormat()`

UnsetFileFormat ensures that no value is present for FileFormat, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


