# ExecutionExplainResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cancelled** | **bool** | Whether the run was cancelled | 
**CodeOrigin** | Pointer to **string** | Where the executed code originates from | [optional] 
**ExecutedTargets** | Pointer to **[]string** | Concrete targets that end up executed -- process names, module paths, or shellcode buffers | [optional] 
**FunctionId** | **int64** | ID of the explained function | 
**FunctionName** | Pointer to **string** | Name of the explained function | [optional] 
**FunctionsInvolved** | Pointer to [**[]ExecutionExplainedFunction**](ExecutionExplainedFunction.md) | Other functions involved in the code execution | [optional] 
**Purpose** | Pointer to **string** | Purpose of the code execution | [optional] 
**Summary** | Pointer to **string** | Explanation of the code execution the function performs | [optional] 
**Trigger** | Pointer to **string** | What triggers the execution | [optional] 
**WhatExecutes** | Pointer to **string** | What code ends up executing | [optional] 

## Methods

### NewExecutionExplainResult

`func NewExecutionExplainResult(cancelled bool, functionId int64, ) *ExecutionExplainResult`

NewExecutionExplainResult instantiates a new ExecutionExplainResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionExplainResultWithDefaults

`func NewExecutionExplainResultWithDefaults() *ExecutionExplainResult`

NewExecutionExplainResultWithDefaults instantiates a new ExecutionExplainResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelled

`func (o *ExecutionExplainResult) GetCancelled() bool`

GetCancelled returns the Cancelled field if non-nil, zero value otherwise.

### GetCancelledOk

`func (o *ExecutionExplainResult) GetCancelledOk() (*bool, bool)`

GetCancelledOk returns a tuple with the Cancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelled

`func (o *ExecutionExplainResult) SetCancelled(v bool)`

SetCancelled sets Cancelled field to given value.


### GetCodeOrigin

`func (o *ExecutionExplainResult) GetCodeOrigin() string`

GetCodeOrigin returns the CodeOrigin field if non-nil, zero value otherwise.

### GetCodeOriginOk

`func (o *ExecutionExplainResult) GetCodeOriginOk() (*string, bool)`

GetCodeOriginOk returns a tuple with the CodeOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeOrigin

`func (o *ExecutionExplainResult) SetCodeOrigin(v string)`

SetCodeOrigin sets CodeOrigin field to given value.

### HasCodeOrigin

`func (o *ExecutionExplainResult) HasCodeOrigin() bool`

HasCodeOrigin returns a boolean if a field has been set.

### GetExecutedTargets

`func (o *ExecutionExplainResult) GetExecutedTargets() []string`

GetExecutedTargets returns the ExecutedTargets field if non-nil, zero value otherwise.

### GetExecutedTargetsOk

`func (o *ExecutionExplainResult) GetExecutedTargetsOk() (*[]string, bool)`

GetExecutedTargetsOk returns a tuple with the ExecutedTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedTargets

`func (o *ExecutionExplainResult) SetExecutedTargets(v []string)`

SetExecutedTargets sets ExecutedTargets field to given value.

### HasExecutedTargets

`func (o *ExecutionExplainResult) HasExecutedTargets() bool`

HasExecutedTargets returns a boolean if a field has been set.

### SetExecutedTargetsNil

`func (o *ExecutionExplainResult) SetExecutedTargetsNil(b bool)`

 SetExecutedTargetsNil sets the value for ExecutedTargets to be an explicit nil

### UnsetExecutedTargets
`func (o *ExecutionExplainResult) UnsetExecutedTargets()`

UnsetExecutedTargets ensures that no value is present for ExecutedTargets, not even an explicit nil
### GetFunctionId

`func (o *ExecutionExplainResult) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *ExecutionExplainResult) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *ExecutionExplainResult) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetFunctionName

`func (o *ExecutionExplainResult) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *ExecutionExplainResult) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *ExecutionExplainResult) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.

### HasFunctionName

`func (o *ExecutionExplainResult) HasFunctionName() bool`

HasFunctionName returns a boolean if a field has been set.

### GetFunctionsInvolved

`func (o *ExecutionExplainResult) GetFunctionsInvolved() []ExecutionExplainedFunction`

GetFunctionsInvolved returns the FunctionsInvolved field if non-nil, zero value otherwise.

### GetFunctionsInvolvedOk

`func (o *ExecutionExplainResult) GetFunctionsInvolvedOk() (*[]ExecutionExplainedFunction, bool)`

GetFunctionsInvolvedOk returns a tuple with the FunctionsInvolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionsInvolved

`func (o *ExecutionExplainResult) SetFunctionsInvolved(v []ExecutionExplainedFunction)`

SetFunctionsInvolved sets FunctionsInvolved field to given value.

### HasFunctionsInvolved

`func (o *ExecutionExplainResult) HasFunctionsInvolved() bool`

HasFunctionsInvolved returns a boolean if a field has been set.

### SetFunctionsInvolvedNil

`func (o *ExecutionExplainResult) SetFunctionsInvolvedNil(b bool)`

 SetFunctionsInvolvedNil sets the value for FunctionsInvolved to be an explicit nil

### UnsetFunctionsInvolved
`func (o *ExecutionExplainResult) UnsetFunctionsInvolved()`

UnsetFunctionsInvolved ensures that no value is present for FunctionsInvolved, not even an explicit nil
### GetPurpose

`func (o *ExecutionExplainResult) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *ExecutionExplainResult) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *ExecutionExplainResult) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *ExecutionExplainResult) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetSummary

`func (o *ExecutionExplainResult) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ExecutionExplainResult) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ExecutionExplainResult) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ExecutionExplainResult) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTrigger

`func (o *ExecutionExplainResult) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *ExecutionExplainResult) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *ExecutionExplainResult) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *ExecutionExplainResult) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetWhatExecutes

`func (o *ExecutionExplainResult) GetWhatExecutes() string`

GetWhatExecutes returns the WhatExecutes field if non-nil, zero value otherwise.

### GetWhatExecutesOk

`func (o *ExecutionExplainResult) GetWhatExecutesOk() (*string, bool)`

GetWhatExecutesOk returns a tuple with the WhatExecutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhatExecutes

`func (o *ExecutionExplainResult) SetWhatExecutes(v string)`

SetWhatExecutes sets WhatExecutes field to given value.

### HasWhatExecutes

`func (o *ExecutionExplainResult) HasWhatExecutes() bool`

HasWhatExecutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


