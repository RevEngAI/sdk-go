# QueuePositionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**QueuePosition** | **int64** | Number of Processing analyses ahead of this one in the queue. 0 if this analysis is not Processing or has no analyses ahead of it. | 

## Methods

### NewQueuePositionResponse

`func NewQueuePositionResponse(queuePosition int64, ) *QueuePositionResponse`

NewQueuePositionResponse instantiates a new QueuePositionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueuePositionResponseWithDefaults

`func NewQueuePositionResponseWithDefaults() *QueuePositionResponse`

NewQueuePositionResponseWithDefaults instantiates a new QueuePositionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *QueuePositionResponse) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *QueuePositionResponse) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *QueuePositionResponse) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *QueuePositionResponse) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetQueuePosition

`func (o *QueuePositionResponse) GetQueuePosition() int64`

GetQueuePosition returns the QueuePosition field if non-nil, zero value otherwise.

### GetQueuePositionOk

`func (o *QueuePositionResponse) GetQueuePositionOk() (*int64, bool)`

GetQueuePositionOk returns a tuple with the QueuePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuePosition

`func (o *QueuePositionResponse) SetQueuePosition(v int64)`

SetQueuePosition sets QueuePosition field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


