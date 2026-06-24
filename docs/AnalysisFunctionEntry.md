# AnalysisFunctionEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinaryId** | **int64** |  | 
**Debug** | **bool** |  | 
**FunctionId** | **int64** |  | 
**FunctionName** | **string** |  | 
**FunctionSize** | **int64** |  | 
**FunctionVaddr** | **int64** |  | 
**MangledName** | Pointer to **NullableString** |  | [optional] 
**SourceBinaryId** | Pointer to **int64** |  | [optional] 
**SourceType** | **string** |  | 

## Methods

### NewAnalysisFunctionEntry

`func NewAnalysisFunctionEntry(binaryId int64, debug bool, functionId int64, functionName string, functionSize int64, functionVaddr int64, sourceType string, ) *AnalysisFunctionEntry`

NewAnalysisFunctionEntry instantiates a new AnalysisFunctionEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisFunctionEntryWithDefaults

`func NewAnalysisFunctionEntryWithDefaults() *AnalysisFunctionEntry`

NewAnalysisFunctionEntryWithDefaults instantiates a new AnalysisFunctionEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaryId

`func (o *AnalysisFunctionEntry) GetBinaryId() int64`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *AnalysisFunctionEntry) GetBinaryIdOk() (*int64, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *AnalysisFunctionEntry) SetBinaryId(v int64)`

SetBinaryId sets BinaryId field to given value.


### GetDebug

`func (o *AnalysisFunctionEntry) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *AnalysisFunctionEntry) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *AnalysisFunctionEntry) SetDebug(v bool)`

SetDebug sets Debug field to given value.


### GetFunctionId

`func (o *AnalysisFunctionEntry) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *AnalysisFunctionEntry) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *AnalysisFunctionEntry) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetFunctionName

`func (o *AnalysisFunctionEntry) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *AnalysisFunctionEntry) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *AnalysisFunctionEntry) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetFunctionSize

`func (o *AnalysisFunctionEntry) GetFunctionSize() int64`

GetFunctionSize returns the FunctionSize field if non-nil, zero value otherwise.

### GetFunctionSizeOk

`func (o *AnalysisFunctionEntry) GetFunctionSizeOk() (*int64, bool)`

GetFunctionSizeOk returns a tuple with the FunctionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionSize

`func (o *AnalysisFunctionEntry) SetFunctionSize(v int64)`

SetFunctionSize sets FunctionSize field to given value.


### GetFunctionVaddr

`func (o *AnalysisFunctionEntry) GetFunctionVaddr() int64`

GetFunctionVaddr returns the FunctionVaddr field if non-nil, zero value otherwise.

### GetFunctionVaddrOk

`func (o *AnalysisFunctionEntry) GetFunctionVaddrOk() (*int64, bool)`

GetFunctionVaddrOk returns a tuple with the FunctionVaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionVaddr

`func (o *AnalysisFunctionEntry) SetFunctionVaddr(v int64)`

SetFunctionVaddr sets FunctionVaddr field to given value.


### GetMangledName

`func (o *AnalysisFunctionEntry) GetMangledName() string`

GetMangledName returns the MangledName field if non-nil, zero value otherwise.

### GetMangledNameOk

`func (o *AnalysisFunctionEntry) GetMangledNameOk() (*string, bool)`

GetMangledNameOk returns a tuple with the MangledName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMangledName

`func (o *AnalysisFunctionEntry) SetMangledName(v string)`

SetMangledName sets MangledName field to given value.

### HasMangledName

`func (o *AnalysisFunctionEntry) HasMangledName() bool`

HasMangledName returns a boolean if a field has been set.

### SetMangledNameNil

`func (o *AnalysisFunctionEntry) SetMangledNameNil(b bool)`

 SetMangledNameNil sets the value for MangledName to be an explicit nil

### UnsetMangledName
`func (o *AnalysisFunctionEntry) UnsetMangledName()`

UnsetMangledName ensures that no value is present for MangledName, not even an explicit nil
### GetSourceBinaryId

`func (o *AnalysisFunctionEntry) GetSourceBinaryId() int64`

GetSourceBinaryId returns the SourceBinaryId field if non-nil, zero value otherwise.

### GetSourceBinaryIdOk

`func (o *AnalysisFunctionEntry) GetSourceBinaryIdOk() (*int64, bool)`

GetSourceBinaryIdOk returns a tuple with the SourceBinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBinaryId

`func (o *AnalysisFunctionEntry) SetSourceBinaryId(v int64)`

SetSourceBinaryId sets SourceBinaryId field to given value.

### HasSourceBinaryId

`func (o *AnalysisFunctionEntry) HasSourceBinaryId() bool`

HasSourceBinaryId returns a boolean if a field has been set.

### GetSourceType

`func (o *AnalysisFunctionEntry) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *AnalysisFunctionEntry) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *AnalysisFunctionEntry) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


