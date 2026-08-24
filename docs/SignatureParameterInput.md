# SignatureParameterInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BitLength** | Pointer to **int64** | Width in bits, when the parameter occupies less than its type&#39;s full size. | [optional] 
**DataTypeId** | Pointer to **int64** | The parameter&#39;s type, which must belong to this analysis. Omit for an unresolved type. | [optional] 
**Name** | Pointer to **string** | Parameter name. Omit for an unnamed parameter. | [optional] 
**Ordinal** | **int64** | Zero-based argument position. Must equal the parameter&#39;s index in the list. | 
**Storage** | Pointer to [**SignatureStorageInput**](SignatureStorageInput.md) | Where the parameter is passed. | [optional] 

## Methods

### NewSignatureParameterInput

`func NewSignatureParameterInput(ordinal int64, ) *SignatureParameterInput`

NewSignatureParameterInput instantiates a new SignatureParameterInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureParameterInputWithDefaults

`func NewSignatureParameterInputWithDefaults() *SignatureParameterInput`

NewSignatureParameterInputWithDefaults instantiates a new SignatureParameterInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBitLength

`func (o *SignatureParameterInput) GetBitLength() int64`

GetBitLength returns the BitLength field if non-nil, zero value otherwise.

### GetBitLengthOk

`func (o *SignatureParameterInput) GetBitLengthOk() (*int64, bool)`

GetBitLengthOk returns a tuple with the BitLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitLength

`func (o *SignatureParameterInput) SetBitLength(v int64)`

SetBitLength sets BitLength field to given value.

### HasBitLength

`func (o *SignatureParameterInput) HasBitLength() bool`

HasBitLength returns a boolean if a field has been set.

### GetDataTypeId

`func (o *SignatureParameterInput) GetDataTypeId() int64`

GetDataTypeId returns the DataTypeId field if non-nil, zero value otherwise.

### GetDataTypeIdOk

`func (o *SignatureParameterInput) GetDataTypeIdOk() (*int64, bool)`

GetDataTypeIdOk returns a tuple with the DataTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypeId

`func (o *SignatureParameterInput) SetDataTypeId(v int64)`

SetDataTypeId sets DataTypeId field to given value.

### HasDataTypeId

`func (o *SignatureParameterInput) HasDataTypeId() bool`

HasDataTypeId returns a boolean if a field has been set.

### GetName

`func (o *SignatureParameterInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SignatureParameterInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SignatureParameterInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SignatureParameterInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrdinal

`func (o *SignatureParameterInput) GetOrdinal() int64`

GetOrdinal returns the Ordinal field if non-nil, zero value otherwise.

### GetOrdinalOk

`func (o *SignatureParameterInput) GetOrdinalOk() (*int64, bool)`

GetOrdinalOk returns a tuple with the Ordinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdinal

`func (o *SignatureParameterInput) SetOrdinal(v int64)`

SetOrdinal sets Ordinal field to given value.


### GetStorage

`func (o *SignatureParameterInput) GetStorage() SignatureStorageInput`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *SignatureParameterInput) GetStorageOk() (*SignatureStorageInput, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *SignatureParameterInput) SetStorage(v SignatureStorageInput)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *SignatureParameterInput) HasStorage() bool`

HasStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


