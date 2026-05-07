# ReportOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveEntryPath** | Pointer to **string** |  | [optional] 
**ExtractArchive** | Pointer to **bool** |  | [optional] 
**GuestTargetDirectory** | Pointer to **string** |  | [optional] 
**GuestWorkingDirectory** | Pointer to **string** |  | [optional] 
**NetEnable** | Pointer to **bool** |  | [optional] 
**OsProfile** | Pointer to **string** |  | [optional] 
**Plugins** | Pointer to **[]string** |  | [optional] 
**Preset** | Pointer to **string** |  | [optional] 
**SampleFilename** | Pointer to **string** |  | [optional] 
**StartCommand** | Pointer to **string** |  | [optional] 
**StartMethod** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to **int64** |  | [optional] 

## Methods

### NewReportOptions

`func NewReportOptions() *ReportOptions`

NewReportOptions instantiates a new ReportOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportOptionsWithDefaults

`func NewReportOptionsWithDefaults() *ReportOptions`

NewReportOptionsWithDefaults instantiates a new ReportOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveEntryPath

`func (o *ReportOptions) GetArchiveEntryPath() string`

GetArchiveEntryPath returns the ArchiveEntryPath field if non-nil, zero value otherwise.

### GetArchiveEntryPathOk

`func (o *ReportOptions) GetArchiveEntryPathOk() (*string, bool)`

GetArchiveEntryPathOk returns a tuple with the ArchiveEntryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveEntryPath

`func (o *ReportOptions) SetArchiveEntryPath(v string)`

SetArchiveEntryPath sets ArchiveEntryPath field to given value.

### HasArchiveEntryPath

`func (o *ReportOptions) HasArchiveEntryPath() bool`

HasArchiveEntryPath returns a boolean if a field has been set.

### GetExtractArchive

`func (o *ReportOptions) GetExtractArchive() bool`

GetExtractArchive returns the ExtractArchive field if non-nil, zero value otherwise.

### GetExtractArchiveOk

`func (o *ReportOptions) GetExtractArchiveOk() (*bool, bool)`

GetExtractArchiveOk returns a tuple with the ExtractArchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractArchive

`func (o *ReportOptions) SetExtractArchive(v bool)`

SetExtractArchive sets ExtractArchive field to given value.

### HasExtractArchive

`func (o *ReportOptions) HasExtractArchive() bool`

HasExtractArchive returns a boolean if a field has been set.

### GetGuestTargetDirectory

`func (o *ReportOptions) GetGuestTargetDirectory() string`

GetGuestTargetDirectory returns the GuestTargetDirectory field if non-nil, zero value otherwise.

### GetGuestTargetDirectoryOk

`func (o *ReportOptions) GetGuestTargetDirectoryOk() (*string, bool)`

GetGuestTargetDirectoryOk returns a tuple with the GuestTargetDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestTargetDirectory

`func (o *ReportOptions) SetGuestTargetDirectory(v string)`

SetGuestTargetDirectory sets GuestTargetDirectory field to given value.

### HasGuestTargetDirectory

`func (o *ReportOptions) HasGuestTargetDirectory() bool`

HasGuestTargetDirectory returns a boolean if a field has been set.

### GetGuestWorkingDirectory

`func (o *ReportOptions) GetGuestWorkingDirectory() string`

GetGuestWorkingDirectory returns the GuestWorkingDirectory field if non-nil, zero value otherwise.

### GetGuestWorkingDirectoryOk

`func (o *ReportOptions) GetGuestWorkingDirectoryOk() (*string, bool)`

GetGuestWorkingDirectoryOk returns a tuple with the GuestWorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestWorkingDirectory

`func (o *ReportOptions) SetGuestWorkingDirectory(v string)`

SetGuestWorkingDirectory sets GuestWorkingDirectory field to given value.

### HasGuestWorkingDirectory

`func (o *ReportOptions) HasGuestWorkingDirectory() bool`

HasGuestWorkingDirectory returns a boolean if a field has been set.

### GetNetEnable

`func (o *ReportOptions) GetNetEnable() bool`

