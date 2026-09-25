# ExecutionVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Confidence** | Pointer to **string** | LLM&#39;s confidence in the verdict. Absent when verified is null. | [optional] 
**Reasoning** | **string** | LLM&#39;s explanation for the verdict, or the reason verification could not be completed | 
**Verified** | **NullableBool** | Whether an LLM confirmed the finding against its decompilation; null if verification could not be completed, in which case the finding is kept unverified rather than dropped | 

## Methods

### NewExecutionVerification

`func NewExecutionVerification(reasoning string, verified NullableBool, ) *ExecutionVerification`

NewExecutionVerification instantiates a new ExecutionVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionVerificationWithDefaults

`func NewExecutionVerificationWithDefaults() *ExecutionVerification`

NewExecutionVerificationWithDefaults instantiates a new ExecutionVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfidence

`func (o *ExecutionVerification) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *ExecutionVerification) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *ExecutionVerification) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *ExecutionVerification) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetReasoning

`func (o *ExecutionVerification) GetReasoning() string`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *ExecutionVerification) GetReasoningOk() (*string, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *ExecutionVerification) SetReasoning(v string)`

SetReasoning sets Reasoning field to given value.


### GetVerified

`func (o *ExecutionVerification) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *ExecutionVerification) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *ExecutionVerification) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### SetVerifiedNil

`func (o *ExecutionVerification) SetVerifiedNil(b bool)`

 SetVerifiedNil sets the value for Verified to be an explicit nil

### UnsetVerified
`func (o *ExecutionVerification) UnsetVerified()`

UnsetVerified ensures that no value is present for Verified, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


