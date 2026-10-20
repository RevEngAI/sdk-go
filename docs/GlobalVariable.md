# GlobalVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DType** | Pointer to **string** | Name of the assigned data type. | [optional] 
**DataTypeId** | Pointer to **int64** | Analysis data type, when one is assigned. | [optional] 
**FunctionCount** | **int64** | How many functions&#39; AI decompilations reference the global. | 
**IndirectTarget** | Pointer to **int64** | Address a pointer slot resolves to. | [optional] 
**IsBss** | **bool** | Whether the address falls in a zero-initialised section. | 
**Name** | Pointer to **string** | Recovered name, absent when nothing named the address. | [optional] 
**NameSource** | Pointer to **string** | What named it. | [optional] 
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

### NewGlobalVariable

`func NewGlobalVariable(functionCount int64, isBss bool, totalRefCount int64, vaddr int64, ) *GlobalVariable`

NewGlobalVariable instantiates a new GlobalVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalVariableWithDefaults

`func NewGlobalVariableWithDefaults() *GlobalVariable`

NewGlobalVariableWithDefaults instantiates a new GlobalVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDType

`func (o *GlobalVariable) GetDType() string`

GetDType returns the DType field if non-nil, zero value otherwise.

### GetDTypeOk

`func (o *GlobalVariable) GetDTypeOk() (*string, bool)`

GetDTypeOk returns a tuple with the DType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDType

`func (o *GlobalVariable) SetDType(v string)`

SetDType sets DType field to given value.

### HasDType

`func (o *GlobalVariable) HasDType() bool`

HasDType returns a boolean if a field has been set.

### GetDataTypeId

`func (o *GlobalVariable) GetDataTypeId() int64`

GetDataTypeId returns the DataTypeId field if non-nil, zero value otherwise.

### GetDataTypeIdOk

`func (o *GlobalVariable) GetDataTypeIdOk() (*int64, bool)`

GetDataTypeIdOk returns a tuple with the DataTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypeId

`func (o *GlobalVariable) SetDataTypeId(v int64)`

SetDataTypeId sets DataTypeId field to given value.

### HasDataTypeId

`func (o *GlobalVariable) HasDataTypeId() bool`

HasDataTypeId returns a boolean if a field has been set.

### GetFunctionCount

`func (o *GlobalVariable) GetFunctionCount() int64`

GetFunctionCount returns the FunctionCount field if non-nil, zero value otherwise.

### GetFunctionCountOk

`func (o *GlobalVariable) GetFunctionCountOk() (*int64, bool)`

GetFunctionCountOk returns a tuple with the FunctionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionCount

`func (o *GlobalVariable) SetFunctionCount(v int64)`

SetFunctionCount sets FunctionCount field to given value.


### GetIndirectTarget

`func (o *GlobalVariable) GetIndirectTarget() int64`

GetIndirectTarget returns the IndirectTarget field if non-nil, zero value otherwise.

### GetIndirectTargetOk

`func (o *GlobalVariable) GetIndirectTargetOk() (*int64, bool)`

GetIndirectTargetOk returns a tuple with the IndirectTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndirectTarget

`func (o *GlobalVariable) SetIndirectTarget(v int64)`

SetIndirectTarget sets IndirectTarget field to given value.

### HasIndirectTarget

`func (o *GlobalVariable) HasIndirectTarget() bool`

HasIndirectTarget returns a boolean if a field has been set.

### GetIsBss

`func (o *GlobalVariable) GetIsBss() bool`

GetIsBss returns the IsBss field if non-nil, zero value otherwise.

### GetIsBssOk

`func (o *GlobalVariable) GetIsBssOk() (*bool, bool)`

GetIsBssOk returns a tuple with the IsBss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBss

`func (o *GlobalVariable) SetIsBss(v bool)`

SetIsBss sets IsBss field to given value.


### GetName

`func (o *GlobalVariable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlobalVariable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlobalVariable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GlobalVariable) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameSource

`func (o *GlobalVariable) GetNameSource() string`

GetNameSource returns the NameSource field if non-nil, zero value otherwise.

### GetNameSourceOk

`func (o *GlobalVariable) GetNameSourceOk() (*string, bool)`

GetNameSourceOk returns a tuple with the NameSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSource

`func (o *GlobalVariable) SetNameSource(v string)`

SetNameSource sets NameSource field to given value.

### HasNameSource

`func (o *GlobalVariable) HasNameSource() bool`

HasNameSource returns a boolean if a field has been set.

### GetSize

`func (o *GlobalVariable) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GlobalVariable) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GlobalVariable) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *GlobalVariable) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSourceFunctionId

