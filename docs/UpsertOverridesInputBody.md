# UpsertOverridesInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**Overrides** | **map[string]string** | Token to name mappings. Empty string removes the override. | 

## Methods

### NewUpsertOverridesInputBody

`func NewUpsertOverridesInputBody(overrides map[string]string, ) *UpsertOverridesInputBody`

NewUpsertOverridesInputBody instantiates a new UpsertOverridesInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertOverridesInputBodyWithDefaults

`func NewUpsertOverridesInputBodyWithDefaults() *UpsertOverridesInputBody`

NewUpsertOverridesInputBodyWithDefaults instantiates a new UpsertOverridesInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *UpsertOverridesInputBody) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *UpsertOverridesInputBody) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *UpsertOverridesInputBody) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *UpsertOverridesInputBody) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetOverrides

`func (o *UpsertOverridesInputBody) GetOverrides() map[string]string`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *UpsertOverridesInputBody) GetOverridesOk() (*map[string]string, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *UpsertOverridesInputBody) SetOverrides(v map[string]string)`

SetOverrides sets Overrides field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


