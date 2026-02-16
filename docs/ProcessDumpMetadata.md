# ProcessDumpMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sha256** | **string** |  | 
**Type** | **string** |  | 
**Size** | **int32** |  | 

## Methods

### NewProcessDumpMetadata

`func NewProcessDumpMetadata(sha256 string, type_ string, size int32, ) *ProcessDumpMetadata`

NewProcessDumpMetadata instantiates a new ProcessDumpMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessDumpMetadataWithDefaults

`func NewProcessDumpMetadataWithDefaults() *ProcessDumpMetadata`

NewProcessDumpMetadataWithDefaults instantiates a new ProcessDumpMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSha256

`func (o *ProcessDumpMetadata) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *ProcessDumpMetadata) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *ProcessDumpMetadata) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.


### GetType

`func (o *ProcessDumpMetadata) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProcessDumpMetadata) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProcessDumpMetadata) SetType(v string)`

SetType sets Type field to given value.


### GetSize

`func (o *ProcessDumpMetadata) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ProcessDumpMetadata) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ProcessDumpMetadata) SetSize(v int32)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


