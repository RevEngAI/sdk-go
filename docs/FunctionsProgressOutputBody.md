# FunctionsProgressOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionEmbeddingsCount** | **int64** | Functions whose embedding has been written back | 
**FunctionsCount** | **int64** | Functions found in the analysis | 
**PercentageCompleted** | **float64** | Embeddings as a percentage of functions, to 2 decimal places. 0 when the analysis has no functions | 

## Methods

### NewFunctionsProgressOutputBody

`func NewFunctionsProgressOutputBody(functionEmbeddingsCount int64, functionsCount int64, percentageCompleted float64, ) *FunctionsProgressOutputBody`

NewFunctionsProgressOutputBody instantiates a new FunctionsProgressOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionsProgressOutputBodyWithDefaults

`func NewFunctionsProgressOutputBodyWithDefaults() *FunctionsProgressOutputBody`

NewFunctionsProgressOutputBodyWithDefaults instantiates a new FunctionsProgressOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionEmbeddingsCount

`func (o *FunctionsProgressOutputBody) GetFunctionEmbeddingsCount() int64`

GetFunctionEmbeddingsCount returns the FunctionEmbeddingsCount field if non-nil, zero value otherwise.

### GetFunctionEmbeddingsCountOk

`func (o *FunctionsProgressOutputBody) GetFunctionEmbeddingsCountOk() (*int64, bool)`

GetFunctionEmbeddingsCountOk returns a tuple with the FunctionEmbeddingsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionEmbeddingsCount

`func (o *FunctionsProgressOutputBody) SetFunctionEmbeddingsCount(v int64)`

SetFunctionEmbeddingsCount sets FunctionEmbeddingsCount field to given value.


### GetFunctionsCount

`func (o *FunctionsProgressOutputBody) GetFunctionsCount() int64`

GetFunctionsCount returns the FunctionsCount field if non-nil, zero value otherwise.

### GetFunctionsCountOk

`func (o *FunctionsProgressOutputBody) GetFunctionsCountOk() (*int64, bool)`

GetFunctionsCountOk returns a tuple with the FunctionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionsCount

`func (o *FunctionsProgressOutputBody) SetFunctionsCount(v int64)`

SetFunctionsCount sets FunctionsCount field to given value.


### GetPercentageCompleted

`func (o *FunctionsProgressOutputBody) GetPercentageCompleted() float64`

GetPercentageCompleted returns the PercentageCompleted field if non-nil, zero value otherwise.

### GetPercentageCompletedOk

`func (o *FunctionsProgressOutputBody) GetPercentageCompletedOk() (*float64, bool)`

GetPercentageCompletedOk returns a tuple with the PercentageCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageCompleted

`func (o *FunctionsProgressOutputBody) SetPercentageCompleted(v float64)`

SetPercentageCompleted sets PercentageCompleted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


