# AgentStatusCountsBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentsRunning** | **int64** |  | 
**Complete** | **int64** |  | 
**Queued** | **int64** |  | 

## Methods

### NewAgentStatusCountsBody

`func NewAgentStatusCountsBody(agentsRunning int64, complete int64, queued int64, ) *AgentStatusCountsBody`

NewAgentStatusCountsBody instantiates a new AgentStatusCountsBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentStatusCountsBodyWithDefaults

`func NewAgentStatusCountsBodyWithDefaults() *AgentStatusCountsBody`

NewAgentStatusCountsBodyWithDefaults instantiates a new AgentStatusCountsBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentsRunning

`func (o *AgentStatusCountsBody) GetAgentsRunning() int64`

GetAgentsRunning returns the AgentsRunning field if non-nil, zero value otherwise.

### GetAgentsRunningOk

`func (o *AgentStatusCountsBody) GetAgentsRunningOk() (*int64, bool)`

GetAgentsRunningOk returns a tuple with the AgentsRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentsRunning

`func (o *AgentStatusCountsBody) SetAgentsRunning(v int64)`

SetAgentsRunning sets AgentsRunning field to given value.


### GetComplete

`func (o *AgentStatusCountsBody) GetComplete() int64`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *AgentStatusCountsBody) GetCompleteOk() (*int64, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *AgentStatusCountsBody) SetComplete(v int64)`

SetComplete sets Complete field to given value.


### GetQueued

`func (o *AgentStatusCountsBody) GetQueued() int64`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *AgentStatusCountsBody) GetQueuedOk() (*int64, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *AgentStatusCountsBody) SetQueued(v int64)`

SetQueued sets Queued field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


