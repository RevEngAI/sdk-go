# TriageFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **int64** | Virtual address of the function | 
**Capabilities** | **[]string** | Capability categories the function exhibits | 
**Id** | **int64** | ID of the function | 
**Score** | **float64** | Maliciousness score for the function, 0 to 1 | 
**Summary** | **string** | What the function does | 

## Methods

### NewTriageFunction

`func NewTriageFunction(address int64, capabilities []string, id int64, score float64, summary string, ) *TriageFunction`

NewTriageFunction instantiates a new TriageFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriageFunctionWithDefaults

`func NewTriageFunctionWithDefaults() *TriageFunction`

NewTriageFunctionWithDefaults instantiates a new TriageFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *TriageFunction) GetAddress() int64`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TriageFunction) GetAddressOk() (*int64, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TriageFunction) SetAddress(v int64)`

SetAddress sets Address field to given value.


### GetCapabilities

`func (o *TriageFunction) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *TriageFunction) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *TriageFunction) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.


### SetCapabilitiesNil

`func (o *TriageFunction) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *TriageFunction) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetId

`func (o *TriageFunction) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TriageFunction) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TriageFunction) SetId(v int64)`

SetId sets Id field to given value.


### GetScore

`func (o *TriageFunction) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *TriageFunction) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *TriageFunction) SetScore(v float64)`

SetScore sets Score field to given value.


### GetSummary

`func (o *TriageFunction) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *TriageFunction) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *TriageFunction) SetSummary(v string)`

SetSummary sets Summary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


