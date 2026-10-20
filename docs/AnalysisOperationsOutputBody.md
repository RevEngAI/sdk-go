# AnalysisOperationsOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operations** | [**[]AnalysisOperationBody**](AnalysisOperationBody.md) | The analysis&#39; static-analysis operation, followed by the latest run of each other operation started on its binary. Operations that have never run are omitted | 

## Methods

### NewAnalysisOperationsOutputBody

`func NewAnalysisOperationsOutputBody(operations []AnalysisOperationBody, ) *AnalysisOperationsOutputBody`

NewAnalysisOperationsOutputBody instantiates a new AnalysisOperationsOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisOperationsOutputBodyWithDefaults

`func NewAnalysisOperationsOutputBodyWithDefaults() *AnalysisOperationsOutputBody`

NewAnalysisOperationsOutputBodyWithDefaults instantiates a new AnalysisOperationsOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperations

`func (o *AnalysisOperationsOutputBody) GetOperations() []AnalysisOperationBody`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *AnalysisOperationsOutputBody) GetOperationsOk() (*[]AnalysisOperationBody, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *AnalysisOperationsOutputBody) SetOperations(v []AnalysisOperationBody)`

SetOperations sets Operations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


