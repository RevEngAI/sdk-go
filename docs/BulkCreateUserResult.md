# BulkCreateUserResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **NullableString** |  | 
**Error** | Pointer to **string** | Error description; present on failure | [optional] 
**Index** | **int64** | 1-based row index in the CSV | 
**Success** | **bool** |  | 
**User** | Pointer to [**User**](User.md) | Created user; present on success | [optional] 
**Username** | **NullableString** |  | 

## Methods

### NewBulkCreateUserResult

`func NewBulkCreateUserResult(email NullableString, index int64, success bool, username NullableString, ) *BulkCreateUserResult`

NewBulkCreateUserResult instantiates a new BulkCreateUserResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkCreateUserResultWithDefaults

`func NewBulkCreateUserResultWithDefaults() *BulkCreateUserResult`

NewBulkCreateUserResultWithDefaults instantiates a new BulkCreateUserResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *BulkCreateUserResult) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BulkCreateUserResult) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BulkCreateUserResult) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *BulkCreateUserResult) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *BulkCreateUserResult) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetError

`func (o *BulkCreateUserResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BulkCreateUserResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BulkCreateUserResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *BulkCreateUserResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIndex

`func (o *BulkCreateUserResult) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *BulkCreateUserResult) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *BulkCreateUserResult) SetIndex(v int64)`

SetIndex sets Index field to given value.


### GetSuccess

`func (o *BulkCreateUserResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *BulkCreateUserResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *BulkCreateUserResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetUser

`func (o *BulkCreateUserResult) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *BulkCreateUserResult) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *BulkCreateUserResult) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *BulkCreateUserResult) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUsername

`func (o *BulkCreateUserResult) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *BulkCreateUserResult) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *BulkCreateUserResult) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *BulkCreateUserResult) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *BulkCreateUserResult) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


