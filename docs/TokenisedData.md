# TokenisedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**FunctionMapping** | Pointer to [**FunctionMapping**](FunctionMapping.md) | Complete mapping data for token resolution | [optional] 
**PredictedFunctionName** | Pointer to **string** | Predicted function name from the AI model | [optional] 
**Status** | **string** | Task status | 
**TokenisedDecompilation** | Pointer to **string** | Source code with placeholder tokens | [optional] 

## Methods

### NewTokenisedData

`func NewTokenisedData(status string, ) *TokenisedData`

NewTokenisedData instantiates a new TokenisedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenisedDataWithDefaults

`func NewTokenisedDataWithDefaults() *TokenisedData`

NewTokenisedDataWithDefaults instantiates a new TokenisedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *TokenisedData) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *TokenisedData) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *TokenisedData) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *TokenisedData) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetFunctionMapping

`func (o *TokenisedData) GetFunctionMapping() FunctionMapping`

GetFunctionMapping returns the FunctionMapping field if non-nil, zero value otherwise.

### GetFunctionMappingOk

`func (o *TokenisedData) GetFunctionMappingOk() (*FunctionMapping, bool)`

GetFunctionMappingOk returns a tuple with the FunctionMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionMapping

`func (o *TokenisedData) SetFunctionMapping(v FunctionMapping)`

SetFunctionMapping sets FunctionMapping field to given value.

### HasFunctionMapping

`func (o *TokenisedData) HasFunctionMapping() bool`

HasFunctionMapping returns a boolean if a field has been set.

### GetPredictedFunctionName

`func (o *TokenisedData) GetPredictedFunctionName() string`

GetPredictedFunctionName returns the PredictedFunctionName field if non-nil, zero value otherwise.

### GetPredictedFunctionNameOk

`func (o *TokenisedData) GetPredictedFunctionNameOk() (*string, bool)`

GetPredictedFunctionNameOk returns a tuple with the PredictedFunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedFunctionName

`func (o *TokenisedData) SetPredictedFunctionName(v string)`

SetPredictedFunctionName sets PredictedFunctionName field to given value.

### HasPredictedFunctionName

`func (o *TokenisedData) HasPredictedFunctionName() bool`

HasPredictedFunctionName returns a boolean if a field has been set.

### GetStatus

`func (o *TokenisedData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TokenisedData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TokenisedData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTokenisedDecompilation

`func (o *TokenisedData) GetTokenisedDecompilation() string`

GetTokenisedDecompilation returns the TokenisedDecompilation field if non-nil, zero value otherwise.

### GetTokenisedDecompilationOk

`func (o *TokenisedData) GetTokenisedDecompilationOk() (*string, bool)`

GetTokenisedDecompilationOk returns a tuple with the TokenisedDecompilation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenisedDecompilation

`func (o *TokenisedData) SetTokenisedDecompilation(v string)`

SetTokenisedDecompilation sets TokenisedDecompilation field to given value.

### HasTokenisedDecompilation

`func (o *TokenisedData) HasTokenisedDecompilation() bool`

HasTokenisedDecompilation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


