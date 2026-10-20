# KnownConstantEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effect** | **string** | Whether the evidence supports or refutes the finding. | 
**EvidenceKind** | **string** | Kind of evidence. | 
**Kind** | **string** | How the evidence was derived. | 
**Strength** | **string** | How directly the evidence bears on the finding. | 
**Values** | [**[]ReportBytesConstant**](ReportBytesConstant.md) | The well-known constants found in the binary. | 

## Methods

### NewKnownConstantEvidence

`func NewKnownConstantEvidence(effect string, evidenceKind string, kind string, strength string, values []ReportBytesConstant, ) *KnownConstantEvidence`

NewKnownConstantEvidence instantiates a new KnownConstantEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnownConstantEvidenceWithDefaults

`func NewKnownConstantEvidenceWithDefaults() *KnownConstantEvidence`

NewKnownConstantEvidenceWithDefaults instantiates a new KnownConstantEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffect

`func (o *KnownConstantEvidence) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *KnownConstantEvidence) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *KnownConstantEvidence) SetEffect(v string)`

SetEffect sets Effect field to given value.


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


### GetValues

`func (o *KnownConstantEvidence) GetValues() []ReportBytesConstant`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *KnownConstantEvidence) GetValuesOk() (*[]ReportBytesConstant, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *KnownConstantEvidence) SetValues(v []ReportBytesConstant)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


