# ApiCombinationEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvidenceKind** | Pointer to **string** |  | [optional] [default to "api_combination"]
**Kind** | Pointer to **string** |  | [optional] [default to "deterministic_derivation"]
**Effect** | [**EvidenceEffect**](EvidenceEffect.md) |  | 
**Strength** | Pointer to **string** |  | [optional] [default to "indirect"]
**Apis** | [**[]ImportedApi**](ImportedApi.md) |  | 

## Methods

### NewApiCombinationEvidence

`func NewApiCombinationEvidence(effect EvidenceEffect, apis []ImportedApi, ) *ApiCombinationEvidence`

NewApiCombinationEvidence instantiates a new ApiCombinationEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCombinationEvidenceWithDefaults

`func NewApiCombinationEvidenceWithDefaults() *ApiCombinationEvidence`

NewApiCombinationEvidenceWithDefaults instantiates a new ApiCombinationEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvidenceKind

`func (o *ApiCombinationEvidence) GetEvidenceKind() string`

GetEvidenceKind returns the EvidenceKind field if non-nil, zero value otherwise.

### GetEvidenceKindOk

`func (o *ApiCombinationEvidence) GetEvidenceKindOk() (*string, bool)`

GetEvidenceKindOk returns a tuple with the EvidenceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceKind

`func (o *ApiCombinationEvidence) SetEvidenceKind(v string)`

SetEvidenceKind sets EvidenceKind field to given value.

### HasEvidenceKind

`func (o *ApiCombinationEvidence) HasEvidenceKind() bool`

HasEvidenceKind returns a boolean if a field has been set.

### GetKind

`func (o *ApiCombinationEvidence) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ApiCombinationEvidence) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ApiCombinationEvidence) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ApiCombinationEvidence) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetEffect

`func (o *ApiCombinationEvidence) GetEffect() EvidenceEffect`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *ApiCombinationEvidence) GetEffectOk() (*EvidenceEffect, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *ApiCombinationEvidence) SetEffect(v EvidenceEffect)`

SetEffect sets Effect field to given value.


### GetStrength

`func (o *ApiCombinationEvidence) GetStrength() string`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *ApiCombinationEvidence) GetStrengthOk() (*string, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *ApiCombinationEvidence) SetStrength(v string)`

SetStrength sets Strength field to given value.

### HasStrength

`func (o *ApiCombinationEvidence) HasStrength() bool`

HasStrength returns a boolean if a field has been set.

### GetApis

`func (o *ApiCombinationEvidence) GetApis() []ImportedApi`

GetApis returns the Apis field if non-nil, zero value otherwise.

### GetApisOk

`func (o *ApiCombinationEvidence) GetApisOk() (*[]ImportedApi, bool)`

GetApisOk returns a tuple with the Apis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApis

`func (o *ApiCombinationEvidence) SetApis(v []ImportedApi)`

SetApis sets Apis field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


