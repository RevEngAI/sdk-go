# ReportFindingEvidenceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallChain** | **[]int64** | Addresses along the call chain. | 
**Effect** | **string** | Whether the evidence supports or refutes the finding. | 
**EvidenceKind** | **string** | Kind of evidence. | 
**Kind** | **string** | How the evidence was derived. | 
**Strength** | **string** | How directly the evidence bears on the finding. | 
**Constants** | [**[]ReportReferencedConstant**](ReportReferencedConstant.md) | The constants and where they are referenced. | 
**Calls** | [**[]ReportImportedAPICall**](ReportImportedAPICall.md) | The observed calls. | 
**Strings** | [**[]ReportStringMatch**](ReportStringMatch.md) | The matched strings. | 
**Similarities** | [**[]ReportFunctionSimilarity**](ReportFunctionSimilarity.md) | The similar functions. | 
**Apis** | [**[]ReportImportedAPI**](ReportImportedAPI.md) | The APIs in the combination. | 
**Summaries** | [**[]ReportCodeSummary**](ReportCodeSummary.md) | Summaries of the disassembled code. | 
**Interpretations** | [**[]ReportModelInterpretation**](ReportModelInterpretation.md) | Interpretations of the code by an AI model. | 
**Values** | [**[]ReportBytesConstant**](ReportBytesConstant.md) | The well-known constants found in the binary. | 
**Description** | **string** | What the rule detects. | 
**Entropy** | **float64** | Shannon entropy of the secret, in bits per byte. | 
**References** | **[]int64** | Address of the secret in the binary&#39;s data. | 
**RuleId** | **string** | Identifier of the secret-scanning rule that matched. | 
**Secret** | **string** | The secret as it appears in the binary. | 
**SecretKind** | **string** | Category of the secret. | 
**Size** | **int64** | Size of the secret in bytes. | 

## Methods

### NewReportFindingEvidenceInner

`func NewReportFindingEvidenceInner(callChain []int64, effect string, evidenceKind string, kind string, strength string, constants []ReportReferencedConstant, calls []ReportImportedAPICall, strings []ReportStringMatch, similarities []ReportFunctionSimilarity, apis []ReportImportedAPI, summaries []ReportCodeSummary, interpretations []ReportModelInterpretation, values []ReportBytesConstant, description string, entropy float64, references []int64, ruleId string, secret string, secretKind string, size int64, ) *ReportFindingEvidenceInner`

NewReportFindingEvidenceInner instantiates a new ReportFindingEvidenceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportFindingEvidenceInnerWithDefaults

`func NewReportFindingEvidenceInnerWithDefaults() *ReportFindingEvidenceInner`

NewReportFindingEvidenceInnerWithDefaults instantiates a new ReportFindingEvidenceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallChain

`func (o *ReportFindingEvidenceInner) GetCallChain() []int64`

GetCallChain returns the CallChain field if non-nil, zero value otherwise.

### GetCallChainOk

`func (o *ReportFindingEvidenceInner) GetCallChainOk() (*[]int64, bool)`

GetCallChainOk returns a tuple with the CallChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallChain

`func (o *ReportFindingEvidenceInner) SetCallChain(v []int64)`

SetCallChain sets CallChain field to given value.


### GetEffect

`func (o *ReportFindingEvidenceInner) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *ReportFindingEvidenceInner) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *ReportFindingEvidenceInner) SetEffect(v string)`

SetEffect sets Effect field to given value.


### GetEvidenceKind

`func (o *ReportFindingEvidenceInner) GetEvidenceKind() string`

GetEvidenceKind returns the EvidenceKind field if non-nil, zero value otherwise.

### GetEvidenceKindOk

`func (o *ReportFindingEvidenceInner) GetEvidenceKindOk() (*string, bool)`

GetEvidenceKindOk returns a tuple with the EvidenceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceKind

`func (o *ReportFindingEvidenceInner) SetEvidenceKind(v string)`

SetEvidenceKind sets EvidenceKind field to given value.


### GetKind

`func (o *ReportFindingEvidenceInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ReportFindingEvidenceInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ReportFindingEvidenceInner) SetKind(v string)`

SetKind sets Kind field to given value.


### GetStrength

`func (o *ReportFindingEvidenceInner) GetStrength() string`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *ReportFindingEvidenceInner) GetStrengthOk() (*string, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *ReportFindingEvidenceInner) SetStrength(v string)`

SetStrength sets Strength field to given value.


### GetConstants

`func (o *ReportFindingEvidenceInner) GetConstants() []ReportReferencedConstant`

GetConstants returns the Constants field if non-nil, zero value otherwise.

### GetConstantsOk

`func (o *ReportFindingEvidenceInner) GetConstantsOk() (*[]ReportReferencedConstant, bool)`

GetConstantsOk returns a tuple with the Constants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstants

`func (o *ReportFindingEvidenceInner) SetConstants(v []ReportReferencedConstant)`

SetConstants sets Constants field to given value.


### GetCalls

`func (o *ReportFindingEvidenceInner) GetCalls() []ReportImportedAPICall`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *ReportFindingEvidenceInner) GetCallsOk() (*[]ReportImportedAPICall, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *ReportFindingEvidenceInner) SetCalls(v []ReportImportedAPICall)`

