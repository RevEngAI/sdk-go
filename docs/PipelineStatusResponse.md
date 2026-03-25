# PipelineStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stages** | [**[]StageStatus**](StageStatus.md) |  | 

## Methods

### NewPipelineStatusResponse

`func NewPipelineStatusResponse(stages []StageStatus, ) *PipelineStatusResponse`

NewPipelineStatusResponse instantiates a new PipelineStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineStatusResponseWithDefaults

`func NewPipelineStatusResponseWithDefaults() *PipelineStatusResponse`

NewPipelineStatusResponseWithDefaults instantiates a new PipelineStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStages

`func (o *PipelineStatusResponse) GetStages() []StageStatus`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *PipelineStatusResponse) GetStagesOk() (*[]StageStatus, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *PipelineStatusResponse) SetStages(v []StageStatus)`

SetStages sets Stages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


