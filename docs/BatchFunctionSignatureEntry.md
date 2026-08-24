# BatchFunctionSignatureEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisId** | **int64** | The analysis this signature belongs to. Its data_type_ids resolve against this analysis only. | 
**CallingConvention** | Pointer to **string** | Calling convention, when the producer reported one. | [optional] 
**CreatedAt** | Pointer to **time.Time** | When this signature was extracted. | [optional] 
**FunctionId** | **int64** |  | 
**FunctionName** | **string** | Current name of the function. | 
**HasSignature** | **bool** | Whether a signature was extracted for this function. False is a normal result: no signature is recorded unless data type extraction ran for the analysis, and thunks and external functions are skipped when it does. | 
**Parameters** | [**[]SignatureParameterEntry**](SignatureParameterEntry.md) | Parameters in argument order. Empty with has_signature true means the function is known to take no arguments. | 
**ReturnDataTypeId** | Pointer to **int64** | Return type, resolvable against the analysis data types list. Absent for an unresolved return type. | [optional] 
**SourceFunctionId** | Pointer to **int64** | The function this signature was copied from, when it was transferred rather than extracted. | [optional] 
**SourceType** | Pointer to **string** | Where this signature came from. | [optional] 

## Methods

### NewBatchFunctionSignatureEntry

`func NewBatchFunctionSignatureEntry(analysisId int64, functionId int64, functionName string, hasSignature bool, parameters []SignatureParameterEntry, ) *BatchFunctionSignatureEntry`

NewBatchFunctionSignatureEntry instantiates a new BatchFunctionSignatureEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchFunctionSignatureEntryWithDefaults

`func NewBatchFunctionSignatureEntryWithDefaults() *BatchFunctionSignatureEntry`

NewBatchFunctionSignatureEntryWithDefaults instantiates a new BatchFunctionSignatureEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisId

`func (o *BatchFunctionSignatureEntry) GetAnalysisId() int64`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *BatchFunctionSignatureEntry) GetAnalysisIdOk() (*int64, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *BatchFunctionSignatureEntry) SetAnalysisId(v int64)`

SetAnalysisId sets AnalysisId field to given value.


### GetCallingConvention

`func (o *BatchFunctionSignatureEntry) GetCallingConvention() string`

GetCallingConvention returns the CallingConvention field if non-nil, zero value otherwise.

### GetCallingConventionOk

`func (o *BatchFunctionSignatureEntry) GetCallingConventionOk() (*string, bool)`

GetCallingConventionOk returns a tuple with the CallingConvention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallingConvention

`func (o *BatchFunctionSignatureEntry) SetCallingConvention(v string)`

SetCallingConvention sets CallingConvention field to given value.

### HasCallingConvention

`func (o *BatchFunctionSignatureEntry) HasCallingConvention() bool`

HasCallingConvention returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BatchFunctionSignatureEntry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BatchFunctionSignatureEntry) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BatchFunctionSignatureEntry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BatchFunctionSignatureEntry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFunctionId

`func (o *BatchFunctionSignatureEntry) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *BatchFunctionSignatureEntry) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *BatchFunctionSignatureEntry) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetFunctionName

`func (o *BatchFunctionSignatureEntry) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *BatchFunctionSignatureEntry) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *BatchFunctionSignatureEntry) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetHasSignature

`func (o *BatchFunctionSignatureEntry) GetHasSignature() bool`

GetHasSignature returns the HasSignature field if non-nil, zero value otherwise.

### GetHasSignatureOk

`func (o *BatchFunctionSignatureEntry) GetHasSignatureOk() (*bool, bool)`

GetHasSignatureOk returns a tuple with the HasSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSignature

`func (o *BatchFunctionSignatureEntry) SetHasSignature(v bool)`

SetHasSignature sets HasSignature field to given value.


### GetParameters

`func (o *BatchFunctionSignatureEntry) GetParameters() []SignatureParameterEntry`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *BatchFunctionSignatureEntry) GetParametersOk() (*[]SignatureParameterEntry, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *BatchFunctionSignatureEntry) SetParameters(v []SignatureParameterEntry)`

SetParameters sets Parameters field to given value.


### SetParametersNil

`func (o *BatchFunctionSignatureEntry) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *BatchFunctionSignatureEntry) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetReturnDataTypeId

`func (o *BatchFunctionSignatureEntry) GetReturnDataTypeId() int64`

GetReturnDataTypeId returns the ReturnDataTypeId field if non-nil, zero value otherwise.

### GetReturnDataTypeIdOk

`func (o *BatchFunctionSignatureEntry) GetReturnDataTypeIdOk() (*int64, bool)`

GetReturnDataTypeIdOk returns a tuple with the ReturnDataTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnDataTypeId

`func (o *BatchFunctionSignatureEntry) SetReturnDataTypeId(v int64)`

SetReturnDataTypeId sets ReturnDataTypeId field to given value.

### HasReturnDataTypeId

`func (o *BatchFunctionSignatureEntry) HasReturnDataTypeId() bool`

HasReturnDataTypeId returns a boolean if a field has been set.

### GetSourceFunctionId

`func (o *BatchFunctionSignatureEntry) GetSourceFunctionId() int64`

GetSourceFunctionId returns the SourceFunctionId field if non-nil, zero value otherwise.

### GetSourceFunctionIdOk

`func (o *BatchFunctionSignatureEntry) GetSourceFunctionIdOk() (*int64, bool)`

GetSourceFunctionIdOk returns a tuple with the SourceFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFunctionId

`func (o *BatchFunctionSignatureEntry) SetSourceFunctionId(v int64)`

SetSourceFunctionId sets SourceFunctionId field to given value.

### HasSourceFunctionId

`func (o *BatchFunctionSignatureEntry) HasSourceFunctionId() bool`

HasSourceFunctionId returns a boolean if a field has been set.

### GetSourceType

`func (o *BatchFunctionSignatureEntry) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *BatchFunctionSignatureEntry) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *BatchFunctionSignatureEntry) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *BatchFunctionSignatureEntry) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


