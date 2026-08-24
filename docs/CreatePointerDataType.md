# CreatePointerDataType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Definition** | [**PointerDefinition**](PointerDefinition.md) |  | 
**Kind** | **string** |  | 
**Name** | **string** | Type name. Unique within the analysis for a given namespace and kind. | 
**Namespace** | Pointer to **string** | The scope qualifying the type name. Omit for a type of the binary&#39;s own. | [optional] 
**Size** | Pointer to **int64** | Size in bytes. Omit when it is not known. | [optional] 

## Methods

### NewCreatePointerDataType

`func NewCreatePointerDataType(definition PointerDefinition, kind string, name string, ) *CreatePointerDataType`

NewCreatePointerDataType instantiates a new CreatePointerDataType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePointerDataTypeWithDefaults

`func NewCreatePointerDataTypeWithDefaults() *CreatePointerDataType`

NewCreatePointerDataTypeWithDefaults instantiates a new CreatePointerDataType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefinition

`func (o *CreatePointerDataType) GetDefinition() PointerDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *CreatePointerDataType) GetDefinitionOk() (*PointerDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *CreatePointerDataType) SetDefinition(v PointerDefinition)`

SetDefinition sets Definition field to given value.


### GetKind

`func (o *CreatePointerDataType) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreatePointerDataType) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreatePointerDataType) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *CreatePointerDataType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePointerDataType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePointerDataType) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *CreatePointerDataType) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CreatePointerDataType) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CreatePointerDataType) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CreatePointerDataType) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetSize

`func (o *CreatePointerDataType) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreatePointerDataType) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreatePointerDataType) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreatePointerDataType) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


