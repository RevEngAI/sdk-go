# UpdateUnknownDataType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTypeId** | **int64** | The type to replace, as returned by the data types list for this analysis. | 
**Kind** | **string** |  | 
**Name** | **string** | Type name. Unique within the analysis for a given namespace and kind. | 
**Namespace** | Pointer to **string** | The scope qualifying the type name. Omit for a type of the binary&#39;s own. | [optional] 
**Size** | Pointer to **int64** | Size in bytes. Omit when it is not known. | [optional] 

## Methods

### NewUpdateUnknownDataType

`func NewUpdateUnknownDataType(dataTypeId int64, kind string, name string, ) *UpdateUnknownDataType`

NewUpdateUnknownDataType instantiates a new UpdateUnknownDataType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUnknownDataTypeWithDefaults

`func NewUpdateUnknownDataTypeWithDefaults() *UpdateUnknownDataType`

NewUpdateUnknownDataTypeWithDefaults instantiates a new UpdateUnknownDataType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTypeId

`func (o *UpdateUnknownDataType) GetDataTypeId() int64`

GetDataTypeId returns the DataTypeId field if non-nil, zero value otherwise.

### GetDataTypeIdOk

`func (o *UpdateUnknownDataType) GetDataTypeIdOk() (*int64, bool)`

GetDataTypeIdOk returns a tuple with the DataTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypeId

`func (o *UpdateUnknownDataType) SetDataTypeId(v int64)`

SetDataTypeId sets DataTypeId field to given value.


### GetKind

`func (o *UpdateUnknownDataType) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *UpdateUnknownDataType) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *UpdateUnknownDataType) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *UpdateUnknownDataType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateUnknownDataType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateUnknownDataType) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *UpdateUnknownDataType) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *UpdateUnknownDataType) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *UpdateUnknownDataType) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *UpdateUnknownDataType) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetSize

`func (o *UpdateUnknownDataType) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *UpdateUnknownDataType) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *UpdateUnknownDataType) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *UpdateUnknownDataType) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


