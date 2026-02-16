# FunctionMatchingFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinaryIds** | Pointer to **[]int32** | ID&#39;s of binaries to limit the search to, if empty, search all scoped binaries | [optional] [default to {}]
**CollectionIds** | Pointer to **[]int32** | ID&#39;s of collections to limit the search to, if empty, search all scoped collections | [optional] [default to {}]
**FunctionIds** | Pointer to **[]int64** | ID&#39;s of functions to limit the search to, if empty, search all scoped functions | [optional] [default to {}]
**DebugTypes** | Pointer to **[]string** | Limit the search to specific debug types, if empty, search all scoped debug &amp; non-debug functions | [optional] [default to {}]

## Methods

### NewFunctionMatchingFilters

`func NewFunctionMatchingFilters() *FunctionMatchingFilters`

NewFunctionMatchingFilters instantiates a new FunctionMatchingFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionMatchingFiltersWithDefaults

`func NewFunctionMatchingFiltersWithDefaults() *FunctionMatchingFilters`

NewFunctionMatchingFiltersWithDefaults instantiates a new FunctionMatchingFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaryIds

`func (o *FunctionMatchingFilters) GetBinaryIds() []int32`

GetBinaryIds returns the BinaryIds field if non-nil, zero value otherwise.

### GetBinaryIdsOk

`func (o *FunctionMatchingFilters) GetBinaryIdsOk() (*[]int32, bool)`

GetBinaryIdsOk returns a tuple with the BinaryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryIds

`func (o *FunctionMatchingFilters) SetBinaryIds(v []int32)`

SetBinaryIds sets BinaryIds field to given value.

### HasBinaryIds

`func (o *FunctionMatchingFilters) HasBinaryIds() bool`

HasBinaryIds returns a boolean if a field has been set.

### GetCollectionIds

`func (o *FunctionMatchingFilters) GetCollectionIds() []int32`

GetCollectionIds returns the CollectionIds field if non-nil, zero value otherwise.

### GetCollectionIdsOk

`func (o *FunctionMatchingFilters) GetCollectionIdsOk() (*[]int32, bool)`

GetCollectionIdsOk returns a tuple with the CollectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionIds

`func (o *FunctionMatchingFilters) SetCollectionIds(v []int32)`

SetCollectionIds sets CollectionIds field to given value.

### HasCollectionIds

`func (o *FunctionMatchingFilters) HasCollectionIds() bool`

HasCollectionIds returns a boolean if a field has been set.

### GetFunctionIds

`func (o *FunctionMatchingFilters) GetFunctionIds() []int64`

GetFunctionIds returns the FunctionIds field if non-nil, zero value otherwise.

### GetFunctionIdsOk

`func (o *FunctionMatchingFilters) GetFunctionIdsOk() (*[]int64, bool)`

GetFunctionIdsOk returns a tuple with the FunctionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionIds

`func (o *FunctionMatchingFilters) SetFunctionIds(v []int64)`

SetFunctionIds sets FunctionIds field to given value.

### HasFunctionIds

`func (o *FunctionMatchingFilters) HasFunctionIds() bool`

HasFunctionIds returns a boolean if a field has been set.

### GetDebugTypes

`func (o *FunctionMatchingFilters) GetDebugTypes() []string`

GetDebugTypes returns the DebugTypes field if non-nil, zero value otherwise.

### GetDebugTypesOk

`func (o *FunctionMatchingFilters) GetDebugTypesOk() (*[]string, bool)`

GetDebugTypesOk returns a tuple with the DebugTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugTypes

`func (o *FunctionMatchingFilters) SetDebugTypes(v []string)`

SetDebugTypes sets DebugTypes field to given value.

### HasDebugTypes

`func (o *FunctionMatchingFilters) HasDebugTypes() bool`

HasDebugTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


