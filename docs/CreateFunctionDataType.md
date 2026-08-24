# CreateFunctionDataType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Definition** | [**FunctionTypeDefinition**](FunctionTypeDefinition.md) |  | 
**Kind** | **string** |  | 
**Name** | **string** | Type name. Unique within the analysis for a given namespace and kind. | 
**Namespace** | Pointer to **string** | The scope qualifying the type name. Omit for a type of the binary&#39;s own. | [optional] 
**Size** | Pointer to **int64** | Size in bytes. Omit when it is not known. | [optional] 

## Methods

### NewCreateFunctionDataType

`func NewCreateFunctionDataType(definition FunctionTypeDefinition, kind string, name string, ) *CreateFunctionDataType`

NewCreateFunctionDataType instantiates a new CreateFunctionDataType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFunctionDataTypeWithDefaults

`func NewCreateFunctionDataTypeWithDefaults() *CreateFunctionDataType`

NewCreateFunctionDataTypeWithDefaults instantiates a new CreateFunctionDataType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefinition

`func (o *CreateFunctionDataType) GetDefinition() FunctionTypeDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *CreateFunctionDataType) GetDefinitionOk() (*FunctionTypeDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *CreateFunctionDataType) SetDefinition(v FunctionTypeDefinition)`

SetDefinition sets Definition field to given value.


### GetKind

`func (o *CreateFunctionDataType) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateFunctionDataType) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateFunctionDataType) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *CreateFunctionDataType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFunctionDataType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFunctionDataType) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *CreateFunctionDataType) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CreateFunctionDataType) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CreateFunctionDataType) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CreateFunctionDataType) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetSize

`func (o *CreateFunctionDataType) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateFunctionDataType) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateFunctionDataType) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateFunctionDataType) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


