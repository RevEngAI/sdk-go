# FunctionCallEdges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Callees** | [**[]CallEdge**](CallEdge.md) |  | 
**Callers** | [**[]CallEdge**](CallEdge.md) |  | 
**FunctionId** | **int64** |  | 

## Methods

### NewFunctionCallEdges

`func NewFunctionCallEdges(callees []CallEdge, callers []CallEdge, functionId int64, ) *FunctionCallEdges`

NewFunctionCallEdges instantiates a new FunctionCallEdges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionCallEdgesWithDefaults

`func NewFunctionCallEdgesWithDefaults() *FunctionCallEdges`

NewFunctionCallEdgesWithDefaults instantiates a new FunctionCallEdges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallees

`func (o *FunctionCallEdges) GetCallees() []CallEdge`

GetCallees returns the Callees field if non-nil, zero value otherwise.

### GetCalleesOk

`func (o *FunctionCallEdges) GetCalleesOk() (*[]CallEdge, bool)`

GetCalleesOk returns a tuple with the Callees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallees

`func (o *FunctionCallEdges) SetCallees(v []CallEdge)`

SetCallees sets Callees field to given value.


### SetCalleesNil

`func (o *FunctionCallEdges) SetCalleesNil(b bool)`

 SetCalleesNil sets the value for Callees to be an explicit nil

### UnsetCallees
`func (o *FunctionCallEdges) UnsetCallees()`

UnsetCallees ensures that no value is present for Callees, not even an explicit nil
### GetCallers

`func (o *FunctionCallEdges) GetCallers() []CallEdge`

GetCallers returns the Callers field if non-nil, zero value otherwise.

### GetCallersOk

`func (o *FunctionCallEdges) GetCallersOk() (*[]CallEdge, bool)`

GetCallersOk returns a tuple with the Callers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallers

`func (o *FunctionCallEdges) SetCallers(v []CallEdge)`

SetCallers sets Callers field to given value.


### SetCallersNil

`func (o *FunctionCallEdges) SetCallersNil(b bool)`

 SetCallersNil sets the value for Callers to be an explicit nil

### UnsetCallers
`func (o *FunctionCallEdges) UnsetCallers()`

UnsetCallers ensures that no value is present for Callers, not even an explicit nil
### GetFunctionId

`func (o *FunctionCallEdges) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *FunctionCallEdges) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *FunctionCallEdges) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


