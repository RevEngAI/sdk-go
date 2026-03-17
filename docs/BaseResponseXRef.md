# BaseResponseXRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** | Response status on whether the request succeeded | [optional] [default to true]
**Data** | Pointer to [**NullableXRef**](XRef.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Errors** | Pointer to [**[]ErrorModel**](ErrorModel.md) |  | [optional] 
**Meta** | Pointer to [**MetaModel**](MetaModel.md) | Metadata | [optional] 

## Methods

### NewBaseResponseXRef

`func NewBaseResponseXRef() *BaseResponseXRef`

NewBaseResponseXRef instantiates a new BaseResponseXRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseResponseXRefWithDefaults

`func NewBaseResponseXRefWithDefaults() *BaseResponseXRef`

NewBaseResponseXRefWithDefaults instantiates a new BaseResponseXRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BaseResponseXRef) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseResponseXRef) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseResponseXRef) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BaseResponseXRef) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *BaseResponseXRef) GetData() XRef`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BaseResponseXRef) GetDataOk() (*XRef, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BaseResponseXRef) SetData(v XRef)`

SetData sets Data field to given value.

### HasData

`func (o *BaseResponseXRef) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *BaseResponseXRef) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *BaseResponseXRef) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMessage

`func (o *BaseResponseXRef) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BaseResponseXRef) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BaseResponseXRef) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BaseResponseXRef) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *BaseResponseXRef) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *BaseResponseXRef) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetErrors

`func (o *BaseResponseXRef) GetErrors() []ErrorModel`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BaseResponseXRef) GetErrorsOk() (*[]ErrorModel, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BaseResponseXRef) SetErrors(v []ErrorModel)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BaseResponseXRef) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *BaseResponseXRef) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *BaseResponseXRef) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetMeta

`func (o *BaseResponseXRef) GetMeta() MetaModel`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BaseResponseXRef) GetMetaOk() (*MetaModel, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BaseResponseXRef) SetMeta(v MetaModel)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BaseResponseXRef) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


