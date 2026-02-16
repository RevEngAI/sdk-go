# BinaryAdditionalDetailsDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | [**FileMetadata**](FileMetadata.md) |  | 
**Pe** | Pointer to [**NullablePEModel**](PEModel.md) |  | [optional] 
**Elf** | Pointer to [**NullableELFModel**](ELFModel.md) |  | [optional] 

## Methods

### NewBinaryAdditionalDetailsDataResponse

`func NewBinaryAdditionalDetailsDataResponse(file FileMetadata, ) *BinaryAdditionalDetailsDataResponse`

NewBinaryAdditionalDetailsDataResponse instantiates a new BinaryAdditionalDetailsDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBinaryAdditionalDetailsDataResponseWithDefaults

`func NewBinaryAdditionalDetailsDataResponseWithDefaults() *BinaryAdditionalDetailsDataResponse`

NewBinaryAdditionalDetailsDataResponseWithDefaults instantiates a new BinaryAdditionalDetailsDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *BinaryAdditionalDetailsDataResponse) GetFile() FileMetadata`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *BinaryAdditionalDetailsDataResponse) GetFileOk() (*FileMetadata, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *BinaryAdditionalDetailsDataResponse) SetFile(v FileMetadata)`

SetFile sets File field to given value.


### GetPe

`func (o *BinaryAdditionalDetailsDataResponse) GetPe() PEModel`

GetPe returns the Pe field if non-nil, zero value otherwise.

### GetPeOk

`func (o *BinaryAdditionalDetailsDataResponse) GetPeOk() (*PEModel, bool)`

GetPeOk returns a tuple with the Pe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPe

`func (o *BinaryAdditionalDetailsDataResponse) SetPe(v PEModel)`

SetPe sets Pe field to given value.

### HasPe

`func (o *BinaryAdditionalDetailsDataResponse) HasPe() bool`

HasPe returns a boolean if a field has been set.

### SetPeNil

`func (o *BinaryAdditionalDetailsDataResponse) SetPeNil(b bool)`

 SetPeNil sets the value for Pe to be an explicit nil

### UnsetPe
`func (o *BinaryAdditionalDetailsDataResponse) UnsetPe()`

UnsetPe ensures that no value is present for Pe, not even an explicit nil
### GetElf

`func (o *BinaryAdditionalDetailsDataResponse) GetElf() ELFModel`

GetElf returns the Elf field if non-nil, zero value otherwise.

### GetElfOk

`func (o *BinaryAdditionalDetailsDataResponse) GetElfOk() (*ELFModel, bool)`

GetElfOk returns a tuple with the Elf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElf

`func (o *BinaryAdditionalDetailsDataResponse) SetElf(v ELFModel)`

SetElf sets Elf field to given value.

### HasElf

`func (o *BinaryAdditionalDetailsDataResponse) HasElf() bool`

HasElf returns a boolean if a field has been set.

### SetElfNil

`func (o *BinaryAdditionalDetailsDataResponse) SetElfNil(b bool)`

 SetElfNil sets the value for Elf to be an explicit nil

### UnsetElf
`func (o *BinaryAdditionalDetailsDataResponse) UnsetElf()`

UnsetElf ensures that no value is present for Elf, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


