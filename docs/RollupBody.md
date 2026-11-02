# RollupBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activity** | **[]int64** | Files added and findings recorded in each of the last 24 hours, oldest first. | 
**AgentStatus** | **NullableString** | State of the agents as a whole: queued while every run is pending, running while any is pending or running, complete otherwise. Null when no agent has run. | 
**Agents** | [**[]AgentBody**](AgentBody.md) | Agents run on the investigation&#39;s files that you may see, each with its state across the files. Empty until an agent has run. | 
**FileCount** | **int64** | Files in the investigation that you are permitted to see. | 
**Findings** | [**FindingCountsBody**](FindingCountsBody.md) |  | 
**FindingsSeverity** | **NullableString** | The highest severity among the findings, independent of the investigation&#39;s own severity. Null when there are no findings. | 
**FirstFileName** | **NullableString** | Name of the first file added that you are permitted to see. Null when there is none. | 
**LastActivityAt** | **time.Time** | The latest of an edit to the investigation, a file being added and a finding being recorded. | 

## Methods

### NewRollupBody

`func NewRollupBody(activity []int64, agentStatus NullableString, agents []AgentBody, fileCount int64, findings FindingCountsBody, findingsSeverity NullableString, firstFileName NullableString, lastActivityAt time.Time, ) *RollupBody`

NewRollupBody instantiates a new RollupBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRollupBodyWithDefaults

`func NewRollupBodyWithDefaults() *RollupBody`

NewRollupBodyWithDefaults instantiates a new RollupBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivity

`func (o *RollupBody) GetActivity() []int64`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *RollupBody) GetActivityOk() (*[]int64, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *RollupBody) SetActivity(v []int64)`

SetActivity sets Activity field to given value.


### GetAgentStatus

`func (o *RollupBody) GetAgentStatus() string`

GetAgentStatus returns the AgentStatus field if non-nil, zero value otherwise.

### GetAgentStatusOk

`func (o *RollupBody) GetAgentStatusOk() (*string, bool)`

GetAgentStatusOk returns a tuple with the AgentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentStatus

`func (o *RollupBody) SetAgentStatus(v string)`

SetAgentStatus sets AgentStatus field to given value.


### SetAgentStatusNil

`func (o *RollupBody) SetAgentStatusNil(b bool)`

 SetAgentStatusNil sets the value for AgentStatus to be an explicit nil

### UnsetAgentStatus
`func (o *RollupBody) UnsetAgentStatus()`

UnsetAgentStatus ensures that no value is present for AgentStatus, not even an explicit nil
### GetAgents

`func (o *RollupBody) GetAgents() []AgentBody`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *RollupBody) GetAgentsOk() (*[]AgentBody, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *RollupBody) SetAgents(v []AgentBody)`

SetAgents sets Agents field to given value.


### GetFileCount

`func (o *RollupBody) GetFileCount() int64`

GetFileCount returns the FileCount field if non-nil, zero value otherwise.

### GetFileCountOk

`func (o *RollupBody) GetFileCountOk() (*int64, bool)`

GetFileCountOk returns a tuple with the FileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCount

`func (o *RollupBody) SetFileCount(v int64)`

SetFileCount sets FileCount field to given value.


### GetFindings

`func (o *RollupBody) GetFindings() FindingCountsBody`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *RollupBody) GetFindingsOk() (*FindingCountsBody, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *RollupBody) SetFindings(v FindingCountsBody)`

SetFindings sets Findings field to given value.


### GetFindingsSeverity

`func (o *RollupBody) GetFindingsSeverity() string`

GetFindingsSeverity returns the FindingsSeverity field if non-nil, zero value otherwise.

### GetFindingsSeverityOk

`func (o *RollupBody) GetFindingsSeverityOk() (*string, bool)`

GetFindingsSeverityOk returns a tuple with the FindingsSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingsSeverity

`func (o *RollupBody) SetFindingsSeverity(v string)`

SetFindingsSeverity sets FindingsSeverity field to given value.


### SetFindingsSeverityNil

`func (o *RollupBody) SetFindingsSeverityNil(b bool)`

 SetFindingsSeverityNil sets the value for FindingsSeverity to be an explicit nil

### UnsetFindingsSeverity
`func (o *RollupBody) UnsetFindingsSeverity()`

UnsetFindingsSeverity ensures that no value is present for FindingsSeverity, not even an explicit nil
### GetFirstFileName

`func (o *RollupBody) GetFirstFileName() string`

GetFirstFileName returns the FirstFileName field if non-nil, zero value otherwise.

### GetFirstFileNameOk

`func (o *RollupBody) GetFirstFileNameOk() (*string, bool)`

GetFirstFileNameOk returns a tuple with the FirstFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstFileName

`func (o *RollupBody) SetFirstFileName(v string)`

SetFirstFileName sets FirstFileName field to given value.


### SetFirstFileNameNil

`func (o *RollupBody) SetFirstFileNameNil(b bool)`

 SetFirstFileNameNil sets the value for FirstFileName to be an explicit nil

### UnsetFirstFileName
`func (o *RollupBody) UnsetFirstFileName()`

UnsetFirstFileName ensures that no value is present for FirstFileName, not even an explicit nil
### GetLastActivityAt

`func (o *RollupBody) GetLastActivityAt() time.Time`

GetLastActivityAt returns the LastActivityAt field if non-nil, zero value otherwise.

### GetLastActivityAtOk

`func (o *RollupBody) GetLastActivityAtOk() (*time.Time, bool)`

GetLastActivityAtOk returns a tuple with the LastActivityAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityAt

`func (o *RollupBody) SetLastActivityAt(v time.Time)`

SetLastActivityAt sets LastActivityAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


