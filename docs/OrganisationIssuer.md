# OrganisationIssuer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | OIDC client ID registered with the identity provider. Used to validate the audience (aud) claim in JWTs. | [optional] 
**CreatedAt** | **time.Time** |  | 
**Enabled** | **bool** |  | 
**IssuerUrl** | **NullableString** |  | 
**JwksUri** | Pointer to **string** | JSON Web Key Set URI discovered from the issuer&#39;s OIDC configuration. Populated automatically during issuer registration. | [optional] 
**OrganisationId** | **int64** |  | 
**OrganisationIssuerId** | **int64** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewOrganisationIssuer

`func NewOrganisationIssuer(createdAt time.Time, enabled bool, issuerUrl NullableString, organisationId int64, organisationIssuerId int64, updatedAt time.Time, ) *OrganisationIssuer`

NewOrganisationIssuer instantiates a new OrganisationIssuer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganisationIssuerWithDefaults

`func NewOrganisationIssuerWithDefaults() *OrganisationIssuer`

NewOrganisationIssuerWithDefaults instantiates a new OrganisationIssuer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *OrganisationIssuer) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OrganisationIssuer) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OrganisationIssuer) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OrganisationIssuer) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrganisationIssuer) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganisationIssuer) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganisationIssuer) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEnabled

`func (o *OrganisationIssuer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OrganisationIssuer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OrganisationIssuer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIssuerUrl

`func (o *OrganisationIssuer) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### GetIssuerUrlOk

`func (o *OrganisationIssuer) GetIssuerUrlOk() (*string, bool)`

GetIssuerUrlOk returns a tuple with the IssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUrl

`func (o *OrganisationIssuer) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.


### SetIssuerUrlNil

`func (o *OrganisationIssuer) SetIssuerUrlNil(b bool)`

 SetIssuerUrlNil sets the value for IssuerUrl to be an explicit nil

### UnsetIssuerUrl
`func (o *OrganisationIssuer) UnsetIssuerUrl()`

UnsetIssuerUrl ensures that no value is present for IssuerUrl, not even an explicit nil
### GetJwksUri

`func (o *OrganisationIssuer) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *OrganisationIssuer) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *OrganisationIssuer) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.

### HasJwksUri

`func (o *OrganisationIssuer) HasJwksUri() bool`

HasJwksUri returns a boolean if a field has been set.

### GetOrganisationId

`func (o *OrganisationIssuer) GetOrganisationId() int64`

GetOrganisationId returns the OrganisationId field if non-nil, zero value otherwise.

### GetOrganisationIdOk

`func (o *OrganisationIssuer) GetOrganisationIdOk() (*int64, bool)`

GetOrganisationIdOk returns a tuple with the OrganisationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisationId

`func (o *OrganisationIssuer) SetOrganisationId(v int64)`

SetOrganisationId sets OrganisationId field to given value.


### GetOrganisationIssuerId

`func (o *OrganisationIssuer) GetOrganisationIssuerId() int64`

GetOrganisationIssuerId returns the OrganisationIssuerId field if non-nil, zero value otherwise.

### GetOrganisationIssuerIdOk

`func (o *OrganisationIssuer) GetOrganisationIssuerIdOk() (*int64, bool)`

GetOrganisationIssuerIdOk returns a tuple with the OrganisationIssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisationIssuerId

`func (o *OrganisationIssuer) SetOrganisationIssuerId(v int64)`

SetOrganisationIssuerId sets OrganisationIssuerId field to given value.


### GetUpdatedAt

`func (o *OrganisationIssuer) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganisationIssuer) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganisationIssuer) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


