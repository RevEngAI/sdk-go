# GrantAccessInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permission** | **string** | The level of access to grant | 
**TeamId** | Pointer to **int64** | The team to share with. Exactly one of user_id and team_id is required | [optional] 
**UserId** | Pointer to **int64** | The user to share with. Exactly one of user_id and team_id is required | [optional] 

## Methods

### NewGrantAccessInputBody

`func NewGrantAccessInputBody(permission string, ) *GrantAccessInputBody`

NewGrantAccessInputBody instantiates a new GrantAccessInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantAccessInputBodyWithDefaults

`func NewGrantAccessInputBodyWithDefaults() *GrantAccessInputBody`

NewGrantAccessInputBodyWithDefaults instantiates a new GrantAccessInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermission

`func (o *GrantAccessInputBody) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *GrantAccessInputBody) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *GrantAccessInputBody) SetPermission(v string)`

SetPermission sets Permission field to given value.


### GetTeamId

`func (o *GrantAccessInputBody) GetTeamId() int64`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *GrantAccessInputBody) GetTeamIdOk() (*int64, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *GrantAccessInputBody) SetTeamId(v int64)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *GrantAccessInputBody) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetUserId

`func (o *GrantAccessInputBody) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GrantAccessInputBody) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GrantAccessInputBody) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GrantAccessInputBody) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


