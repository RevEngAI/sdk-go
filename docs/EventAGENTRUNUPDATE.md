# EventAGENTRUNUPDATE

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**SseEventAgentRunUpdateData**](SseEventAgentRunUpdateData.md) |  | 
**Event** | **string** | The event name. | 
**Id** | Pointer to **int32** | The event ID. | [optional] 
**Retry** | Pointer to **int32** | The retry time in milliseconds. | [optional] 

## Methods

### NewEventAGENTRUNUPDATE

`func NewEventAGENTRUNUPDATE(data SseEventAgentRunUpdateData, event string, ) *EventAGENTRUNUPDATE`

NewEventAGENTRUNUPDATE instantiates a new EventAGENTRUNUPDATE object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventAGENTRUNUPDATEWithDefaults

`func NewEventAGENTRUNUPDATEWithDefaults() *EventAGENTRUNUPDATE`

NewEventAGENTRUNUPDATEWithDefaults instantiates a new EventAGENTRUNUPDATE object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EventAGENTRUNUPDATE) GetData() SseEventAgentRunUpdateData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventAGENTRUNUPDATE) GetDataOk() (*SseEventAgentRunUpdateData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventAGENTRUNUPDATE) SetData(v SseEventAgentRunUpdateData)`

SetData sets Data field to given value.


### GetEvent

`func (o *EventAGENTRUNUPDATE) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventAGENTRUNUPDATE) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventAGENTRUNUPDATE) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetId

`func (o *EventAGENTRUNUPDATE) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventAGENTRUNUPDATE) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventAGENTRUNUPDATE) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EventAGENTRUNUPDATE) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRetry

`func (o *EventAGENTRUNUPDATE) GetRetry() int32`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *EventAGENTRUNUPDATE) GetRetryOk() (*int32, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *EventAGENTRUNUPDATE) SetRetry(v int32)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *EventAGENTRUNUPDATE) HasRetry() bool`

HasRetry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


