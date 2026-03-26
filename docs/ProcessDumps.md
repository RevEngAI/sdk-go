# ProcessDumps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**ProcessDumpsData**](ProcessDumpsData.md) |  | 
**Success** | **bool** |  | 

## Methods

### NewProcessDumps

`func NewProcessDumps(data ProcessDumpsData, success bool, ) *ProcessDumps`

NewProcessDumps instantiates a new ProcessDumps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessDumpsWithDefaults

`func NewProcessDumpsWithDefaults() *ProcessDumps`

NewProcessDumpsWithDefaults instantiates a new ProcessDumps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ProcessDumps) GetData() ProcessDumpsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProcessDumps) GetDataOk() (*ProcessDumpsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProcessDumps) SetData(v ProcessDumpsData)`

SetData sets Data field to given value.


### GetSuccess

`func (o *ProcessDumps) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ProcessDumps) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ProcessDumps) SetSuccess(v bool)`

SetSuccess sets Success field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


