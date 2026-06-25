# SSOProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthUrl** | **string** | Base authorization URL. Append code_challenge, code_challenge_method, redirect_uri, and state query parameters to complete the PKCE flow. | 
**Issuer** | **string** | Issuer URL of the identity provider. Pass this value as the issuer field in the SSO callback request. | 

## Methods

### NewSSOProvider

`func NewSSOProvider(authUrl string, issuer string, ) *SSOProvider`

NewSSOProvider instantiates a new SSOProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSOProviderWithDefaults

`func NewSSOProviderWithDefaults() *SSOProvider`

NewSSOProviderWithDefaults instantiates a new SSOProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthUrl

`func (o *SSOProvider) GetAuthUrl() string`

GetAuthUrl returns the AuthUrl field if non-nil, zero value otherwise.

### GetAuthUrlOk

`func (o *SSOProvider) GetAuthUrlOk() (*string, bool)`

GetAuthUrlOk returns a tuple with the AuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUrl

`func (o *SSOProvider) SetAuthUrl(v string)`

SetAuthUrl sets AuthUrl field to given value.


### GetIssuer

`func (o *SSOProvider) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *SSOProvider) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *SSOProvider) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


