# EvidenceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallChain** | **[]int64** | Addresses along the call chain. | 
**Effect** | **string** | Whether the evidence supports or refutes the finding. | 
**EvidenceKind** | **string** | Kind of evidence. | 
**Kind** | **string** | How the evidence was derived. | 
**Strength** | **string** | How directly the evidence bears on the finding. | 
**Constants** | [**[]ReportReferencedConstant**](ReportReferencedConstant.md) | The constants and where they are referenced. | 
**Calls** | [**[]ImportedApiCall**](ImportedApiCall.md) |  | 
**Strings** | [**[]ReportStringMatch**](ReportStringMatch.md) | The matched strings. | 
**Similarities** | [**[]ReportFunctionSimilarity**](ReportFunctionSimilarity.md) | The similar functions. | 
**Apis** | [**[]ImportedApi**](ImportedApi.md) |  | 
**Summaries** | [**[]ReportCodeSummary**](ReportCodeSummary.md) | Summaries of the decompiled code. | 
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

### NewEvidenceInner

`func NewEvidenceInner(callChain []int64, effect string, evidenceKind string, kind string, strength string, constants []ReportReferencedConstant, calls []ImportedApiCall, strings []ReportStringMatch, similarities []ReportFunctionSimilarity, apis []ImportedApi, summaries []ReportCodeSummary, interpretations []ReportModelInterpretation, values []ReportBytesConstant, description string, entropy float64, references []int64, ruleId string, secret string, secretKind string, size int64, ) *EvidenceInner`

NewEvidenceInner instantiates a new EvidenceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvidenceInnerWithDefaults

`func NewEvidenceInnerWithDefaults() *EvidenceInner`

NewEvidenceInnerWithDefaults instantiates a new EvidenceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallChain

`func (o *EvidenceInner) GetCallChain() []int64`

GetCallChain returns the CallChain field if non-nil, zero value otherwise.

### GetCallChainOk

`func (o *EvidenceInner) GetCallChainOk() (*[]int64, bool)`

GetCallChainOk returns a tuple with the CallChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallChain

`func (o *EvidenceInner) SetCallChain(v []int64)`

SetCallChain sets CallChain field to given value.


### GetEffect

`func (o *EvidenceInner) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *EvidenceInner) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *EvidenceInner) SetEffect(v string)`

SetEffect sets Effect field to given value.


### GetEvidenceKind

`func (o *EvidenceInner) GetEvidenceKind() string`

GetEvidenceKind returns the EvidenceKind field if non-nil, zero value otherwise.

### GetEvidenceKindOk

`func (o *EvidenceInner) GetEvidenceKindOk() (*string, bool)`

GetEvidenceKindOk returns a tuple with the EvidenceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceKind

`func (o *EvidenceInner) SetEvidenceKind(v string)`

SetEvidenceKind sets EvidenceKind field to given value.


### GetKind

`func (o *EvidenceInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EvidenceInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EvidenceInner) SetKind(v string)`

SetKind sets Kind field to given value.


### GetStrength

`func (o *EvidenceInner) GetStrength() string`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *EvidenceInner) GetStrengthOk() (*string, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *EvidenceInner) SetStrength(v string)`

SetStrength sets Strength field to given value.


### GetConstants

`func (o *EvidenceInner) GetConstants() []ReportReferencedConstant`

GetConstants returns the Constants field if non-nil, zero value otherwise.

### GetConstantsOk

`func (o *EvidenceInner) GetConstantsOk() (*[]ReportReferencedConstant, bool)`

GetConstantsOk returns a tuple with the Constants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstants

`func (o *EvidenceInner) SetConstants(v []ReportReferencedConstant)`

SetConstants sets Constants field to given value.


### GetCalls

`func (o *EvidenceInner) GetCalls() []ImportedApiCall`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *EvidenceInner) GetCallsOk() (*[]ImportedApiCall, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *EvidenceInner) SetCalls(v []ImportedApiCall)`

SetCalls sets Calls field to given value.


### GetStrings

`func (o *EvidenceInner) GetStrings() []ReportStringMatch`

GetStrings returns the Strings field if non-nil, zero value otherwise.

