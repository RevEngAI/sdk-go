# SseEventRunStartedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**RunStartedData**](RunStartedData.md) | The payload of the event named in type. | 
**EventId** | **int64** | The stored event&#39;s ID, the same as the SSE id field. Pass the last one received as last_event_id to resume after a reconnection. | 
**SourceRunId** | Pointer to **string** |  | [optional] 
**Type** | **string** | The event name, the same as the SSE event field. | 

## Methods

### NewSseEventRunStartedData

`func NewSseEventRunStartedData(data RunStartedData, eventId int64, type_ string, ) *SseEventRunStartedData`

NewSseEventRunStartedData instantiates a new SseEventRunStartedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSseEventRunStartedDataWithDefaults

`func NewSseEventRunStartedDataWithDefaults() *SseEventRunStartedData`

NewSseEventRunStartedDataWithDefaults instantiates a new SseEventRunStartedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SseEventRunStartedData) GetData() RunStartedData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SseEventRunStartedData) GetDataOk() (*RunStartedData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SseEventRunStartedData) SetData(v RunStartedData)`

SetData sets Data field to given value.


### GetEventId

`func (o *SseEventRunStartedData) GetEventId() int64`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *SseEventRunStartedData) GetEventIdOk() (*int64, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *SseEventRunStartedData) SetEventId(v int64)`

SetEventId sets EventId field to given value.


### GetSourceRunId

`func (o *SseEventRunStartedData) GetSourceRunId() string`

GetSourceRunId returns the SourceRunId field if non-nil, zero value otherwise.

### GetSourceRunIdOk

`func (o *SseEventRunStartedData) GetSourceRunIdOk() (*string, bool)`

GetSourceRunIdOk returns a tuple with the SourceRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRunId

`func (o *SseEventRunStartedData) SetSourceRunId(v string)`

SetSourceRunId sets SourceRunId field to given value.

### HasSourceRunId

`func (o *SseEventRunStartedData) HasSourceRunId() bool`

HasSourceRunId returns a boolean if a field has been set.

### GetType

`func (o *SseEventRunStartedData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SseEventRunStartedData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SseEventRunStartedData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


