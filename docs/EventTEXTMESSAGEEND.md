# EventTEXTMESSAGEEND

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**SseEventTextMessageEndData**](SseEventTextMessageEndData.md) |  | 
**Event** | **string** | The event name. | 
**Id** | Pointer to **int32** | The event ID. | [optional] 
**Retry** | Pointer to **int32** | The retry time in milliseconds. | [optional] 

## Methods

### NewEventTEXTMESSAGEEND

`func NewEventTEXTMESSAGEEND(data SseEventTextMessageEndData, event string, ) *EventTEXTMESSAGEEND`

NewEventTEXTMESSAGEEND instantiates a new EventTEXTMESSAGEEND object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventTEXTMESSAGEENDWithDefaults

`func NewEventTEXTMESSAGEENDWithDefaults() *EventTEXTMESSAGEEND`

NewEventTEXTMESSAGEENDWithDefaults instantiates a new EventTEXTMESSAGEEND object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EventTEXTMESSAGEEND) GetData() SseEventTextMessageEndData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventTEXTMESSAGEEND) GetDataOk() (*SseEventTextMessageEndData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventTEXTMESSAGEEND) SetData(v SseEventTextMessageEndData)`

SetData sets Data field to given value.


### GetEvent

`func (o *EventTEXTMESSAGEEND) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventTEXTMESSAGEEND) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventTEXTMESSAGEEND) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetId

`func (o *EventTEXTMESSAGEEND) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventTEXTMESSAGEEND) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventTEXTMESSAGEEND) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EventTEXTMESSAGEEND) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRetry

`func (o *EventTEXTMESSAGEEND) GetRetry() int32`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *EventTEXTMESSAGEEND) GetRetryOk() (*int32, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *EventTEXTMESSAGEEND) SetRetry(v int32)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *EventTEXTMESSAGEEND) HasRetry() bool`

HasRetry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


