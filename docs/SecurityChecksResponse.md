# SecurityChecksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinaryId** | **int32** |  | 
**TotalResults** | **int32** |  | 
**Results** | [**[]SecurityChecksResult**](SecurityChecksResult.md) |  | 

## Methods

### NewSecurityChecksResponse

`func NewSecurityChecksResponse(binaryId int32, totalResults int32, results []SecurityChecksResult, ) *SecurityChecksResponse`

NewSecurityChecksResponse instantiates a new SecurityChecksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityChecksResponseWithDefaults

`func NewSecurityChecksResponseWithDefaults() *SecurityChecksResponse`

NewSecurityChecksResponseWithDefaults instantiates a new SecurityChecksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaryId

`func (o *SecurityChecksResponse) GetBinaryId() int32`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *SecurityChecksResponse) GetBinaryIdOk() (*int32, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *SecurityChecksResponse) SetBinaryId(v int32)`

SetBinaryId sets BinaryId field to given value.


### GetTotalResults

`func (o *SecurityChecksResponse) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *SecurityChecksResponse) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *SecurityChecksResponse) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.


### GetResults

`func (o *SecurityChecksResponse) GetResults() []SecurityChecksResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SecurityChecksResponse) GetResultsOk() (*[]SecurityChecksResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SecurityChecksResponse) SetResults(v []SecurityChecksResult)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


