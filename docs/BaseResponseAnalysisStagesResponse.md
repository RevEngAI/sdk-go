# BaseResponseAnalysisStagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** | Response status on whether the request succeeded | [optional] [default to true]
**Data** | Pointer to [**NullableAnalysisStagesResponse**](AnalysisStagesResponse.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Errors** | Pointer to [**[]ErrorModel**](ErrorModel.md) |  | [optional] 
**Meta** | Pointer to [**MetaModel**](MetaModel.md) | Metadata | [optional] 

## Methods

### NewBaseResponseAnalysisStagesResponse

`func NewBaseResponseAnalysisStagesResponse() *BaseResponseAnalysisStagesResponse`

NewBaseResponseAnalysisStagesResponse instantiates a new BaseResponseAnalysisStagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseResponseAnalysisStagesResponseWithDefaults

`func NewBaseResponseAnalysisStagesResponseWithDefaults() *BaseResponseAnalysisStagesResponse`

NewBaseResponseAnalysisStagesResponseWithDefaults instantiates a new BaseResponseAnalysisStagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BaseResponseAnalysisStagesResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseResponseAnalysisStagesResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseResponseAnalysisStagesResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BaseResponseAnalysisStagesResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *BaseResponseAnalysisStagesResponse) GetData() AnalysisStagesResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BaseResponseAnalysisStagesResponse) GetDataOk() (*AnalysisStagesResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BaseResponseAnalysisStagesResponse) SetData(v AnalysisStagesResponse)`

SetData sets Data field to given value.

### HasData

`func (o *BaseResponseAnalysisStagesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *BaseResponseAnalysisStagesResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *BaseResponseAnalysisStagesResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMessage

`func (o *BaseResponseAnalysisStagesResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BaseResponseAnalysisStagesResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BaseResponseAnalysisStagesResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BaseResponseAnalysisStagesResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *BaseResponseAnalysisStagesResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *BaseResponseAnalysisStagesResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetErrors

`func (o *BaseResponseAnalysisStagesResponse) GetErrors() []ErrorModel`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BaseResponseAnalysisStagesResponse) GetErrorsOk() (*[]ErrorModel, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BaseResponseAnalysisStagesResponse) SetErrors(v []ErrorModel)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BaseResponseAnalysisStagesResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *BaseResponseAnalysisStagesResponse) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *BaseResponseAnalysisStagesResponse) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetMeta

`func (o *BaseResponseAnalysisStagesResponse) GetMeta() MetaModel`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BaseResponseAnalysisStagesResponse) GetMetaOk() (*MetaModel, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BaseResponseAnalysisStagesResponse) SetMeta(v MetaModel)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BaseResponseAnalysisStagesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


