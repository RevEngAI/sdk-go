# BaseResponseXrefResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**NullableXrefResponse**](XrefResponse.md) |  | [optional] 
**Errors** | Pointer to [**[]ErrorModel**](ErrorModel.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Meta** | Pointer to [**MetaModel**](MetaModel.md) | Metadata | [optional] 
**Status** | Pointer to **bool** | Response status on whether the request succeeded | [optional] [default to true]

## Methods

### NewBaseResponseXrefResponse

`func NewBaseResponseXrefResponse() *BaseResponseXrefResponse`

NewBaseResponseXrefResponse instantiates a new BaseResponseXrefResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseResponseXrefResponseWithDefaults

`func NewBaseResponseXrefResponseWithDefaults() *BaseResponseXrefResponse`

NewBaseResponseXrefResponseWithDefaults instantiates a new BaseResponseXrefResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BaseResponseXrefResponse) GetData() XrefResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BaseResponseXrefResponse) GetDataOk() (*XrefResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BaseResponseXrefResponse) SetData(v XrefResponse)`

SetData sets Data field to given value.

### HasData

`func (o *BaseResponseXrefResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *BaseResponseXrefResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *BaseResponseXrefResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetErrors

`func (o *BaseResponseXrefResponse) GetErrors() []ErrorModel`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BaseResponseXrefResponse) GetErrorsOk() (*[]ErrorModel, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BaseResponseXrefResponse) SetErrors(v []ErrorModel)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BaseResponseXrefResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *BaseResponseXrefResponse) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *BaseResponseXrefResponse) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetMessage

`func (o *BaseResponseXrefResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BaseResponseXrefResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BaseResponseXrefResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BaseResponseXrefResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *BaseResponseXrefResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *BaseResponseXrefResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetMeta

`func (o *BaseResponseXrefResponse) GetMeta() MetaModel`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BaseResponseXrefResponse) GetMetaOk() (*MetaModel, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BaseResponseXrefResponse) SetMeta(v MetaModel)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BaseResponseXrefResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetStatus

`func (o *BaseResponseXrefResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseResponseXrefResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseResponseXrefResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BaseResponseXrefResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


