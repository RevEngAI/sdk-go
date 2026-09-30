# GetPublicUserOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int64** | The user&#39;s ID | 
**Username** | **string** | The user&#39;s display name | 

## Methods

### NewGetPublicUserOutputBody

`func NewGetPublicUserOutputBody(userId int64, username string, ) *GetPublicUserOutputBody`

NewGetPublicUserOutputBody instantiates a new GetPublicUserOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicUserOutputBodyWithDefaults

`func NewGetPublicUserOutputBodyWithDefaults() *GetPublicUserOutputBody`

NewGetPublicUserOutputBodyWithDefaults instantiates a new GetPublicUserOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *GetPublicUserOutputBody) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetPublicUserOutputBody) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetPublicUserOutputBody) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetUsername

`func (o *GetPublicUserOutputBody) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetPublicUserOutputBody) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetPublicUserOutputBody) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


