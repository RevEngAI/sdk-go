# ArrayDataType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | When this type was extracted. | 
**DataTypeId** | **int64** | Identifies the type within its analysis. 0 is a valid id. | 
**Definition** | Pointer to [**ArrayDefinition**](ArrayDefinition.md) | Absent only for a type referenced but never defined. | [optional] 
**HasDefinition** | **bool** | Whether this type carries a definition. False for the kinds that never have one and for a type referenced but never defined. | 
**Kind** | **string** |  | 
**Name** | **string** | Type name. | 
**Namespace** | **string** | The scope qualifying the type name. Empty for a program-defined type. | 
**Size** | Pointer to **int64** | Size in bytes, absent when it could not be determined. | [optional] 
**SourceFunctionId** | Pointer to **int64** | The function this type was copied from, when transferred rather than extracted. | [optional] 
**SourceType** | **string** | Where this type came from. | 

## Methods

### NewArrayDataType

`func NewArrayDataType(createdAt time.Time, dataTypeId int64, hasDefinition bool, kind string, name string, namespace string, sourceType string, ) *ArrayDataType`

NewArrayDataType instantiates a new ArrayDataType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArrayDataTypeWithDefaults

`func NewArrayDataTypeWithDefaults() *ArrayDataType`

NewArrayDataTypeWithDefaults instantiates a new ArrayDataType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ArrayDataType) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArrayDataType) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArrayDataType) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDataTypeId

`func (o *ArrayDataType) GetDataTypeId() int64`

GetDataTypeId returns the DataTypeId field if non-nil, zero value otherwise.

### GetDataTypeIdOk

`func (o *ArrayDataType) GetDataTypeIdOk() (*int64, bool)`

GetDataTypeIdOk returns a tuple with the DataTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypeId

`func (o *ArrayDataType) SetDataTypeId(v int64)`

SetDataTypeId sets DataTypeId field to given value.


### GetDefinition

`func (o *ArrayDataType) GetDefinition() ArrayDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *ArrayDataType) GetDefinitionOk() (*ArrayDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *ArrayDataType) SetDefinition(v ArrayDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *ArrayDataType) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetHasDefinition

`func (o *ArrayDataType) GetHasDefinition() bool`

GetHasDefinition returns the HasDefinition field if non-nil, zero value otherwise.

### GetHasDefinitionOk

`func (o *ArrayDataType) GetHasDefinitionOk() (*bool, bool)`

GetHasDefinitionOk returns a tuple with the HasDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDefinition

`func (o *ArrayDataType) SetHasDefinition(v bool)`

SetHasDefinition sets HasDefinition field to given value.


### GetKind

`func (o *ArrayDataType) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ArrayDataType) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ArrayDataType) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *ArrayDataType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArrayDataType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArrayDataType) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *ArrayDataType) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ArrayDataType) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ArrayDataType) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetSize

`func (o *ArrayDataType) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ArrayDataType) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ArrayDataType) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ArrayDataType) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSourceFunctionId

`func (o *ArrayDataType) GetSourceFunctionId() int64`

GetSourceFunctionId returns the SourceFunctionId field if non-nil, zero value otherwise.

### GetSourceFunctionIdOk

`func (o *ArrayDataType) GetSourceFunctionIdOk() (*int64, bool)`

GetSourceFunctionIdOk returns a tuple with the SourceFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFunctionId

`func (o *ArrayDataType) SetSourceFunctionId(v int64)`

SetSourceFunctionId sets SourceFunctionId field to given value.

### HasSourceFunctionId

`func (o *ArrayDataType) HasSourceFunctionId() bool`

HasSourceFunctionId returns a boolean if a field has been set.

### GetSourceType

`func (o *ArrayDataType) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *ArrayDataType) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *ArrayDataType) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


