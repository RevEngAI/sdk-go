# ReportResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Findings** | **interface{}** |  | 
**Meta** | **interface{}** |  | 

## Methods

### NewReportResult

`func NewReportResult(findings interface{}, meta interface{}, ) *ReportResult`

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

`func (o *ReportResult) GetFindings() interface{}`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *ReportResult) GetFindingsOk() (*interface{}, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *ReportResult) SetFindings(v interface{})`

SetFindings sets Findings field to given value.


### SetFindingsNil

`func (o *ReportResult) SetFindingsNil(b bool)`

 SetFindingsNil sets the value for Findings to be an explicit nil

### UnsetFindings
`func (o *ReportResult) UnsetFindings()`

UnsetFindings ensures that no value is present for Findings, not even an explicit nil
### GetMeta

`func (o *ReportResult) GetMeta() interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ReportResult) GetMetaOk() (*interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ReportResult) SetMeta(v interface{})`

SetMeta sets Meta field to given value.


### SetMetaNil

`func (o *ReportResult) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *ReportResult) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


