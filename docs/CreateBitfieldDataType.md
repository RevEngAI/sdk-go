# CreateBitfieldDataType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Name** | **string** | Type name. Unique within the analysis for a given namespace and kind. | 
**Namespace** | Pointer to **string** | The scope qualifying the type name. Omit for a type of the binary&#39;s own. | [optional] 
**Size** | Pointer to **int64** | Size in bytes. Omit when it is not known. | [optional] 

## Methods

### NewCreateBitfieldDataType

`func NewCreateBitfieldDataType(kind string, name string, ) *CreateBitfieldDataType`

NewCreateBitfieldDataType instantiates a new CreateBitfieldDataType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBitfieldDataTypeWithDefaults

`func NewCreateBitfieldDataTypeWithDefaults() *CreateBitfieldDataType`

NewCreateBitfieldDataTypeWithDefaults instantiates a new CreateBitfieldDataType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CreateBitfieldDataType) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateBitfieldDataType) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateBitfieldDataType) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *CreateBitfieldDataType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateBitfieldDataType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateBitfieldDataType) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *CreateBitfieldDataType) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CreateBitfieldDataType) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CreateBitfieldDataType) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CreateBitfieldDataType) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetSize

`func (o *CreateBitfieldDataType) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateBitfieldDataType) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateBitfieldDataType) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateBitfieldDataType) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


