# GetInvestigationOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisIds** | **[]int64** | Analyses in the investigation that you are permitted to see. | 
**Budget** | [**BudgetBody**](BudgetBody.md) |  | 
**CaseType** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**InvestigationId** | **int64** |  | 
**Name** | **string** |  | 
**OwnerId** | **int64** |  | 
**OwnerUsername** | **string** |  | 
**Rollup** | [**RollupBody**](RollupBody.md) |  | 
**Scope** | **string** |  | 
**Severity** | **string** |  | 
**Status** | **string** |  | 
**TeamId** | **int64** | Team the investigation belongs to; 0 when its owner had no team | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewGetInvestigationOutputBody

`func NewGetInvestigationOutputBody(analysisIds []int64, budget BudgetBody, caseType string, createdAt time.Time, investigationId int64, name string, ownerId int64, ownerUsername string, rollup RollupBody, scope string, severity string, status string, teamId int64, updatedAt time.Time, ) *GetInvestigationOutputBody`

NewGetInvestigationOutputBody instantiates a new GetInvestigationOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInvestigationOutputBodyWithDefaults

`func NewGetInvestigationOutputBodyWithDefaults() *GetInvestigationOutputBody`

NewGetInvestigationOutputBodyWithDefaults instantiates a new GetInvestigationOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisIds

`func (o *GetInvestigationOutputBody) GetAnalysisIds() []int64`

GetAnalysisIds returns the AnalysisIds field if non-nil, zero value otherwise.

### GetAnalysisIdsOk

`func (o *GetInvestigationOutputBody) GetAnalysisIdsOk() (*[]int64, bool)`

GetAnalysisIdsOk returns a tuple with the AnalysisIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisIds

`func (o *GetInvestigationOutputBody) SetAnalysisIds(v []int64)`

SetAnalysisIds sets AnalysisIds field to given value.


### GetBudget

`func (o *GetInvestigationOutputBody) GetBudget() BudgetBody`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *GetInvestigationOutputBody) GetBudgetOk() (*BudgetBody, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *GetInvestigationOutputBody) SetBudget(v BudgetBody)`

SetBudget sets Budget field to given value.


### GetCaseType

`func (o *GetInvestigationOutputBody) GetCaseType() string`

GetCaseType returns the CaseType field if non-nil, zero value otherwise.

### GetCaseTypeOk

`func (o *GetInvestigationOutputBody) GetCaseTypeOk() (*string, bool)`

GetCaseTypeOk returns a tuple with the CaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseType

`func (o *GetInvestigationOutputBody) SetCaseType(v string)`

SetCaseType sets CaseType field to given value.


### GetCreatedAt

`func (o *GetInvestigationOutputBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetInvestigationOutputBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetInvestigationOutputBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetInvestigationId

`func (o *GetInvestigationOutputBody) GetInvestigationId() int64`

GetInvestigationId returns the InvestigationId field if non-nil, zero value otherwise.

### GetInvestigationIdOk

`func (o *GetInvestigationOutputBody) GetInvestigationIdOk() (*int64, bool)`

GetInvestigationIdOk returns a tuple with the InvestigationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigationId

`func (o *GetInvestigationOutputBody) SetInvestigationId(v int64)`

SetInvestigationId sets InvestigationId field to given value.


### GetName

`func (o *GetInvestigationOutputBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetInvestigationOutputBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetInvestigationOutputBody) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *GetInvestigationOutputBody) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *GetInvestigationOutputBody) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *GetInvestigationOutputBody) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.


### GetOwnerUsername

`func (o *GetInvestigationOutputBody) GetOwnerUsername() string`

GetOwnerUsername returns the OwnerUsername field if non-nil, zero value otherwise.

### GetOwnerUsernameOk

`func (o *GetInvestigationOutputBody) GetOwnerUsernameOk() (*string, bool)`

GetOwnerUsernameOk returns a tuple with the OwnerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUsername

`func (o *GetInvestigationOutputBody) SetOwnerUsername(v string)`

SetOwnerUsername sets OwnerUsername field to given value.


### GetRollup

`func (o *GetInvestigationOutputBody) GetRollup() RollupBody`

GetRollup returns the Rollup field if non-nil, zero value otherwise.

### GetRollupOk

`func (o *GetInvestigationOutputBody) GetRollupOk() (*RollupBody, bool)`

GetRollupOk returns a tuple with the Rollup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollup

`func (o *GetInvestigationOutputBody) SetRollup(v RollupBody)`

SetRollup sets Rollup field to given value.


### GetScope

`func (o *GetInvestigationOutputBody) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GetInvestigationOutputBody) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GetInvestigationOutputBody) SetScope(v string)`

SetScope sets Scope field to given value.


### GetSeverity

`func (o *GetInvestigationOutputBody) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetInvestigationOutputBody) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetInvestigationOutputBody) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetStatus

`func (o *GetInvestigationOutputBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetInvestigationOutputBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetInvestigationOutputBody) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTeamId

`func (o *GetInvestigationOutputBody) GetTeamId() int64`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *GetInvestigationOutputBody) GetTeamIdOk() (*int64, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *GetInvestigationOutputBody) SetTeamId(v int64)`

SetTeamId sets TeamId field to given value.


### GetUpdatedAt

`func (o *GetInvestigationOutputBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetInvestigationOutputBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetInvestigationOutputBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


