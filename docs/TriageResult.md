# TriageResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Functions** | [**[]TriageFunction**](TriageFunction.md) | Per-function assessments. A function whose address no longer resolves within the analysis is omitted. | 
**SoftwareScore** | **float64** | Maliciousness score for the binary, 0 to 1 | 
**Summary** | **string** | Summary of the triage assessment | 

## Methods

### NewTriageResult

`func NewTriageResult(functions []TriageFunction, softwareScore float64, summary string, ) *TriageResult`

NewTriageResult instantiates a new TriageResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriageResultWithDefaults

`func NewTriageResultWithDefaults() *TriageResult`

NewTriageResultWithDefaults instantiates a new TriageResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctions

`func (o *TriageResult) GetFunctions() []TriageFunction`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *TriageResult) GetFunctionsOk() (*[]TriageFunction, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *TriageResult) SetFunctions(v []TriageFunction)`

SetFunctions sets Functions field to given value.


### SetFunctionsNil

`func (o *TriageResult) SetFunctionsNil(b bool)`

 SetFunctionsNil sets the value for Functions to be an explicit nil

### UnsetFunctions
`func (o *TriageResult) UnsetFunctions()`

UnsetFunctions ensures that no value is present for Functions, not even an explicit nil
### GetSoftwareScore

`func (o *TriageResult) GetSoftwareScore() float64`

GetSoftwareScore returns the SoftwareScore field if non-nil, zero value otherwise.

### GetSoftwareScoreOk

`func (o *TriageResult) GetSoftwareScoreOk() (*float64, bool)`

GetSoftwareScoreOk returns a tuple with the SoftwareScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareScore

`func (o *TriageResult) SetSoftwareScore(v float64)`

SetSoftwareScore sets SoftwareScore field to given value.


### GetSummary

`func (o *TriageResult) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *TriageResult) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *TriageResult) SetSummary(v string)`

SetSummary sets Summary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


