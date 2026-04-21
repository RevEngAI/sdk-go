# TaskStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**BinaryTaskStatus**](BinaryTaskStatus.md) |  | 

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


