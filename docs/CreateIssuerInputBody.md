# CreateIssuerInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | OIDC client ID for audience validation (recommended) | [optional] 
**IssuerUrl** | **string** | OIDC issuer URL to trust | 

## Methods

### NewCreateIssuerInputBody

`func NewCreateIssuerInputBody(issuerUrl string, ) *CreateIssuerInputBody`

NewCreateIssuerInputBody instantiates a new CreateIssuerInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIssuerInputBodyWithDefaults

`func NewCreateIssuerInputBodyWithDefaults() *CreateIssuerInputBody`

NewCreateIssuerInputBodyWithDefaults instantiates a new CreateIssuerInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *CreateIssuerInputBody) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CreateIssuerInputBody) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CreateIssuerInputBody) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *CreateIssuerInputBody) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetIssuerUrl

`func (o *CreateIssuerInputBody) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### GetIssuerUrlOk

`func (o *CreateIssuerInputBody) GetIssuerUrlOk() (*string, bool)`

GetIssuerUrlOk returns a tuple with the IssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUrl

`func (o *CreateIssuerInputBody) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


