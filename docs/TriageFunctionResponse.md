# TriageFunctionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the function | 
**Address** | **int32** | Address of the function in the binary | 
**Summary** | **string** | Summary of the function&#39;s behaviour | 
**Score** | **float32** | Score indicating the function&#39;s relevance | 
**Capabilities** | **[]string** | List of capabilities exhibited by the function | 

## Methods

### NewTriageFunctionResponse

`func NewTriageFunctionResponse(id int32, address int32, summary string, score float32, capabilities []string, ) *TriageFunctionResponse`

NewTriageFunctionResponse instantiates a new TriageFunctionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriageFunctionResponseWithDefaults

`func NewTriageFunctionResponseWithDefaults() *TriageFunctionResponse`

NewTriageFunctionResponseWithDefaults instantiates a new TriageFunctionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TriageFunctionResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TriageFunctionResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TriageFunctionResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetAddress

`func (o *TriageFunctionResponse) GetAddress() int32`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TriageFunctionResponse) GetAddressOk() (*int32, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TriageFunctionResponse) SetAddress(v int32)`

SetAddress sets Address field to given value.


### GetSummary

`func (o *TriageFunctionResponse) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *TriageFunctionResponse) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *TriageFunctionResponse) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetScore

`func (o *TriageFunctionResponse) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *TriageFunctionResponse) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *TriageFunctionResponse) SetScore(v float32)`

SetScore sets Score field to given value.


### GetCapabilities

`func (o *TriageFunctionResponse) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *TriageFunctionResponse) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *TriageFunctionResponse) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


