# AnalysisDetailOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | [**AnalysisAccessBody**](AnalysisAccessBody.md) |  | 
**AnalysisId** | **int64** |  | 
**AnalysisScope** | **string** |  | 
**Architecture** | **string** |  | 
**AutoRunAgents** | [**AutoRunAgentsBody**](AutoRunAgentsBody.md) |  | 
**BinaryDynamic** | **bool** |  | 
**BinaryFormat** | **string** |  | 
**BinaryName** | **string** |  | 
**BinarySize** | **int64** |  | 
**BinaryType** | **string** |  | 
**Creation** | **string** |  | 
**DashboardUrl** | **string** | URL to view this analysis in the dashboard | 
**Debug** | **bool** |  | 
**ModelName** | **string** |  | 
**RequestedConfig** | [**RequestedConfigBody**](RequestedConfigBody.md) | Snapshot of the configuration the analysis was submitted with | 
**Sha256Hash** | **string** |  | 

## Methods

### NewAnalysisDetailOutputBody

`func NewAnalysisDetailOutputBody(access AnalysisAccessBody, analysisId int64, analysisScope string, architecture string, autoRunAgents AutoRunAgentsBody, binaryDynamic bool, binaryFormat string, binaryName string, binarySize int64, binaryType string, creation string, dashboardUrl string, debug bool, modelName string, requestedConfig RequestedConfigBody, sha256Hash string, ) *AnalysisDetailOutputBody`

NewAnalysisDetailOutputBody instantiates a new AnalysisDetailOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisDetailOutputBodyWithDefaults

`func NewAnalysisDetailOutputBodyWithDefaults() *AnalysisDetailOutputBody`

NewAnalysisDetailOutputBodyWithDefaults instantiates a new AnalysisDetailOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *AnalysisDetailOutputBody) GetAccess() AnalysisAccessBody`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AnalysisDetailOutputBody) GetAccessOk() (*AnalysisAccessBody, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AnalysisDetailOutputBody) SetAccess(v AnalysisAccessBody)`

SetAccess sets Access field to given value.


### GetAnalysisId

`func (o *AnalysisDetailOutputBody) GetAnalysisId() int64`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *AnalysisDetailOutputBody) GetAnalysisIdOk() (*int64, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *AnalysisDetailOutputBody) SetAnalysisId(v int64)`

SetAnalysisId sets AnalysisId field to given value.


### GetAnalysisScope

`func (o *AnalysisDetailOutputBody) GetAnalysisScope() string`

GetAnalysisScope returns the AnalysisScope field if non-nil, zero value otherwise.

### GetAnalysisScopeOk

`func (o *AnalysisDetailOutputBody) GetAnalysisScopeOk() (*string, bool)`

GetAnalysisScopeOk returns a tuple with the AnalysisScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisScope

`func (o *AnalysisDetailOutputBody) SetAnalysisScope(v string)`

SetAnalysisScope sets AnalysisScope field to given value.


### GetArchitecture

`func (o *AnalysisDetailOutputBody) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *AnalysisDetailOutputBody) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *AnalysisDetailOutputBody) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.


### GetAutoRunAgents

`func (o *AnalysisDetailOutputBody) GetAutoRunAgents() AutoRunAgentsBody`

GetAutoRunAgents returns the AutoRunAgents field if non-nil, zero value otherwise.

### GetAutoRunAgentsOk

`func (o *AnalysisDetailOutputBody) GetAutoRunAgentsOk() (*AutoRunAgentsBody, bool)`

GetAutoRunAgentsOk returns a tuple with the AutoRunAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRunAgents

`func (o *AnalysisDetailOutputBody) SetAutoRunAgents(v AutoRunAgentsBody)`

SetAutoRunAgents sets AutoRunAgents field to given value.


### GetBinaryDynamic

`func (o *AnalysisDetailOutputBody) GetBinaryDynamic() bool`

GetBinaryDynamic returns the BinaryDynamic field if non-nil, zero value otherwise.

### GetBinaryDynamicOk

`func (o *AnalysisDetailOutputBody) GetBinaryDynamicOk() (*bool, bool)`

GetBinaryDynamicOk returns a tuple with the BinaryDynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDynamic

`func (o *AnalysisDetailOutputBody) SetBinaryDynamic(v bool)`

