# WorkflowDayBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyseCapabilities** | [**AgentCompletedBody**](AgentCompletedBody.md) | ANALYSE_CAPABILITIES agent-task stats for the day | 
**CryptoScan** | [**AgentCompletedBody**](AgentCompletedBody.md) | CRYPTOGRAPHY agent-task stats for the day | 
**CryptographyBinary** | [**AgentCompletedBody**](AgentCompletedBody.md) | CRYPTOGRAPHY_BINARY agent-task stats for the day | 
**Date** | **string** | Day this entry covers | 
**ExecutionScan** | [**AgentCompletedBody**](AgentCompletedBody.md) | EXECUTION agent-task stats for the day | 
**FilesystemScan** | [**AgentCompletedBody**](AgentCompletedBody.md) | FILESYSTEM agent-task stats for the day | 
**NetworkingScan** | [**AgentCompletedBody**](AgentCompletedBody.md) | NETWORKING agent-task stats for the day | 
**Protocols** | [**AgentCompletedBody**](AgentCompletedBody.md) | PROTOCOLS agent-task stats for the day | 
**Remediation** | [**AgentCompletedBody**](AgentCompletedBody.md) | REMEDIATION agent-task stats for the day | 
**RenameUnnamedFunctions** | [**AgentCompletedBody**](AgentCompletedBody.md) | RENAME_UNNAMED_FUNCTIONS agent-task stats for the day | 
**ReportAnalysis** | [**ReportAnalysisBody**](ReportAnalysisBody.md) | REPORT_ANALYSIS agent-task stats for the day | 
**Secrets** | [**AgentCompletedBody**](AgentCompletedBody.md) | SECRETS agent-task stats for the day | 
**SecurityScan** | [**AgentCompletedBody**](AgentCompletedBody.md) | SECURITY_SCAN agent-task stats for the day | 
**Triage** | [**AgentCompletedBody**](AgentCompletedBody.md) | TRIAGE agent-task stats for the day | 

## Methods

### NewWorkflowDayBody

`func NewWorkflowDayBody(analyseCapabilities AgentCompletedBody, cryptoScan AgentCompletedBody, cryptographyBinary AgentCompletedBody, date string, executionScan AgentCompletedBody, filesystemScan AgentCompletedBody, networkingScan AgentCompletedBody, protocols AgentCompletedBody, remediation AgentCompletedBody, renameUnnamedFunctions AgentCompletedBody, reportAnalysis ReportAnalysisBody, secrets AgentCompletedBody, securityScan AgentCompletedBody, triage AgentCompletedBody, ) *WorkflowDayBody`

NewWorkflowDayBody instantiates a new WorkflowDayBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowDayBodyWithDefaults

`func NewWorkflowDayBodyWithDefaults() *WorkflowDayBody`

NewWorkflowDayBodyWithDefaults instantiates a new WorkflowDayBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyseCapabilities

`func (o *WorkflowDayBody) GetAnalyseCapabilities() AgentCompletedBody`

GetAnalyseCapabilities returns the AnalyseCapabilities field if non-nil, zero value otherwise.

### GetAnalyseCapabilitiesOk

`func (o *WorkflowDayBody) GetAnalyseCapabilitiesOk() (*AgentCompletedBody, bool)`

GetAnalyseCapabilitiesOk returns a tuple with the AnalyseCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyseCapabilities

`func (o *WorkflowDayBody) SetAnalyseCapabilities(v AgentCompletedBody)`

SetAnalyseCapabilities sets AnalyseCapabilities field to given value.


### GetCryptoScan

`func (o *WorkflowDayBody) GetCryptoScan() AgentCompletedBody`

GetCryptoScan returns the CryptoScan field if non-nil, zero value otherwise.

### GetCryptoScanOk

`func (o *WorkflowDayBody) GetCryptoScanOk() (*AgentCompletedBody, bool)`

GetCryptoScanOk returns a tuple with the CryptoScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoScan

`func (o *WorkflowDayBody) SetCryptoScan(v AgentCompletedBody)`

SetCryptoScan sets CryptoScan field to given value.


### GetCryptographyBinary

`func (o *WorkflowDayBody) GetCryptographyBinary() AgentCompletedBody`

GetCryptographyBinary returns the CryptographyBinary field if non-nil, zero value otherwise.

### GetCryptographyBinaryOk

`func (o *WorkflowDayBody) GetCryptographyBinaryOk() (*AgentCompletedBody, bool)`

GetCryptographyBinaryOk returns a tuple with the CryptographyBinary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptographyBinary

`func (o *WorkflowDayBody) SetCryptographyBinary(v AgentCompletedBody)`

SetCryptographyBinary sets CryptographyBinary field to given value.


### GetDate

`func (o *WorkflowDayBody) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *WorkflowDayBody) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *WorkflowDayBody) SetDate(v string)`

SetDate sets Date field to given value.


### GetExecutionScan

`func (o *WorkflowDayBody) GetExecutionScan() AgentCompletedBody`

GetExecutionScan returns the ExecutionScan field if non-nil, zero value otherwise.

