# SseEventRunErrorData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **interface{}** |  | 
**EventId** | **int64** |  | 
**Type** | **string** |  | 

## Methods

### NewSseEventRunErrorData

`func NewSseEventRunErrorData(data interface{}, eventId int64, type_ string, ) *SseEventRunErrorData`

NewSseEventRunErrorData instantiates a new SseEventRunErrorData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSseEventRunErrorDataWithDefaults

`func NewSseEventRunErrorDataWithDefaults() *SseEventRunErrorData`

NewSseEventRunErrorDataWithDefaults instantiates a new SseEventRunErrorData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SseEventRunErrorData) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SseEventRunErrorData) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SseEventRunErrorData) SetData(v interface{})`

SetData sets Data field to given value.


### SetDataNil

`func (o *SseEventRunErrorData) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *SseEventRunErrorData) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetEventId

`func (o *SseEventRunErrorData) GetEventId() int64`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *SseEventRunErrorData) GetEventIdOk() (*int64, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *SseEventRunErrorData) SetEventId(v int64)`

SetEventId sets EventId field to given value.


### GetType

`func (o *SseEventRunErrorData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SseEventRunErrorData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SseEventRunErrorData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


