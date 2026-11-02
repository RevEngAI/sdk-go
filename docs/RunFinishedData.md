# RunFinishedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversationUuid** | Pointer to **string** |  | [optional] 
**DefinitionName** | Pointer to **string** |  | [optional] 
**InputTokens** | Pointer to **int32** |  | [optional] 
**OutputTokens** | Pointer to **int32** |  | [optional] 
**TotalTokens** | Pointer to **int32** |  | [optional] 

## Methods

### NewRunFinishedData

`func NewRunFinishedData() *RunFinishedData`

NewRunFinishedData instantiates a new RunFinishedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunFinishedDataWithDefaults

`func NewRunFinishedDataWithDefaults() *RunFinishedData`

NewRunFinishedDataWithDefaults instantiates a new RunFinishedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversationUuid

`func (o *RunFinishedData) GetConversationUuid() string`

GetConversationUuid returns the ConversationUuid field if non-nil, zero value otherwise.

### GetConversationUuidOk

`func (o *RunFinishedData) GetConversationUuidOk() (*string, bool)`

GetConversationUuidOk returns a tuple with the ConversationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationUuid

`func (o *RunFinishedData) SetConversationUuid(v string)`

SetConversationUuid sets ConversationUuid field to given value.

### HasConversationUuid

`func (o *RunFinishedData) HasConversationUuid() bool`

HasConversationUuid returns a boolean if a field has been set.

### GetDefinitionName

`func (o *RunFinishedData) GetDefinitionName() string`

GetDefinitionName returns the DefinitionName field if non-nil, zero value otherwise.

### GetDefinitionNameOk

`func (o *RunFinishedData) GetDefinitionNameOk() (*string, bool)`

GetDefinitionNameOk returns a tuple with the DefinitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionName

`func (o *RunFinishedData) SetDefinitionName(v string)`

SetDefinitionName sets DefinitionName field to given value.

### HasDefinitionName

`func (o *RunFinishedData) HasDefinitionName() bool`

HasDefinitionName returns a boolean if a field has been set.

### GetInputTokens

`func (o *RunFinishedData) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *RunFinishedData) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *RunFinishedData) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *RunFinishedData) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### GetOutputTokens

`func (o *RunFinishedData) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *RunFinishedData) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *RunFinishedData) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.

### HasOutputTokens

`func (o *RunFinishedData) HasOutputTokens() bool`

HasOutputTokens returns a boolean if a field has been set.

### GetTotalTokens

`func (o *RunFinishedData) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *RunFinishedData) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *RunFinishedData) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *RunFinishedData) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


