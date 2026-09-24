# EvidenceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvidenceKind** | Pointer to **string** |  | [optional] [default to "hardcoded_secret"]
**Kind** | Pointer to **string** |  | [optional] [default to "deterministic_derivation"]
**Effect** | **string** |  | [default to "supports"]
**Strength** | [**EvidenceStrength**](EvidenceStrength.md) |  | 
**CallChain** | **[]int32** |  | 
**Constants** | [**[]ReferencedConstant**](ReferencedConstant.md) |  | 
**Calls** | [**[]ImportedApiCall**](ImportedApiCall.md) |  | 
**Strings** | [**[]StringMatch**](StringMatch.md) |  | 
**Similarities** | [**[]FunctionSimilarity**](FunctionSimilarity.md) |  | 
**Apis** | [**[]ImportedApi**](ImportedApi.md) |  | 
**Summaries** | [**[]DecompilerSummary**](DecompilerSummary.md) |  | 
**Interpretations** | [**[]ModelInterpretation**](ModelInterpretation.md) |  | 
**Values** | [**[]BytesConstant**](BytesConstant.md) |  | 
**RuleId** | **string** |  | 
**Description** | **string** |  | 
**SecretKind** | [**RuleKind**](RuleKind.md) |  | 
**Secret** | **string** |  | 
**Entropy** | **float32** |  | 
**Size** | **int32** |  | 
**References** | **[]int32** |  | 

## Methods

### NewEvidenceInner

`func NewEvidenceInner(effect string, strength EvidenceStrength, callChain []int32, constants []ReferencedConstant, calls []ImportedApiCall, strings []StringMatch, similarities []FunctionSimilarity, apis []ImportedApi, summaries []DecompilerSummary, interpretations []ModelInterpretation, values []BytesConstant, ruleId string, description string, secretKind RuleKind, secret string, entropy float32, size int32, references []int32, ) *EvidenceInner`

NewEvidenceInner instantiates a new EvidenceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvidenceInnerWithDefaults

`func NewEvidenceInnerWithDefaults() *EvidenceInner`

NewEvidenceInnerWithDefaults instantiates a new EvidenceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### HasEvidenceKind

`func (o *EvidenceInner) HasEvidenceKind() bool`

HasEvidenceKind returns a boolean if a field has been set.

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

### HasKind

`func (o *EvidenceInner) HasKind() bool`

HasKind returns a boolean if a field has been set.

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


### GetStrength

`func (o *EvidenceInner) GetStrength() EvidenceStrength`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *EvidenceInner) GetStrengthOk() (*EvidenceStrength, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *EvidenceInner) SetStrength(v EvidenceStrength)`

SetStrength sets Strength field to given value.


### GetCallChain

`func (o *EvidenceInner) GetCallChain() []int32`

GetCallChain returns the CallChain field if non-nil, zero value otherwise.

### GetCallChainOk

`func (o *EvidenceInner) GetCallChainOk() (*[]int32, bool)`

GetCallChainOk returns a tuple with the CallChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallChain

`func (o *EvidenceInner) SetCallChain(v []int32)`

SetCallChain sets CallChain field to given value.


### GetConstants

`func (o *EvidenceInner) GetConstants() []ReferencedConstant`

GetConstants returns the Constants field if non-nil, zero value otherwise.

### GetConstantsOk

`func (o *EvidenceInner) GetConstantsOk() (*[]ReferencedConstant, bool)`

GetConstantsOk returns a tuple with the Constants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstants

`func (o *EvidenceInner) SetConstants(v []ReferencedConstant)`

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

`func (o *EvidenceInner) GetStrings() []StringMatch`

GetStrings returns the Strings field if non-nil, zero value otherwise.

### GetStringsOk

`func (o *EvidenceInner) GetStringsOk() (*[]StringMatch, bool)`

GetStringsOk returns a tuple with the Strings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrings

`func (o *EvidenceInner) SetStrings(v []StringMatch)`

SetStrings sets Strings field to given value.


### GetSimilarities

`func (o *EvidenceInner) GetSimilarities() []FunctionSimilarity`

GetSimilarities returns the Similarities field if non-nil, zero value otherwise.

### GetSimilaritiesOk

`func (o *EvidenceInner) GetSimilaritiesOk() (*[]FunctionSimilarity, bool)`

GetSimilaritiesOk returns a tuple with the Similarities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimilarities

`func (o *EvidenceInner) SetSimilarities(v []FunctionSimilarity)`

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

`func (o *EvidenceInner) GetSummaries() []DecompilerSummary`

GetSummaries returns the Summaries field if non-nil, zero value otherwise.

### GetSummariesOk

`func (o *EvidenceInner) GetSummariesOk() (*[]DecompilerSummary, bool)`

GetSummariesOk returns a tuple with the Summaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaries

`func (o *EvidenceInner) SetSummaries(v []DecompilerSummary)`

SetSummaries sets Summaries field to given value.


### GetInterpretations

`func (o *EvidenceInner) GetInterpretations() []ModelInterpretation`

GetInterpretations returns the Interpretations field if non-nil, zero value otherwise.

### GetInterpretationsOk

`func (o *EvidenceInner) GetInterpretationsOk() (*[]ModelInterpretation, bool)`

GetInterpretationsOk returns a tuple with the Interpretations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterpretations

`func (o *EvidenceInner) SetInterpretations(v []ModelInterpretation)`

SetInterpretations sets Interpretations field to given value.


### GetValues

`func (o *EvidenceInner) GetValues() []BytesConstant`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *EvidenceInner) GetValuesOk() (*[]BytesConstant, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *EvidenceInner) SetValues(v []BytesConstant)`

SetValues sets Values field to given value.


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


### GetSecretKind

`func (o *EvidenceInner) GetSecretKind() RuleKind`

GetSecretKind returns the SecretKind field if non-nil, zero value otherwise.

### GetSecretKindOk

`func (o *EvidenceInner) GetSecretKindOk() (*RuleKind, bool)`

GetSecretKindOk returns a tuple with the SecretKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKind

`func (o *EvidenceInner) SetSecretKind(v RuleKind)`

SetSecretKind sets SecretKind field to given value.


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


### GetEntropy

`func (o *EvidenceInner) GetEntropy() float32`

GetEntropy returns the Entropy field if non-nil, zero value otherwise.

### GetEntropyOk

`func (o *EvidenceInner) GetEntropyOk() (*float32, bool)`

GetEntropyOk returns a tuple with the Entropy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntropy

`func (o *EvidenceInner) SetEntropy(v float32)`

SetEntropy sets Entropy field to given value.


### GetSize

`func (o *EvidenceInner) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *EvidenceInner) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *EvidenceInner) SetSize(v int32)`

SetSize sets Size field to given value.


### GetReferences

`func (o *EvidenceInner) GetReferences() []int32`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *EvidenceInner) GetReferencesOk() (*[]int32, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *EvidenceInner) SetReferences(v []int32)`

SetReferences sets References field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


