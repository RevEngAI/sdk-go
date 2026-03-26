# BaseResponseAnalysisBulkAddTagsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**NullableAnalysisBulkAddTagsResponse**](AnalysisBulkAddTagsResponse.md) |  | [optional] 
**Errors** | Pointer to [**[]ErrorModel**](ErrorModel.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Meta** | Pointer to [**MetaModel**](MetaModel.md) | Metadata | [optional] 
**Status** | Pointer to **bool** | Response status on whether the request succeeded | [optional] [default to true]

## Methods

### NewBaseResponseAnalysisBulkAddTagsResponse

`func NewBaseResponseAnalysisBulkAddTagsResponse() *BaseResponseAnalysisBulkAddTagsResponse`

NewBaseResponseAnalysisBulkAddTagsResponse instantiates a new BaseResponseAnalysisBulkAddTagsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseResponseAnalysisBulkAddTagsResponseWithDefaults

`func NewBaseResponseAnalysisBulkAddTagsResponseWithDefaults() *BaseResponseAnalysisBulkAddTagsResponse`

NewBaseResponseAnalysisBulkAddTagsResponseWithDefaults instantiates a new BaseResponseAnalysisBulkAddTagsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BaseResponseAnalysisBulkAddTagsResponse) GetData() AnalysisBulkAddTagsResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BaseResponseAnalysisBulkAddTagsResponse) GetDataOk() (*AnalysisBulkAddTagsResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BaseResponseAnalysisBulkAddTagsResponse) SetData(v AnalysisBulkAddTagsResponse)`

SetData sets Data field to given value.

### HasData

`func (o *BaseResponseAnalysisBulkAddTagsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *BaseResponseAnalysisBulkAddTagsResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *BaseResponseAnalysisBulkAddTagsResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetErrors

`func (o *BaseResponseAnalysisBulkAddTagsResponse) GetErrors() []ErrorModel`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BaseResponseAnalysisBulkAddTagsResponse) GetErrorsOk() (*[]ErrorModel, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BaseResponseAnalysisBulkAddTagsResponse) SetErrors(v []ErrorModel)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BaseResponseAnalysisBulkAddTagsResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *BaseResponseAnalysisBulkAddTagsResponse) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *BaseResponseAnalysisBulkAddTagsResponse) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetMessage

`func (o *BaseResponseAnalysisBulkAddTagsResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BaseResponseAnalysisBulkAddTagsResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BaseResponseAnalysisBulkAddTagsResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BaseResponseAnalysisBulkAddTagsResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *BaseResponseAnalysisBulkAddTagsResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *BaseResponseAnalysisBulkAddTagsResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetMeta

`func (o *BaseResponseAnalysisBulkAddTagsResponse) GetMeta() MetaModel`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BaseResponseAnalysisBulkAddTagsResponse) GetMetaOk() (*MetaModel, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BaseResponseAnalysisBulkAddTagsResponse) SetMeta(v MetaModel)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BaseResponseAnalysisBulkAddTagsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetStatus

`func (o *BaseResponseAnalysisBulkAddTagsResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseResponseAnalysisBulkAddTagsResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseResponseAnalysisBulkAddTagsResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BaseResponseAnalysisBulkAddTagsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


