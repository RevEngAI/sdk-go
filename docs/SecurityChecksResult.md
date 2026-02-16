# SecurityChecksResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionId** | **int64** |  | 
**FunctionName** | **string** |  | 
**Name** | **string** |  | 
**VulnClass** | [**VulnerabilityType**](VulnerabilityType.md) |  | 
**Description** | **string** |  | 
**Remediation** | **string** |  | 
**Confidence** | [**ConfidenceType**](ConfidenceType.md) |  | 
**Severity** | [**SeverityType**](SeverityType.md) |  | 

## Methods

### NewSecurityChecksResult

`func NewSecurityChecksResult(functionId int64, functionName string, name string, vulnClass VulnerabilityType, description string, remediation string, confidence ConfidenceType, severity SeverityType, ) *SecurityChecksResult`

NewSecurityChecksResult instantiates a new SecurityChecksResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityChecksResultWithDefaults

`func NewSecurityChecksResultWithDefaults() *SecurityChecksResult`

NewSecurityChecksResultWithDefaults instantiates a new SecurityChecksResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionId

`func (o *SecurityChecksResult) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *SecurityChecksResult) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *SecurityChecksResult) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetFunctionName

`func (o *SecurityChecksResult) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *SecurityChecksResult) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *SecurityChecksResult) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetName

`func (o *SecurityChecksResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityChecksResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityChecksResult) SetName(v string)`

SetName sets Name field to given value.


### GetVulnClass

`func (o *SecurityChecksResult) GetVulnClass() VulnerabilityType`

GetVulnClass returns the VulnClass field if non-nil, zero value otherwise.

### GetVulnClassOk

`func (o *SecurityChecksResult) GetVulnClassOk() (*VulnerabilityType, bool)`

GetVulnClassOk returns a tuple with the VulnClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnClass

`func (o *SecurityChecksResult) SetVulnClass(v VulnerabilityType)`

SetVulnClass sets VulnClass field to given value.


### GetDescription

`func (o *SecurityChecksResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityChecksResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityChecksResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetRemediation

`func (o *SecurityChecksResult) GetRemediation() string`

GetRemediation returns the Remediation field if non-nil, zero value otherwise.

### GetRemediationOk

`func (o *SecurityChecksResult) GetRemediationOk() (*string, bool)`

GetRemediationOk returns a tuple with the Remediation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediation

`func (o *SecurityChecksResult) SetRemediation(v string)`

SetRemediation sets Remediation field to given value.


### GetConfidence

`func (o *SecurityChecksResult) GetConfidence() ConfidenceType`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *SecurityChecksResult) GetConfidenceOk() (*ConfidenceType, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *SecurityChecksResult) SetConfidence(v ConfidenceType)`

SetConfidence sets Confidence field to given value.


### GetSeverity

`func (o *SecurityChecksResult) GetSeverity() SeverityType`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *SecurityChecksResult) GetSeverityOk() (*SeverityType, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *SecurityChecksResult) SetSeverity(v SeverityType)`

SetSeverity sets Severity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


