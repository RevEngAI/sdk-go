# DynamicExecutionSummaryMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Status of the most recent summary attempt. UNINITIALISED means no summary has been attempted for this analysis. | 

## Methods

### NewDynamicExecutionSummaryMetadata

`func NewDynamicExecutionSummaryMetadata(status string, ) *DynamicExecutionSummaryMetadata`

NewDynamicExecutionSummaryMetadata instantiates a new DynamicExecutionSummaryMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicExecutionSummaryMetadataWithDefaults

`func NewDynamicExecutionSummaryMetadataWithDefaults() *DynamicExecutionSummaryMetadata`

NewDynamicExecutionSummaryMetadataWithDefaults instantiates a new DynamicExecutionSummaryMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DynamicExecutionSummaryMetadata) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DynamicExecutionSummaryMetadata) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DynamicExecutionSummaryMetadata) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


