# EventTOOLCALLEND

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**SseEventToolCallEndData**](SseEventToolCallEndData.md) |  | 
**Event** | **string** | The event name. | 
**Id** | Pointer to **int32** | The event ID. | [optional] 
**Retry** | Pointer to **int32** | The retry time in milliseconds. | [optional] 

## Methods

### NewEventTOOLCALLEND

`func NewEventTOOLCALLEND(data SseEventToolCallEndData, event string, ) *EventTOOLCALLEND`

NewEventTOOLCALLEND instantiates a new EventTOOLCALLEND object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventTOOLCALLENDWithDefaults

`func NewEventTOOLCALLENDWithDefaults() *EventTOOLCALLEND`

NewEventTOOLCALLENDWithDefaults instantiates a new EventTOOLCALLEND object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EventTOOLCALLEND) GetData() SseEventToolCallEndData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventTOOLCALLEND) GetDataOk() (*SseEventToolCallEndData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventTOOLCALLEND) SetData(v SseEventToolCallEndData)`

SetData sets Data field to given value.


### GetEvent

`func (o *EventTOOLCALLEND) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventTOOLCALLEND) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventTOOLCALLEND) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetId

`func (o *EventTOOLCALLEND) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventTOOLCALLEND) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventTOOLCALLEND) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EventTOOLCALLEND) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRetry

`func (o *EventTOOLCALLEND) GetRetry() int32`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *EventTOOLCALLEND) GetRetryOk() (*int32, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *EventTOOLCALLEND) SetRetry(v int32)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *EventTOOLCALLEND) HasRetry() bool`

HasRetry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


