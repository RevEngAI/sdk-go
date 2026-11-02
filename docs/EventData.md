# EventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversationUuid** | Pointer to **string** |  | [optional] 
**DefinitionName** | Pointer to **string** |  | [optional] 
**InputTokens** | Pointer to **int32** |  | [optional] 
**OutputTokens** | Pointer to **int32** |  | [optional] 
**TotalTokens** | Pointer to **int32** |  | [optional] 
**Error** | **string** | Why the run failed. Null unless it failed with a reason. | 
**StepIndex** | Pointer to **int32** |  | [optional] 
**FinishReason** | Pointer to **string** |  | [optional] 
**MessageId** | Pointer to **string** |  | [optional] 
**Delta** | Pointer to **string** |  | [optional] 
**ToolArgs** | Pointer to **string** |  | [optional] 
**ToolCallId** | Pointer to **string** |  | [optional] 
**ToolName** | Pointer to **string** |  | [optional] 
**IsError** | Pointer to **bool** |  | [optional] 
**Result** | Pointer to **string** |  | [optional] 
**Updated** | Pointer to [**[]EntityUpdate**](EntityUpdate.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**MessagesAfter** | Pointer to **int32** |  | [optional] 
**MessagesBefore** | Pointer to **int32** |  | [optional] 
**RemovedTurns** | Pointer to **int32** |  | [optional] 
**Strategy** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Agent** | [**AgentRunUpdateAgent**](AgentRunUpdateAgent.md) |  | 
**AnalysisId** | **int64** | The earliest-added analysis of the file in the investigation that the conversation&#39;s starter may see. | 
**BinaryName** | **string** |  | 
**Card** | Pointer to [**AgentRunUpdateCard**](AgentRunUpdateCard.md) | A summary of the agent&#39;s result. Absent until the run is done, and when its result cannot be summarised. | [optional] 
**CreatedAt** | **time.Time** |  | 
**InvestigationId** | **int64** |  | 
**TaskId** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewEventData

`func NewEventData(error_ string, agent AgentRunUpdateAgent, analysisId int64, binaryName string, createdAt time.Time, investigationId int64, taskId string, updatedAt time.Time, ) *EventData`

NewEventData instantiates a new EventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventDataWithDefaults

`func NewEventDataWithDefaults() *EventData`

NewEventDataWithDefaults instantiates a new EventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversationUuid

`func (o *EventData) GetConversationUuid() string`

GetConversationUuid returns the ConversationUuid field if non-nil, zero value otherwise.

### GetConversationUuidOk

`func (o *EventData) GetConversationUuidOk() (*string, bool)`

GetConversationUuidOk returns a tuple with the ConversationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationUuid

`func (o *EventData) SetConversationUuid(v string)`

SetConversationUuid sets ConversationUuid field to given value.

### HasConversationUuid

`func (o *EventData) HasConversationUuid() bool`

HasConversationUuid returns a boolean if a field has been set.

### GetDefinitionName

`func (o *EventData) GetDefinitionName() string`

GetDefinitionName returns the DefinitionName field if non-nil, zero value otherwise.

### GetDefinitionNameOk

`func (o *EventData) GetDefinitionNameOk() (*string, bool)`

GetDefinitionNameOk returns a tuple with the DefinitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionName

`func (o *EventData) SetDefinitionName(v string)`

SetDefinitionName sets DefinitionName field to given value.

### HasDefinitionName

`func (o *EventData) HasDefinitionName() bool`

HasDefinitionName returns a boolean if a field has been set.

### GetInputTokens

`func (o *EventData) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *EventData) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *EventData) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *EventData) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### GetOutputTokens

`func (o *EventData) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *EventData) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *EventData) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.

### HasOutputTokens

`func (o *EventData) HasOutputTokens() bool`

HasOutputTokens returns a boolean if a field has been set.

### GetTotalTokens

