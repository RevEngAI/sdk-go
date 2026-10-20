# FunctionUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DecompLines** | **[]int32** |  | 
**FunctionId** | **int64** |  | 
**FunctionName** | **string** |  | 
**FunctionVaddr** | **int64** |  | 
**RefCount** | **int64** |  | 

## Methods

### NewFunctionUsage

`func NewFunctionUsage(decompLines []int32, functionId int64, functionName string, functionVaddr int64, refCount int64, ) *FunctionUsage`

NewFunctionUsage instantiates a new FunctionUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionUsageWithDefaults

`func NewFunctionUsageWithDefaults() *FunctionUsage`

NewFunctionUsageWithDefaults instantiates a new FunctionUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecompLines

`func (o *FunctionUsage) GetDecompLines() []int32`

GetDecompLines returns the DecompLines field if non-nil, zero value otherwise.

### GetDecompLinesOk

`func (o *FunctionUsage) GetDecompLinesOk() (*[]int32, bool)`

GetDecompLinesOk returns a tuple with the DecompLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecompLines

`func (o *FunctionUsage) SetDecompLines(v []int32)`

SetDecompLines sets DecompLines field to given value.


### SetDecompLinesNil

`func (o *FunctionUsage) SetDecompLinesNil(b bool)`

 SetDecompLinesNil sets the value for DecompLines to be an explicit nil

### UnsetDecompLines
`func (o *FunctionUsage) UnsetDecompLines()`

UnsetDecompLines ensures that no value is present for DecompLines, not even an explicit nil
### GetFunctionId

`func (o *FunctionUsage) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *FunctionUsage) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *FunctionUsage) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetFunctionName

`func (o *FunctionUsage) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *FunctionUsage) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *FunctionUsage) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetFunctionVaddr

`func (o *FunctionUsage) GetFunctionVaddr() int64`

GetFunctionVaddr returns the FunctionVaddr field if non-nil, zero value otherwise.

### GetFunctionVaddrOk

`func (o *FunctionUsage) GetFunctionVaddrOk() (*int64, bool)`

GetFunctionVaddrOk returns a tuple with the FunctionVaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionVaddr

`func (o *FunctionUsage) SetFunctionVaddr(v int64)`

SetFunctionVaddr sets FunctionVaddr field to given value.


### GetRefCount

`func (o *FunctionUsage) GetRefCount() int64`

GetRefCount returns the RefCount field if non-nil, zero value otherwise.

### GetRefCountOk

`func (o *FunctionUsage) GetRefCountOk() (*int64, bool)`

GetRefCountOk returns a tuple with the RefCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefCount

`func (o *FunctionUsage) SetRefCount(v int64)`

SetRefCount sets RefCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


