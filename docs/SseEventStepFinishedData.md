# SseEventStepFinishedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **interface{}** |  | 
**EventId** | **int64** |  | 
**Type** | **string** |  | 

## Methods

### NewSseEventStepFinishedData

`func NewSseEventStepFinishedData(data interface{}, eventId int64, type_ string, ) *SseEventStepFinishedData`

NewSseEventStepFinishedData instantiates a new SseEventStepFinishedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSseEventStepFinishedDataWithDefaults

`func NewSseEventStepFinishedDataWithDefaults() *SseEventStepFinishedData`

NewSseEventStepFinishedDataWithDefaults instantiates a new SseEventStepFinishedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SseEventStepFinishedData) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SseEventStepFinishedData) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SseEventStepFinishedData) SetData(v interface{})`

SetData sets Data field to given value.


### SetDataNil

`func (o *SseEventStepFinishedData) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *SseEventStepFinishedData) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetEventId

`func (o *SseEventStepFinishedData) GetEventId() int64`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *SseEventStepFinishedData) GetEventIdOk() (*int64, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *SseEventStepFinishedData) SetEventId(v int64)`

SetEventId sets EventId field to given value.


### GetType

`func (o *SseEventStepFinishedData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SseEventStepFinishedData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SseEventStepFinishedData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


