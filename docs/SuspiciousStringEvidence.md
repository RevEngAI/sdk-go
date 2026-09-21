# SuspiciousStringEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvidenceKind** | Pointer to **string** |  | [optional] [default to "suspicious_string"]
**Kind** | Pointer to **string** |  | [optional] [default to "deterministic_derivation"]
**Effect** | [**EvidenceEffect**](EvidenceEffect.md) |  | 
**Strength** | Pointer to **string** |  | [optional] [default to "indirect"]
**Strings** | [**[]SuspiciousString**](SuspiciousString.md) |  | 

## Methods

### NewSuspiciousStringEvidence

`func NewSuspiciousStringEvidence(effect EvidenceEffect, strings []SuspiciousString, ) *SuspiciousStringEvidence`

NewSuspiciousStringEvidence instantiates a new SuspiciousStringEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuspiciousStringEvidenceWithDefaults

`func NewSuspiciousStringEvidenceWithDefaults() *SuspiciousStringEvidence`

NewSuspiciousStringEvidenceWithDefaults instantiates a new SuspiciousStringEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvidenceKind

`func (o *SuspiciousStringEvidence) GetEvidenceKind() string`

GetEvidenceKind returns the EvidenceKind field if non-nil, zero value otherwise.

### GetEvidenceKindOk

`func (o *SuspiciousStringEvidence) GetEvidenceKindOk() (*string, bool)`

GetEvidenceKindOk returns a tuple with the EvidenceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceKind

`func (o *SuspiciousStringEvidence) SetEvidenceKind(v string)`

SetEvidenceKind sets EvidenceKind field to given value.

### HasEvidenceKind

`func (o *SuspiciousStringEvidence) HasEvidenceKind() bool`

HasEvidenceKind returns a boolean if a field has been set.

### GetKind

`func (o *SuspiciousStringEvidence) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SuspiciousStringEvidence) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SuspiciousStringEvidence) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SuspiciousStringEvidence) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetEffect

`func (o *SuspiciousStringEvidence) GetEffect() EvidenceEffect`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *SuspiciousStringEvidence) GetEffectOk() (*EvidenceEffect, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *SuspiciousStringEvidence) SetEffect(v EvidenceEffect)`

SetEffect sets Effect field to given value.


### GetStrength

`func (o *SuspiciousStringEvidence) GetStrength() string`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *SuspiciousStringEvidence) GetStrengthOk() (*string, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *SuspiciousStringEvidence) SetStrength(v string)`

SetStrength sets Strength field to given value.

### HasStrength

`func (o *SuspiciousStringEvidence) HasStrength() bool`

HasStrength returns a boolean if a field has been set.

### GetStrings

`func (o *SuspiciousStringEvidence) GetStrings() []SuspiciousString`

GetStrings returns the Strings field if non-nil, zero value otherwise.

### GetStringsOk

`func (o *SuspiciousStringEvidence) GetStringsOk() (*[]SuspiciousString, bool)`

GetStringsOk returns a tuple with the Strings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrings

`func (o *SuspiciousStringEvidence) SetStrings(v []SuspiciousString)`

SetStrings sets Strings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


