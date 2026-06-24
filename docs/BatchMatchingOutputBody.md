# BatchMatchingOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PerBinary** | [**[]BatchBinaryMatchResult**](BatchBinaryMatchResult.md) | Per-binary status (order matches the request). | 
**Status** | **string** | Aggregate status across the batch: COMPLETED when every binary is completed, FAILED if any failed, RUNNING/PENDING otherwise. | 

## Methods

### NewBatchMatchingOutputBody

`func NewBatchMatchingOutputBody(perBinary []BatchBinaryMatchResult, status string, ) *BatchMatchingOutputBody`

NewBatchMatchingOutputBody instantiates a new BatchMatchingOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchMatchingOutputBodyWithDefaults

`func NewBatchMatchingOutputBodyWithDefaults() *BatchMatchingOutputBody`

NewBatchMatchingOutputBodyWithDefaults instantiates a new BatchMatchingOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPerBinary

`func (o *BatchMatchingOutputBody) GetPerBinary() []BatchBinaryMatchResult`

GetPerBinary returns the PerBinary field if non-nil, zero value otherwise.

### GetPerBinaryOk

`func (o *BatchMatchingOutputBody) GetPerBinaryOk() (*[]BatchBinaryMatchResult, bool)`

GetPerBinaryOk returns a tuple with the PerBinary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerBinary

`func (o *BatchMatchingOutputBody) SetPerBinary(v []BatchBinaryMatchResult)`

SetPerBinary sets PerBinary field to given value.


### SetPerBinaryNil

`func (o *BatchMatchingOutputBody) SetPerBinaryNil(b bool)`

 SetPerBinaryNil sets the value for PerBinary to be an explicit nil

### UnsetPerBinary
`func (o *BatchMatchingOutputBody) UnsetPerBinary()`

UnsetPerBinary ensures that no value is present for PerBinary, not even an explicit nil
### GetStatus

`func (o *BatchMatchingOutputBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchMatchingOutputBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchMatchingOutputBody) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


