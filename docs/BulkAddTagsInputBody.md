# BulkAddTagsInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisIds** | **[]int64** | IDs of the analyses to tag. The caller must own every one, or none are changed | 
**Tags** | **[]string** | Tags to add. A binary that already carries a tag, under any origin, is left as-is for that name | 

## Methods

### NewBulkAddTagsInputBody

`func NewBulkAddTagsInputBody(analysisIds []int64, tags []string, ) *BulkAddTagsInputBody`

NewBulkAddTagsInputBody instantiates a new BulkAddTagsInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkAddTagsInputBodyWithDefaults

`func NewBulkAddTagsInputBodyWithDefaults() *BulkAddTagsInputBody`

NewBulkAddTagsInputBodyWithDefaults instantiates a new BulkAddTagsInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisIds

`func (o *BulkAddTagsInputBody) GetAnalysisIds() []int64`

GetAnalysisIds returns the AnalysisIds field if non-nil, zero value otherwise.

### GetAnalysisIdsOk

`func (o *BulkAddTagsInputBody) GetAnalysisIdsOk() (*[]int64, bool)`

GetAnalysisIdsOk returns a tuple with the AnalysisIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisIds

`func (o *BulkAddTagsInputBody) SetAnalysisIds(v []int64)`

SetAnalysisIds sets AnalysisIds field to given value.


### SetAnalysisIdsNil

`func (o *BulkAddTagsInputBody) SetAnalysisIdsNil(b bool)`

 SetAnalysisIdsNil sets the value for AnalysisIds to be an explicit nil

### UnsetAnalysisIds
`func (o *BulkAddTagsInputBody) UnsetAnalysisIds()`

UnsetAnalysisIds ensures that no value is present for AnalysisIds, not even an explicit nil
### GetTags

`func (o *BulkAddTagsInputBody) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkAddTagsInputBody) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkAddTagsInputBody) SetTags(v []string)`

SetTags sets Tags field to given value.


### SetTagsNil

`func (o *BulkAddTagsInputBody) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *BulkAddTagsInputBody) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


