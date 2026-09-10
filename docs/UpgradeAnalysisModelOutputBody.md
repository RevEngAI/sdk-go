# UpgradeAnalysisModelOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisId** | **int64** | ID of the re-queued analysis. Unchanged — the upgrade moves the existing analysis rather than creating a new one. | 
**BinaryId** | **int64** | ID of the binary the analysis belongs to | 
**ModelId** | **int64** | Model the analysis is now queued against | 

## Methods

### NewUpgradeAnalysisModelOutputBody

`func NewUpgradeAnalysisModelOutputBody(analysisId int64, binaryId int64, modelId int64, ) *UpgradeAnalysisModelOutputBody`

NewUpgradeAnalysisModelOutputBody instantiates a new UpgradeAnalysisModelOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeAnalysisModelOutputBodyWithDefaults

`func NewUpgradeAnalysisModelOutputBodyWithDefaults() *UpgradeAnalysisModelOutputBody`

NewUpgradeAnalysisModelOutputBodyWithDefaults instantiates a new UpgradeAnalysisModelOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisId

`func (o *UpgradeAnalysisModelOutputBody) GetAnalysisId() int64`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *UpgradeAnalysisModelOutputBody) GetAnalysisIdOk() (*int64, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *UpgradeAnalysisModelOutputBody) SetAnalysisId(v int64)`

SetAnalysisId sets AnalysisId field to given value.


### GetBinaryId

`func (o *UpgradeAnalysisModelOutputBody) GetBinaryId() int64`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *UpgradeAnalysisModelOutputBody) GetBinaryIdOk() (*int64, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *UpgradeAnalysisModelOutputBody) SetBinaryId(v int64)`

SetBinaryId sets BinaryId field to given value.


### GetModelId

`func (o *UpgradeAnalysisModelOutputBody) GetModelId() int64`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *UpgradeAnalysisModelOutputBody) GetModelIdOk() (*int64, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *UpgradeAnalysisModelOutputBody) SetModelId(v int64)`

SetModelId sets ModelId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


