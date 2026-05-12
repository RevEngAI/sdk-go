# CommentsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**InlineComments** | [**[]InlineComment**](InlineComment.md) | Structured inline comments with line numbers | 
**TaskStatus** | **string** | Task status | 

## Methods

### NewCommentsData

`func NewCommentsData(inlineComments []InlineComment, taskStatus string, ) *CommentsData`

NewCommentsData instantiates a new CommentsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentsDataWithDefaults

`func NewCommentsDataWithDefaults() *CommentsData`

NewCommentsDataWithDefaults instantiates a new CommentsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *CommentsData) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *CommentsData) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *CommentsData) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *CommentsData) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetInlineComments

`func (o *CommentsData) GetInlineComments() []InlineComment`

GetInlineComments returns the InlineComments field if non-nil, zero value otherwise.

### GetInlineCommentsOk

`func (o *CommentsData) GetInlineCommentsOk() (*[]InlineComment, bool)`

GetInlineCommentsOk returns a tuple with the InlineComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineComments

`func (o *CommentsData) SetInlineComments(v []InlineComment)`

SetInlineComments sets InlineComments field to given value.


### SetInlineCommentsNil

`func (o *CommentsData) SetInlineCommentsNil(b bool)`

 SetInlineCommentsNil sets the value for InlineComments to be an explicit nil

### UnsetInlineComments
`func (o *CommentsData) UnsetInlineComments()`

UnsetInlineComments ensures that no value is present for InlineComments, not even an explicit nil
### GetTaskStatus

`func (o *CommentsData) GetTaskStatus() string`

GetTaskStatus returns the TaskStatus field if non-nil, zero value otherwise.

### GetTaskStatusOk

`func (o *CommentsData) GetTaskStatusOk() (*string, bool)`

GetTaskStatusOk returns a tuple with the TaskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStatus

`func (o *CommentsData) SetTaskStatus(v string)`

SetTaskStatus sets TaskStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


