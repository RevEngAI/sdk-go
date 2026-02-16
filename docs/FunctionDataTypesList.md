# FunctionDataTypesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** | Total number of functions in analysis | [optional] [default to 0]
**TotalDataTypesCount** | Pointer to **int32** | Total number of functions with data types | [optional] [default to 0]
**Items** | [**[]FunctionDataTypesListItem**](FunctionDataTypesListItem.md) | List of function data types information | 

## Methods

### NewFunctionDataTypesList

`func NewFunctionDataTypesList(items []FunctionDataTypesListItem, ) *FunctionDataTypesList`

NewFunctionDataTypesList instantiates a new FunctionDataTypesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionDataTypesListWithDefaults

`func NewFunctionDataTypesListWithDefaults() *FunctionDataTypesList`

NewFunctionDataTypesListWithDefaults instantiates a new FunctionDataTypesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *FunctionDataTypesList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *FunctionDataTypesList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *FunctionDataTypesList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *FunctionDataTypesList) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetTotalDataTypesCount

`func (o *FunctionDataTypesList) GetTotalDataTypesCount() int32`

GetTotalDataTypesCount returns the TotalDataTypesCount field if non-nil, zero value otherwise.

### GetTotalDataTypesCountOk

`func (o *FunctionDataTypesList) GetTotalDataTypesCountOk() (*int32, bool)`

GetTotalDataTypesCountOk returns a tuple with the TotalDataTypesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDataTypesCount

`func (o *FunctionDataTypesList) SetTotalDataTypesCount(v int32)`

SetTotalDataTypesCount sets TotalDataTypesCount field to given value.

### HasTotalDataTypesCount

`func (o *FunctionDataTypesList) HasTotalDataTypesCount() bool`

HasTotalDataTypesCount returns a boolean if a field has been set.

### GetItems

`func (o *FunctionDataTypesList) GetItems() []FunctionDataTypesListItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FunctionDataTypesList) GetItemsOk() (*[]FunctionDataTypesListItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FunctionDataTypesList) SetItems(v []FunctionDataTypesListItem)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


