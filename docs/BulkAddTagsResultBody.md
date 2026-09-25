# BulkAddTagsResultBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedTags** | **[]string** | Tags actually added; empty when the binary already carried every requested name | 
**AnalysisId** | **int64** | ID of the analysis | 

## Methods

### NewBulkAddTagsResultBody

`func NewBulkAddTagsResultBody(addedTags []string, analysisId int64, ) *BulkAddTagsResultBody`

NewBulkAddTagsResultBody instantiates a new BulkAddTagsResultBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkAddTagsResultBodyWithDefaults

`func NewBulkAddTagsResultBodyWithDefaults() *BulkAddTagsResultBody`

NewBulkAddTagsResultBodyWithDefaults instantiates a new BulkAddTagsResultBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedTags

`func (o *BulkAddTagsResultBody) GetAddedTags() []string`

GetAddedTags returns the AddedTags field if non-nil, zero value otherwise.

### GetAddedTagsOk

`func (o *BulkAddTagsResultBody) GetAddedTagsOk() (*[]string, bool)`

GetAddedTagsOk returns a tuple with the AddedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedTags

`func (o *BulkAddTagsResultBody) SetAddedTags(v []string)`

SetAddedTags sets AddedTags field to given value.


### SetAddedTagsNil

`func (o *BulkAddTagsResultBody) SetAddedTagsNil(b bool)`

 SetAddedTagsNil sets the value for AddedTags to be an explicit nil

### UnsetAddedTags
`func (o *BulkAddTagsResultBody) UnsetAddedTags()`

UnsetAddedTags ensures that no value is present for AddedTags, not even an explicit nil
### GetAnalysisId

`func (o *BulkAddTagsResultBody) GetAnalysisId() int64`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *BulkAddTagsResultBody) GetAnalysisIdOk() (*int64, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *BulkAddTagsResultBody) SetAnalysisId(v int64)`

SetAnalysisId sets AnalysisId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


