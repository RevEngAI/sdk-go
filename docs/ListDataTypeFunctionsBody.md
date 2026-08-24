# ListDataTypeFunctionsBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]DataTypeFunctionEntry**](DataTypeFunctionEntry.md) | The page of matching functions, in ascending function ID order. Empty when no function uses the type. | 
**NextAfterFunctionId** | Pointer to **int64** | Pass as after_function_id to fetch the next page. Absent on the last page. | [optional] 

## Methods

### NewListDataTypeFunctionsBody

`func NewListDataTypeFunctionsBody(items []DataTypeFunctionEntry, ) *ListDataTypeFunctionsBody`

NewListDataTypeFunctionsBody instantiates a new ListDataTypeFunctionsBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDataTypeFunctionsBodyWithDefaults

`func NewListDataTypeFunctionsBodyWithDefaults() *ListDataTypeFunctionsBody`

NewListDataTypeFunctionsBodyWithDefaults instantiates a new ListDataTypeFunctionsBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListDataTypeFunctionsBody) GetItems() []DataTypeFunctionEntry`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListDataTypeFunctionsBody) GetItemsOk() (*[]DataTypeFunctionEntry, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListDataTypeFunctionsBody) SetItems(v []DataTypeFunctionEntry)`

SetItems sets Items field to given value.


### GetNextAfterFunctionId

`func (o *ListDataTypeFunctionsBody) GetNextAfterFunctionId() int64`

GetNextAfterFunctionId returns the NextAfterFunctionId field if non-nil, zero value otherwise.

### GetNextAfterFunctionIdOk

`func (o *ListDataTypeFunctionsBody) GetNextAfterFunctionIdOk() (*int64, bool)`

GetNextAfterFunctionIdOk returns a tuple with the NextAfterFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAfterFunctionId

`func (o *ListDataTypeFunctionsBody) SetNextAfterFunctionId(v int64)`

SetNextAfterFunctionId sets NextAfterFunctionId field to given value.

### HasNextAfterFunctionId

`func (o *ListDataTypeFunctionsBody) HasNextAfterFunctionId() bool`

HasNextAfterFunctionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


