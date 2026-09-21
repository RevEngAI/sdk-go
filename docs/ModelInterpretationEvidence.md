# ModelInterpretationEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvidenceKind** | Pointer to **string** |  | [optional] [default to "model_interpretation"]
**Kind** | Pointer to **string** |  | [optional] [default to "model_interpretation"]
**Effect** | [**EvidenceEffect**](EvidenceEffect.md) |  | 
**Strength** | Pointer to [**EvidenceStrength**](EvidenceStrength.md) |  | [optional] [default to EVIDENCESTRENGTH_INDIRECT]
**Interpretations** | [**[]ModelInterpretation**](ModelInterpretation.md) |  | 

## Methods

### NewModelInterpretationEvidence

`func NewModelInterpretationEvidence(effect EvidenceEffect, interpretations []ModelInterpretation, ) *ModelInterpretationEvidence`

NewModelInterpretationEvidence instantiates a new ModelInterpretationEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelInterpretationEvidenceWithDefaults

`func NewModelInterpretationEvidenceWithDefaults() *ModelInterpretationEvidence`

NewModelInterpretationEvidenceWithDefaults instantiates a new ModelInterpretationEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### HasEvidenceKind

`func (o *ModelInterpretationEvidence) HasEvidenceKind() bool`

HasEvidenceKind returns a boolean if a field has been set.

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

### HasKind

`func (o *ModelInterpretationEvidence) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetEffect

`func (o *ModelInterpretationEvidence) GetEffect() EvidenceEffect`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *ModelInterpretationEvidence) GetEffectOk() (*EvidenceEffect, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *ModelInterpretationEvidence) SetEffect(v EvidenceEffect)`

SetEffect sets Effect field to given value.


### GetStrength

`func (o *ModelInterpretationEvidence) GetStrength() EvidenceStrength`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *ModelInterpretationEvidence) GetStrengthOk() (*EvidenceStrength, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *ModelInterpretationEvidence) SetStrength(v EvidenceStrength)`

SetStrength sets Strength field to given value.

### HasStrength

`func (o *ModelInterpretationEvidence) HasStrength() bool`

HasStrength returns a boolean if a field has been set.

### GetInterpretations

`func (o *ModelInterpretationEvidence) GetInterpretations() []ModelInterpretation`

GetInterpretations returns the Interpretations field if non-nil, zero value otherwise.

### GetInterpretationsOk

`func (o *ModelInterpretationEvidence) GetInterpretationsOk() (*[]ModelInterpretation, bool)`

GetInterpretationsOk returns a tuple with the Interpretations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterpretations

`func (o *ModelInterpretationEvidence) SetInterpretations(v []ModelInterpretation)`

SetInterpretations sets Interpretations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


