# SendMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**Content** | **string** |  | 
**Context** | Pointer to [**ConversationContext**](ConversationContext.md) |  | [optional] 

## Methods

### NewSendMessageRequest

`func NewSendMessageRequest(content string, ) *SendMessageRequest`

NewSendMessageRequest instantiates a new SendMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendMessageRequestWithDefaults

`func NewSendMessageRequestWithDefaults() *SendMessageRequest`

NewSendMessageRequestWithDefaults instantiates a new SendMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *SendMessageRequest) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *SendMessageRequest) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *SendMessageRequest) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *SendMessageRequest) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetContent

`func (o *SendMessageRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SendMessageRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SendMessageRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetContext

`func (o *SendMessageRequest) GetContext() ConversationContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *SendMessageRequest) GetContextOk() (*ConversationContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *SendMessageRequest) SetContext(v ConversationContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *SendMessageRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


