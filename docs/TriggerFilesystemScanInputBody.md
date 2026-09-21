# TriggerFilesystemScanInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to **[]string** | Restrict findings to these categories. Omit to scan the default filesystem-modifying set. | [optional] 
**DirectOnly** | Pointer to **bool** | Only report functions whose own name matches a known filesystem API; skips the calls-into-filesystem pass, avoiding a bulk call-graph fetch. | [optional] 
**Sources** | Pointer to **[]string** | Restrict findings to these sources. Omit to scan every source. | [optional] 

## Methods

### NewTriggerFilesystemScanInputBody

`func NewTriggerFilesystemScanInputBody() *TriggerFilesystemScanInputBody`

NewTriggerFilesystemScanInputBody instantiates a new TriggerFilesystemScanInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerFilesystemScanInputBodyWithDefaults

`func NewTriggerFilesystemScanInputBodyWithDefaults() *TriggerFilesystemScanInputBody`

NewTriggerFilesystemScanInputBodyWithDefaults instantiates a new TriggerFilesystemScanInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *TriggerFilesystemScanInputBody) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *TriggerFilesystemScanInputBody) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *TriggerFilesystemScanInputBody) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *TriggerFilesystemScanInputBody) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategoriesNil

`func (o *TriggerFilesystemScanInputBody) SetCategoriesNil(b bool)`

 SetCategoriesNil sets the value for Categories to be an explicit nil

### UnsetCategories
`func (o *TriggerFilesystemScanInputBody) UnsetCategories()`

UnsetCategories ensures that no value is present for Categories, not even an explicit nil
### GetDirectOnly

`func (o *TriggerFilesystemScanInputBody) GetDirectOnly() bool`

GetDirectOnly returns the DirectOnly field if non-nil, zero value otherwise.

### GetDirectOnlyOk

`func (o *TriggerFilesystemScanInputBody) GetDirectOnlyOk() (*bool, bool)`

GetDirectOnlyOk returns a tuple with the DirectOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectOnly

`func (o *TriggerFilesystemScanInputBody) SetDirectOnly(v bool)`

SetDirectOnly sets DirectOnly field to given value.

### HasDirectOnly

`func (o *TriggerFilesystemScanInputBody) HasDirectOnly() bool`

HasDirectOnly returns a boolean if a field has been set.

### GetSources

`func (o *TriggerFilesystemScanInputBody) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *TriggerFilesystemScanInputBody) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *TriggerFilesystemScanInputBody) SetSources(v []string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *TriggerFilesystemScanInputBody) HasSources() bool`

HasSources returns a boolean if a field has been set.

### SetSourcesNil

`func (o *TriggerFilesystemScanInputBody) SetSourcesNil(b bool)`

 SetSourcesNil sets the value for Sources to be an explicit nil

### UnsetSources
`func (o *TriggerFilesystemScanInputBody) UnsetSources()`

UnsetSources ensures that no value is present for Sources, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


