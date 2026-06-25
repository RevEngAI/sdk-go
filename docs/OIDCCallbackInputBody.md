# OIDCCallbackInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Authorization code from the identity provider | 
**CodeVerifier** | **string** | PKCE code verifier (the raw secret, not the challenge) | 
**Issuer** | **string** | OIDC issuer URL (as returned by the SSO providers endpoint) | 
**RedirectUri** | **string** | Redirect URI used when initiating the authorization request; must match exactly | 

## Methods

### NewOIDCCallbackInputBody

`func NewOIDCCallbackInputBody(code string, codeVerifier string, issuer string, redirectUri string, ) *OIDCCallbackInputBody`

NewOIDCCallbackInputBody instantiates a new OIDCCallbackInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDCCallbackInputBodyWithDefaults

`func NewOIDCCallbackInputBodyWithDefaults() *OIDCCallbackInputBody`

NewOIDCCallbackInputBodyWithDefaults instantiates a new OIDCCallbackInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *OIDCCallbackInputBody) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OIDCCallbackInputBody) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OIDCCallbackInputBody) SetCode(v string)`

SetCode sets Code field to given value.


### GetCodeVerifier

`func (o *OIDCCallbackInputBody) GetCodeVerifier() string`

GetCodeVerifier returns the CodeVerifier field if non-nil, zero value otherwise.

### GetCodeVerifierOk

`func (o *OIDCCallbackInputBody) GetCodeVerifierOk() (*string, bool)`

GetCodeVerifierOk returns a tuple with the CodeVerifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeVerifier

`func (o *OIDCCallbackInputBody) SetCodeVerifier(v string)`

SetCodeVerifier sets CodeVerifier field to given value.


### GetIssuer

`func (o *OIDCCallbackInputBody) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OIDCCallbackInputBody) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OIDCCallbackInputBody) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetRedirectUri

`func (o *OIDCCallbackInputBody) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *OIDCCallbackInputBody) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *OIDCCallbackInputBody) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


