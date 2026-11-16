# CurrentUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Email** | **string** |  | 
**IsActive** | **bool** | Inactive users cannot authenticate | 
**Profile** | [**UserProfile**](UserProfile.md) |  | 
**Role** | **string** |  | 
**SignsInWith** | **string** | How the user re-authenticates to delete their account: &#39;password&#39; sends their current password, &#39;sso&#39; signs in again at their identity provider. | 
**Tier** | Pointer to **string** |  | [optional] 
**UserId** | **int64** |  | 

## Methods

### NewCurrentUser

`func NewCurrentUser(createdAt time.Time, email string, isActive bool, profile UserProfile, role string, signsInWith string, userId int64, ) *CurrentUser`

NewCurrentUser instantiates a new CurrentUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentUserWithDefaults

`func NewCurrentUserWithDefaults() *CurrentUser`

NewCurrentUserWithDefaults instantiates a new CurrentUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CurrentUser) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CurrentUser) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CurrentUser) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEmail

`func (o *CurrentUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CurrentUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CurrentUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetIsActive

`func (o *CurrentUser) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CurrentUser) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CurrentUser) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetProfile

`func (o *CurrentUser) GetProfile() UserProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *CurrentUser) GetProfileOk() (*UserProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *CurrentUser) SetProfile(v UserProfile)`

SetProfile sets Profile field to given value.


### GetRole

`func (o *CurrentUser) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CurrentUser) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CurrentUser) SetRole(v string)`

SetRole sets Role field to given value.


### GetSignsInWith

`func (o *CurrentUser) GetSignsInWith() string`

GetSignsInWith returns the SignsInWith field if non-nil, zero value otherwise.

### GetSignsInWithOk

`func (o *CurrentUser) GetSignsInWithOk() (*string, bool)`

GetSignsInWithOk returns a tuple with the SignsInWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignsInWith

`func (o *CurrentUser) SetSignsInWith(v string)`

SetSignsInWith sets SignsInWith field to given value.


### GetTier

`func (o *CurrentUser) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *CurrentUser) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *CurrentUser) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *CurrentUser) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetUserId

`func (o *CurrentUser) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CurrentUser) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CurrentUser) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


