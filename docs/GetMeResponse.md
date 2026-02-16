# GetMeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**UserId** | **int32** |  | 
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**Email** | **string** |  | 
**Creation** | **time.Time** |  | 
**TutorialSeen** | **bool** |  | 
**Role** | **string** |  | 

## Methods

### NewGetMeResponse

`func NewGetMeResponse(username string, userId int32, firstName string, lastName string, email string, creation time.Time, tutorialSeen bool, role string, ) *GetMeResponse`

NewGetMeResponse instantiates a new GetMeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMeResponseWithDefaults

`func NewGetMeResponseWithDefaults() *GetMeResponse`

NewGetMeResponseWithDefaults instantiates a new GetMeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *GetMeResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetMeResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetMeResponse) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetUserId

`func (o *GetMeResponse) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetMeResponse) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetMeResponse) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetFirstName

`func (o *GetMeResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *GetMeResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *GetMeResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *GetMeResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *GetMeResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *GetMeResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetEmail

`func (o *GetMeResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetMeResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetMeResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetCreation

`func (o *GetMeResponse) GetCreation() time.Time`

GetCreation returns the Creation field if non-nil, zero value otherwise.

### GetCreationOk

`func (o *GetMeResponse) GetCreationOk() (*time.Time, bool)`

GetCreationOk returns a tuple with the Creation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreation

`func (o *GetMeResponse) SetCreation(v time.Time)`

SetCreation sets Creation field to given value.


### GetTutorialSeen

`func (o *GetMeResponse) GetTutorialSeen() bool`

GetTutorialSeen returns the TutorialSeen field if non-nil, zero value otherwise.

### GetTutorialSeenOk

`func (o *GetMeResponse) GetTutorialSeenOk() (*bool, bool)`

GetTutorialSeenOk returns a tuple with the TutorialSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTutorialSeen

`func (o *GetMeResponse) SetTutorialSeen(v bool)`

SetTutorialSeen sets TutorialSeen field to given value.


### GetRole

`func (o *GetMeResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GetMeResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GetMeResponse) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


