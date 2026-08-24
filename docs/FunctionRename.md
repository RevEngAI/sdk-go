# FunctionRename

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewName** | **string** | The new name for the function | 
**NewMangledName** | **string** | The new mangled name for the function | 
**SourceType** | Pointer to [**FunctionSourceType**](FunctionSourceType.md) | The source that triggered the rename | [optional] [default to FUNCTIONSOURCETYPE_USER]

## Methods

### NewFunctionRename

`func NewFunctionRename(newName string, newMangledName string, ) *FunctionRename`

NewFunctionRename instantiates a new FunctionRename object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionRenameWithDefaults

`func NewFunctionRenameWithDefaults() *FunctionRename`

NewFunctionRenameWithDefaults instantiates a new FunctionRename object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewName

`func (o *FunctionRename) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *FunctionRename) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *FunctionRename) SetNewName(v string)`

SetNewName sets NewName field to given value.


### GetNewMangledName

`func (o *FunctionRename) GetNewMangledName() string`

GetNewMangledName returns the NewMangledName field if non-nil, zero value otherwise.

### GetNewMangledNameOk

`func (o *FunctionRename) GetNewMangledNameOk() (*string, bool)`

GetNewMangledNameOk returns a tuple with the NewMangledName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewMangledName

`func (o *FunctionRename) SetNewMangledName(v string)`

SetNewMangledName sets NewMangledName field to given value.


### GetSourceType

`func (o *FunctionRename) GetSourceType() FunctionSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *FunctionRename) GetSourceTypeOk() (*FunctionSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *FunctionRename) SetSourceType(v FunctionSourceType)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *FunctionRename) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


