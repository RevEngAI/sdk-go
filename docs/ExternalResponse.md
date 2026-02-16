# ExternalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sha256Hash** | **string** |  | 
**Data** | **map[string]interface{}** |  | 
**LastUpdated** | **time.Time** |  | 

## Methods

### NewExternalResponse

`func NewExternalResponse(sha256Hash string, data map[string]interface{}, lastUpdated time.Time, ) *ExternalResponse`

NewExternalResponse instantiates a new ExternalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalResponseWithDefaults

`func NewExternalResponseWithDefaults() *ExternalResponse`

NewExternalResponseWithDefaults instantiates a new ExternalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSha256Hash

`func (o *ExternalResponse) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *ExternalResponse) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *ExternalResponse) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.


### GetData

`func (o *ExternalResponse) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExternalResponse) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExternalResponse) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetLastUpdated

`func (o *ExternalResponse) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ExternalResponse) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ExternalResponse) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


