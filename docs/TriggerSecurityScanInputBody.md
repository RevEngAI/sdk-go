# TriggerSecurityScanInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxFunctionsToScan** | Pointer to **int64** | Stop after decompiling and scanning this many functions. Omit to process every function in the analysis. | [optional] 

## Methods

### NewTriggerSecurityScanInputBody

`func NewTriggerSecurityScanInputBody() *TriggerSecurityScanInputBody`

NewTriggerSecurityScanInputBody instantiates a new TriggerSecurityScanInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerSecurityScanInputBodyWithDefaults

`func NewTriggerSecurityScanInputBodyWithDefaults() *TriggerSecurityScanInputBody`

NewTriggerSecurityScanInputBodyWithDefaults instantiates a new TriggerSecurityScanInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxFunctionsToScan

`func (o *TriggerSecurityScanInputBody) GetMaxFunctionsToScan() int64`

GetMaxFunctionsToScan returns the MaxFunctionsToScan field if non-nil, zero value otherwise.

### GetMaxFunctionsToScanOk

`func (o *TriggerSecurityScanInputBody) GetMaxFunctionsToScanOk() (*int64, bool)`

GetMaxFunctionsToScanOk returns a tuple with the MaxFunctionsToScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFunctionsToScan

`func (o *TriggerSecurityScanInputBody) SetMaxFunctionsToScan(v int64)`

SetMaxFunctionsToScan sets MaxFunctionsToScan field to given value.

### HasMaxFunctionsToScan

`func (o *TriggerSecurityScanInputBody) HasMaxFunctionsToScan() bool`

HasMaxFunctionsToScan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


