# FunctionListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Function id | 
**Name** | **string** | Name of the function | 
**NameSourceType** | **string** | The source (process) the function name came from | 
**NameSource** | [**NameSourceType**](NameSourceType.md) | The source of the current function name. | 
**MangledName** | **string** | Mangled name of the function | 
**Vaddr** | **int64** | Function virtual address | 
**Size** | **int32** | Function size in bytes | 
**Debug** | **bool** | Whether the function has debug information | 

## Methods

### NewFunctionListItem

`func NewFunctionListItem(id int64, name string, nameSourceType string, nameSource NameSourceType, mangledName string, vaddr int64, size int32, debug bool, ) *FunctionListItem`

NewFunctionListItem instantiates a new FunctionListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionListItemWithDefaults

`func NewFunctionListItemWithDefaults() *FunctionListItem`

NewFunctionListItemWithDefaults instantiates a new FunctionListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FunctionListItem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FunctionListItem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FunctionListItem) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *FunctionListItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionListItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionListItem) SetName(v string)`

SetName sets Name field to given value.


### GetNameSourceType

`func (o *FunctionListItem) GetNameSourceType() string`

GetNameSourceType returns the NameSourceType field if non-nil, zero value otherwise.

### GetNameSourceTypeOk

`func (o *FunctionListItem) GetNameSourceTypeOk() (*string, bool)`

GetNameSourceTypeOk returns a tuple with the NameSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSourceType

`func (o *FunctionListItem) SetNameSourceType(v string)`

SetNameSourceType sets NameSourceType field to given value.


### GetNameSource

`func (o *FunctionListItem) GetNameSource() NameSourceType`

GetNameSource returns the NameSource field if non-nil, zero value otherwise.

### GetNameSourceOk

`func (o *FunctionListItem) GetNameSourceOk() (*NameSourceType, bool)`

GetNameSourceOk returns a tuple with the NameSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSource

`func (o *FunctionListItem) SetNameSource(v NameSourceType)`

SetNameSource sets NameSource field to given value.


### GetMangledName

`func (o *FunctionListItem) GetMangledName() string`

GetMangledName returns the MangledName field if non-nil, zero value otherwise.

### GetMangledNameOk

`func (o *FunctionListItem) GetMangledNameOk() (*string, bool)`

GetMangledNameOk returns a tuple with the MangledName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMangledName

`func (o *FunctionListItem) SetMangledName(v string)`

SetMangledName sets MangledName field to given value.


### GetVaddr

`func (o *FunctionListItem) GetVaddr() int64`

GetVaddr returns the Vaddr field if non-nil, zero value otherwise.

### GetVaddrOk

`func (o *FunctionListItem) GetVaddrOk() (*int64, bool)`

GetVaddrOk returns a tuple with the Vaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaddr

`func (o *FunctionListItem) SetVaddr(v int64)`

SetVaddr sets Vaddr field to given value.


### GetSize

`func (o *FunctionListItem) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FunctionListItem) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FunctionListItem) SetSize(v int32)`

SetSize sets Size field to given value.


### GetDebug

`func (o *FunctionListItem) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *FunctionListItem) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *FunctionListItem) SetDebug(v bool)`

SetDebug sets Debug field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


