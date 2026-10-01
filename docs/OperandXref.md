# OperandXref

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstructionVaddr** | **int64** | Vaddr of the instruction containing the operand. | 
**PointedVaddr** | **int64** | Address stored in the pointer slot. Resolve this to a function, import or global to name the reference. | 
**TargetVaddr** | **int64** | Vaddr of the pointer slot the operand references. | 

## Methods

### NewOperandXref

`func NewOperandXref(instructionVaddr int64, pointedVaddr int64, targetVaddr int64, ) *OperandXref`

NewOperandXref instantiates a new OperandXref object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperandXrefWithDefaults

`func NewOperandXrefWithDefaults() *OperandXref`

NewOperandXrefWithDefaults instantiates a new OperandXref object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstructionVaddr

`func (o *OperandXref) GetInstructionVaddr() int64`

GetInstructionVaddr returns the InstructionVaddr field if non-nil, zero value otherwise.

### GetInstructionVaddrOk

`func (o *OperandXref) GetInstructionVaddrOk() (*int64, bool)`

GetInstructionVaddrOk returns a tuple with the InstructionVaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructionVaddr

`func (o *OperandXref) SetInstructionVaddr(v int64)`

SetInstructionVaddr sets InstructionVaddr field to given value.


### GetPointedVaddr

`func (o *OperandXref) GetPointedVaddr() int64`

GetPointedVaddr returns the PointedVaddr field if non-nil, zero value otherwise.

### GetPointedVaddrOk

`func (o *OperandXref) GetPointedVaddrOk() (*int64, bool)`

GetPointedVaddrOk returns a tuple with the PointedVaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointedVaddr

`func (o *OperandXref) SetPointedVaddr(v int64)`

SetPointedVaddr sets PointedVaddr field to given value.


### GetTargetVaddr

`func (o *OperandXref) GetTargetVaddr() int64`

GetTargetVaddr returns the TargetVaddr field if non-nil, zero value otherwise.

### GetTargetVaddrOk

`func (o *OperandXref) GetTargetVaddrOk() (*int64, bool)`

GetTargetVaddrOk returns a tuple with the TargetVaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVaddr

`func (o *OperandXref) SetTargetVaddr(v int64)`

SetTargetVaddr sets TargetVaddr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


