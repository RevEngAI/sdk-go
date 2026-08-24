# PointerDataType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | When this type was extracted. | 
**DataTypeId** | **int64** | Identifies the type within its analysis. 0 is a valid id. | 
**Definition** | Pointer to [**PointerDefinition**](PointerDefinition.md) | Absent only for a type referenced but never defined. | [optional] 
**HasDefinition** | **bool** | Whether this type carries a definition. False for the kinds that never have one and for a type referenced but never defined. | 
**Kind** | **string** |  | 
**Name** | **string** | Type name. | 
**Namespace** | **string** | The scope qualifying the type name. Empty for a program-defined type. | 
**Size** | Pointer to **int64** | Size in bytes, absent when it could not be determined. | [optional] 
**SourceFunctionId** | Pointer to **int64** | The function this type was copied from, when transferred rather than extracted. | [optional] 
**SourceType** | **string** | Where this type came from. | 

## Methods

### NewPointerDataType

`func NewPointerDataType(createdAt time.Time, dataTypeId int64, hasDefinition bool, kind string, name string, namespace string, sourceType string, ) *PointerDataType`

NewPointerDataType instantiates a new PointerDataType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointerDataTypeWithDefaults

`func NewPointerDataTypeWithDefaults() *PointerDataType`

NewPointerDataTypeWithDefaults instantiates a new PointerDataType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PointerDataType) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PointerDataType) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PointerDataType) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDataTypeId

`func (o *PointerDataType) GetDataTypeId() int64`

GetDataTypeId returns the DataTypeId field if non-nil, zero value otherwise.

### GetDataTypeIdOk

`func (o *PointerDataType) GetDataTypeIdOk() (*int64, bool)`

GetDataTypeIdOk returns a tuple with the DataTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypeId

`func (o *PointerDataType) SetDataTypeId(v int64)`

SetDataTypeId sets DataTypeId field to given value.


### GetDefinition

`func (o *PointerDataType) GetDefinition() PointerDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *PointerDataType) GetDefinitionOk() (*PointerDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *PointerDataType) SetDefinition(v PointerDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *PointerDataType) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetHasDefinition

`func (o *PointerDataType) GetHasDefinition() bool`

GetHasDefinition returns the HasDefinition field if non-nil, zero value otherwise.

### GetHasDefinitionOk

`func (o *PointerDataType) GetHasDefinitionOk() (*bool, bool)`

GetHasDefinitionOk returns a tuple with the HasDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDefinition

`func (o *PointerDataType) SetHasDefinition(v bool)`

SetHasDefinition sets HasDefinition field to given value.


### GetKind

`func (o *PointerDataType) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PointerDataType) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PointerDataType) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *PointerDataType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PointerDataType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PointerDataType) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *PointerDataType) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PointerDataType) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PointerDataType) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetSize

`func (o *PointerDataType) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PointerDataType) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PointerDataType) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *PointerDataType) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSourceFunctionId

`func (o *PointerDataType) GetSourceFunctionId() int64`

GetSourceFunctionId returns the SourceFunctionId field if non-nil, zero value otherwise.

### GetSourceFunctionIdOk

`func (o *PointerDataType) GetSourceFunctionIdOk() (*int64, bool)`

GetSourceFunctionIdOk returns a tuple with the SourceFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFunctionId

`func (o *PointerDataType) SetSourceFunctionId(v int64)`

SetSourceFunctionId sets SourceFunctionId field to given value.

### HasSourceFunctionId

`func (o *PointerDataType) HasSourceFunctionId() bool`

HasSourceFunctionId returns a boolean if a field has been set.

### GetSourceType

`func (o *PointerDataType) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *PointerDataType) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *PointerDataType) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


