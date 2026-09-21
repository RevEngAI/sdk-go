# TriggerExecutionScanInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to **[]string** | Restrict findings to these categories. Omit to scan the default executes-code set. | [optional] 
**DirectOnly** | Pointer to **bool** | Only report functions whose own name matches a known execution API; skips the calls-into-execution pass, avoiding a bulk call-graph fetch. | [optional] 
**Sources** | Pointer to **[]string** | Restrict findings to these sources. Omit to scan every source. | [optional] 

## Methods

### NewTriggerExecutionScanInputBody

`func NewTriggerExecutionScanInputBody() *TriggerExecutionScanInputBody`

NewTriggerExecutionScanInputBody instantiates a new TriggerExecutionScanInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerExecutionScanInputBodyWithDefaults

`func NewTriggerExecutionScanInputBodyWithDefaults() *TriggerExecutionScanInputBody`

NewTriggerExecutionScanInputBodyWithDefaults instantiates a new TriggerExecutionScanInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *TriggerExecutionScanInputBody) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *TriggerExecutionScanInputBody) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *TriggerExecutionScanInputBody) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *TriggerExecutionScanInputBody) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategoriesNil

`func (o *TriggerExecutionScanInputBody) SetCategoriesNil(b bool)`

 SetCategoriesNil sets the value for Categories to be an explicit nil

### UnsetCategories
`func (o *TriggerExecutionScanInputBody) UnsetCategories()`

UnsetCategories ensures that no value is present for Categories, not even an explicit nil
### GetDirectOnly

`func (o *TriggerExecutionScanInputBody) GetDirectOnly() bool`

GetDirectOnly returns the DirectOnly field if non-nil, zero value otherwise.

### GetDirectOnlyOk

`func (o *TriggerExecutionScanInputBody) GetDirectOnlyOk() (*bool, bool)`

GetDirectOnlyOk returns a tuple with the DirectOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectOnly

`func (o *TriggerExecutionScanInputBody) SetDirectOnly(v bool)`

SetDirectOnly sets DirectOnly field to given value.

### HasDirectOnly

`func (o *TriggerExecutionScanInputBody) HasDirectOnly() bool`

HasDirectOnly returns a boolean if a field has been set.

### GetSources

`func (o *TriggerExecutionScanInputBody) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *TriggerExecutionScanInputBody) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *TriggerExecutionScanInputBody) SetSources(v []string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *TriggerExecutionScanInputBody) HasSources() bool`

HasSources returns a boolean if a field has been set.

### SetSourcesNil

`func (o *TriggerExecutionScanInputBody) SetSourcesNil(b bool)`

 SetSourcesNil sets the value for Sources to be an explicit nil

### UnsetSources
`func (o *TriggerExecutionScanInputBody) UnsetSources()`

UnsetSources ensures that no value is present for Sources, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


