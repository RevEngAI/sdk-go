# UpdateArrayDataType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTypeId** | **int64** | The type to replace, as returned by the data types list for this analysis. | 
**Definition** | [**ArrayDefinition**](ArrayDefinition.md) |  | 
**Kind** | **string** |  | 
**Name** | **string** | Type name. Unique within the analysis for a given namespace and kind. | 
**Namespace** | Pointer to **string** | The scope qualifying the type name. Omit for a type of the binary&#39;s own. | [optional] 
**Size** | Pointer to **int64** | Size in bytes. Omit when it is not known. | [optional] 

## Methods

### NewUpdateArrayDataType

`func NewUpdateArrayDataType(dataTypeId int64, definition ArrayDefinition, kind string, name string, ) *UpdateArrayDataType`

NewUpdateArrayDataType instantiates a new UpdateArrayDataType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateArrayDataTypeWithDefaults

`func NewUpdateArrayDataTypeWithDefaults() *UpdateArrayDataType`

NewUpdateArrayDataTypeWithDefaults instantiates a new UpdateArrayDataType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTypeId

`func (o *UpdateArrayDataType) GetDataTypeId() int64`

GetDataTypeId returns the DataTypeId field if non-nil, zero value otherwise.

### GetDataTypeIdOk

`func (o *UpdateArrayDataType) GetDataTypeIdOk() (*int64, bool)`

GetDataTypeIdOk returns a tuple with the DataTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypeId

`func (o *UpdateArrayDataType) SetDataTypeId(v int64)`

SetDataTypeId sets DataTypeId field to given value.


### GetDefinition

`func (o *UpdateArrayDataType) GetDefinition() ArrayDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *UpdateArrayDataType) GetDefinitionOk() (*ArrayDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *UpdateArrayDataType) SetDefinition(v ArrayDefinition)`

SetDefinition sets Definition field to given value.


### GetKind

`func (o *UpdateArrayDataType) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *UpdateArrayDataType) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *UpdateArrayDataType) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *UpdateArrayDataType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateArrayDataType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateArrayDataType) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *UpdateArrayDataType) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *UpdateArrayDataType) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *UpdateArrayDataType) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *UpdateArrayDataType) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetSize

`func (o *UpdateArrayDataType) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *UpdateArrayDataType) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *UpdateArrayDataType) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *UpdateArrayDataType) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


