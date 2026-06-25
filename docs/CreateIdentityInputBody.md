# CreateIdentityInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssuerUrl** | **string** | Issuer URL (must be trusted by this organisation) | 
**Subject** | **string** | External subject identifier from the issuer | 
**UserId** | **int64** | Internal user ID to link | 

## Methods

### NewCreateIdentityInputBody

`func NewCreateIdentityInputBody(issuerUrl string, subject string, userId int64, ) *CreateIdentityInputBody`

NewCreateIdentityInputBody instantiates a new CreateIdentityInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIdentityInputBodyWithDefaults

`func NewCreateIdentityInputBodyWithDefaults() *CreateIdentityInputBody`

NewCreateIdentityInputBodyWithDefaults instantiates a new CreateIdentityInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuerUrl

`func (o *CreateIdentityInputBody) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### GetIssuerUrlOk

`func (o *CreateIdentityInputBody) GetIssuerUrlOk() (*string, bool)`

GetIssuerUrlOk returns a tuple with the IssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUrl

`func (o *CreateIdentityInputBody) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.


### GetSubject

`func (o *CreateIdentityInputBody) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CreateIdentityInputBody) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CreateIdentityInputBody) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetUserId

`func (o *CreateIdentityInputBody) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateIdentityInputBody) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateIdentityInputBody) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


