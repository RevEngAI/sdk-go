# WorkflowProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**Messages** | [**[]ProgressMessage**](ProgressMessage.md) | Log messages emitted during execution | 
**Status** | **string** | Current workflow status | 
**Step** | **string** | Name of the current step | 
**StepIndex** | **int64** | Zero-based index of the current step | 
**StepsTotal** | **int64** | Total number of steps in the workflow | 

## Methods

### NewWorkflowProgress

`func NewWorkflowProgress(messages []ProgressMessage, status string, step string, stepIndex int64, stepsTotal int64, ) *WorkflowProgress`

NewWorkflowProgress instantiates a new WorkflowProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowProgressWithDefaults

`func NewWorkflowProgressWithDefaults() *WorkflowProgress`

NewWorkflowProgressWithDefaults instantiates a new WorkflowProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *WorkflowProgress) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *WorkflowProgress) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *WorkflowProgress) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *WorkflowProgress) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


