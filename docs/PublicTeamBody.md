# PublicTeamBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TeamId** | **int64** | The team&#39;s ID | 
**TeamName** | **string** | The team&#39;s name | 

## Methods

### NewPublicTeamBody

`func NewPublicTeamBody(teamId int64, teamName string, ) *PublicTeamBody`

NewPublicTeamBody instantiates a new PublicTeamBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicTeamBodyWithDefaults

`func NewPublicTeamBodyWithDefaults() *PublicTeamBody`

NewPublicTeamBodyWithDefaults instantiates a new PublicTeamBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeamId

`func (o *PublicTeamBody) GetTeamId() int64`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *PublicTeamBody) GetTeamIdOk() (*int64, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *PublicTeamBody) SetTeamId(v int64)`

SetTeamId sets TeamId field to given value.


### GetTeamName

`func (o *PublicTeamBody) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *PublicTeamBody) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *PublicTeamBody) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


