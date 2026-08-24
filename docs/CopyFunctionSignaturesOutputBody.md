# CopyFunctionSignaturesOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTypes** | [**[]DataTypeEntry**](DataTypeEntry.md) | The data types this analysis gained or had replaced, ordered by data_type_id. Empty when every type the copied signatures need was already stored unchanged. | 
**Signatures** | [**[]FunctionSignatureEntry**](FunctionSignatureEntry.md) | The stored signatures of the target functions, in request order. | 

## Methods

### NewCopyFunctionSignaturesOutputBody

`func NewCopyFunctionSignaturesOutputBody(dataTypes []DataTypeEntry, signatures []FunctionSignatureEntry, ) *CopyFunctionSignaturesOutputBody`

NewCopyFunctionSignaturesOutputBody instantiates a new CopyFunctionSignaturesOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyFunctionSignaturesOutputBodyWithDefaults

`func NewCopyFunctionSignaturesOutputBodyWithDefaults() *CopyFunctionSignaturesOutputBody`

NewCopyFunctionSignaturesOutputBodyWithDefaults instantiates a new CopyFunctionSignaturesOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTypes

`func (o *CopyFunctionSignaturesOutputBody) GetDataTypes() []DataTypeEntry`

GetDataTypes returns the DataTypes field if non-nil, zero value otherwise.

### GetDataTypesOk

`func (o *CopyFunctionSignaturesOutputBody) GetDataTypesOk() (*[]DataTypeEntry, bool)`

GetDataTypesOk returns a tuple with the DataTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypes

`func (o *CopyFunctionSignaturesOutputBody) SetDataTypes(v []DataTypeEntry)`

SetDataTypes sets DataTypes field to given value.


### SetDataTypesNil

`func (o *CopyFunctionSignaturesOutputBody) SetDataTypesNil(b bool)`

 SetDataTypesNil sets the value for DataTypes to be an explicit nil

### UnsetDataTypes
`func (o *CopyFunctionSignaturesOutputBody) UnsetDataTypes()`

UnsetDataTypes ensures that no value is present for DataTypes, not even an explicit nil
### GetSignatures

`func (o *CopyFunctionSignaturesOutputBody) GetSignatures() []FunctionSignatureEntry`

GetSignatures returns the Signatures field if non-nil, zero value otherwise.

### GetSignaturesOk

`func (o *CopyFunctionSignaturesOutputBody) GetSignaturesOk() (*[]FunctionSignatureEntry, bool)`

GetSignaturesOk returns a tuple with the Signatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatures

`func (o *CopyFunctionSignaturesOutputBody) SetSignatures(v []FunctionSignatureEntry)`

SetSignatures sets Signatures field to given value.


### SetSignaturesNil

`func (o *CopyFunctionSignaturesOutputBody) SetSignaturesNil(b bool)`

 SetSignaturesNil sets the value for Signatures to be an explicit nil

### UnsetSignatures
`func (o *CopyFunctionSignaturesOutputBody) UnsetSignatures()`

UnsetSignatures ensures that no value is present for Signatures, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


