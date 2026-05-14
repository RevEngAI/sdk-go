# BatchRenameItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionId** | **int64** | Function ID to rename | 
**NewMangledName** | Pointer to **string** | New mangled function name | [optional] 
**NewName** | **string** | New function name | 

## Methods

### NewBatchRenameItem

`func NewBatchRenameItem(functionId int64, newName string, ) *BatchRenameItem`

NewBatchRenameItem instantiates a new BatchRenameItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchRenameItemWithDefaults

`func NewBatchRenameItemWithDefaults() *BatchRenameItem`

NewBatchRenameItemWithDefaults instantiates a new BatchRenameItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionId

`func (o *BatchRenameItem) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *BatchRenameItem) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *BatchRenameItem) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetNewMangledName

`func (o *BatchRenameItem) GetNewMangledName() string`

GetNewMangledName returns the NewMangledName field if non-nil, zero value otherwise.

### GetNewMangledNameOk

`func (o *BatchRenameItem) GetNewMangledNameOk() (*string, bool)`

GetNewMangledNameOk returns a tuple with the NewMangledName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewMangledName

`func (o *BatchRenameItem) SetNewMangledName(v string)`

SetNewMangledName sets NewMangledName field to given value.

### HasNewMangledName

`func (o *BatchRenameItem) HasNewMangledName() bool`

HasNewMangledName returns a boolean if a field has been set.

### GetNewName

`func (o *BatchRenameItem) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *BatchRenameItem) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *BatchRenameItem) SetNewName(v string)`

SetNewName sets NewName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


