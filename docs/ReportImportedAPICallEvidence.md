# ReportImportedAPICallEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Calls** | [**[]ReportImportedAPICall**](ReportImportedAPICall.md) | The observed calls. | 
**Effect** | **string** | Whether the evidence supports or refutes the finding. | 
**EvidenceKind** | **string** | Kind of evidence. | 
**Kind** | **string** | How the evidence was derived. | 
**Strength** | **string** | How directly the evidence bears on the finding. | 

## Methods

### NewReportImportedAPICallEvidence

`func NewReportImportedAPICallEvidence(calls []ReportImportedAPICall, effect string, evidenceKind string, kind string, strength string, ) *ReportImportedAPICallEvidence`

NewReportImportedAPICallEvidence instantiates a new ReportImportedAPICallEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportImportedAPICallEvidenceWithDefaults

`func NewReportImportedAPICallEvidenceWithDefaults() *ReportImportedAPICallEvidence`

NewReportImportedAPICallEvidenceWithDefaults instantiates a new ReportImportedAPICallEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalls

`func (o *ReportImportedAPICallEvidence) GetCalls() []ReportImportedAPICall`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *ReportImportedAPICallEvidence) GetCallsOk() (*[]ReportImportedAPICall, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *ReportImportedAPICallEvidence) SetCalls(v []ReportImportedAPICall)`

SetCalls sets Calls field to given value.


### GetEffect

`func (o *ReportImportedAPICallEvidence) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *ReportImportedAPICallEvidence) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *ReportImportedAPICallEvidence) SetEffect(v string)`

SetEffect sets Effect field to given value.


### GetEvidenceKind

`func (o *ReportImportedAPICallEvidence) GetEvidenceKind() string`

GetEvidenceKind returns the EvidenceKind field if non-nil, zero value otherwise.

### GetEvidenceKindOk

`func (o *ReportImportedAPICallEvidence) GetEvidenceKindOk() (*string, bool)`

GetEvidenceKindOk returns a tuple with the EvidenceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceKind

`func (o *ReportImportedAPICallEvidence) SetEvidenceKind(v string)`

SetEvidenceKind sets EvidenceKind field to given value.


### GetKind

`func (o *ReportImportedAPICallEvidence) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ReportImportedAPICallEvidence) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ReportImportedAPICallEvidence) SetKind(v string)`

SetKind sets Kind field to given value.


### GetStrength

`func (o *ReportImportedAPICallEvidence) GetStrength() string`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *ReportImportedAPICallEvidence) GetStrengthOk() (*string, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *ReportImportedAPICallEvidence) SetStrength(v string)`

SetStrength sets Strength field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


