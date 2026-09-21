# BaseDataType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | When this type was extracted. | 
**DataTypeId** | **int64** | Identifies the type within its analysis. 0 is a valid id. | 
**DefinitionPresent** | **bool** | Whether this type carries a definition. False for the kinds that never have one and for a type referenced but never defined. | 
**Kind** | **string** |  | 
**Name** | **string** | Type name. | 
**Namespace** | **string** | The scope qualifying the type name. Empty for a program-defined type. | 
**Size** | Pointer to **int64** | Size in bytes, absent when it could not be determined. | [optional] 
**SourceAnalysisId** | Pointer to **int64** | ID of the analysis the source function belongs to, when it could be resolved. | [optional] 
**SourceFunctionId** | Pointer to **int64** | The function this type was copied from, when transferred rather than extracted. | [optional] 
**SourceType** | **string** | Where this type came from. | 

## Methods

### NewBaseDataType

`func NewBaseDataType(createdAt time.Time, dataTypeId int64, definitionPresent bool, kind string, name string, namespace string, sourceType string, ) *BaseDataType`

NewBaseDataType instantiates a new BaseDataType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseDataTypeWithDefaults

`func NewBaseDataTypeWithDefaults() *BaseDataType`

NewBaseDataTypeWithDefaults instantiates a new BaseDataType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *BaseDataType) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseDataType) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseDataType) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDataTypeId

`func (o *BaseDataType) GetDataTypeId() int64`

GetDataTypeId returns the DataTypeId field if non-nil, zero value otherwise.

### GetDataTypeIdOk

`func (o *BaseDataType) GetDataTypeIdOk() (*int64, bool)`

GetDataTypeIdOk returns a tuple with the DataTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypeId

`func (o *BaseDataType) SetDataTypeId(v int64)`

SetDataTypeId sets DataTypeId field to given value.


### GetDefinitionPresent

`func (o *BaseDataType) GetDefinitionPresent() bool`

GetDefinitionPresent returns the DefinitionPresent field if non-nil, zero value otherwise.

### GetDefinitionPresentOk

`func (o *BaseDataType) GetDefinitionPresentOk() (*bool, bool)`

GetDefinitionPresentOk returns a tuple with the DefinitionPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionPresent

`func (o *BaseDataType) SetDefinitionPresent(v bool)`

SetDefinitionPresent sets DefinitionPresent field to given value.


### GetKind

`func (o *BaseDataType) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BaseDataType) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BaseDataType) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *BaseDataType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseDataType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseDataType) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *BaseDataType) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *BaseDataType) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *BaseDataType) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetSize

`func (o *BaseDataType) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BaseDataType) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BaseDataType) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *BaseDataType) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSourceAnalysisId

`func (o *BaseDataType) GetSourceAnalysisId() int64`

GetSourceAnalysisId returns the SourceAnalysisId field if non-nil, zero value otherwise.

### GetSourceAnalysisIdOk

`func (o *BaseDataType) GetSourceAnalysisIdOk() (*int64, bool)`

GetSourceAnalysisIdOk returns a tuple with the SourceAnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAnalysisId

`func (o *BaseDataType) SetSourceAnalysisId(v int64)`

SetSourceAnalysisId sets SourceAnalysisId field to given value.

### HasSourceAnalysisId

`func (o *BaseDataType) HasSourceAnalysisId() bool`

HasSourceAnalysisId returns a boolean if a field has been set.

### GetSourceFunctionId

`func (o *BaseDataType) GetSourceFunctionId() int64`

GetSourceFunctionId returns the SourceFunctionId field if non-nil, zero value otherwise.

### GetSourceFunctionIdOk

`func (o *BaseDataType) GetSourceFunctionIdOk() (*int64, bool)`

GetSourceFunctionIdOk returns a tuple with the SourceFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFunctionId

`func (o *BaseDataType) SetSourceFunctionId(v int64)`

SetSourceFunctionId sets SourceFunctionId field to given value.

### HasSourceFunctionId

`func (o *BaseDataType) HasSourceFunctionId() bool`

HasSourceFunctionId returns a boolean if a field has been set.

### GetSourceType

`func (o *BaseDataType) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *BaseDataType) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *BaseDataType) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


