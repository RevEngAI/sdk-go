# GetPublicUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** |  | 
**Username** | **string** |  | 

## Methods

### NewGetPublicUserResponse

`func NewGetPublicUserResponse(userId int32, username string, ) *GetPublicUserResponse`

NewGetPublicUserResponse instantiates a new GetPublicUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicUserResponseWithDefaults

`func NewGetPublicUserResponseWithDefaults() *GetPublicUserResponse`

NewGetPublicUserResponseWithDefaults instantiates a new GetPublicUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *GetPublicUserResponse) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetPublicUserResponse) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetPublicUserResponse) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetUsername

`func (o *GetPublicUserResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetPublicUserResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetPublicUserResponse) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


