# EventNamesFinished

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**NamesFinishedEvent**](NamesFinishedEvent.md) |  | 
**Event** | **string** | The event name. | 
**Id** | Pointer to **int32** | The event ID. | [optional] 
**Retry** | Pointer to **int32** | The retry time in milliseconds. | [optional] 

## Methods

### NewEventNamesFinished

`func NewEventNamesFinished(data NamesFinishedEvent, event string, ) *EventNamesFinished`

NewEventNamesFinished instantiates a new EventNamesFinished object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventNamesFinishedWithDefaults

`func NewEventNamesFinishedWithDefaults() *EventNamesFinished`

NewEventNamesFinishedWithDefaults instantiates a new EventNamesFinished object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EventNamesFinished) GetData() NamesFinishedEvent`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventNamesFinished) GetDataOk() (*NamesFinishedEvent, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventNamesFinished) SetData(v NamesFinishedEvent)`

SetData sets Data field to given value.


### GetEvent

`func (o *EventNamesFinished) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventNamesFinished) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventNamesFinished) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetId

`func (o *EventNamesFinished) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventNamesFinished) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventNamesFinished) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EventNamesFinished) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRetry

`func (o *EventNamesFinished) GetRetry() int32`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *EventNamesFinished) GetRetryOk() (*int32, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *EventNamesFinished) SetRetry(v int32)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *EventNamesFinished) HasRetry() bool`

HasRetry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


