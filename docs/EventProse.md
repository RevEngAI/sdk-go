# EventProse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**ProseEvent**](ProseEvent.md) |  | 
**Event** | **string** | The event name. | 
**Id** | Pointer to **int32** | The event ID. | [optional] 
**Retry** | Pointer to **int32** | The retry time in milliseconds. | [optional] 

## Methods

### NewEventProse

`func NewEventProse(data ProseEvent, event string, ) *EventProse`

NewEventProse instantiates a new EventProse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventProseWithDefaults

`func NewEventProseWithDefaults() *EventProse`

NewEventProseWithDefaults instantiates a new EventProse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EventProse) GetData() ProseEvent`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventProse) GetDataOk() (*ProseEvent, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventProse) SetData(v ProseEvent)`

SetData sets Data field to given value.


### GetEvent

`func (o *EventProse) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventProse) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventProse) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetId

`func (o *EventProse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventProse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventProse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EventProse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRetry

`func (o *EventProse) GetRetry() int32`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *EventProse) GetRetryOk() (*int32, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *EventProse) SetRetry(v int32)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *EventProse) HasRetry() bool`

HasRetry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


