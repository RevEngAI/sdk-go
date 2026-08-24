# CopyFunctionSignaturesInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Copies** | [**[]CopySignatureItem**](CopySignatureItem.md) | Signatures to copy. No target may repeat, and no pair may name the same function twice. | 

## Methods

### NewCopyFunctionSignaturesInputBody

`func NewCopyFunctionSignaturesInputBody(copies []CopySignatureItem, ) *CopyFunctionSignaturesInputBody`

NewCopyFunctionSignaturesInputBody instantiates a new CopyFunctionSignaturesInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyFunctionSignaturesInputBodyWithDefaults

`func NewCopyFunctionSignaturesInputBodyWithDefaults() *CopyFunctionSignaturesInputBody`

NewCopyFunctionSignaturesInputBodyWithDefaults instantiates a new CopyFunctionSignaturesInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopies

`func (o *CopyFunctionSignaturesInputBody) GetCopies() []CopySignatureItem`

GetCopies returns the Copies field if non-nil, zero value otherwise.

### GetCopiesOk

`func (o *CopyFunctionSignaturesInputBody) GetCopiesOk() (*[]CopySignatureItem, bool)`

GetCopiesOk returns a tuple with the Copies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopies

`func (o *CopyFunctionSignaturesInputBody) SetCopies(v []CopySignatureItem)`

SetCopies sets Copies field to given value.


### SetCopiesNil

`func (o *CopyFunctionSignaturesInputBody) SetCopiesNil(b bool)`

 SetCopiesNil sets the value for Copies to be an explicit nil

### UnsetCopies
`func (o *CopyFunctionSignaturesInputBody) UnsetCopies()`

UnsetCopies ensures that no value is present for Copies, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


