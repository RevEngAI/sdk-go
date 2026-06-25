# CreateUserInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credits** | Pointer to **int64** | Initial credit balance in credits (1 credit &#x3D; 1000 units); defaults to 10 credits | [optional] 
**Email** | **string** | Email address | 
**FirstName** | Pointer to **string** | First name | [optional] 
**LastName** | Pointer to **string** | Last name | [optional] 
**Password** | **string** | Initial password | 
**Role** | Pointer to **string** | User role (defaults to USER) | [optional] 
**Tier** | Pointer to **string** | User tier (defaults to ENTHUSIAST) | [optional] 
**TimeZone** | Pointer to **string** | IANA time zone | [optional] 
**Username** | **string** | Username | 

## Methods

### NewCreateUserInputBody

`func NewCreateUserInputBody(email string, password string, username string, ) *CreateUserInputBody`

NewCreateUserInputBody instantiates a new CreateUserInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserInputBodyWithDefaults

`func NewCreateUserInputBodyWithDefaults() *CreateUserInputBody`

NewCreateUserInputBodyWithDefaults instantiates a new CreateUserInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredits

`func (o *CreateUserInputBody) GetCredits() int64`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *CreateUserInputBody) GetCreditsOk() (*int64, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *CreateUserInputBody) SetCredits(v int64)`

SetCredits sets Credits field to given value.

### HasCredits

`func (o *CreateUserInputBody) HasCredits() bool`

HasCredits returns a boolean if a field has been set.

### GetEmail

`func (o *CreateUserInputBody) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateUserInputBody) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateUserInputBody) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *CreateUserInputBody) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CreateUserInputBody) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CreateUserInputBody) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CreateUserInputBody) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *CreateUserInputBody) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CreateUserInputBody) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CreateUserInputBody) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CreateUserInputBody) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPassword

`func (o *CreateUserInputBody) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateUserInputBody) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateUserInputBody) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetRole

`func (o *CreateUserInputBody) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CreateUserInputBody) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CreateUserInputBody) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *CreateUserInputBody) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTier

`func (o *CreateUserInputBody) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *CreateUserInputBody) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *CreateUserInputBody) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *CreateUserInputBody) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetTimeZone

`func (o *CreateUserInputBody) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *CreateUserInputBody) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *CreateUserInputBody) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *CreateUserInputBody) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetUsername

`func (o *CreateUserInputBody) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateUserInputBody) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateUserInputBody) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


