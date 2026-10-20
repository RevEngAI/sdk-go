# ReportStringMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to [**ReportSubject**](ReportSubject.md) | Where the string is located, absent when unknown. | [optional] 
**Value** | **string** | The matched string. | 

## Methods

### NewReportStringMatch

`func NewReportStringMatch(value string, ) *ReportStringMatch`

NewReportStringMatch instantiates a new ReportStringMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportStringMatchWithDefaults

`func NewReportStringMatchWithDefaults() *ReportStringMatch`

NewReportStringMatchWithDefaults instantiates a new ReportStringMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *ReportStringMatch) GetSubject() ReportSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ReportStringMatch) GetSubjectOk() (*ReportSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ReportStringMatch) SetSubject(v ReportSubject)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ReportStringMatch) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetValue

`func (o *ReportStringMatch) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ReportStringMatch) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ReportStringMatch) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


