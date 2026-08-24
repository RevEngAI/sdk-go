# HistoryActor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int64** | The user who made the change. | 
**Username** | Pointer to **string** | Absent when the user no longer exists. | [optional] 

## Methods

### NewHistoryActor

`func NewHistoryActor(userId int64, ) *HistoryActor`

NewHistoryActor instantiates a new HistoryActor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryActorWithDefaults

`func NewHistoryActorWithDefaults() *HistoryActor`

NewHistoryActorWithDefaults instantiates a new HistoryActor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *HistoryActor) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *HistoryActor) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *HistoryActor) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetUsername

`func (o *HistoryActor) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *HistoryActor) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *HistoryActor) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *HistoryActor) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


