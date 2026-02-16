# SingleCodeSignatureModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificates** | [**[]SingleCodeCertificateModel**](SingleCodeCertificateModel.md) |  | 
**AuthenticodeDigest** | **string** |  | 

## Methods

### NewSingleCodeSignatureModel

`func NewSingleCodeSignatureModel(certificates []SingleCodeCertificateModel, authenticodeDigest string, ) *SingleCodeSignatureModel`

NewSingleCodeSignatureModel instantiates a new SingleCodeSignatureModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleCodeSignatureModelWithDefaults

`func NewSingleCodeSignatureModelWithDefaults() *SingleCodeSignatureModel`

NewSingleCodeSignatureModelWithDefaults instantiates a new SingleCodeSignatureModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificates

`func (o *SingleCodeSignatureModel) GetCertificates() []SingleCodeCertificateModel`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *SingleCodeSignatureModel) GetCertificatesOk() (*[]SingleCodeCertificateModel, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *SingleCodeSignatureModel) SetCertificates(v []SingleCodeCertificateModel)`

SetCertificates sets Certificates field to given value.


### GetAuthenticodeDigest

`func (o *SingleCodeSignatureModel) GetAuthenticodeDigest() string`

GetAuthenticodeDigest returns the AuthenticodeDigest field if non-nil, zero value otherwise.

### GetAuthenticodeDigestOk

`func (o *SingleCodeSignatureModel) GetAuthenticodeDigestOk() (*string, bool)`

GetAuthenticodeDigestOk returns a tuple with the AuthenticodeDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticodeDigest

`func (o *SingleCodeSignatureModel) SetAuthenticodeDigest(v string)`

SetAuthenticodeDigest sets AuthenticodeDigest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


