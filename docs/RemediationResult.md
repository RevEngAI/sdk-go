# RemediationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnortRules** | **[]string** | Generated Snort rules | 
**StixRules** | **[]string** | Generated STIX rules | 
**YaraRules** | **[]string** | Generated YARA rules | 

## Methods

### NewRemediationResult

`func NewRemediationResult(snortRules []string, stixRules []string, yaraRules []string, ) *RemediationResult`

NewRemediationResult instantiates a new RemediationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemediationResultWithDefaults

`func NewRemediationResultWithDefaults() *RemediationResult`

NewRemediationResultWithDefaults instantiates a new RemediationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnortRules

`func (o *RemediationResult) GetSnortRules() []string`

GetSnortRules returns the SnortRules field if non-nil, zero value otherwise.

### GetSnortRulesOk

`func (o *RemediationResult) GetSnortRulesOk() (*[]string, bool)`

GetSnortRulesOk returns a tuple with the SnortRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnortRules

`func (o *RemediationResult) SetSnortRules(v []string)`

SetSnortRules sets SnortRules field to given value.


### SetSnortRulesNil

`func (o *RemediationResult) SetSnortRulesNil(b bool)`

 SetSnortRulesNil sets the value for SnortRules to be an explicit nil

### UnsetSnortRules
`func (o *RemediationResult) UnsetSnortRules()`

UnsetSnortRules ensures that no value is present for SnortRules, not even an explicit nil
### GetStixRules

`func (o *RemediationResult) GetStixRules() []string`

GetStixRules returns the StixRules field if non-nil, zero value otherwise.

### GetStixRulesOk

`func (o *RemediationResult) GetStixRulesOk() (*[]string, bool)`

GetStixRulesOk returns a tuple with the StixRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStixRules

`func (o *RemediationResult) SetStixRules(v []string)`

SetStixRules sets StixRules field to given value.


### SetStixRulesNil

`func (o *RemediationResult) SetStixRulesNil(b bool)`

 SetStixRulesNil sets the value for StixRules to be an explicit nil

### UnsetStixRules
`func (o *RemediationResult) UnsetStixRules()`

UnsetStixRules ensures that no value is present for StixRules, not even an explicit nil
### GetYaraRules

`func (o *RemediationResult) GetYaraRules() []string`

GetYaraRules returns the YaraRules field if non-nil, zero value otherwise.

### GetYaraRulesOk

`func (o *RemediationResult) GetYaraRulesOk() (*[]string, bool)`

GetYaraRulesOk returns a tuple with the YaraRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaraRules

`func (o *RemediationResult) SetYaraRules(v []string)`

SetYaraRules sets YaraRules field to given value.


### SetYaraRulesNil

`func (o *RemediationResult) SetYaraRulesNil(b bool)`

 SetYaraRulesNil sets the value for YaraRules to be an explicit nil

### UnsetYaraRules
`func (o *RemediationResult) UnsetYaraRules()`

UnsetYaraRules ensures that no value is present for YaraRules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


