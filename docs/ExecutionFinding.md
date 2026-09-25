# ExecutionFinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Function&#39;s virtual address, hex-encoded | 
**Categories** | **[]string** | Distinct execution categories evidenced by this function | 
**Confidence** | **string** | High when a direct name match was found, medium when the function only calls into execution APIs | 
**DirectMatches** | Pointer to [**[]ExecutionDirectMatch**](ExecutionDirectMatch.md) | Matches against the function&#39;s own name | [optional] 
**EvidenceCount** | **int64** | Total number of direct matches and execution calls | 
**Executes** | **bool** | Whether this function evidences executing code rather than only supporting it | 
**ExecutionCalls** | Pointer to [**[]ExecutionCall**](ExecutionCall.md) | Matches against names this function calls | [optional] 
**FunctionId** | **int64** | ID of the function the finding was reported in | 
**FunctionName** | **string** | Name of the function the finding was reported in | 
**FunctionSize** | **int64** | Size of the function in bytes | 
**Sources** | **[]string** | Distinct execution sources evidenced by this function | 
**Verification** | Pointer to [**ExecutionVerification**](ExecutionVerification.md) | LLM verdict checking this finding against its decompilation. Present only when the run verified this finding. | [optional] 

## Methods

### NewExecutionFinding

`func NewExecutionFinding(address string, categories []string, confidence string, evidenceCount int64, executes bool, functionId int64, functionName string, functionSize int64, sources []string, ) *ExecutionFinding`

NewExecutionFinding instantiates a new ExecutionFinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionFindingWithDefaults

`func NewExecutionFindingWithDefaults() *ExecutionFinding`

NewExecutionFindingWithDefaults instantiates a new ExecutionFinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ExecutionFinding) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ExecutionFinding) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ExecutionFinding) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCategories

`func (o *ExecutionFinding) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ExecutionFinding) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ExecutionFinding) SetCategories(v []string)`

SetCategories sets Categories field to given value.


### SetCategoriesNil

`func (o *ExecutionFinding) SetCategoriesNil(b bool)`

 SetCategoriesNil sets the value for Categories to be an explicit nil

### UnsetCategories
`func (o *ExecutionFinding) UnsetCategories()`

UnsetCategories ensures that no value is present for Categories, not even an explicit nil
### GetConfidence

`func (o *ExecutionFinding) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *ExecutionFinding) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *ExecutionFinding) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.


### GetDirectMatches

`func (o *ExecutionFinding) GetDirectMatches() []ExecutionDirectMatch`

GetDirectMatches returns the DirectMatches field if non-nil, zero value otherwise.

### GetDirectMatchesOk

`func (o *ExecutionFinding) GetDirectMatchesOk() (*[]ExecutionDirectMatch, bool)`

GetDirectMatchesOk returns a tuple with the DirectMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectMatches

`func (o *ExecutionFinding) SetDirectMatches(v []ExecutionDirectMatch)`

SetDirectMatches sets DirectMatches field to given value.

### HasDirectMatches

`func (o *ExecutionFinding) HasDirectMatches() bool`

HasDirectMatches returns a boolean if a field has been set.

### SetDirectMatchesNil

`func (o *ExecutionFinding) SetDirectMatchesNil(b bool)`

 SetDirectMatchesNil sets the value for DirectMatches to be an explicit nil

### UnsetDirectMatches
`func (o *ExecutionFinding) UnsetDirectMatches()`

UnsetDirectMatches ensures that no value is present for DirectMatches, not even an explicit nil
### GetEvidenceCount

`func (o *ExecutionFinding) GetEvidenceCount() int64`

GetEvidenceCount returns the EvidenceCount field if non-nil, zero value otherwise.

### GetEvidenceCountOk

`func (o *ExecutionFinding) GetEvidenceCountOk() (*int64, bool)`

GetEvidenceCountOk returns a tuple with the EvidenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceCount

`func (o *ExecutionFinding) SetEvidenceCount(v int64)`

SetEvidenceCount sets EvidenceCount field to given value.


### GetExecutes

`func (o *ExecutionFinding) GetExecutes() bool`

GetExecutes returns the Executes field if non-nil, zero value otherwise.

### GetExecutesOk

`func (o *ExecutionFinding) GetExecutesOk() (*bool, bool)`

GetExecutesOk returns a tuple with the Executes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutes

`func (o *ExecutionFinding) SetExecutes(v bool)`

SetExecutes sets Executes field to given value.


### GetExecutionCalls

`func (o *ExecutionFinding) GetExecutionCalls() []ExecutionCall`

GetExecutionCalls returns the ExecutionCalls field if non-nil, zero value otherwise.

### GetExecutionCallsOk

`func (o *ExecutionFinding) GetExecutionCallsOk() (*[]ExecutionCall, bool)`

GetExecutionCallsOk returns a tuple with the ExecutionCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCalls

`func (o *ExecutionFinding) SetExecutionCalls(v []ExecutionCall)`

SetExecutionCalls sets ExecutionCalls field to given value.

### HasExecutionCalls

`func (o *ExecutionFinding) HasExecutionCalls() bool`

HasExecutionCalls returns a boolean if a field has been set.

### SetExecutionCallsNil

`func (o *ExecutionFinding) SetExecutionCallsNil(b bool)`

 SetExecutionCallsNil sets the value for ExecutionCalls to be an explicit nil

### UnsetExecutionCalls
`func (o *ExecutionFinding) UnsetExecutionCalls()`

UnsetExecutionCalls ensures that no value is present for ExecutionCalls, not even an explicit nil
### GetFunctionId

`func (o *ExecutionFinding) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *ExecutionFinding) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *ExecutionFinding) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetFunctionName

`func (o *ExecutionFinding) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *ExecutionFinding) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *ExecutionFinding) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetFunctionSize

`func (o *ExecutionFinding) GetFunctionSize() int64`

GetFunctionSize returns the FunctionSize field if non-nil, zero value otherwise.

### GetFunctionSizeOk

`func (o *ExecutionFinding) GetFunctionSizeOk() (*int64, bool)`

GetFunctionSizeOk returns a tuple with the FunctionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionSize

`func (o *ExecutionFinding) SetFunctionSize(v int64)`

SetFunctionSize sets FunctionSize field to given value.


### GetSources

`func (o *ExecutionFinding) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *ExecutionFinding) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *ExecutionFinding) SetSources(v []string)`

SetSources sets Sources field to given value.


### SetSourcesNil

`func (o *ExecutionFinding) SetSourcesNil(b bool)`

 SetSourcesNil sets the value for Sources to be an explicit nil

### UnsetSources
`func (o *ExecutionFinding) UnsetSources()`

UnsetSources ensures that no value is present for Sources, not even an explicit nil
### GetVerification

`func (o *ExecutionFinding) GetVerification() ExecutionVerification`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *ExecutionFinding) GetVerificationOk() (*ExecutionVerification, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *ExecutionFinding) SetVerification(v ExecutionVerification)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *ExecutionFinding) HasVerification() bool`

HasVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


