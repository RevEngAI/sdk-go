# UpdatePasswordInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Password reset code received by email | 
**Password** | **string** | New password (minimum 10 characters) | 

## Methods

### NewUpdatePasswordInputBody

`func NewUpdatePasswordInputBody(code string, password string, ) *UpdatePasswordInputBody`

NewUpdatePasswordInputBody instantiates a new UpdatePasswordInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePasswordInputBodyWithDefaults

`func NewUpdatePasswordInputBodyWithDefaults() *UpdatePasswordInputBody`

NewUpdatePasswordInputBodyWithDefaults instantiates a new UpdatePasswordInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *UpdatePasswordInputBody) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdatePasswordInputBody) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdatePasswordInputBody) SetCode(v string)`

SetCode sets Code field to given value.


### GetPassword

`func (o *UpdatePasswordInputBody) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdatePasswordInputBody) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdatePasswordInputBody) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


