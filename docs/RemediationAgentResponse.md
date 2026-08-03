# RemediationAgentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**YaraRules** | **[]string** | Generated YARA rules for the binary | 
**SnortRules** | **[]string** | Generated Snort rules for the binary | 
**StixRules** | **[]string** | Generated STIX rules for the binary | 

## Methods

### NewRemediationAgentResponse

`func NewRemediationAgentResponse(yaraRules []string, snortRules []string, stixRules []string, ) *RemediationAgentResponse`

NewRemediationAgentResponse instantiates a new RemediationAgentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemediationAgentResponseWithDefaults

`func NewRemediationAgentResponseWithDefaults() *RemediationAgentResponse`

NewRemediationAgentResponseWithDefaults instantiates a new RemediationAgentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetYaraRules

`func (o *RemediationAgentResponse) GetYaraRules() []string`

GetYaraRules returns the YaraRules field if non-nil, zero value otherwise.

### GetYaraRulesOk

`func (o *RemediationAgentResponse) GetYaraRulesOk() (*[]string, bool)`

GetYaraRulesOk returns a tuple with the YaraRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaraRules

`func (o *RemediationAgentResponse) SetYaraRules(v []string)`

SetYaraRules sets YaraRules field to given value.


### GetSnortRules

`func (o *RemediationAgentResponse) GetSnortRules() []string`

GetSnortRules returns the SnortRules field if non-nil, zero value otherwise.

### GetSnortRulesOk

`func (o *RemediationAgentResponse) GetSnortRulesOk() (*[]string, bool)`

GetSnortRulesOk returns a tuple with the SnortRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnortRules

`func (o *RemediationAgentResponse) SetSnortRules(v []string)`

SetSnortRules sets SnortRules field to given value.


### GetStixRules

`func (o *RemediationAgentResponse) GetStixRules() []string`

GetStixRules returns the StixRules field if non-nil, zero value otherwise.

### GetStixRulesOk

`func (o *RemediationAgentResponse) GetStixRulesOk() (*[]string, bool)`

GetStixRulesOk returns a tuple with the StixRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStixRules

`func (o *RemediationAgentResponse) SetStixRules(v []string)`

SetStixRules sets StixRules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


