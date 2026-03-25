# StageStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stage** | [**AnalysisStage**](AnalysisStage.md) |  | 
**Status** | [**PipelineStageStatus**](PipelineStageStatus.md) |  | 
**NumAhead** | **int32** |  | 

## Methods

### NewStageStatus

`func NewStageStatus(stage AnalysisStage, status PipelineStageStatus, numAhead int32, ) *StageStatus`

NewStageStatus instantiates a new StageStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStageStatusWithDefaults

`func NewStageStatusWithDefaults() *StageStatus`

NewStageStatusWithDefaults instantiates a new StageStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStage

`func (o *StageStatus) GetStage() AnalysisStage`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *StageStatus) GetStageOk() (*AnalysisStage, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *StageStatus) SetStage(v AnalysisStage)`

SetStage sets Stage field to given value.


### GetStatus

`func (o *StageStatus) GetStatus() PipelineStageStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StageStatus) GetStatusOk() (*PipelineStageStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StageStatus) SetStatus(v PipelineStageStatus)`

SetStatus sets Status field to given value.


### GetNumAhead

`func (o *StageStatus) GetNumAhead() int32`

GetNumAhead returns the NumAhead field if non-nil, zero value otherwise.

### GetNumAheadOk

`func (o *StageStatus) GetNumAheadOk() (*int32, bool)`

GetNumAheadOk returns a tuple with the NumAhead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAhead

`func (o *StageStatus) SetNumAhead(v int32)`

SetNumAhead sets NumAhead field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


