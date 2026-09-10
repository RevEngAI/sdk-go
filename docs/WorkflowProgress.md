# WorkflowProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]ProgressMessage**](ProgressMessage.md) | Log messages emitted during execution | 
**Percent** | **int64** | Overall completion as a percentage, weighted by step duration | 
**Status** | **string** | Current workflow status | 
**Step** | **string** | Name of the current step | 
**StepIndex** | **int64** | Zero-based index of the current step | 
**StepShare** | **int64** | Percentage points the current step contributes when it completes | 
**StepsTotal** | **int64** | Total number of steps in the workflow | 
**SubStep** | Pointer to **string** | Phase within the current step, when the step reports one | [optional] 
**SubStepDone** | Pointer to **int64** | Items completed in the current phase | [optional] 
**SubStepTotal** | Pointer to **int64** | Items the current phase will process, 0 when unknown | [optional] 

## Methods

### NewWorkflowProgress

`func NewWorkflowProgress(messages []ProgressMessage, percent int64, status string, step string, stepIndex int64, stepShare int64, stepsTotal int64, ) *WorkflowProgress`

NewWorkflowProgress instantiates a new WorkflowProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowProgressWithDefaults

`func NewWorkflowProgressWithDefaults() *WorkflowProgress`

NewWorkflowProgressWithDefaults instantiates a new WorkflowProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *WorkflowProgress) GetMessages() []ProgressMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *WorkflowProgress) GetMessagesOk() (*[]ProgressMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *WorkflowProgress) SetMessages(v []ProgressMessage)`

SetMessages sets Messages field to given value.


### SetMessagesNil

`func (o *WorkflowProgress) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *WorkflowProgress) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
### GetPercent

`func (o *WorkflowProgress) GetPercent() int64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *WorkflowProgress) GetPercentOk() (*int64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *WorkflowProgress) SetPercent(v int64)`

SetPercent sets Percent field to given value.


### GetStatus

`func (o *WorkflowProgress) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowProgress) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowProgress) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStep

`func (o *WorkflowProgress) GetStep() string`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *WorkflowProgress) GetStepOk() (*string, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *WorkflowProgress) SetStep(v string)`

SetStep sets Step field to given value.


### GetStepIndex

`func (o *WorkflowProgress) GetStepIndex() int64`

GetStepIndex returns the StepIndex field if non-nil, zero value otherwise.

### GetStepIndexOk

`func (o *WorkflowProgress) GetStepIndexOk() (*int64, bool)`

GetStepIndexOk returns a tuple with the StepIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepIndex

`func (o *WorkflowProgress) SetStepIndex(v int64)`

SetStepIndex sets StepIndex field to given value.


### GetStepShare

`func (o *WorkflowProgress) GetStepShare() int64`

GetStepShare returns the StepShare field if non-nil, zero value otherwise.

### GetStepShareOk

`func (o *WorkflowProgress) GetStepShareOk() (*int64, bool)`

GetStepShareOk returns a tuple with the StepShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepShare

`func (o *WorkflowProgress) SetStepShare(v int64)`

SetStepShare sets StepShare field to given value.


### GetStepsTotal

`func (o *WorkflowProgress) GetStepsTotal() int64`

GetStepsTotal returns the StepsTotal field if non-nil, zero value otherwise.

### GetStepsTotalOk

`func (o *WorkflowProgress) GetStepsTotalOk() (*int64, bool)`

GetStepsTotalOk returns a tuple with the StepsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepsTotal

`func (o *WorkflowProgress) SetStepsTotal(v int64)`

SetStepsTotal sets StepsTotal field to given value.


### GetSubStep

`func (o *WorkflowProgress) GetSubStep() string`

GetSubStep returns the SubStep field if non-nil, zero value otherwise.

### GetSubStepOk

`func (o *WorkflowProgress) GetSubStepOk() (*string, bool)`

GetSubStepOk returns a tuple with the SubStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStep

`func (o *WorkflowProgress) SetSubStep(v string)`

SetSubStep sets SubStep field to given value.

### HasSubStep

`func (o *WorkflowProgress) HasSubStep() bool`

HasSubStep returns a boolean if a field has been set.

### GetSubStepDone

`func (o *WorkflowProgress) GetSubStepDone() int64`

GetSubStepDone returns the SubStepDone field if non-nil, zero value otherwise.

### GetSubStepDoneOk

`func (o *WorkflowProgress) GetSubStepDoneOk() (*int64, bool)`

GetSubStepDoneOk returns a tuple with the SubStepDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStepDone

`func (o *WorkflowProgress) SetSubStepDone(v int64)`

SetSubStepDone sets SubStepDone field to given value.

### HasSubStepDone

`func (o *WorkflowProgress) HasSubStepDone() bool`

HasSubStepDone returns a boolean if a field has been set.

### GetSubStepTotal

`func (o *WorkflowProgress) GetSubStepTotal() int64`

GetSubStepTotal returns the SubStepTotal field if non-nil, zero value otherwise.

### GetSubStepTotalOk

`func (o *WorkflowProgress) GetSubStepTotalOk() (*int64, bool)`

GetSubStepTotalOk returns a tuple with the SubStepTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStepTotal

`func (o *WorkflowProgress) SetSubStepTotal(v int64)`

SetSubStepTotal sets SubStepTotal field to given value.

### HasSubStepTotal

`func (o *WorkflowProgress) HasSubStepTotal() bool`

HasSubStepTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


