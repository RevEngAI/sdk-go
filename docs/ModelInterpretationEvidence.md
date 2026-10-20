# ModelInterpretationEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effect** | **string** | Whether the evidence supports or refutes the finding. | 
**EvidenceKind** | **string** | Kind of evidence. | 
**Interpretations** | [**[]ReportModelInterpretation**](ReportModelInterpretation.md) | Interpretations of the code by an AI model. | 
**Kind** | **string** | How the evidence was derived. | 
**Strength** | **string** | How directly the evidence bears on the finding. | 

## Methods

### NewModelInterpretationEvidence

`func NewModelInterpretationEvidence(effect string, evidenceKind string, interpretations []ReportModelInterpretation, kind string, strength string, ) *ModelInterpretationEvidence`

NewModelInterpretationEvidence instantiates a new ModelInterpretationEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelInterpretationEvidenceWithDefaults

`func NewModelInterpretationEvidenceWithDefaults() *ModelInterpretationEvidence`

NewModelInterpretationEvidenceWithDefaults instantiates a new ModelInterpretationEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffect

`func (o *ModelInterpretationEvidence) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *ModelInterpretationEvidence) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *ModelInterpretationEvidence) SetEffect(v string)`

SetEffect sets Effect field to given value.


### GetEvidenceKind

`func (o *ModelInterpretationEvidence) GetEvidenceKind() string`

GetEvidenceKind returns the EvidenceKind field if non-nil, zero value otherwise.

### GetEvidenceKindOk

`func (o *ModelInterpretationEvidence) GetEvidenceKindOk() (*string, bool)`

GetEvidenceKindOk returns a tuple with the EvidenceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceKind

`func (o *ModelInterpretationEvidence) SetEvidenceKind(v string)`

SetEvidenceKind sets EvidenceKind field to given value.


### GetInterpretations

`func (o *ModelInterpretationEvidence) GetInterpretations() []ReportModelInterpretation`

GetInterpretations returns the Interpretations field if non-nil, zero value otherwise.

### GetInterpretationsOk

`func (o *ModelInterpretationEvidence) GetInterpretationsOk() (*[]ReportModelInterpretation, bool)`

GetInterpretationsOk returns a tuple with the Interpretations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterpretations

`func (o *ModelInterpretationEvidence) SetInterpretations(v []ReportModelInterpretation)`

SetInterpretations sets Interpretations field to given value.


### GetKind

`func (o *ModelInterpretationEvidence) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ModelInterpretationEvidence) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ModelInterpretationEvidence) SetKind(v string)`

SetKind sets Kind field to given value.


### GetStrength

`func (o *ModelInterpretationEvidence) GetStrength() string`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *ModelInterpretationEvidence) GetStrengthOk() (*string, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *ModelInterpretationEvidence) SetStrength(v string)`

SetStrength sets Strength field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


