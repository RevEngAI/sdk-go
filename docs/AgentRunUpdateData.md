# AgentRunUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | [**AgentRunUpdateAgent**](AgentRunUpdateAgent.md) |  | 
**AnalysisId** | **int64** | The earliest-added analysis of the file in the investigation that the conversation&#39;s starter may see. | 
**BinaryName** | **string** |  | 
**Card** | Pointer to [**AgentRunUpdateCard**](AgentRunUpdateCard.md) | A summary of the agent&#39;s result. Absent until the run is done, and when its result cannot be summarised. | [optional] 
**CreatedAt** | **time.Time** |  | 
**Error** | **NullableString** | Why the run failed. Null unless it failed with a reason. | 
**InvestigationId** | **int64** |  | 
**TaskId** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewAgentRunUpdateData

`func NewAgentRunUpdateData(agent AgentRunUpdateAgent, analysisId int64, binaryName string, createdAt time.Time, error_ NullableString, investigationId int64, taskId string, updatedAt time.Time, ) *AgentRunUpdateData`

NewAgentRunUpdateData instantiates a new AgentRunUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentRunUpdateDataWithDefaults

`func NewAgentRunUpdateDataWithDefaults() *AgentRunUpdateData`

NewAgentRunUpdateDataWithDefaults instantiates a new AgentRunUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *AgentRunUpdateData) GetAgent() AgentRunUpdateAgent`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *AgentRunUpdateData) GetAgentOk() (*AgentRunUpdateAgent, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *AgentRunUpdateData) SetAgent(v AgentRunUpdateAgent)`

SetAgent sets Agent field to given value.


### GetAnalysisId

`func (o *AgentRunUpdateData) GetAnalysisId() int64`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *AgentRunUpdateData) GetAnalysisIdOk() (*int64, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *AgentRunUpdateData) SetAnalysisId(v int64)`

SetAnalysisId sets AnalysisId field to given value.


### GetBinaryName

`func (o *AgentRunUpdateData) GetBinaryName() string`

GetBinaryName returns the BinaryName field if non-nil, zero value otherwise.

### GetBinaryNameOk

`func (o *AgentRunUpdateData) GetBinaryNameOk() (*string, bool)`

GetBinaryNameOk returns a tuple with the BinaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryName

`func (o *AgentRunUpdateData) SetBinaryName(v string)`

SetBinaryName sets BinaryName field to given value.


### GetCard

`func (o *AgentRunUpdateData) GetCard() AgentRunUpdateCard`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *AgentRunUpdateData) GetCardOk() (*AgentRunUpdateCard, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *AgentRunUpdateData) SetCard(v AgentRunUpdateCard)`

SetCard sets Card field to given value.

### HasCard

`func (o *AgentRunUpdateData) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AgentRunUpdateData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentRunUpdateData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentRunUpdateData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetError

`func (o *AgentRunUpdateData) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AgentRunUpdateData) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AgentRunUpdateData) SetError(v string)`

SetError sets Error field to given value.


### SetErrorNil

`func (o *AgentRunUpdateData) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *AgentRunUpdateData) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetInvestigationId

`func (o *AgentRunUpdateData) GetInvestigationId() int64`

GetInvestigationId returns the InvestigationId field if non-nil, zero value otherwise.

### GetInvestigationIdOk

`func (o *AgentRunUpdateData) GetInvestigationIdOk() (*int64, bool)`

GetInvestigationIdOk returns a tuple with the InvestigationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigationId

`func (o *AgentRunUpdateData) SetInvestigationId(v int64)`

SetInvestigationId sets InvestigationId field to given value.


### GetTaskId

`func (o *AgentRunUpdateData) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *AgentRunUpdateData) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *AgentRunUpdateData) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetUpdatedAt

`func (o *AgentRunUpdateData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AgentRunUpdateData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AgentRunUpdateData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


