# EventRUNERROR

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**SseEventRunErrorData**](SseEventRunErrorData.md) |  | 
**Event** | **string** | The event name. | 
**Id** | Pointer to **int32** | The event ID. | [optional] 
**Retry** | Pointer to **int32** | The retry time in milliseconds. | [optional] 

## Methods

### NewEventRUNERROR

`func NewEventRUNERROR(data SseEventRunErrorData, event string, ) *EventRUNERROR`

NewEventRUNERROR instantiates a new EventRUNERROR object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventRUNERRORWithDefaults

`func NewEventRUNERRORWithDefaults() *EventRUNERROR`

NewEventRUNERRORWithDefaults instantiates a new EventRUNERROR object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EventRUNERROR) GetData() SseEventRunErrorData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventRUNERROR) GetDataOk() (*SseEventRunErrorData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventRUNERROR) SetData(v SseEventRunErrorData)`

SetData sets Data field to given value.


### GetEvent

`func (o *EventRUNERROR) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventRUNERROR) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventRUNERROR) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetId

`func (o *EventRUNERROR) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventRUNERROR) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventRUNERROR) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EventRUNERROR) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRetry

`func (o *EventRUNERROR) GetRetry() int32`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *EventRUNERROR) GetRetryOk() (*int32, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *EventRUNERROR) SetRetry(v int32)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *EventRUNERROR) HasRetry() bool`

HasRetry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


