# FilesystemScanResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisId** | **int64** | Analysis the run was performed against | 
**Findings** | Pointer to [**[]FilesystemFinding**](FilesystemFinding.md) | Functions with filesystem-related evidence, sorted by whether they modify the filesystem, then confidence, then evidence count | [optional] 
**TotalFunctions** | **int64** | Functions the run considered | 

## Methods

### NewFilesystemScanResult

`func NewFilesystemScanResult(analysisId int64, totalFunctions int64, ) *FilesystemScanResult`

NewFilesystemScanResult instantiates a new FilesystemScanResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesystemScanResultWithDefaults

`func NewFilesystemScanResultWithDefaults() *FilesystemScanResult`

NewFilesystemScanResultWithDefaults instantiates a new FilesystemScanResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisId

`func (o *FilesystemScanResult) GetAnalysisId() int64`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *FilesystemScanResult) GetAnalysisIdOk() (*int64, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *FilesystemScanResult) SetAnalysisId(v int64)`

SetAnalysisId sets AnalysisId field to given value.


### GetFindings

`func (o *FilesystemScanResult) GetFindings() []FilesystemFinding`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *FilesystemScanResult) GetFindingsOk() (*[]FilesystemFinding, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *FilesystemScanResult) SetFindings(v []FilesystemFinding)`

SetFindings sets Findings field to given value.

### HasFindings

`func (o *FilesystemScanResult) HasFindings() bool`

HasFindings returns a boolean if a field has been set.

### SetFindingsNil

`func (o *FilesystemScanResult) SetFindingsNil(b bool)`

 SetFindingsNil sets the value for Findings to be an explicit nil

### UnsetFindings
`func (o *FilesystemScanResult) UnsetFindings()`

UnsetFindings ensures that no value is present for Findings, not even an explicit nil
### GetTotalFunctions

`func (o *FilesystemScanResult) GetTotalFunctions() int64`

GetTotalFunctions returns the TotalFunctions field if non-nil, zero value otherwise.

### GetTotalFunctionsOk

`func (o *FilesystemScanResult) GetTotalFunctionsOk() (*int64, bool)`

GetTotalFunctionsOk returns a tuple with the TotalFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFunctions

`func (o *FilesystemScanResult) SetTotalFunctions(v int64)`

SetTotalFunctions sets TotalFunctions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


