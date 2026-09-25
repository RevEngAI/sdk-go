# CryptoFinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Function&#39;s virtual address, hex-encoded | 
**Categories** | **[]string** | Distinct crypto categories evidenced by this function | 
**Confidence** | **string** | High when a direct name match was found, medium when the function only calls into crypto APIs | 
**CryptoCalls** | Pointer to [**[]CryptoCall**](CryptoCall.md) | Matches against names this function calls | [optional] 
**DirectMatches** | Pointer to [**[]CryptoDirectMatch**](CryptoDirectMatch.md) | Matches against the function&#39;s own name | [optional] 
**EvidenceCount** | **int64** | Total number of direct matches and crypto calls | 
**FunctionId** | **int64** | ID of the function the finding was reported in | 
**FunctionName** | **string** | Name of the function the finding was reported in | 
**FunctionSize** | **int64** | Size of the function in bytes | 
**Libraries** | **[]string** | Distinct crypto libraries evidenced by this function | 
**Verification** | Pointer to [**CryptoVerification**](CryptoVerification.md) | LLM verdict checking this finding against its decompilation. Present only when the run verified this finding. | [optional] 

## Methods

### NewCryptoFinding

`func NewCryptoFinding(address string, categories []string, confidence string, evidenceCount int64, functionId int64, functionName string, functionSize int64, libraries []string, ) *CryptoFinding`

NewCryptoFinding instantiates a new CryptoFinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptoFindingWithDefaults

`func NewCryptoFindingWithDefaults() *CryptoFinding`

NewCryptoFindingWithDefaults instantiates a new CryptoFinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CryptoFinding) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CryptoFinding) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CryptoFinding) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCategories

`func (o *CryptoFinding) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *CryptoFinding) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *CryptoFinding) SetCategories(v []string)`

SetCategories sets Categories field to given value.


### SetCategoriesNil

`func (o *CryptoFinding) SetCategoriesNil(b bool)`

 SetCategoriesNil sets the value for Categories to be an explicit nil

### UnsetCategories
`func (o *CryptoFinding) UnsetCategories()`

UnsetCategories ensures that no value is present for Categories, not even an explicit nil
### GetConfidence

`func (o *CryptoFinding) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *CryptoFinding) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *CryptoFinding) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.


### GetCryptoCalls

`func (o *CryptoFinding) GetCryptoCalls() []CryptoCall`

GetCryptoCalls returns the CryptoCalls field if non-nil, zero value otherwise.

### GetCryptoCallsOk

`func (o *CryptoFinding) GetCryptoCallsOk() (*[]CryptoCall, bool)`

GetCryptoCallsOk returns a tuple with the CryptoCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoCalls

`func (o *CryptoFinding) SetCryptoCalls(v []CryptoCall)`

SetCryptoCalls sets CryptoCalls field to given value.

### HasCryptoCalls

`func (o *CryptoFinding) HasCryptoCalls() bool`

HasCryptoCalls returns a boolean if a field has been set.

### SetCryptoCallsNil

`func (o *CryptoFinding) SetCryptoCallsNil(b bool)`

 SetCryptoCallsNil sets the value for CryptoCalls to be an explicit nil

### UnsetCryptoCalls
`func (o *CryptoFinding) UnsetCryptoCalls()`

UnsetCryptoCalls ensures that no value is present for CryptoCalls, not even an explicit nil
### GetDirectMatches

`func (o *CryptoFinding) GetDirectMatches() []CryptoDirectMatch`

GetDirectMatches returns the DirectMatches field if non-nil, zero value otherwise.

### GetDirectMatchesOk

`func (o *CryptoFinding) GetDirectMatchesOk() (*[]CryptoDirectMatch, bool)`

GetDirectMatchesOk returns a tuple with the DirectMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectMatches

`func (o *CryptoFinding) SetDirectMatches(v []CryptoDirectMatch)`

SetDirectMatches sets DirectMatches field to given value.

### HasDirectMatches

`func (o *CryptoFinding) HasDirectMatches() bool`

HasDirectMatches returns a boolean if a field has been set.

### SetDirectMatchesNil

`func (o *CryptoFinding) SetDirectMatchesNil(b bool)`

 SetDirectMatchesNil sets the value for DirectMatches to be an explicit nil

### UnsetDirectMatches
`func (o *CryptoFinding) UnsetDirectMatches()`

UnsetDirectMatches ensures that no value is present for DirectMatches, not even an explicit nil
### GetEvidenceCount

`func (o *CryptoFinding) GetEvidenceCount() int64`

GetEvidenceCount returns the EvidenceCount field if non-nil, zero value otherwise.

### GetEvidenceCountOk

`func (o *CryptoFinding) GetEvidenceCountOk() (*int64, bool)`

GetEvidenceCountOk returns a tuple with the EvidenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceCount

`func (o *CryptoFinding) SetEvidenceCount(v int64)`

SetEvidenceCount sets EvidenceCount field to given value.


### GetFunctionId

`func (o *CryptoFinding) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *CryptoFinding) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *CryptoFinding) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetFunctionName

`func (o *CryptoFinding) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *CryptoFinding) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *CryptoFinding) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetFunctionSize

`func (o *CryptoFinding) GetFunctionSize() int64`

GetFunctionSize returns the FunctionSize field if non-nil, zero value otherwise.

### GetFunctionSizeOk

`func (o *CryptoFinding) GetFunctionSizeOk() (*int64, bool)`

GetFunctionSizeOk returns a tuple with the FunctionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionSize

`func (o *CryptoFinding) SetFunctionSize(v int64)`

SetFunctionSize sets FunctionSize field to given value.


### GetLibraries

`func (o *CryptoFinding) GetLibraries() []string`

GetLibraries returns the Libraries field if non-nil, zero value otherwise.

### GetLibrariesOk

`func (o *CryptoFinding) GetLibrariesOk() (*[]string, bool)`

GetLibrariesOk returns a tuple with the Libraries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraries

`func (o *CryptoFinding) SetLibraries(v []string)`

SetLibraries sets Libraries field to given value.


### SetLibrariesNil

`func (o *CryptoFinding) SetLibrariesNil(b bool)`

 SetLibrariesNil sets the value for Libraries to be an explicit nil

### UnsetLibraries
`func (o *CryptoFinding) UnsetLibraries()`

UnsetLibraries ensures that no value is present for Libraries, not even an explicit nil
### GetVerification

`func (o *CryptoFinding) GetVerification() CryptoVerification`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *CryptoFinding) GetVerificationOk() (*CryptoVerification, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *CryptoFinding) SetVerification(v CryptoVerification)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *CryptoFinding) HasVerification() bool`

HasVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


