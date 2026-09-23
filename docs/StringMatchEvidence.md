# StringMatchEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvidenceKind** | Pointer to **string** |  | [optional] [default to "string_match"]
**Kind** | Pointer to **string** |  | [optional] [default to "deterministic_derivation"]
**Effect** | [**EvidenceEffect**](EvidenceEffect.md) |  | 
**Strength** | Pointer to **string** |  | [optional] [default to "indirect"]
**Strings** | [**[]StringMatch**](StringMatch.md) |  | 

## Methods

### NewStringMatchEvidence

`func NewStringMatchEvidence(effect EvidenceEffect, strings []StringMatch, ) *StringMatchEvidence`

NewStringMatchEvidence instantiates a new StringMatchEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStringMatchEvidenceWithDefaults

`func NewStringMatchEvidenceWithDefaults() *StringMatchEvidence`

NewStringMatchEvidenceWithDefaults instantiates a new StringMatchEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvidenceKind

`func (o *StringMatchEvidence) GetEvidenceKind() string`

GetEvidenceKind returns the EvidenceKind field if non-nil, zero value otherwise.

### GetEvidenceKindOk

`func (o *StringMatchEvidence) GetEvidenceKindOk() (*string, bool)`

GetEvidenceKindOk returns a tuple with the EvidenceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceKind

`func (o *StringMatchEvidence) SetEvidenceKind(v string)`

SetEvidenceKind sets EvidenceKind field to given value.

### HasEvidenceKind

`func (o *StringMatchEvidence) HasEvidenceKind() bool`

HasEvidenceKind returns a boolean if a field has been set.

### GetKind

`func (o *StringMatchEvidence) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *StringMatchEvidence) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *StringMatchEvidence) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *StringMatchEvidence) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetEffect

`func (o *StringMatchEvidence) GetEffect() EvidenceEffect`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *StringMatchEvidence) GetEffectOk() (*EvidenceEffect, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *StringMatchEvidence) SetEffect(v EvidenceEffect)`

SetEffect sets Effect field to given value.


### GetStrength

`func (o *StringMatchEvidence) GetStrength() string`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *StringMatchEvidence) GetStrengthOk() (*string, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *StringMatchEvidence) SetStrength(v string)`

SetStrength sets Strength field to given value.

### HasStrength

`func (o *StringMatchEvidence) HasStrength() bool`

HasStrength returns a boolean if a field has been set.

### GetStrings

`func (o *StringMatchEvidence) GetStrings() []StringMatch`

GetStrings returns the Strings field if non-nil, zero value otherwise.

### GetStringsOk

`func (o *StringMatchEvidence) GetStringsOk() (*[]StringMatch, bool)`

GetStringsOk returns a tuple with the Strings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrings

`func (o *StringMatchEvidence) SetStrings(v []StringMatch)`

SetStrings sets Strings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


