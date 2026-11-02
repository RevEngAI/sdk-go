# ConversationBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisId** | **NullableInt64** | The analysis the conversation is about. Null when it is about the whole investigation. | 
**ConversationUuid** | **string** | Use with the /v2/conversations endpoints to read, stream and send messages. | 
**CreatedAt** | **time.Time** |  | 
**InvestigationId** | **int64** |  | 
**StarterId** | **int64** | The user who started the conversation, and the only one who can send messages in it. | 
**StarterUsername** | **string** |  | 
**Title** | **string** |  | 
**UpdatedAt** | **time.Time** | When the conversation was last active: its newest message or event, or a change such as a new title. | 

## Methods

### NewConversationBody

`func NewConversationBody(analysisId NullableInt64, conversationUuid string, createdAt time.Time, investigationId int64, starterId int64, starterUsername string, title string, updatedAt time.Time, ) *ConversationBody`

NewConversationBody instantiates a new ConversationBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationBodyWithDefaults

`func NewConversationBodyWithDefaults() *ConversationBody`

NewConversationBodyWithDefaults instantiates a new ConversationBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisId

`func (o *ConversationBody) GetAnalysisId() int64`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *ConversationBody) GetAnalysisIdOk() (*int64, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *ConversationBody) SetAnalysisId(v int64)`

SetAnalysisId sets AnalysisId field to given value.


### SetAnalysisIdNil

`func (o *ConversationBody) SetAnalysisIdNil(b bool)`

 SetAnalysisIdNil sets the value for AnalysisId to be an explicit nil

### UnsetAnalysisId
`func (o *ConversationBody) UnsetAnalysisId()`

UnsetAnalysisId ensures that no value is present for AnalysisId, not even an explicit nil
### GetConversationUuid

`func (o *ConversationBody) GetConversationUuid() string`

GetConversationUuid returns the ConversationUuid field if non-nil, zero value otherwise.

### GetConversationUuidOk

`func (o *ConversationBody) GetConversationUuidOk() (*string, bool)`

GetConversationUuidOk returns a tuple with the ConversationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationUuid

`func (o *ConversationBody) SetConversationUuid(v string)`

SetConversationUuid sets ConversationUuid field to given value.


### GetCreatedAt

`func (o *ConversationBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConversationBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConversationBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetInvestigationId

`func (o *ConversationBody) GetInvestigationId() int64`

GetInvestigationId returns the InvestigationId field if non-nil, zero value otherwise.

### GetInvestigationIdOk

`func (o *ConversationBody) GetInvestigationIdOk() (*int64, bool)`

GetInvestigationIdOk returns a tuple with the InvestigationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigationId

`func (o *ConversationBody) SetInvestigationId(v int64)`

SetInvestigationId sets InvestigationId field to given value.


### GetStarterId

`func (o *ConversationBody) GetStarterId() int64`

GetStarterId returns the StarterId field if non-nil, zero value otherwise.

### GetStarterIdOk

`func (o *ConversationBody) GetStarterIdOk() (*int64, bool)`

GetStarterIdOk returns a tuple with the StarterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarterId

`func (o *ConversationBody) SetStarterId(v int64)`

SetStarterId sets StarterId field to given value.


### GetStarterUsername

`func (o *ConversationBody) GetStarterUsername() string`

GetStarterUsername returns the StarterUsername field if non-nil, zero value otherwise.

### GetStarterUsernameOk

`func (o *ConversationBody) GetStarterUsernameOk() (*string, bool)`

GetStarterUsernameOk returns a tuple with the StarterUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarterUsername

`func (o *ConversationBody) SetStarterUsername(v string)`

SetStarterUsername sets StarterUsername field to given value.


### GetTitle

`func (o *ConversationBody) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ConversationBody) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ConversationBody) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUpdatedAt

`func (o *ConversationBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ConversationBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ConversationBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


