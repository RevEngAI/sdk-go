# OrganisationGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Name** | **NullableString** |  | 
**OrganisationGroupId** | **int64** |  | 
**OrganisationId** | **int64** |  | 
**TeamId** | **int64** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewOrganisationGroup

`func NewOrganisationGroup(createdAt time.Time, name NullableString, organisationGroupId int64, organisationId int64, teamId int64, updatedAt time.Time, ) *OrganisationGroup`

NewOrganisationGroup instantiates a new OrganisationGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganisationGroupWithDefaults

`func NewOrganisationGroupWithDefaults() *OrganisationGroup`

NewOrganisationGroupWithDefaults instantiates a new OrganisationGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *OrganisationGroup) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganisationGroup) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganisationGroup) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetName

`func (o *OrganisationGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganisationGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganisationGroup) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *OrganisationGroup) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OrganisationGroup) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrganisationGroupId

`func (o *OrganisationGroup) GetOrganisationGroupId() int64`

GetOrganisationGroupId returns the OrganisationGroupId field if non-nil, zero value otherwise.

### GetOrganisationGroupIdOk

`func (o *OrganisationGroup) GetOrganisationGroupIdOk() (*int64, bool)`

GetOrganisationGroupIdOk returns a tuple with the OrganisationGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisationGroupId

`func (o *OrganisationGroup) SetOrganisationGroupId(v int64)`

SetOrganisationGroupId sets OrganisationGroupId field to given value.


### GetOrganisationId

`func (o *OrganisationGroup) GetOrganisationId() int64`

GetOrganisationId returns the OrganisationId field if non-nil, zero value otherwise.

### GetOrganisationIdOk

`func (o *OrganisationGroup) GetOrganisationIdOk() (*int64, bool)`

GetOrganisationIdOk returns a tuple with the OrganisationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisationId

`func (o *OrganisationGroup) SetOrganisationId(v int64)`

SetOrganisationId sets OrganisationId field to given value.


### GetTeamId

`func (o *OrganisationGroup) GetTeamId() int64`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *OrganisationGroup) GetTeamIdOk() (*int64, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *OrganisationGroup) SetTeamId(v int64)`

SetTeamId sets TeamId field to given value.


### GetUpdatedAt

`func (o *OrganisationGroup) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganisationGroup) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganisationGroup) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


