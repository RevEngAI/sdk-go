# ModuleLoadEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Modules** | Pointer to **map[string]string** |  | [optional] 
**Pid** | **int64** |  | 
**ProcessName** | Pointer to **string** |  | [optional] 
**ProcessSeqid** | Pointer to **int64** |  | [optional] 

## Methods

### NewModuleLoadEntry

`func NewModuleLoadEntry(pid int64, ) *ModuleLoadEntry`

NewModuleLoadEntry instantiates a new ModuleLoadEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleLoadEntryWithDefaults

`func NewModuleLoadEntryWithDefaults() *ModuleLoadEntry`

NewModuleLoadEntryWithDefaults instantiates a new ModuleLoadEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModules

`func (o *ModuleLoadEntry) GetModules() map[string]string`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *ModuleLoadEntry) GetModulesOk() (*map[string]string, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *ModuleLoadEntry) SetModules(v map[string]string)`

SetModules sets Modules field to given value.

### HasModules

`func (o *ModuleLoadEntry) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetPid

`func (o *ModuleLoadEntry) GetPid() int64`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *ModuleLoadEntry) GetPidOk() (*int64, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *ModuleLoadEntry) SetPid(v int64)`

SetPid sets Pid field to given value.


### GetProcessName

`func (o *ModuleLoadEntry) GetProcessName() string`

GetProcessName returns the ProcessName field if non-nil, zero value otherwise.

### GetProcessNameOk

`func (o *ModuleLoadEntry) GetProcessNameOk() (*string, bool)`

GetProcessNameOk returns a tuple with the ProcessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessName

`func (o *ModuleLoadEntry) SetProcessName(v string)`

SetProcessName sets ProcessName field to given value.

### HasProcessName

`func (o *ModuleLoadEntry) HasProcessName() bool`

HasProcessName returns a boolean if a field has been set.

### GetProcessSeqid

`func (o *ModuleLoadEntry) GetProcessSeqid() int64`

GetProcessSeqid returns the ProcessSeqid field if non-nil, zero value otherwise.

### GetProcessSeqidOk

`func (o *ModuleLoadEntry) GetProcessSeqidOk() (*int64, bool)`

GetProcessSeqidOk returns a tuple with the ProcessSeqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessSeqid

`func (o *ModuleLoadEntry) SetProcessSeqid(v int64)`

SetProcessSeqid sets ProcessSeqid field to given value.

### HasProcessSeqid

`func (o *ModuleLoadEntry) HasProcessSeqid() bool`

HasProcessSeqid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


