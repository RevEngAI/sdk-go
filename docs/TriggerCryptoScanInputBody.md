# TriggerCryptoScanInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to **[]string** | Restrict findings to these categories. Omit to scan every category. | [optional] 
**DirectOnly** | Pointer to **bool** | Only report functions whose own name matches a known crypto API; skips the calls-into-crypto pass, avoiding a bulk call-graph fetch. | [optional] 
**Libraries** | Pointer to **[]string** | Restrict findings to these libraries. Omit to scan every library. | [optional] 

## Methods

### NewTriggerCryptoScanInputBody

`func NewTriggerCryptoScanInputBody() *TriggerCryptoScanInputBody`

NewTriggerCryptoScanInputBody instantiates a new TriggerCryptoScanInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerCryptoScanInputBodyWithDefaults

`func NewTriggerCryptoScanInputBodyWithDefaults() *TriggerCryptoScanInputBody`

NewTriggerCryptoScanInputBodyWithDefaults instantiates a new TriggerCryptoScanInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *TriggerCryptoScanInputBody) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *TriggerCryptoScanInputBody) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *TriggerCryptoScanInputBody) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *TriggerCryptoScanInputBody) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategoriesNil

`func (o *TriggerCryptoScanInputBody) SetCategoriesNil(b bool)`

 SetCategoriesNil sets the value for Categories to be an explicit nil

### UnsetCategories
`func (o *TriggerCryptoScanInputBody) UnsetCategories()`

UnsetCategories ensures that no value is present for Categories, not even an explicit nil
### GetDirectOnly

`func (o *TriggerCryptoScanInputBody) GetDirectOnly() bool`

GetDirectOnly returns the DirectOnly field if non-nil, zero value otherwise.

### GetDirectOnlyOk

`func (o *TriggerCryptoScanInputBody) GetDirectOnlyOk() (*bool, bool)`

GetDirectOnlyOk returns a tuple with the DirectOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectOnly

`func (o *TriggerCryptoScanInputBody) SetDirectOnly(v bool)`

SetDirectOnly sets DirectOnly field to given value.

### HasDirectOnly

`func (o *TriggerCryptoScanInputBody) HasDirectOnly() bool`

HasDirectOnly returns a boolean if a field has been set.

### GetLibraries

`func (o *TriggerCryptoScanInputBody) GetLibraries() []string`

GetLibraries returns the Libraries field if non-nil, zero value otherwise.

### GetLibrariesOk

`func (o *TriggerCryptoScanInputBody) GetLibrariesOk() (*[]string, bool)`

GetLibrariesOk returns a tuple with the Libraries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraries

`func (o *TriggerCryptoScanInputBody) SetLibraries(v []string)`

SetLibraries sets Libraries field to given value.

### HasLibraries

`func (o *TriggerCryptoScanInputBody) HasLibraries() bool`

HasLibraries returns a boolean if a field has been set.

### SetLibrariesNil

`func (o *TriggerCryptoScanInputBody) SetLibrariesNil(b bool)`

 SetLibrariesNil sets the value for Libraries to be an explicit nil

### UnsetLibraries
`func (o *TriggerCryptoScanInputBody) UnsetLibraries()`

UnsetLibraries ensures that no value is present for Libraries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


