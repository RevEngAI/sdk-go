# OrganisationOwner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**OrganisationId** | **int64** |  | 
**OrganisationOwnerId** | **int64** |  | 
**UserId** | **int64** |  | 

## Methods

### NewOrganisationOwner

`func NewOrganisationOwner(createdAt time.Time, organisationId int64, organisationOwnerId int64, userId int64, ) *OrganisationOwner`

NewOrganisationOwner instantiates a new OrganisationOwner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganisationOwnerWithDefaults

`func NewOrganisationOwnerWithDefaults() *OrganisationOwner`

NewOrganisationOwnerWithDefaults instantiates a new OrganisationOwner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *OrganisationOwner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganisationOwner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganisationOwner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetOrganisationId

`func (o *OrganisationOwner) GetOrganisationId() int64`

GetOrganisationId returns the OrganisationId field if non-nil, zero value otherwise.

### GetOrganisationIdOk

`func (o *OrganisationOwner) GetOrganisationIdOk() (*int64, bool)`

GetOrganisationIdOk returns a tuple with the OrganisationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisationId

`func (o *OrganisationOwner) SetOrganisationId(v int64)`

SetOrganisationId sets OrganisationId field to given value.


### GetOrganisationOwnerId

`func (o *OrganisationOwner) GetOrganisationOwnerId() int64`

GetOrganisationOwnerId returns the OrganisationOwnerId field if non-nil, zero value otherwise.

### GetOrganisationOwnerIdOk

`func (o *OrganisationOwner) GetOrganisationOwnerIdOk() (*int64, bool)`

GetOrganisationOwnerIdOk returns a tuple with the OrganisationOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisationOwnerId

`func (o *OrganisationOwner) SetOrganisationOwnerId(v int64)`

SetOrganisationOwnerId sets OrganisationOwnerId field to given value.


### GetUserId

`func (o *OrganisationOwner) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OrganisationOwner) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OrganisationOwner) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


