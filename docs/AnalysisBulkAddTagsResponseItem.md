# AnalysisBulkAddTagsResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisId** | **int32** |  | 
**Error** | Pointer to **NullableString** |  | [optional] 
**Message** | **NullableString** |  | 

## Methods

### NewAnalysisBulkAddTagsResponseItem

`func NewAnalysisBulkAddTagsResponseItem(analysisId int32, message NullableString, ) *AnalysisBulkAddTagsResponseItem`

NewAnalysisBulkAddTagsResponseItem instantiates a new AnalysisBulkAddTagsResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisBulkAddTagsResponseItemWithDefaults

`func NewAnalysisBulkAddTagsResponseItemWithDefaults() *AnalysisBulkAddTagsResponseItem`

NewAnalysisBulkAddTagsResponseItemWithDefaults instantiates a new AnalysisBulkAddTagsResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisId

`func (o *AnalysisBulkAddTagsResponseItem) GetAnalysisId() int32`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *AnalysisBulkAddTagsResponseItem) GetAnalysisIdOk() (*int32, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *AnalysisBulkAddTagsResponseItem) SetAnalysisId(v int32)`

SetAnalysisId sets AnalysisId field to given value.


### GetError

`func (o *AnalysisBulkAddTagsResponseItem) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AnalysisBulkAddTagsResponseItem) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AnalysisBulkAddTagsResponseItem) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AnalysisBulkAddTagsResponseItem) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *AnalysisBulkAddTagsResponseItem) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *AnalysisBulkAddTagsResponseItem) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetMessage

`func (o *AnalysisBulkAddTagsResponseItem) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AnalysisBulkAddTagsResponseItem) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AnalysisBulkAddTagsResponseItem) SetMessage(v string)`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *AnalysisBulkAddTagsResponseItem) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *AnalysisBulkAddTagsResponseItem) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


