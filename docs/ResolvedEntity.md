# ResolvedEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddrToken** | **NullableString** |  | 
**AllAddrTokens** | **[]string** |  | 
**BitOffset** | **NullableInt64** |  | 
**ByteOffset** | **NullableInt64** |  | 
**ByteSize** | **NullableInt64** |  | 
**Count** | **int64** |  | 
**DataTypeIndex** | **NullableInt64** |  | 
**FieldStatus** | Pointer to **string** |  | [optional] 
**FunctionId** | Pointer to **int64** |  | [optional] 
**ImportedFunctionId** | Pointer to **int64** |  | [optional] 
**Kind** | **NullableString** |  | 
**Name** | **NullableString** |  | 
**NameSource** | **string** |  | 
**NeedsNaming** | **bool** |  | 
**Provenance** | **NullableString** |  | 
**ResolvedName** | **NullableString** |  | 
**SuggestedName** | Pointer to **NullableString** |  | [optional] 
**SuggestedType** | **NullableString** |  | 
**SuggestionConfidence** | **NullableString** |  | 
**Token** | **NullableString** |  | 
**TypeIndex** | **NullableInt64** |  | 
**Vaddr** | **NullableInt64** |  | 
**Value** | **NullableString** |  | 
**ValueConfidence** | **NullableString** |  | 
**ValueType** | **NullableString** |  | 

## Methods

### NewResolvedEntity

`func NewResolvedEntity(addrToken NullableString, allAddrTokens []string, bitOffset NullableInt64, byteOffset NullableInt64, byteSize NullableInt64, count int64, dataTypeIndex NullableInt64, kind NullableString, name NullableString, nameSource string, needsNaming bool, provenance NullableString, resolvedName NullableString, suggestedType NullableString, suggestionConfidence NullableString, token NullableString, typeIndex NullableInt64, vaddr NullableInt64, value NullableString, valueConfidence NullableString, valueType NullableString, ) *ResolvedEntity`

NewResolvedEntity instantiates a new ResolvedEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResolvedEntityWithDefaults

`func NewResolvedEntityWithDefaults() *ResolvedEntity`

NewResolvedEntityWithDefaults instantiates a new ResolvedEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddrToken

`func (o *ResolvedEntity) GetAddrToken() string`

GetAddrToken returns the AddrToken field if non-nil, zero value otherwise.

### GetAddrTokenOk

`func (o *ResolvedEntity) GetAddrTokenOk() (*string, bool)`

GetAddrTokenOk returns a tuple with the AddrToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddrToken

`func (o *ResolvedEntity) SetAddrToken(v string)`

SetAddrToken sets AddrToken field to given value.


### SetAddrTokenNil

`func (o *ResolvedEntity) SetAddrTokenNil(b bool)`

 SetAddrTokenNil sets the value for AddrToken to be an explicit nil

### UnsetAddrToken
`func (o *ResolvedEntity) UnsetAddrToken()`

UnsetAddrToken ensures that no value is present for AddrToken, not even an explicit nil
### GetAllAddrTokens

`func (o *ResolvedEntity) GetAllAddrTokens() []string`

GetAllAddrTokens returns the AllAddrTokens field if non-nil, zero value otherwise.

### GetAllAddrTokensOk

`func (o *ResolvedEntity) GetAllAddrTokensOk() (*[]string, bool)`

GetAllAddrTokensOk returns a tuple with the AllAddrTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllAddrTokens

`func (o *ResolvedEntity) SetAllAddrTokens(v []string)`

SetAllAddrTokens sets AllAddrTokens field to given value.


### SetAllAddrTokensNil

`func (o *ResolvedEntity) SetAllAddrTokensNil(b bool)`

 SetAllAddrTokensNil sets the value for AllAddrTokens to be an explicit nil

### UnsetAllAddrTokens
`func (o *ResolvedEntity) UnsetAllAddrTokens()`

UnsetAllAddrTokens ensures that no value is present for AllAddrTokens, not even an explicit nil
### GetBitOffset

`func (o *ResolvedEntity) GetBitOffset() int64`

GetBitOffset returns the BitOffset field if non-nil, zero value otherwise.

### GetBitOffsetOk

`func (o *ResolvedEntity) GetBitOffsetOk() (*int64, bool)`

GetBitOffsetOk returns a tuple with the BitOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitOffset

`func (o *ResolvedEntity) SetBitOffset(v int64)`

SetBitOffset sets BitOffset field to given value.


### SetBitOffsetNil

`func (o *ResolvedEntity) SetBitOffsetNil(b bool)`

 SetBitOffsetNil sets the value for BitOffset to be an explicit nil

