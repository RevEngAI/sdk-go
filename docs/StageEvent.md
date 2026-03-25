# StageEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stage** | [**AnalysisStage**](AnalysisStage.md) |  | 
**Status** | [**AnalysisStageStatus**](AnalysisStageStatus.md) |  | 
**Timestamp** | **string** |  | 

## Methods

### NewStageEvent

`func NewStageEvent(stage AnalysisStage, status AnalysisStageStatus, timestamp string, ) *StageEvent`

NewStageEvent instantiates a new StageEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStageEventWithDefaults

`func NewStageEventWithDefaults() *StageEvent`

NewStageEventWithDefaults instantiates a new StageEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStage

`func (o *StageEvent) GetStage() AnalysisStage`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *StageEvent) GetStageOk() (*AnalysisStage, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *StageEvent) SetStage(v AnalysisStage)`

SetStage sets Stage field to given value.


### GetStatus

`func (o *StageEvent) GetStatus() AnalysisStageStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StageEvent) GetStatusOk() (*AnalysisStageStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StageEvent) SetStatus(v AnalysisStageStatus)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *StageEvent) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *StageEvent) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *StageEvent) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


