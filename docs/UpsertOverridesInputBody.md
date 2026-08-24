# UpsertOverridesInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overrides** | [**map[string]Token**](Token.md) | Overrides keyed by placeholder token. An entry whose value is an empty string removes that override. | 

## Methods

### NewUpsertOverridesInputBody

`func NewUpsertOverridesInputBody(overrides map[string]Token, ) *UpsertOverridesInputBody`

NewUpsertOverridesInputBody instantiates a new UpsertOverridesInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertOverridesInputBodyWithDefaults

`func NewUpsertOverridesInputBodyWithDefaults() *UpsertOverridesInputBody`

NewUpsertOverridesInputBodyWithDefaults instantiates a new UpsertOverridesInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverrides

`func (o *UpsertOverridesInputBody) GetOverrides() map[string]Token`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *UpsertOverridesInputBody) GetOverridesOk() (*map[string]Token, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *UpsertOverridesInputBody) SetOverrides(v map[string]Token)`

SetOverrides sets Overrides field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


