# ListDynamicExecutionRunsOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Runs** | [**[]DynamicExecutionRun**](DynamicExecutionRun.md) | The analysis&#39;s runs, newest first | 

## Methods

### NewListDynamicExecutionRunsOutputBody

`func NewListDynamicExecutionRunsOutputBody(runs []DynamicExecutionRun, ) *ListDynamicExecutionRunsOutputBody`

NewListDynamicExecutionRunsOutputBody instantiates a new ListDynamicExecutionRunsOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDynamicExecutionRunsOutputBodyWithDefaults

`func NewListDynamicExecutionRunsOutputBodyWithDefaults() *ListDynamicExecutionRunsOutputBody`

NewListDynamicExecutionRunsOutputBodyWithDefaults instantiates a new ListDynamicExecutionRunsOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuns

`func (o *ListDynamicExecutionRunsOutputBody) GetRuns() []DynamicExecutionRun`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *ListDynamicExecutionRunsOutputBody) GetRunsOk() (*[]DynamicExecutionRun, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *ListDynamicExecutionRunsOutputBody) SetRuns(v []DynamicExecutionRun)`

SetRuns sets Runs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


