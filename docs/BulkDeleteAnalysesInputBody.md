# BulkDeleteAnalysesInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisIds** | **[]int64** | IDs of the analyses to delete. The caller must own every one, or none are deleted | 

## Methods

### NewBulkDeleteAnalysesInputBody

`func NewBulkDeleteAnalysesInputBody(analysisIds []int64, ) *BulkDeleteAnalysesInputBody`

NewBulkDeleteAnalysesInputBody instantiates a new BulkDeleteAnalysesInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkDeleteAnalysesInputBodyWithDefaults

`func NewBulkDeleteAnalysesInputBodyWithDefaults() *BulkDeleteAnalysesInputBody`

NewBulkDeleteAnalysesInputBodyWithDefaults instantiates a new BulkDeleteAnalysesInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisIds

`func (o *BulkDeleteAnalysesInputBody) GetAnalysisIds() []int64`

GetAnalysisIds returns the AnalysisIds field if non-nil, zero value otherwise.

### GetAnalysisIdsOk

`func (o *BulkDeleteAnalysesInputBody) GetAnalysisIdsOk() (*[]int64, bool)`

GetAnalysisIdsOk returns a tuple with the AnalysisIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisIds

`func (o *BulkDeleteAnalysesInputBody) SetAnalysisIds(v []int64)`

SetAnalysisIds sets AnalysisIds field to given value.


### SetAnalysisIdsNil

`func (o *BulkDeleteAnalysesInputBody) SetAnalysisIdsNil(b bool)`

 SetAnalysisIdsNil sets the value for AnalysisIds to be an explicit nil

### UnsetAnalysisIds
`func (o *BulkDeleteAnalysesInputBody) UnsetAnalysisIds()`

UnsetAnalysisIds ensures that no value is present for AnalysisIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


