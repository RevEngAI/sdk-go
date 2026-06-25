# RegisterUserInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email address | 
**FirstName** | **string** | First name | 
**InviteCode** | Pointer to **string** | Invite code from the registration email | [optional] 
**LastName** | **string** | Last name | 
**Password** | **string** | Password (minimum 10 characters) | 
**Username** | **string** | Username | 

## Methods

### NewRegisterUserInputBody

`func NewRegisterUserInputBody(email string, firstName string, lastName string, password string, username string, ) *RegisterUserInputBody`

NewRegisterUserInputBody instantiates a new RegisterUserInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterUserInputBodyWithDefaults

`func NewRegisterUserInputBodyWithDefaults() *RegisterUserInputBody`

NewRegisterUserInputBodyWithDefaults instantiates a new RegisterUserInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *RegisterUserInputBody) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RegisterUserInputBody) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RegisterUserInputBody) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *RegisterUserInputBody) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *RegisterUserInputBody) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *RegisterUserInputBody) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetInviteCode

`func (o *RegisterUserInputBody) GetInviteCode() string`

GetInviteCode returns the InviteCode field if non-nil, zero value otherwise.

### GetInviteCodeOk

`func (o *RegisterUserInputBody) GetInviteCodeOk() (*string, bool)`

GetInviteCodeOk returns a tuple with the InviteCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteCode

`func (o *RegisterUserInputBody) SetInviteCode(v string)`

SetInviteCode sets InviteCode field to given value.

### HasInviteCode

`func (o *RegisterUserInputBody) HasInviteCode() bool`

HasInviteCode returns a boolean if a field has been set.

### GetLastName

`func (o *RegisterUserInputBody) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *RegisterUserInputBody) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *RegisterUserInputBody) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetPassword

`func (o *RegisterUserInputBody) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RegisterUserInputBody) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RegisterUserInputBody) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUsername

`func (o *RegisterUserInputBody) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RegisterUserInputBody) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RegisterUserInputBody) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


