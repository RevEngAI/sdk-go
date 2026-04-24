# TaskStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**BinaryTaskStatus**](BinaryTaskStatus.md) |  | 
**LogHistory** | Pointer to **[][]interface{}** |  | [optional] 

## Methods

### NewTaskStatusResponse

`func NewTaskStatusResponse(status BinaryTaskStatus, ) *TaskStatusResponse`

NewTaskStatusResponse instantiates a new TaskStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskStatusResponseWithDefaults

`func NewTaskStatusResponseWithDefaults() *TaskStatusResponse`

NewTaskStatusResponseWithDefaults instantiates a new TaskStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TaskStatusResponse) GetStatus() BinaryTaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskStatusResponse) GetStatusOk() (*BinaryTaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskStatusResponse) SetStatus(v BinaryTaskStatus)`

SetStatus sets Status field to given value.


### GetLogHistory

`func (o *TaskStatusResponse) GetLogHistory() [][]interface{}`

GetLogHistory returns the LogHistory field if non-nil, zero value otherwise.

### GetLogHistoryOk

`func (o *TaskStatusResponse) GetLogHistoryOk() (*[][]interface{}, bool)`

GetLogHistoryOk returns a tuple with the LogHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogHistory

`func (o *TaskStatusResponse) SetLogHistory(v [][]interface{})`

SetLogHistory sets LogHistory field to given value.

### HasLogHistory

`func (o *TaskStatusResponse) HasLogHistory() bool`

HasLogHistory returns a boolean if a field has been set.

### SetLogHistoryNil

`func (o *TaskStatusResponse) SetLogHistoryNil(b bool)`

 SetLogHistoryNil sets the value for LogHistory to be an explicit nil

### UnsetLogHistory
`func (o *TaskStatusResponse) UnsetLogHistory()`

UnsetLogHistory ensures that no value is present for LogHistory, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


