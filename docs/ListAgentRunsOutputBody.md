# ListAgentRunsOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNextPage** | **bool** |  | 
**Results** | [**[]AgentRunBody**](AgentRunBody.md) |  | 
**TotalCount** | **int64** | Runs matching the filters across all pages. | 

## Methods

### NewListAgentRunsOutputBody

`func NewListAgentRunsOutputBody(hasNextPage bool, results []AgentRunBody, totalCount int64, ) *ListAgentRunsOutputBody`

NewListAgentRunsOutputBody instantiates a new ListAgentRunsOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAgentRunsOutputBodyWithDefaults

`func NewListAgentRunsOutputBodyWithDefaults() *ListAgentRunsOutputBody`

NewListAgentRunsOutputBodyWithDefaults instantiates a new ListAgentRunsOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNextPage

`func (o *ListAgentRunsOutputBody) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *ListAgentRunsOutputBody) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *ListAgentRunsOutputBody) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.


### GetResults

`func (o *ListAgentRunsOutputBody) GetResults() []AgentRunBody`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListAgentRunsOutputBody) GetResultsOk() (*[]AgentRunBody, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListAgentRunsOutputBody) SetResults(v []AgentRunBody)`

SetResults sets Results field to given value.


### GetTotalCount

`func (o *ListAgentRunsOutputBody) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListAgentRunsOutputBody) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListAgentRunsOutputBody) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