SetCalls sets Calls field to given value.


### GetStrings

`func (o *ReportFindingEvidenceInner) GetStrings() []ReportStringMatch`

GetStrings returns the Strings field if non-nil, zero value otherwise.

### GetStringsOk

`func (o *ReportFindingEvidenceInner) GetStringsOk() (*[]ReportStringMatch, bool)`

GetStringsOk returns a tuple with the Strings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrings

`func (o *ReportFindingEvidenceInner) SetStrings(v []ReportStringMatch)`

SetStrings sets Strings field to given value.


### GetSimilarities

`func (o *ReportFindingEvidenceInner) GetSimilarities() []ReportFunctionSimilarity`

GetSimilarities returns the Similarities field if non-nil, zero value otherwise.

### GetSimilaritiesOk

`func (o *ReportFindingEvidenceInner) GetSimilaritiesOk() (*[]ReportFunctionSimilarity, bool)`

GetSimilaritiesOk returns a tuple with the Similarities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimilarities

`func (o *ReportFindingEvidenceInner) SetSimilarities(v []ReportFunctionSimilarity)`

SetSimilarities sets Similarities field to given value.


### GetApis

`func (o *ReportFindingEvidenceInner) GetApis() []ReportImportedAPI`

GetApis returns the Apis field if non-nil, zero value otherwise.

### GetApisOk

`func (o *ReportFindingEvidenceInner) GetApisOk() (*[]ReportImportedAPI, bool)`

GetApisOk returns a tuple with the Apis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApis

`func (o *ReportFindingEvidenceInner) SetApis(v []ReportImportedAPI)`

SetApis sets Apis field to given value.


### GetSummaries

`func (o *ReportFindingEvidenceInner) GetSummaries() []ReportCodeSummary`

GetSummaries returns the Summaries field if non-nil, zero value otherwise.

### GetSummariesOk

`func (o *ReportFindingEvidenceInner) GetSummariesOk() (*[]ReportCodeSummary, bool)`

GetSummariesOk returns a tuple with the Summaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaries

`func (o *ReportFindingEvidenceInner) SetSummaries(v []ReportCodeSummary)`

SetSummaries sets Summaries field to given value.


### GetInterpretations

`func (o *ReportFindingEvidenceInner) GetInterpretations() []ReportModelInterpretation`

GetInterpretations returns the Interpretations field if non-nil, zero value otherwise.

### GetInterpretationsOk

`func (o *ReportFindingEvidenceInner) GetInterpretationsOk() (*[]ReportModelInterpretation, bool)`

GetInterpretationsOk returns a tuple with the Interpretations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterpretations

`func (o *ReportFindingEvidenceInner) SetInterpretations(v []ReportModelInterpretation)`

SetInterpretations sets Interpretations field to given value.


### GetValues

`func (o *ReportFindingEvidenceInner) GetValues() []ReportBytesConstant`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ReportFindingEvidenceInner) GetValuesOk() (*[]ReportBytesConstant, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ReportFindingEvidenceInner) SetValues(v []ReportBytesConstant)`

SetValues sets Values field to given value.


### GetDescription

`func (o *ReportFindingEvidenceInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReportFindingEvidenceInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReportFindingEvidenceInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEntropy

`func (o *ReportFindingEvidenceInner) GetEntropy() float64`

GetEntropy returns the Entropy field if non-nil, zero value otherwise.

### GetEntropyOk

`func (o *ReportFindingEvidenceInner) GetEntropyOk() (*float64, bool)`

GetEntropyOk returns a tuple with the Entropy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntropy

`func (o *ReportFindingEvidenceInner) SetEntropy(v float64)`

SetEntropy sets Entropy field to given value.


### GetReferences

`func (o *ReportFindingEvidenceInner) GetReferences() []int64`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ReportFindingEvidenceInner) GetReferencesOk() (*[]int64, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ReportFindingEvidenceInner) SetReferences(v []int64)`

SetReferences sets References field to given value.


### GetRuleId

`func (o *ReportFindingEvidenceInner) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *ReportFindingEvidenceInner) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *ReportFindingEvidenceInner) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.


### GetSecret

`func (o *ReportFindingEvidenceInner) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ReportFindingEvidenceInner) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ReportFindingEvidenceInner) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetSecretKind

`func (o *ReportFindingEvidenceInner) GetSecretKind() string`

GetSecretKind returns the SecretKind field if non-nil, zero value otherwise.

### GetSecretKindOk

`func (o *ReportFindingEvidenceInner) GetSecretKindOk() (*string, bool)`

GetSecretKindOk returns a tuple with the SecretKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKind

`func (o *ReportFindingEvidenceInner) SetSecretKind(v string)`

SetSecretKind sets SecretKind field to given value.


### GetSize

`func (o *ReportFindingEvidenceInner) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ReportFindingEvidenceInner) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ReportFindingEvidenceInner) SetSize(v int64)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


