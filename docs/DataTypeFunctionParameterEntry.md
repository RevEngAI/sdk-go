# DataTypeFunctionParameterEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTypeId** | Pointer to **int64** | The parameter&#39;s type. | [optional] 
**Name** | Pointer to **string** | Parameter name, when the producer had one. | [optional] 
**Ordinal** | **int64** | Zero-based argument position. | 
**Size** | **int64** | Parameter size in bytes. | 

## Methods

### NewDataTypeFunctionParameterEntry

`func NewDataTypeFunctionParameterEntry(ordinal int64, size int64, ) *DataTypeFunctionParameterEntry`

NewDataTypeFunctionParameterEntry instantiates a new DataTypeFunctionParameterEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTypeFunctionParameterEntryWithDefaults

`func NewDataTypeFunctionParameterEntryWithDefaults() *DataTypeFunctionParameterEntry`

NewDataTypeFunctionParameterEntryWithDefaults instantiates a new DataTypeFunctionParameterEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTypeId

`func (o *DataTypeFunctionParameterEntry) GetDataTypeId() int64`

GetDataTypeId returns the DataTypeId field if non-nil, zero value otherwise.

### GetDataTypeIdOk

`func (o *DataTypeFunctionParameterEntry) GetDataTypeIdOk() (*int64, bool)`

GetDataTypeIdOk returns a tuple with the DataTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypeId

`func (o *DataTypeFunctionParameterEntry) SetDataTypeId(v int64)`

SetDataTypeId sets DataTypeId field to given value.

### HasDataTypeId

`func (o *DataTypeFunctionParameterEntry) HasDataTypeId() bool`

HasDataTypeId returns a boolean if a field has been set.

### GetName

`func (o *DataTypeFunctionParameterEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataTypeFunctionParameterEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataTypeFunctionParameterEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataTypeFunctionParameterEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrdinal

`func (o *DataTypeFunctionParameterEntry) GetOrdinal() int64`

GetOrdinal returns the Ordinal field if non-nil, zero value otherwise.

### GetOrdinalOk

`func (o *DataTypeFunctionParameterEntry) GetOrdinalOk() (*int64, bool)`

GetOrdinalOk returns a tuple with the Ordinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdinal

`func (o *DataTypeFunctionParameterEntry) SetOrdinal(v int64)`

SetOrdinal sets Ordinal field to given value.


### GetSize

`func (o *DataTypeFunctionParameterEntry) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DataTypeFunctionParameterEntry) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DataTypeFunctionParameterEntry) SetSize(v int64)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


