# ELFRelocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **int32** |  | 
**Type** | **string** |  | 
**Size** | **int32** |  | 
**Addend** | **int32** |  | 
**SymbolName** | **string** |  | 
**IsDynamic** | **bool** |  | 
**IsPltgot** | **bool** |  | 

## Methods

### NewELFRelocation

`func NewELFRelocation(address int32, type_ string, size int32, addend int32, symbolName string, isDynamic bool, isPltgot bool, ) *ELFRelocation`

NewELFRelocation instantiates a new ELFRelocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewELFRelocationWithDefaults

`func NewELFRelocationWithDefaults() *ELFRelocation`

NewELFRelocationWithDefaults instantiates a new ELFRelocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ELFRelocation) GetAddress() int32`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ELFRelocation) GetAddressOk() (*int32, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ELFRelocation) SetAddress(v int32)`

SetAddress sets Address field to given value.


### GetType

`func (o *ELFRelocation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ELFRelocation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ELFRelocation) SetType(v string)`

SetType sets Type field to given value.


### GetSize

`func (o *ELFRelocation) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ELFRelocation) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ELFRelocation) SetSize(v int32)`

SetSize sets Size field to given value.


### GetAddend

`func (o *ELFRelocation) GetAddend() int32`

GetAddend returns the Addend field if non-nil, zero value otherwise.

### GetAddendOk

`func (o *ELFRelocation) GetAddendOk() (*int32, bool)`

GetAddendOk returns a tuple with the Addend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddend

`func (o *ELFRelocation) SetAddend(v int32)`

SetAddend sets Addend field to given value.


### GetSymbolName

`func (o *ELFRelocation) GetSymbolName() string`

GetSymbolName returns the SymbolName field if non-nil, zero value otherwise.

### GetSymbolNameOk

`func (o *ELFRelocation) GetSymbolNameOk() (*string, bool)`

GetSymbolNameOk returns a tuple with the SymbolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbolName

`func (o *ELFRelocation) SetSymbolName(v string)`

SetSymbolName sets SymbolName field to given value.


### GetIsDynamic

`func (o *ELFRelocation) GetIsDynamic() bool`

GetIsDynamic returns the IsDynamic field if non-nil, zero value otherwise.

### GetIsDynamicOk

`func (o *ELFRelocation) GetIsDynamicOk() (*bool, bool)`

GetIsDynamicOk returns a tuple with the IsDynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDynamic

`func (o *ELFRelocation) SetIsDynamic(v bool)`

SetIsDynamic sets IsDynamic field to given value.


### GetIsPltgot

`func (o *ELFRelocation) GetIsPltgot() bool`

GetIsPltgot returns the IsPltgot field if non-nil, zero value otherwise.

### GetIsPltgotOk

`func (o *ELFRelocation) GetIsPltgotOk() (*bool, bool)`

GetIsPltgotOk returns a tuple with the IsPltgot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPltgot

`func (o *ELFRelocation) SetIsPltgot(v bool)`

SetIsPltgot sets IsPltgot field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


