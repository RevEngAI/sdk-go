# SizeBucketBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BelowBytes** | **NullableInt64** | Exclusive upper bound of the bucket, in bytes. Null for the last bucket, which has no upper bound. | 
**Functions** | **int64** | Functions a binary this size is estimated to have | 
**Uncertain** | **bool** | True where too few operations have run on binaries this size for the estimate to be reliable | 

## Methods

### NewSizeBucketBody

`func NewSizeBucketBody(belowBytes NullableInt64, functions int64, uncertain bool, ) *SizeBucketBody`

NewSizeBucketBody instantiates a new SizeBucketBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSizeBucketBodyWithDefaults

`func NewSizeBucketBodyWithDefaults() *SizeBucketBody`

NewSizeBucketBodyWithDefaults instantiates a new SizeBucketBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBelowBytes

`func (o *SizeBucketBody) GetBelowBytes() int64`

GetBelowBytes returns the BelowBytes field if non-nil, zero value otherwise.

### GetBelowBytesOk

`func (o *SizeBucketBody) GetBelowBytesOk() (*int64, bool)`

GetBelowBytesOk returns a tuple with the BelowBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBelowBytes

`func (o *SizeBucketBody) SetBelowBytes(v int64)`

SetBelowBytes sets BelowBytes field to given value.


### SetBelowBytesNil

`func (o *SizeBucketBody) SetBelowBytesNil(b bool)`

 SetBelowBytesNil sets the value for BelowBytes to be an explicit nil

### UnsetBelowBytes
`func (o *SizeBucketBody) UnsetBelowBytes()`

UnsetBelowBytes ensures that no value is present for BelowBytes, not even an explicit nil
### GetFunctions

`func (o *SizeBucketBody) GetFunctions() int64`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *SizeBucketBody) GetFunctionsOk() (*int64, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *SizeBucketBody) SetFunctions(v int64)`

SetFunctions sets Functions field to given value.


### GetUncertain

`func (o *SizeBucketBody) GetUncertain() bool`

GetUncertain returns the Uncertain field if non-nil, zero value otherwise.

### GetUncertainOk

`func (o *SizeBucketBody) GetUncertainOk() (*bool, bool)`

GetUncertainOk returns a tuple with the Uncertain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncertain

`func (o *SizeBucketBody) SetUncertain(v bool)`

SetUncertain sets Uncertain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


