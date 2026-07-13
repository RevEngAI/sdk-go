# SandboxOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**CommandLineArgs** | Pointer to **string** | The command line parameters to pass to the dynamic execution sandbox. Requires &#x60;sandbox&#x60; to be True. | [optional] [default to ""]
**StartMethod** | Pointer to [**NullableSandboxStartMethod**](SandboxStartMethod.md) |  | [optional] 
**Timeout** | Pointer to [**SandboxTimeout**](SandboxTimeout.md) | Maximum execution time for the sandbox run, in seconds. Allowed values: 120 (2m), 180 (3m), 300 (5m), 600 (10m). | [optional] [default to SANDBOXTIMEOUT__120]
**ArchiveSha256Hash** | Pointer to **NullableString** |  | [optional] 
**ArchiveEntryPath** | Pointer to **NullableString** |  | [optional] 
**ArchivePassword** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSandboxOptions

`func NewSandboxOptions() *SandboxOptions`

NewSandboxOptions instantiates a new SandboxOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxOptionsWithDefaults

`func NewSandboxOptionsWithDefaults() *SandboxOptions`

NewSandboxOptionsWithDefaults instantiates a new SandboxOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SandboxOptions) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SandboxOptions) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SandboxOptions) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SandboxOptions) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCommandLineArgs

`func (o *SandboxOptions) GetCommandLineArgs() string`

GetCommandLineArgs returns the CommandLineArgs field if non-nil, zero value otherwise.

### GetCommandLineArgsOk

`func (o *SandboxOptions) GetCommandLineArgsOk() (*string, bool)`

GetCommandLineArgsOk returns a tuple with the CommandLineArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineArgs

`func (o *SandboxOptions) SetCommandLineArgs(v string)`

SetCommandLineArgs sets CommandLineArgs field to given value.

### HasCommandLineArgs

`func (o *SandboxOptions) HasCommandLineArgs() bool`

HasCommandLineArgs returns a boolean if a field has been set.

### GetStartMethod

`func (o *SandboxOptions) GetStartMethod() SandboxStartMethod`

GetStartMethod returns the StartMethod field if non-nil, zero value otherwise.

### GetStartMethodOk

`func (o *SandboxOptions) GetStartMethodOk() (*SandboxStartMethod, bool)`

GetStartMethodOk returns a tuple with the StartMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMethod

`func (o *SandboxOptions) SetStartMethod(v SandboxStartMethod)`

SetStartMethod sets StartMethod field to given value.

### HasStartMethod

`func (o *SandboxOptions) HasStartMethod() bool`

HasStartMethod returns a boolean if a field has been set.

### SetStartMethodNil

`func (o *SandboxOptions) SetStartMethodNil(b bool)`

 SetStartMethodNil sets the value for StartMethod to be an explicit nil

### UnsetStartMethod
`func (o *SandboxOptions) UnsetStartMethod()`

UnsetStartMethod ensures that no value is present for StartMethod, not even an explicit nil
### GetTimeout

`func (o *SandboxOptions) GetTimeout() SandboxTimeout`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SandboxOptions) GetTimeoutOk() (*SandboxTimeout, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SandboxOptions) SetTimeout(v SandboxTimeout)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *SandboxOptions) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetArchiveSha256Hash

`func (o *SandboxOptions) GetArchiveSha256Hash() string`

GetArchiveSha256Hash returns the ArchiveSha256Hash field if non-nil, zero value otherwise.

### GetArchiveSha256HashOk

`func (o *SandboxOptions) GetArchiveSha256HashOk() (*string, bool)`

GetArchiveSha256HashOk returns a tuple with the ArchiveSha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveSha256Hash

`func (o *SandboxOptions) SetArchiveSha256Hash(v string)`

SetArchiveSha256Hash sets ArchiveSha256Hash field to given value.

### HasArchiveSha256Hash

`func (o *SandboxOptions) HasArchiveSha256Hash() bool`

HasArchiveSha256Hash returns a boolean if a field has been set.

### SetArchiveSha256HashNil

`func (o *SandboxOptions) SetArchiveSha256HashNil(b bool)`

 SetArchiveSha256HashNil sets the value for ArchiveSha256Hash to be an explicit nil

### UnsetArchiveSha256Hash
`func (o *SandboxOptions) UnsetArchiveSha256Hash()`

UnsetArchiveSha256Hash ensures that no value is present for ArchiveSha256Hash, not even an explicit nil
### GetArchiveEntryPath

`func (o *SandboxOptions) GetArchiveEntryPath() string`

GetArchiveEntryPath returns the ArchiveEntryPath field if non-nil, zero value otherwise.

### GetArchiveEntryPathOk

`func (o *SandboxOptions) GetArchiveEntryPathOk() (*string, bool)`

GetArchiveEntryPathOk returns a tuple with the ArchiveEntryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveEntryPath

`func (o *SandboxOptions) SetArchiveEntryPath(v string)`

SetArchiveEntryPath sets ArchiveEntryPath field to given value.

### HasArchiveEntryPath

`func (o *SandboxOptions) HasArchiveEntryPath() bool`

HasArchiveEntryPath returns a boolean if a field has been set.

### SetArchiveEntryPathNil

`func (o *SandboxOptions) SetArchiveEntryPathNil(b bool)`

 SetArchiveEntryPathNil sets the value for ArchiveEntryPath to be an explicit nil

### UnsetArchiveEntryPath
`func (o *SandboxOptions) UnsetArchiveEntryPath()`

UnsetArchiveEntryPath ensures that no value is present for ArchiveEntryPath, not even an explicit nil
### GetArchivePassword

`func (o *SandboxOptions) GetArchivePassword() string`

GetArchivePassword returns the ArchivePassword field if non-nil, zero value otherwise.

### GetArchivePasswordOk

`func (o *SandboxOptions) GetArchivePasswordOk() (*string, bool)`

GetArchivePasswordOk returns a tuple with the ArchivePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivePassword

`func (o *SandboxOptions) SetArchivePassword(v string)`

SetArchivePassword sets ArchivePassword field to given value.

### HasArchivePassword

`func (o *SandboxOptions) HasArchivePassword() bool`

HasArchivePassword returns a boolean if a field has been set.

### SetArchivePasswordNil

`func (o *SandboxOptions) SetArchivePasswordNil(b bool)`

 SetArchivePasswordNil sets the value for ArchivePassword to be an explicit nil

### UnsetArchivePassword
`func (o *SandboxOptions) UnsetArchivePassword()`

UnsetArchivePassword ensures that no value is present for ArchivePassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


