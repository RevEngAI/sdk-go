# SseEventTextMessageEndData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **interface{}** |  | 
**EventId** | **int64** |  | 
**Type** | **string** |  | 

## Methods

### NewSseEventTextMessageEndData

`func NewSseEventTextMessageEndData(data interface{}, eventId int64, type_ string, ) *SseEventTextMessageEndData`

NewSseEventTextMessageEndData instantiates a new SseEventTextMessageEndData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSseEventTextMessageEndDataWithDefaults

`func NewSseEventTextMessageEndDataWithDefaults() *SseEventTextMessageEndData`

NewSseEventTextMessageEndDataWithDefaults instantiates a new SseEventTextMessageEndData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SseEventTextMessageEndData) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SseEventTextMessageEndData) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SseEventTextMessageEndData) SetData(v interface{})`

SetData sets Data field to given value.


### SetDataNil

`func (o *SseEventTextMessageEndData) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *SseEventTextMessageEndData) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetEventId

`func (o *SseEventTextMessageEndData) GetEventId() int64`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *SseEventTextMessageEndData) GetEventIdOk() (*int64, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *SseEventTextMessageEndData) SetEventId(v int64)`

SetEventId sets EventId field to given value.


### GetType

`func (o *SseEventTextMessageEndData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SseEventTextMessageEndData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SseEventTextMessageEndData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


