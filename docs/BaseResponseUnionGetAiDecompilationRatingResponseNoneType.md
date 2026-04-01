# BaseResponseUnionGetAiDecompilationRatingResponseNoneType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** | Response status on whether the request succeeded | [optional] [default to true]
**Data** | Pointer to [**NullableGetAiDecompilationRatingResponse**](GetAiDecompilationRatingResponse.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Errors** | Pointer to [**[]ErrorModel**](ErrorModel.md) |  | [optional] 
**Meta** | Pointer to [**MetaModel**](MetaModel.md) | Metadata | [optional] 

## Methods

### NewBaseResponseUnionGetAiDecompilationRatingResponseNoneType

`func NewBaseResponseUnionGetAiDecompilationRatingResponseNoneType() *BaseResponseUnionGetAiDecompilationRatingResponseNoneType`

NewBaseResponseUnionGetAiDecompilationRatingResponseNoneType instantiates a new BaseResponseUnionGetAiDecompilationRatingResponseNoneType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseResponseUnionGetAiDecompilationRatingResponseNoneTypeWithDefaults

`func NewBaseResponseUnionGetAiDecompilationRatingResponseNoneTypeWithDefaults() *BaseResponseUnionGetAiDecompilationRatingResponseNoneType`

NewBaseResponseUnionGetAiDecompilationRatingResponseNoneTypeWithDefaults instantiates a new BaseResponseUnionGetAiDecompilationRatingResponseNoneType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BaseResponseUnionGetAiDecompilationRatingResponseNoneType) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseResponseUnionGetAiDecompilationRatingResponseNoneType) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseResponseUnionGetAiDecompilationRatingResponseNoneType) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BaseResponseUnionGetAiDecompilationRatingResponseNoneType) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *BaseResponseUnionGetAiDecompilationRatingResponseNoneType) GetData() GetAiDecompilationRatingResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BaseResponseUnionGetAiDecompilationRatingResponseNoneType) GetDataOk() (*GetAiDecompilationRatingResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BaseResponseUnionGetAiDecompilationRatingResponseNoneType) SetData(v GetAiDecompilationRatingResponse)`

SetData sets Data field to given value.

### HasData

`func (o *BaseResponseUnionGetAiDecompilationRatingResponseNoneType) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *BaseResponseUnionGetAiDecompilationRatingResponseNoneType) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *BaseResponseUnionGetAiDecompilationRatingResponseNoneType) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMessage

`func (o *BaseResponseUnionGetAiDecompilationRatingResponseNoneType) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BaseResponseUnionGetAiDecompilationRatingResponseNoneType) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BaseResponseUnionGetAiDecompilationRatingResponseNoneType) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BaseResponseUnionGetAiDecompilationRatingResponseNoneType) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *BaseResponseUnionGetAiDecompilationRatingResponseNoneType) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *BaseResponseUnionGetAiDecompilationRatingResponseNoneType) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetErrors

`func (o *BaseResponseUnionGetAiDecompilationRatingResponseNoneType) GetErrors() []ErrorModel`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BaseResponseUnionGetAiDecompilationRatingResponseNoneType) GetErrorsOk() (*[]ErrorModel, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BaseResponseUnionGetAiDecompilationRatingResponseNoneType) SetErrors(v []ErrorModel)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BaseResponseUnionGetAiDecompilationRatingResponseNoneType) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *BaseResponseUnionGetAiDecompilationRatingResponseNoneType) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *BaseResponseUnionGetAiDecompilationRatingResponseNoneType) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetMeta

`func (o *BaseResponseUnionGetAiDecompilationRatingResponseNoneType) GetMeta() MetaModel`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BaseResponseUnionGetAiDecompilationRatingResponseNoneType) GetMetaOk() (*MetaModel, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BaseResponseUnionGetAiDecompilationRatingResponseNoneType) SetMeta(v MetaModel)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BaseResponseUnionGetAiDecompilationRatingResponseNoneType) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


