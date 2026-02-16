# FunctionTaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**FunctionTaskStatus**](FunctionTaskStatus.md) |  | [optional] [default to FUNCTIONTASKSTATUS_UNINITIALISED]
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFunctionTaskResponse

`func NewFunctionTaskResponse() *FunctionTaskResponse`

NewFunctionTaskResponse instantiates a new FunctionTaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionTaskResponseWithDefaults

`func NewFunctionTaskResponseWithDefaults() *FunctionTaskResponse`

NewFunctionTaskResponseWithDefaults instantiates a new FunctionTaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *FunctionTaskResponse) GetStatus() FunctionTaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FunctionTaskResponse) GetStatusOk() (*FunctionTaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FunctionTaskResponse) SetStatus(v FunctionTaskStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FunctionTaskResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *FunctionTaskResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *FunctionTaskResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *FunctionTaskResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *FunctionTaskResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *FunctionTaskResponse) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *FunctionTaskResponse) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