### UnsetBitOffset
`func (o *ResolvedEntity) UnsetBitOffset()`

UnsetBitOffset ensures that no value is present for BitOffset, not even an explicit nil
### GetByteOffset

`func (o *ResolvedEntity) GetByteOffset() int64`

GetByteOffset returns the ByteOffset field if non-nil, zero value otherwise.

### GetByteOffsetOk

`func (o *ResolvedEntity) GetByteOffsetOk() (*int64, bool)`

GetByteOffsetOk returns a tuple with the ByteOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteOffset

`func (o *ResolvedEntity) SetByteOffset(v int64)`

SetByteOffset sets ByteOffset field to given value.


### SetByteOffsetNil

`func (o *ResolvedEntity) SetByteOffsetNil(b bool)`

 SetByteOffsetNil sets the value for ByteOffset to be an explicit nil

### UnsetByteOffset
`func (o *ResolvedEntity) UnsetByteOffset()`

UnsetByteOffset ensures that no value is present for ByteOffset, not even an explicit nil
### GetByteSize

`func (o *ResolvedEntity) GetByteSize() int64`

GetByteSize returns the ByteSize field if non-nil, zero value otherwise.

### GetByteSizeOk

`func (o *ResolvedEntity) GetByteSizeOk() (*int64, bool)`

GetByteSizeOk returns a tuple with the ByteSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteSize

`func (o *ResolvedEntity) SetByteSize(v int64)`

SetByteSize sets ByteSize field to given value.


### SetByteSizeNil

`func (o *ResolvedEntity) SetByteSizeNil(b bool)`

 SetByteSizeNil sets the value for ByteSize to be an explicit nil

### UnsetByteSize
`func (o *ResolvedEntity) UnsetByteSize()`

UnsetByteSize ensures that no value is present for ByteSize, not even an explicit nil
### GetCount

`func (o *ResolvedEntity) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ResolvedEntity) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ResolvedEntity) SetCount(v int64)`

SetCount sets Count field to given value.


### GetDataTypeIndex

`func (o *ResolvedEntity) GetDataTypeIndex() int64`

GetDataTypeIndex returns the DataTypeIndex field if non-nil, zero value otherwise.

### GetDataTypeIndexOk

`func (o *ResolvedEntity) GetDataTypeIndexOk() (*int64, bool)`

GetDataTypeIndexOk returns a tuple with the DataTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypeIndex

`func (o *ResolvedEntity) SetDataTypeIndex(v int64)`

SetDataTypeIndex sets DataTypeIndex field to given value.


### SetDataTypeIndexNil

`func (o *ResolvedEntity) SetDataTypeIndexNil(b bool)`

 SetDataTypeIndexNil sets the value for DataTypeIndex to be an explicit nil

### UnsetDataTypeIndex
`func (o *ResolvedEntity) UnsetDataTypeIndex()`

UnsetDataTypeIndex ensures that no value is present for DataTypeIndex, not even an explicit nil
### GetFieldStatus

`func (o *ResolvedEntity) GetFieldStatus() string`

GetFieldStatus returns the FieldStatus field if non-nil, zero value otherwise.

### GetFieldStatusOk

`func (o *ResolvedEntity) GetFieldStatusOk() (*string, bool)`

GetFieldStatusOk returns a tuple with the FieldStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldStatus

`func (o *ResolvedEntity) SetFieldStatus(v string)`

SetFieldStatus sets FieldStatus field to given value.

### HasFieldStatus

`func (o *ResolvedEntity) HasFieldStatus() bool`

HasFieldStatus returns a boolean if a field has been set.

### GetFunctionId

`func (o *ResolvedEntity) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *ResolvedEntity) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *ResolvedEntity) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.

### HasFunctionId

`func (o *ResolvedEntity) HasFunctionId() bool`

HasFunctionId returns a boolean if a field has been set.

### GetImportedFunctionId

`func (o *ResolvedEntity) GetImportedFunctionId() int64`

GetImportedFunctionId returns the ImportedFunctionId field if non-nil, zero value otherwise.

### GetImportedFunctionIdOk

`func (o *ResolvedEntity) GetImportedFunctionIdOk() (*int64, bool)`

GetImportedFunctionIdOk returns a tuple with the ImportedFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedFunctionId

`func (o *ResolvedEntity) SetImportedFunctionId(v int64)`

SetImportedFunctionId sets ImportedFunctionId field to given value.

### HasImportedFunctionId

`func (o *ResolvedEntity) HasImportedFunctionId() bool`

HasImportedFunctionId returns a boolean if a field has been set.

### GetKind

`func (o *ResolvedEntity) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ResolvedEntity) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ResolvedEntity) SetKind(v string)`

SetKind sets Kind field to given value.


### SetKindNil

`func (o *ResolvedEntity) SetKindNil(b bool)`

 SetKindNil sets the value for Kind to be an explicit nil

### UnsetKind
`func (o *ResolvedEntity) UnsetKind()`

UnsetKind ensures that no value is present for Kind, not even an explicit nil
### GetName

`func (o *ResolvedEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResolvedEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResolvedEntity) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ResolvedEntity) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ResolvedEntity) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNameSource

