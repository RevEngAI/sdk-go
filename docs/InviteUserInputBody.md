# InviteUserInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credits** | **int64** | Initial credit balance for the invited user in raw units (1 credit &#x3D; 1000 units). | 
**Email** | **string** | Email address to invite | 
**FirstName** | **string** | First name included in the invite email and registration URL | 
**LastName** | Pointer to **string** | Last name pre-filled in the registration URL | [optional] 
**TeamId** | Pointer to **int64** | Team to assign the invited user to | [optional] 
**Username** | Pointer to **string** | Username pre-filled in the registration URL | [optional] 

## Methods

### NewInviteUserInputBody

`func NewInviteUserInputBody(credits int64, email string, firstName string, ) *InviteUserInputBody`

NewInviteUserInputBody instantiates a new InviteUserInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteUserInputBodyWithDefaults

`func NewInviteUserInputBodyWithDefaults() *InviteUserInputBody`

NewInviteUserInputBodyWithDefaults instantiates a new InviteUserInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredits

`func (o *InviteUserInputBody) GetCredits() int64`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *InviteUserInputBody) GetCreditsOk() (*int64, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *InviteUserInputBody) SetCredits(v int64)`

SetCredits sets Credits field to given value.


### GetEmail

`func (o *InviteUserInputBody) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InviteUserInputBody) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InviteUserInputBody) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *InviteUserInputBody) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *InviteUserInputBody) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *InviteUserInputBody) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *InviteUserInputBody) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *InviteUserInputBody) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *InviteUserInputBody) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *InviteUserInputBody) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetTeamId

`func (o *InviteUserInputBody) GetTeamId() int64`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *InviteUserInputBody) GetTeamIdOk() (*int64, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *InviteUserInputBody) SetTeamId(v int64)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *InviteUserInputBody) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetUsername

`func (o *InviteUserInputBody) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InviteUserInputBody) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InviteUserInputBody) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *InviteUserInputBody) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


