# AgentWorkflowUsageOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Report** | [**[]WorkflowDayBody**](WorkflowDayBody.md) | One entry per day in the requested window, oldest first | 

## Methods

### NewAgentWorkflowUsageOutputBody

`func NewAgentWorkflowUsageOutputBody(report []WorkflowDayBody, ) *AgentWorkflowUsageOutputBody`

NewAgentWorkflowUsageOutputBody instantiates a new AgentWorkflowUsageOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentWorkflowUsageOutputBodyWithDefaults

`func NewAgentWorkflowUsageOutputBodyWithDefaults() *AgentWorkflowUsageOutputBody`

NewAgentWorkflowUsageOutputBodyWithDefaults instantiates a new AgentWorkflowUsageOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReport

`func (o *AgentWorkflowUsageOutputBody) GetReport() []WorkflowDayBody`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *AgentWorkflowUsageOutputBody) GetReportOk() (*[]WorkflowDayBody, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *AgentWorkflowUsageOutputBody) SetReport(v []WorkflowDayBody)`

SetReport sets Report field to given value.


### SetReportNil

`func (o *AgentWorkflowUsageOutputBody) SetReportNil(b bool)`

 SetReportNil sets the value for Report to be an explicit nil

### UnsetReport
`func (o *AgentWorkflowUsageOutputBody) UnsetReport()`

UnsetReport ensures that no value is present for Report, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