SetBinaryDynamic sets BinaryDynamic field to given value.


### GetBinaryFormat

`func (o *AnalysisDetailOutputBody) GetBinaryFormat() string`

GetBinaryFormat returns the BinaryFormat field if non-nil, zero value otherwise.

### GetBinaryFormatOk

`func (o *AnalysisDetailOutputBody) GetBinaryFormatOk() (*string, bool)`

GetBinaryFormatOk returns a tuple with the BinaryFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryFormat

`func (o *AnalysisDetailOutputBody) SetBinaryFormat(v string)`

SetBinaryFormat sets BinaryFormat field to given value.


### GetBinaryName

`func (o *AnalysisDetailOutputBody) GetBinaryName() string`

GetBinaryName returns the BinaryName field if non-nil, zero value otherwise.

### GetBinaryNameOk

`func (o *AnalysisDetailOutputBody) GetBinaryNameOk() (*string, bool)`

GetBinaryNameOk returns a tuple with the BinaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryName

`func (o *AnalysisDetailOutputBody) SetBinaryName(v string)`

SetBinaryName sets BinaryName field to given value.


### GetBinarySize

`func (o *AnalysisDetailOutputBody) GetBinarySize() int64`

GetBinarySize returns the BinarySize field if non-nil, zero value otherwise.

### GetBinarySizeOk

`func (o *AnalysisDetailOutputBody) GetBinarySizeOk() (*int64, bool)`

GetBinarySizeOk returns a tuple with the BinarySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinarySize

`func (o *AnalysisDetailOutputBody) SetBinarySize(v int64)`

SetBinarySize sets BinarySize field to given value.


### GetBinaryType

`func (o *AnalysisDetailOutputBody) GetBinaryType() string`

GetBinaryType returns the BinaryType field if non-nil, zero value otherwise.

### GetBinaryTypeOk

`func (o *AnalysisDetailOutputBody) GetBinaryTypeOk() (*string, bool)`

GetBinaryTypeOk returns a tuple with the BinaryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryType

`func (o *AnalysisDetailOutputBody) SetBinaryType(v string)`

SetBinaryType sets BinaryType field to given value.


### GetCreation

`func (o *AnalysisDetailOutputBody) GetCreation() string`

GetCreation returns the Creation field if non-nil, zero value otherwise.

### GetCreationOk

`func (o *AnalysisDetailOutputBody) GetCreationOk() (*string, bool)`

GetCreationOk returns a tuple with the Creation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreation

`func (o *AnalysisDetailOutputBody) SetCreation(v string)`

SetCreation sets Creation field to given value.


### GetDashboardUrl

`func (o *AnalysisDetailOutputBody) GetDashboardUrl() string`

GetDashboardUrl returns the DashboardUrl field if non-nil, zero value otherwise.

### GetDashboardUrlOk

`func (o *AnalysisDetailOutputBody) GetDashboardUrlOk() (*string, bool)`

GetDashboardUrlOk returns a tuple with the DashboardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardUrl

`func (o *AnalysisDetailOutputBody) SetDashboardUrl(v string)`

SetDashboardUrl sets DashboardUrl field to given value.


### GetDebug

`func (o *AnalysisDetailOutputBody) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *AnalysisDetailOutputBody) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *AnalysisDetailOutputBody) SetDebug(v bool)`

SetDebug sets Debug field to given value.


### GetModelName

`func (o *AnalysisDetailOutputBody) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *AnalysisDetailOutputBody) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *AnalysisDetailOutputBody) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetRequestedConfig

`func (o *AnalysisDetailOutputBody) GetRequestedConfig() RequestedConfigBody`

GetRequestedConfig returns the RequestedConfig field if non-nil, zero value otherwise.

### GetRequestedConfigOk

`func (o *AnalysisDetailOutputBody) GetRequestedConfigOk() (*RequestedConfigBody, bool)`

GetRequestedConfigOk returns a tuple with the RequestedConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedConfig

`func (o *AnalysisDetailOutputBody) SetRequestedConfig(v RequestedConfigBody)`

SetRequestedConfig sets RequestedConfig field to given value.


### GetSha256Hash

`func (o *AnalysisDetailOutputBody) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *AnalysisDetailOutputBody) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *AnalysisDetailOutputBody) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


