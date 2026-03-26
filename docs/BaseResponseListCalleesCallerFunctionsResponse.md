# BaseResponseListCalleesCallerFunctionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CalleesCallerFunctionsResponse**](CalleesCallerFunctionsResponse.md) |  | [optional] 
**Errors** | Pointer to [**[]ErrorModel**](ErrorModel.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Meta** | Pointer to [**MetaModel**](MetaModel.md) | Metadata | [optional] 
**Status** | Pointer to **bool** | Response status on whether the request succeeded | [optional] [default to true]

## Methods

### NewBaseResponseListCalleesCallerFunctionsResponse

`func NewBaseResponseListCalleesCallerFunctionsResponse() *BaseResponseListCalleesCallerFunctionsResponse`

NewBaseResponseListCalleesCallerFunctionsResponse instantiates a new BaseResponseListCalleesCallerFunctionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseResponseListCalleesCallerFunctionsResponseWithDefaults

`func NewBaseResponseListCalleesCallerFunctionsResponseWithDefaults() *BaseResponseListCalleesCallerFunctionsResponse`

NewBaseResponseListCalleesCallerFunctionsResponseWithDefaults instantiates a new BaseResponseListCalleesCallerFunctionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BaseResponseListCalleesCallerFunctionsResponse) GetData() []CalleesCallerFunctionsResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BaseResponseListCalleesCallerFunctionsResponse) GetDataOk() (*[]CalleesCallerFunctionsResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BaseResponseListCalleesCallerFunctionsResponse) SetData(v []CalleesCallerFunctionsResponse)`

SetData sets Data field to given value.

### HasData

`func (o *BaseResponseListCalleesCallerFunctionsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *BaseResponseListCalleesCallerFunctionsResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *BaseResponseListCalleesCallerFunctionsResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetErrors

`func (o *BaseResponseListCalleesCallerFunctionsResponse) GetErrors() []ErrorModel`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BaseResponseListCalleesCallerFunctionsResponse) GetErrorsOk() (*[]ErrorModel, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BaseResponseListCalleesCallerFunctionsResponse) SetErrors(v []ErrorModel)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BaseResponseListCalleesCallerFunctionsResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *BaseResponseListCalleesCallerFunctionsResponse) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *BaseResponseListCalleesCallerFunctionsResponse) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetMessage

`func (o *BaseResponseListCalleesCallerFunctionsResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BaseResponseListCalleesCallerFunctionsResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BaseResponseListCalleesCallerFunctionsResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BaseResponseListCalleesCallerFunctionsResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *BaseResponseListCalleesCallerFunctionsResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *BaseResponseListCalleesCallerFunctionsResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetMeta

`func (o *BaseResponseListCalleesCallerFunctionsResponse) GetMeta() MetaModel`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BaseResponseListCalleesCallerFunctionsResponse) GetMetaOk() (*MetaModel, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BaseResponseListCalleesCallerFunctionsResponse) SetMeta(v MetaModel)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BaseResponseListCalleesCallerFunctionsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetStatus

`func (o *BaseResponseListCalleesCallerFunctionsResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseResponseListCalleesCallerFunctionsResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseResponseListCalleesCallerFunctionsResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BaseResponseListCalleesCallerFunctionsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


