# ImportedAPICallEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Calls** | [**[]ReportImportedAPICall**](ReportImportedAPICall.md) | The observed calls. | 
**Effect** | **string** | Whether the evidence supports or refutes the finding. | 
**EvidenceKind** | **string** | Kind of evidence. | 
**Kind** | **string** | How the evidence was derived. | 
**Strength** | **string** | How directly the evidence bears on the finding. | 

## Methods

### NewImportedAPICallEvidence

`func NewImportedAPICallEvidence(calls []ReportImportedAPICall, effect string, evidenceKind string, kind string, strength string, ) *ImportedAPICallEvidence`

NewImportedAPICallEvidence instantiates a new ImportedAPICallEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportedAPICallEvidenceWithDefaults

`func NewImportedAPICallEvidenceWithDefaults() *ImportedAPICallEvidence`

NewImportedAPICallEvidenceWithDefaults instantiates a new ImportedAPICallEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalls

`func (o *ImportedAPICallEvidence) GetCalls() []ReportImportedAPICall`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *ImportedAPICallEvidence) GetCallsOk() (*[]ReportImportedAPICall, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *ImportedAPICallEvidence) SetCalls(v []ReportImportedAPICall)`

SetCalls sets Calls field to given value.


### GetEffect

`func (o *ImportedAPICallEvidence) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *ImportedAPICallEvidence) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *ImportedAPICallEvidence) SetEffect(v string)`

SetEffect sets Effect field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


