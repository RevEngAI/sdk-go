# AccessRuleBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permission** | **string** | VIEW lets the principal read the analysis; EDIT also lets them modify it | 
**TeamId** | Pointer to **int64** | The team the rule applies to. Set for a team rule | [optional] 
**TeamName** | Pointer to **string** | The team&#39;s name. Set for a team rule | [optional] 
**UserId** | Pointer to **int64** | The user the rule applies to. Set for a user rule | [optional] 
**Username** | Pointer to **string** | The user&#39;s username. Set for a user rule | [optional] 

## Methods

### NewAccessRuleBody

`func NewAccessRuleBody(permission string, ) *AccessRuleBody`

NewAccessRuleBody instantiates a new AccessRuleBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRuleBodyWithDefaults

`func NewAccessRuleBodyWithDefaults() *AccessRuleBody`

NewAccessRuleBodyWithDefaults instantiates a new AccessRuleBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermission

`func (o *AccessRuleBody) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *AccessRuleBody) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *AccessRuleBody) SetPermission(v string)`

SetPermission sets Permission field to given value.


### GetTeamId

`func (o *AccessRuleBody) GetTeamId() int64`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *AccessRuleBody) GetTeamIdOk() (*int64, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *AccessRuleBody) SetTeamId(v int64)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *AccessRuleBody) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetTeamName

`func (o *AccessRuleBody) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *AccessRuleBody) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *AccessRuleBody) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.

### HasTeamName

`func (o *AccessRuleBody) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.

### GetUserId

`func (o *AccessRuleBody) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AccessRuleBody) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AccessRuleBody) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AccessRuleBody) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsername

`func (o *AccessRuleBody) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AccessRuleBody) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AccessRuleBody) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AccessRuleBody) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


