# AgentRunUpdateCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iocs** | [**[]AgentRunUpdateIOC**](AgentRunUpdateIOC.md) |  | 
**Score** | **NullableFloat64** | The agent&#39;s suspicion score from 0 to 1. Null when the agent gives none. | 
**Summary** | **string** |  | 

## Methods

### NewAgentRunUpdateCard

`func NewAgentRunUpdateCard(iocs []AgentRunUpdateIOC, score NullableFloat64, summary string, ) *AgentRunUpdateCard`

NewAgentRunUpdateCard instantiates a new AgentRunUpdateCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentRunUpdateCardWithDefaults

`func NewAgentRunUpdateCardWithDefaults() *AgentRunUpdateCard`

NewAgentRunUpdateCardWithDefaults instantiates a new AgentRunUpdateCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIocs

`func (o *AgentRunUpdateCard) GetIocs() []AgentRunUpdateIOC`

GetIocs returns the Iocs field if non-nil, zero value otherwise.

### GetIocsOk

`func (o *AgentRunUpdateCard) GetIocsOk() (*[]AgentRunUpdateIOC, bool)`

GetIocsOk returns a tuple with the Iocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIocs

`func (o *AgentRunUpdateCard) SetIocs(v []AgentRunUpdateIOC)`

SetIocs sets Iocs field to given value.


### GetScore

`func (o *AgentRunUpdateCard) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *AgentRunUpdateCard) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *AgentRunUpdateCard) SetScore(v float64)`

SetScore sets Score field to given value.


### SetScoreNil

`func (o *AgentRunUpdateCard) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *AgentRunUpdateCard) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetSummary

`func (o *AgentRunUpdateCard) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AgentRunUpdateCard) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AgentRunUpdateCard) SetSummary(v string)`

SetSummary sets Summary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


