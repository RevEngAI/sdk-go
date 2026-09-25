# BulkAddTagsOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]BulkAddTagsResultBody**](BulkAddTagsResultBody.md) | One entry per requested analysis, in the order they were given | 

## Methods

### NewBulkAddTagsOutputBody

`func NewBulkAddTagsOutputBody(results []BulkAddTagsResultBody, ) *BulkAddTagsOutputBody`

NewBulkAddTagsOutputBody instantiates a new BulkAddTagsOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkAddTagsOutputBodyWithDefaults

`func NewBulkAddTagsOutputBodyWithDefaults() *BulkAddTagsOutputBody`

NewBulkAddTagsOutputBodyWithDefaults instantiates a new BulkAddTagsOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *BulkAddTagsOutputBody) GetResults() []BulkAddTagsResultBody`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BulkAddTagsOutputBody) GetResultsOk() (*[]BulkAddTagsResultBody, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BulkAddTagsOutputBody) SetResults(v []BulkAddTagsResultBody)`

SetResults sets Results field to given value.


### SetResultsNil

`func (o *BulkAddTagsOutputBody) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *BulkAddTagsOutputBody) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


