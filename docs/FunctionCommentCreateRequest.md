# FunctionCommentCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | Comment text content | 
**Context** | Pointer to [**DecompilationCommentContext**](DecompilationCommentContext.md) | Comment context for a function decompilation | [optional] 

## Methods

### NewFunctionCommentCreateRequest

`func NewFunctionCommentCreateRequest(content string, ) *FunctionCommentCreateRequest`

NewFunctionCommentCreateRequest instantiates a new FunctionCommentCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionCommentCreateRequestWithDefaults

`func NewFunctionCommentCreateRequestWithDefaults() *FunctionCommentCreateRequest`

NewFunctionCommentCreateRequestWithDefaults instantiates a new FunctionCommentCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *FunctionCommentCreateRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *FunctionCommentCreateRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *FunctionCommentCreateRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetContext

`func (o *FunctionCommentCreateRequest) GetContext() DecompilationCommentContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *FunctionCommentCreateRequest) GetContextOk() (*DecompilationCommentContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *FunctionCommentCreateRequest) SetContext(v DecompilationCommentContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *FunctionCommentCreateRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


