# SandboxConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveEntryPath** | Pointer to **string** |  | [optional] 
**ArchivePassword** | Pointer to **string** |  | [optional] 
**ArchiveSha256Hash** | Pointer to **string** |  | [optional] 
**CommandLineArgs** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**StartMethod** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to **int64** |  | [optional] [default to 120]

## Methods

### NewSandboxConfig

`func NewSandboxConfig() *SandboxConfig`

NewSandboxConfig instantiates a new SandboxConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxConfigWithDefaults

`func NewSandboxConfigWithDefaults() *SandboxConfig`

NewSandboxConfigWithDefaults instantiates a new SandboxConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveEntryPath

`func (o *SandboxConfig) GetArchiveEntryPath() string`

GetArchiveEntryPath returns the ArchiveEntryPath field if non-nil, zero value otherwise.

### GetArchiveEntryPathOk

`func (o *SandboxConfig) GetArchiveEntryPathOk() (*string, bool)`

GetArchiveEntryPathOk returns a tuple with the ArchiveEntryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveEntryPath

`func (o *SandboxConfig) SetArchiveEntryPath(v string)`

SetArchiveEntryPath sets ArchiveEntryPath field to given value.

### HasArchiveEntryPath

`func (o *SandboxConfig) HasArchiveEntryPath() bool`

HasArchiveEntryPath returns a boolean if a field has been set.

### GetArchivePassword

`func (o *SandboxConfig) GetArchivePassword() string`

GetArchivePassword returns the ArchivePassword field if non-nil, zero value otherwise.

### GetArchivePasswordOk

`func (o *SandboxConfig) GetArchivePasswordOk() (*string, bool)`

GetArchivePasswordOk returns a tuple with the ArchivePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivePassword

`func (o *SandboxConfig) SetArchivePassword(v string)`

SetArchivePassword sets ArchivePassword field to given value.

### HasArchivePassword

`func (o *SandboxConfig) HasArchivePassword() bool`

HasArchivePassword returns a boolean if a field has been set.

### GetArchiveSha256Hash

`func (o *SandboxConfig) GetArchiveSha256Hash() string`

GetArchiveSha256Hash returns the ArchiveSha256Hash field if non-nil, zero value otherwise.

### GetArchiveSha256HashOk

`func (o *SandboxConfig) GetArchiveSha256HashOk() (*string, bool)`

GetArchiveSha256HashOk returns a tuple with the ArchiveSha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveSha256Hash

`func (o *SandboxConfig) SetArchiveSha256Hash(v string)`

SetArchiveSha256Hash sets ArchiveSha256Hash field to given value.

### HasArchiveSha256Hash

`func (o *SandboxConfig) HasArchiveSha256Hash() bool`

HasArchiveSha256Hash returns a boolean if a field has been set.

### GetCommandLineArgs

`func (o *SandboxConfig) GetCommandLineArgs() string`

GetCommandLineArgs returns the CommandLineArgs field if non-nil, zero value otherwise.

### GetCommandLineArgsOk

`func (o *SandboxConfig) GetCommandLineArgsOk() (*string, bool)`

GetCommandLineArgsOk returns a tuple with the CommandLineArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineArgs

`func (o *SandboxConfig) SetCommandLineArgs(v string)`

SetCommandLineArgs sets CommandLineArgs field to given value.

### HasCommandLineArgs

`func (o *SandboxConfig) HasCommandLineArgs() bool`

HasCommandLineArgs returns a boolean if a field has been set.

### GetEnabled

`func (o *SandboxConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SandboxConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SandboxConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SandboxConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStartMethod

`func (o *SandboxConfig) GetStartMethod() string`

GetStartMethod returns the StartMethod field if non-nil, zero value otherwise.

### GetStartMethodOk

`func (o *SandboxConfig) GetStartMethodOk() (*string, bool)`

GetStartMethodOk returns a tuple with the StartMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMethod

`func (o *SandboxConfig) SetStartMethod(v string)`

SetStartMethod sets StartMethod field to given value.

### HasStartMethod

`func (o *SandboxConfig) HasStartMethod() bool`

HasStartMethod returns a boolean if a field has been set.

### GetTimeout

`func (o *SandboxConfig) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SandboxConfig) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SandboxConfig) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *SandboxConfig) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


