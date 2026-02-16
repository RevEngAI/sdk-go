# Process

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pid** | **int32** |  | 
**Procname** | **string** |  | 
**ExecutableName** | **string** |  | 
**Args** | **[]string** |  | 
**TsFrom** | **float32** |  | 
**TsTo** | **NullableFloat32** |  | 
**Children** | **[]interface{}** |  | 

## Methods

### NewProcess

`func NewProcess(pid int32, procname string, executableName string, args []string, tsFrom float32, tsTo NullableFloat32, children []interface{}, ) *Process`

NewProcess instantiates a new Process object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessWithDefaults

`func NewProcessWithDefaults() *Process`

NewProcessWithDefaults instantiates a new Process object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPid

`func (o *Process) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *Process) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *Process) SetPid(v int32)`

SetPid sets Pid field to given value.


### GetProcname

`func (o *Process) GetProcname() string`

GetProcname returns the Procname field if non-nil, zero value otherwise.

### GetProcnameOk

`func (o *Process) GetProcnameOk() (*string, bool)`

GetProcnameOk returns a tuple with the Procname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcname

`func (o *Process) SetProcname(v string)`

SetProcname sets Procname field to given value.


### GetExecutableName

`func (o *Process) GetExecutableName() string`

GetExecutableName returns the ExecutableName field if non-nil, zero value otherwise.

### GetExecutableNameOk

`func (o *Process) GetExecutableNameOk() (*string, bool)`

GetExecutableNameOk returns a tuple with the ExecutableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutableName

`func (o *Process) SetExecutableName(v string)`

SetExecutableName sets ExecutableName field to given value.


### GetArgs

`func (o *Process) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *Process) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *Process) SetArgs(v []string)`

SetArgs sets Args field to given value.


### GetTsFrom

`func (o *Process) GetTsFrom() float32`

GetTsFrom returns the TsFrom field if non-nil, zero value otherwise.

### GetTsFromOk

`func (o *Process) GetTsFromOk() (*float32, bool)`

GetTsFromOk returns a tuple with the TsFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsFrom

`func (o *Process) SetTsFrom(v float32)`

SetTsFrom sets TsFrom field to given value.


### GetTsTo

`func (o *Process) GetTsTo() float32`

GetTsTo returns the TsTo field if non-nil, zero value otherwise.

### GetTsToOk

`func (o *Process) GetTsToOk() (*float32, bool)`

GetTsToOk returns a tuple with the TsTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsTo

`func (o *Process) SetTsTo(v float32)`

SetTsTo sets TsTo field to given value.


### SetTsToNil

`func (o *Process) SetTsToNil(b bool)`

 SetTsToNil sets the value for TsTo to be an explicit nil

### UnsetTsTo
`func (o *Process) UnsetTsTo()`

UnsetTsTo ensures that no value is present for TsTo, not even an explicit nil
### GetChildren

`func (o *Process) GetChildren() []interface{}`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Process) GetChildrenOk() (*[]interface{}, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Process) SetChildren(v []interface{})`

SetChildren sets Children field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