GetNetEnable returns the NetEnable field if non-nil, zero value otherwise.

### GetNetEnableOk

`func (o *ReportOptions) GetNetEnableOk() (*bool, bool)`

GetNetEnableOk returns a tuple with the NetEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetEnable

`func (o *ReportOptions) SetNetEnable(v bool)`

SetNetEnable sets NetEnable field to given value.

### HasNetEnable

`func (o *ReportOptions) HasNetEnable() bool`

HasNetEnable returns a boolean if a field has been set.

### GetOsProfile

`func (o *ReportOptions) GetOsProfile() string`

GetOsProfile returns the OsProfile field if non-nil, zero value otherwise.

### GetOsProfileOk

`func (o *ReportOptions) GetOsProfileOk() (*string, bool)`

GetOsProfileOk returns a tuple with the OsProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsProfile

`func (o *ReportOptions) SetOsProfile(v string)`

SetOsProfile sets OsProfile field to given value.

### HasOsProfile

`func (o *ReportOptions) HasOsProfile() bool`

HasOsProfile returns a boolean if a field has been set.

### GetPlugins

`func (o *ReportOptions) GetPlugins() []string`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *ReportOptions) GetPluginsOk() (*[]string, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *ReportOptions) SetPlugins(v []string)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *ReportOptions) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.

### SetPluginsNil

`func (o *ReportOptions) SetPluginsNil(b bool)`

 SetPluginsNil sets the value for Plugins to be an explicit nil

### UnsetPlugins
`func (o *ReportOptions) UnsetPlugins()`

UnsetPlugins ensures that no value is present for Plugins, not even an explicit nil
### GetPreset

`func (o *ReportOptions) GetPreset() string`

GetPreset returns the Preset field if non-nil, zero value otherwise.

### GetPresetOk

`func (o *ReportOptions) GetPresetOk() (*string, bool)`

GetPresetOk returns a tuple with the Preset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreset

`func (o *ReportOptions) SetPreset(v string)`

SetPreset sets Preset field to given value.

### HasPreset

`func (o *ReportOptions) HasPreset() bool`

HasPreset returns a boolean if a field has been set.

### GetSampleFilename

`func (o *ReportOptions) GetSampleFilename() string`

GetSampleFilename returns the SampleFilename field if non-nil, zero value otherwise.

### GetSampleFilenameOk

`func (o *ReportOptions) GetSampleFilenameOk() (*string, bool)`

GetSampleFilenameOk returns a tuple with the SampleFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleFilename

`func (o *ReportOptions) SetSampleFilename(v string)`

SetSampleFilename sets SampleFilename field to given value.

### HasSampleFilename

`func (o *ReportOptions) HasSampleFilename() bool`

HasSampleFilename returns a boolean if a field has been set.

### GetStartCommand

`func (o *ReportOptions) GetStartCommand() string`

GetStartCommand returns the StartCommand field if non-nil, zero value otherwise.

### GetStartCommandOk

`func (o *ReportOptions) GetStartCommandOk() (*string, bool)`

GetStartCommandOk returns a tuple with the StartCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCommand

`func (o *ReportOptions) SetStartCommand(v string)`

SetStartCommand sets StartCommand field to given value.

### HasStartCommand

`func (o *ReportOptions) HasStartCommand() bool`

HasStartCommand returns a boolean if a field has been set.

### GetStartMethod

`func (o *ReportOptions) GetStartMethod() string`

GetStartMethod returns the StartMethod field if non-nil, zero value otherwise.

### GetStartMethodOk

`func (o *ReportOptions) GetStartMethodOk() (*string, bool)`

GetStartMethodOk returns a tuple with the StartMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMethod

`func (o *ReportOptions) SetStartMethod(v string)`

SetStartMethod sets StartMethod field to given value.

### HasStartMethod

`func (o *ReportOptions) HasStartMethod() bool`

HasStartMethod returns a boolean if a field has been set.

### GetTimeout

`func (o *ReportOptions) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ReportOptions) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ReportOptions) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ReportOptions) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


