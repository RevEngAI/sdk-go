# CreateConversationInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisId** | Pointer to **int64** | Scopes the conversation to one of the investigation&#39;s analyses. Omit it to cover the whole investigation. | [optional] 
**Title** | Pointer to **string** | A title for the conversation. When empty, one is generated from the first message. | [optional] 

## Methods

### NewCreateConversationInputBody

`func NewCreateConversationInputBody() *CreateConversationInputBody`

NewCreateConversationInputBody instantiates a new CreateConversationInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConversationInputBodyWithDefaults

`func NewCreateConversationInputBodyWithDefaults() *CreateConversationInputBody`

NewCreateConversationInputBodyWithDefaults instantiates a new CreateConversationInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisId

`func (o *CreateConversationInputBody) GetAnalysisId() int64`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *CreateConversationInputBody) GetAnalysisIdOk() (*int64, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *CreateConversationInputBody) SetAnalysisId(v int64)`

SetAnalysisId sets AnalysisId field to given value.

### HasAnalysisId

`func (o *CreateConversationInputBody) HasAnalysisId() bool`

HasAnalysisId returns a boolean if a field has been set.

### GetTitle

`func (o *CreateConversationInputBody) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateConversationInputBody) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateConversationInputBody) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateConversationInputBody) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


