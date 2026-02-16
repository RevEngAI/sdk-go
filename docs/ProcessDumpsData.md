# ProcessDumpsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Dumps** | [**[]ProcessDump**](ProcessDump.md) |  | 

## Methods

### NewProcessDumpsData

`func NewProcessDumpsData(count int32, dumps []ProcessDump, ) *ProcessDumpsData`

NewProcessDumpsData instantiates a new ProcessDumpsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessDumpsDataWithDefaults

`func NewProcessDumpsDataWithDefaults() *ProcessDumpsData`

NewProcessDumpsDataWithDefaults instantiates a new ProcessDumpsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ProcessDumpsData) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ProcessDumpsData) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ProcessDumpsData) SetCount(v int32)`

SetCount sets Count field to given value.


### GetDumps

`func (o *ProcessDumpsData) GetDumps() []ProcessDump`

GetDumps returns the Dumps field if non-nil, zero value otherwise.

### GetDumpsOk

`func (o *ProcessDumpsData) GetDumpsOk() (*[]ProcessDump, bool)`

GetDumpsOk returns a tuple with the Dumps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDumps

`func (o *ProcessDumpsData) SetDumps(v []ProcessDump)`

SetDumps sets Dumps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


