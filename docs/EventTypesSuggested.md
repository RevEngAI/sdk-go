# EventTypesSuggested

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**TypesSuggestedEvent**](TypesSuggestedEvent.md) |  | 
**Event** | **string** | The event name. | 
**Id** | Pointer to **int32** | The event ID. | [optional] 
**Retry** | Pointer to **int32** | The retry time in milliseconds. | [optional] 

## Methods

### NewEventTypesSuggested

`func NewEventTypesSuggested(data TypesSuggestedEvent, event string, ) *EventTypesSuggested`

NewEventTypesSuggested instantiates a new EventTypesSuggested object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventTypesSuggestedWithDefaults

`func NewEventTypesSuggestedWithDefaults() *EventTypesSuggested`

NewEventTypesSuggestedWithDefaults instantiates a new EventTypesSuggested object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EventTypesSuggested) GetData() TypesSuggestedEvent`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventTypesSuggested) GetDataOk() (*TypesSuggestedEvent, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventTypesSuggested) SetData(v TypesSuggestedEvent)`

SetData sets Data field to given value.


### GetEvent

`func (o *EventTypesSuggested) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventTypesSuggested) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventTypesSuggested) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetId

`func (o *EventTypesSuggested) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventTypesSuggested) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventTypesSuggested) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EventTypesSuggested) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRetry

`func (o *EventTypesSuggested) GetRetry() int32`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *EventTypesSuggested) GetRetryOk() (*int32, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *EventTypesSuggested) SetRetry(v int32)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *EventTypesSuggested) HasRetry() bool`

HasRetry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


