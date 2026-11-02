# ImportedAPICallEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvidenceKind** | Pointer to **string** |  | [optional] [default to "imported_api_call"]
**Kind** | Pointer to **string** |  | [optional] [default to "deterministic_derivation"]
**Effect** | [**EvidenceEffect**](EvidenceEffect.md) |  | 
**Strength** | Pointer to **string** |  | [optional] [default to "direct"]
**Calls** | [**[]ImportedApiCall**](ImportedApiCall.md) |  | 

## Methods

### NewImportedAPICallEvidence

`func NewImportedAPICallEvidence(effect EvidenceEffect, calls []ImportedApiCall, ) *ImportedAPICallEvidence`

NewImportedAPICallEvidence instantiates a new ImportedAPICallEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportedAPICallEvidenceWithDefaults

`func NewImportedAPICallEvidenceWithDefaults() *ImportedAPICallEvidence`

NewImportedAPICallEvidenceWithDefaults instantiates a new ImportedAPICallEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvidenceKind

`func (o *ImportedAPICallEvidence) GetEvidenceKind() string`

GetEvidenceKind returns the EvidenceKind field if non-nil, zero value otherwise.

### GetEvidenceKindOk

`func (o *ImportedAPICallEvidence) GetEvidenceKindOk() (*string, bool)`

GetEvidenceKindOk returns a tuple with the EvidenceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceKind

`func (o *ImportedAPICallEvidence) SetEvidenceKind(v string)`

SetEvidenceKind sets EvidenceKind field to given value.

### HasEvidenceKind

`func (o *ImportedAPICallEvidence) HasEvidenceKind() bool`

HasEvidenceKind returns a boolean if a field has been set.

### GetKind

`func (o *ImportedAPICallEvidence) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ImportedAPICallEvidence) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ImportedAPICallEvidence) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ImportedAPICallEvidence) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetEffect

`func (o *ImportedAPICallEvidence) GetEffect() EvidenceEffect`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *ImportedAPICallEvidence) GetEffectOk() (*EvidenceEffect, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *ImportedAPICallEvidence) SetEffect(v EvidenceEffect)`

SetEffect sets Effect field to given value.


### GetStrength

`func (o *ImportedAPICallEvidence) GetStrength() string`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *ImportedAPICallEvidence) GetStrengthOk() (*string, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *ImportedAPICallEvidence) SetStrength(v string)`

SetStrength sets Strength field to given value.

### HasStrength

`func (o *ImportedAPICallEvidence) HasStrength() bool`

HasStrength returns a boolean if a field has been set.

### GetCalls

`func (o *ImportedAPICallEvidence) GetCalls() []ImportedApiCall`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *ImportedAPICallEvidence) GetCallsOk() (*[]ImportedApiCall, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *ImportedAPICallEvidence) SetCalls(v []ImportedApiCall)`

SetCalls sets Calls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


