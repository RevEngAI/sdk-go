# FindingBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewFindingBody

`func NewFindingBody(createdAt time.Time, details map[string]interface{}, findingId int64, findingType string, investigationId int64, severity string, source string, summary string, title string, ) *FindingBody`

NewFindingBody instantiates a new FindingBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingBodyWithDefaults

`func NewFindingBodyWithDefaults() *FindingBody`

NewFindingBodyWithDefaults instantiates a new FindingBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *FindingBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindingBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindingBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDetails

`func (o *FindingBody) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *FindingBody) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *FindingBody) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.


### GetFindingId

`func (o *FindingBody) GetFindingId() int64`

GetFindingId returns the FindingId field if non-nil, zero value otherwise.

### GetFindingIdOk

`func (o *FindingBody) GetFindingIdOk() (*int64, bool)`

GetFindingIdOk returns a tuple with the FindingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingId

`func (o *FindingBody) SetFindingId(v int64)`

SetFindingId sets FindingId field to given value.


### GetFindingType

`func (o *FindingBody) GetFindingType() string`

GetFindingType returns the FindingType field if non-nil, zero value otherwise.

### GetFindingTypeOk

`func (o *FindingBody) GetFindingTypeOk() (*string, bool)`

GetFindingTypeOk returns a tuple with the FindingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingType

`func (o *FindingBody) SetFindingType(v string)`

SetFindingType sets FindingType field to given value.


### GetInvestigationId

`func (o *FindingBody) GetInvestigationId() int64`

GetInvestigationId returns the InvestigationId field if non-nil, zero value otherwise.

### GetInvestigationIdOk

`func (o *FindingBody) GetInvestigationIdOk() (*int64, bool)`

GetInvestigationIdOk returns a tuple with the InvestigationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigationId

`func (o *FindingBody) SetInvestigationId(v int64)`

SetInvestigationId sets InvestigationId field to given value.


### GetSeverity

`func (o *FindingBody) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *FindingBody) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *FindingBody) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetSource

`func (o *FindingBody) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FindingBody) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FindingBody) SetSource(v string)`

SetSource sets Source field to given value.


### GetSummary

`func (o *FindingBody) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *FindingBody) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *FindingBody) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetTitle

`func (o *FindingBody) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FindingBody) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FindingBody) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


