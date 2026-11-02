# InvestigationListItemBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewInvestigationListItemBody

`func NewInvestigationListItemBody(budget BudgetBody, caseType string, createdAt time.Time, investigationId int64, name string, ownerId int64, ownerUsername string, rollup RollupBody, scope string, severity string, status string, teamId int64, updatedAt time.Time, ) *InvestigationListItemBody`

NewInvestigationListItemBody instantiates a new InvestigationListItemBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestigationListItemBodyWithDefaults

`func NewInvestigationListItemBodyWithDefaults() *InvestigationListItemBody`

NewInvestigationListItemBodyWithDefaults instantiates a new InvestigationListItemBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudget

`func (o *InvestigationListItemBody) GetBudget() BudgetBody`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *InvestigationListItemBody) GetBudgetOk() (*BudgetBody, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *InvestigationListItemBody) SetBudget(v BudgetBody)`

SetBudget sets Budget field to given value.


### GetCaseType

`func (o *InvestigationListItemBody) GetCaseType() string`

GetCaseType returns the CaseType field if non-nil, zero value otherwise.

### GetCaseTypeOk

`func (o *InvestigationListItemBody) GetCaseTypeOk() (*string, bool)`

GetCaseTypeOk returns a tuple with the CaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseType

`func (o *InvestigationListItemBody) SetCaseType(v string)`

SetCaseType sets CaseType field to given value.


### GetCreatedAt

`func (o *InvestigationListItemBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvestigationListItemBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvestigationListItemBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetInvestigationId

`func (o *InvestigationListItemBody) GetInvestigationId() int64`

GetInvestigationId returns the InvestigationId field if non-nil, zero value otherwise.

### GetInvestigationIdOk

`func (o *InvestigationListItemBody) GetInvestigationIdOk() (*int64, bool)`

GetInvestigationIdOk returns a tuple with the InvestigationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigationId

`func (o *InvestigationListItemBody) SetInvestigationId(v int64)`

SetInvestigationId sets InvestigationId field to given value.


### GetName

`func (o *InvestigationListItemBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvestigationListItemBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvestigationListItemBody) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *InvestigationListItemBody) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *InvestigationListItemBody) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *InvestigationListItemBody) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.


### GetOwnerUsername

`func (o *InvestigationListItemBody) GetOwnerUsername() string`

GetOwnerUsername returns the OwnerUsername field if non-nil, zero value otherwise.

### GetOwnerUsernameOk

`func (o *InvestigationListItemBody) GetOwnerUsernameOk() (*string, bool)`

GetOwnerUsernameOk returns a tuple with the OwnerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUsername

`func (o *InvestigationListItemBody) SetOwnerUsername(v string)`

SetOwnerUsername sets OwnerUsername field to given value.


### GetRollup

`func (o *InvestigationListItemBody) GetRollup() RollupBody`

GetRollup returns the Rollup field if non-nil, zero value otherwise.

### GetRollupOk

`func (o *InvestigationListItemBody) GetRollupOk() (*RollupBody, bool)`

GetRollupOk returns a tuple with the Rollup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollup

`func (o *InvestigationListItemBody) SetRollup(v RollupBody)`

SetRollup sets Rollup field to given value.


### GetScope

`func (o *InvestigationListItemBody) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *InvestigationListItemBody) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *InvestigationListItemBody) SetScope(v string)`

SetScope sets Scope field to given value.


### GetSeverity

`func (o *InvestigationListItemBody) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *InvestigationListItemBody) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *InvestigationListItemBody) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetStatus

`func (o *InvestigationListItemBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvestigationListItemBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvestigationListItemBody) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTeamId

`func (o *InvestigationListItemBody) GetTeamId() int64`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *InvestigationListItemBody) GetTeamIdOk() (*int64, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *InvestigationListItemBody) SetTeamId(v int64)`

SetTeamId sets TeamId field to given value.


### GetUpdatedAt

`func (o *InvestigationListItemBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InvestigationListItemBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InvestigationListItemBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


