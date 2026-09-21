# HardcodedSecretEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvidenceKind** | Pointer to **string** |  | [optional] [default to "hardcoded_secret"]
**Kind** | Pointer to **string** |  | [optional] [default to "deterministic_derivation"]
**Effect** | Pointer to **string** |  | [optional] [default to "supports"]
**Strength** | [**EvidenceStrength**](EvidenceStrength.md) |  | 
**RuleId** | **string** |  | 
**Description** | **string** |  | 
**SecretKind** | [**RuleKind**](RuleKind.md) |  | 
**Secret** | **string** |  | 
**Entropy** | **float32** |  | 
**Size** | **int32** |  | 
**References** | **[]int32** |  | 

## Methods

### NewHardcodedSecretEvidence

`func NewHardcodedSecretEvidence(strength EvidenceStrength, ruleId string, description string, secretKind RuleKind, secret string, entropy float32, size int32, references []int32, ) *HardcodedSecretEvidence`

NewHardcodedSecretEvidence instantiates a new HardcodedSecretEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHardcodedSecretEvidenceWithDefaults

`func NewHardcodedSecretEvidenceWithDefaults() *HardcodedSecretEvidence`

NewHardcodedSecretEvidenceWithDefaults instantiates a new HardcodedSecretEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvidenceKind

`func (o *HardcodedSecretEvidence) GetEvidenceKind() string`

GetEvidenceKind returns the EvidenceKind field if non-nil, zero value otherwise.

### GetEvidenceKindOk

`func (o *HardcodedSecretEvidence) GetEvidenceKindOk() (*string, bool)`

GetEvidenceKindOk returns a tuple with the EvidenceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceKind

`func (o *HardcodedSecretEvidence) SetEvidenceKind(v string)`

SetEvidenceKind sets EvidenceKind field to given value.

### HasEvidenceKind

`func (o *HardcodedSecretEvidence) HasEvidenceKind() bool`

HasEvidenceKind returns a boolean if a field has been set.

### GetKind

`func (o *HardcodedSecretEvidence) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *HardcodedSecretEvidence) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *HardcodedSecretEvidence) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *HardcodedSecretEvidence) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetEffect

`func (o *HardcodedSecretEvidence) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *HardcodedSecretEvidence) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *HardcodedSecretEvidence) SetEffect(v string)`

SetEffect sets Effect field to given value.

### HasEffect

`func (o *HardcodedSecretEvidence) HasEffect() bool`

HasEffect returns a boolean if a field has been set.

### GetStrength

`func (o *HardcodedSecretEvidence) GetStrength() EvidenceStrength`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *HardcodedSecretEvidence) GetStrengthOk() (*EvidenceStrength, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *HardcodedSecretEvidence) SetStrength(v EvidenceStrength)`

SetStrength sets Strength field to given value.


### GetRuleId

`func (o *HardcodedSecretEvidence) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *HardcodedSecretEvidence) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *HardcodedSecretEvidence) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.


### GetDescription

`func (o *HardcodedSecretEvidence) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HardcodedSecretEvidence) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HardcodedSecretEvidence) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSecretKind

`func (o *HardcodedSecretEvidence) GetSecretKind() RuleKind`

GetSecretKind returns the SecretKind field if non-nil, zero value otherwise.

### GetSecretKindOk

`func (o *HardcodedSecretEvidence) GetSecretKindOk() (*RuleKind, bool)`

GetSecretKindOk returns a tuple with the SecretKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKind

`func (o *HardcodedSecretEvidence) SetSecretKind(v RuleKind)`

SetSecretKind sets SecretKind field to given value.


### GetSecret

`func (o *HardcodedSecretEvidence) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *HardcodedSecretEvidence) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *HardcodedSecretEvidence) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetEntropy

`func (o *HardcodedSecretEvidence) GetEntropy() float32`

GetEntropy returns the Entropy field if non-nil, zero value otherwise.

### GetEntropyOk

`func (o *HardcodedSecretEvidence) GetEntropyOk() (*float32, bool)`

GetEntropyOk returns a tuple with the Entropy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntropy

`func (o *HardcodedSecretEvidence) SetEntropy(v float32)`

SetEntropy sets Entropy field to given value.


### GetSize

`func (o *HardcodedSecretEvidence) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *HardcodedSecretEvidence) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *HardcodedSecretEvidence) SetSize(v int32)`

SetSize sets Size field to given value.


### GetReferences

`func (o *HardcodedSecretEvidence) GetReferences() []int32`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *HardcodedSecretEvidence) GetReferencesOk() (*[]int32, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *HardcodedSecretEvidence) SetReferences(v []int32)`

SetReferences sets References field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


