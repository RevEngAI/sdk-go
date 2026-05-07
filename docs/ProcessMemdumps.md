# ProcessMemdumps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dumps** | Pointer to [**[]MemdumpEntry**](MemdumpEntry.md) |  | [optional] 
**ProcessSeqid** | **int64** |  | 

## Methods

### NewProcessMemdumps

`func NewProcessMemdumps(processSeqid int64, ) *ProcessMemdumps`

NewProcessMemdumps instantiates a new ProcessMemdumps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessMemdumpsWithDefaults

`func NewProcessMemdumpsWithDefaults() *ProcessMemdumps`

NewProcessMemdumpsWithDefaults instantiates a new ProcessMemdumps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDumps

`func (o *ProcessMemdumps) GetDumps() []MemdumpEntry`

GetDumps returns the Dumps field if non-nil, zero value otherwise.

### GetDumpsOk

`func (o *ProcessMemdumps) GetDumpsOk() (*[]MemdumpEntry, bool)`

GetDumpsOk returns a tuple with the Dumps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDumps

`func (o *ProcessMemdumps) SetDumps(v []MemdumpEntry)`

SetDumps sets Dumps field to given value.

### HasDumps

`func (o *ProcessMemdumps) HasDumps() bool`

HasDumps returns a boolean if a field has been set.

### SetDumpsNil

`func (o *ProcessMemdumps) SetDumpsNil(b bool)`

 SetDumpsNil sets the value for Dumps to be an explicit nil

### UnsetDumps
`func (o *ProcessMemdumps) UnsetDumps()`

UnsetDumps ensures that no value is present for Dumps, not even an explicit nil
### GetProcessSeqid

`func (o *ProcessMemdumps) GetProcessSeqid() int64`

GetProcessSeqid returns the ProcessSeqid field if non-nil, zero value otherwise.

### GetProcessSeqidOk

`func (o *ProcessMemdumps) GetProcessSeqidOk() (*int64, bool)`

GetProcessSeqidOk returns a tuple with the ProcessSeqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessSeqid

`func (o *ProcessMemdumps) SetProcessSeqid(v int64)`

SetProcessSeqid sets ProcessSeqid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


