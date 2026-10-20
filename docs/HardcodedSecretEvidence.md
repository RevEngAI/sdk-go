# HardcodedSecretEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | What the rule detects. | 
**Effect** | **string** | Whether the evidence supports or refutes the finding. | 
**Entropy** | **float64** | Shannon entropy of the secret, in bits per byte. | 
**EvidenceKind** | **string** | Kind of evidence. | 
**Kind** | **string** | How the evidence was derived. | 
**References** | **[]int64** | Address of the secret in the binary&#39;s data. | 
**RuleId** | **string** | Identifier of the secret-scanning rule that matched. | 
**Secret** | **string** | The secret as it appears in the binary. | 
**SecretKind** | **string** | Category of the secret. | 
**Size** | **int64** | Size of the secret in bytes. | 
**Strength** | **string** | How directly the evidence bears on the finding. | 

## Methods

### NewHardcodedSecretEvidence

`func NewHardcodedSecretEvidence(description string, effect string, entropy float64, evidenceKind string, kind string, references []int64, ruleId string, secret string, secretKind string, size int64, strength string, ) *HardcodedSecretEvidence`

NewHardcodedSecretEvidence instantiates a new HardcodedSecretEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHardcodedSecretEvidenceWithDefaults

`func NewHardcodedSecretEvidenceWithDefaults() *HardcodedSecretEvidence`

NewHardcodedSecretEvidenceWithDefaults instantiates a new HardcodedSecretEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetEntropy

`func (o *HardcodedSecretEvidence) GetEntropy() float64`

GetEntropy returns the Entropy field if non-nil, zero value otherwise.

### GetEntropyOk

`func (o *HardcodedSecretEvidence) GetEntropyOk() (*float64, bool)`

GetEntropyOk returns a tuple with the Entropy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntropy

`func (o *HardcodedSecretEvidence) SetEntropy(v float64)`

SetEntropy sets Entropy field to given value.


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


### GetReferences

`func (o *HardcodedSecretEvidence) GetReferences() []int64`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *HardcodedSecretEvidence) GetReferencesOk() (*[]int64, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *HardcodedSecretEvidence) SetReferences(v []int64)`

SetReferences sets References field to given value.


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


### GetSecretKind

`func (o *HardcodedSecretEvidence) GetSecretKind() string`

GetSecretKind returns the SecretKind field if non-nil, zero value otherwise.

### GetSecretKindOk

`func (o *HardcodedSecretEvidence) GetSecretKindOk() (*string, bool)`

GetSecretKindOk returns a tuple with the SecretKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKind

`func (o *HardcodedSecretEvidence) SetSecretKind(v string)`

SetSecretKind sets SecretKind field to given value.


### GetSize

`func (o *HardcodedSecretEvidence) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *HardcodedSecretEvidence) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *HardcodedSecretEvidence) SetSize(v int64)`

SetSize sets Size field to given value.


### GetStrength

`func (o *HardcodedSecretEvidence) GetStrength() string`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *HardcodedSecretEvidence) GetStrengthOk() (*string, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *HardcodedSecretEvidence) SetStrength(v string)`

SetStrength sets Strength field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


