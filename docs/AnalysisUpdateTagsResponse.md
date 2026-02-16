# AnalysisUpdateTagsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | [**[]TagResponse**](TagResponse.md) | The analysis tags after updating | 

## Methods

### NewAnalysisUpdateTagsResponse

`func NewAnalysisUpdateTagsResponse(tags []TagResponse, ) *AnalysisUpdateTagsResponse`

NewAnalysisUpdateTagsResponse instantiates a new AnalysisUpdateTagsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisUpdateTagsResponseWithDefaults

`func NewAnalysisUpdateTagsResponseWithDefaults() *AnalysisUpdateTagsResponse`

NewAnalysisUpdateTagsResponseWithDefaults instantiates a new AnalysisUpdateTagsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *AnalysisUpdateTagsResponse) GetTags() []TagResponse`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AnalysisUpdateTagsResponse) GetTagsOk() (*[]TagResponse, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AnalysisUpdateTagsResponse) SetTags(v []TagResponse)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


