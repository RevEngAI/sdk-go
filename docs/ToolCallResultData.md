# ToolCallResultData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsError** | Pointer to **bool** |  | [optional] 
**Result** | Pointer to **string** |  | [optional] 
**ToolCallId** | Pointer to **string** |  | [optional] 
**ToolName** | Pointer to **string** |  | [optional] 
**Updated** | Pointer to [**[]EntityUpdate**](EntityUpdate.md) |  | [optional] 

## Methods

### NewToolCallResultData

`func NewToolCallResultData() *ToolCallResultData`

NewToolCallResultData instantiates a new ToolCallResultData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolCallResultDataWithDefaults

`func NewToolCallResultDataWithDefaults() *ToolCallResultData`

NewToolCallResultDataWithDefaults instantiates a new ToolCallResultData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsError

`func (o *ToolCallResultData) GetIsError() bool`

GetIsError returns the IsError field if non-nil, zero value otherwise.

### GetIsErrorOk

`func (o *ToolCallResultData) GetIsErrorOk() (*bool, bool)`

GetIsErrorOk returns a tuple with the IsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsError

`func (o *ToolCallResultData) SetIsError(v bool)`

SetIsError sets IsError field to given value.

### HasIsError

`func (o *ToolCallResultData) HasIsError() bool`

HasIsError returns a boolean if a field has been set.

### GetResult

`func (o *ToolCallResultData) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ToolCallResultData) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ToolCallResultData) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *ToolCallResultData) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetToolCallId

`func (o *ToolCallResultData) GetToolCallId() string`

GetToolCallId returns the ToolCallId field if non-nil, zero value otherwise.

### GetToolCallIdOk

`func (o *ToolCallResultData) GetToolCallIdOk() (*string, bool)`

GetToolCallIdOk returns a tuple with the ToolCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCallId

`func (o *ToolCallResultData) SetToolCallId(v string)`

SetToolCallId sets ToolCallId field to given value.

### HasToolCallId

`func (o *ToolCallResultData) HasToolCallId() bool`

HasToolCallId returns a boolean if a field has been set.

### GetToolName

`func (o *ToolCallResultData) GetToolName() string`

GetToolName returns the ToolName field if non-nil, zero value otherwise.

### GetToolNameOk

`func (o *ToolCallResultData) GetToolNameOk() (*string, bool)`

GetToolNameOk returns a tuple with the ToolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolName

`func (o *ToolCallResultData) SetToolName(v string)`

SetToolName sets ToolName field to given value.

### HasToolName

`func (o *ToolCallResultData) HasToolName() bool`

HasToolName returns a boolean if a field has been set.

### GetUpdated

`func (o *ToolCallResultData) GetUpdated() []EntityUpdate`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ToolCallResultData) GetUpdatedOk() (*[]EntityUpdate, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ToolCallResultData) SetUpdated(v []EntityUpdate)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ToolCallResultData) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *ToolCallResultData) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *ToolCallResultData) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


