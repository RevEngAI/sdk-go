# AnalysisReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileActivity** | Pointer to [**[]FileActivityEntry**](FileActivityEntry.md) |  | [optional] 
**Info** | [**ReportInfo**](ReportInfo.md) |  | 
**Memdumps** | Pointer to [**[]ProcessMemdumps**](ProcessMemdumps.md) |  | [optional] 
**ModuleLoadAddresses** | Pointer to [**[]ModuleLoadEntry**](ModuleLoadEntry.md) |  | [optional] 
**Mutexes** | Pointer to [**[]MutexEntry**](MutexEntry.md) |  | [optional] 
**NetworkActivity** | Pointer to [**NetworkActivity**](NetworkActivity.md) |  | [optional] 
**ProcessActivity** | Pointer to [**[]ProcessActivityEntry**](ProcessActivityEntry.md) |  | [optional] 
**ProcessTree** | Pointer to [**ProcessTree**](ProcessTree.md) |  | [optional] 
**RegistryOperations** | Pointer to [**[]RegistryOperation**](RegistryOperation.md) |  | [optional] 
**ScheduledTasks** | Pointer to [**[]ScheduledTaskEntry**](ScheduledTaskEntry.md) |  | [optional] 
**Services** | Pointer to [**[]ServiceEntry**](ServiceEntry.md) |  | [optional] 
**Startup** | Pointer to [**StartupInfo**](StartupInfo.md) |  | [optional] 
**ThreatScore** | **int64** |  | 
**Ttps** | Pointer to [**[]Ttp**](Ttp.md) |  | [optional] 

## Methods

### NewAnalysisReport

`func NewAnalysisReport(info ReportInfo, threatScore int64, ) *AnalysisReport`

NewAnalysisReport instantiates a new AnalysisReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisReportWithDefaults

`func NewAnalysisReportWithDefaults() *AnalysisReport`

NewAnalysisReportWithDefaults instantiates a new AnalysisReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileActivity

`func (o *AnalysisReport) GetFileActivity() []FileActivityEntry`

GetFileActivity returns the FileActivity field if non-nil, zero value otherwise.

### GetFileActivityOk

`func (o *AnalysisReport) GetFileActivityOk() (*[]FileActivityEntry, bool)`

GetFileActivityOk returns a tuple with the FileActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileActivity

`func (o *AnalysisReport) SetFileActivity(v []FileActivityEntry)`

SetFileActivity sets FileActivity field to given value.

### HasFileActivity

`func (o *AnalysisReport) HasFileActivity() bool`

HasFileActivity returns a boolean if a field has been set.

### SetFileActivityNil

`func (o *AnalysisReport) SetFileActivityNil(b bool)`

 SetFileActivityNil sets the value for FileActivity to be an explicit nil

### UnsetFileActivity
`func (o *AnalysisReport) UnsetFileActivity()`

UnsetFileActivity ensures that no value is present for FileActivity, not even an explicit nil
### GetInfo

`func (o *AnalysisReport) GetInfo() ReportInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AnalysisReport) GetInfoOk() (*ReportInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AnalysisReport) SetInfo(v ReportInfo)`

SetInfo sets Info field to given value.


### GetMemdumps

`func (o *AnalysisReport) GetMemdumps() []ProcessMemdumps`

GetMemdumps returns the Memdumps field if non-nil, zero value otherwise.

### GetMemdumpsOk

`func (o *AnalysisReport) GetMemdumpsOk() (*[]ProcessMemdumps, bool)`

GetMemdumpsOk returns a tuple with the Memdumps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemdumps

`func (o *AnalysisReport) SetMemdumps(v []ProcessMemdumps)`

SetMemdumps sets Memdumps field to given value.

### HasMemdumps

`func (o *AnalysisReport) HasMemdumps() bool`

HasMemdumps returns a boolean if a field has been set.

### SetMemdumpsNil

`func (o *AnalysisReport) SetMemdumpsNil(b bool)`

 SetMemdumpsNil sets the value for Memdumps to be an explicit nil

### UnsetMemdumps
`func (o *AnalysisReport) UnsetMemdumps()`

UnsetMemdumps ensures that no value is present for Memdumps, not even an explicit nil
### GetModuleLoadAddresses

`func (o *AnalysisReport) GetModuleLoadAddresses() []ModuleLoadEntry`

GetModuleLoadAddresses returns the ModuleLoadAddresses field if non-nil, zero value otherwise.

### GetModuleLoadAddressesOk

