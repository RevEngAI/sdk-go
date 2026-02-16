# ProcessTree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | [**[]Process**](Process.md) |  | 

## Methods

### NewProcessTree

`func NewProcessTree(success bool, data []Process, ) *ProcessTree`

NewProcessTree instantiates a new ProcessTree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessTreeWithDefaults

`func NewProcessTreeWithDefaults() *ProcessTree`

NewProcessTreeWithDefaults instantiates a new ProcessTree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ProcessTree) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ProcessTree) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ProcessTree) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *ProcessTree) GetData() []Process`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProcessTree) GetDataOk() (*[]Process, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProcessTree) SetData(v []Process)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


