# FunctionGlobalVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DType** | Pointer to **string** | Name of the assigned data type. | [optional] 
**DataTypeId** | Pointer to **int64** | Analysis data type, when one is assigned. | [optional] 
**DecompLines** | **[]int32** | 1-indexed lines of the function&#39;s AI decompilation that reference the global. | 
**FunctionCount** | **int64** | How many functions&#39; AI decompilations reference the global. | 
**IndirectTarget** | Pointer to **int64** | Address a pointer slot resolves to. | [optional] 
**IsBss** | **bool** | Whether the address falls in a zero-initialised section. | 
**Name** | Pointer to **string** | Recovered name, absent when nothing named the address. | [optional] 
**NameSource** | Pointer to **string** | What named it. | [optional] 
**RefCount** | **int64** | How many times the function&#39;s AI decompilation references the global. | 
**Size** | Pointer to **int64** | Size in bytes, when known. | [optional] 
**SourceFunctionId** | Pointer to **int64** | Function a constructed static was initialised in. | [optional] 
**TotalRefCount** | **int64** | How many times those AI decompilations reference the global, summed across functions. | 
**TypeSource** | Pointer to **string** | What typed it. | [optional] 
**Vaddr** | **int64** | Virtual address the global lives at. | 
**Value** | Pointer to **string** | Decoded initial value, rendered for display. | [optional] 
**ValueConfidence** | Pointer to **string** | How much to trust the decoded value. | [optional] 
**ValueKind** | Pointer to **string** | How the value was interpreted. | [optional] 
**ValueSource** | Pointer to **string** | What produced the value. | [optional] 

## Methods

### NewFunctionGlobalVariable

`func NewFunctionGlobalVariable(decompLines []int32, functionCount int64, isBss bool, refCount int64, totalRefCount int64, vaddr int64, ) *FunctionGlobalVariable`

NewFunctionGlobalVariable instantiates a new FunctionGlobalVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionGlobalVariableWithDefaults

`func NewFunctionGlobalVariableWithDefaults() *FunctionGlobalVariable`

NewFunctionGlobalVariableWithDefaults instantiates a new FunctionGlobalVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDType

`func (o *FunctionGlobalVariable) GetDType() string`

GetDType returns the DType field if non-nil, zero value otherwise.

### GetDTypeOk

`func (o *FunctionGlobalVariable) GetDTypeOk() (*string, bool)`

GetDTypeOk returns a tuple with the DType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDType

`func (o *FunctionGlobalVariable) SetDType(v string)`

SetDType sets DType field to given value.

### HasDType

`func (o *FunctionGlobalVariable) HasDType() bool`

HasDType returns a boolean if a field has been set.

### GetDataTypeId

`func (o *FunctionGlobalVariable) GetDataTypeId() int64`

GetDataTypeId returns the DataTypeId field if non-nil, zero value otherwise.

### GetDataTypeIdOk

`func (o *FunctionGlobalVariable) GetDataTypeIdOk() (*int64, bool)`

GetDataTypeIdOk returns a tuple with the DataTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypeId

`func (o *FunctionGlobalVariable) SetDataTypeId(v int64)`

SetDataTypeId sets DataTypeId field to given value.

### HasDataTypeId

`func (o *FunctionGlobalVariable) HasDataTypeId() bool`

HasDataTypeId returns a boolean if a field has been set.

### GetDecompLines

`func (o *FunctionGlobalVariable) GetDecompLines() []int32`

GetDecompLines returns the DecompLines field if non-nil, zero value otherwise.

### GetDecompLinesOk

`func (o *FunctionGlobalVariable) GetDecompLinesOk() (*[]int32, bool)`

GetDecompLinesOk returns a tuple with the DecompLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecompLines

`func (o *FunctionGlobalVariable) SetDecompLines(v []int32)`

SetDecompLines sets DecompLines field to given value.


### SetDecompLinesNil

