# SseEventToolConfirmationRequiredData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**ToolConfirmationRequiredData**](ToolConfirmationRequiredData.md) | The payload of the event named in type. | 
**EventId** | **int64** | The stored event&#39;s ID, the same as the SSE id field. Pass the last one received as last_event_id to resume after a reconnection. | 
**SourceRunId** | Pointer to **string** |  | [optional] 
**Type** | **string** | The event name, the same as the SSE event field. | 

## Methods

### NewSseEventToolConfirmationRequiredData

`func NewSseEventToolConfirmationRequiredData(data ToolConfirmationRequiredData, eventId int64, type_ string, ) *SseEventToolConfirmationRequiredData`

NewSseEventToolConfirmationRequiredData instantiates a new SseEventToolConfirmationRequiredData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSseEventToolConfirmationRequiredDataWithDefaults

`func NewSseEventToolConfirmationRequiredDataWithDefaults() *SseEventToolConfirmationRequiredData`

NewSseEventToolConfirmationRequiredDataWithDefaults instantiates a new SseEventToolConfirmationRequiredData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SseEventToolConfirmationRequiredData) GetData() ToolConfirmationRequiredData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SseEventToolConfirmationRequiredData) GetDataOk() (*ToolConfirmationRequiredData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SseEventToolConfirmationRequiredData) SetData(v ToolConfirmationRequiredData)`

SetData sets Data field to given value.


### GetEventId

`func (o *SseEventToolConfirmationRequiredData) GetEventId() int64`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *SseEventToolConfirmationRequiredData) GetEventIdOk() (*int64, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *SseEventToolConfirmationRequiredData) SetEventId(v int64)`

SetEventId sets EventId field to given value.


### GetSourceRunId

`func (o *SseEventToolConfirmationRequiredData) GetSourceRunId() string`

GetSourceRunId returns the SourceRunId field if non-nil, zero value otherwise.

### GetSourceRunIdOk

`func (o *SseEventToolConfirmationRequiredData) GetSourceRunIdOk() (*string, bool)`

GetSourceRunIdOk returns a tuple with the SourceRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRunId

`func (o *SseEventToolConfirmationRequiredData) SetSourceRunId(v string)`

SetSourceRunId sets SourceRunId field to given value.

### HasSourceRunId

`func (o *SseEventToolConfirmationRequiredData) HasSourceRunId() bool`

HasSourceRunId returns a boolean if a field has been set.

### GetType

`func (o *SseEventToolConfirmationRequiredData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SseEventToolConfirmationRequiredData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SseEventToolConfirmationRequiredData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


