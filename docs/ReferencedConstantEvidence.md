# ReferencedConstantEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvidenceKind** | Pointer to **string** |  | [optional] [default to "referenced_constant"]
**Kind** | Pointer to **string** |  | [optional] [default to "deterministic_derivation"]
**Effect** | [**EvidenceEffect**](EvidenceEffect.md) |  | 
**Strength** | Pointer to **string** |  | [optional] [default to "direct"]
**Constants** | [**[]ReferencedConstant**](ReferencedConstant.md) |  | 

## Methods

### NewReferencedConstantEvidence

`func NewReferencedConstantEvidence(effect EvidenceEffect, constants []ReferencedConstant, ) *ReferencedConstantEvidence`

NewReferencedConstantEvidence instantiates a new ReferencedConstantEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferencedConstantEvidenceWithDefaults

`func NewReferencedConstantEvidenceWithDefaults() *ReferencedConstantEvidence`

NewReferencedConstantEvidenceWithDefaults instantiates a new ReferencedConstantEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvidenceKind

`func (o *ReferencedConstantEvidence) GetEvidenceKind() string`

GetEvidenceKind returns the EvidenceKind field if non-nil, zero value otherwise.

### GetEvidenceKindOk

`func (o *ReferencedConstantEvidence) GetEvidenceKindOk() (*string, bool)`

GetEvidenceKindOk returns a tuple with the EvidenceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceKind

`func (o *ReferencedConstantEvidence) SetEvidenceKind(v string)`

SetEvidenceKind sets EvidenceKind field to given value.

### HasEvidenceKind

`func (o *ReferencedConstantEvidence) HasEvidenceKind() bool`

HasEvidenceKind returns a boolean if a field has been set.

### GetKind

`func (o *ReferencedConstantEvidence) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ReferencedConstantEvidence) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ReferencedConstantEvidence) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ReferencedConstantEvidence) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetEffect

`func (o *ReferencedConstantEvidence) GetEffect() EvidenceEffect`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *ReferencedConstantEvidence) GetEffectOk() (*EvidenceEffect, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *ReferencedConstantEvidence) SetEffect(v EvidenceEffect)`

SetEffect sets Effect field to given value.


### GetStrength

`func (o *ReferencedConstantEvidence) GetStrength() string`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *ReferencedConstantEvidence) GetStrengthOk() (*string, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *ReferencedConstantEvidence) SetStrength(v string)`

SetStrength sets Strength field to given value.

### HasStrength

`func (o *ReferencedConstantEvidence) HasStrength() bool`

HasStrength returns a boolean if a field has been set.

### GetConstants

`func (o *ReferencedConstantEvidence) GetConstants() []ReferencedConstant`

GetConstants returns the Constants field if non-nil, zero value otherwise.

### GetConstantsOk

`func (o *ReferencedConstantEvidence) GetConstantsOk() (*[]ReferencedConstant, bool)`

GetConstantsOk returns a tuple with the Constants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstants

`func (o *ReferencedConstantEvidence) SetConstants(v []ReferencedConstant)`

SetConstants sets Constants field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


