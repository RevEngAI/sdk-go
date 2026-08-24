# CreateUnknownDataType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Name** | **string** | Type name. Unique within the analysis for a given namespace and kind. | 
**Namespace** | Pointer to **string** | The scope qualifying the type name. Omit for a type of the binary&#39;s own. | [optional] 
**Size** | Pointer to **int64** | Size in bytes. Omit when it is not known. | [optional] 

## Methods

### NewCreateUnknownDataType

`func NewCreateUnknownDataType(kind string, name string, ) *CreateUnknownDataType`

NewCreateUnknownDataType instantiates a new CreateUnknownDataType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUnknownDataTypeWithDefaults

`func NewCreateUnknownDataTypeWithDefaults() *CreateUnknownDataType`

NewCreateUnknownDataTypeWithDefaults instantiates a new CreateUnknownDataType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CreateUnknownDataType) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateUnknownDataType) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateUnknownDataType) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *CreateUnknownDataType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUnknownDataType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUnknownDataType) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *CreateUnknownDataType) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CreateUnknownDataType) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CreateUnknownDataType) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CreateUnknownDataType) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetSize

`func (o *CreateUnknownDataType) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateUnknownDataType) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateUnknownDataType) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateUnknownDataType) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


