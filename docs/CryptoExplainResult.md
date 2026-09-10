# CryptoExplainResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cancelled** | **bool** | Whether the run was cancelled | 
**Error** | Pointer to **string** | Why no explanation could be produced. Empty when the run succeeded. | [optional] 
**FunctionId** | **int64** | ID of the explained function | 
**FunctionName** | Pointer to **string** | Name of the explained function | [optional] 
**FunctionsInvolved** | Pointer to [**[]CryptoExplainedFunction**](CryptoExplainedFunction.md) | Other functions involved in the cryptographic operation | [optional] 
**Summary** | Pointer to **string** | Explanation of the cryptography the function performs | [optional] 

## Methods

### NewCryptoExplainResult

`func NewCryptoExplainResult(cancelled bool, functionId int64, ) *CryptoExplainResult`

NewCryptoExplainResult instantiates a new CryptoExplainResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptoExplainResultWithDefaults

`func NewCryptoExplainResultWithDefaults() *CryptoExplainResult`

NewCryptoExplainResultWithDefaults instantiates a new CryptoExplainResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelled

`func (o *CryptoExplainResult) GetCancelled() bool`

GetCancelled returns the Cancelled field if non-nil, zero value otherwise.

### GetCancelledOk

`func (o *CryptoExplainResult) GetCancelledOk() (*bool, bool)`

GetCancelledOk returns a tuple with the Cancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelled

`func (o *CryptoExplainResult) SetCancelled(v bool)`

SetCancelled sets Cancelled field to given value.


### GetError

`func (o *CryptoExplainResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CryptoExplainResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CryptoExplainResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CryptoExplainResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFunctionId

`func (o *CryptoExplainResult) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *CryptoExplainResult) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *CryptoExplainResult) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetFunctionName

`func (o *CryptoExplainResult) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *CryptoExplainResult) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *CryptoExplainResult) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.

### HasFunctionName

`func (o *CryptoExplainResult) HasFunctionName() bool`

HasFunctionName returns a boolean if a field has been set.

### GetFunctionsInvolved

`func (o *CryptoExplainResult) GetFunctionsInvolved() []CryptoExplainedFunction`

GetFunctionsInvolved returns the FunctionsInvolved field if non-nil, zero value otherwise.

### GetFunctionsInvolvedOk

`func (o *CryptoExplainResult) GetFunctionsInvolvedOk() (*[]CryptoExplainedFunction, bool)`

GetFunctionsInvolvedOk returns a tuple with the FunctionsInvolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionsInvolved

`func (o *CryptoExplainResult) SetFunctionsInvolved(v []CryptoExplainedFunction)`

SetFunctionsInvolved sets FunctionsInvolved field to given value.

### HasFunctionsInvolved

`func (o *CryptoExplainResult) HasFunctionsInvolved() bool`

HasFunctionsInvolved returns a boolean if a field has been set.

### SetFunctionsInvolvedNil

`func (o *CryptoExplainResult) SetFunctionsInvolvedNil(b bool)`

 SetFunctionsInvolvedNil sets the value for FunctionsInvolved to be an explicit nil

### UnsetFunctionsInvolved
`func (o *CryptoExplainResult) UnsetFunctionsInvolved()`

UnsetFunctionsInvolved ensures that no value is present for FunctionsInvolved, not even an explicit nil
### GetSummary

`func (o *CryptoExplainResult) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *CryptoExplainResult) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *CryptoExplainResult) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *CryptoExplainResult) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


