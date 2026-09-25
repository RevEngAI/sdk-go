# WorkflowDayBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyseCapabilities** | [**AnalyseCapabilitiesBody**](AnalyseCapabilitiesBody.md) | ANALYSE_CAPABILITIES agent-task stats for the day | 
**Date** | **string** | Day this entry covers | 
**ReportAnalysis** | [**ReportAnalysisBody**](ReportAnalysisBody.md) | REPORT_ANALYSIS agent-task stats for the day | 

## Methods

### NewWorkflowDayBody

`func NewWorkflowDayBody(analyseCapabilities AnalyseCapabilitiesBody, date string, reportAnalysis ReportAnalysisBody, ) *WorkflowDayBody`

NewWorkflowDayBody instantiates a new WorkflowDayBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowDayBodyWithDefaults

`func NewWorkflowDayBodyWithDefaults() *WorkflowDayBody`

NewWorkflowDayBodyWithDefaults instantiates a new WorkflowDayBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyseCapabilities

`func (o *WorkflowDayBody) GetAnalyseCapabilities() AnalyseCapabilitiesBody`

GetAnalyseCapabilities returns the AnalyseCapabilities field if non-nil, zero value otherwise.

### GetAnalyseCapabilitiesOk

`func (o *WorkflowDayBody) GetAnalyseCapabilitiesOk() (*AnalyseCapabilitiesBody, bool)`

GetAnalyseCapabilitiesOk returns a tuple with the AnalyseCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyseCapabilities

`func (o *WorkflowDayBody) SetAnalyseCapabilities(v AnalyseCapabilitiesBody)`

SetAnalyseCapabilities sets AnalyseCapabilities field to given value.


### GetDate

`func (o *WorkflowDayBody) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *WorkflowDayBody) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *WorkflowDayBody) SetDate(v string)`

SetDate sets Date field to given value.


### GetReportAnalysis

`func (o *WorkflowDayBody) GetReportAnalysis() ReportAnalysisBody`

GetReportAnalysis returns the ReportAnalysis field if non-nil, zero value otherwise.

### GetReportAnalysisOk

`func (o *WorkflowDayBody) GetReportAnalysisOk() (*ReportAnalysisBody, bool)`

GetReportAnalysisOk returns a tuple with the ReportAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportAnalysis

`func (o *WorkflowDayBody) SetReportAnalysis(v ReportAnalysisBody)`

SetReportAnalysis sets ReportAnalysis field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


