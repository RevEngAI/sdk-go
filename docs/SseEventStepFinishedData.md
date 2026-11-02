# SseEventStepFinishedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**StepFinishedData**](StepFinishedData.md) | The payload of the event named in type. | 
**EventId** | **int64** | The stored event&#39;s ID, the same as the SSE id field. Pass the last one received as last_event_id to resume after a reconnection. | 
**SourceRunId** | Pointer to **string** |  | [optional] 
**Type** | **string** | The event name, the same as the SSE event field. | 

## Methods

### NewSseEventStepFinishedData

`func NewSseEventStepFinishedData(data StepFinishedData, eventId int64, type_ string, ) *SseEventStepFinishedData`

NewSseEventStepFinishedData instantiates a new SseEventStepFinishedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSseEventStepFinishedDataWithDefaults

`func NewSseEventStepFinishedDataWithDefaults() *SseEventStepFinishedData`

NewSseEventStepFinishedDataWithDefaults instantiates a new SseEventStepFinishedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SseEventStepFinishedData) GetData() StepFinishedData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SseEventStepFinishedData) GetDataOk() (*StepFinishedData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SseEventStepFinishedData) SetData(v StepFinishedData)`

SetData sets Data field to given value.


### GetEventId

`func (o *SseEventStepFinishedData) GetEventId() int64`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *SseEventStepFinishedData) GetEventIdOk() (*int64, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *SseEventStepFinishedData) SetEventId(v int64)`

SetEventId sets EventId field to given value.


### GetSourceRunId

`func (o *SseEventStepFinishedData) GetSourceRunId() string`

GetSourceRunId returns the SourceRunId field if non-nil, zero value otherwise.

### GetSourceRunIdOk

`func (o *SseEventStepFinishedData) GetSourceRunIdOk() (*string, bool)`

GetSourceRunIdOk returns a tuple with the SourceRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRunId

`func (o *SseEventStepFinishedData) SetSourceRunId(v string)`

SetSourceRunId sets SourceRunId field to given value.

### HasSourceRunId

`func (o *SseEventStepFinishedData) HasSourceRunId() bool`

HasSourceRunId returns a boolean if a field has been set.

### GetType

`func (o *SseEventStepFinishedData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SseEventStepFinishedData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SseEventStepFinishedData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


