# TagSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagId** | **int32** | The ID of the tag | 
**Tag** | **string** | The name of the tag | 

## Methods

### NewTagSearchResult

`func NewTagSearchResult(tagId int32, tag string, ) *TagSearchResult`

NewTagSearchResult instantiates a new TagSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagSearchResultWithDefaults

`func NewTagSearchResultWithDefaults() *TagSearchResult`

NewTagSearchResultWithDefaults instantiates a new TagSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagId

`func (o *TagSearchResult) GetTagId() int32`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *TagSearchResult) GetTagIdOk() (*int32, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *TagSearchResult) SetTagId(v int32)`

SetTagId sets TagId field to given value.


### GetTag

`func (o *TagSearchResult) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *TagSearchResult) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *TagSearchResult) SetTag(v string)`

SetTag sets Tag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


