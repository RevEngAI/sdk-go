# ReportReferencedConstant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Display** | Pointer to **string** | Suggested format for displaying the value at this reference, absent to use the constant&#39;s own. | [optional] 
**Subject** | [**ReportSubject**](ReportSubject.md) | Where the constant is referenced. | 
**Value** | [**ReportBytesConstant**](ReportBytesConstant.md) | The constant. | 

## Methods

### NewReportReferencedConstant

`func NewReportReferencedConstant(subject ReportSubject, value ReportBytesConstant, ) *ReportReferencedConstant`

NewReportReferencedConstant instantiates a new ReportReferencedConstant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportReferencedConstantWithDefaults

`func NewReportReferencedConstantWithDefaults() *ReportReferencedConstant`

NewReportReferencedConstantWithDefaults instantiates a new ReportReferencedConstant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplay

`func (o *ReportReferencedConstant) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ReportReferencedConstant) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ReportReferencedConstant) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *ReportReferencedConstant) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetSubject

`func (o *ReportReferencedConstant) GetSubject() ReportSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ReportReferencedConstant) GetSubjectOk() (*ReportSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ReportReferencedConstant) SetSubject(v ReportSubject)`

SetSubject sets Subject field to given value.


### GetValue

`func (o *ReportReferencedConstant) GetValue() ReportBytesConstant`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ReportReferencedConstant) GetValueOk() (*ReportBytesConstant, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ReportReferencedConstant) SetValue(v ReportBytesConstant)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


