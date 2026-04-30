# ConversationWithEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**Context** | Pointer to **interface{}** |  | [optional] 
**ConversationUuid** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Events** | [**[]Event**](Event.md) |  | 
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

### GetSchema

`func (o *ConversationWithEvents) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *ConversationWithEvents) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *ConversationWithEvents) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *ConversationWithEvents) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetContext

`func (o *ConversationWithEvents) GetContext() interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ConversationWithEvents) GetContextOk() (*interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ConversationWithEvents) SetContext(v interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *ConversationWithEvents) HasContext() bool`

HasContext returns a boolean if a field has been set.

### SetContextNil

`func (o *ConversationWithEvents) SetContextNil(b bool)`

 SetContextNil sets the value for Context to be an explicit nil

### UnsetContext
`func (o *ConversationWithEvents) UnsetContext()`

UnsetContext ensures that no value is present for Context, not even an explicit nil
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


