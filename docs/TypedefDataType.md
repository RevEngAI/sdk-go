# TypedefDataType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | When this type was extracted. | 
**DataTypeId** | **int64** | Identifies the type within its analysis. 0 is a valid id. | 
**Definition** | Pointer to [**TypedefDefinition**](TypedefDefinition.md) | Absent only for a type referenced but never defined. | [optional] 
**DefinitionPresent** | **bool** | Whether this type carries a definition. False for the kinds that never have one and for a type referenced but never defined. | 
**Kind** | **string** |  | 
**Name** | **string** | Type name. | 
**Namespace** | **string** | The scope qualifying the type name. Empty for a program-defined type. | 
**Size** | Pointer to **int64** | Size in bytes, absent when it could not be determined. | [optional] 
**SourceFunctionId** | Pointer to **int64** | The function this type was copied from, when transferred rather than extracted. | [optional] 
**SourceType** | **string** | Where this type came from. | 

## Methods

### NewTypedefDataType

`func NewTypedefDataType(createdAt time.Time, dataTypeId int64, definitionPresent bool, kind string, name string, namespace string, sourceType string, ) *TypedefDataType`

NewTypedefDataType instantiates a new TypedefDataType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypedefDataTypeWithDefaults

`func NewTypedefDataTypeWithDefaults() *TypedefDataType`

NewTypedefDataTypeWithDefaults instantiates a new TypedefDataType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *TypedefDataType) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TypedefDataType) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TypedefDataType) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDataTypeId

`func (o *TypedefDataType) GetDataTypeId() int64`

GetDataTypeId returns the DataTypeId field if non-nil, zero value otherwise.

### GetDataTypeIdOk

`func (o *TypedefDataType) GetDataTypeIdOk() (*int64, bool)`

GetDataTypeIdOk returns a tuple with the DataTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypeId

`func (o *TypedefDataType) SetDataTypeId(v int64)`

SetDataTypeId sets DataTypeId field to given value.


### GetDefinition

`func (o *TypedefDataType) GetDefinition() TypedefDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *TypedefDataType) GetDefinitionOk() (*TypedefDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *TypedefDataType) SetDefinition(v TypedefDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *TypedefDataType) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetDefinitionPresent

`func (o *TypedefDataType) GetDefinitionPresent() bool`

GetDefinitionPresent returns the DefinitionPresent field if non-nil, zero value otherwise.

### GetDefinitionPresentOk

`func (o *TypedefDataType) GetDefinitionPresentOk() (*bool, bool)`

GetDefinitionPresentOk returns a tuple with the DefinitionPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionPresent

`func (o *TypedefDataType) SetDefinitionPresent(v bool)`

SetDefinitionPresent sets DefinitionPresent field to given value.


### GetKind

`func (o *TypedefDataType) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TypedefDataType) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TypedefDataType) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *TypedefDataType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TypedefDataType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TypedefDataType) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *TypedefDataType) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *TypedefDataType) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *TypedefDataType) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetSize

`func (o *TypedefDataType) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TypedefDataType) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TypedefDataType) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *TypedefDataType) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSourceFunctionId

`func (o *TypedefDataType) GetSourceFunctionId() int64`

GetSourceFunctionId returns the SourceFunctionId field if non-nil, zero value otherwise.

### GetSourceFunctionIdOk

`func (o *TypedefDataType) GetSourceFunctionIdOk() (*int64, bool)`

GetSourceFunctionIdOk returns a tuple with the SourceFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFunctionId

`func (o *TypedefDataType) SetSourceFunctionId(v int64)`

SetSourceFunctionId sets SourceFunctionId field to given value.

### HasSourceFunctionId

`func (o *TypedefDataType) HasSourceFunctionId() bool`

HasSourceFunctionId returns a boolean if a field has been set.

### GetSourceType

`func (o *TypedefDataType) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *TypedefDataType) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *TypedefDataType) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


