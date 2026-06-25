# UserIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**IssuerUrl** | **NullableString** |  | 
**Subject** | **NullableString** |  | 
**UpdatedAt** | **time.Time** |  | 
**UserId** | **int64** |  | 
**UserIdentityId** | **int64** |  | 

## Methods

### NewUserIdentity

`func NewUserIdentity(createdAt time.Time, issuerUrl NullableString, subject NullableString, updatedAt time.Time, userId int64, userIdentityId int64, ) *UserIdentity`

NewUserIdentity instantiates a new UserIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserIdentityWithDefaults

`func NewUserIdentityWithDefaults() *UserIdentity`

NewUserIdentityWithDefaults instantiates a new UserIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *UserIdentity) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserIdentity) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserIdentity) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetIssuerUrl

`func (o *UserIdentity) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### GetIssuerUrlOk

`func (o *UserIdentity) GetIssuerUrlOk() (*string, bool)`

GetIssuerUrlOk returns a tuple with the IssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUrl

`func (o *UserIdentity) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.


### SetIssuerUrlNil

`func (o *UserIdentity) SetIssuerUrlNil(b bool)`

 SetIssuerUrlNil sets the value for IssuerUrl to be an explicit nil

### UnsetIssuerUrl
`func (o *UserIdentity) UnsetIssuerUrl()`

UnsetIssuerUrl ensures that no value is present for IssuerUrl, not even an explicit nil
### GetSubject

`func (o *UserIdentity) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *UserIdentity) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *UserIdentity) SetSubject(v string)`

SetSubject sets Subject field to given value.


### SetSubjectNil

`func (o *UserIdentity) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *UserIdentity) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetUpdatedAt

`func (o *UserIdentity) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserIdentity) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserIdentity) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *UserIdentity) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserIdentity) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserIdentity) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetUserIdentityId

`func (o *UserIdentity) GetUserIdentityId() int64`

GetUserIdentityId returns the UserIdentityId field if non-nil, zero value otherwise.

### GetUserIdentityIdOk

`func (o *UserIdentity) GetUserIdentityIdOk() (*int64, bool)`

GetUserIdentityIdOk returns a tuple with the UserIdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdentityId

`func (o *UserIdentity) SetUserIdentityId(v int64)`

SetUserIdentityId sets UserIdentityId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


