# AppApiRestV2FunctionsTypesFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionId** | **int64** | Function id | 
**FunctionName** | **string** | Demangled name of the function | 
**FunctionMangledName** | **string** | Mangled name of the function | 
**FunctionVaddr** | **int64** | Function virtual address | 
**FunctionSize** | **int32** | Function size | 
**Debug** | **bool** | Whether the function is debug | 
**Embedding3d** | Pointer to **[]float32** |  | [optional] 
**Embedding1d** | Pointer to **[]float32** |  | [optional] 

## Methods

### NewAppApiRestV2FunctionsTypesFunction

`func NewAppApiRestV2FunctionsTypesFunction(functionId int64, functionName string, functionMangledName string, functionVaddr int64, functionSize int32, debug bool, ) *AppApiRestV2FunctionsTypesFunction`

NewAppApiRestV2FunctionsTypesFunction instantiates a new AppApiRestV2FunctionsTypesFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppApiRestV2FunctionsTypesFunctionWithDefaults

`func NewAppApiRestV2FunctionsTypesFunctionWithDefaults() *AppApiRestV2FunctionsTypesFunction`

NewAppApiRestV2FunctionsTypesFunctionWithDefaults instantiates a new AppApiRestV2FunctionsTypesFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionId

`func (o *AppApiRestV2FunctionsTypesFunction) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *AppApiRestV2FunctionsTypesFunction) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *AppApiRestV2FunctionsTypesFunction) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetFunctionName

`func (o *AppApiRestV2FunctionsTypesFunction) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *AppApiRestV2FunctionsTypesFunction) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *AppApiRestV2FunctionsTypesFunction) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetFunctionMangledName

`func (o *AppApiRestV2FunctionsTypesFunction) GetFunctionMangledName() string`

GetFunctionMangledName returns the FunctionMangledName field if non-nil, zero value otherwise.

### GetFunctionMangledNameOk

`func (o *AppApiRestV2FunctionsTypesFunction) GetFunctionMangledNameOk() (*string, bool)`

GetFunctionMangledNameOk returns a tuple with the FunctionMangledName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionMangledName

`func (o *AppApiRestV2FunctionsTypesFunction) SetFunctionMangledName(v string)`

SetFunctionMangledName sets FunctionMangledName field to given value.


### GetFunctionVaddr

`func (o *AppApiRestV2FunctionsTypesFunction) GetFunctionVaddr() int64`

GetFunctionVaddr returns the FunctionVaddr field if non-nil, zero value otherwise.

### GetFunctionVaddrOk

`func (o *AppApiRestV2FunctionsTypesFunction) GetFunctionVaddrOk() (*int64, bool)`

GetFunctionVaddrOk returns a tuple with the FunctionVaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionVaddr

`func (o *AppApiRestV2FunctionsTypesFunction) SetFunctionVaddr(v int64)`

SetFunctionVaddr sets FunctionVaddr field to given value.


### GetFunctionSize

`func (o *AppApiRestV2FunctionsTypesFunction) GetFunctionSize() int32`

GetFunctionSize returns the FunctionSize field if non-nil, zero value otherwise.

### GetFunctionSizeOk

`func (o *AppApiRestV2FunctionsTypesFunction) GetFunctionSizeOk() (*int32, bool)`

GetFunctionSizeOk returns a tuple with the FunctionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionSize

`func (o *AppApiRestV2FunctionsTypesFunction) SetFunctionSize(v int32)`

SetFunctionSize sets FunctionSize field to given value.


### GetDebug

`func (o *AppApiRestV2FunctionsTypesFunction) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *AppApiRestV2FunctionsTypesFunction) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *AppApiRestV2FunctionsTypesFunction) SetDebug(v bool)`

SetDebug sets Debug field to given value.


### GetEmbedding3d

`func (o *AppApiRestV2FunctionsTypesFunction) GetEmbedding3d() []float32`

GetEmbedding3d returns the Embedding3d field if non-nil, zero value otherwise.

### GetEmbedding3dOk

`func (o *AppApiRestV2FunctionsTypesFunction) GetEmbedding3dOk() (*[]float32, bool)`

GetEmbedding3dOk returns a tuple with the Embedding3d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedding3d

`func (o *AppApiRestV2FunctionsTypesFunction) SetEmbedding3d(v []float32)`

SetEmbedding3d sets Embedding3d field to given value.

### HasEmbedding3d

`func (o *AppApiRestV2FunctionsTypesFunction) HasEmbedding3d() bool`

HasEmbedding3d returns a boolean if a field has been set.

### SetEmbedding3dNil

`func (o *AppApiRestV2FunctionsTypesFunction) SetEmbedding3dNil(b bool)`

 SetEmbedding3dNil sets the value for Embedding3d to be an explicit nil

### UnsetEmbedding3d
`func (o *AppApiRestV2FunctionsTypesFunction) UnsetEmbedding3d()`

UnsetEmbedding3d ensures that no value is present for Embedding3d, not even an explicit nil
### GetEmbedding1d

`func (o *AppApiRestV2FunctionsTypesFunction) GetEmbedding1d() []float32`

GetEmbedding1d returns the Embedding1d field if non-nil, zero value otherwise.

### GetEmbedding1dOk

`func (o *AppApiRestV2FunctionsTypesFunction) GetEmbedding1dOk() (*[]float32, bool)`

GetEmbedding1dOk returns a tuple with the Embedding1d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedding1d

`func (o *AppApiRestV2FunctionsTypesFunction) SetEmbedding1d(v []float32)`

SetEmbedding1d sets Embedding1d field to given value.

### HasEmbedding1d

`func (o *AppApiRestV2FunctionsTypesFunction) HasEmbedding1d() bool`

HasEmbedding1d returns a boolean if a field has been set.

### SetEmbedding1dNil

`func (o *AppApiRestV2FunctionsTypesFunction) SetEmbedding1dNil(b bool)`

 SetEmbedding1dNil sets the value for Embedding1d to be an explicit nil

### UnsetEmbedding1d
`func (o *AppApiRestV2FunctionsTypesFunction) UnsetEmbedding1d()`

UnsetEmbedding1d ensures that no value is present for Embedding1d, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


