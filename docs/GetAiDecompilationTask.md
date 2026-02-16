# GetAiDecompilationTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Decompilation** | **NullableString** |  | 
**RawDecompilation** | **NullableString** |  | 
**FunctionMapping** | [**map[string]InverseFunctionMapItem**](InverseFunctionMapItem.md) |  | 
**FunctionMappingFull** | [**NullableFunctionMappingFull**](FunctionMappingFull.md) |  | 
**Summary** | Pointer to **NullableString** |  | [optional] 
**AiSummary** | Pointer to **NullableString** |  | [optional] 
**RawAiSummary** | Pointer to **NullableString** |  | [optional] 
**PredictedFunctionName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetAiDecompilationTask

`func NewGetAiDecompilationTask(status string, decompilation NullableString, rawDecompilation NullableString, functionMapping map[string]InverseFunctionMapItem, functionMappingFull NullableFunctionMappingFull, ) *GetAiDecompilationTask`

NewGetAiDecompilationTask instantiates a new GetAiDecompilationTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAiDecompilationTaskWithDefaults

`func NewGetAiDecompilationTaskWithDefaults() *GetAiDecompilationTask`

NewGetAiDecompilationTaskWithDefaults instantiates a new GetAiDecompilationTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetAiDecompilationTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAiDecompilationTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAiDecompilationTask) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDecompilation

`func (o *GetAiDecompilationTask) GetDecompilation() string`

GetDecompilation returns the Decompilation field if non-nil, zero value otherwise.

### GetDecompilationOk

`func (o *GetAiDecompilationTask) GetDecompilationOk() (*string, bool)`

GetDecompilationOk returns a tuple with the Decompilation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecompilation

`func (o *GetAiDecompilationTask) SetDecompilation(v string)`

SetDecompilation sets Decompilation field to given value.


### SetDecompilationNil

`func (o *GetAiDecompilationTask) SetDecompilationNil(b bool)`

 SetDecompilationNil sets the value for Decompilation to be an explicit nil

### UnsetDecompilation
`func (o *GetAiDecompilationTask) UnsetDecompilation()`

UnsetDecompilation ensures that no value is present for Decompilation, not even an explicit nil
### GetRawDecompilation

`func (o *GetAiDecompilationTask) GetRawDecompilation() string`

GetRawDecompilation returns the RawDecompilation field if non-nil, zero value otherwise.

### GetRawDecompilationOk

`func (o *GetAiDecompilationTask) GetRawDecompilationOk() (*string, bool)`

GetRawDecompilationOk returns a tuple with the RawDecompilation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawDecompilation

`func (o *GetAiDecompilationTask) SetRawDecompilation(v string)`

SetRawDecompilation sets RawDecompilation field to given value.


### SetRawDecompilationNil

`func (o *GetAiDecompilationTask) SetRawDecompilationNil(b bool)`

 SetRawDecompilationNil sets the value for RawDecompilation to be an explicit nil

### UnsetRawDecompilation
`func (o *GetAiDecompilationTask) UnsetRawDecompilation()`

UnsetRawDecompilation ensures that no value is present for RawDecompilation, not even an explicit nil
### GetFunctionMapping

`func (o *GetAiDecompilationTask) GetFunctionMapping() map[string]InverseFunctionMapItem`

GetFunctionMapping returns the FunctionMapping field if non-nil, zero value otherwise.

### GetFunctionMappingOk

`func (o *GetAiDecompilationTask) GetFunctionMappingOk() (*map[string]InverseFunctionMapItem, bool)`

GetFunctionMappingOk returns a tuple with the FunctionMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionMapping

`func (o *GetAiDecompilationTask) SetFunctionMapping(v map[string]InverseFunctionMapItem)`

SetFunctionMapping sets FunctionMapping field to given value.


### SetFunctionMappingNil

`func (o *GetAiDecompilationTask) SetFunctionMappingNil(b bool)`

 SetFunctionMappingNil sets the value for FunctionMapping to be an explicit nil

### UnsetFunctionMapping
`func (o *GetAiDecompilationTask) UnsetFunctionMapping()`

UnsetFunctionMapping ensures that no value is present for FunctionMapping, not even an explicit nil
### GetFunctionMappingFull

`func (o *GetAiDecompilationTask) GetFunctionMappingFull() FunctionMappingFull`

GetFunctionMappingFull returns the FunctionMappingFull field if non-nil, zero value otherwise.

### GetFunctionMappingFullOk

