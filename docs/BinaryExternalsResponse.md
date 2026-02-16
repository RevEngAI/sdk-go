# BinaryExternalsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sha256Hash** | **string** | SHA256 hash of the binary | 
**Vt** | **map[string]interface{}** | VirusTotal information | 
**VtLastUpdated** | **time.Time** | VirusTotal last updated date | 
**Mb** | **map[string]interface{}** | MalwareBazaar information | 
**MbLastUpdated** | **time.Time** | MalwareBazaar last updated date | 

## Methods

### NewBinaryExternalsResponse

`func NewBinaryExternalsResponse(sha256Hash string, vt map[string]interface{}, vtLastUpdated time.Time, mb map[string]interface{}, mbLastUpdated time.Time, ) *BinaryExternalsResponse`

NewBinaryExternalsResponse instantiates a new BinaryExternalsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBinaryExternalsResponseWithDefaults

`func NewBinaryExternalsResponseWithDefaults() *BinaryExternalsResponse`

NewBinaryExternalsResponseWithDefaults instantiates a new BinaryExternalsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSha256Hash

`func (o *BinaryExternalsResponse) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *BinaryExternalsResponse) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *BinaryExternalsResponse) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.


### GetVt

`func (o *BinaryExternalsResponse) GetVt() map[string]interface{}`

GetVt returns the Vt field if non-nil, zero value otherwise.

### GetVtOk

`func (o *BinaryExternalsResponse) GetVtOk() (*map[string]interface{}, bool)`

GetVtOk returns a tuple with the Vt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVt

`func (o *BinaryExternalsResponse) SetVt(v map[string]interface{})`

SetVt sets Vt field to given value.


### GetVtLastUpdated

`func (o *BinaryExternalsResponse) GetVtLastUpdated() time.Time`

GetVtLastUpdated returns the VtLastUpdated field if non-nil, zero value otherwise.

### GetVtLastUpdatedOk

`func (o *BinaryExternalsResponse) GetVtLastUpdatedOk() (*time.Time, bool)`

GetVtLastUpdatedOk returns a tuple with the VtLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtLastUpdated

`func (o *BinaryExternalsResponse) SetVtLastUpdated(v time.Time)`

SetVtLastUpdated sets VtLastUpdated field to given value.


### GetMb

`func (o *BinaryExternalsResponse) GetMb() map[string]interface{}`

GetMb returns the Mb field if non-nil, zero value otherwise.

### GetMbOk

`func (o *BinaryExternalsResponse) GetMbOk() (*map[string]interface{}, bool)`

GetMbOk returns a tuple with the Mb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMb

`func (o *BinaryExternalsResponse) SetMb(v map[string]interface{})`

SetMb sets Mb field to given value.


### GetMbLastUpdated

`func (o *BinaryExternalsResponse) GetMbLastUpdated() time.Time`

GetMbLastUpdated returns the MbLastUpdated field if non-nil, zero value otherwise.

### GetMbLastUpdatedOk

`func (o *BinaryExternalsResponse) GetMbLastUpdatedOk() (*time.Time, bool)`

GetMbLastUpdatedOk returns a tuple with the MbLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbLastUpdated

`func (o *BinaryExternalsResponse) SetMbLastUpdated(v time.Time)`

SetMbLastUpdated sets MbLastUpdated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


