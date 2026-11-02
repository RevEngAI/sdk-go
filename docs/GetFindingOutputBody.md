# GetFindingOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisIds** | **[]int64** | Analyses the finding was derived from that you are permitted to see. | 
**CreatedAt** | **time.Time** |  | 
**Details** | **map[string]interface{}** | Fields specific to the finding type. An IOC has kind and value, a capability has name, a vulnerability has optional cve, cwe and location, and a cross-file flow has steps. | 
**FindingId** | **int64** |  | 
**FindingType** | **string** |  | 
**InvestigationId** | **int64** |  | 
**Severity** | **string** |  | 
**Source** | **string** | The agent or tool that produced the finding. | 
**Summary** | **string** |  | 
**Title** | **string** |  | 

## Methods

### NewGetFindingOutputBody

`func NewGetFindingOutputBody(analysisIds []int64, createdAt time.Time, details map[string]interface{}, findingId int64, findingType string, investigationId int64, severity string, source string, summary string, title string, ) *GetFindingOutputBody`

NewGetFindingOutputBody instantiates a new GetFindingOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFindingOutputBodyWithDefaults

`func NewGetFindingOutputBodyWithDefaults() *GetFindingOutputBody`

NewGetFindingOutputBodyWithDefaults instantiates a new GetFindingOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisIds

`func (o *GetFindingOutputBody) GetAnalysisIds() []int64`

GetAnalysisIds returns the AnalysisIds field if non-nil, zero value otherwise.

### GetAnalysisIdsOk

`func (o *GetFindingOutputBody) GetAnalysisIdsOk() (*[]int64, bool)`

GetAnalysisIdsOk returns a tuple with the AnalysisIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisIds

`func (o *GetFindingOutputBody) SetAnalysisIds(v []int64)`

SetAnalysisIds sets AnalysisIds field to given value.


### GetCreatedAt

`func (o *GetFindingOutputBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetFindingOutputBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetFindingOutputBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDetails

`func (o *GetFindingOutputBody) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetFindingOutputBody) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetFindingOutputBody) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.


### GetFindingId

`func (o *GetFindingOutputBody) GetFindingId() int64`

GetFindingId returns the FindingId field if non-nil, zero value otherwise.

### GetFindingIdOk

`func (o *GetFindingOutputBody) GetFindingIdOk() (*int64, bool)`

GetFindingIdOk returns a tuple with the FindingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingId

`func (o *GetFindingOutputBody) SetFindingId(v int64)`

SetFindingId sets FindingId field to given value.


### GetFindingType

`func (o *GetFindingOutputBody) GetFindingType() string`

GetFindingType returns the FindingType field if non-nil, zero value otherwise.

### GetFindingTypeOk

`func (o *GetFindingOutputBody) GetFindingTypeOk() (*string, bool)`

GetFindingTypeOk returns a tuple with the FindingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingType

`func (o *GetFindingOutputBody) SetFindingType(v string)`

SetFindingType sets FindingType field to given value.


### GetInvestigationId

`func (o *GetFindingOutputBody) GetInvestigationId() int64`

GetInvestigationId returns the InvestigationId field if non-nil, zero value otherwise.

### GetInvestigationIdOk

`func (o *GetFindingOutputBody) GetInvestigationIdOk() (*int64, bool)`

GetInvestigationIdOk returns a tuple with the InvestigationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigationId

`func (o *GetFindingOutputBody) SetInvestigationId(v int64)`

SetInvestigationId sets InvestigationId field to given value.


### GetSeverity

`func (o *GetFindingOutputBody) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetFindingOutputBody) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetFindingOutputBody) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetSource

`func (o *GetFindingOutputBody) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetFindingOutputBody) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetFindingOutputBody) SetSource(v string)`

SetSource sets Source field to given value.


### GetSummary

`func (o *GetFindingOutputBody) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *GetFindingOutputBody) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *GetFindingOutputBody) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetTitle

`func (o *GetFindingOutputBody) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetFindingOutputBody) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetFindingOutputBody) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


