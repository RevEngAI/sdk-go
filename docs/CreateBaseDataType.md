# CreateBaseDataType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Name** | **string** | Type name. Unique within the analysis for a given namespace and kind. | 
**Namespace** | Pointer to **string** | The scope qualifying the type name. Omit for a type of the binary&#39;s own. | [optional] 
**Size** | Pointer to **int64** | Size in bytes. Omit when it is not known. | [optional] 

## Methods

### NewCreateBaseDataType

`func NewCreateBaseDataType(kind string, name string, ) *CreateBaseDataType`

NewCreateBaseDataType instantiates a new CreateBaseDataType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBaseDataTypeWithDefaults

`func NewCreateBaseDataTypeWithDefaults() *CreateBaseDataType`

NewCreateBaseDataTypeWithDefaults instantiates a new CreateBaseDataType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CreateBaseDataType) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateBaseDataType) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateBaseDataType) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *CreateBaseDataType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateBaseDataType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateBaseDataType) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *CreateBaseDataType) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CreateBaseDataType) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CreateBaseDataType) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CreateBaseDataType) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetSize

`func (o *CreateBaseDataType) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateBaseDataType) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateBaseDataType) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateBaseDataType) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


