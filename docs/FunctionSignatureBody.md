# FunctionSignatureBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallingConvention** | Pointer to **string** | Calling convention, when the producer reported one. | [optional] 
**CreatedAt** | Pointer to **time.Time** | When this signature was extracted. | [optional] 
**DataTypes** | Pointer to [**[]DataTypeEntry**](DataTypeEntry.md) | The types the signature names — its parameter types and its return type — ordered by data_type_id. Each entry is identical to the one the data types endpoints serve for that id. Returned only when include_data_types is true. | [optional] 
**FunctionId** | **int64** |  | 
**FunctionName** | **string** | Current name of the function. | 
**HasSignature** | **bool** | Whether a signature was extracted for this function. False is a normal result: no signature is recorded unless data type extraction ran for the analysis, and thunks and external functions are skipped when it does. | 
**Parameters** | [**[]SignatureParameterEntry**](SignatureParameterEntry.md) | Parameters in argument order. Empty with has_signature true means the function is known to take no arguments. | 
**ReturnDataTypeId** | Pointer to **int64** | Return type, resolvable against the analysis data types list. Absent for an unresolved return type. | [optional] 
**SourceFunctionId** | Pointer to **int64** | The function this signature was copied from, when it was transferred rather than extracted. | [optional] 
**SourceType** | Pointer to **string** | Where this signature came from. | [optional] 

## Methods

### NewFunctionSignatureBody

`func NewFunctionSignatureBody(functionId int64, functionName string, hasSignature bool, parameters []SignatureParameterEntry, ) *FunctionSignatureBody`

NewFunctionSignatureBody instantiates a new FunctionSignatureBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionSignatureBodyWithDefaults

`func NewFunctionSignatureBodyWithDefaults() *FunctionSignatureBody`

NewFunctionSignatureBodyWithDefaults instantiates a new FunctionSignatureBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallingConvention

`func (o *FunctionSignatureBody) GetCallingConvention() string`

GetCallingConvention returns the CallingConvention field if non-nil, zero value otherwise.

### GetCallingConventionOk

`func (o *FunctionSignatureBody) GetCallingConventionOk() (*string, bool)`

GetCallingConventionOk returns a tuple with the CallingConvention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallingConvention

`func (o *FunctionSignatureBody) SetCallingConvention(v string)`

SetCallingConvention sets CallingConvention field to given value.

### HasCallingConvention

`func (o *FunctionSignatureBody) HasCallingConvention() bool`

HasCallingConvention returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FunctionSignatureBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FunctionSignatureBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FunctionSignatureBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FunctionSignatureBody) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDataTypes

`func (o *FunctionSignatureBody) GetDataTypes() []DataTypeEntry`

GetDataTypes returns the DataTypes field if non-nil, zero value otherwise.

### GetDataTypesOk

`func (o *FunctionSignatureBody) GetDataTypesOk() (*[]DataTypeEntry, bool)`

GetDataTypesOk returns a tuple with the DataTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypes

`func (o *FunctionSignatureBody) SetDataTypes(v []DataTypeEntry)`

SetDataTypes sets DataTypes field to given value.

### HasDataTypes

`func (o *FunctionSignatureBody) HasDataTypes() bool`

HasDataTypes returns a boolean if a field has been set.

### SetDataTypesNil

`func (o *FunctionSignatureBody) SetDataTypesNil(b bool)`

 SetDataTypesNil sets the value for DataTypes to be an explicit nil

### UnsetDataTypes
`func (o *FunctionSignatureBody) UnsetDataTypes()`

UnsetDataTypes ensures that no value is present for DataTypes, not even an explicit nil
### GetFunctionId

`func (o *FunctionSignatureBody) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *FunctionSignatureBody) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *FunctionSignatureBody) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetFunctionName

`func (o *FunctionSignatureBody) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *FunctionSignatureBody) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *FunctionSignatureBody) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetHasSignature

`func (o *FunctionSignatureBody) GetHasSignature() bool`

GetHasSignature returns the HasSignature field if non-nil, zero value otherwise.

### GetHasSignatureOk

`func (o *FunctionSignatureBody) GetHasSignatureOk() (*bool, bool)`

GetHasSignatureOk returns a tuple with the HasSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSignature

`func (o *FunctionSignatureBody) SetHasSignature(v bool)`

SetHasSignature sets HasSignature field to given value.


### GetParameters

`func (o *FunctionSignatureBody) GetParameters() []SignatureParameterEntry`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *FunctionSignatureBody) GetParametersOk() (*[]SignatureParameterEntry, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *FunctionSignatureBody) SetParameters(v []SignatureParameterEntry)`

SetParameters sets Parameters field to given value.


### SetParametersNil

`func (o *FunctionSignatureBody) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *FunctionSignatureBody) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetReturnDataTypeId

`func (o *FunctionSignatureBody) GetReturnDataTypeId() int64`

GetReturnDataTypeId returns the ReturnDataTypeId field if non-nil, zero value otherwise.

### GetReturnDataTypeIdOk

`func (o *FunctionSignatureBody) GetReturnDataTypeIdOk() (*int64, bool)`

GetReturnDataTypeIdOk returns a tuple with the ReturnDataTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnDataTypeId

`func (o *FunctionSignatureBody) SetReturnDataTypeId(v int64)`

SetReturnDataTypeId sets ReturnDataTypeId field to given value.

### HasReturnDataTypeId

`func (o *FunctionSignatureBody) HasReturnDataTypeId() bool`

HasReturnDataTypeId returns a boolean if a field has been set.

### GetSourceFunctionId

`func (o *FunctionSignatureBody) GetSourceFunctionId() int64`

GetSourceFunctionId returns the SourceFunctionId field if non-nil, zero value otherwise.

### GetSourceFunctionIdOk

`func (o *FunctionSignatureBody) GetSourceFunctionIdOk() (*int64, bool)`

GetSourceFunctionIdOk returns a tuple with the SourceFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFunctionId

`func (o *FunctionSignatureBody) SetSourceFunctionId(v int64)`

SetSourceFunctionId sets SourceFunctionId field to given value.

### HasSourceFunctionId

`func (o *FunctionSignatureBody) HasSourceFunctionId() bool`

HasSourceFunctionId returns a boolean if a field has been set.

### GetSourceType

`func (o *FunctionSignatureBody) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *FunctionSignatureBody) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *FunctionSignatureBody) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *FunctionSignatureBody) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


