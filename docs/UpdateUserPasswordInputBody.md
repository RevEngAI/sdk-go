# UpdateUserPasswordInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPassword** | Pointer to **string** | Current password (required when changing your own password; ignored for SUPERADMIN updates of other users) | [optional] 
**NewPassword** | **string** | New password (minimum 10 characters) | 

## Methods

### NewUpdateUserPasswordInputBody

`func NewUpdateUserPasswordInputBody(newPassword string, ) *UpdateUserPasswordInputBody`

NewUpdateUserPasswordInputBody instantiates a new UpdateUserPasswordInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserPasswordInputBodyWithDefaults

`func NewUpdateUserPasswordInputBodyWithDefaults() *UpdateUserPasswordInputBody`

NewUpdateUserPasswordInputBodyWithDefaults instantiates a new UpdateUserPasswordInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPassword

`func (o *UpdateUserPasswordInputBody) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *UpdateUserPasswordInputBody) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *UpdateUserPasswordInputBody) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.

### HasCurrentPassword

`func (o *UpdateUserPasswordInputBody) HasCurrentPassword() bool`

HasCurrentPassword returns a boolean if a field has been set.

### GetNewPassword

`func (o *UpdateUserPasswordInputBody) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *UpdateUserPasswordInputBody) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *UpdateUserPasswordInputBody) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


