# SignatureParameterEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BitLength** | Pointer to **int64** | Width in bits, when the parameter occupies less than its type&#39;s full size. | [optional] 
**DataTypeId** | Pointer to **int64** | The parameter&#39;s type, resolvable against the analysis data types list. Absent when the type could not be resolved. | [optional] 
**Name** | Pointer to **string** | Parameter name, absent when the producer had none. | [optional] 
**Ordinal** | **int64** | Zero-based argument position. | 
**Storage** | Pointer to [**SignatureStorageEntry**](SignatureStorageEntry.md) | Where the parameter is passed. | [optional] 

## Methods

### NewSignatureParameterEntry

`func NewSignatureParameterEntry(ordinal int64, ) *SignatureParameterEntry`

NewSignatureParameterEntry instantiates a new SignatureParameterEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureParameterEntryWithDefaults

`func NewSignatureParameterEntryWithDefaults() *SignatureParameterEntry`

NewSignatureParameterEntryWithDefaults instantiates a new SignatureParameterEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBitLength

`func (o *SignatureParameterEntry) GetBitLength() int64`

GetBitLength returns the BitLength field if non-nil, zero value otherwise.

### GetBitLengthOk

`func (o *SignatureParameterEntry) GetBitLengthOk() (*int64, bool)`

GetBitLengthOk returns a tuple with the BitLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitLength

`func (o *SignatureParameterEntry) SetBitLength(v int64)`

SetBitLength sets BitLength field to given value.

### HasBitLength

`func (o *SignatureParameterEntry) HasBitLength() bool`

HasBitLength returns a boolean if a field has been set.

### GetDataTypeId

`func (o *SignatureParameterEntry) GetDataTypeId() int64`

GetDataTypeId returns the DataTypeId field if non-nil, zero value otherwise.

### GetDataTypeIdOk

`func (o *SignatureParameterEntry) GetDataTypeIdOk() (*int64, bool)`

GetDataTypeIdOk returns a tuple with the DataTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypeId

`func (o *SignatureParameterEntry) SetDataTypeId(v int64)`

SetDataTypeId sets DataTypeId field to given value.

### HasDataTypeId

`func (o *SignatureParameterEntry) HasDataTypeId() bool`

HasDataTypeId returns a boolean if a field has been set.

### GetName

`func (o *SignatureParameterEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SignatureParameterEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SignatureParameterEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SignatureParameterEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrdinal

`func (o *SignatureParameterEntry) GetOrdinal() int64`

GetOrdinal returns the Ordinal field if non-nil, zero value otherwise.

### GetOrdinalOk

`func (o *SignatureParameterEntry) GetOrdinalOk() (*int64, bool)`

GetOrdinalOk returns a tuple with the Ordinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdinal

`func (o *SignatureParameterEntry) SetOrdinal(v int64)`

SetOrdinal sets Ordinal field to given value.


### GetStorage

`func (o *SignatureParameterEntry) GetStorage() SignatureStorageEntry`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *SignatureParameterEntry) GetStorageOk() (*SignatureStorageEntry, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *SignatureParameterEntry) SetStorage(v SignatureStorageEntry)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *SignatureParameterEntry) HasStorage() bool`

HasStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


