# FunctionRenameMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionId** | **int64** | The ID of the function to rename | 
**NewName** | **string** | The new name for the function | 
**NewMangledName** | **string** | The new mangled name for the function | 

## Methods

### NewFunctionRenameMap

`func NewFunctionRenameMap(functionId int64, newName string, newMangledName string, ) *FunctionRenameMap`

NewFunctionRenameMap instantiates a new FunctionRenameMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionRenameMapWithDefaults

`func NewFunctionRenameMapWithDefaults() *FunctionRenameMap`

NewFunctionRenameMapWithDefaults instantiates a new FunctionRenameMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionId

`func (o *FunctionRenameMap) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *FunctionRenameMap) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *FunctionRenameMap) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetNewName

`func (o *FunctionRenameMap) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *FunctionRenameMap) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *FunctionRenameMap) SetNewName(v string)`

SetNewName sets NewName field to given value.


### GetNewMangledName

`func (o *FunctionRenameMap) GetNewMangledName() string`

GetNewMangledName returns the NewMangledName field if non-nil, zero value otherwise.

### GetNewMangledNameOk

`func (o *FunctionRenameMap) GetNewMangledNameOk() (*string, bool)`

GetNewMangledNameOk returns a tuple with the NewMangledName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewMangledName

`func (o *FunctionRenameMap) SetNewMangledName(v string)`

SetNewMangledName sets NewMangledName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


