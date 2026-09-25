# CryptoVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Confidence** | Pointer to **string** | LLM&#39;s confidence in the verdict. Absent when verified is null. | [optional] 
**Reasoning** | **string** | LLM&#39;s explanation for the verdict, or the reason verification could not be completed | 
**Verified** | **NullableBool** | Whether an LLM confirmed the finding against its decompilation; null if verification could not be completed, in which case the finding is kept unverified rather than dropped | 

## Methods

### NewCryptoVerification

`func NewCryptoVerification(reasoning string, verified NullableBool, ) *CryptoVerification`

NewCryptoVerification instantiates a new CryptoVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptoVerificationWithDefaults

`func NewCryptoVerificationWithDefaults() *CryptoVerification`

NewCryptoVerificationWithDefaults instantiates a new CryptoVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfidence

`func (o *CryptoVerification) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *CryptoVerification) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *CryptoVerification) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *CryptoVerification) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetReasoning

`func (o *CryptoVerification) GetReasoning() string`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *CryptoVerification) GetReasoningOk() (*string, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *CryptoVerification) SetReasoning(v string)`

SetReasoning sets Reasoning field to given value.


### GetVerified

`func (o *CryptoVerification) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *CryptoVerification) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *CryptoVerification) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### SetVerifiedNil

`func (o *CryptoVerification) SetVerifiedNil(b bool)`

 SetVerifiedNil sets the value for Verified to be an explicit nil

### UnsetVerified
`func (o *CryptoVerification) UnsetVerified()`

UnsetVerified ensures that no value is present for Verified, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


