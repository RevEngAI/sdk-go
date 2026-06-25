# IssuerAllowedDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChallengeToken** | Pointer to **string** | DNS TXT challenge token. Add a TXT record at _reveng-verification.&lt;domain&gt; with this value. | [optional] 
**CreatedAt** | **time.Time** |  | 
**Domain** | **string** | Email domain (e.g. acme.com) | 
**IssuerAllowedDomainId** | **int64** |  | 
**OrganisationIssuerId** | **int64** |  | 
**VerificationStatus** | **string** | Domain ownership verification status: PENDING, VERIFIED, or FAILED | 

## Methods

### NewIssuerAllowedDomain

`func NewIssuerAllowedDomain(createdAt time.Time, domain string, issuerAllowedDomainId int64, organisationIssuerId int64, verificationStatus string, ) *IssuerAllowedDomain`

NewIssuerAllowedDomain instantiates a new IssuerAllowedDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuerAllowedDomainWithDefaults

`func NewIssuerAllowedDomainWithDefaults() *IssuerAllowedDomain`

NewIssuerAllowedDomainWithDefaults instantiates a new IssuerAllowedDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChallengeToken

`func (o *IssuerAllowedDomain) GetChallengeToken() string`

GetChallengeToken returns the ChallengeToken field if non-nil, zero value otherwise.

### GetChallengeTokenOk

`func (o *IssuerAllowedDomain) GetChallengeTokenOk() (*string, bool)`

GetChallengeTokenOk returns a tuple with the ChallengeToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeToken

`func (o *IssuerAllowedDomain) SetChallengeToken(v string)`

SetChallengeToken sets ChallengeToken field to given value.

### HasChallengeToken

`func (o *IssuerAllowedDomain) HasChallengeToken() bool`

HasChallengeToken returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IssuerAllowedDomain) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IssuerAllowedDomain) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IssuerAllowedDomain) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDomain

`func (o *IssuerAllowedDomain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *IssuerAllowedDomain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *IssuerAllowedDomain) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetIssuerAllowedDomainId

`func (o *IssuerAllowedDomain) GetIssuerAllowedDomainId() int64`

GetIssuerAllowedDomainId returns the IssuerAllowedDomainId field if non-nil, zero value otherwise.

### GetIssuerAllowedDomainIdOk

`func (o *IssuerAllowedDomain) GetIssuerAllowedDomainIdOk() (*int64, bool)`

GetIssuerAllowedDomainIdOk returns a tuple with the IssuerAllowedDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerAllowedDomainId

`func (o *IssuerAllowedDomain) SetIssuerAllowedDomainId(v int64)`

SetIssuerAllowedDomainId sets IssuerAllowedDomainId field to given value.


### GetOrganisationIssuerId

`func (o *IssuerAllowedDomain) GetOrganisationIssuerId() int64`

GetOrganisationIssuerId returns the OrganisationIssuerId field if non-nil, zero value otherwise.

### GetOrganisationIssuerIdOk

`func (o *IssuerAllowedDomain) GetOrganisationIssuerIdOk() (*int64, bool)`

GetOrganisationIssuerIdOk returns a tuple with the OrganisationIssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisationIssuerId

`func (o *IssuerAllowedDomain) SetOrganisationIssuerId(v int64)`

SetOrganisationIssuerId sets OrganisationIssuerId field to given value.


### GetVerificationStatus

`func (o *IssuerAllowedDomain) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *IssuerAllowedDomain) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *IssuerAllowedDomain) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


