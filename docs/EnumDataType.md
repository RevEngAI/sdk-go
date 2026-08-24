# EnumDataType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | When this type was extracted. | 
**DataTypeId** | **int64** | Identifies the type within its analysis. 0 is a valid id. | 
**Definition** | Pointer to [**EnumDefinition**](EnumDefinition.md) | Absent only for a type referenced but never defined. | [optional] 
**HasDefinition** | **bool** | Whether this type carries a definition. False for the kinds that never have one and for a type referenced but never defined. | 
**Kind** | **string** |  | 
**Name** | **string** | Type name. | 
**Namespace** | **string** | The scope qualifying the type name. Empty for a program-defined type. | 
**Size** | Pointer to **int64** | Size in bytes, absent when it could not be determined. | [optional] 
**SourceFunctionId** | Pointer to **int64** | The function this type was copied from, when transferred rather than extracted. | [optional] 
**SourceType** | **string** | Where this type came from. | 

## Methods

### NewEnumDataType

`func NewEnumDataType(createdAt time.Time, dataTypeId int64, hasDefinition bool, kind string, name string, namespace string, sourceType string, ) *EnumDataType`

NewEnumDataType instantiates a new EnumDataType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnumDataTypeWithDefaults

`func NewEnumDataTypeWithDefaults() *EnumDataType`

NewEnumDataTypeWithDefaults instantiates a new EnumDataType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *EnumDataType) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnumDataType) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnumDataType) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDataTypeId

`func (o *EnumDataType) GetDataTypeId() int64`

GetDataTypeId returns the DataTypeId field if non-nil, zero value otherwise.

### GetDataTypeIdOk

`func (o *EnumDataType) GetDataTypeIdOk() (*int64, bool)`

GetDataTypeIdOk returns a tuple with the DataTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypeId

`func (o *EnumDataType) SetDataTypeId(v int64)`

SetDataTypeId sets DataTypeId field to given value.


### GetDefinition

`func (o *EnumDataType) GetDefinition() EnumDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *EnumDataType) GetDefinitionOk() (*EnumDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *EnumDataType) SetDefinition(v EnumDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *EnumDataType) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetHasDefinition

`func (o *EnumDataType) GetHasDefinition() bool`

GetHasDefinition returns the HasDefinition field if non-nil, zero value otherwise.

### GetHasDefinitionOk

`func (o *EnumDataType) GetHasDefinitionOk() (*bool, bool)`

GetHasDefinitionOk returns a tuple with the HasDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDefinition

`func (o *EnumDataType) SetHasDefinition(v bool)`

SetHasDefinition sets HasDefinition field to given value.


### GetKind

`func (o *EnumDataType) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EnumDataType) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EnumDataType) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *EnumDataType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnumDataType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnumDataType) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *EnumDataType) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *EnumDataType) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *EnumDataType) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetSize

`func (o *EnumDataType) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *EnumDataType) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *EnumDataType) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *EnumDataType) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSourceFunctionId

`func (o *EnumDataType) GetSourceFunctionId() int64`

GetSourceFunctionId returns the SourceFunctionId field if non-nil, zero value otherwise.

### GetSourceFunctionIdOk

`func (o *EnumDataType) GetSourceFunctionIdOk() (*int64, bool)`

GetSourceFunctionIdOk returns a tuple with the SourceFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFunctionId

`func (o *EnumDataType) SetSourceFunctionId(v int64)`

SetSourceFunctionId sets SourceFunctionId field to given value.

### HasSourceFunctionId

`func (o *EnumDataType) HasSourceFunctionId() bool`

HasSourceFunctionId returns a boolean if a field has been set.

### GetSourceType

`func (o *EnumDataType) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *EnumDataType) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *EnumDataType) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


