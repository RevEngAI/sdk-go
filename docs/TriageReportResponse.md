# TriageReportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SoftwareScore** | **float32** | Overall triage score for the software | 
**Summary** | **string** | Summary of the triage analysis | 
**Functions** | [**[]TriageFunctionResponse**](TriageFunctionResponse.md) | List of triaged functions | 

## Methods

### NewTriageReportResponse

`func NewTriageReportResponse(softwareScore float32, summary string, functions []TriageFunctionResponse, ) *TriageReportResponse`

NewTriageReportResponse instantiates a new TriageReportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriageReportResponseWithDefaults

`func NewTriageReportResponseWithDefaults() *TriageReportResponse`

NewTriageReportResponseWithDefaults instantiates a new TriageReportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSoftwareScore

`func (o *TriageReportResponse) GetSoftwareScore() float32`

GetSoftwareScore returns the SoftwareScore field if non-nil, zero value otherwise.

### GetSoftwareScoreOk

`func (o *TriageReportResponse) GetSoftwareScoreOk() (*float32, bool)`

GetSoftwareScoreOk returns a tuple with the SoftwareScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareScore

`func (o *TriageReportResponse) SetSoftwareScore(v float32)`

SetSoftwareScore sets SoftwareScore field to given value.


### GetSummary

`func (o *TriageReportResponse) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *TriageReportResponse) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *TriageReportResponse) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetFunctions

`func (o *TriageReportResponse) GetFunctions() []TriageFunctionResponse`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *TriageReportResponse) GetFunctionsOk() (*[]TriageFunctionResponse, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *TriageReportResponse) SetFunctions(v []TriageFunctionResponse)`

SetFunctions sets Functions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