`func (o *AnalysisReport) GetModuleLoadAddressesOk() (*[]ModuleLoadEntry, bool)`

GetModuleLoadAddressesOk returns a tuple with the ModuleLoadAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleLoadAddresses

`func (o *AnalysisReport) SetModuleLoadAddresses(v []ModuleLoadEntry)`

SetModuleLoadAddresses sets ModuleLoadAddresses field to given value.

### HasModuleLoadAddresses

`func (o *AnalysisReport) HasModuleLoadAddresses() bool`

HasModuleLoadAddresses returns a boolean if a field has been set.

### SetModuleLoadAddressesNil

`func (o *AnalysisReport) SetModuleLoadAddressesNil(b bool)`

 SetModuleLoadAddressesNil sets the value for ModuleLoadAddresses to be an explicit nil

### UnsetModuleLoadAddresses
`func (o *AnalysisReport) UnsetModuleLoadAddresses()`

UnsetModuleLoadAddresses ensures that no value is present for ModuleLoadAddresses, not even an explicit nil
### GetMutexes

`func (o *AnalysisReport) GetMutexes() []MutexEntry`

GetMutexes returns the Mutexes field if non-nil, zero value otherwise.

### GetMutexesOk

`func (o *AnalysisReport) GetMutexesOk() (*[]MutexEntry, bool)`

GetMutexesOk returns a tuple with the Mutexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutexes

`func (o *AnalysisReport) SetMutexes(v []MutexEntry)`

SetMutexes sets Mutexes field to given value.

### HasMutexes

`func (o *AnalysisReport) HasMutexes() bool`

HasMutexes returns a boolean if a field has been set.

### SetMutexesNil

`func (o *AnalysisReport) SetMutexesNil(b bool)`

 SetMutexesNil sets the value for Mutexes to be an explicit nil

### UnsetMutexes
`func (o *AnalysisReport) UnsetMutexes()`

UnsetMutexes ensures that no value is present for Mutexes, not even an explicit nil
### GetNetworkActivity

`func (o *AnalysisReport) GetNetworkActivity() NetworkActivity`

GetNetworkActivity returns the NetworkActivity field if non-nil, zero value otherwise.

### GetNetworkActivityOk

`func (o *AnalysisReport) GetNetworkActivityOk() (*NetworkActivity, bool)`

GetNetworkActivityOk returns a tuple with the NetworkActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkActivity

`func (o *AnalysisReport) SetNetworkActivity(v NetworkActivity)`

SetNetworkActivity sets NetworkActivity field to given value.

### HasNetworkActivity

`func (o *AnalysisReport) HasNetworkActivity() bool`

HasNetworkActivity returns a boolean if a field has been set.

### GetProcessActivity

`func (o *AnalysisReport) GetProcessActivity() []ProcessActivityEntry`

GetProcessActivity returns the ProcessActivity field if non-nil, zero value otherwise.

### GetProcessActivityOk

`func (o *AnalysisReport) GetProcessActivityOk() (*[]ProcessActivityEntry, bool)`

GetProcessActivityOk returns a tuple with the ProcessActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessActivity

`func (o *AnalysisReport) SetProcessActivity(v []ProcessActivityEntry)`

SetProcessActivity sets ProcessActivity field to given value.

### HasProcessActivity

`func (o *AnalysisReport) HasProcessActivity() bool`

HasProcessActivity returns a boolean if a field has been set.

### SetProcessActivityNil

`func (o *AnalysisReport) SetProcessActivityNil(b bool)`

 SetProcessActivityNil sets the value for ProcessActivity to be an explicit nil

### UnsetProcessActivity
`func (o *AnalysisReport) UnsetProcessActivity()`

UnsetProcessActivity ensures that no value is present for ProcessActivity, not even an explicit nil
### GetProcessTree

`func (o *AnalysisReport) GetProcessTree() ProcessTree`

GetProcessTree returns the ProcessTree field if non-nil, zero value otherwise.

### GetProcessTreeOk

`func (o *AnalysisReport) GetProcessTreeOk() (*ProcessTree, bool)`

GetProcessTreeOk returns a tuple with the ProcessTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessTree

`func (o *AnalysisReport) SetProcessTree(v ProcessTree)`

SetProcessTree sets ProcessTree field to given value.

### HasProcessTree

`func (o *AnalysisReport) HasProcessTree() bool`

HasProcessTree returns a boolean if a field has been set.

### GetRegistryOperations

