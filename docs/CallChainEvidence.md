# CallChainEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallChain** | **[]int64** | Addresses along the call chain. | 
**Effect** | **string** | Whether the evidence supports or refutes the finding. | 
**EvidenceKind** | **string** | Kind of evidence. | 
**Kind** | **string** | How the evidence was derived. | 
**Strength** | **string** | How directly the evidence bears on the finding. | 

## Methods

### NewCallChainEvidence

`func NewCallChainEvidence(callChain []int64, effect string, evidenceKind string, kind string, strength string, ) *CallChainEvidence`

NewCallChainEvidence instantiates a new CallChainEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallChainEvidenceWithDefaults

`func NewCallChainEvidenceWithDefaults() *CallChainEvidence`

NewCallChainEvidenceWithDefaults instantiates a new CallChainEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallChain

`func (o *CallChainEvidence) GetCallChain() []int64`

GetCallChain returns the CallChain field if non-nil, zero value otherwise.

### GetCallChainOk

`func (o *CallChainEvidence) GetCallChainOk() (*[]int64, bool)`

GetCallChainOk returns a tuple with the CallChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallChain

`func (o *CallChainEvidence) SetCallChain(v []int64)`

SetCallChain sets CallChain field to given value.


### GetEffect

`func (o *CallChainEvidence) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *CallChainEvidence) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *CallChainEvidence) SetEffect(v string)`

SetEffect sets Effect field to given value.


### GetEvidenceKind

`func (o *CallChainEvidence) GetEvidenceKind() string`

GetEvidenceKind returns the EvidenceKind field if non-nil, zero value otherwise.

### GetEvidenceKindOk

`func (o *CallChainEvidence) GetEvidenceKindOk() (*string, bool)`

GetEvidenceKindOk returns a tuple with the EvidenceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceKind

`func (o *CallChainEvidence) SetEvidenceKind(v string)`

SetEvidenceKind sets EvidenceKind field to given value.


### GetKind

`func (o *CallChainEvidence) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CallChainEvidence) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CallChainEvidence) SetKind(v string)`

SetKind sets Kind field to given value.


### GetStrength

`func (o *CallChainEvidence) GetStrength() string`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *CallChainEvidence) GetStrengthOk() (*string, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *CallChainEvidence) SetStrength(v string)`

SetStrength sets Strength field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