`func (o *GetAiDecompilationTask) GetFunctionMappingFullOk() (*FunctionMappingFull, bool)`

GetFunctionMappingFullOk returns a tuple with the FunctionMappingFull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionMappingFull

`func (o *GetAiDecompilationTask) SetFunctionMappingFull(v FunctionMappingFull)`

SetFunctionMappingFull sets FunctionMappingFull field to given value.


### SetFunctionMappingFullNil

`func (o *GetAiDecompilationTask) SetFunctionMappingFullNil(b bool)`

 SetFunctionMappingFullNil sets the value for FunctionMappingFull to be an explicit nil

### UnsetFunctionMappingFull
`func (o *GetAiDecompilationTask) UnsetFunctionMappingFull()`

UnsetFunctionMappingFull ensures that no value is present for FunctionMappingFull, not even an explicit nil
### GetSummary

`func (o *GetAiDecompilationTask) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *GetAiDecompilationTask) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *GetAiDecompilationTask) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *GetAiDecompilationTask) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *GetAiDecompilationTask) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *GetAiDecompilationTask) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetAiSummary

`func (o *GetAiDecompilationTask) GetAiSummary() string`

GetAiSummary returns the AiSummary field if non-nil, zero value otherwise.

### GetAiSummaryOk

`func (o *GetAiDecompilationTask) GetAiSummaryOk() (*string, bool)`

GetAiSummaryOk returns a tuple with the AiSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiSummary

`func (o *GetAiDecompilationTask) SetAiSummary(v string)`

SetAiSummary sets AiSummary field to given value.

### HasAiSummary

`func (o *GetAiDecompilationTask) HasAiSummary() bool`

HasAiSummary returns a boolean if a field has been set.

### SetAiSummaryNil

`func (o *GetAiDecompilationTask) SetAiSummaryNil(b bool)`

 SetAiSummaryNil sets the value for AiSummary to be an explicit nil

### UnsetAiSummary
`func (o *GetAiDecompilationTask) UnsetAiSummary()`

UnsetAiSummary ensures that no value is present for AiSummary, not even an explicit nil
### GetRawAiSummary

`func (o *GetAiDecompilationTask) GetRawAiSummary() string`

GetRawAiSummary returns the RawAiSummary field if non-nil, zero value otherwise.

### GetRawAiSummaryOk

`func (o *GetAiDecompilationTask) GetRawAiSummaryOk() (*string, bool)`

GetRawAiSummaryOk returns a tuple with the RawAiSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawAiSummary

`func (o *GetAiDecompilationTask) SetRawAiSummary(v string)`

SetRawAiSummary sets RawAiSummary field to given value.

### HasRawAiSummary

`func (o *GetAiDecompilationTask) HasRawAiSummary() bool`

HasRawAiSummary returns a boolean if a field has been set.

### SetRawAiSummaryNil

`func (o *GetAiDecompilationTask) SetRawAiSummaryNil(b bool)`

 SetRawAiSummaryNil sets the value for RawAiSummary to be an explicit nil

### UnsetRawAiSummary
`func (o *GetAiDecompilationTask) UnsetRawAiSummary()`

UnsetRawAiSummary ensures that no value is present for RawAiSummary, not even an explicit nil
### GetPredictedFunctionName

`func (o *GetAiDecompilationTask) GetPredictedFunctionName() string`

GetPredictedFunctionName returns the PredictedFunctionName field if non-nil, zero value otherwise.

### GetPredictedFunctionNameOk

`func (o *GetAiDecompilationTask) GetPredictedFunctionNameOk() (*string, bool)`

GetPredictedFunctionNameOk returns a tuple with the PredictedFunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedFunctionName

`func (o *GetAiDecompilationTask) SetPredictedFunctionName(v string)`

SetPredictedFunctionName sets PredictedFunctionName field to given value.

### HasPredictedFunctionName

`func (o *GetAiDecompilationTask) HasPredictedFunctionName() bool`

HasPredictedFunctionName returns a boolean if a field has been set.

### SetPredictedFunctionNameNil

`func (o *GetAiDecompilationTask) SetPredictedFunctionNameNil(b bool)`

 SetPredictedFunctionNameNil sets the value for PredictedFunctionName to be an explicit nil

### UnsetPredictedFunctionName
`func (o *GetAiDecompilationTask) UnsetPredictedFunctionName()`

UnsetPredictedFunctionName ensures that no value is present for PredictedFunctionName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


