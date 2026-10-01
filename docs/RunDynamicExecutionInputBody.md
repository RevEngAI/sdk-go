# RunDynamicExecutionInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveEntryPath** | Pointer to **string** | Relative path of the entry inside the archive to execute | [optional] 
**ArchivePassword** | Pointer to **string** | Password for an encrypted archive | [optional] 
**ArchiveSha256Hash** | Pointer to **string** | SHA-256 of the archive object to send to the sandbox instead of the analysed binary | [optional] 
**CommandLineArgs** | Pointer to **string** | Command-line arguments passed to the sample when the sandbox launches it | [optional] 
**StartMethod** | Pointer to **string** | How the sandbox launches the sample. Defaults to the sandbox&#39;s standard behaviour when omitted. | [optional] 
**Timeout** | Pointer to **int64** | Maximum sandbox execution time in seconds | [optional] [default to 120]

## Methods

### NewRunDynamicExecutionInputBody

`func NewRunDynamicExecutionInputBody() *RunDynamicExecutionInputBody`

NewRunDynamicExecutionInputBody instantiates a new RunDynamicExecutionInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunDynamicExecutionInputBodyWithDefaults

`func NewRunDynamicExecutionInputBodyWithDefaults() *RunDynamicExecutionInputBody`

NewRunDynamicExecutionInputBodyWithDefaults instantiates a new RunDynamicExecutionInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveEntryPath

`func (o *RunDynamicExecutionInputBody) GetArchiveEntryPath() string`

GetArchiveEntryPath returns the ArchiveEntryPath field if non-nil, zero value otherwise.

### GetArchiveEntryPathOk

`func (o *RunDynamicExecutionInputBody) GetArchiveEntryPathOk() (*string, bool)`

GetArchiveEntryPathOk returns a tuple with the ArchiveEntryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveEntryPath

`func (o *RunDynamicExecutionInputBody) SetArchiveEntryPath(v string)`

SetArchiveEntryPath sets ArchiveEntryPath field to given value.

### HasArchiveEntryPath

`func (o *RunDynamicExecutionInputBody) HasArchiveEntryPath() bool`

HasArchiveEntryPath returns a boolean if a field has been set.

### GetArchivePassword

`func (o *RunDynamicExecutionInputBody) GetArchivePassword() string`

GetArchivePassword returns the ArchivePassword field if non-nil, zero value otherwise.

### GetArchivePasswordOk

`func (o *RunDynamicExecutionInputBody) GetArchivePasswordOk() (*string, bool)`

GetArchivePasswordOk returns a tuple with the ArchivePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivePassword

`func (o *RunDynamicExecutionInputBody) SetArchivePassword(v string)`

SetArchivePassword sets ArchivePassword field to given value.

### HasArchivePassword

`func (o *RunDynamicExecutionInputBody) HasArchivePassword() bool`

HasArchivePassword returns a boolean if a field has been set.

### GetArchiveSha256Hash

`func (o *RunDynamicExecutionInputBody) GetArchiveSha256Hash() string`

GetArchiveSha256Hash returns the ArchiveSha256Hash field if non-nil, zero value otherwise.

### GetArchiveSha256HashOk

`func (o *RunDynamicExecutionInputBody) GetArchiveSha256HashOk() (*string, bool)`

GetArchiveSha256HashOk returns a tuple with the ArchiveSha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveSha256Hash

`func (o *RunDynamicExecutionInputBody) SetArchiveSha256Hash(v string)`

SetArchiveSha256Hash sets ArchiveSha256Hash field to given value.

### HasArchiveSha256Hash

`func (o *RunDynamicExecutionInputBody) HasArchiveSha256Hash() bool`

HasArchiveSha256Hash returns a boolean if a field has been set.

### GetCommandLineArgs

`func (o *RunDynamicExecutionInputBody) GetCommandLineArgs() string`

GetCommandLineArgs returns the CommandLineArgs field if non-nil, zero value otherwise.

### GetCommandLineArgsOk

`func (o *RunDynamicExecutionInputBody) GetCommandLineArgsOk() (*string, bool)`

GetCommandLineArgsOk returns a tuple with the CommandLineArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineArgs

`func (o *RunDynamicExecutionInputBody) SetCommandLineArgs(v string)`

SetCommandLineArgs sets CommandLineArgs field to given value.

### HasCommandLineArgs

`func (o *RunDynamicExecutionInputBody) HasCommandLineArgs() bool`

HasCommandLineArgs returns a boolean if a field has been set.

### GetStartMethod

`func (o *RunDynamicExecutionInputBody) GetStartMethod() string`

GetStartMethod returns the StartMethod field if non-nil, zero value otherwise.

### GetStartMethodOk

`func (o *RunDynamicExecutionInputBody) GetStartMethodOk() (*string, bool)`

GetStartMethodOk returns a tuple with the StartMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMethod

`func (o *RunDynamicExecutionInputBody) SetStartMethod(v string)`

SetStartMethod sets StartMethod field to given value.

### HasStartMethod

`func (o *RunDynamicExecutionInputBody) HasStartMethod() bool`

HasStartMethod returns a boolean if a field has been set.

### GetTimeout

`func (o *RunDynamicExecutionInputBody) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *RunDynamicExecutionInputBody) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *RunDynamicExecutionInputBody) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *RunDynamicExecutionInputBody) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