### GetExecutionScanOk

`func (o *WorkflowDayBody) GetExecutionScanOk() (*AgentCompletedBody, bool)`

GetExecutionScanOk returns a tuple with the ExecutionScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionScan

`func (o *WorkflowDayBody) SetExecutionScan(v AgentCompletedBody)`

SetExecutionScan sets ExecutionScan field to given value.


### GetFilesystemScan

`func (o *WorkflowDayBody) GetFilesystemScan() AgentCompletedBody`

GetFilesystemScan returns the FilesystemScan field if non-nil, zero value otherwise.

### GetFilesystemScanOk

`func (o *WorkflowDayBody) GetFilesystemScanOk() (*AgentCompletedBody, bool)`

GetFilesystemScanOk returns a tuple with the FilesystemScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemScan

`func (o *WorkflowDayBody) SetFilesystemScan(v AgentCompletedBody)`

SetFilesystemScan sets FilesystemScan field to given value.


### GetNetworkingScan

`func (o *WorkflowDayBody) GetNetworkingScan() AgentCompletedBody`

GetNetworkingScan returns the NetworkingScan field if non-nil, zero value otherwise.

### GetNetworkingScanOk

`func (o *WorkflowDayBody) GetNetworkingScanOk() (*AgentCompletedBody, bool)`

GetNetworkingScanOk returns a tuple with the NetworkingScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkingScan

`func (o *WorkflowDayBody) SetNetworkingScan(v AgentCompletedBody)`

SetNetworkingScan sets NetworkingScan field to given value.


### GetProtocols

`func (o *WorkflowDayBody) GetProtocols() AgentCompletedBody`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *WorkflowDayBody) GetProtocolsOk() (*AgentCompletedBody, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *WorkflowDayBody) SetProtocols(v AgentCompletedBody)`

SetProtocols sets Protocols field to given value.


### GetRemediation

`func (o *WorkflowDayBody) GetRemediation() AgentCompletedBody`

GetRemediation returns the Remediation field if non-nil, zero value otherwise.

### GetRemediationOk

`func (o *WorkflowDayBody) GetRemediationOk() (*AgentCompletedBody, bool)`

GetRemediationOk returns a tuple with the Remediation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediation

`func (o *WorkflowDayBody) SetRemediation(v AgentCompletedBody)`

SetRemediation sets Remediation field to given value.


### GetRenameUnnamedFunctions

`func (o *WorkflowDayBody) GetRenameUnnamedFunctions() AgentCompletedBody`

GetRenameUnnamedFunctions returns the RenameUnnamedFunctions field if non-nil, zero value otherwise.

### GetRenameUnnamedFunctionsOk

`func (o *WorkflowDayBody) GetRenameUnnamedFunctionsOk() (*AgentCompletedBody, bool)`

GetRenameUnnamedFunctionsOk returns a tuple with the RenameUnnamedFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameUnnamedFunctions

`func (o *WorkflowDayBody) SetRenameUnnamedFunctions(v AgentCompletedBody)`

SetRenameUnnamedFunctions sets RenameUnnamedFunctions field to given value.


### GetReportAnalysis

`func (o *WorkflowDayBody) GetReportAnalysis() ReportAnalysisBody`

GetReportAnalysis returns the ReportAnalysis field if non-nil, zero value otherwise.

### GetReportAnalysisOk

`func (o *WorkflowDayBody) GetReportAnalysisOk() (*ReportAnalysisBody, bool)`

GetReportAnalysisOk returns a tuple with the ReportAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportAnalysis

`func (o *WorkflowDayBody) SetReportAnalysis(v ReportAnalysisBody)`

SetReportAnalysis sets ReportAnalysis field to given value.


### GetSecrets

`func (o *WorkflowDayBody) GetSecrets() AgentCompletedBody`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *WorkflowDayBody) GetSecretsOk() (*AgentCompletedBody, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *WorkflowDayBody) SetSecrets(v AgentCompletedBody)`

SetSecrets sets Secrets field to given value.


### GetSecurityScan

`func (o *WorkflowDayBody) GetSecurityScan() AgentCompletedBody`

GetSecurityScan returns the SecurityScan field if non-nil, zero value otherwise.

### GetSecurityScanOk

`func (o *WorkflowDayBody) GetSecurityScanOk() (*AgentCompletedBody, bool)`

GetSecurityScanOk returns a tuple with the SecurityScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScan

`func (o *WorkflowDayBody) SetSecurityScan(v AgentCompletedBody)`

SetSecurityScan sets SecurityScan field to given value.


### GetTriage

`func (o *WorkflowDayBody) GetTriage() AgentCompletedBody`

GetTriage returns the Triage field if non-nil, zero value otherwise.

### GetTriageOk

`func (o *WorkflowDayBody) GetTriageOk() (*AgentCompletedBody, bool)`

GetTriageOk returns a tuple with the Triage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriage

`func (o *WorkflowDayBody) SetTriage(v AgentCompletedBody)`

SetTriage sets Triage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


