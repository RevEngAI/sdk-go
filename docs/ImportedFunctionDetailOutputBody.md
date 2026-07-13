# ImportedFunctionDetailOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Callers** | [**[]ImportedFunctionCallerEntry**](ImportedFunctionCallerEntry.md) | Internal functions that call this import, resolved via its PLT/stub addresses. | 
**ImportedFunctionId** | **int64** |  | 
**IsFunction** | **bool** | False for imported data symbols. | 
**LibraryName** | **string** | Library the symbol is imported from. &#39;&lt;EXTERNAL&gt;&#39; for unattributed imports. | 
**LibraryVersion** | Pointer to **string** | Versioned symbol tag, when the loader records one. | [optional] 
**Name** | **string** |  | 
**OriginalName** | Pointer to **string** | Pre-demangling / pre-aliasing name, when it differs from name. | [optional] 
**StubVaddrs** | **[]int64** | PLT/stub addresses that resolve external call edges (function_call_edges.callee_vaddr) to this import. Use these to link a caller&#39;s external callee to this import. | 
**Vaddr** | Pointer to **int64** | Virtual address of the import, when known. | [optional] 

## Methods

### NewImportedFunctionDetailOutputBody

`func NewImportedFunctionDetailOutputBody(callers []ImportedFunctionCallerEntry, importedFunctionId int64, isFunction bool, libraryName string, name string, stubVaddrs []int64, ) *ImportedFunctionDetailOutputBody`

NewImportedFunctionDetailOutputBody instantiates a new ImportedFunctionDetailOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportedFunctionDetailOutputBodyWithDefaults

`func NewImportedFunctionDetailOutputBodyWithDefaults() *ImportedFunctionDetailOutputBody`

NewImportedFunctionDetailOutputBodyWithDefaults instantiates a new ImportedFunctionDetailOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallers

`func (o *ImportedFunctionDetailOutputBody) GetCallers() []ImportedFunctionCallerEntry`

GetCallers returns the Callers field if non-nil, zero value otherwise.

### GetCallersOk

`func (o *ImportedFunctionDetailOutputBody) GetCallersOk() (*[]ImportedFunctionCallerEntry, bool)`

GetCallersOk returns a tuple with the Callers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallers

`func (o *ImportedFunctionDetailOutputBody) SetCallers(v []ImportedFunctionCallerEntry)`

SetCallers sets Callers field to given value.


### SetCallersNil

`func (o *ImportedFunctionDetailOutputBody) SetCallersNil(b bool)`

 SetCallersNil sets the value for Callers to be an explicit nil

### UnsetCallers
`func (o *ImportedFunctionDetailOutputBody) UnsetCallers()`

UnsetCallers ensures that no value is present for Callers, not even an explicit nil
### GetImportedFunctionId

`func (o *ImportedFunctionDetailOutputBody) GetImportedFunctionId() int64`

GetImportedFunctionId returns the ImportedFunctionId field if non-nil, zero value otherwise.

### GetImportedFunctionIdOk

`func (o *ImportedFunctionDetailOutputBody) GetImportedFunctionIdOk() (*int64, bool)`

GetImportedFunctionIdOk returns a tuple with the ImportedFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedFunctionId

`func (o *ImportedFunctionDetailOutputBody) SetImportedFunctionId(v int64)`

SetImportedFunctionId sets ImportedFunctionId field to given value.


### GetIsFunction

`func (o *ImportedFunctionDetailOutputBody) GetIsFunction() bool`

GetIsFunction returns the IsFunction field if non-nil, zero value otherwise.

### GetIsFunctionOk

`func (o *ImportedFunctionDetailOutputBody) GetIsFunctionOk() (*bool, bool)`

GetIsFunctionOk returns a tuple with the IsFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFunction

`func (o *ImportedFunctionDetailOutputBody) SetIsFunction(v bool)`

SetIsFunction sets IsFunction field to given value.


### GetLibraryName

`func (o *ImportedFunctionDetailOutputBody) GetLibraryName() string`

GetLibraryName returns the LibraryName field if non-nil, zero value otherwise.

### GetLibraryNameOk

`func (o *ImportedFunctionDetailOutputBody) GetLibraryNameOk() (*string, bool)`

GetLibraryNameOk returns a tuple with the LibraryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryName

`func (o *ImportedFunctionDetailOutputBody) SetLibraryName(v string)`

SetLibraryName sets LibraryName field to given value.


### GetLibraryVersion

`func (o *ImportedFunctionDetailOutputBody) GetLibraryVersion() string`

GetLibraryVersion returns the LibraryVersion field if non-nil, zero value otherwise.

### GetLibraryVersionOk

`func (o *ImportedFunctionDetailOutputBody) GetLibraryVersionOk() (*string, bool)`

GetLibraryVersionOk returns a tuple with the LibraryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryVersion

`func (o *ImportedFunctionDetailOutputBody) SetLibraryVersion(v string)`

SetLibraryVersion sets LibraryVersion field to given value.

### HasLibraryVersion

`func (o *ImportedFunctionDetailOutputBody) HasLibraryVersion() bool`

HasLibraryVersion returns a boolean if a field has been set.

### GetName

`func (o *ImportedFunctionDetailOutputBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImportedFunctionDetailOutputBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImportedFunctionDetailOutputBody) SetName(v string)`

SetName sets Name field to given value.


### GetOriginalName

`func (o *ImportedFunctionDetailOutputBody) GetOriginalName() string`

GetOriginalName returns the OriginalName field if non-nil, zero value otherwise.

### GetOriginalNameOk

`func (o *ImportedFunctionDetailOutputBody) GetOriginalNameOk() (*string, bool)`

GetOriginalNameOk returns a tuple with the OriginalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalName

`func (o *ImportedFunctionDetailOutputBody) SetOriginalName(v string)`

SetOriginalName sets OriginalName field to given value.

### HasOriginalName

`func (o *ImportedFunctionDetailOutputBody) HasOriginalName() bool`

HasOriginalName returns a boolean if a field has been set.

### GetStubVaddrs

`func (o *ImportedFunctionDetailOutputBody) GetStubVaddrs() []int64`

GetStubVaddrs returns the StubVaddrs field if non-nil, zero value otherwise.

### GetStubVaddrsOk

`func (o *ImportedFunctionDetailOutputBody) GetStubVaddrsOk() (*[]int64, bool)`

GetStubVaddrsOk returns a tuple with the StubVaddrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStubVaddrs

`func (o *ImportedFunctionDetailOutputBody) SetStubVaddrs(v []int64)`

SetStubVaddrs sets StubVaddrs field to given value.


### SetStubVaddrsNil

`func (o *ImportedFunctionDetailOutputBody) SetStubVaddrsNil(b bool)`

 SetStubVaddrsNil sets the value for StubVaddrs to be an explicit nil

### UnsetStubVaddrs
`func (o *ImportedFunctionDetailOutputBody) UnsetStubVaddrs()`

UnsetStubVaddrs ensures that no value is present for StubVaddrs, not even an explicit nil
### GetVaddr

`func (o *ImportedFunctionDetailOutputBody) GetVaddr() int64`

GetVaddr returns the Vaddr field if non-nil, zero value otherwise.

### GetVaddrOk

`func (o *ImportedFunctionDetailOutputBody) GetVaddrOk() (*int64, bool)`

GetVaddrOk returns a tuple with the Vaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaddr

`func (o *ImportedFunctionDetailOutputBody) SetVaddr(v int64)`

SetVaddr sets Vaddr field to given value.

### HasVaddr

`func (o *ImportedFunctionDetailOutputBody) HasVaddr() bool`

HasVaddr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


