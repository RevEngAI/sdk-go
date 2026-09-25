# BinaryExternalsBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mb** | **interface{}** |  | 
**MbLastUpdated** | **NullableTime** | When the MalwareBazaar lookup last ran, null if never looked up | 
**Sha256Hash** | **string** | SHA-256 hash the lookups were keyed by | 
**Vt** | **interface{}** |  | 
**VtLastUpdated** | **time.Time** | When the VirusTotal lookup last ran | 

## Methods

### NewBinaryExternalsBody

`func NewBinaryExternalsBody(mb interface{}, mbLastUpdated NullableTime, sha256Hash string, vt interface{}, vtLastUpdated time.Time, ) *BinaryExternalsBody`

NewBinaryExternalsBody instantiates a new BinaryExternalsBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBinaryExternalsBodyWithDefaults

`func NewBinaryExternalsBodyWithDefaults() *BinaryExternalsBody`

NewBinaryExternalsBodyWithDefaults instantiates a new BinaryExternalsBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMb

`func (o *BinaryExternalsBody) GetMb() interface{}`

GetMb returns the Mb field if non-nil, zero value otherwise.

### GetMbOk

`func (o *BinaryExternalsBody) GetMbOk() (*interface{}, bool)`

GetMbOk returns a tuple with the Mb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMb

`func (o *BinaryExternalsBody) SetMb(v interface{})`

SetMb sets Mb field to given value.


### SetMbNil

`func (o *BinaryExternalsBody) SetMbNil(b bool)`

 SetMbNil sets the value for Mb to be an explicit nil

### UnsetMb
`func (o *BinaryExternalsBody) UnsetMb()`

UnsetMb ensures that no value is present for Mb, not even an explicit nil
### GetMbLastUpdated

`func (o *BinaryExternalsBody) GetMbLastUpdated() time.Time`

GetMbLastUpdated returns the MbLastUpdated field if non-nil, zero value otherwise.

### GetMbLastUpdatedOk

`func (o *BinaryExternalsBody) GetMbLastUpdatedOk() (*time.Time, bool)`

GetMbLastUpdatedOk returns a tuple with the MbLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbLastUpdated

`func (o *BinaryExternalsBody) SetMbLastUpdated(v time.Time)`

SetMbLastUpdated sets MbLastUpdated field to given value.


### SetMbLastUpdatedNil

`func (o *BinaryExternalsBody) SetMbLastUpdatedNil(b bool)`

 SetMbLastUpdatedNil sets the value for MbLastUpdated to be an explicit nil

### UnsetMbLastUpdated
`func (o *BinaryExternalsBody) UnsetMbLastUpdated()`

UnsetMbLastUpdated ensures that no value is present for MbLastUpdated, not even an explicit nil
### GetSha256Hash

`func (o *BinaryExternalsBody) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *BinaryExternalsBody) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *BinaryExternalsBody) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.


### GetVt

`func (o *BinaryExternalsBody) GetVt() interface{}`

GetVt returns the Vt field if non-nil, zero value otherwise.

### GetVtOk

`func (o *BinaryExternalsBody) GetVtOk() (*interface{}, bool)`

GetVtOk returns a tuple with the Vt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVt

`func (o *BinaryExternalsBody) SetVt(v interface{})`

SetVt sets Vt field to given value.


### SetVtNil

`func (o *BinaryExternalsBody) SetVtNil(b bool)`

 SetVtNil sets the value for Vt to be an explicit nil

### UnsetVt
`func (o *BinaryExternalsBody) UnsetVt()`

UnsetVt ensures that no value is present for Vt, not even an explicit nil
### GetVtLastUpdated

`func (o *BinaryExternalsBody) GetVtLastUpdated() time.Time`

GetVtLastUpdated returns the VtLastUpdated field if non-nil, zero value otherwise.

### GetVtLastUpdatedOk

`func (o *BinaryExternalsBody) GetVtLastUpdatedOk() (*time.Time, bool)`

GetVtLastUpdatedOk returns a tuple with the VtLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtLastUpdated

`func (o *BinaryExternalsBody) SetVtLastUpdated(v time.Time)`

SetVtLastUpdated sets VtLastUpdated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


