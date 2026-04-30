# EventRUNCANCELLED

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**SseEventRunCancelledData**](SseEventRunCancelledData.md) |  | 
**Event** | **string** | The event name. | 
**Id** | Pointer to **int32** | The event ID. | [optional] 
**Retry** | Pointer to **int32** | The retry time in milliseconds. | [optional] 

## Methods

### NewEventRUNCANCELLED

`func NewEventRUNCANCELLED(data SseEventRunCancelledData, event string, ) *EventRUNCANCELLED`

NewEventRUNCANCELLED instantiates a new EventRUNCANCELLED object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventRUNCANCELLEDWithDefaults

`func NewEventRUNCANCELLEDWithDefaults() *EventRUNCANCELLED`

NewEventRUNCANCELLEDWithDefaults instantiates a new EventRUNCANCELLED object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EventRUNCANCELLED) GetData() SseEventRunCancelledData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventRUNCANCELLED) GetDataOk() (*SseEventRunCancelledData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventRUNCANCELLED) SetData(v SseEventRunCancelledData)`

SetData sets Data field to given value.


### GetEvent

`func (o *EventRUNCANCELLED) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventRUNCANCELLED) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventRUNCANCELLED) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetId

`func (o *EventRUNCANCELLED) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventRUNCANCELLED) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventRUNCANCELLED) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EventRUNCANCELLED) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRetry

`func (o *EventRUNCANCELLED) GetRetry() int32`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *EventRUNCANCELLED) GetRetryOk() (*int32, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *EventRUNCANCELLED) SetRetry(v int32)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *EventRUNCANCELLED) HasRetry() bool`

HasRetry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


