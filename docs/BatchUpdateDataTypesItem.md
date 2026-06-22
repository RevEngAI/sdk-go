# BatchUpdateDataTypesItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTypes** | **interface{}** |  | 
**DataTypesVersion** | **int64** | Current stored version. Pass 0 on the first write. | 
**FunctionId** | **int64** | Function ID | 

## Methods

### NewBatchUpdateDataTypesItem

`func NewBatchUpdateDataTypesItem(dataTypes interface{}, dataTypesVersion int64, functionId int64, ) *BatchUpdateDataTypesItem`

NewBatchUpdateDataTypesItem instantiates a new BatchUpdateDataTypesItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchUpdateDataTypesItemWithDefaults

`func NewBatchUpdateDataTypesItemWithDefaults() *BatchUpdateDataTypesItem`

NewBatchUpdateDataTypesItemWithDefaults instantiates a new BatchUpdateDataTypesItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTypes

`func (o *BatchUpdateDataTypesItem) GetDataTypes() interface{}`

GetDataTypes returns the DataTypes field if non-nil, zero value otherwise.

### GetDataTypesOk

`func (o *BatchUpdateDataTypesItem) GetDataTypesOk() (*interface{}, bool)`

GetDataTypesOk returns a tuple with the DataTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypes

`func (o *BatchUpdateDataTypesItem) SetDataTypes(v interface{})`

SetDataTypes sets DataTypes field to given value.


### SetDataTypesNil

`func (o *BatchUpdateDataTypesItem) SetDataTypesNil(b bool)`

 SetDataTypesNil sets the value for DataTypes to be an explicit nil

### UnsetDataTypes
`func (o *BatchUpdateDataTypesItem) UnsetDataTypes()`

UnsetDataTypes ensures that no value is present for DataTypes, not even an explicit nil
### GetDataTypesVersion

`func (o *BatchUpdateDataTypesItem) GetDataTypesVersion() int64`

GetDataTypesVersion returns the DataTypesVersion field if non-nil, zero value otherwise.

### GetDataTypesVersionOk

`func (o *BatchUpdateDataTypesItem) GetDataTypesVersionOk() (*int64, bool)`

GetDataTypesVersionOk returns a tuple with the DataTypesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypesVersion

`func (o *BatchUpdateDataTypesItem) SetDataTypesVersion(v int64)`

SetDataTypesVersion sets DataTypesVersion field to given value.


### GetFunctionId

`func (o *BatchUpdateDataTypesItem) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *BatchUpdateDataTypesItem) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *BatchUpdateDataTypesItem) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


