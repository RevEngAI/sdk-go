# GenerationStatusList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** | Total number of functions in analysis | [optional] [default to 0]
**TotalDataTypesCount** | Pointer to **int32** | Total number of functions with data types | [optional] [default to 0]
**Items** | [**[]FunctionDataTypesStatus**](FunctionDataTypesStatus.md) | List of function data types information | 

## Methods

### NewGenerationStatusList

`func NewGenerationStatusList(items []FunctionDataTypesStatus, ) *GenerationStatusList`

NewGenerationStatusList instantiates a new GenerationStatusList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerationStatusListWithDefaults

`func NewGenerationStatusListWithDefaults() *GenerationStatusList`

NewGenerationStatusListWithDefaults instantiates a new GenerationStatusList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *GenerationStatusList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GenerationStatusList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GenerationStatusList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GenerationStatusList) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetTotalDataTypesCount

`func (o *GenerationStatusList) GetTotalDataTypesCount() int32`

GetTotalDataTypesCount returns the TotalDataTypesCount field if non-nil, zero value otherwise.

### GetTotalDataTypesCountOk

`func (o *GenerationStatusList) GetTotalDataTypesCountOk() (*int32, bool)`

GetTotalDataTypesCountOk returns a tuple with the TotalDataTypesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDataTypesCount

`func (o *GenerationStatusList) SetTotalDataTypesCount(v int32)`

SetTotalDataTypesCount sets TotalDataTypesCount field to given value.

### HasTotalDataTypesCount

`func (o *GenerationStatusList) HasTotalDataTypesCount() bool`

HasTotalDataTypesCount returns a boolean if a field has been set.

### GetItems

`func (o *GenerationStatusList) GetItems() []FunctionDataTypesStatus`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GenerationStatusList) GetItemsOk() (*[]FunctionDataTypesStatus, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GenerationStatusList) SetItems(v []FunctionDataTypesStatus)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


