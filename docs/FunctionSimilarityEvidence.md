# FunctionSimilarityEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvidenceKind** | Pointer to **string** |  | [optional] [default to "function_similarity"]
**Kind** | Pointer to **string** |  | [optional] [default to "deterministic_derivation"]
**Effect** | [**EvidenceEffect**](EvidenceEffect.md) |  | 
**Strength** | Pointer to **string** |  | [optional] [default to "indirect"]
**Similarities** | [**[]FunctionSimilarity**](FunctionSimilarity.md) |  | 

## Methods

### NewFunctionSimilarityEvidence

`func NewFunctionSimilarityEvidence(effect EvidenceEffect, similarities []FunctionSimilarity, ) *FunctionSimilarityEvidence`

NewFunctionSimilarityEvidence instantiates a new FunctionSimilarityEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionSimilarityEvidenceWithDefaults

`func NewFunctionSimilarityEvidenceWithDefaults() *FunctionSimilarityEvidence`

NewFunctionSimilarityEvidenceWithDefaults instantiates a new FunctionSimilarityEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvidenceKind

`func (o *FunctionSimilarityEvidence) GetEvidenceKind() string`

GetEvidenceKind returns the EvidenceKind field if non-nil, zero value otherwise.

### GetEvidenceKindOk

`func (o *FunctionSimilarityEvidence) GetEvidenceKindOk() (*string, bool)`

GetEvidenceKindOk returns a tuple with the EvidenceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceKind

`func (o *FunctionSimilarityEvidence) SetEvidenceKind(v string)`

SetEvidenceKind sets EvidenceKind field to given value.

### HasEvidenceKind

`func (o *FunctionSimilarityEvidence) HasEvidenceKind() bool`

HasEvidenceKind returns a boolean if a field has been set.

### GetKind

`func (o *FunctionSimilarityEvidence) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FunctionSimilarityEvidence) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FunctionSimilarityEvidence) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FunctionSimilarityEvidence) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetEffect

`func (o *FunctionSimilarityEvidence) GetEffect() EvidenceEffect`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *FunctionSimilarityEvidence) GetEffectOk() (*EvidenceEffect, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *FunctionSimilarityEvidence) SetEffect(v EvidenceEffect)`

SetEffect sets Effect field to given value.


### GetStrength

`func (o *FunctionSimilarityEvidence) GetStrength() string`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *FunctionSimilarityEvidence) GetStrengthOk() (*string, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *FunctionSimilarityEvidence) SetStrength(v string)`

SetStrength sets Strength field to given value.

### HasStrength

`func (o *FunctionSimilarityEvidence) HasStrength() bool`

HasStrength returns a boolean if a field has been set.

### GetSimilarities

`func (o *FunctionSimilarityEvidence) GetSimilarities() []FunctionSimilarity`

GetSimilarities returns the Similarities field if non-nil, zero value otherwise.

### GetSimilaritiesOk

`func (o *FunctionSimilarityEvidence) GetSimilaritiesOk() (*[]FunctionSimilarity, bool)`

GetSimilaritiesOk returns a tuple with the Similarities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimilarities

`func (o *FunctionSimilarityEvidence) SetSimilarities(v []FunctionSimilarity)`

SetSimilarities sets Similarities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


