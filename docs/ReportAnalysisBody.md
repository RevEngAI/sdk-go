# ReportAnalysisBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completed** | **int64** | REPORT_ANALYSIS agent tasks completed that day | 
**SoftwareTypes** | [**SoftwareTypeCountsBody**](SoftwareTypeCountsBody.md) | Breakdown of completed reports by the software type the agent classified | 

## Methods

### NewReportAnalysisBody

`func NewReportAnalysisBody(completed int64, softwareTypes SoftwareTypeCountsBody, ) *ReportAnalysisBody`

NewReportAnalysisBody instantiates a new ReportAnalysisBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportAnalysisBodyWithDefaults

`func NewReportAnalysisBodyWithDefaults() *ReportAnalysisBody`

NewReportAnalysisBodyWithDefaults instantiates a new ReportAnalysisBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleted

`func (o *ReportAnalysisBody) GetCompleted() int64`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *ReportAnalysisBody) GetCompletedOk() (*int64, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *ReportAnalysisBody) SetCompleted(v int64)`

SetCompleted sets Completed field to given value.


### GetSoftwareTypes

`func (o *ReportAnalysisBody) GetSoftwareTypes() SoftwareTypeCountsBody`

GetSoftwareTypes returns the SoftwareTypes field if non-nil, zero value otherwise.

### GetSoftwareTypesOk

`func (o *ReportAnalysisBody) GetSoftwareTypesOk() (*SoftwareTypeCountsBody, bool)`

GetSoftwareTypesOk returns a tuple with the SoftwareTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTypes

`func (o *ReportAnalysisBody) SetSoftwareTypes(v SoftwareTypeCountsBody)`

SetSoftwareTypes sets SoftwareTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


