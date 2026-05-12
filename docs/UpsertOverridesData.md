# UpsertOverridesData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**UserOverrideMappings** | **map[string]string** | Merged override mappings after applying changes | 

## Methods

### NewUpsertOverridesData

`func NewUpsertOverridesData(userOverrideMappings map[string]string, ) *UpsertOverridesData`

NewUpsertOverridesData instantiates a new UpsertOverridesData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertOverridesDataWithDefaults

`func NewUpsertOverridesDataWithDefaults() *UpsertOverridesData`

NewUpsertOverridesDataWithDefaults instantiates a new UpsertOverridesData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *UpsertOverridesData) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *UpsertOverridesData) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *UpsertOverridesData) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *UpsertOverridesData) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetUserOverrideMappings

`func (o *UpsertOverridesData) GetUserOverrideMappings() map[string]string`

GetUserOverrideMappings returns the UserOverrideMappings field if non-nil, zero value otherwise.

### GetUserOverrideMappingsOk

`func (o *UpsertOverridesData) GetUserOverrideMappingsOk() (*map[string]string, bool)`

GetUserOverrideMappingsOk returns a tuple with the UserOverrideMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserOverrideMappings

`func (o *UpsertOverridesData) SetUserOverrideMappings(v map[string]string)`

SetUserOverrideMappings sets UserOverrideMappings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


