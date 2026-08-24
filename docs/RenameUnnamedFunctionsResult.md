# RenameUnnamedFunctionsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisId** | **int64** | Analysis the run was performed against | 
**Failed** | **int64** | Functions whose rename attempt errored | 
**Renamed** | **int64** | Functions successfully renamed | 
**Skipped** | **int64** | Functions the agent chose not to rename | 
**Total** | **int64** | Unnamed functions the run considered | 

## Methods

### NewRenameUnnamedFunctionsResult

`func NewRenameUnnamedFunctionsResult(analysisId int64, failed int64, renamed int64, skipped int64, total int64, ) *RenameUnnamedFunctionsResult`

NewRenameUnnamedFunctionsResult instantiates a new RenameUnnamedFunctionsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenameUnnamedFunctionsResultWithDefaults

`func NewRenameUnnamedFunctionsResultWithDefaults() *RenameUnnamedFunctionsResult`

NewRenameUnnamedFunctionsResultWithDefaults instantiates a new RenameUnnamedFunctionsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisId

`func (o *RenameUnnamedFunctionsResult) GetAnalysisId() int64`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *RenameUnnamedFunctionsResult) GetAnalysisIdOk() (*int64, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *RenameUnnamedFunctionsResult) SetAnalysisId(v int64)`

SetAnalysisId sets AnalysisId field to given value.


### GetFailed

`func (o *RenameUnnamedFunctionsResult) GetFailed() int64`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *RenameUnnamedFunctionsResult) GetFailedOk() (*int64, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *RenameUnnamedFunctionsResult) SetFailed(v int64)`

SetFailed sets Failed field to given value.


### GetRenamed

`func (o *RenameUnnamedFunctionsResult) GetRenamed() int64`

GetRenamed returns the Renamed field if non-nil, zero value otherwise.

### GetRenamedOk

`func (o *RenameUnnamedFunctionsResult) GetRenamedOk() (*int64, bool)`

GetRenamedOk returns a tuple with the Renamed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenamed

`func (o *RenameUnnamedFunctionsResult) SetRenamed(v int64)`

SetRenamed sets Renamed field to given value.


### GetSkipped

`func (o *RenameUnnamedFunctionsResult) GetSkipped() int64`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *RenameUnnamedFunctionsResult) GetSkippedOk() (*int64, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *RenameUnnamedFunctionsResult) SetSkipped(v int64)`

SetSkipped sets Skipped field to given value.


### GetTotal

`func (o *RenameUnnamedFunctionsResult) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RenameUnnamedFunctionsResult) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RenameUnnamedFunctionsResult) SetTotal(v int64)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


