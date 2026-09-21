# KnownConstantEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvidenceKind** | Pointer to **string** |  | [optional] [default to "known_constant"]
**Kind** | Pointer to **string** |  | [optional] [default to "deterministic_derivation"]
**Effect** | [**EvidenceEffect**](EvidenceEffect.md) |  | 
**Strength** | Pointer to **string** |  | [optional] [default to "indirect"]
**Values** | [**[]BytesConstant**](BytesConstant.md) |  | 

## Methods

### NewKnownConstantEvidence

`func NewKnownConstantEvidence(effect EvidenceEffect, values []BytesConstant, ) *KnownConstantEvidence`

NewKnownConstantEvidence instantiates a new KnownConstantEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnownConstantEvidenceWithDefaults

`func NewKnownConstantEvidenceWithDefaults() *KnownConstantEvidence`

NewKnownConstantEvidenceWithDefaults instantiates a new KnownConstantEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvidenceKind

`func (o *KnownConstantEvidence) GetEvidenceKind() string`

GetEvidenceKind returns the EvidenceKind field if non-nil, zero value otherwise.

### GetEvidenceKindOk

`func (o *KnownConstantEvidence) GetEvidenceKindOk() (*string, bool)`

GetEvidenceKindOk returns a tuple with the EvidenceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceKind

`func (o *KnownConstantEvidence) SetEvidenceKind(v string)`

SetEvidenceKind sets EvidenceKind field to given value.

### HasEvidenceKind

`func (o *KnownConstantEvidence) HasEvidenceKind() bool`

HasEvidenceKind returns a boolean if a field has been set.

### GetKind

`func (o *KnownConstantEvidence) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KnownConstantEvidence) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KnownConstantEvidence) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *KnownConstantEvidence) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetEffect

`func (o *KnownConstantEvidence) GetEffect() EvidenceEffect`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *KnownConstantEvidence) GetEffectOk() (*EvidenceEffect, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *KnownConstantEvidence) SetEffect(v EvidenceEffect)`

SetEffect sets Effect field to given value.


### GetStrength

`func (o *KnownConstantEvidence) GetStrength() string`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *KnownConstantEvidence) GetStrengthOk() (*string, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *KnownConstantEvidence) SetStrength(v string)`

SetStrength sets Strength field to given value.

### HasStrength

`func (o *KnownConstantEvidence) HasStrength() bool`

HasStrength returns a boolean if a field has been set.

### GetValues

`func (o *KnownConstantEvidence) GetValues() []BytesConstant`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *KnownConstantEvidence) GetValuesOk() (*[]BytesConstant, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *KnownConstantEvidence) SetValues(v []BytesConstant)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


