# ProcessActivityEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **[]string** |  | [optional] 
**ChildSeqid** | **int64** |  | 
**Events** | Pointer to [**[]ReportEvent**](ReportEvent.md) |  | [optional] 
**ExitCode** | Pointer to **int64** |  | [optional] 
**ExitCodeStr** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Pid** | **int64** |  | 

## Methods

### NewProcessActivityEntry

`func NewProcessActivityEntry(childSeqid int64, name string, pid int64, ) *ProcessActivityEntry`

NewProcessActivityEntry instantiates a new ProcessActivityEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessActivityEntryWithDefaults

`func NewProcessActivityEntryWithDefaults() *ProcessActivityEntry`

NewProcessActivityEntryWithDefaults instantiates a new ProcessActivityEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *ProcessActivityEntry) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *ProcessActivityEntry) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *ProcessActivityEntry) SetArgs(v []string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *ProcessActivityEntry) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### SetArgsNil

`func (o *ProcessActivityEntry) SetArgsNil(b bool)`

 SetArgsNil sets the value for Args to be an explicit nil

### UnsetArgs
`func (o *ProcessActivityEntry) UnsetArgs()`

UnsetArgs ensures that no value is present for Args, not even an explicit nil
### GetChildSeqid

`func (o *ProcessActivityEntry) GetChildSeqid() int64`

GetChildSeqid returns the ChildSeqid field if non-nil, zero value otherwise.

### GetChildSeqidOk

`func (o *ProcessActivityEntry) GetChildSeqidOk() (*int64, bool)`

GetChildSeqidOk returns a tuple with the ChildSeqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildSeqid

`func (o *ProcessActivityEntry) SetChildSeqid(v int64)`

SetChildSeqid sets ChildSeqid field to given value.


### GetEvents

`func (o *ProcessActivityEntry) GetEvents() []ReportEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ProcessActivityEntry) GetEventsOk() (*[]ReportEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ProcessActivityEntry) SetEvents(v []ReportEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ProcessActivityEntry) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *ProcessActivityEntry) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *ProcessActivityEntry) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetExitCode

`func (o *ProcessActivityEntry) GetExitCode() int64`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *ProcessActivityEntry) GetExitCodeOk() (*int64, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *ProcessActivityEntry) SetExitCode(v int64)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *ProcessActivityEntry) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetExitCodeStr

`func (o *ProcessActivityEntry) GetExitCodeStr() string`

GetExitCodeStr returns the ExitCodeStr field if non-nil, zero value otherwise.

### GetExitCodeStrOk

`func (o *ProcessActivityEntry) GetExitCodeStrOk() (*string, bool)`

GetExitCodeStrOk returns a tuple with the ExitCodeStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCodeStr

`func (o *ProcessActivityEntry) SetExitCodeStr(v string)`

SetExitCodeStr sets ExitCodeStr field to given value.

### HasExitCodeStr

`func (o *ProcessActivityEntry) HasExitCodeStr() bool`

HasExitCodeStr returns a boolean if a field has been set.

### GetName

`func (o *ProcessActivityEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProcessActivityEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProcessActivityEntry) SetName(v string)`

SetName sets Name field to given value.


### GetPid

`func (o *ProcessActivityEntry) GetPid() int64`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *ProcessActivityEntry) GetPidOk() (*int64, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *ProcessActivityEntry) SetPid(v int64)`

SetPid sets Pid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


