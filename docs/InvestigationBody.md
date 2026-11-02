# InvestigationBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseType** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**InvestigationId** | **int64** |  | 
**Name** | **string** |  | 
**OwnerId** | **int64** |  | 
**OwnerUsername** | **string** |  | 
**Scope** | **string** |  | 
**Severity** | **string** |  | 
**Status** | **string** |  | 
**TeamId** | **int64** | Team the investigation belongs to; 0 when its owner had no team | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewInvestigationBody

`func NewInvestigationBody(caseType string, createdAt time.Time, investigationId int64, name string, ownerId int64, ownerUsername string, scope string, severity string, status string, teamId int64, updatedAt time.Time, ) *InvestigationBody`

NewInvestigationBody instantiates a new InvestigationBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestigationBodyWithDefaults

`func NewInvestigationBodyWithDefaults() *InvestigationBody`

NewInvestigationBodyWithDefaults instantiates a new InvestigationBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaseType

`func (o *InvestigationBody) GetCaseType() string`

GetCaseType returns the CaseType field if non-nil, zero value otherwise.

### GetCaseTypeOk

`func (o *InvestigationBody) GetCaseTypeOk() (*string, bool)`

GetCaseTypeOk returns a tuple with the CaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseType

`func (o *InvestigationBody) SetCaseType(v string)`

SetCaseType sets CaseType field to given value.


### GetCreatedAt

`func (o *InvestigationBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvestigationBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvestigationBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetInvestigationId

`func (o *InvestigationBody) GetInvestigationId() int64`

GetInvestigationId returns the InvestigationId field if non-nil, zero value otherwise.

### GetInvestigationIdOk

`func (o *InvestigationBody) GetInvestigationIdOk() (*int64, bool)`

GetInvestigationIdOk returns a tuple with the InvestigationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigationId

`func (o *InvestigationBody) SetInvestigationId(v int64)`

SetInvestigationId sets InvestigationId field to given value.


### GetName

`func (o *InvestigationBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvestigationBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvestigationBody) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *InvestigationBody) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *InvestigationBody) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *InvestigationBody) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.


### GetOwnerUsername

`func (o *InvestigationBody) GetOwnerUsername() string`

GetOwnerUsername returns the OwnerUsername field if non-nil, zero value otherwise.

### GetOwnerUsernameOk

`func (o *InvestigationBody) GetOwnerUsernameOk() (*string, bool)`

GetOwnerUsernameOk returns a tuple with the OwnerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUsername

`func (o *InvestigationBody) SetOwnerUsername(v string)`

SetOwnerUsername sets OwnerUsername field to given value.


### GetScope

`func (o *InvestigationBody) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *InvestigationBody) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *InvestigationBody) SetScope(v string)`

SetScope sets Scope field to given value.


### GetSeverity

`func (o *InvestigationBody) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *InvestigationBody) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *InvestigationBody) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetStatus

`func (o *InvestigationBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvestigationBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvestigationBody) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTeamId

`func (o *InvestigationBody) GetTeamId() int64`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *InvestigationBody) GetTeamIdOk() (*int64, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *InvestigationBody) SetTeamId(v int64)`

SetTeamId sets TeamId field to given value.


### GetUpdatedAt

`func (o *InvestigationBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InvestigationBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InvestigationBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


