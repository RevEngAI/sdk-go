# BinaryAdditionalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinaryId** | **int32** |  | 
**Details** | [**NullableBinaryAdditionalDetailsDataResponse**](BinaryAdditionalDetailsDataResponse.md) |  | 
**Creation** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBinaryAdditionalResponse

`func NewBinaryAdditionalResponse(binaryId int32, details NullableBinaryAdditionalDetailsDataResponse, ) *BinaryAdditionalResponse`

NewBinaryAdditionalResponse instantiates a new BinaryAdditionalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBinaryAdditionalResponseWithDefaults

`func NewBinaryAdditionalResponseWithDefaults() *BinaryAdditionalResponse`

NewBinaryAdditionalResponseWithDefaults instantiates a new BinaryAdditionalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaryId

`func (o *BinaryAdditionalResponse) GetBinaryId() int32`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *BinaryAdditionalResponse) GetBinaryIdOk() (*int32, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *BinaryAdditionalResponse) SetBinaryId(v int32)`

SetBinaryId sets BinaryId field to given value.


### GetDetails

`func (o *BinaryAdditionalResponse) GetDetails() BinaryAdditionalDetailsDataResponse`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *BinaryAdditionalResponse) GetDetailsOk() (*BinaryAdditionalDetailsDataResponse, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *BinaryAdditionalResponse) SetDetails(v BinaryAdditionalDetailsDataResponse)`

SetDetails sets Details field to given value.


### SetDetailsNil

`func (o *BinaryAdditionalResponse) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *BinaryAdditionalResponse) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetCreation

`func (o *BinaryAdditionalResponse) GetCreation() time.Time`

GetCreation returns the Creation field if non-nil, zero value otherwise.

### GetCreationOk

`func (o *BinaryAdditionalResponse) GetCreationOk() (*time.Time, bool)`

GetCreationOk returns a tuple with the Creation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreation

`func (o *BinaryAdditionalResponse) SetCreation(v time.Time)`

SetCreation sets Creation field to given value.

### HasCreation

`func (o *BinaryAdditionalResponse) HasCreation() bool`

HasCreation returns a boolean if a field has been set.

### SetCreationNil

`func (o *BinaryAdditionalResponse) SetCreationNil(b bool)`

 SetCreationNil sets the value for Creation to be an explicit nil

### UnsetCreation
`func (o *BinaryAdditionalResponse) UnsetCreation()`

UnsetCreation ensures that no value is present for Creation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


