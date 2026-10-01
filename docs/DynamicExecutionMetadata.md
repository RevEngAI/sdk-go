# DynamicExecutionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | [**AnalysisLogs**](AnalysisLogs.md) | Sandbox status log messages captured during the run. Empty when none have been captured yet. | 
**Status** | **string** | Run status. UNINITIALISED means this analysis has never had a run triggered. | 

## Methods

### NewDynamicExecutionMetadata

`func NewDynamicExecutionMetadata(logs AnalysisLogs, status string, ) *DynamicExecutionMetadata`

NewDynamicExecutionMetadata instantiates a new DynamicExecutionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicExecutionMetadataWithDefaults

`func NewDynamicExecutionMetadataWithDefaults() *DynamicExecutionMetadata`

NewDynamicExecutionMetadataWithDefaults instantiates a new DynamicExecutionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *DynamicExecutionMetadata) GetLogs() AnalysisLogs`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *DynamicExecutionMetadata) GetLogsOk() (*AnalysisLogs, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *DynamicExecutionMetadata) SetLogs(v AnalysisLogs)`

SetLogs sets Logs field to given value.


### GetStatus

`func (o *DynamicExecutionMetadata) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DynamicExecutionMetadata) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DynamicExecutionMetadata) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


