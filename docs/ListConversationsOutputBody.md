# ListConversationsOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNextPage** | **bool** |  | 
**Results** | [**[]ConversationBody**](ConversationBody.md) |  | 
**TotalCount** | **int64** | Conversations in the investigation across all pages. | 

## Methods

### NewListConversationsOutputBody

`func NewListConversationsOutputBody(hasNextPage bool, results []ConversationBody, totalCount int64, ) *ListConversationsOutputBody`

NewListConversationsOutputBody instantiates a new ListConversationsOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConversationsOutputBodyWithDefaults

`func NewListConversationsOutputBodyWithDefaults() *ListConversationsOutputBody`

NewListConversationsOutputBodyWithDefaults instantiates a new ListConversationsOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNextPage

`func (o *ListConversationsOutputBody) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *ListConversationsOutputBody) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *ListConversationsOutputBody) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.


### GetResults

`func (o *ListConversationsOutputBody) GetResults() []ConversationBody`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListConversationsOutputBody) GetResultsOk() (*[]ConversationBody, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListConversationsOutputBody) SetResults(v []ConversationBody)`

SetResults sets Results field to given value.


### GetTotalCount

`func (o *ListConversationsOutputBody) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListConversationsOutputBody) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListConversationsOutputBody) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


