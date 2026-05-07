# ProcessNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **[]string** |  | [optional] 
**Attributed** | **bool** |  | 
**Children** | Pointer to [**[]ProcessNode**](ProcessNode.md) |  | [optional] 
**ExitCode** | Pointer to **int64** |  | [optional] 
**ExitCodeStr** | Pointer to **string** |  | [optional] 
**ExitedAt** | Pointer to **float64** |  | [optional] 
**KilledBy** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**Pid** | **int64** |  | 
**Seqid** | **int64** |  | 
**StartedAt** | Pointer to **float64** |  | [optional] 

## Methods

### NewProcessNode

`func NewProcessNode(attributed bool, name string, pid int64, seqid int64, ) *ProcessNode`

NewProcessNode instantiates a new ProcessNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessNodeWithDefaults

`func NewProcessNodeWithDefaults() *ProcessNode`

NewProcessNodeWithDefaults instantiates a new ProcessNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *ProcessNode) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *ProcessNode) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *ProcessNode) SetArgs(v []string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *ProcessNode) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### SetArgsNil

`func (o *ProcessNode) SetArgsNil(b bool)`

 SetArgsNil sets the value for Args to be an explicit nil

### UnsetArgs
`func (o *ProcessNode) UnsetArgs()`

UnsetArgs ensures that no value is present for Args, not even an explicit nil
### GetAttributed

`func (o *ProcessNode) GetAttributed() bool`

GetAttributed returns the Attributed field if non-nil, zero value otherwise.

### GetAttributedOk

`func (o *ProcessNode) GetAttributedOk() (*bool, bool)`

GetAttributedOk returns a tuple with the Attributed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributed

`func (o *ProcessNode) SetAttributed(v bool)`

SetAttributed sets Attributed field to given value.


### GetChildren

`func (o *ProcessNode) GetChildren() []ProcessNode`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ProcessNode) GetChildrenOk() (*[]ProcessNode, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ProcessNode) SetChildren(v []ProcessNode)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *ProcessNode) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### SetChildrenNil

`func (o *ProcessNode) SetChildrenNil(b bool)`

 SetChildrenNil sets the value for Children to be an explicit nil

### UnsetChildren
`func (o *ProcessNode) UnsetChildren()`

UnsetChildren ensures that no value is present for Children, not even an explicit nil
### GetExitCode

`func (o *ProcessNode) GetExitCode() int64`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *ProcessNode) GetExitCodeOk() (*int64, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *ProcessNode) SetExitCode(v int64)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *ProcessNode) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetExitCodeStr

`func (o *ProcessNode) GetExitCodeStr() string`

GetExitCodeStr returns the ExitCodeStr field if non-nil, zero value otherwise.

### GetExitCodeStrOk

`func (o *ProcessNode) GetExitCodeStrOk() (*string, bool)`

GetExitCodeStrOk returns a tuple with the ExitCodeStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCodeStr

`func (o *ProcessNode) SetExitCodeStr(v string)`

SetExitCodeStr sets ExitCodeStr field to given value.

### HasExitCodeStr

`func (o *ProcessNode) HasExitCodeStr() bool`

HasExitCodeStr returns a boolean if a field has been set.

### GetExitedAt

`func (o *ProcessNode) GetExitedAt() float64`

GetExitedAt returns the ExitedAt field if non-nil, zero value otherwise.

### GetExitedAtOk

`func (o *ProcessNode) GetExitedAtOk() (*float64, bool)`

GetExitedAtOk returns a tuple with the ExitedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitedAt

`func (o *ProcessNode) SetExitedAt(v float64)`

SetExitedAt sets ExitedAt field to given value.

### HasExitedAt

`func (o *ProcessNode) HasExitedAt() bool`

HasExitedAt returns a boolean if a field has been set.

### GetKilledBy

`func (o *ProcessNode) GetKilledBy() int64`

GetKilledBy returns the KilledBy field if non-nil, zero value otherwise.

### GetKilledByOk

`func (o *ProcessNode) GetKilledByOk() (*int64, bool)`

GetKilledByOk returns a tuple with the KilledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKilledBy

`func (o *ProcessNode) SetKilledBy(v int64)`

SetKilledBy sets KilledBy field to given value.

### HasKilledBy

`func (o *ProcessNode) HasKilledBy() bool`

HasKilledBy returns a boolean if a field has been set.

### GetName

`func (o *ProcessNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProcessNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProcessNode) SetName(v string)`

SetName sets Name field to given value.


### GetPid

`func (o *ProcessNode) GetPid() int64`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *ProcessNode) GetPidOk() (*int64, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *ProcessNode) SetPid(v int64)`

SetPid sets Pid field to given value.


### GetSeqid

`func (o *ProcessNode) GetSeqid() int64`

GetSeqid returns the Seqid field if non-nil, zero value otherwise.

### GetSeqidOk

`func (o *ProcessNode) GetSeqidOk() (*int64, bool)`

GetSeqidOk returns a tuple with the Seqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqid

`func (o *ProcessNode) SetSeqid(v int64)`

SetSeqid sets Seqid field to given value.


### GetStartedAt

`func (o *ProcessNode) GetStartedAt() float64`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ProcessNode) GetStartedAtOk() (*float64, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ProcessNode) SetStartedAt(v float64)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ProcessNode) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


