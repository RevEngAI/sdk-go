# EventRUNSTARTED

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**SseEventRunStartedData**](SseEventRunStartedData.md) |  | 
**Event** | **string** | The event name. | 
**Id** | Pointer to **int32** | The event ID. | [optional] 
**Retry** | Pointer to **int32** | The retry time in milliseconds. | [optional] 

## Methods

### NewEventRUNSTARTED

`func NewEventRUNSTARTED(data SseEventRunStartedData, event string, ) *EventRUNSTARTED`

NewEventRUNSTARTED instantiates a new EventRUNSTARTED object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventRUNSTARTEDWithDefaults

`func NewEventRUNSTARTEDWithDefaults() *EventRUNSTARTED`

NewEventRUNSTARTEDWithDefaults instantiates a new EventRUNSTARTED object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EventRUNSTARTED) GetData() SseEventRunStartedData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventRUNSTARTED) GetDataOk() (*SseEventRunStartedData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventRUNSTARTED) SetData(v SseEventRunStartedData)`

SetData sets Data field to given value.


### GetEvent

`func (o *EventRUNSTARTED) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventRUNSTARTED) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventRUNSTARTED) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetId

`func (o *EventRUNSTARTED) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventRUNSTARTED) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventRUNSTARTED) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EventRUNSTARTED) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRetry

`func (o *EventRUNSTARTED) GetRetry() int32`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *EventRUNSTARTED) GetRetryOk() (*int32, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *EventRUNSTARTED) SetRetry(v int32)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *EventRUNSTARTED) HasRetry() bool`

HasRetry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


