# SignatureStorageInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Storage class — typically reg, stack, mem or unknown. Not restricted to a fixed set. | [optional] 
**Location** | Pointer to **string** | Register name or stack slot. | [optional] 

## Methods

### NewSignatureStorageInput

`func NewSignatureStorageInput() *SignatureStorageInput`

NewSignatureStorageInput instantiates a new SignatureStorageInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureStorageInputWithDefaults

`func NewSignatureStorageInputWithDefaults() *SignatureStorageInput`

NewSignatureStorageInputWithDefaults instantiates a new SignatureStorageInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *SignatureStorageInput) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SignatureStorageInput) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SignatureStorageInput) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SignatureStorageInput) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLocation

`func (o *SignatureStorageInput) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SignatureStorageInput) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SignatureStorageInput) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SignatureStorageInput) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


