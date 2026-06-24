# DataTypesEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTypes** | Pointer to [**FunctionInfo**](FunctionInfo.md) |  | [optional] 
**FunctionId** | **int64** |  | 

## Methods

### NewDataTypesEntry

`func NewDataTypesEntry(functionId int64, ) *DataTypesEntry`

NewDataTypesEntry instantiates a new DataTypesEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTypesEntryWithDefaults

`func NewDataTypesEntryWithDefaults() *DataTypesEntry`

NewDataTypesEntryWithDefaults instantiates a new DataTypesEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTypes

`func (o *DataTypesEntry) GetDataTypes() FunctionInfo`

GetDataTypes returns the DataTypes field if non-nil, zero value otherwise.

### GetDataTypesOk

`func (o *DataTypesEntry) GetDataTypesOk() (*FunctionInfo, bool)`

GetDataTypesOk returns a tuple with the DataTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypes

`func (o *DataTypesEntry) SetDataTypes(v FunctionInfo)`

SetDataTypes sets DataTypes field to given value.

### HasDataTypes

`func (o *DataTypesEntry) HasDataTypes() bool`

HasDataTypes returns a boolean if a field has been set.

### GetFunctionId

`func (o *DataTypesEntry) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *DataTypesEntry) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *DataTypesEntry) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


