# ReportAnalysisResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | **string** | A markdown summary of the report | 
**SoftwareType** | **string** | The type of software being analyzed | 
**TotalNumberOfFunctions** | **int32** | The total number of functions identified in the binary | 
**NumberOfAnalysedFunctions** | **int32** | The number of functions that were analyzed in the binary | 
**AttackFlowSummary** | **string** | A summary in markdown format of the attack flow | 
**IOCs** | [**[]IOC**](IOC.md) | A list of IOCs (Indicators of Compromise) found in the analysis | 
**ExecutableTechniques** | [**[]MITRETechnique**](MITRETechnique.md) | A series of MITRE Techniques found | 
**YaraRule** | **string** | The YARA rule generated for the binary | 

## Methods

### NewReportAnalysisResponse

`func NewReportAnalysisResponse(summary string, softwareType string, totalNumberOfFunctions int32, numberOfAnalysedFunctions int32, attackFlowSummary string, iOCs []IOC, executableTechniques []MITRETechnique, yaraRule string, ) *ReportAnalysisResponse`

NewReportAnalysisResponse instantiates a new ReportAnalysisResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportAnalysisResponseWithDefaults

`func NewReportAnalysisResponseWithDefaults() *ReportAnalysisResponse`

NewReportAnalysisResponseWithDefaults instantiates a new ReportAnalysisResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *ReportAnalysisResponse) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ReportAnalysisResponse) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ReportAnalysisResponse) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetSoftwareType

`func (o *ReportAnalysisResponse) GetSoftwareType() string`

GetSoftwareType returns the SoftwareType field if non-nil, zero value otherwise.

### GetSoftwareTypeOk

`func (o *ReportAnalysisResponse) GetSoftwareTypeOk() (*string, bool)`

GetSoftwareTypeOk returns a tuple with the SoftwareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareType

`func (o *ReportAnalysisResponse) SetSoftwareType(v string)`

SetSoftwareType sets SoftwareType field to given value.


### GetTotalNumberOfFunctions

`func (o *ReportAnalysisResponse) GetTotalNumberOfFunctions() int32`

GetTotalNumberOfFunctions returns the TotalNumberOfFunctions field if non-nil, zero value otherwise.

### GetTotalNumberOfFunctionsOk

`func (o *ReportAnalysisResponse) GetTotalNumberOfFunctionsOk() (*int32, bool)`

GetTotalNumberOfFunctionsOk returns a tuple with the TotalNumberOfFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfFunctions

`func (o *ReportAnalysisResponse) SetTotalNumberOfFunctions(v int32)`

SetTotalNumberOfFunctions sets TotalNumberOfFunctions field to given value.


### GetNumberOfAnalysedFunctions

`func (o *ReportAnalysisResponse) GetNumberOfAnalysedFunctions() int32`

GetNumberOfAnalysedFunctions returns the NumberOfAnalysedFunctions field if non-nil, zero value otherwise.

### GetNumberOfAnalysedFunctionsOk

`func (o *ReportAnalysisResponse) GetNumberOfAnalysedFunctionsOk() (*int32, bool)`

GetNumberOfAnalysedFunctionsOk returns a tuple with the NumberOfAnalysedFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfAnalysedFunctions

`func (o *ReportAnalysisResponse) SetNumberOfAnalysedFunctions(v int32)`

SetNumberOfAnalysedFunctions sets NumberOfAnalysedFunctions field to given value.


### GetAttackFlowSummary

`func (o *ReportAnalysisResponse) GetAttackFlowSummary() string`

GetAttackFlowSummary returns the AttackFlowSummary field if non-nil, zero value otherwise.

### GetAttackFlowSummaryOk

`func (o *ReportAnalysisResponse) GetAttackFlowSummaryOk() (*string, bool)`

GetAttackFlowSummaryOk returns a tuple with the AttackFlowSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackFlowSummary

`func (o *ReportAnalysisResponse) SetAttackFlowSummary(v string)`

SetAttackFlowSummary sets AttackFlowSummary field to given value.


### GetIOCs

`func (o *ReportAnalysisResponse) GetIOCs() []IOC`

GetIOCs returns the IOCs field if non-nil, zero value otherwise.

### GetIOCsOk

`func (o *ReportAnalysisResponse) GetIOCsOk() (*[]IOC, bool)`

GetIOCsOk returns a tuple with the IOCs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIOCs

`func (o *ReportAnalysisResponse) SetIOCs(v []IOC)`

SetIOCs sets IOCs field to given value.


### GetExecutableTechniques

`func (o *ReportAnalysisResponse) GetExecutableTechniques() []MITRETechnique`

GetExecutableTechniques returns the ExecutableTechniques field if non-nil, zero value otherwise.

### GetExecutableTechniquesOk

`func (o *ReportAnalysisResponse) GetExecutableTechniquesOk() (*[]MITRETechnique, bool)`

GetExecutableTechniquesOk returns a tuple with the ExecutableTechniques field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutableTechniques

`func (o *ReportAnalysisResponse) SetExecutableTechniques(v []MITRETechnique)`

SetExecutableTechniques sets ExecutableTechniques field to given value.


### GetYaraRule

`func (o *ReportAnalysisResponse) GetYaraRule() string`

GetYaraRule returns the YaraRule field if non-nil, zero value otherwise.

### GetYaraRuleOk

`func (o *ReportAnalysisResponse) GetYaraRuleOk() (*string, bool)`

GetYaraRuleOk returns a tuple with the YaraRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaraRule

`func (o *ReportAnalysisResponse) SetYaraRule(v string)`

SetYaraRule sets YaraRule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


