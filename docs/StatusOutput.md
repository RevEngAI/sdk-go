# StatusOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisId** | **int32** | The ID corresponding to the checked status | 
**AnalysisStatus** | **string** | The status of the checked analysis | 

## Methods

### NewStatusOutput

`func NewStatusOutput(analysisId int32, analysisStatus string, ) *StatusOutput`

NewStatusOutput instantiates a new StatusOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusOutputWithDefaults

`func NewStatusOutputWithDefaults() *StatusOutput`

NewStatusOutputWithDefaults instantiates a new StatusOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisId

`func (o *StatusOutput) GetAnalysisId() int32`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *StatusOutput) GetAnalysisIdOk() (*int32, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *StatusOutput) SetAnalysisId(v int32)`

SetAnalysisId sets AnalysisId field to given value.


### GetAnalysisStatus

`func (o *StatusOutput) GetAnalysisStatus() string`

GetAnalysisStatus returns the AnalysisStatus field if non-nil, zero value otherwise.

### GetAnalysisStatusOk

`func (o *StatusOutput) GetAnalysisStatusOk() (*string, bool)`

GetAnalysisStatusOk returns a tuple with the AnalysisStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisStatus

`func (o *StatusOutput) SetAnalysisStatus(v string)`

SetAnalysisStatus sets AnalysisStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


