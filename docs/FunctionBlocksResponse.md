# FunctionBlocksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocks** | [**[]FunctionBlockResponse**](FunctionBlockResponse.md) | Disassembly is broken into control flow blocks | 
**LocalVariables** | [**[]FunctionLocalVariableResponse**](FunctionLocalVariableResponse.md) | Local variables associated with this function | 
**Params** | [**[]FunctionParamResponse**](FunctionParamResponse.md) | Params associated with this function | 
**OverviewComment** | **NullableString** |  | 

## Methods

### NewFunctionBlocksResponse

`func NewFunctionBlocksResponse(blocks []FunctionBlockResponse, localVariables []FunctionLocalVariableResponse, params []FunctionParamResponse, overviewComment NullableString, ) *FunctionBlocksResponse`

NewFunctionBlocksResponse instantiates a new FunctionBlocksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionBlocksResponseWithDefaults

`func NewFunctionBlocksResponseWithDefaults() *FunctionBlocksResponse`

NewFunctionBlocksResponseWithDefaults instantiates a new FunctionBlocksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocks

`func (o *FunctionBlocksResponse) GetBlocks() []FunctionBlockResponse`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *FunctionBlocksResponse) GetBlocksOk() (*[]FunctionBlockResponse, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *FunctionBlocksResponse) SetBlocks(v []FunctionBlockResponse)`

SetBlocks sets Blocks field to given value.


### GetLocalVariables

`func (o *FunctionBlocksResponse) GetLocalVariables() []FunctionLocalVariableResponse`

GetLocalVariables returns the LocalVariables field if non-nil, zero value otherwise.

### GetLocalVariablesOk

`func (o *FunctionBlocksResponse) GetLocalVariablesOk() (*[]FunctionLocalVariableResponse, bool)`

GetLocalVariablesOk returns a tuple with the LocalVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVariables

`func (o *FunctionBlocksResponse) SetLocalVariables(v []FunctionLocalVariableResponse)`

SetLocalVariables sets LocalVariables field to given value.


### GetParams

`func (o *FunctionBlocksResponse) GetParams() []FunctionParamResponse`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *FunctionBlocksResponse) GetParamsOk() (*[]FunctionParamResponse, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *FunctionBlocksResponse) SetParams(v []FunctionParamResponse)`

SetParams sets Params field to given value.


### GetOverviewComment

`func (o *FunctionBlocksResponse) GetOverviewComment() string`

GetOverviewComment returns the OverviewComment field if non-nil, zero value otherwise.

### GetOverviewCommentOk

`func (o *FunctionBlocksResponse) GetOverviewCommentOk() (*string, bool)`

GetOverviewCommentOk returns a tuple with the OverviewComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverviewComment

`func (o *FunctionBlocksResponse) SetOverviewComment(v string)`

SetOverviewComment sets OverviewComment field to given value.


### SetOverviewCommentNil

`func (o *FunctionBlocksResponse) SetOverviewCommentNil(b bool)`

 SetOverviewCommentNil sets the value for OverviewComment to be an explicit nil

### UnsetOverviewComment
`func (o *FunctionBlocksResponse) UnsetOverviewComment()`

UnsetOverviewComment ensures that no value is present for OverviewComment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