`func (o *AnalysisReport) GetRegistryOperations() []RegistryOperation`

GetRegistryOperations returns the RegistryOperations field if non-nil, zero value otherwise.

### GetRegistryOperationsOk

`func (o *AnalysisReport) GetRegistryOperationsOk() (*[]RegistryOperation, bool)`

GetRegistryOperationsOk returns a tuple with the RegistryOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryOperations

`func (o *AnalysisReport) SetRegistryOperations(v []RegistryOperation)`

SetRegistryOperations sets RegistryOperations field to given value.

### HasRegistryOperations

`func (o *AnalysisReport) HasRegistryOperations() bool`

HasRegistryOperations returns a boolean if a field has been set.

### SetRegistryOperationsNil

`func (o *AnalysisReport) SetRegistryOperationsNil(b bool)`

 SetRegistryOperationsNil sets the value for RegistryOperations to be an explicit nil

### UnsetRegistryOperations
`func (o *AnalysisReport) UnsetRegistryOperations()`

UnsetRegistryOperations ensures that no value is present for RegistryOperations, not even an explicit nil
### GetScheduledTasks

`func (o *AnalysisReport) GetScheduledTasks() []ScheduledTaskEntry`

GetScheduledTasks returns the ScheduledTasks field if non-nil, zero value otherwise.

### GetScheduledTasksOk

`func (o *AnalysisReport) GetScheduledTasksOk() (*[]ScheduledTaskEntry, bool)`

GetScheduledTasksOk returns a tuple with the ScheduledTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledTasks

`func (o *AnalysisReport) SetScheduledTasks(v []ScheduledTaskEntry)`

SetScheduledTasks sets ScheduledTasks field to given value.

### HasScheduledTasks

`func (o *AnalysisReport) HasScheduledTasks() bool`

HasScheduledTasks returns a boolean if a field has been set.

### SetScheduledTasksNil

`func (o *AnalysisReport) SetScheduledTasksNil(b bool)`

 SetScheduledTasksNil sets the value for ScheduledTasks to be an explicit nil

### UnsetScheduledTasks
`func (o *AnalysisReport) UnsetScheduledTasks()`

UnsetScheduledTasks ensures that no value is present for ScheduledTasks, not even an explicit nil
### GetServices

`func (o *AnalysisReport) GetServices() []ServiceEntry`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AnalysisReport) GetServicesOk() (*[]ServiceEntry, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *AnalysisReport) SetServices(v []ServiceEntry)`

SetServices sets Services field to given value.

### HasServices

`func (o *AnalysisReport) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *AnalysisReport) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *AnalysisReport) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil
### GetStartup

`func (o *AnalysisReport) GetStartup() StartupInfo`

GetStartup returns the Startup field if non-nil, zero value otherwise.

### GetStartupOk

`func (o *AnalysisReport) GetStartupOk() (*StartupInfo, bool)`

GetStartupOk returns a tuple with the Startup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartup

`func (o *AnalysisReport) SetStartup(v StartupInfo)`

SetStartup sets Startup field to given value.

### HasStartup

`func (o *AnalysisReport) HasStartup() bool`

HasStartup returns a boolean if a field has been set.

### GetThreatScore

`func (o *AnalysisReport) GetThreatScore() int64`

GetThreatScore returns the ThreatScore field if non-nil, zero value otherwise.

### GetThreatScoreOk

`func (o *AnalysisReport) GetThreatScoreOk() (*int64, bool)`

GetThreatScoreOk returns a tuple with the ThreatScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatScore

`func (o *AnalysisReport) SetThreatScore(v int64)`

SetThreatScore sets ThreatScore field to given value.


### GetTtps

`func (o *AnalysisReport) GetTtps() []Ttp`

GetTtps returns the Ttps field if non-nil, zero value otherwise.

### GetTtpsOk

`func (o *AnalysisReport) GetTtpsOk() (*[]Ttp, bool)`

GetTtpsOk returns a tuple with the Ttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtps

`func (o *AnalysisReport) SetTtps(v []Ttp)`

SetTtps sets Ttps field to given value.

### HasTtps

`func (o *AnalysisReport) HasTtps() bool`

HasTtps returns a boolean if a field has been set.

### SetTtpsNil

`func (o *AnalysisReport) SetTtpsNil(b bool)`

 SetTtpsNil sets the value for Ttps to be an explicit nil

### UnsetTtps
`func (o *AnalysisReport) UnsetTtps()`

UnsetTtps ensures that no value is present for Ttps, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


