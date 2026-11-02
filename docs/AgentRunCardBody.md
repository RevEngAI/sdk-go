# AgentRunCardBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iocs** | [**[]AgentRunIOCBody**](AgentRunIOCBody.md) |  | 
**Score** | **NullableFloat64** | The agent&#39;s suspicion score from 0 to 1. Null when the agent gives none. | 
**Summary** | **string** |  | 

## Methods

### NewAgentRunCardBody

`func NewAgentRunCardBody(iocs []AgentRunIOCBody, score NullableFloat64, summary string, ) *AgentRunCardBody`

NewAgentRunCardBody instantiates a new AgentRunCardBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentRunCardBodyWithDefaults

`func NewAgentRunCardBodyWithDefaults() *AgentRunCardBody`

NewAgentRunCardBodyWithDefaults instantiates a new AgentRunCardBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIocs

`func (o *AgentRunCardBody) GetIocs() []AgentRunIOCBody`

GetIocs returns the Iocs field if non-nil, zero value otherwise.

### GetIocsOk

`func (o *AgentRunCardBody) GetIocsOk() (*[]AgentRunIOCBody, bool)`

GetIocsOk returns a tuple with the Iocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIocs

`func (o *AgentRunCardBody) SetIocs(v []AgentRunIOCBody)`

SetIocs sets Iocs field to given value.


### GetScore

`func (o *AgentRunCardBody) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *AgentRunCardBody) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *AgentRunCardBody) SetScore(v float64)`

SetScore sets Score field to given value.


### SetScoreNil

`func (o *AgentRunCardBody) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *AgentRunCardBody) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetSummary

`func (o *AgentRunCardBody) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AgentRunCardBody) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AgentRunCardBody) SetSummary(v string)`

SetSummary sets Summary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


