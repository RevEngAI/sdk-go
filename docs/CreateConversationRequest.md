# CreateConversationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**ConversationContext**](ConversationContext.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateConversationRequest

`func NewCreateConversationRequest() *CreateConversationRequest`

NewCreateConversationRequest instantiates a new CreateConversationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConversationRequestWithDefaults

`func NewCreateConversationRequestWithDefaults() *CreateConversationRequest`

NewCreateConversationRequestWithDefaults instantiates a new CreateConversationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *CreateConversationRequest) GetContext() ConversationContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *CreateConversationRequest) GetContextOk() (*ConversationContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *CreateConversationRequest) SetContext(v ConversationContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *CreateConversationRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetTitle

`func (o *CreateConversationRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateConversationRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateConversationRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateConversationRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


