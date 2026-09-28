# VirusTotalScanResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sha256Hash** | **string** | Content hash the lookup was keyed by | 
**Vt** | **interface{}** |  | 
**VtLastUpdated** | **time.Time** | When this result was recorded | 

## Methods

### NewVirusTotalScanResult

`func NewVirusTotalScanResult(sha256Hash string, vt interface{}, vtLastUpdated time.Time, ) *VirusTotalScanResult`

NewVirusTotalScanResult instantiates a new VirusTotalScanResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirusTotalScanResultWithDefaults

`func NewVirusTotalScanResultWithDefaults() *VirusTotalScanResult`

NewVirusTotalScanResultWithDefaults instantiates a new VirusTotalScanResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSha256Hash

`func (o *VirusTotalScanResult) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *VirusTotalScanResult) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *VirusTotalScanResult) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.


### GetVt

`func (o *VirusTotalScanResult) GetVt() interface{}`

GetVt returns the Vt field if non-nil, zero value otherwise.

### GetVtOk

`func (o *VirusTotalScanResult) GetVtOk() (*interface{}, bool)`

GetVtOk returns a tuple with the Vt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVt

`func (o *VirusTotalScanResult) SetVt(v interface{})`

SetVt sets Vt field to given value.


### SetVtNil

`func (o *VirusTotalScanResult) SetVtNil(b bool)`

 SetVtNil sets the value for Vt to be an explicit nil

### UnsetVt
`func (o *VirusTotalScanResult) UnsetVt()`

UnsetVt ensures that no value is present for Vt, not even an explicit nil
### GetVtLastUpdated

`func (o *VirusTotalScanResult) GetVtLastUpdated() time.Time`

GetVtLastUpdated returns the VtLastUpdated field if non-nil, zero value otherwise.

### GetVtLastUpdatedOk

`func (o *VirusTotalScanResult) GetVtLastUpdatedOk() (*time.Time, bool)`

GetVtLastUpdatedOk returns a tuple with the VtLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtLastUpdated

`func (o *VirusTotalScanResult) SetVtLastUpdated(v time.Time)`

SetVtLastUpdated sets VtLastUpdated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


