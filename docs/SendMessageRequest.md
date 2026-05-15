# SendMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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