`func (o *GlobalVariable) GetSourceFunctionId() int64`

GetSourceFunctionId returns the SourceFunctionId field if non-nil, zero value otherwise.

### GetSourceFunctionIdOk

`func (o *GlobalVariable) GetSourceFunctionIdOk() (*int64, bool)`

GetSourceFunctionIdOk returns a tuple with the SourceFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFunctionId

`func (o *GlobalVariable) SetSourceFunctionId(v int64)`

SetSourceFunctionId sets SourceFunctionId field to given value.

### HasSourceFunctionId

`func (o *GlobalVariable) HasSourceFunctionId() bool`

HasSourceFunctionId returns a boolean if a field has been set.

### GetTotalRefCount

`func (o *GlobalVariable) GetTotalRefCount() int64`

GetTotalRefCount returns the TotalRefCount field if non-nil, zero value otherwise.

### GetTotalRefCountOk

`func (o *GlobalVariable) GetTotalRefCountOk() (*int64, bool)`

GetTotalRefCountOk returns a tuple with the TotalRefCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRefCount

`func (o *GlobalVariable) SetTotalRefCount(v int64)`

SetTotalRefCount sets TotalRefCount field to given value.


### GetTypeSource

`func (o *GlobalVariable) GetTypeSource() string`

GetTypeSource returns the TypeSource field if non-nil, zero value otherwise.

### GetTypeSourceOk

`func (o *GlobalVariable) GetTypeSourceOk() (*string, bool)`

GetTypeSourceOk returns a tuple with the TypeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeSource

`func (o *GlobalVariable) SetTypeSource(v string)`

SetTypeSource sets TypeSource field to given value.

### HasTypeSource

`func (o *GlobalVariable) HasTypeSource() bool`

HasTypeSource returns a boolean if a field has been set.

### GetVaddr

`func (o *GlobalVariable) GetVaddr() int64`

GetVaddr returns the Vaddr field if non-nil, zero value otherwise.

### GetVaddrOk

`func (o *GlobalVariable) GetVaddrOk() (*int64, bool)`

GetVaddrOk returns a tuple with the Vaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaddr

`func (o *GlobalVariable) SetVaddr(v int64)`

SetVaddr sets Vaddr field to given value.


### GetValue

`func (o *GlobalVariable) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GlobalVariable) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GlobalVariable) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *GlobalVariable) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueConfidence

`func (o *GlobalVariable) GetValueConfidence() string`

GetValueConfidence returns the ValueConfidence field if non-nil, zero value otherwise.

### GetValueConfidenceOk

`func (o *GlobalVariable) GetValueConfidenceOk() (*string, bool)`

GetValueConfidenceOk returns a tuple with the ValueConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueConfidence

`func (o *GlobalVariable) SetValueConfidence(v string)`

SetValueConfidence sets ValueConfidence field to given value.

### HasValueConfidence

`func (o *GlobalVariable) HasValueConfidence() bool`

HasValueConfidence returns a boolean if a field has been set.

### GetValueKind

`func (o *GlobalVariable) GetValueKind() string`

GetValueKind returns the ValueKind field if non-nil, zero value otherwise.

### GetValueKindOk

`func (o *GlobalVariable) GetValueKindOk() (*string, bool)`

GetValueKindOk returns a tuple with the ValueKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueKind

`func (o *GlobalVariable) SetValueKind(v string)`

SetValueKind sets ValueKind field to given value.

### HasValueKind

`func (o *GlobalVariable) HasValueKind() bool`

HasValueKind returns a boolean if a field has been set.

### GetValueSource

`func (o *GlobalVariable) GetValueSource() string`

GetValueSource returns the ValueSource field if non-nil, zero value otherwise.

### GetValueSourceOk

`func (o *GlobalVariable) GetValueSourceOk() (*string, bool)`

GetValueSourceOk returns a tuple with the ValueSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueSource

`func (o *GlobalVariable) SetValueSource(v string)`

SetValueSource sets ValueSource field to given value.

### HasValueSource

`func (o *GlobalVariable) HasValueSource() bool`

HasValueSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


