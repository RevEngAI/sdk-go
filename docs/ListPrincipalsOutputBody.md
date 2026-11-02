# ListPrincipalsOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]PrincipalBody**](PrincipalBody.md) | Up to 5 matching users followed by up to 5 matching teams, each ordered by name | 

## Methods

### NewListPrincipalsOutputBody

`func NewListPrincipalsOutputBody(results []PrincipalBody, ) *ListPrincipalsOutputBody`

NewListPrincipalsOutputBody instantiates a new ListPrincipalsOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPrincipalsOutputBodyWithDefaults

`func NewListPrincipalsOutputBodyWithDefaults() *ListPrincipalsOutputBody`

NewListPrincipalsOutputBodyWithDefaults instantiates a new ListPrincipalsOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ListPrincipalsOutputBody) GetResults() []PrincipalBody`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListPrincipalsOutputBody) GetResultsOk() (*[]PrincipalBody, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListPrincipalsOutputBody) SetResults(v []PrincipalBody)`

SetResults sets Results field to given value.


### SetResultsNil

`func (o *ListPrincipalsOutputBody) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *ListPrincipalsOutputBody) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


