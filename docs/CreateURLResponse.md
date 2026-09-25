# CreateURLResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisId** | **int64** |  | 
**BinaryId** | **int64** |  | 
**Sha256Hash** | **string** |  | 

## Methods

### NewCreateURLResponse

`func NewCreateURLResponse(analysisId int64, binaryId int64, sha256Hash string, ) *CreateURLResponse`

NewCreateURLResponse instantiates a new CreateURLResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateURLResponseWithDefaults

`func NewCreateURLResponseWithDefaults() *CreateURLResponse`

NewCreateURLResponseWithDefaults instantiates a new CreateURLResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisId

`func (o *CreateURLResponse) GetAnalysisId() int64`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *CreateURLResponse) GetAnalysisIdOk() (*int64, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *CreateURLResponse) SetAnalysisId(v int64)`

SetAnalysisId sets AnalysisId field to given value.


### GetBinaryId

`func (o *CreateURLResponse) GetBinaryId() int64`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *CreateURLResponse) GetBinaryIdOk() (*int64, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *CreateURLResponse) SetBinaryId(v int64)`

SetBinaryId sets BinaryId field to given value.


### GetSha256Hash

`func (o *CreateURLResponse) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *CreateURLResponse) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *CreateURLResponse) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


