# TokenInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** | API key to exchange for a JWT | [optional] 
**Email** | Pointer to **string** | User email address (required with password) | [optional] 
**Password** | Pointer to **string** | User password (required with email or username) | [optional] 
**Username** | Pointer to **string** | Username (alternative to email, required with password) | [optional] 

## Methods

### NewTokenInputBody

`func NewTokenInputBody() *TokenInputBody`

NewTokenInputBody instantiates a new TokenInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenInputBodyWithDefaults

`func NewTokenInputBodyWithDefaults() *TokenInputBody`

NewTokenInputBodyWithDefaults instantiates a new TokenInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *TokenInputBody) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *TokenInputBody) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *TokenInputBody) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *TokenInputBody) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetEmail

`func (o *TokenInputBody) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TokenInputBody) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TokenInputBody) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *TokenInputBody) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPassword

`func (o *TokenInputBody) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TokenInputBody) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TokenInputBody) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TokenInputBody) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *TokenInputBody) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TokenInputBody) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TokenInputBody) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TokenInputBody) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


