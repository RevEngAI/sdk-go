# UserProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultTeamId** | Pointer to **int64** |  | [optional] 
**FirstName** | **string** |  | 
**HideExampleBinaries** | **bool** |  | 
**LastName** | **string** |  | 
**TimeZone** | **string** |  | 
**Username** | **string** |  | 

## Methods

### NewUserProfile

`func NewUserProfile(firstName string, hideExampleBinaries bool, lastName string, timeZone string, username string, ) *UserProfile`

NewUserProfile instantiates a new UserProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserProfileWithDefaults

`func NewUserProfileWithDefaults() *UserProfile`

NewUserProfileWithDefaults instantiates a new UserProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultTeamId

`func (o *UserProfile) GetDefaultTeamId() int64`

GetDefaultTeamId returns the DefaultTeamId field if non-nil, zero value otherwise.

### GetDefaultTeamIdOk

`func (o *UserProfile) GetDefaultTeamIdOk() (*int64, bool)`

GetDefaultTeamIdOk returns a tuple with the DefaultTeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTeamId

`func (o *UserProfile) SetDefaultTeamId(v int64)`

SetDefaultTeamId sets DefaultTeamId field to given value.

### HasDefaultTeamId

`func (o *UserProfile) HasDefaultTeamId() bool`

HasDefaultTeamId returns a boolean if a field has been set.

### GetFirstName

`func (o *UserProfile) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserProfile) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserProfile) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetHideExampleBinaries

`func (o *UserProfile) GetHideExampleBinaries() bool`

GetHideExampleBinaries returns the HideExampleBinaries field if non-nil, zero value otherwise.

### GetHideExampleBinariesOk

`func (o *UserProfile) GetHideExampleBinariesOk() (*bool, bool)`

GetHideExampleBinariesOk returns a tuple with the HideExampleBinaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideExampleBinaries

`func (o *UserProfile) SetHideExampleBinaries(v bool)`

SetHideExampleBinaries sets HideExampleBinaries field to given value.


### GetLastName

`func (o *UserProfile) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserProfile) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserProfile) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetTimeZone

`func (o *UserProfile) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *UserProfile) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *UserProfile) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.


### GetUsername

`func (o *UserProfile) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserProfile) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserProfile) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


