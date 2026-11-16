# DynamicExecutionRunConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveEntryPath** | **string** | Relative path of the archive entry executed. Empty when no archive was used. | 
**ArchiveSha256Hash** | **string** | SHA-256 of the archive sent to the sandbox instead of the analysed binary. Empty when no archive was used. | 
**CommandLineArgs** | **string** | Command-line arguments passed to the sample. Empty when none were given. | 
**HasArchivePassword** | **bool** | Whether a password was given for the archive. The password itself is never returned. | 
**StartMethod** | **NullableString** | How the sandbox launched the sample. Null when the sandbox&#39;s standard behaviour was used. | 
**TimeoutSeconds** | **int64** | Maximum sandbox execution time in seconds | 
**Url** | **string** | URL opened in the sandbox&#39;s browser instead of a sample. Empty for a sample run. | 

## Methods

### NewDynamicExecutionRunConfig

`func NewDynamicExecutionRunConfig(archiveEntryPath string, archiveSha256Hash string, commandLineArgs string, hasArchivePassword bool, startMethod NullableString, timeoutSeconds int64, url string, ) *DynamicExecutionRunConfig`

NewDynamicExecutionRunConfig instantiates a new DynamicExecutionRunConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicExecutionRunConfigWithDefaults

`func NewDynamicExecutionRunConfigWithDefaults() *DynamicExecutionRunConfig`

NewDynamicExecutionRunConfigWithDefaults instantiates a new DynamicExecutionRunConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveEntryPath

`func (o *DynamicExecutionRunConfig) GetArchiveEntryPath() string`

GetArchiveEntryPath returns the ArchiveEntryPath field if non-nil, zero value otherwise.

### GetArchiveEntryPathOk

`func (o *DynamicExecutionRunConfig) GetArchiveEntryPathOk() (*string, bool)`

GetArchiveEntryPathOk returns a tuple with the ArchiveEntryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveEntryPath

`func (o *DynamicExecutionRunConfig) SetArchiveEntryPath(v string)`

SetArchiveEntryPath sets ArchiveEntryPath field to given value.


### GetArchiveSha256Hash

`func (o *DynamicExecutionRunConfig) GetArchiveSha256Hash() string`

GetArchiveSha256Hash returns the ArchiveSha256Hash field if non-nil, zero value otherwise.

### GetArchiveSha256HashOk

`func (o *DynamicExecutionRunConfig) GetArchiveSha256HashOk() (*string, bool)`

GetArchiveSha256HashOk returns a tuple with the ArchiveSha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveSha256Hash

`func (o *DynamicExecutionRunConfig) SetArchiveSha256Hash(v string)`

SetArchiveSha256Hash sets ArchiveSha256Hash field to given value.


### GetCommandLineArgs

`func (o *DynamicExecutionRunConfig) GetCommandLineArgs() string`

GetCommandLineArgs returns the CommandLineArgs field if non-nil, zero value otherwise.

### GetCommandLineArgsOk

`func (o *DynamicExecutionRunConfig) GetCommandLineArgsOk() (*string, bool)`

GetCommandLineArgsOk returns a tuple with the CommandLineArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineArgs

`func (o *DynamicExecutionRunConfig) SetCommandLineArgs(v string)`

SetCommandLineArgs sets CommandLineArgs field to given value.


### GetHasArchivePassword

`func (o *DynamicExecutionRunConfig) GetHasArchivePassword() bool`

GetHasArchivePassword returns the HasArchivePassword field if non-nil, zero value otherwise.

### GetHasArchivePasswordOk

`func (o *DynamicExecutionRunConfig) GetHasArchivePasswordOk() (*bool, bool)`

GetHasArchivePasswordOk returns a tuple with the HasArchivePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasArchivePassword

`func (o *DynamicExecutionRunConfig) SetHasArchivePassword(v bool)`

SetHasArchivePassword sets HasArchivePassword field to given value.


### GetStartMethod

`func (o *DynamicExecutionRunConfig) GetStartMethod() string`

GetStartMethod returns the StartMethod field if non-nil, zero value otherwise.

### GetStartMethodOk

`func (o *DynamicExecutionRunConfig) GetStartMethodOk() (*string, bool)`

GetStartMethodOk returns a tuple with the StartMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMethod

`func (o *DynamicExecutionRunConfig) SetStartMethod(v string)`

SetStartMethod sets StartMethod field to given value.


### SetStartMethodNil

`func (o *DynamicExecutionRunConfig) SetStartMethodNil(b bool)`

 SetStartMethodNil sets the value for StartMethod to be an explicit nil

### UnsetStartMethod
`func (o *DynamicExecutionRunConfig) UnsetStartMethod()`

UnsetStartMethod ensures that no value is present for StartMethod, not even an explicit nil
### GetTimeoutSeconds

`func (o *DynamicExecutionRunConfig) GetTimeoutSeconds() int64`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *DynamicExecutionRunConfig) GetTimeoutSecondsOk() (*int64, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *DynamicExecutionRunConfig) SetTimeoutSeconds(v int64)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.


### GetUrl

`func (o *DynamicExecutionRunConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DynamicExecutionRunConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DynamicExecutionRunConfig) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


