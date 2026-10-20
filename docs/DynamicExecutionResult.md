# DynamicExecutionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportAvailable** | **bool** | Whether the run&#39;s report can be fetched from the report endpoint. Always true when the operation succeeded. | 

## Methods

### NewDynamicExecutionResult

`func NewDynamicExecutionResult(reportAvailable bool, ) *DynamicExecutionResult`

NewDynamicExecutionResult instantiates a new DynamicExecutionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicExecutionResultWithDefaults

`func NewDynamicExecutionResultWithDefaults() *DynamicExecutionResult`

NewDynamicExecutionResultWithDefaults instantiates a new DynamicExecutionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportAvailable

`func (o *DynamicExecutionResult) GetReportAvailable() bool`

GetReportAvailable returns the ReportAvailable field if non-nil, zero value otherwise.

### GetReportAvailableOk

`func (o *DynamicExecutionResult) GetReportAvailableOk() (*bool, bool)`

GetReportAvailableOk returns a tuple with the ReportAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportAvailable

`func (o *DynamicExecutionResult) SetReportAvailable(v bool)`

SetReportAvailable sets ReportAvailable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


