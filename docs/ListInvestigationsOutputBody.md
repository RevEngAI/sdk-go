# ListInvestigationsOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentStatusCounts** | [**AgentStatusCountsBody**](AgentStatusCountsBody.md) | Open investigations you can see at each agent status, whatever the filters and page. | 
**HasNextPage** | **bool** |  | 
**Results** | [**[]InvestigationListItemBody**](InvestigationListItemBody.md) |  | 
**TotalCount** | **int64** | Investigations matching the filters across all pages. | 

## Methods

### NewListInvestigationsOutputBody

`func NewListInvestigationsOutputBody(agentStatusCounts AgentStatusCountsBody, hasNextPage bool, results []InvestigationListItemBody, totalCount int64, ) *ListInvestigationsOutputBody`

NewListInvestigationsOutputBody instantiates a new ListInvestigationsOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInvestigationsOutputBodyWithDefaults

`func NewListInvestigationsOutputBodyWithDefaults() *ListInvestigationsOutputBody`

NewListInvestigationsOutputBodyWithDefaults instantiates a new ListInvestigationsOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentStatusCounts

`func (o *ListInvestigationsOutputBody) GetAgentStatusCounts() AgentStatusCountsBody`

GetAgentStatusCounts returns the AgentStatusCounts field if non-nil, zero value otherwise.

### GetAgentStatusCountsOk

`func (o *ListInvestigationsOutputBody) GetAgentStatusCountsOk() (*AgentStatusCountsBody, bool)`

GetAgentStatusCountsOk returns a tuple with the AgentStatusCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentStatusCounts

`func (o *ListInvestigationsOutputBody) SetAgentStatusCounts(v AgentStatusCountsBody)`

SetAgentStatusCounts sets AgentStatusCounts field to given value.


### GetHasNextPage

`func (o *ListInvestigationsOutputBody) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *ListInvestigationsOutputBody) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *ListInvestigationsOutputBody) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.


### GetResults

`func (o *ListInvestigationsOutputBody) GetResults() []InvestigationListItemBody`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListInvestigationsOutputBody) GetResultsOk() (*[]InvestigationListItemBody, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListInvestigationsOutputBody) SetResults(v []InvestigationListItemBody)`

SetResults sets Results field to given value.


### GetTotalCount

`func (o *ListInvestigationsOutputBody) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListInvestigationsOutputBody) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListInvestigationsOutputBody) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


