# ReportCodeSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reachability** | Pointer to [**ReportCallChain**](ReportCallChain.md) | A call chain reaching the entity, absent when none was found. | [optional] 
**Subject** | [**ReportSubject**](ReportSubject.md) | The program location whose code is summarised. | 
**Summary** | **string** | What the code does. | 

## Methods

### NewReportCodeSummary

`func NewReportCodeSummary(subject ReportSubject, summary string, ) *ReportCodeSummary`

NewReportCodeSummary instantiates a new ReportCodeSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportCodeSummaryWithDefaults

`func NewReportCodeSummaryWithDefaults() *ReportCodeSummary`

NewReportCodeSummaryWithDefaults instantiates a new ReportCodeSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReachability

`func (o *ReportCodeSummary) GetReachability() ReportCallChain`

GetReachability returns the Reachability field if non-nil, zero value otherwise.

### GetReachabilityOk

`func (o *ReportCodeSummary) GetReachabilityOk() (*ReportCallChain, bool)`

GetReachabilityOk returns a tuple with the Reachability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachability

`func (o *ReportCodeSummary) SetReachability(v ReportCallChain)`

SetReachability sets Reachability field to given value.

### HasReachability

`func (o *ReportCodeSummary) HasReachability() bool`

HasReachability returns a boolean if a field has been set.

### GetSubject

`func (o *ReportCodeSummary) GetSubject() ReportSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ReportCodeSummary) GetSubjectOk() (*ReportSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ReportCodeSummary) SetSubject(v ReportSubject)`

SetSubject sets Subject field to given value.


### GetSummary

`func (o *ReportCodeSummary) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ReportCodeSummary) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ReportCodeSummary) SetSummary(v string)`

SetSummary sets Summary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


