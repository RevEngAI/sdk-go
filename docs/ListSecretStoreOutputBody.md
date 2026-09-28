# ListSecretStoreOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secrets** | [**[]SecretBody**](SecretBody.md) |  | 

## Methods

### NewListSecretStoreOutputBody

`func NewListSecretStoreOutputBody(secrets []SecretBody, ) *ListSecretStoreOutputBody`

NewListSecretStoreOutputBody instantiates a new ListSecretStoreOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSecretStoreOutputBodyWithDefaults

`func NewListSecretStoreOutputBodyWithDefaults() *ListSecretStoreOutputBody`

NewListSecretStoreOutputBodyWithDefaults instantiates a new ListSecretStoreOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecrets

`func (o *ListSecretStoreOutputBody) GetSecrets() []SecretBody`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *ListSecretStoreOutputBody) GetSecretsOk() (*[]SecretBody, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *ListSecretStoreOutputBody) SetSecrets(v []SecretBody)`

SetSecrets sets Secrets field to given value.


### SetSecretsNil

`func (o *ListSecretStoreOutputBody) SetSecretsNil(b bool)`

 SetSecretsNil sets the value for Secrets to be an explicit nil

### UnsetSecrets
`func (o *ListSecretStoreOutputBody) UnsetSecrets()`

UnsetSecrets ensures that no value is present for Secrets, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


