# DynamicExecutionSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneratedAt** | **time.Time** | When the summary was generated | 
**Summary** | **string** | Markdown summary of the run | 

## Methods

### NewDynamicExecutionSummary

`func NewDynamicExecutionSummary(generatedAt time.Time, summary string, ) *DynamicExecutionSummary`

NewDynamicExecutionSummary instantiates a new DynamicExecutionSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicExecutionSummaryWithDefaults

`func NewDynamicExecutionSummaryWithDefaults() *DynamicExecutionSummary`

NewDynamicExecutionSummaryWithDefaults instantiates a new DynamicExecutionSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneratedAt

`func (o *DynamicExecutionSummary) GetGeneratedAt() time.Time`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *DynamicExecutionSummary) GetGeneratedAtOk() (*time.Time, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *DynamicExecutionSummary) SetGeneratedAt(v time.Time)`

SetGeneratedAt sets GeneratedAt field to given value.


### GetSummary

`func (o *DynamicExecutionSummary) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *DynamicExecutionSummary) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *DynamicExecutionSummary) SetSummary(v string)`

SetSummary sets Summary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