`func (o *ResolvedEntity) GetNameSource() string`

GetNameSource returns the NameSource field if non-nil, zero value otherwise.

### GetNameSourceOk

`func (o *ResolvedEntity) GetNameSourceOk() (*string, bool)`

GetNameSourceOk returns a tuple with the NameSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSource

`func (o *ResolvedEntity) SetNameSource(v string)`

SetNameSource sets NameSource field to given value.


### GetNeedsNaming

`func (o *ResolvedEntity) GetNeedsNaming() bool`

GetNeedsNaming returns the NeedsNaming field if non-nil, zero value otherwise.

### GetNeedsNamingOk

`func (o *ResolvedEntity) GetNeedsNamingOk() (*bool, bool)`

GetNeedsNamingOk returns a tuple with the NeedsNaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsNaming

`func (o *ResolvedEntity) SetNeedsNaming(v bool)`

SetNeedsNaming sets NeedsNaming field to given value.


### GetProvenance

`func (o *ResolvedEntity) GetProvenance() string`

GetProvenance returns the Provenance field if non-nil, zero value otherwise.

### GetProvenanceOk

`func (o *ResolvedEntity) GetProvenanceOk() (*string, bool)`

GetProvenanceOk returns a tuple with the Provenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenance

`func (o *ResolvedEntity) SetProvenance(v string)`

SetProvenance sets Provenance field to given value.


### SetProvenanceNil

`func (o *ResolvedEntity) SetProvenanceNil(b bool)`

 SetProvenanceNil sets the value for Provenance to be an explicit nil

### UnsetProvenance
`func (o *ResolvedEntity) UnsetProvenance()`

UnsetProvenance ensures that no value is present for Provenance, not even an explicit nil
### GetResolvedName

`func (o *ResolvedEntity) GetResolvedName() string`

GetResolvedName returns the ResolvedName field if non-nil, zero value otherwise.

### GetResolvedNameOk

`func (o *ResolvedEntity) GetResolvedNameOk() (*string, bool)`

GetResolvedNameOk returns a tuple with the ResolvedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedName

`func (o *ResolvedEntity) SetResolvedName(v string)`

SetResolvedName sets ResolvedName field to given value.


### SetResolvedNameNil

`func (o *ResolvedEntity) SetResolvedNameNil(b bool)`

 SetResolvedNameNil sets the value for ResolvedName to be an explicit nil

### UnsetResolvedName
`func (o *ResolvedEntity) UnsetResolvedName()`

UnsetResolvedName ensures that no value is present for ResolvedName, not even an explicit nil
### GetSuggestedName

`func (o *ResolvedEntity) GetSuggestedName() string`

GetSuggestedName returns the SuggestedName field if non-nil, zero value otherwise.

### GetSuggestedNameOk

`func (o *ResolvedEntity) GetSuggestedNameOk() (*string, bool)`

GetSuggestedNameOk returns a tuple with the SuggestedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedName

`func (o *ResolvedEntity) SetSuggestedName(v string)`

SetSuggestedName sets SuggestedName field to given value.

### HasSuggestedName

`func (o *ResolvedEntity) HasSuggestedName() bool`

HasSuggestedName returns a boolean if a field has been set.

### SetSuggestedNameNil

`func (o *ResolvedEntity) SetSuggestedNameNil(b bool)`

 SetSuggestedNameNil sets the value for SuggestedName to be an explicit nil

### UnsetSuggestedName
`func (o *ResolvedEntity) UnsetSuggestedName()`

UnsetSuggestedName ensures that no value is present for SuggestedName, not even an explicit nil
### GetSuggestedType

`func (o *ResolvedEntity) GetSuggestedType() string`

GetSuggestedType returns the SuggestedType field if non-nil, zero value otherwise.

### GetSuggestedTypeOk

`func (o *ResolvedEntity) GetSuggestedTypeOk() (*string, bool)`

GetSuggestedTypeOk returns a tuple with the SuggestedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedType

`func (o *ResolvedEntity) SetSuggestedType(v string)`

SetSuggestedType sets SuggestedType field to given value.


### SetSuggestedTypeNil

`func (o *ResolvedEntity) SetSuggestedTypeNil(b bool)`

 SetSuggestedTypeNil sets the value for SuggestedType to be an explicit nil

