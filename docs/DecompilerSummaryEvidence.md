# DecompilerSummaryEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effect** | **string** | Whether the evidence supports or refutes the finding. | 
**EvidenceKind** | **string** | Kind of evidence. | 
**Kind** | **string** | How the evidence was derived. | 
**Strength** | **string** | How directly the evidence bears on the finding. | 
**Summaries** | [**[]ReportCodeSummary**](ReportCodeSummary.md) | Summaries of the decompiled code. | 

## Methods

### NewDecompilerSummaryEvidence

`func NewDecompilerSummaryEvidence(effect string, evidenceKind string, kind string, strength string, summaries []ReportCodeSummary, ) *DecompilerSummaryEvidence`

NewDecompilerSummaryEvidence instantiates a new DecompilerSummaryEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecompilerSummaryEvidenceWithDefaults

`func NewDecompilerSummaryEvidenceWithDefaults() *DecompilerSummaryEvidence`

NewDecompilerSummaryEvidenceWithDefaults instantiates a new DecompilerSummaryEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffect

`func (o *DecompilerSummaryEvidence) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *DecompilerSummaryEvidence) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *DecompilerSummaryEvidence) SetEffect(v string)`

SetEffect sets Effect field to given value.


### GetEvidenceKind

`func (o *DecompilerSummaryEvidence) GetEvidenceKind() string`

GetEvidenceKind returns the EvidenceKind field if non-nil, zero value otherwise.

### GetEvidenceKindOk

`func (o *DecompilerSummaryEvidence) GetEvidenceKindOk() (*string, bool)`

GetEvidenceKindOk returns a tuple with the EvidenceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceKind

`func (o *DecompilerSummaryEvidence) SetEvidenceKind(v string)`

SetEvidenceKind sets EvidenceKind field to given value.


### GetKind

`func (o *DecompilerSummaryEvidence) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DecompilerSummaryEvidence) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DecompilerSummaryEvidence) SetKind(v string)`

SetKind sets Kind field to given value.


### GetStrength

`func (o *DecompilerSummaryEvidence) GetStrength() string`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *DecompilerSummaryEvidence) GetStrengthOk() (*string, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *DecompilerSummaryEvidence) SetStrength(v string)`

SetStrength sets Strength field to given value.


### GetSummaries

`func (o *DecompilerSummaryEvidence) GetSummaries() []ReportCodeSummary`

GetSummaries returns the Summaries field if non-nil, zero value otherwise.

### GetSummariesOk

`func (o *DecompilerSummaryEvidence) GetSummariesOk() (*[]ReportCodeSummary, bool)`

GetSummariesOk returns a tuple with the Summaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaries

`func (o *DecompilerSummaryEvidence) SetSummaries(v []ReportCodeSummary)`

SetSummaries sets Summaries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