### GetStringsOk

`func (o *EvidenceInner) GetStringsOk() (*[]ReportStringMatch, bool)`

GetStringsOk returns a tuple with the Strings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrings

`func (o *EvidenceInner) SetStrings(v []ReportStringMatch)`

SetStrings sets Strings field to given value.


### GetSimilarities

`func (o *EvidenceInner) GetSimilarities() []ReportFunctionSimilarity`

GetSimilarities returns the Similarities field if non-nil, zero value otherwise.

### GetSimilaritiesOk

`func (o *EvidenceInner) GetSimilaritiesOk() (*[]ReportFunctionSimilarity, bool)`

GetSimilaritiesOk returns a tuple with the Similarities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimilarities

`func (o *EvidenceInner) SetSimilarities(v []ReportFunctionSimilarity)`

SetSimilarities sets Similarities field to given value.


### GetApis

`func (o *EvidenceInner) GetApis() []ImportedApi`

GetApis returns the Apis field if non-nil, zero value otherwise.

### GetApisOk

`func (o *EvidenceInner) GetApisOk() (*[]ImportedApi, bool)`

GetApisOk returns a tuple with the Apis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApis

`func (o *EvidenceInner) SetApis(v []ImportedApi)`

SetApis sets Apis field to given value.


### GetSummaries

`func (o *EvidenceInner) GetSummaries() []ReportCodeSummary`

GetSummaries returns the Summaries field if non-nil, zero value otherwise.

### GetSummariesOk

`func (o *EvidenceInner) GetSummariesOk() (*[]ReportCodeSummary, bool)`

GetSummariesOk returns a tuple with the Summaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaries

`func (o *EvidenceInner) SetSummaries(v []ReportCodeSummary)`

SetSummaries sets Summaries field to given value.


### GetInterpretations

`func (o *EvidenceInner) GetInterpretations() []ReportModelInterpretation`

GetInterpretations returns the Interpretations field if non-nil, zero value otherwise.

### GetInterpretationsOk

`func (o *EvidenceInner) GetInterpretationsOk() (*[]ReportModelInterpretation, bool)`

GetInterpretationsOk returns a tuple with the Interpretations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterpretations

`func (o *EvidenceInner) SetInterpretations(v []ReportModelInterpretation)`

SetInterpretations sets Interpretations field to given value.


### GetValues

`func (o *EvidenceInner) GetValues() []ReportBytesConstant`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *EvidenceInner) GetValuesOk() (*[]ReportBytesConstant, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *EvidenceInner) SetValues(v []ReportBytesConstant)`

SetValues sets Values field to given value.


### GetDescription

`func (o *EvidenceInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EvidenceInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EvidenceInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEntropy

`func (o *EvidenceInner) GetEntropy() float64`

GetEntropy returns the Entropy field if non-nil, zero value otherwise.

### GetEntropyOk

`func (o *EvidenceInner) GetEntropyOk() (*float64, bool)`

GetEntropyOk returns a tuple with the Entropy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntropy

`func (o *EvidenceInner) SetEntropy(v float64)`

SetEntropy sets Entropy field to given value.


### GetReferences

`func (o *EvidenceInner) GetReferences() []int64`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *EvidenceInner) GetReferencesOk() (*[]int64, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *EvidenceInner) SetReferences(v []int64)`

SetReferences sets References field to given value.


### GetRuleId

`func (o *EvidenceInner) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *EvidenceInner) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *EvidenceInner) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.


### GetSecret

`func (o *EvidenceInner) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *EvidenceInner) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *EvidenceInner) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetSecretKind

`func (o *EvidenceInner) GetSecretKind() string`

GetSecretKind returns the SecretKind field if non-nil, zero value otherwise.

### GetSecretKindOk

`func (o *EvidenceInner) GetSecretKindOk() (*string, bool)`

GetSecretKindOk returns a tuple with the SecretKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKind

`func (o *EvidenceInner) SetSecretKind(v string)`

SetSecretKind sets SecretKind field to given value.


### GetSize

`func (o *EvidenceInner) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *EvidenceInner) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *EvidenceInner) SetSize(v int64)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