`func (o *FunctionGlobalVariable) SetDecompLinesNil(b bool)`

 SetDecompLinesNil sets the value for DecompLines to be an explicit nil

### UnsetDecompLines
`func (o *FunctionGlobalVariable) UnsetDecompLines()`

UnsetDecompLines ensures that no value is present for DecompLines, not even an explicit nil
### GetFunctionCount

`func (o *FunctionGlobalVariable) GetFunctionCount() int64`

GetFunctionCount returns the FunctionCount field if non-nil, zero value otherwise.

### GetFunctionCountOk

`func (o *FunctionGlobalVariable) GetFunctionCountOk() (*int64, bool)`

GetFunctionCountOk returns a tuple with the FunctionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionCount

`func (o *FunctionGlobalVariable) SetFunctionCount(v int64)`

SetFunctionCount sets FunctionCount field to given value.


### GetIndirectTarget

`func (o *FunctionGlobalVariable) GetIndirectTarget() int64`

GetIndirectTarget returns the IndirectTarget field if non-nil, zero value otherwise.

### GetIndirectTargetOk

`func (o *FunctionGlobalVariable) GetIndirectTargetOk() (*int64, bool)`

GetIndirectTargetOk returns a tuple with the IndirectTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndirectTarget

`func (o *FunctionGlobalVariable) SetIndirectTarget(v int64)`

SetIndirectTarget sets IndirectTarget field to given value.

### HasIndirectTarget

`func (o *FunctionGlobalVariable) HasIndirectTarget() bool`

HasIndirectTarget returns a boolean if a field has been set.

### GetIsBss

`func (o *FunctionGlobalVariable) GetIsBss() bool`

GetIsBss returns the IsBss field if non-nil, zero value otherwise.

### GetIsBssOk

`func (o *FunctionGlobalVariable) GetIsBssOk() (*bool, bool)`

GetIsBssOk returns a tuple with the IsBss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBss

`func (o *FunctionGlobalVariable) SetIsBss(v bool)`

SetIsBss sets IsBss field to given value.


### GetName

`func (o *FunctionGlobalVariable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionGlobalVariable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionGlobalVariable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FunctionGlobalVariable) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameSource

`func (o *FunctionGlobalVariable) GetNameSource() string`

GetNameSource returns the NameSource field if non-nil, zero value otherwise.

### GetNameSourceOk

`func (o *FunctionGlobalVariable) GetNameSourceOk() (*string, bool)`

GetNameSourceOk returns a tuple with the NameSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSource

`func (o *FunctionGlobalVariable) SetNameSource(v string)`

SetNameSource sets NameSource field to given value.

### HasNameSource

`func (o *FunctionGlobalVariable) HasNameSource() bool`

HasNameSource returns a boolean if a field has been set.

### GetRefCount

`func (o *FunctionGlobalVariable) GetRefCount() int64`

GetRefCount returns the RefCount field if non-nil, zero value otherwise.

### GetRefCountOk

`func (o *FunctionGlobalVariable) GetRefCountOk() (*int64, bool)`

GetRefCountOk returns a tuple with the RefCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefCount

`func (o *FunctionGlobalVariable) SetRefCount(v int64)`

SetRefCount sets RefCount field to given value.


### GetSize

`func (o *FunctionGlobalVariable) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FunctionGlobalVariable) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FunctionGlobalVariable) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *FunctionGlobalVariable) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSourceFunctionId

`func (o *FunctionGlobalVariable) GetSourceFunctionId() int64`

GetSourceFunctionId returns the SourceFunctionId field if non-nil, zero value otherwise.

### GetSourceFunctionIdOk

`func (o *FunctionGlobalVariable) GetSourceFunctionIdOk() (*int64, bool)`

GetSourceFunctionIdOk returns a tuple with the SourceFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFunctionId

`func (o *FunctionGlobalVariable) SetSourceFunctionId(v int64)`

