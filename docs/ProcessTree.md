# ProcessTree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to [**[]ProcessNode**](ProcessNode.md) |  | [optional] 
**SampleSeqid** | Pointer to **int64** |  | [optional] 

## Methods

### NewProcessTree

`func NewProcessTree() *ProcessTree`

NewProcessTree instantiates a new ProcessTree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessTreeWithDefaults

`func NewProcessTreeWithDefaults() *ProcessTree`

NewProcessTreeWithDefaults instantiates a new ProcessTree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *ProcessTree) GetNodes() []ProcessNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ProcessTree) GetNodesOk() (*[]ProcessNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ProcessTree) SetNodes(v []ProcessNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *ProcessTree) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### SetNodesNil

`func (o *ProcessTree) SetNodesNil(b bool)`

 SetNodesNil sets the value for Nodes to be an explicit nil

### UnsetNodes
`func (o *ProcessTree) UnsetNodes()`

UnsetNodes ensures that no value is present for Nodes, not even an explicit nil
### GetSampleSeqid

`func (o *ProcessTree) GetSampleSeqid() int64`

GetSampleSeqid returns the SampleSeqid field if non-nil, zero value otherwise.

### GetSampleSeqidOk

`func (o *ProcessTree) GetSampleSeqidOk() (*int64, bool)`

GetSampleSeqidOk returns a tuple with the SampleSeqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleSeqid

`func (o *ProcessTree) SetSampleSeqid(v int64)`

SetSampleSeqid sets SampleSeqid field to given value.

### HasSampleSeqid

`func (o *ProcessTree) HasSampleSeqid() bool`

HasSampleSeqid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


