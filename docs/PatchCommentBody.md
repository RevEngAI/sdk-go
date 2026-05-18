# PatchCommentBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | **string** | Comment text | 
**Line** | **int64** | Line number to set the comment on | 

## Methods

### NewPatchCommentBody

`func NewPatchCommentBody(comment string, line int64, ) *PatchCommentBody`

NewPatchCommentBody instantiates a new PatchCommentBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchCommentBodyWithDefaults

`func NewPatchCommentBodyWithDefaults() *PatchCommentBody`

NewPatchCommentBodyWithDefaults instantiates a new PatchCommentBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *PatchCommentBody) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PatchCommentBody) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PatchCommentBody) SetComment(v string)`

SetComment sets Comment field to given value.


### GetLine

`func (o *PatchCommentBody) GetLine() int64`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *PatchCommentBody) GetLineOk() (*int64, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *PatchCommentBody) SetLine(v int64)`

SetLine sets Line field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


