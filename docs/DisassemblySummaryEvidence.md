# DisassemblySummaryEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effect** | **string** | Whether the evidence supports or refutes the finding. | 
**EvidenceKind** | **string** | Kind of evidence. | 
**Kind** | **string** | How the evidence was derived. | 
**Strength** | **string** | How directly the evidence bears on the finding. | 
**Summaries** | [**[]ReportCodeSummary**](ReportCodeSummary.md) | Summaries of the disassembled code. | 

## Methods

### NewDisassemblySummaryEvidence

`func NewDisassemblySummaryEvidence(effect string, evidenceKind string, kind string, strength string, summaries []ReportCodeSummary, ) *DisassemblySummaryEvidence`

NewDisassemblySummaryEvidence instantiates a new DisassemblySummaryEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisassemblySummaryEvidenceWithDefaults

`func NewDisassemblySummaryEvidenceWithDefaults() *DisassemblySummaryEvidence`

NewDisassemblySummaryEvidenceWithDefaults instantiates a new DisassemblySummaryEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffect

`func (o *DisassemblySummaryEvidence) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *DisassemblySummaryEvidence) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *DisassemblySummaryEvidence) SetEffect(v string)`

SetEffect sets Effect field to given value.


### GetEvidenceKind

`func (o *DisassemblySummaryEvidence) GetEvidenceKind() string`

GetEvidenceKind returns the EvidenceKind field if non-nil, zero value otherwise.

### GetEvidenceKindOk

`func (o *DisassemblySummaryEvidence) GetEvidenceKindOk() (*string, bool)`

GetEvidenceKindOk returns a tuple with the EvidenceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceKind

`func (o *DisassemblySummaryEvidence) SetEvidenceKind(v string)`

SetEvidenceKind sets EvidenceKind field to given value.


### GetKind

`func (o *DisassemblySummaryEvidence) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DisassemblySummaryEvidence) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DisassemblySummaryEvidence) SetKind(v string)`

SetKind sets Kind field to given value.


### GetStrength

`func (o *DisassemblySummaryEvidence) GetStrength() string`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *DisassemblySummaryEvidence) GetStrengthOk() (*string, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *DisassemblySummaryEvidence) SetStrength(v string)`

SetStrength sets Strength field to given value.


### GetSummaries

`func (o *DisassemblySummaryEvidence) GetSummaries() []ReportCodeSummary`

GetSummaries returns the Summaries field if non-nil, zero value otherwise.

### GetSummariesOk

`func (o *DisassemblySummaryEvidence) GetSummariesOk() (*[]ReportCodeSummary, bool)`

GetSummariesOk returns a tuple with the Summaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaries

`func (o *DisassemblySummaryEvidence) SetSummaries(v []ReportCodeSummary)`

SetSummaries sets Summaries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


