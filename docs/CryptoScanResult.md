# CryptoScanResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisId** | **int64** | Analysis the run was performed against | 
**Findings** | Pointer to [**[]CryptoFinding**](CryptoFinding.md) | Functions with crypto-related evidence, sorted by confidence then evidence count | [optional] 
**TotalFunctions** | **int64** | Functions the run considered | 

## Methods

### NewCryptoScanResult

`func NewCryptoScanResult(analysisId int64, totalFunctions int64, ) *CryptoScanResult`

NewCryptoScanResult instantiates a new CryptoScanResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptoScanResultWithDefaults

`func NewCryptoScanResultWithDefaults() *CryptoScanResult`

NewCryptoScanResultWithDefaults instantiates a new CryptoScanResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisId

`func (o *CryptoScanResult) GetAnalysisId() int64`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *CryptoScanResult) GetAnalysisIdOk() (*int64, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *CryptoScanResult) SetAnalysisId(v int64)`

SetAnalysisId sets AnalysisId field to given value.


### GetFindings

`func (o *CryptoScanResult) GetFindings() []CryptoFinding`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *CryptoScanResult) GetFindingsOk() (*[]CryptoFinding, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *CryptoScanResult) SetFindings(v []CryptoFinding)`

SetFindings sets Findings field to given value.

### HasFindings

`func (o *CryptoScanResult) HasFindings() bool`

HasFindings returns a boolean if a field has been set.

### SetFindingsNil

`func (o *CryptoScanResult) SetFindingsNil(b bool)`

 SetFindingsNil sets the value for Findings to be an explicit nil

### UnsetFindings
`func (o *CryptoScanResult) UnsetFindings()`

UnsetFindings ensures that no value is present for Findings, not even an explicit nil
### GetTotalFunctions

`func (o *CryptoScanResult) GetTotalFunctions() int64`

GetTotalFunctions returns the TotalFunctions field if non-nil, zero value otherwise.

### GetTotalFunctionsOk

`func (o *CryptoScanResult) GetTotalFunctionsOk() (*int64, bool)`

GetTotalFunctionsOk returns a tuple with the TotalFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFunctions

`func (o *CryptoScanResult) SetTotalFunctions(v int64)`

SetTotalFunctions sets TotalFunctions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


