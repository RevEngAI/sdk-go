# FunctionBlockDestinationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationBlockId** | **NullableInt32** |  | 
**Flowtype** | **string** | The type of execution flow between chunks | 
**Vaddr** | **string** | The vaddr of the destination where the execution flow continues from | 

## Methods

### NewFunctionBlockDestinationResponse

`func NewFunctionBlockDestinationResponse(destinationBlockId NullableInt32, flowtype string, vaddr string, ) *FunctionBlockDestinationResponse`

NewFunctionBlockDestinationResponse instantiates a new FunctionBlockDestinationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionBlockDestinationResponseWithDefaults

`func NewFunctionBlockDestinationResponseWithDefaults() *FunctionBlockDestinationResponse`

NewFunctionBlockDestinationResponseWithDefaults instantiates a new FunctionBlockDestinationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationBlockId

`func (o *FunctionBlockDestinationResponse) GetDestinationBlockId() int32`

GetDestinationBlockId returns the DestinationBlockId field if non-nil, zero value otherwise.

### GetDestinationBlockIdOk

`func (o *FunctionBlockDestinationResponse) GetDestinationBlockIdOk() (*int32, bool)`

GetDestinationBlockIdOk returns a tuple with the DestinationBlockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationBlockId

`func (o *FunctionBlockDestinationResponse) SetDestinationBlockId(v int32)`

SetDestinationBlockId sets DestinationBlockId field to given value.


### SetDestinationBlockIdNil

`func (o *FunctionBlockDestinationResponse) SetDestinationBlockIdNil(b bool)`

 SetDestinationBlockIdNil sets the value for DestinationBlockId to be an explicit nil

### UnsetDestinationBlockId
`func (o *FunctionBlockDestinationResponse) UnsetDestinationBlockId()`

UnsetDestinationBlockId ensures that no value is present for DestinationBlockId, not even an explicit nil
### GetFlowtype

`func (o *FunctionBlockDestinationResponse) GetFlowtype() string`

GetFlowtype returns the Flowtype field if non-nil, zero value otherwise.

### GetFlowtypeOk

`func (o *FunctionBlockDestinationResponse) GetFlowtypeOk() (*string, bool)`

GetFlowtypeOk returns a tuple with the Flowtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowtype

`func (o *FunctionBlockDestinationResponse) SetFlowtype(v string)`

SetFlowtype sets Flowtype field to given value.


### GetVaddr

`func (o *FunctionBlockDestinationResponse) GetVaddr() string`

GetVaddr returns the Vaddr field if non-nil, zero value otherwise.

### GetVaddrOk

`func (o *FunctionBlockDestinationResponse) GetVaddrOk() (*string, bool)`

GetVaddrOk returns a tuple with the Vaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaddr

`func (o *FunctionBlockDestinationResponse) SetVaddr(v string)`

SetVaddr sets Vaddr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


