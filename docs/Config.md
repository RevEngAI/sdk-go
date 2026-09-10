# Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GenerateCapabilities** | Pointer to **bool** |  | [optional] 
**NoCache** | Pointer to **bool** |  | [optional] 
**SandboxConfig** | Pointer to [**SandboxConfig**](SandboxConfig.md) |  | [optional] 
**ScrapeThirdPartyConfig** | Pointer to [**ScrapeThirdPartyConfig**](ScrapeThirdPartyConfig.md) |  | [optional] 

## Methods

### NewConfig

`func NewConfig() *Config`

NewConfig instantiates a new Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigWithDefaults

`func NewConfigWithDefaults() *Config`

NewConfigWithDefaults instantiates a new Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenerateCapabilities

`func (o *Config) GetGenerateCapabilities() bool`

GetGenerateCapabilities returns the GenerateCapabilities field if non-nil, zero value otherwise.

### GetGenerateCapabilitiesOk

`func (o *Config) GetGenerateCapabilitiesOk() (*bool, bool)`

GetGenerateCapabilitiesOk returns a tuple with the GenerateCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateCapabilities

`func (o *Config) SetGenerateCapabilities(v bool)`

SetGenerateCapabilities sets GenerateCapabilities field to given value.

### HasGenerateCapabilities

`func (o *Config) HasGenerateCapabilities() bool`

HasGenerateCapabilities returns a boolean if a field has been set.

### GetNoCache

`func (o *Config) GetNoCache() bool`

GetNoCache returns the NoCache field if non-nil, zero value otherwise.

### GetNoCacheOk

`func (o *Config) GetNoCacheOk() (*bool, bool)`

GetNoCacheOk returns a tuple with the NoCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoCache

`func (o *Config) SetNoCache(v bool)`

SetNoCache sets NoCache field to given value.

### HasNoCache

`func (o *Config) HasNoCache() bool`

HasNoCache returns a boolean if a field has been set.

### GetSandboxConfig

`func (o *Config) GetSandboxConfig() SandboxConfig`

GetSandboxConfig returns the SandboxConfig field if non-nil, zero value otherwise.

### GetSandboxConfigOk

`func (o *Config) GetSandboxConfigOk() (*SandboxConfig, bool)`

GetSandboxConfigOk returns a tuple with the SandboxConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxConfig

`func (o *Config) SetSandboxConfig(v SandboxConfig)`

SetSandboxConfig sets SandboxConfig field to given value.

### HasSandboxConfig

`func (o *Config) HasSandboxConfig() bool`

HasSandboxConfig returns a boolean if a field has been set.

### GetScrapeThirdPartyConfig

`func (o *Config) GetScrapeThirdPartyConfig() ScrapeThirdPartyConfig`

GetScrapeThirdPartyConfig returns the ScrapeThirdPartyConfig field if non-nil, zero value otherwise.

### GetScrapeThirdPartyConfigOk

`func (o *Config) GetScrapeThirdPartyConfigOk() (*ScrapeThirdPartyConfig, bool)`

GetScrapeThirdPartyConfigOk returns a tuple with the ScrapeThirdPartyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrapeThirdPartyConfig

`func (o *Config) SetScrapeThirdPartyConfig(v ScrapeThirdPartyConfig)`

SetScrapeThirdPartyConfig sets ScrapeThirdPartyConfig field to given value.

### HasScrapeThirdPartyConfig

`func (o *Config) HasScrapeThirdPartyConfig() bool`

HasScrapeThirdPartyConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


