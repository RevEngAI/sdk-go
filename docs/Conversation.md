# Conversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **interface{}** |  | [optional] 
**ConversationUuid** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Title** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 
**UserId** | **int64** |  | 

## Methods

### NewConversation

`func NewConversation(conversationUuid string, createdAt time.Time, title string, updatedAt time.Time, userId int64, ) *Conversation`

NewConversation instantiates a new Conversation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationWithDefaults

`func NewConversationWithDefaults() *Conversation`

NewConversationWithDefaults instantiates a new Conversation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *Conversation) GetContext() interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Conversation) GetContextOk() (*interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Conversation) SetContext(v interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *Conversation) HasContext() bool`

HasContext returns a boolean if a field has been set.

### SetContextNil

`func (o *Conversation) SetContextNil(b bool)`

 SetContextNil sets the value for Context to be an explicit nil

### UnsetContext
`func (o *Conversation) UnsetContext()`

UnsetContext ensures that no value is present for Context, not even an explicit nil
### GetConversationUuid

`func (o *Conversation) GetConversationUuid() string`

GetConversationUuid returns the ConversationUuid field if non-nil, zero value otherwise.

### GetConversationUuidOk

`func (o *Conversation) GetConversationUuidOk() (*string, bool)`

GetConversationUuidOk returns a tuple with the ConversationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationUuid

`func (o *Conversation) SetConversationUuid(v string)`

SetConversationUuid sets ConversationUuid field to given value.


### GetCreatedAt

`func (o *Conversation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Conversation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Conversation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetTitle

`func (o *Conversation) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Conversation) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Conversation) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUpdatedAt

`func (o *Conversation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Conversation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Conversation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *Conversation) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Conversation) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Conversation) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


