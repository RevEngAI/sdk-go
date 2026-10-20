# APICombinationEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apis** | [**[]ReportImportedAPI**](ReportImportedAPI.md) | The APIs in the combination. | 
**Effect** | **string** | Whether the evidence supports or refutes the finding. | 
**EvidenceKind** | **string** | Kind of evidence. | 
**Kind** | **string** | How the evidence was derived. | 
**Strength** | **string** | How directly the evidence bears on the finding. | 

## Methods

### NewAPICombinationEvidence

`func NewAPICombinationEvidence(apis []ReportImportedAPI, effect string, evidenceKind string, kind string, strength string, ) *APICombinationEvidence`

NewAPICombinationEvidence instantiates a new APICombinationEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPICombinationEvidenceWithDefaults

`func NewAPICombinationEvidenceWithDefaults() *APICombinationEvidence`

NewAPICombinationEvidenceWithDefaults instantiates a new APICombinationEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApis

`func (o *APICombinationEvidence) GetApis() []ReportImportedAPI`

GetApis returns the Apis field if non-nil, zero value otherwise.

### GetApisOk

`func (o *APICombinationEvidence) GetApisOk() (*[]ReportImportedAPI, bool)`

GetApisOk returns a tuple with the Apis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApis

`func (o *APICombinationEvidence) SetApis(v []ReportImportedAPI)`

SetApis sets Apis field to given value.


### GetEffect

`func (o *APICombinationEvidence) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *APICombinationEvidence) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *APICombinationEvidence) SetEffect(v string)`

SetEffect sets Effect field to given value.


### GetEvidenceKind

`func (o *APICombinationEvidence) GetEvidenceKind() string`

GetEvidenceKind returns the EvidenceKind field if non-nil, zero value otherwise.

### GetEvidenceKindOk

`func (o *APICombinationEvidence) GetEvidenceKindOk() (*string, bool)`

GetEvidenceKindOk returns a tuple with the EvidenceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceKind

`func (o *APICombinationEvidence) SetEvidenceKind(v string)`

SetEvidenceKind sets EvidenceKind field to given value.


### GetKind

`func (o *APICombinationEvidence) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *APICombinationEvidence) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *APICombinationEvidence) SetKind(v string)`

SetKind sets Kind field to given value.


### GetStrength

`func (o *APICombinationEvidence) GetStrength() string`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *APICombinationEvidence) GetStrengthOk() (*string, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *APICombinationEvidence) SetStrength(v string)`

SetStrength sets Strength field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


