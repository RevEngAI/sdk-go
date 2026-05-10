# ReportInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to [**DrakvufFileMetadata**](DrakvufFileMetadata.md) |  | [optional] 
**Id** | **string** |  | 
**Options** | Pointer to [**ReportOptions**](ReportOptions.md) |  | [optional] 
**OsProfile** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TimeAnalysisFinished** | Pointer to **string** |  | [optional] 
**TimeExecutionStarted** | Pointer to **string** |  | [optional] 
**TimeStarted** | Pointer to **string** |  | [optional] 

## Methods

### NewReportInfo

`func NewReportInfo(id string, ) *ReportInfo`

NewReportInfo instantiates a new ReportInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportInfoWithDefaults

`func NewReportInfoWithDefaults() *ReportInfo`

NewReportInfoWithDefaults instantiates a new ReportInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *ReportInfo) GetFile() DrakvufFileMetadata`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ReportInfo) GetFileOk() (*DrakvufFileMetadata, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ReportInfo) SetFile(v DrakvufFileMetadata)`

SetFile sets File field to given value.

### HasFile

`func (o *ReportInfo) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetId

`func (o *ReportInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReportInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReportInfo) SetId(v string)`

SetId sets Id field to given value.


### GetOptions

`func (o *ReportInfo) GetOptions() ReportOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ReportInfo) GetOptionsOk() (*ReportOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ReportInfo) SetOptions(v ReportOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ReportInfo) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetOsProfile

`func (o *ReportInfo) GetOsProfile() string`

GetOsProfile returns the OsProfile field if non-nil, zero value otherwise.

### GetOsProfileOk

`func (o *ReportInfo) GetOsProfileOk() (*string, bool)`

GetOsProfileOk returns a tuple with the OsProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsProfile

`func (o *ReportInfo) SetOsProfile(v string)`

SetOsProfile sets OsProfile field to given value.

### HasOsProfile

`func (o *ReportInfo) HasOsProfile() bool`

HasOsProfile returns a boolean if a field has been set.

### GetStatus

`func (o *ReportInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReportInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReportInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReportInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeAnalysisFinished

`func (o *ReportInfo) GetTimeAnalysisFinished() string`

GetTimeAnalysisFinished returns the TimeAnalysisFinished field if non-nil, zero value otherwise.

### GetTimeAnalysisFinishedOk

`func (o *ReportInfo) GetTimeAnalysisFinishedOk() (*string, bool)`

GetTimeAnalysisFinishedOk returns a tuple with the TimeAnalysisFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAnalysisFinished

`func (o *ReportInfo) SetTimeAnalysisFinished(v string)`

SetTimeAnalysisFinished sets TimeAnalysisFinished field to given value.

### HasTimeAnalysisFinished

`func (o *ReportInfo) HasTimeAnalysisFinished() bool`

HasTimeAnalysisFinished returns a boolean if a field has been set.

### GetTimeExecutionStarted

`func (o *ReportInfo) GetTimeExecutionStarted() string`

GetTimeExecutionStarted returns the TimeExecutionStarted field if non-nil, zero value otherwise.

### GetTimeExecutionStartedOk

`func (o *ReportInfo) GetTimeExecutionStartedOk() (*string, bool)`

GetTimeExecutionStartedOk returns a tuple with the TimeExecutionStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeExecutionStarted

`func (o *ReportInfo) SetTimeExecutionStarted(v string)`

SetTimeExecutionStarted sets TimeExecutionStarted field to given value.

### HasTimeExecutionStarted

`func (o *ReportInfo) HasTimeExecutionStarted() bool`

HasTimeExecutionStarted returns a boolean if a field has been set.

### GetTimeStarted

`func (o *ReportInfo) GetTimeStarted() string`

GetTimeStarted returns the TimeStarted field if non-nil, zero value otherwise.

### GetTimeStartedOk

`func (o *ReportInfo) GetTimeStartedOk() (*string, bool)`

GetTimeStartedOk returns a tuple with the TimeStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStarted

`func (o *ReportInfo) SetTimeStarted(v string)`

SetTimeStarted sets TimeStarted field to given value.

### HasTimeStarted

`func (o *ReportInfo) HasTimeStarted() bool`

HasTimeStarted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