### UnsetSuggestedType
`func (o *ResolvedEntity) UnsetSuggestedType()`

UnsetSuggestedType ensures that no value is present for SuggestedType, not even an explicit nil
### GetSuggestionConfidence

`func (o *ResolvedEntity) GetSuggestionConfidence() string`

GetSuggestionConfidence returns the SuggestionConfidence field if non-nil, zero value otherwise.

### GetSuggestionConfidenceOk

`func (o *ResolvedEntity) GetSuggestionConfidenceOk() (*string, bool)`

GetSuggestionConfidenceOk returns a tuple with the SuggestionConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestionConfidence

`func (o *ResolvedEntity) SetSuggestionConfidence(v string)`

SetSuggestionConfidence sets SuggestionConfidence field to given value.


### SetSuggestionConfidenceNil

`func (o *ResolvedEntity) SetSuggestionConfidenceNil(b bool)`

 SetSuggestionConfidenceNil sets the value for SuggestionConfidence to be an explicit nil

### UnsetSuggestionConfidence
`func (o *ResolvedEntity) UnsetSuggestionConfidence()`

UnsetSuggestionConfidence ensures that no value is present for SuggestionConfidence, not even an explicit nil
### GetToken

`func (o *ResolvedEntity) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ResolvedEntity) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ResolvedEntity) SetToken(v string)`

SetToken sets Token field to given value.


### SetTokenNil

`func (o *ResolvedEntity) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *ResolvedEntity) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetTypeIndex

`func (o *ResolvedEntity) GetTypeIndex() int64`

GetTypeIndex returns the TypeIndex field if non-nil, zero value otherwise.

### GetTypeIndexOk

`func (o *ResolvedEntity) GetTypeIndexOk() (*int64, bool)`

GetTypeIndexOk returns a tuple with the TypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeIndex

`func (o *ResolvedEntity) SetTypeIndex(v int64)`

SetTypeIndex sets TypeIndex field to given value.


### SetTypeIndexNil

`func (o *ResolvedEntity) SetTypeIndexNil(b bool)`

 SetTypeIndexNil sets the value for TypeIndex to be an explicit nil

### UnsetTypeIndex
`func (o *ResolvedEntity) UnsetTypeIndex()`

UnsetTypeIndex ensures that no value is present for TypeIndex, not even an explicit nil
### GetVaddr

`func (o *ResolvedEntity) GetVaddr() int64`

GetVaddr returns the Vaddr field if non-nil, zero value otherwise.

### GetVaddrOk

`func (o *ResolvedEntity) GetVaddrOk() (*int64, bool)`

GetVaddrOk returns a tuple with the Vaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaddr

`func (o *ResolvedEntity) SetVaddr(v int64)`

SetVaddr sets Vaddr field to given value.


### SetVaddrNil

`func (o *ResolvedEntity) SetVaddrNil(b bool)`

 SetVaddrNil sets the value for Vaddr to be an explicit nil

### UnsetVaddr
`func (o *ResolvedEntity) UnsetVaddr()`

UnsetVaddr ensures that no value is present for Vaddr, not even an explicit nil
### GetValue

`func (o *ResolvedEntity) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ResolvedEntity) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ResolvedEntity) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *ResolvedEntity) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ResolvedEntity) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetValueConfidence

`func (o *ResolvedEntity) GetValueConfidence() string`

GetValueConfidence returns the ValueConfidence field if non-nil, zero value otherwise.

### GetValueConfidenceOk

`func (o *ResolvedEntity) GetValueConfidenceOk() (*string, bool)`

GetValueConfidenceOk returns a tuple with the ValueConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueConfidence

`func (o *ResolvedEntity) SetValueConfidence(v string)`

SetValueConfidence sets ValueConfidence field to given value.


### SetValueConfidenceNil

`func (o *ResolvedEntity) SetValueConfidenceNil(b bool)`

 SetValueConfidenceNil sets the value for ValueConfidence to be an explicit nil

### UnsetValueConfidence
`func (o *ResolvedEntity) UnsetValueConfidence()`

UnsetValueConfidence ensures that no value is present for ValueConfidence, not even an explicit nil
### GetValueType

`func (o *ResolvedEntity) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *ResolvedEntity) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *ResolvedEntity) SetValueType(v string)`

SetValueType sets ValueType field to given value.


### SetValueTypeNil

`func (o *ResolvedEntity) SetValueTypeNil(b bool)`

 SetValueTypeNil sets the value for ValueType to be an explicit nil

### UnsetValueType
`func (o *ResolvedEntity) UnsetValueType()`

UnsetValueType ensures that no value is present for ValueType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


