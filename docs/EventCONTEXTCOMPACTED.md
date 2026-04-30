# EventCONTEXTCOMPACTED

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**SseEventContextCompactedData**](SseEventContextCompactedData.md) |  | 
**Event** | **string** | The event name. | 
**Id** | Pointer to **int32** | The event ID. | [optional] 
**Retry** | Pointer to **int32** | The retry time in milliseconds. | [optional] 

## Methods

### NewEventCONTEXTCOMPACTED

`func NewEventCONTEXTCOMPACTED(data SseEventContextCompactedData, event string, ) *EventCONTEXTCOMPACTED`

NewEventCONTEXTCOMPACTED instantiates a new EventCONTEXTCOMPACTED object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventCONTEXTCOMPACTEDWithDefaults

`func NewEventCONTEXTCOMPACTEDWithDefaults() *EventCONTEXTCOMPACTED`

NewEventCONTEXTCOMPACTEDWithDefaults instantiates a new EventCONTEXTCOMPACTED object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EventCONTEXTCOMPACTED) GetData() SseEventContextCompactedData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventCONTEXTCOMPACTED) GetDataOk() (*SseEventContextCompactedData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventCONTEXTCOMPACTED) SetData(v SseEventContextCompactedData)`

SetData sets Data field to given value.


### GetEvent

`func (o *EventCONTEXTCOMPACTED) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventCONTEXTCOMPACTED) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventCONTEXTCOMPACTED) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetId

`func (o *EventCONTEXTCOMPACTED) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventCONTEXTCOMPACTED) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventCONTEXTCOMPACTED) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EventCONTEXTCOMPACTED) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRetry

`func (o *EventCONTEXTCOMPACTED) GetRetry() int32`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *EventCONTEXTCOMPACTED) GetRetryOk() (*int32, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *EventCONTEXTCOMPACTED) SetRetry(v int32)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *EventCONTEXTCOMPACTED) HasRetry() bool`

HasRetry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


