# DataTypeEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | When this type was extracted. | 
**DataTypeId** | **int64** | Identifies the type within its analysis. 0 is a valid id. | 
**Definition** | Pointer to [**FunctionTypeDefinition**](FunctionTypeDefinition.md) | Absent only for a type referenced but never defined. | [optional] 
**DefinitionPresent** | **bool** | Whether this type carries a definition. False for the kinds that never have one and for a type referenced but never defined. | 
**Kind** | **string** |  | 
**Name** | **string** | Type name. | 
**Namespace** | **string** | The scope qualifying the type name. Empty for a program-defined type. | 
**Size** | Pointer to **int64** | Size in bytes, absent when it could not be determined. | [optional] 
**SourceAnalysisId** | Pointer to **int64** | ID of the analysis the source function belongs to, when it could be resolved. | [optional] 
**SourceFunctionId** | Pointer to **int64** | The function this type was copied from, when transferred rather than extracted. | [optional] 
**SourceType** | **string** | Where this type came from. | 

## Methods

### NewDataTypeEntry

`func NewDataTypeEntry(createdAt time.Time, dataTypeId int64, definitionPresent bool, kind string, name string, namespace string, sourceType string, ) *DataTypeEntry`

NewDataTypeEntry instantiates a new DataTypeEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTypeEntryWithDefaults

`func NewDataTypeEntryWithDefaults() *DataTypeEntry`

NewDataTypeEntryWithDefaults instantiates a new DataTypeEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DataTypeEntry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DataTypeEntry) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DataTypeEntry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDataTypeId

`func (o *DataTypeEntry) GetDataTypeId() int64`

GetDataTypeId returns the DataTypeId field if non-nil, zero value otherwise.

### GetDataTypeIdOk

`func (o *DataTypeEntry) GetDataTypeIdOk() (*int64, bool)`

GetDataTypeIdOk returns a tuple with the DataTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypeId

`func (o *DataTypeEntry) SetDataTypeId(v int64)`

SetDataTypeId sets DataTypeId field to given value.


### GetDefinition

`func (o *DataTypeEntry) GetDefinition() FunctionTypeDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *DataTypeEntry) GetDefinitionOk() (*FunctionTypeDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *DataTypeEntry) SetDefinition(v FunctionTypeDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *DataTypeEntry) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetDefinitionPresent

`func (o *DataTypeEntry) GetDefinitionPresent() bool`

GetDefinitionPresent returns the DefinitionPresent field if non-nil, zero value otherwise.

### GetDefinitionPresentOk

`func (o *DataTypeEntry) GetDefinitionPresentOk() (*bool, bool)`

GetDefinitionPresentOk returns a tuple with the DefinitionPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionPresent

`func (o *DataTypeEntry) SetDefinitionPresent(v bool)`

SetDefinitionPresent sets DefinitionPresent field to given value.


### GetKind

`func (o *DataTypeEntry) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DataTypeEntry) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DataTypeEntry) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *DataTypeEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataTypeEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataTypeEntry) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *DataTypeEntry) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *DataTypeEntry) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *DataTypeEntry) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetSize

`func (o *DataTypeEntry) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DataTypeEntry) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DataTypeEntry) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *DataTypeEntry) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSourceAnalysisId

`func (o *DataTypeEntry) GetSourceAnalysisId() int64`

GetSourceAnalysisId returns the SourceAnalysisId field if non-nil, zero value otherwise.

### GetSourceAnalysisIdOk

`func (o *DataTypeEntry) GetSourceAnalysisIdOk() (*int64, bool)`

GetSourceAnalysisIdOk returns a tuple with the SourceAnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAnalysisId

`func (o *DataTypeEntry) SetSourceAnalysisId(v int64)`

SetSourceAnalysisId sets SourceAnalysisId field to given value.

### HasSourceAnalysisId

`func (o *DataTypeEntry) HasSourceAnalysisId() bool`

HasSourceAnalysisId returns a boolean if a field has been set.

### GetSourceFunctionId

`func (o *DataTypeEntry) GetSourceFunctionId() int64`

GetSourceFunctionId returns the SourceFunctionId field if non-nil, zero value otherwise.

### GetSourceFunctionIdOk

`func (o *DataTypeEntry) GetSourceFunctionIdOk() (*int64, bool)`

GetSourceFunctionIdOk returns a tuple with the SourceFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFunctionId

`func (o *DataTypeEntry) SetSourceFunctionId(v int64)`

SetSourceFunctionId sets SourceFunctionId field to given value.

### HasSourceFunctionId

`func (o *DataTypeEntry) HasSourceFunctionId() bool`

HasSourceFunctionId returns a boolean if a field has been set.

### GetSourceType

`func (o *DataTypeEntry) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *DataTypeEntry) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *DataTypeEntry) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