SetSourceFunctionId sets SourceFunctionId field to given value.

### HasSourceFunctionId

`func (o *FunctionGlobalVariable) HasSourceFunctionId() bool`

HasSourceFunctionId returns a boolean if a field has been set.

### GetTotalRefCount

`func (o *FunctionGlobalVariable) GetTotalRefCount() int64`

GetTotalRefCount returns the TotalRefCount field if non-nil, zero value otherwise.

### GetTotalRefCountOk

`func (o *FunctionGlobalVariable) GetTotalRefCountOk() (*int64, bool)`

GetTotalRefCountOk returns a tuple with the TotalRefCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRefCount

`func (o *FunctionGlobalVariable) SetTotalRefCount(v int64)`

SetTotalRefCount sets TotalRefCount field to given value.


### GetTypeSource

`func (o *FunctionGlobalVariable) GetTypeSource() string`

GetTypeSource returns the TypeSource field if non-nil, zero value otherwise.

### GetTypeSourceOk

`func (o *FunctionGlobalVariable) GetTypeSourceOk() (*string, bool)`

GetTypeSourceOk returns a tuple with the TypeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeSource

`func (o *FunctionGlobalVariable) SetTypeSource(v string)`

SetTypeSource sets TypeSource field to given value.

### HasTypeSource

`func (o *FunctionGlobalVariable) HasTypeSource() bool`

HasTypeSource returns a boolean if a field has been set.

### GetVaddr

`func (o *FunctionGlobalVariable) GetVaddr() int64`

GetVaddr returns the Vaddr field if non-nil, zero value otherwise.

### GetVaddrOk

`func (o *FunctionGlobalVariable) GetVaddrOk() (*int64, bool)`

GetVaddrOk returns a tuple with the Vaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaddr

`func (o *FunctionGlobalVariable) SetVaddr(v int64)`

SetVaddr sets Vaddr field to given value.


### GetValue

`func (o *FunctionGlobalVariable) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FunctionGlobalVariable) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FunctionGlobalVariable) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *FunctionGlobalVariable) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueConfidence

`func (o *FunctionGlobalVariable) GetValueConfidence() string`

GetValueConfidence returns the ValueConfidence field if non-nil, zero value otherwise.

### GetValueConfidenceOk

`func (o *FunctionGlobalVariable) GetValueConfidenceOk() (*string, bool)`

GetValueConfidenceOk returns a tuple with the ValueConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueConfidence

`func (o *FunctionGlobalVariable) SetValueConfidence(v string)`

SetValueConfidence sets ValueConfidence field to given value.

### HasValueConfidence

`func (o *FunctionGlobalVariable) HasValueConfidence() bool`

HasValueConfidence returns a boolean if a field has been set.

### GetValueKind

`func (o *FunctionGlobalVariable) GetValueKind() string`

GetValueKind returns the ValueKind field if non-nil, zero value otherwise.

### GetValueKindOk

`func (o *FunctionGlobalVariable) GetValueKindOk() (*string, bool)`

GetValueKindOk returns a tuple with the ValueKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueKind

`func (o *FunctionGlobalVariable) SetValueKind(v string)`

SetValueKind sets ValueKind field to given value.

### HasValueKind

`func (o *FunctionGlobalVariable) HasValueKind() bool`

HasValueKind returns a boolean if a field has been set.

### GetValueSource

`func (o *FunctionGlobalVariable) GetValueSource() string`

GetValueSource returns the ValueSource field if non-nil, zero value otherwise.

### GetValueSourceOk

`func (o *FunctionGlobalVariable) GetValueSourceOk() (*string, bool)`

GetValueSourceOk returns a tuple with the ValueSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueSource

`func (o *FunctionGlobalVariable) SetValueSource(v string)`

SetValueSource sets ValueSource field to given value.

### HasValueSource

`func (o *FunctionGlobalVariable) HasValueSource() bool`

HasValueSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


