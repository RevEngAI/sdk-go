# FunctionDetailsOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisId** | **int64** |  | 
**BinaryId** | **int64** |  | 
**Creation** | **time.Time** |  | 
**Debug** | **bool** |  | 
**FunctionId** | **int64** |  | 
**FunctionName** | **string** |  | 
**FunctionSize** | **int64** |  | 
**FunctionVaddr** | **int64** |  | 
**MangledName** | Pointer to **string** |  | [optional] 
**SourceFunctionId** | Pointer to **int64** |  | [optional] 

## Methods

### NewFunctionDetailsOutputBody

`func NewFunctionDetailsOutputBody(analysisId int64, binaryId int64, creation time.Time, debug bool, functionId int64, functionName string, functionSize int64, functionVaddr int64, ) *FunctionDetailsOutputBody`

NewFunctionDetailsOutputBody instantiates a new FunctionDetailsOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionDetailsOutputBodyWithDefaults

`func NewFunctionDetailsOutputBodyWithDefaults() *FunctionDetailsOutputBody`

NewFunctionDetailsOutputBodyWithDefaults instantiates a new FunctionDetailsOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisId

`func (o *FunctionDetailsOutputBody) GetAnalysisId() int64`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *FunctionDetailsOutputBody) GetAnalysisIdOk() (*int64, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *FunctionDetailsOutputBody) SetAnalysisId(v int64)`

SetAnalysisId sets AnalysisId field to given value.


### GetBinaryId

`func (o *FunctionDetailsOutputBody) GetBinaryId() int64`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *FunctionDetailsOutputBody) GetBinaryIdOk() (*int64, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *FunctionDetailsOutputBody) SetBinaryId(v int64)`

SetBinaryId sets BinaryId field to given value.


### GetCreation

`func (o *FunctionDetailsOutputBody) GetCreation() time.Time`

GetCreation returns the Creation field if non-nil, zero value otherwise.

### GetCreationOk

`func (o *FunctionDetailsOutputBody) GetCreationOk() (*time.Time, bool)`

GetCreationOk returns a tuple with the Creation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreation

`func (o *FunctionDetailsOutputBody) SetCreation(v time.Time)`

SetCreation sets Creation field to given value.


### GetDebug

`func (o *FunctionDetailsOutputBody) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *FunctionDetailsOutputBody) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *FunctionDetailsOutputBody) SetDebug(v bool)`

SetDebug sets Debug field to given value.


### GetFunctionId

`func (o *FunctionDetailsOutputBody) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *FunctionDetailsOutputBody) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *FunctionDetailsOutputBody) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetFunctionName

`func (o *FunctionDetailsOutputBody) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *FunctionDetailsOutputBody) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *FunctionDetailsOutputBody) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetFunctionSize

`func (o *FunctionDetailsOutputBody) GetFunctionSize() int64`

GetFunctionSize returns the FunctionSize field if non-nil, zero value otherwise.

### GetFunctionSizeOk

`func (o *FunctionDetailsOutputBody) GetFunctionSizeOk() (*int64, bool)`

GetFunctionSizeOk returns a tuple with the FunctionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionSize

`func (o *FunctionDetailsOutputBody) SetFunctionSize(v int64)`

SetFunctionSize sets FunctionSize field to given value.


### GetFunctionVaddr

`func (o *FunctionDetailsOutputBody) GetFunctionVaddr() int64`

GetFunctionVaddr returns the FunctionVaddr field if non-nil, zero value otherwise.

### GetFunctionVaddrOk

`func (o *FunctionDetailsOutputBody) GetFunctionVaddrOk() (*int64, bool)`

GetFunctionVaddrOk returns a tuple with the FunctionVaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionVaddr

`func (o *FunctionDetailsOutputBody) SetFunctionVaddr(v int64)`

SetFunctionVaddr sets FunctionVaddr field to given value.


### GetMangledName

`func (o *FunctionDetailsOutputBody) GetMangledName() string`

GetMangledName returns the MangledName field if non-nil, zero value otherwise.

### GetMangledNameOk

`func (o *FunctionDetailsOutputBody) GetMangledNameOk() (*string, bool)`

GetMangledNameOk returns a tuple with the MangledName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMangledName

`func (o *FunctionDetailsOutputBody) SetMangledName(v string)`

SetMangledName sets MangledName field to given value.

### HasMangledName

`func (o *FunctionDetailsOutputBody) HasMangledName() bool`

HasMangledName returns a boolean if a field has been set.

### GetSourceFunctionId

`func (o *FunctionDetailsOutputBody) GetSourceFunctionId() int64`

GetSourceFunctionId returns the SourceFunctionId field if non-nil, zero value otherwise.

### GetSourceFunctionIdOk

`func (o *FunctionDetailsOutputBody) GetSourceFunctionIdOk() (*int64, bool)`

GetSourceFunctionIdOk returns a tuple with the SourceFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFunctionId

`func (o *FunctionDetailsOutputBody) SetSourceFunctionId(v int64)`

SetSourceFunctionId sets SourceFunctionId field to given value.

### HasSourceFunctionId

`func (o *FunctionDetailsOutputBody) HasSourceFunctionId() bool`

HasSourceFunctionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


