# GetGlobalHistoryOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**History** | [**[]GlobalHistoryEntry**](GlobalHistoryEntry.md) | Replaced values, oldest first. | 

## Methods

### NewGetGlobalHistoryOutputBody

`func NewGetGlobalHistoryOutputBody(history []GlobalHistoryEntry, ) *GetGlobalHistoryOutputBody`

NewGetGlobalHistoryOutputBody instantiates a new GetGlobalHistoryOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGlobalHistoryOutputBodyWithDefaults

`func NewGetGlobalHistoryOutputBodyWithDefaults() *GetGlobalHistoryOutputBody`

NewGetGlobalHistoryOutputBodyWithDefaults instantiates a new GetGlobalHistoryOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistory

`func (o *GetGlobalHistoryOutputBody) GetHistory() []GlobalHistoryEntry`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *GetGlobalHistoryOutputBody) GetHistoryOk() (*[]GlobalHistoryEntry, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *GetGlobalHistoryOutputBody) SetHistory(v []GlobalHistoryEntry)`

SetHistory sets History field to given value.


### SetHistoryNil

`func (o *GetGlobalHistoryOutputBody) SetHistoryNil(b bool)`

 SetHistoryNil sets the value for History to be an explicit nil

### UnsetHistory
`func (o *GetGlobalHistoryOutputBody) UnsetHistory()`

UnsetHistory ensures that no value is present for History, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


