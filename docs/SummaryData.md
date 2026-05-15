# SummaryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiSummary** | **string** | Summary with code tags removed | 
**Summary** | **string** | Raw summary from the model | 
**TaskStatus** | **string** | Task status | 

## Methods

### NewSummaryData

`func NewSummaryData(aiSummary string, summary string, taskStatus string, ) *SummaryData`

NewSummaryData instantiates a new SummaryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryDataWithDefaults

`func NewSummaryDataWithDefaults() *SummaryData`

NewSummaryDataWithDefaults instantiates a new SummaryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiSummary

`func (o *SummaryData) GetAiSummary() string`

GetAiSummary returns the AiSummary field if non-nil, zero value otherwise.

### GetAiSummaryOk

`func (o *SummaryData) GetAiSummaryOk() (*string, bool)`

GetAiSummaryOk returns a tuple with the AiSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiSummary

`func (o *SummaryData) SetAiSummary(v string)`

SetAiSummary sets AiSummary field to given value.


### GetSummary

`func (o *SummaryData) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *SummaryData) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *SummaryData) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetTaskStatus

`func (o *SummaryData) GetTaskStatus() string`

GetTaskStatus returns the TaskStatus field if non-nil, zero value otherwise.

### GetTaskStatusOk

`func (o *SummaryData) GetTaskStatusOk() (*string, bool)`

GetTaskStatusOk returns a tuple with the TaskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStatus

`func (o *SummaryData) SetTaskStatus(v string)`

SetTaskStatus sets TaskStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


