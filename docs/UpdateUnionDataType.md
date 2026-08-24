# UpdateUnionDataType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTypeId** | **int64** | The type to replace, as returned by the data types list for this analysis. | 
**Definition** | [**UnionDefinition**](UnionDefinition.md) |  | 
**Kind** | **string** |  | 
**Name** | **string** | Type name. Unique within the analysis for a given namespace and kind. | 
**Namespace** | Pointer to **string** | The scope qualifying the type name. Omit for a type of the binary&#39;s own. | [optional] 
**Size** | Pointer to **int64** | Size in bytes. Omit when it is not known. | [optional] 

## Methods

### NewUpdateUnionDataType

`func NewUpdateUnionDataType(dataTypeId int64, definition UnionDefinition, kind string, name string, ) *UpdateUnionDataType`

NewUpdateUnionDataType instantiates a new UpdateUnionDataType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUnionDataTypeWithDefaults

`func NewUpdateUnionDataTypeWithDefaults() *UpdateUnionDataType`

NewUpdateUnionDataTypeWithDefaults instantiates a new UpdateUnionDataType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTypeId

`func (o *UpdateUnionDataType) GetDataTypeId() int64`

GetDataTypeId returns the DataTypeId field if non-nil, zero value otherwise.

### GetDataTypeIdOk

`func (o *UpdateUnionDataType) GetDataTypeIdOk() (*int64, bool)`

GetDataTypeIdOk returns a tuple with the DataTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypeId

`func (o *UpdateUnionDataType) SetDataTypeId(v int64)`

SetDataTypeId sets DataTypeId field to given value.


### GetDefinition

`func (o *UpdateUnionDataType) GetDefinition() UnionDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *UpdateUnionDataType) GetDefinitionOk() (*UnionDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *UpdateUnionDataType) SetDefinition(v UnionDefinition)`

SetDefinition sets Definition field to given value.


### GetKind

`func (o *UpdateUnionDataType) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *UpdateUnionDataType) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *UpdateUnionDataType) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *UpdateUnionDataType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateUnionDataType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateUnionDataType) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *UpdateUnionDataType) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *UpdateUnionDataType) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *UpdateUnionDataType) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *UpdateUnionDataType) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetSize

`func (o *UpdateUnionDataType) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *UpdateUnionDataType) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *UpdateUnionDataType) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *UpdateUnionDataType) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


