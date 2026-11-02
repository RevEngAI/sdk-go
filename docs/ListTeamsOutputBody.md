# ListTeamsOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]PublicTeamBody**](PublicTeamBody.md) | Up to 5 matching teams, ordered by name | 

## Methods

### NewListTeamsOutputBody

`func NewListTeamsOutputBody(results []PublicTeamBody, ) *ListTeamsOutputBody`

NewListTeamsOutputBody instantiates a new ListTeamsOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTeamsOutputBodyWithDefaults

`func NewListTeamsOutputBodyWithDefaults() *ListTeamsOutputBody`

NewListTeamsOutputBodyWithDefaults instantiates a new ListTeamsOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ListTeamsOutputBody) GetResults() []PublicTeamBody`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListTeamsOutputBody) GetResultsOk() (*[]PublicTeamBody, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListTeamsOutputBody) SetResults(v []PublicTeamBody)`

SetResults sets Results field to given value.


### SetResultsNil

`func (o *ListTeamsOutputBody) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *ListTeamsOutputBody) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


