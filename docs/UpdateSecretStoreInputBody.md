# UpdateSecretStoreInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Activates or deactivates the secret. Omit to leave unchanged. | [optional] 
**Key** | Pointer to **string** | New API key. Omit to leave the current key unchanged. | [optional] 

## Methods

### NewUpdateSecretStoreInputBody

`func NewUpdateSecretStoreInputBody() *UpdateSecretStoreInputBody`

NewUpdateSecretStoreInputBody instantiates a new UpdateSecretStoreInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSecretStoreInputBodyWithDefaults

`func NewUpdateSecretStoreInputBodyWithDefaults() *UpdateSecretStoreInputBody`

NewUpdateSecretStoreInputBodyWithDefaults instantiates a new UpdateSecretStoreInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *UpdateSecretStoreInputBody) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateSecretStoreInputBody) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateSecretStoreInputBody) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateSecretStoreInputBody) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetKey

`func (o *UpdateSecretStoreInputBody) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateSecretStoreInputBody) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateSecretStoreInputBody) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateSecretStoreInputBody) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


