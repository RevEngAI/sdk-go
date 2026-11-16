# Function

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionId** | **int64** | Function id | 
**FunctionName** | **string** | Demangled name of the function | 
**FunctionMangledName** | **string** | Mangled name of the function | 
**FunctionVaddr** | **int64** | Function virtual address | 
**FunctionSize** | **int32** | Function size | 
**Debug** | **bool** | Whether the function is debug | 

## Methods

### NewFunction

`func NewFunction(functionId int64, functionName string, functionMangledName string, functionVaddr int64, functionSize int32, debug bool, ) *Function`

NewFunction instantiates a new Function object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionWithDefaults

`func NewFunctionWithDefaults() *Function`

NewFunctionWithDefaults instantiates a new Function object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionId

`func (o *Function) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *Function) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *Function) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetFunctionName

`func (o *Function) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *Function) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *Function) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetFunctionMangledName

`func (o *Function) GetFunctionMangledName() string`

GetFunctionMangledName returns the FunctionMangledName field if non-nil, zero value otherwise.

### GetFunctionMangledNameOk

`func (o *Function) GetFunctionMangledNameOk() (*string, bool)`

GetFunctionMangledNameOk returns a tuple with the FunctionMangledName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionMangledName

`func (o *Function) SetFunctionMangledName(v string)`

SetFunctionMangledName sets FunctionMangledName field to given value.


### GetFunctionVaddr

`func (o *Function) GetFunctionVaddr() int64`

GetFunctionVaddr returns the FunctionVaddr field if non-nil, zero value otherwise.

### GetFunctionVaddrOk

`func (o *Function) GetFunctionVaddrOk() (*int64, bool)`

GetFunctionVaddrOk returns a tuple with the FunctionVaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionVaddr

`func (o *Function) SetFunctionVaddr(v int64)`

SetFunctionVaddr sets FunctionVaddr field to given value.


### GetFunctionSize

`func (o *Function) GetFunctionSize() int32`

GetFunctionSize returns the FunctionSize field if non-nil, zero value otherwise.

### GetFunctionSizeOk

`func (o *Function) GetFunctionSizeOk() (*int32, bool)`

GetFunctionSizeOk returns a tuple with the FunctionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionSize

`func (o *Function) SetFunctionSize(v int32)`

SetFunctionSize sets FunctionSize field to given value.


### GetDebug

`func (o *Function) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *Function) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *Function) SetDebug(v bool)`

SetDebug sets Debug field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


