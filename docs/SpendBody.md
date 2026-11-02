# SpendBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversationUuid** | **NullableString** | The conversation of a CHAT spend. | 
**CreatedAt** | **time.Time** |  | 
**Credits** | **float64** |  | 
**Source** | **string** | What the credits were spent on: a conversation turn or an agent run started from the investigation. | 
**SpendId** | **int64** |  | 
**TaskId** | **NullableString** | The agent run of an AGENT_RUN spend. | 
**UserId** | **int64** | Who was charged. Zero when that user has been deleted. | 
**Username** | **string** |  | 

## Methods

### NewSpendBody

`func NewSpendBody(conversationUuid NullableString, createdAt time.Time, credits float64, source string, spendId int64, taskId NullableString, userId int64, username string, ) *SpendBody`

NewSpendBody instantiates a new SpendBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpendBodyWithDefaults

`func NewSpendBodyWithDefaults() *SpendBody`

NewSpendBodyWithDefaults instantiates a new SpendBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversationUuid

`func (o *SpendBody) GetConversationUuid() string`

GetConversationUuid returns the ConversationUuid field if non-nil, zero value otherwise.

### GetConversationUuidOk

`func (o *SpendBody) GetConversationUuidOk() (*string, bool)`

GetConversationUuidOk returns a tuple with the ConversationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationUuid

`func (o *SpendBody) SetConversationUuid(v string)`

SetConversationUuid sets ConversationUuid field to given value.


### SetConversationUuidNil

`func (o *SpendBody) SetConversationUuidNil(b bool)`

 SetConversationUuidNil sets the value for ConversationUuid to be an explicit nil

### UnsetConversationUuid
`func (o *SpendBody) UnsetConversationUuid()`

UnsetConversationUuid ensures that no value is present for ConversationUuid, not even an explicit nil
### GetCreatedAt

`func (o *SpendBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SpendBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SpendBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCredits

`func (o *SpendBody) GetCredits() float64`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *SpendBody) GetCreditsOk() (*float64, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *SpendBody) SetCredits(v float64)`

SetCredits sets Credits field to given value.


### GetSource

`func (o *SpendBody) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SpendBody) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SpendBody) SetSource(v string)`

SetSource sets Source field to given value.


### GetSpendId

`func (o *SpendBody) GetSpendId() int64`

GetSpendId returns the SpendId field if non-nil, zero value otherwise.

### GetSpendIdOk

`func (o *SpendBody) GetSpendIdOk() (*int64, bool)`

GetSpendIdOk returns a tuple with the SpendId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendId

`func (o *SpendBody) SetSpendId(v int64)`

SetSpendId sets SpendId field to given value.


### GetTaskId

`func (o *SpendBody) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *SpendBody) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *SpendBody) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### SetTaskIdNil

`func (o *SpendBody) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *SpendBody) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetUserId

`func (o *SpendBody) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SpendBody) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SpendBody) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetUsername

`func (o *SpendBody) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SpendBody) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SpendBody) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


