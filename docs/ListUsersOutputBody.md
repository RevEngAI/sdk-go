# ListUsersOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** |  | 
**Users** | [**[]User**](User.md) |  | 

## Methods

### NewListUsersOutputBody

`func NewListUsersOutputBody(total int64, users []User, ) *ListUsersOutputBody`

NewListUsersOutputBody instantiates a new ListUsersOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUsersOutputBodyWithDefaults

`func NewListUsersOutputBodyWithDefaults() *ListUsersOutputBody`

NewListUsersOutputBodyWithDefaults instantiates a new ListUsersOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ListUsersOutputBody) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListUsersOutputBody) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListUsersOutputBody) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetUsers

`func (o *ListUsersOutputBody) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ListUsersOutputBody) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ListUsersOutputBody) SetUsers(v []User)`

SetUsers sets Users field to given value.


### SetUsersNil

`func (o *ListUsersOutputBody) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *ListUsersOutputBody) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


