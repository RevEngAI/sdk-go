# CreateSecretStoreInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiProvider** | **string** |  | 
**Key** | **string** | The provider&#39;s API key, encrypted before storage. | 
**TeamId** | Pointer to **int64** | Registers a team secret when set; the caller must administer that team. Registers a personal secret when omitted. | [optional] 

## Methods

### NewCreateSecretStoreInputBody

`func NewCreateSecretStoreInputBody(apiProvider string, key string, ) *CreateSecretStoreInputBody`

NewCreateSecretStoreInputBody instantiates a new CreateSecretStoreInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSecretStoreInputBodyWithDefaults

`func NewCreateSecretStoreInputBodyWithDefaults() *CreateSecretStoreInputBody`

NewCreateSecretStoreInputBodyWithDefaults instantiates a new CreateSecretStoreInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiProvider

`func (o *CreateSecretStoreInputBody) GetApiProvider() string`

GetApiProvider returns the ApiProvider field if non-nil, zero value otherwise.

### GetApiProviderOk

`func (o *CreateSecretStoreInputBody) GetApiProviderOk() (*string, bool)`

GetApiProviderOk returns a tuple with the ApiProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiProvider

`func (o *CreateSecretStoreInputBody) SetApiProvider(v string)`

SetApiProvider sets ApiProvider field to given value.


### GetKey

`func (o *CreateSecretStoreInputBody) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateSecretStoreInputBody) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateSecretStoreInputBody) SetKey(v string)`

SetKey sets Key field to given value.


### GetTeamId

`func (o *CreateSecretStoreInputBody) GetTeamId() int64`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *CreateSecretStoreInputBody) GetTeamIdOk() (*int64, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *CreateSecretStoreInputBody) SetTeamId(v int64)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *CreateSecretStoreInputBody) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


