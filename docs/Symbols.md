# Symbols

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseAddress** | **int64** |  | 
**FunctionBoundaries** | Pointer to [**[]FunctionBoundary**](FunctionBoundary.md) |  | [optional] 

## Methods

### NewSymbols

`func NewSymbols(baseAddress int64, ) *Symbols`

NewSymbols instantiates a new Symbols object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSymbolsWithDefaults

`func NewSymbolsWithDefaults() *Symbols`

NewSymbolsWithDefaults instantiates a new Symbols object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseAddress

`func (o *Symbols) GetBaseAddress() int64`

GetBaseAddress returns the BaseAddress field if non-nil, zero value otherwise.

### GetBaseAddressOk

`func (o *Symbols) GetBaseAddressOk() (*int64, bool)`

GetBaseAddressOk returns a tuple with the BaseAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAddress

`func (o *Symbols) SetBaseAddress(v int64)`

SetBaseAddress sets BaseAddress field to given value.


### GetFunctionBoundaries

`func (o *Symbols) GetFunctionBoundaries() []FunctionBoundary`

GetFunctionBoundaries returns the FunctionBoundaries field if non-nil, zero value otherwise.

### GetFunctionBoundariesOk

`func (o *Symbols) GetFunctionBoundariesOk() (*[]FunctionBoundary, bool)`

GetFunctionBoundariesOk returns a tuple with the FunctionBoundaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionBoundaries

`func (o *Symbols) SetFunctionBoundaries(v []FunctionBoundary)`

SetFunctionBoundaries sets FunctionBoundaries field to given value.

### HasFunctionBoundaries

`func (o *Symbols) HasFunctionBoundaries() bool`

HasFunctionBoundaries returns a boolean if a field has been set.

### SetFunctionBoundariesNil

`func (o *Symbols) SetFunctionBoundariesNil(b bool)`

 SetFunctionBoundariesNil sets the value for FunctionBoundaries to be an explicit nil

### UnsetFunctionBoundaries
`func (o *Symbols) UnsetFunctionBoundaries()`

UnsetFunctionBoundaries ensures that no value is present for FunctionBoundaries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


