# CodeSignatureModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signed** | **bool** |  | 
**ValidSignature** | **bool** |  | 
**Signatures** | [**[]SingleCodeSignatureModel**](SingleCodeSignatureModel.md) |  | 

## Methods

### NewCodeSignatureModel

`func NewCodeSignatureModel(signed bool, validSignature bool, signatures []SingleCodeSignatureModel, ) *CodeSignatureModel`

NewCodeSignatureModel instantiates a new CodeSignatureModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeSignatureModelWithDefaults

`func NewCodeSignatureModelWithDefaults() *CodeSignatureModel`

NewCodeSignatureModelWithDefaults instantiates a new CodeSignatureModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSigned

`func (o *CodeSignatureModel) GetSigned() bool`

GetSigned returns the Signed field if non-nil, zero value otherwise.

### GetSignedOk

`func (o *CodeSignatureModel) GetSignedOk() (*bool, bool)`

GetSignedOk returns a tuple with the Signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigned

`func (o *CodeSignatureModel) SetSigned(v bool)`

SetSigned sets Signed field to given value.


### GetValidSignature

`func (o *CodeSignatureModel) GetValidSignature() bool`

GetValidSignature returns the ValidSignature field if non-nil, zero value otherwise.

### GetValidSignatureOk

`func (o *CodeSignatureModel) GetValidSignatureOk() (*bool, bool)`

GetValidSignatureOk returns a tuple with the ValidSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidSignature

`func (o *CodeSignatureModel) SetValidSignature(v bool)`

SetValidSignature sets ValidSignature field to given value.


### GetSignatures

`func (o *CodeSignatureModel) GetSignatures() []SingleCodeSignatureModel`

GetSignatures returns the Signatures field if non-nil, zero value otherwise.

### GetSignaturesOk

`func (o *CodeSignatureModel) GetSignaturesOk() (*[]SingleCodeSignatureModel, bool)`

GetSignaturesOk returns a tuple with the Signatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatures

`func (o *CodeSignatureModel) SetSignatures(v []SingleCodeSignatureModel)`

SetSignatures sets Signatures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


