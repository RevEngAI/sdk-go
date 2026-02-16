# FileMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | **int32** |  | 
**FriendlySize** | **string** |  | 
**Entropy** | **float32** |  | 
**Hashes** | [**FileHashes**](FileHashes.md) |  | 

## Methods

### NewFileMetadata

`func NewFileMetadata(size int32, friendlySize string, entropy float32, hashes FileHashes, ) *FileMetadata`

NewFileMetadata instantiates a new FileMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileMetadataWithDefaults

`func NewFileMetadataWithDefaults() *FileMetadata`

NewFileMetadataWithDefaults instantiates a new FileMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *FileMetadata) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FileMetadata) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FileMetadata) SetSize(v int32)`

SetSize sets Size field to given value.


### GetFriendlySize

`func (o *FileMetadata) GetFriendlySize() string`

GetFriendlySize returns the FriendlySize field if non-nil, zero value otherwise.

### GetFriendlySizeOk

`func (o *FileMetadata) GetFriendlySizeOk() (*string, bool)`

GetFriendlySizeOk returns a tuple with the FriendlySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlySize

`func (o *FileMetadata) SetFriendlySize(v string)`

SetFriendlySize sets FriendlySize field to given value.


### GetEntropy

`func (o *FileMetadata) GetEntropy() float32`

GetEntropy returns the Entropy field if non-nil, zero value otherwise.

### GetEntropyOk

`func (o *FileMetadata) GetEntropyOk() (*float32, bool)`

GetEntropyOk returns a tuple with the Entropy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntropy

`func (o *FileMetadata) SetEntropy(v float32)`

SetEntropy sets Entropy field to given value.


### GetHashes

`func (o *FileMetadata) GetHashes() FileHashes`

GetHashes returns the Hashes field if non-nil, zero value otherwise.

### GetHashesOk

`func (o *FileMetadata) GetHashesOk() (*FileHashes, bool)`

GetHashesOk returns a tuple with the Hashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashes

`func (o *FileMetadata) SetHashes(v FileHashes)`

SetHashes sets Hashes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


