# SecretsAgentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 
**Findings** | Pointer to [**[]Finding**](Finding.md) |  | [optional] 

## Methods

### NewSecretsAgentResponse

`func NewSecretsAgentResponse() *SecretsAgentResponse`

NewSecretsAgentResponse instantiates a new SecretsAgentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretsAgentResponseWithDefaults

`func NewSecretsAgentResponseWithDefaults() *SecretsAgentResponse`

NewSecretsAgentResponseWithDefaults instantiates a new SecretsAgentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *SecretsAgentResponse) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SecretsAgentResponse) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SecretsAgentResponse) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SecretsAgentResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetFindings

`func (o *SecretsAgentResponse) GetFindings() []Finding`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *SecretsAgentResponse) GetFindingsOk() (*[]Finding, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *SecretsAgentResponse) SetFindings(v []Finding)`

SetFindings sets Findings field to given value.

### HasFindings

`func (o *SecretsAgentResponse) HasFindings() bool`

HasFindings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


