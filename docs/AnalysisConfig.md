# AnalysisConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScrapeThirdPartyConfig** | Pointer to [**ScrapeThirdPartyConfig**](ScrapeThirdPartyConfig.md) | Settings to scrape third party sources | [optional] 
**GenerateCves** | Pointer to **bool** | A configuration option for fetching CVEs data. | [optional] [default to false]
**GenerateSbom** | Pointer to **bool** | A configuration option for generating software bill of materials data. | [optional] [default to false]
**GenerateCapabilities** | Pointer to **bool** | A configuration option for generating capabilities of a binary | [optional] [default to false]
**NoCache** | Pointer to **bool** | When enabled, skips using cached data within the processing. | [optional] [default to false]
**AdvancedAnalysis** | Pointer to **bool** | Enables an advanced security analysis. | [optional] [default to false]
**SandboxConfig** | Pointer to [**SandboxOptions**](SandboxOptions.md) | Including a sandbox config enables the dynamic execution sandbox | [optional] 

## Methods

### NewAnalysisConfig

`func NewAnalysisConfig() *AnalysisConfig`

NewAnalysisConfig instantiates a new AnalysisConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisConfigWithDefaults

`func NewAnalysisConfigWithDefaults() *AnalysisConfig`

NewAnalysisConfigWithDefaults instantiates a new AnalysisConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScrapeThirdPartyConfig

`func (o *AnalysisConfig) GetScrapeThirdPartyConfig() ScrapeThirdPartyConfig`

GetScrapeThirdPartyConfig returns the ScrapeThirdPartyConfig field if non-nil, zero value otherwise.

### GetScrapeThirdPartyConfigOk

`func (o *AnalysisConfig) GetScrapeThirdPartyConfigOk() (*ScrapeThirdPartyConfig, bool)`

GetScrapeThirdPartyConfigOk returns a tuple with the ScrapeThirdPartyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrapeThirdPartyConfig

`func (o *AnalysisConfig) SetScrapeThirdPartyConfig(v ScrapeThirdPartyConfig)`

SetScrapeThirdPartyConfig sets ScrapeThirdPartyConfig field to given value.

### HasScrapeThirdPartyConfig

`func (o *AnalysisConfig) HasScrapeThirdPartyConfig() bool`

HasScrapeThirdPartyConfig returns a boolean if a field has been set.

### GetGenerateCves

`func (o *AnalysisConfig) GetGenerateCves() bool`

GetGenerateCves returns the GenerateCves field if non-nil, zero value otherwise.

### GetGenerateCvesOk

`func (o *AnalysisConfig) GetGenerateCvesOk() (*bool, bool)`

GetGenerateCvesOk returns a tuple with the GenerateCves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateCves

`func (o *AnalysisConfig) SetGenerateCves(v bool)`

SetGenerateCves sets GenerateCves field to given value.

### HasGenerateCves

`func (o *AnalysisConfig) HasGenerateCves() bool`

HasGenerateCves returns a boolean if a field has been set.

### GetGenerateSbom

`func (o *AnalysisConfig) GetGenerateSbom() bool`

GetGenerateSbom returns the GenerateSbom field if non-nil, zero value otherwise.

### GetGenerateSbomOk

`func (o *AnalysisConfig) GetGenerateSbomOk() (*bool, bool)`

GetGenerateSbomOk returns a tuple with the GenerateSbom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateSbom

`func (o *AnalysisConfig) SetGenerateSbom(v bool)`

SetGenerateSbom sets GenerateSbom field to given value.

### HasGenerateSbom

`func (o *AnalysisConfig) HasGenerateSbom() bool`

HasGenerateSbom returns a boolean if a field has been set.

### GetGenerateCapabilities

`func (o *AnalysisConfig) GetGenerateCapabilities() bool`

GetGenerateCapabilities returns the GenerateCapabilities field if non-nil, zero value otherwise.

### GetGenerateCapabilitiesOk

`func (o *AnalysisConfig) GetGenerateCapabilitiesOk() (*bool, bool)`

GetGenerateCapabilitiesOk returns a tuple with the GenerateCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateCapabilities

`func (o *AnalysisConfig) SetGenerateCapabilities(v bool)`

SetGenerateCapabilities sets GenerateCapabilities field to given value.

### HasGenerateCapabilities

`func (o *AnalysisConfig) HasGenerateCapabilities() bool`

HasGenerateCapabilities returns a boolean if a field has been set.

### GetNoCache

`func (o *AnalysisConfig) GetNoCache() bool`

GetNoCache returns the NoCache field if non-nil, zero value otherwise.

### GetNoCacheOk

`func (o *AnalysisConfig) GetNoCacheOk() (*bool, bool)`

GetNoCacheOk returns a tuple with the NoCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoCache

`func (o *AnalysisConfig) SetNoCache(v bool)`

SetNoCache sets NoCache field to given value.

### HasNoCache

`func (o *AnalysisConfig) HasNoCache() bool`

HasNoCache returns a boolean if a field has been set.

### GetAdvancedAnalysis

`func (o *AnalysisConfig) GetAdvancedAnalysis() bool`

GetAdvancedAnalysis returns the AdvancedAnalysis field if non-nil, zero value otherwise.

### GetAdvancedAnalysisOk

`func (o *AnalysisConfig) GetAdvancedAnalysisOk() (*bool, bool)`

GetAdvancedAnalysisOk returns a tuple with the AdvancedAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedAnalysis

`func (o *AnalysisConfig) SetAdvancedAnalysis(v bool)`

SetAdvancedAnalysis sets AdvancedAnalysis field to given value.

### HasAdvancedAnalysis

`func (o *AnalysisConfig) HasAdvancedAnalysis() bool`

HasAdvancedAnalysis returns a boolean if a field has been set.

### GetSandboxConfig

`func (o *AnalysisConfig) GetSandboxConfig() SandboxOptions`

GetSandboxConfig returns the SandboxConfig field if non-nil, zero value otherwise.

### GetSandboxConfigOk

`func (o *AnalysisConfig) GetSandboxConfigOk() (*SandboxOptions, bool)`

GetSandboxConfigOk returns a tuple with the SandboxConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxConfig

`func (o *AnalysisConfig) SetSandboxConfig(v SandboxOptions)`

SetSandboxConfig sets SandboxConfig field to given value.

### HasSandboxConfig

`func (o *AnalysisConfig) HasSandboxConfig() bool`

HasSandboxConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


