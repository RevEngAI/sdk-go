# DynamicExecutionStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessage** | Pointer to **string** | Error detail, set when status is FAILED | [optional] 
**Logs** | [**AnalysisLogs**](AnalysisLogs.md) | Sandbox status log messages captured during the run. Contains a single \&quot;No logs available\&quot; message when none have been captured yet. | 
**Status** | **string** | Task status: UNINITIALISED, PENDING, RUNNING, COMPLETED, or FAILED | 

## Methods

### NewDynamicExecutionStatusResponse

`func NewDynamicExecutionStatusResponse(logs AnalysisLogs, status string, ) *DynamicExecutionStatusResponse`

NewDynamicExecutionStatusResponse instantiates a new DynamicExecutionStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicExecutionStatusResponseWithDefaults

`func NewDynamicExecutionStatusResponseWithDefaults() *DynamicExecutionStatusResponse`

NewDynamicExecutionStatusResponseWithDefaults instantiates a new DynamicExecutionStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessage

`func (o *DynamicExecutionStatusResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *DynamicExecutionStatusResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *DynamicExecutionStatusResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *DynamicExecutionStatusResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetLogs

`func (o *DynamicExecutionStatusResponse) GetLogs() AnalysisLogs`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *DynamicExecutionStatusResponse) GetLogsOk() (*AnalysisLogs, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *DynamicExecutionStatusResponse) SetLogs(v AnalysisLogs)`

SetLogs sets Logs field to given value.


### GetStatus

`func (o *DynamicExecutionStatusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DynamicExecutionStatusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DynamicExecutionStatusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


