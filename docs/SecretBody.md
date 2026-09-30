# SecretBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** |  | 
**ApiProvider** | **string** |  | 
**Creation** | **time.Time** |  | 
**DisabledAt** | Pointer to **time.Time** |  | [optional] 
**Id** | **int64** |  | 
**Key** | **string** | Masked API key showing only the last 4 characters | 
**TeamId** | Pointer to **int64** | Null for a personal secret | [optional] 
**UserId** | **int64** |  | 
**Valid** | **bool** |  | 

## Methods

### NewSecretBody

`func NewSecretBody(active bool, apiProvider string, creation time.Time, id int64, key string, userId int64, valid bool, ) *SecretBody`

NewSecretBody instantiates a new SecretBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretBodyWithDefaults

`func NewSecretBodyWithDefaults() *SecretBody`

NewSecretBodyWithDefaults instantiates a new SecretBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *SecretBody) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SecretBody) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SecretBody) SetActive(v bool)`

SetActive sets Active field to given value.


### GetApiProvider

`func (o *SecretBody) GetApiProvider() string`

GetApiProvider returns the ApiProvider field if non-nil, zero value otherwise.

### GetApiProviderOk

`func (o *SecretBody) GetApiProviderOk() (*string, bool)`

GetApiProviderOk returns a tuple with the ApiProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiProvider

`func (o *SecretBody) SetApiProvider(v string)`

SetApiProvider sets ApiProvider field to given value.


### GetCreation

`func (o *SecretBody) GetCreation() time.Time`

GetCreation returns the Creation field if non-nil, zero value otherwise.

### GetCreationOk

`func (o *SecretBody) GetCreationOk() (*time.Time, bool)`

GetCreationOk returns a tuple with the Creation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreation

`func (o *SecretBody) SetCreation(v time.Time)`

SetCreation sets Creation field to given value.


### GetDisabledAt

`func (o *SecretBody) GetDisabledAt() time.Time`

GetDisabledAt returns the DisabledAt field if non-nil, zero value otherwise.

### GetDisabledAtOk

`func (o *SecretBody) GetDisabledAtOk() (*time.Time, bool)`

GetDisabledAtOk returns a tuple with the DisabledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAt

`func (o *SecretBody) SetDisabledAt(v time.Time)`

SetDisabledAt sets DisabledAt field to given value.

### HasDisabledAt

`func (o *SecretBody) HasDisabledAt() bool`

HasDisabledAt returns a boolean if a field has been set.

### GetId

`func (o *SecretBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecretBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecretBody) SetId(v int64)`

SetId sets Id field to given value.


### GetKey

`func (o *SecretBody) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SecretBody) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SecretBody) SetKey(v string)`

SetKey sets Key field to given value.


### GetTeamId

`func (o *SecretBody) GetTeamId() int64`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *SecretBody) GetTeamIdOk() (*int64, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *SecretBody) SetTeamId(v int64)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *SecretBody) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetUserId

`func (o *SecretBody) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SecretBody) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SecretBody) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetValid

`func (o *SecretBody) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *SecretBody) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *SecretBody) SetValid(v bool)`

SetValid sets Valid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