`func (o *EventData) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *EventData) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *EventData) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *EventData) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.

### GetError

`func (o *EventData) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *EventData) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *EventData) SetError(v string)`

SetError sets Error field to given value.


### GetStepIndex

`func (o *EventData) GetStepIndex() int32`

GetStepIndex returns the StepIndex field if non-nil, zero value otherwise.

### GetStepIndexOk

`func (o *EventData) GetStepIndexOk() (*int32, bool)`

GetStepIndexOk returns a tuple with the StepIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepIndex

`func (o *EventData) SetStepIndex(v int32)`

SetStepIndex sets StepIndex field to given value.

### HasStepIndex

`func (o *EventData) HasStepIndex() bool`

HasStepIndex returns a boolean if a field has been set.

### GetFinishReason

`func (o *EventData) GetFinishReason() string`

GetFinishReason returns the FinishReason field if non-nil, zero value otherwise.

### GetFinishReasonOk

`func (o *EventData) GetFinishReasonOk() (*string, bool)`

GetFinishReasonOk returns a tuple with the FinishReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishReason

`func (o *EventData) SetFinishReason(v string)`

SetFinishReason sets FinishReason field to given value.

### HasFinishReason

`func (o *EventData) HasFinishReason() bool`

HasFinishReason returns a boolean if a field has been set.

### GetMessageId

`func (o *EventData) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *EventData) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *EventData) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *EventData) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetDelta

`func (o *EventData) GetDelta() string`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *EventData) GetDeltaOk() (*string, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *EventData) SetDelta(v string)`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *EventData) HasDelta() bool`

HasDelta returns a boolean if a field has been set.

### GetToolArgs

`func (o *EventData) GetToolArgs() string`

GetToolArgs returns the ToolArgs field if non-nil, zero value otherwise.

### GetToolArgsOk

`func (o *EventData) GetToolArgsOk() (*string, bool)`

GetToolArgsOk returns a tuple with the ToolArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolArgs

`func (o *EventData) SetToolArgs(v string)`

SetToolArgs sets ToolArgs field to given value.

### HasToolArgs

`func (o *EventData) HasToolArgs() bool`

HasToolArgs returns a boolean if a field has been set.

### GetToolCallId

`func (o *EventData) GetToolCallId() string`

GetToolCallId returns the ToolCallId field if non-nil, zero value otherwise.

### GetToolCallIdOk

`func (o *EventData) GetToolCallIdOk() (*string, bool)`

GetToolCallIdOk returns a tuple with the ToolCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCallId

`func (o *EventData) SetToolCallId(v string)`

SetToolCallId sets ToolCallId field to given value.

### HasToolCallId

`func (o *EventData) HasToolCallId() bool`

HasToolCallId returns a boolean if a field has been set.

### GetToolName

`func (o *EventData) GetToolName() string`

GetToolName returns the ToolName field if non-nil, zero value otherwise.

### GetToolNameOk

`func (o *EventData) GetToolNameOk() (*string, bool)`

GetToolNameOk returns a tuple with the ToolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolName

`func (o *EventData) SetToolName(v string)`

SetToolName sets ToolName field to given value.

### HasToolName

`func (o *EventData) HasToolName() bool`

HasToolName returns a boolean if a field has been set.

### GetIsError

`func (o *EventData) GetIsError() bool`

GetIsError returns the IsError field if non-nil, zero value otherwise.

### GetIsErrorOk

`func (o *EventData) GetIsErrorOk() (*bool, bool)`

GetIsErrorOk returns a tuple with the IsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsError

`func (o *EventData) SetIsError(v bool)`

SetIsError sets IsError field to given value.

### HasIsError

`func (o *EventData) HasIsError() bool`

HasIsError returns a boolean if a field has been set.

### GetResult

