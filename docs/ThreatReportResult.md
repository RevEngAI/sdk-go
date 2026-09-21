# ThreatReportResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IOCs** | [**[]IOC**](IOC.md) | Indicators of compromise found. An indicator whose source does not resolve to a function is still listed, without function details. | 
**AttackFlowSummary** | **string** | Markdown summary of the attack flow | 
**ExecutableTechniques** | [**[]Technique**](Technique.md) | MITRE ATT&amp;CK techniques found. A technique is listed only when both its function and its ATT&amp;CK catalogue entry resolve. | 
**NumberOfAnalysedFunctions** | **int64** | Functions the agent analysed | 
**SoftwareType** | **string** | Classification of the binary | 
**Summary** | **string** | Summary of the analysis findings | 
**TotalNumberOfFunctions** | **int64** | Functions identified in the binary | 
**YaraRule** | **string** | YARA rule generated for the binary | 

## Methods

### NewThreatReportResult

`func NewThreatReportResult(iOCs []IOC, attackFlowSummary string, executableTechniques []Technique, numberOfAnalysedFunctions int64, softwareType string, summary string, totalNumberOfFunctions int64, yaraRule string, ) *ThreatReportResult`

NewThreatReportResult instantiates a new ThreatReportResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatReportResultWithDefaults

`func NewThreatReportResultWithDefaults() *ThreatReportResult`

NewThreatReportResultWithDefaults instantiates a new ThreatReportResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIOCs

`func (o *ThreatReportResult) GetIOCs() []IOC`

GetIOCs returns the IOCs field if non-nil, zero value otherwise.

### GetIOCsOk

`func (o *ThreatReportResult) GetIOCsOk() (*[]IOC, bool)`

GetIOCsOk returns a tuple with the IOCs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIOCs

`func (o *ThreatReportResult) SetIOCs(v []IOC)`

SetIOCs sets IOCs field to given value.


### SetIOCsNil

`func (o *ThreatReportResult) SetIOCsNil(b bool)`

 SetIOCsNil sets the value for IOCs to be an explicit nil

### UnsetIOCs
`func (o *ThreatReportResult) UnsetIOCs()`

UnsetIOCs ensures that no value is present for IOCs, not even an explicit nil
### GetAttackFlowSummary

`func (o *ThreatReportResult) GetAttackFlowSummary() string`

GetAttackFlowSummary returns the AttackFlowSummary field if non-nil, zero value otherwise.

### GetAttackFlowSummaryOk

`func (o *ThreatReportResult) GetAttackFlowSummaryOk() (*string, bool)`

GetAttackFlowSummaryOk returns a tuple with the AttackFlowSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackFlowSummary

`func (o *ThreatReportResult) SetAttackFlowSummary(v string)`

SetAttackFlowSummary sets AttackFlowSummary field to given value.


### GetExecutableTechniques

`func (o *ThreatReportResult) GetExecutableTechniques() []Technique`

GetExecutableTechniques returns the ExecutableTechniques field if non-nil, zero value otherwise.

### GetExecutableTechniquesOk

`func (o *ThreatReportResult) GetExecutableTechniquesOk() (*[]Technique, bool)`

GetExecutableTechniquesOk returns a tuple with the ExecutableTechniques field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutableTechniques

`func (o *ThreatReportResult) SetExecutableTechniques(v []Technique)`

SetExecutableTechniques sets ExecutableTechniques field to given value.


### SetExecutableTechniquesNil

`func (o *ThreatReportResult) SetExecutableTechniquesNil(b bool)`

 SetExecutableTechniquesNil sets the value for ExecutableTechniques to be an explicit nil

### UnsetExecutableTechniques
`func (o *ThreatReportResult) UnsetExecutableTechniques()`

UnsetExecutableTechniques ensures that no value is present for ExecutableTechniques, not even an explicit nil
### GetNumberOfAnalysedFunctions

`func (o *ThreatReportResult) GetNumberOfAnalysedFunctions() int64`

GetNumberOfAnalysedFunctions returns the NumberOfAnalysedFunctions field if non-nil, zero value otherwise.

### GetNumberOfAnalysedFunctionsOk

`func (o *ThreatReportResult) GetNumberOfAnalysedFunctionsOk() (*int64, bool)`

GetNumberOfAnalysedFunctionsOk returns a tuple with the NumberOfAnalysedFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfAnalysedFunctions

`func (o *ThreatReportResult) SetNumberOfAnalysedFunctions(v int64)`

SetNumberOfAnalysedFunctions sets NumberOfAnalysedFunctions field to given value.


### GetSoftwareType

`func (o *ThreatReportResult) GetSoftwareType() string`

GetSoftwareType returns the SoftwareType field if non-nil, zero value otherwise.

### GetSoftwareTypeOk

`func (o *ThreatReportResult) GetSoftwareTypeOk() (*string, bool)`

GetSoftwareTypeOk returns a tuple with the SoftwareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareType

`func (o *ThreatReportResult) SetSoftwareType(v string)`

SetSoftwareType sets SoftwareType field to given value.


### GetSummary

`func (o *ThreatReportResult) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ThreatReportResult) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ThreatReportResult) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetTotalNumberOfFunctions

`func (o *ThreatReportResult) GetTotalNumberOfFunctions() int64`

GetTotalNumberOfFunctions returns the TotalNumberOfFunctions field if non-nil, zero value otherwise.

### GetTotalNumberOfFunctionsOk

`func (o *ThreatReportResult) GetTotalNumberOfFunctionsOk() (*int64, bool)`

GetTotalNumberOfFunctionsOk returns a tuple with the TotalNumberOfFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfFunctions

`func (o *ThreatReportResult) SetTotalNumberOfFunctions(v int64)`

SetTotalNumberOfFunctions sets TotalNumberOfFunctions field to given value.


### GetYaraRule

`func (o *ThreatReportResult) GetYaraRule() string`

GetYaraRule returns the YaraRule field if non-nil, zero value otherwise.

### GetYaraRuleOk

`func (o *ThreatReportResult) GetYaraRuleOk() (*string, bool)`

GetYaraRuleOk returns a tuple with the YaraRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaraRule

`func (o *ThreatReportResult) SetYaraRule(v string)`

SetYaraRule sets YaraRule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


