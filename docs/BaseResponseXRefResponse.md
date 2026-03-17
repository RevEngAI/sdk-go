# BaseResponseXRefResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** | Response status on whether the request succeeded | [optional] [default to true]
**Data** | Pointer to [**NullableXRefResponse**](XRefResponse.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Errors** | Pointer to [**[]ErrorModel**](ErrorModel.md) |  | [optional] 
**Meta** | Pointer to [**MetaModel**](MetaModel.md) | Metadata | [optional] 

## Methods

### NewBaseResponseXRefResponse

`func NewBaseResponseXRefResponse() *BaseResponseXRefResponse`

NewBaseResponseXRefResponse instantiates a new BaseResponseXRefResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseResponseXRefResponseWithDefaults

`func NewBaseResponseXRefResponseWithDefaults() *BaseResponseXRefResponse`

NewBaseResponseXRefResponseWithDefaults instantiates a new BaseResponseXRefResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BaseResponseXRefResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseResponseXRefResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseResponseXRefResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BaseResponseXRefResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *BaseResponseXRefResponse) GetData() XRefResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BaseResponseXRefResponse) GetDataOk() (*XRefResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BaseResponseXRefResponse) SetData(v XRefResponse)`

SetData sets Data field to given value.

### HasData

`func (o *BaseResponseXRefResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *BaseResponseXRefResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *BaseResponseXRefResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMessage

`func (o *BaseResponseXRefResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BaseResponseXRefResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BaseResponseXRefResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BaseResponseXRefResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *BaseResponseXRefResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *BaseResponseXRefResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetErrors

`func (o *BaseResponseXRefResponse) GetErrors() []ErrorModel`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BaseResponseXRefResponse) GetErrorsOk() (*[]ErrorModel, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BaseResponseXRefResponse) SetErrors(v []ErrorModel)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BaseResponseXRefResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *BaseResponseXRefResponse) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *BaseResponseXRefResponse) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetMeta

`func (o *BaseResponseXRefResponse) GetMeta() MetaModel`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BaseResponseXRefResponse) GetMetaOk() (*MetaModel, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BaseResponseXRefResponse) SetMeta(v MetaModel)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BaseResponseXRefResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


