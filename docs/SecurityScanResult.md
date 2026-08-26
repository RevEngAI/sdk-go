# SecurityScanResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisId** | **int64** | Analysis the run was performed against | 
**Cancelled** | **bool** | Whether the run was cancelled before it covered every function | 
**Decompiled** | **int64** | Functions successfully decompiled and scanned | 
**Failed** | **int64** | Functions whose decompilation or scan attempt errored | 
**SecurityScan** | Pointer to **map[string]interface{}** | Raw semgrep findings, keyed by the scanner&#39;s own result shape | [optional] 
**Source** | **string** | Decompiler that produced the source scanned | 
**Total** | **int64** | Functions the run considered | 

## Methods

### NewSecurityScanResult

`func NewSecurityScanResult(analysisId int64, cancelled bool, decompiled int64, failed int64, source string, total int64, ) *SecurityScanResult`

NewSecurityScanResult instantiates a new SecurityScanResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityScanResultWithDefaults

`func NewSecurityScanResultWithDefaults() *SecurityScanResult`

NewSecurityScanResultWithDefaults instantiates a new SecurityScanResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisId

`func (o *SecurityScanResult) GetAnalysisId() int64`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *SecurityScanResult) GetAnalysisIdOk() (*int64, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *SecurityScanResult) SetAnalysisId(v int64)`

SetAnalysisId sets AnalysisId field to given value.


### GetCancelled

`func (o *SecurityScanResult) GetCancelled() bool`

GetCancelled returns the Cancelled field if non-nil, zero value otherwise.

### GetCancelledOk

`func (o *SecurityScanResult) GetCancelledOk() (*bool, bool)`

GetCancelledOk returns a tuple with the Cancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelled

`func (o *SecurityScanResult) SetCancelled(v bool)`

SetCancelled sets Cancelled field to given value.


### GetDecompiled

`func (o *SecurityScanResult) GetDecompiled() int64`

GetDecompiled returns the Decompiled field if non-nil, zero value otherwise.

### GetDecompiledOk

`func (o *SecurityScanResult) GetDecompiledOk() (*int64, bool)`

GetDecompiledOk returns a tuple with the Decompiled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecompiled

`func (o *SecurityScanResult) SetDecompiled(v int64)`

SetDecompiled sets Decompiled field to given value.


### GetFailed

`func (o *SecurityScanResult) GetFailed() int64`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *SecurityScanResult) GetFailedOk() (*int64, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *SecurityScanResult) SetFailed(v int64)`

SetFailed sets Failed field to given value.


### GetSecurityScan

`func (o *SecurityScanResult) GetSecurityScan() map[string]interface{}`

GetSecurityScan returns the SecurityScan field if non-nil, zero value otherwise.

### GetSecurityScanOk

`func (o *SecurityScanResult) GetSecurityScanOk() (*map[string]interface{}, bool)`

GetSecurityScanOk returns a tuple with the SecurityScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScan

`func (o *SecurityScanResult) SetSecurityScan(v map[string]interface{})`

SetSecurityScan sets SecurityScan field to given value.

### HasSecurityScan

`func (o *SecurityScanResult) HasSecurityScan() bool`

HasSecurityScan returns a boolean if a field has been set.

### GetSource

`func (o *SecurityScanResult) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SecurityScanResult) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SecurityScanResult) SetSource(v string)`

SetSource sets Source field to given value.


### GetTotal

`func (o *SecurityScanResult) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SecurityScanResult) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SecurityScanResult) SetTotal(v int64)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


