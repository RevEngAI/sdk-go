# EventTEXTMESSAGESTART

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**SseEventTextMessageStartData**](SseEventTextMessageStartData.md) |  | 
**Event** | **string** | The event name. | 
**Id** | Pointer to **int32** | The event ID. | [optional] 
**Retry** | Pointer to **int32** | The retry time in milliseconds. | [optional] 

## Methods

### NewEventTEXTMESSAGESTART

`func NewEventTEXTMESSAGESTART(data SseEventTextMessageStartData, event string, ) *EventTEXTMESSAGESTART`

NewEventTEXTMESSAGESTART instantiates a new EventTEXTMESSAGESTART object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventTEXTMESSAGESTARTWithDefaults

`func NewEventTEXTMESSAGESTARTWithDefaults() *EventTEXTMESSAGESTART`

NewEventTEXTMESSAGESTARTWithDefaults instantiates a new EventTEXTMESSAGESTART object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EventTEXTMESSAGESTART) GetData() SseEventTextMessageStartData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventTEXTMESSAGESTART) GetDataOk() (*SseEventTextMessageStartData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventTEXTMESSAGESTART) SetData(v SseEventTextMessageStartData)`

SetData sets Data field to given value.


### GetEvent

`func (o *EventTEXTMESSAGESTART) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventTEXTMESSAGESTART) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventTEXTMESSAGESTART) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetId

`func (o *EventTEXTMESSAGESTART) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventTEXTMESSAGESTART) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventTEXTMESSAGESTART) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EventTEXTMESSAGESTART) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRetry

`func (o *EventTEXTMESSAGESTART) GetRetry() int32`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *EventTEXTMESSAGESTART) GetRetryOk() (*int32, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *EventTEXTMESSAGESTART) SetRetry(v int32)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *EventTEXTMESSAGESTART) HasRetry() bool`

HasRetry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