`func (o *EventData) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *EventData) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *EventData) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *EventData) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetUpdated

`func (o *EventData) GetUpdated() []EntityUpdate`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *EventData) GetUpdatedOk() (*[]EntityUpdate, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *EventData) SetUpdated(v []EntityUpdate)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *EventData) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetTitle

`func (o *EventData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EventData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EventData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EventData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMessagesAfter

`func (o *EventData) GetMessagesAfter() int32`

GetMessagesAfter returns the MessagesAfter field if non-nil, zero value otherwise.

### GetMessagesAfterOk

`func (o *EventData) GetMessagesAfterOk() (*int32, bool)`

GetMessagesAfterOk returns a tuple with the MessagesAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesAfter

`func (o *EventData) SetMessagesAfter(v int32)`

SetMessagesAfter sets MessagesAfter field to given value.

### HasMessagesAfter

`func (o *EventData) HasMessagesAfter() bool`

HasMessagesAfter returns a boolean if a field has been set.

### GetMessagesBefore

`func (o *EventData) GetMessagesBefore() int32`

GetMessagesBefore returns the MessagesBefore field if non-nil, zero value otherwise.

### GetMessagesBeforeOk

`func (o *EventData) GetMessagesBeforeOk() (*int32, bool)`

GetMessagesBeforeOk returns a tuple with the MessagesBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesBefore

`func (o *EventData) SetMessagesBefore(v int32)`

SetMessagesBefore sets MessagesBefore field to given value.

### HasMessagesBefore

`func (o *EventData) HasMessagesBefore() bool`

HasMessagesBefore returns a boolean if a field has been set.

### GetRemovedTurns

`func (o *EventData) GetRemovedTurns() int32`

GetRemovedTurns returns the RemovedTurns field if non-nil, zero value otherwise.

### GetRemovedTurnsOk

`func (o *EventData) GetRemovedTurnsOk() (*int32, bool)`

GetRemovedTurnsOk returns a tuple with the RemovedTurns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedTurns

`func (o *EventData) SetRemovedTurns(v int32)`

SetRemovedTurns sets RemovedTurns field to given value.

### HasRemovedTurns

`func (o *EventData) HasRemovedTurns() bool`

HasRemovedTurns returns a boolean if a field has been set.

### GetStrategy

`func (o *EventData) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *EventData) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *EventData) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *EventData) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetMessage

`func (o *EventData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EventData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EventData) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EventData) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetAgent

`func (o *EventData) GetAgent() AgentRunUpdateAgent`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *EventData) GetAgentOk() (*AgentRunUpdateAgent, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *EventData) SetAgent(v AgentRunUpdateAgent)`

SetAgent sets Agent field to given value.


### GetAnalysisId

`func (o *EventData) GetAnalysisId() int64`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *EventData) GetAnalysisIdOk() (*int64, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *EventData) SetAnalysisId(v int64)`

SetAnalysisId sets AnalysisId field to given value.


### GetBinaryName

`func (o *EventData) GetBinaryName() string`

GetBinaryName returns the BinaryName field if non-nil, zero value otherwise.

### GetBinaryNameOk

`func (o *EventData) GetBinaryNameOk() (*string, bool)`

GetBinaryNameOk returns a tuple with the BinaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryName

`func (o *EventData) SetBinaryName(v string)`

SetBinaryName sets BinaryName field to given value.


### GetCard

`func (o *EventData) GetCard() AgentRunUpdateCard`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *EventData) GetCardOk() (*AgentRunUpdateCard, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *EventData) SetCard(v AgentRunUpdateCard)`

SetCard sets Card field to given value.

### HasCard

`func (o *EventData) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EventData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EventData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EventData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetInvestigationId

`func (o *EventData) GetInvestigationId() int64`

GetInvestigationId returns the InvestigationId field if non-nil, zero value otherwise.

### GetInvestigationIdOk

`func (o *EventData) GetInvestigationIdOk() (*int64, bool)`

GetInvestigationIdOk returns a tuple with the InvestigationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigationId

`func (o *EventData) SetInvestigationId(v int64)`

SetInvestigationId sets InvestigationId field to given value.


### GetTaskId

`func (o *EventData) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *EventData) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *EventData) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetUpdatedAt

`func (o *EventData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EventData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EventData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


