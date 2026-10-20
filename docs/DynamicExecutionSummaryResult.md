# DynamicExecutionSummaryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SummaryAvailable** | **bool** | Whether the summary can be fetched from the summary endpoint. Always true when the operation succeeded. | 

## Methods

### NewDynamicExecutionSummaryResult

`func NewDynamicExecutionSummaryResult(summaryAvailable bool, ) *DynamicExecutionSummaryResult`

NewDynamicExecutionSummaryResult instantiates a new DynamicExecutionSummaryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicExecutionSummaryResultWithDefaults

`func NewDynamicExecutionSummaryResultWithDefaults() *DynamicExecutionSummaryResult`

NewDynamicExecutionSummaryResultWithDefaults instantiates a new DynamicExecutionSummaryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummaryAvailable

`func (o *DynamicExecutionSummaryResult) GetSummaryAvailable() bool`

GetSummaryAvailable returns the SummaryAvailable field if non-nil, zero value otherwise.

### GetSummaryAvailableOk

`func (o *DynamicExecutionSummaryResult) GetSummaryAvailableOk() (*bool, bool)`

GetSummaryAvailableOk returns a tuple with the SummaryAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryAvailable

`func (o *DynamicExecutionSummaryResult) SetSummaryAvailable(v bool)`

SetSummaryAvailable sets SummaryAvailable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


