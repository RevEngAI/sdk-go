# BatchUpdateDataTypesOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]BatchUpdateDataTypesResult**](BatchUpdateDataTypesResult.md) | Per-function outcomes in the same order as the input | 

## Methods

### NewBatchUpdateDataTypesOutputBody

`func NewBatchUpdateDataTypesOutputBody(results []BatchUpdateDataTypesResult, ) *BatchUpdateDataTypesOutputBody`

NewBatchUpdateDataTypesOutputBody instantiates a new BatchUpdateDataTypesOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchUpdateDataTypesOutputBodyWithDefaults

`func NewBatchUpdateDataTypesOutputBodyWithDefaults() *BatchUpdateDataTypesOutputBody`

NewBatchUpdateDataTypesOutputBodyWithDefaults instantiates a new BatchUpdateDataTypesOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *BatchUpdateDataTypesOutputBody) GetResults() []BatchUpdateDataTypesResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchUpdateDataTypesOutputBody) GetResultsOk() (*[]BatchUpdateDataTypesResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchUpdateDataTypesOutputBody) SetResults(v []BatchUpdateDataTypesResult)`

SetResults sets Results field to given value.


### SetResultsNil

`func (o *BatchUpdateDataTypesOutputBody) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *BatchUpdateDataTypesOutputBody) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


