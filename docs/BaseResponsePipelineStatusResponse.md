# BaseResponsePipelineStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**NullablePipelineStatusResponse**](PipelineStatusResponse.md) |  | [optional] 
**Errors** | Pointer to [**[]ErrorModel**](ErrorModel.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Meta** | Pointer to [**MetaModel**](MetaModel.md) | Metadata | [optional] 
**Status** | Pointer to **bool** | Response status on whether the request succeeded | [optional] [default to true]

## Methods

### NewBaseResponsePipelineStatusResponse

`func NewBaseResponsePipelineStatusResponse() *BaseResponsePipelineStatusResponse`

NewBaseResponsePipelineStatusResponse instantiates a new BaseResponsePipelineStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseResponsePipelineStatusResponseWithDefaults

`func NewBaseResponsePipelineStatusResponseWithDefaults() *BaseResponsePipelineStatusResponse`

NewBaseResponsePipelineStatusResponseWithDefaults instantiates a new BaseResponsePipelineStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BaseResponsePipelineStatusResponse) GetData() PipelineStatusResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BaseResponsePipelineStatusResponse) GetDataOk() (*PipelineStatusResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BaseResponsePipelineStatusResponse) SetData(v PipelineStatusResponse)`

SetData sets Data field to given value.

### HasData

`func (o *BaseResponsePipelineStatusResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *BaseResponsePipelineStatusResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *BaseResponsePipelineStatusResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetErrors

`func (o *BaseResponsePipelineStatusResponse) GetErrors() []ErrorModel`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BaseResponsePipelineStatusResponse) GetErrorsOk() (*[]ErrorModel, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BaseResponsePipelineStatusResponse) SetErrors(v []ErrorModel)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BaseResponsePipelineStatusResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *BaseResponsePipelineStatusResponse) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *BaseResponsePipelineStatusResponse) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetMessage

`func (o *BaseResponsePipelineStatusResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BaseResponsePipelineStatusResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BaseResponsePipelineStatusResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BaseResponsePipelineStatusResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *BaseResponsePipelineStatusResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *BaseResponsePipelineStatusResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetMeta

`func (o *BaseResponsePipelineStatusResponse) GetMeta() MetaModel`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BaseResponsePipelineStatusResponse) GetMetaOk() (*MetaModel, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BaseResponsePipelineStatusResponse) SetMeta(v MetaModel)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BaseResponsePipelineStatusResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetStatus

`func (o *BaseResponsePipelineStatusResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseResponsePipelineStatusResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseResponsePipelineStatusResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BaseResponsePipelineStatusResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


