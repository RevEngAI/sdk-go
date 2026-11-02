# ConversationWithEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**ConversationContext**](ConversationContext.md) | The analysis, and optionally the function, the conversation is about. Absent when it is about neither. | [optional] 
**ConversationUuid** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Events** | [**[]Event**](Event.md) |  | 
**InvestigationId** | Pointer to **int64** | The investigation the conversation belongs to. Absent for a conversation outside any investigation. | [optional] 
**Title** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 
**UserId** | **int64** |  | 

## Methods

### NewConversationWithEvents

`func NewConversationWithEvents(conversationUuid string, createdAt time.Time, events []Event, title string, updatedAt time.Time, userId int64, ) *ConversationWithEvents`

NewConversationWithEvents instantiates a new ConversationWithEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationWithEventsWithDefaults

`func NewConversationWithEventsWithDefaults() *ConversationWithEvents`

NewConversationWithEventsWithDefaults instantiates a new ConversationWithEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *ConversationWithEvents) GetContext() ConversationContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ConversationWithEvents) GetContextOk() (*ConversationContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ConversationWithEvents) SetContext(v ConversationContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *ConversationWithEvents) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetConversationUuid

`func (o *ConversationWithEvents) GetConversationUuid() string`

GetConversationUuid returns the ConversationUuid field if non-nil, zero value otherwise.

### GetConversationUuidOk

`func (o *ConversationWithEvents) GetConversationUuidOk() (*string, bool)`

GetConversationUuidOk returns a tuple with the ConversationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationUuid

`func (o *ConversationWithEvents) SetConversationUuid(v string)`

SetConversationUuid sets ConversationUuid field to given value.


### GetCreatedAt

`func (o *ConversationWithEvents) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConversationWithEvents) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConversationWithEvents) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEvents

`func (o *ConversationWithEvents) GetEvents() []Event`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ConversationWithEvents) GetEventsOk() (*[]Event, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ConversationWithEvents) SetEvents(v []Event)`

SetEvents sets Events field to given value.


### SetEventsNil

`func (o *ConversationWithEvents) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *ConversationWithEvents) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetInvestigationId

`func (o *ConversationWithEvents) GetInvestigationId() int64`

GetInvestigationId returns the InvestigationId field if non-nil, zero value otherwise.

### GetInvestigationIdOk

`func (o *ConversationWithEvents) GetInvestigationIdOk() (*int64, bool)`

GetInvestigationIdOk returns a tuple with the InvestigationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigationId

`func (o *ConversationWithEvents) SetInvestigationId(v int64)`

SetInvestigationId sets InvestigationId field to given value.

### HasInvestigationId

`func (o *ConversationWithEvents) HasInvestigationId() bool`

HasInvestigationId returns a boolean if a field has been set.

### GetTitle

`func (o *ConversationWithEvents) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ConversationWithEvents) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ConversationWithEvents) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUpdatedAt

`func (o *ConversationWithEvents) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ConversationWithEvents) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ConversationWithEvents) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *ConversationWithEvents) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ConversationWithEvents) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ConversationWithEvents) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


