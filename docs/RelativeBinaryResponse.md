# RelativeBinaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinaryId** | **int32** | ID of the relative binary | 
**Name** | **string** | Name of the relative binary | 
**Sha256** | **string** | SHA256 hash of the relative binary | 

## Methods

### NewRelativeBinaryResponse

`func NewRelativeBinaryResponse(binaryId int32, name string, sha256 string, ) *RelativeBinaryResponse`

NewRelativeBinaryResponse instantiates a new RelativeBinaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelativeBinaryResponseWithDefaults

`func NewRelativeBinaryResponseWithDefaults() *RelativeBinaryResponse`

NewRelativeBinaryResponseWithDefaults instantiates a new RelativeBinaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaryId

`func (o *RelativeBinaryResponse) GetBinaryId() int32`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *RelativeBinaryResponse) GetBinaryIdOk() (*int32, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *RelativeBinaryResponse) SetBinaryId(v int32)`

SetBinaryId sets BinaryId field to given value.


### GetName

`func (o *RelativeBinaryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RelativeBinaryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RelativeBinaryResponse) SetName(v string)`

SetName sets Name field to given value.


### GetSha256

`func (o *RelativeBinaryResponse) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *RelativeBinaryResponse) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *RelativeBinaryResponse) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


