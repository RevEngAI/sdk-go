# Conversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**ConversationContext**](ConversationContext.md) | The analysis, and optionally the function, the conversation is about. Absent when it is about neither. | [optional] 
**ConversationUuid** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**InvestigationId** | Pointer to **int64** | The investigation the conversation belongs to. Absent for a conversation outside any investigation. | [optional] 
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

`func (o *Conversation) GetContext() ConversationContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Conversation) GetContextOk() (*ConversationContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Conversation) SetContext(v ConversationContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *Conversation) HasContext() bool`

HasContext returns a boolean if a field has been set.

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


### GetInvestigationId

`func (o *Conversation) GetInvestigationId() int64`

GetInvestigationId returns the InvestigationId field if non-nil, zero value otherwise.

### GetInvestigationIdOk

`func (o *Conversation) GetInvestigationIdOk() (*int64, bool)`

GetInvestigationIdOk returns a tuple with the InvestigationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigationId

`func (o *Conversation) SetInvestigationId(v int64)`

SetInvestigationId sets InvestigationId field to given value.

### HasInvestigationId

`func (o *Conversation) HasInvestigationId() bool`

HasInvestigationId returns a boolean if a field has been set.

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


