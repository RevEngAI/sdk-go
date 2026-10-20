# PublicUserBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int64** | The user&#39;s ID | 
**Username** | **string** | The user&#39;s display name | 

## Methods

### NewPublicUserBody

`func NewPublicUserBody(userId int64, username string, ) *PublicUserBody`

NewPublicUserBody instantiates a new PublicUserBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicUserBodyWithDefaults

`func NewPublicUserBodyWithDefaults() *PublicUserBody`

NewPublicUserBodyWithDefaults instantiates a new PublicUserBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *PublicUserBody) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PublicUserBody) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PublicUserBody) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetUsername

`func (o *PublicUserBody) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PublicUserBody) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PublicUserBody) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


