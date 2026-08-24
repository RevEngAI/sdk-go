# UpdateFunctionSignatureInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallingConvention** | Pointer to **string** | Calling convention. Omit when there is none to record. | [optional] 
**Parameters** | [**[]SignatureParameterInput**](SignatureParameterInput.md) | Parameters in argument order. An empty list records a function that takes no arguments. | 
**ReturnDataTypeId** | Pointer to **int64** | Return type, which must belong to this analysis. Omit for an unresolved return type. | [optional] 

## Methods

### NewUpdateFunctionSignatureInputBody

`func NewUpdateFunctionSignatureInputBody(parameters []SignatureParameterInput, ) *UpdateFunctionSignatureInputBody`

NewUpdateFunctionSignatureInputBody instantiates a new UpdateFunctionSignatureInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFunctionSignatureInputBodyWithDefaults

`func NewUpdateFunctionSignatureInputBodyWithDefaults() *UpdateFunctionSignatureInputBody`

NewUpdateFunctionSignatureInputBodyWithDefaults instantiates a new UpdateFunctionSignatureInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallingConvention

`func (o *UpdateFunctionSignatureInputBody) GetCallingConvention() string`

GetCallingConvention returns the CallingConvention field if non-nil, zero value otherwise.

### GetCallingConventionOk

`func (o *UpdateFunctionSignatureInputBody) GetCallingConventionOk() (*string, bool)`

GetCallingConventionOk returns a tuple with the CallingConvention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallingConvention

`func (o *UpdateFunctionSignatureInputBody) SetCallingConvention(v string)`

SetCallingConvention sets CallingConvention field to given value.

### HasCallingConvention

`func (o *UpdateFunctionSignatureInputBody) HasCallingConvention() bool`

HasCallingConvention returns a boolean if a field has been set.

### GetParameters

`func (o *UpdateFunctionSignatureInputBody) GetParameters() []SignatureParameterInput`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *UpdateFunctionSignatureInputBody) GetParametersOk() (*[]SignatureParameterInput, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *UpdateFunctionSignatureInputBody) SetParameters(v []SignatureParameterInput)`

SetParameters sets Parameters field to given value.


### SetParametersNil

`func (o *UpdateFunctionSignatureInputBody) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *UpdateFunctionSignatureInputBody) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetReturnDataTypeId

`func (o *UpdateFunctionSignatureInputBody) GetReturnDataTypeId() int64`

GetReturnDataTypeId returns the ReturnDataTypeId field if non-nil, zero value otherwise.

### GetReturnDataTypeIdOk

`func (o *UpdateFunctionSignatureInputBody) GetReturnDataTypeIdOk() (*int64, bool)`

GetReturnDataTypeIdOk returns a tuple with the ReturnDataTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnDataTypeId

`func (o *UpdateFunctionSignatureInputBody) SetReturnDataTypeId(v int64)`

SetReturnDataTypeId sets ReturnDataTypeId field to given value.

### HasReturnDataTypeId

`func (o *UpdateFunctionSignatureInputBody) HasReturnDataTypeId() bool`

HasReturnDataTypeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


