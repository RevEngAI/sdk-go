# ImportedFunctionCallerEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionId** | **int64** |  | 
**FunctionName** | **string** |  | 
**FunctionVaddr** | **int64** |  | 
**StubVaddr** | **int64** | The PLT/stub address this caller targets. | 

## Methods

### NewImportedFunctionCallerEntry

`func NewImportedFunctionCallerEntry(functionId int64, functionName string, functionVaddr int64, stubVaddr int64, ) *ImportedFunctionCallerEntry`

NewImportedFunctionCallerEntry instantiates a new ImportedFunctionCallerEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportedFunctionCallerEntryWithDefaults

`func NewImportedFunctionCallerEntryWithDefaults() *ImportedFunctionCallerEntry`

NewImportedFunctionCallerEntryWithDefaults instantiates a new ImportedFunctionCallerEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionId

`func (o *ImportedFunctionCallerEntry) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *ImportedFunctionCallerEntry) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *ImportedFunctionCallerEntry) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetFunctionName

`func (o *ImportedFunctionCallerEntry) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *ImportedFunctionCallerEntry) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *ImportedFunctionCallerEntry) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetFunctionVaddr

`func (o *ImportedFunctionCallerEntry) GetFunctionVaddr() int64`

GetFunctionVaddr returns the FunctionVaddr field if non-nil, zero value otherwise.

### GetFunctionVaddrOk

`func (o *ImportedFunctionCallerEntry) GetFunctionVaddrOk() (*int64, bool)`

GetFunctionVaddrOk returns a tuple with the FunctionVaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionVaddr

`func (o *ImportedFunctionCallerEntry) SetFunctionVaddr(v int64)`

SetFunctionVaddr sets FunctionVaddr field to given value.


### GetStubVaddr

`func (o *ImportedFunctionCallerEntry) GetStubVaddr() int64`

GetStubVaddr returns the StubVaddr field if non-nil, zero value otherwise.

### GetStubVaddrOk

`func (o *ImportedFunctionCallerEntry) GetStubVaddrOk() (*int64, bool)`

GetStubVaddrOk returns a tuple with the StubVaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStubVaddr

`func (o *ImportedFunctionCallerEntry) SetStubVaddr(v int64)`

SetStubVaddr sets StubVaddr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


