# GetDieInfoOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Matches** | [**[]DieMatch**](DieMatch.md) | Signatures Detect It Easy recognised in the binary | 

## Methods

### NewGetDieInfoOutputBody

`func NewGetDieInfoOutputBody(matches []DieMatch, ) *GetDieInfoOutputBody`

NewGetDieInfoOutputBody instantiates a new GetDieInfoOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDieInfoOutputBodyWithDefaults

`func NewGetDieInfoOutputBodyWithDefaults() *GetDieInfoOutputBody`

NewGetDieInfoOutputBodyWithDefaults instantiates a new GetDieInfoOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatches

`func (o *GetDieInfoOutputBody) GetMatches() []DieMatch`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *GetDieInfoOutputBody) GetMatchesOk() (*[]DieMatch, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *GetDieInfoOutputBody) SetMatches(v []DieMatch)`

SetMatches sets Matches field to given value.


### SetMatchesNil

`func (o *GetDieInfoOutputBody) SetMatchesNil(b bool)`

 SetMatchesNil sets the value for Matches to be an explicit nil

### UnsetMatches
`func (o *GetDieInfoOutputBody) UnsetMatches()`

UnsetMatches ensures that no value is present for Matches, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


