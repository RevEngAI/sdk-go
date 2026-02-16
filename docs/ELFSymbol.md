# ELFSymbol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Value** | **int32** |  | 
**Size** | **int32** |  | 
**Type** | **string** |  | 
**Binding** | **string** |  | 
**Visibility** | **string** |  | 
**SectionIndex** | **int32** |  | 

## Methods

### NewELFSymbol

`func NewELFSymbol(name string, value int32, size int32, type_ string, binding string, visibility string, sectionIndex int32, ) *ELFSymbol`

NewELFSymbol instantiates a new ELFSymbol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewELFSymbolWithDefaults

`func NewELFSymbolWithDefaults() *ELFSymbol`

NewELFSymbolWithDefaults instantiates a new ELFSymbol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ELFSymbol) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ELFSymbol) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ELFSymbol) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *ELFSymbol) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ELFSymbol) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ELFSymbol) SetValue(v int32)`

SetValue sets Value field to given value.


### GetSize

`func (o *ELFSymbol) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ELFSymbol) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ELFSymbol) SetSize(v int32)`

SetSize sets Size field to given value.


### GetType

`func (o *ELFSymbol) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ELFSymbol) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ELFSymbol) SetType(v string)`

SetType sets Type field to given value.


### GetBinding

`func (o *ELFSymbol) GetBinding() string`

GetBinding returns the Binding field if non-nil, zero value otherwise.

### GetBindingOk

`func (o *ELFSymbol) GetBindingOk() (*string, bool)`

GetBindingOk returns a tuple with the Binding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinding

`func (o *ELFSymbol) SetBinding(v string)`

SetBinding sets Binding field to given value.


### GetVisibility

`func (o *ELFSymbol) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ELFSymbol) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ELFSymbol) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.


### GetSectionIndex

`func (o *ELFSymbol) GetSectionIndex() int32`

GetSectionIndex returns the SectionIndex field if non-nil, zero value otherwise.

### GetSectionIndexOk

`func (o *ELFSymbol) GetSectionIndexOk() (*int32, bool)`

GetSectionIndexOk returns a tuple with the SectionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionIndex

`func (o *ELFSymbol) SetSectionIndex(v int32)`

SetSectionIndex sets SectionIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


