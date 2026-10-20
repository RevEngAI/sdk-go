# ListPublicUsersOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]PublicUserBody**](PublicUserBody.md) | Up to 5 matching active users, ordered by username | 

## Methods

### NewListPublicUsersOutputBody

`func NewListPublicUsersOutputBody(results []PublicUserBody, ) *ListPublicUsersOutputBody`

NewListPublicUsersOutputBody instantiates a new ListPublicUsersOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPublicUsersOutputBodyWithDefaults

`func NewListPublicUsersOutputBodyWithDefaults() *ListPublicUsersOutputBody`

NewListPublicUsersOutputBodyWithDefaults instantiates a new ListPublicUsersOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ListPublicUsersOutputBody) GetResults() []PublicUserBody`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListPublicUsersOutputBody) GetResultsOk() (*[]PublicUserBody, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListPublicUsersOutputBody) SetResults(v []PublicUserBody)`

SetResults sets Results field to given value.


### SetResultsNil

`func (o *ListPublicUsersOutputBody) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *ListPublicUsersOutputBody) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


