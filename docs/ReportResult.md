# ReportResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Findings** | [**[]ReportFinding**](ReportFinding.md) | The findings the agent reported, each with its supporting evidence. | 
**Meta** | [**ReportMeta**](ReportMeta.md) | Report-level metadata. | 

## Methods

### NewReportResult

`func NewReportResult(findings []ReportFinding, meta ReportMeta, ) *ReportResult`

NewReportResult instantiates a new ReportResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportResultWithDefaults

`func NewReportResultWithDefaults() *ReportResult`

NewReportResultWithDefaults instantiates a new ReportResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFindings

`func (o *ReportResult) GetFindings() []ReportFinding`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *ReportResult) GetFindingsOk() (*[]ReportFinding, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *ReportResult) SetFindings(v []ReportFinding)`

SetFindings sets Findings field to given value.


### GetMeta

`func (o *ReportResult) GetMeta() ReportMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ReportResult) GetMetaOk() (*ReportMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ReportResult) SetMeta(v ReportMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


