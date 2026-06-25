# ListTeamsOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Teams** | [**[]Team**](Team.md) |  | 
**Total** | **int64** |  | 

## Methods

### NewListTeamsOutputBody

`func NewListTeamsOutputBody(teams []Team, total int64, ) *ListTeamsOutputBody`

NewListTeamsOutputBody instantiates a new ListTeamsOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTeamsOutputBodyWithDefaults

`func NewListTeamsOutputBodyWithDefaults() *ListTeamsOutputBody`

NewListTeamsOutputBodyWithDefaults instantiates a new ListTeamsOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeams

`func (o *ListTeamsOutputBody) GetTeams() []Team`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *ListTeamsOutputBody) GetTeamsOk() (*[]Team, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *ListTeamsOutputBody) SetTeams(v []Team)`

SetTeams sets Teams field to given value.


### SetTeamsNil

`func (o *ListTeamsOutputBody) SetTeamsNil(b bool)`

 SetTeamsNil sets the value for Teams to be an explicit nil

### UnsetTeams
`func (o *ListTeamsOutputBody) UnsetTeams()`

UnsetTeams ensures that no value is present for Teams, not even an explicit nil
### GetTotal

`func (o *ListTeamsOutputBody) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListTeamsOutputBody) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListTeamsOutputBody) SetTotal(v int64)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


