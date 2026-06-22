# BatchUpdateDataTypesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTypes** | Pointer to **interface{}** |  | [optional] 
**DataTypesVersion** | Pointer to **int64** | Version after update (present when status is &#39;updated&#39;) | [optional] 
**Error** | Pointer to **string** | Error message (present when status is &#39;error&#39;) | [optional] 
**FunctionId** | **int64** | Function ID | 
**Status** | **string** | Outcome for this function | 

## Methods

### NewBatchUpdateDataTypesResult

`func NewBatchUpdateDataTypesResult(functionId int64, status string, ) *BatchUpdateDataTypesResult`

NewBatchUpdateDataTypesResult instantiates a new BatchUpdateDataTypesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchUpdateDataTypesResultWithDefaults

`func NewBatchUpdateDataTypesResultWithDefaults() *BatchUpdateDataTypesResult`

NewBatchUpdateDataTypesResultWithDefaults instantiates a new BatchUpdateDataTypesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTypes

`func (o *BatchUpdateDataTypesResult) GetDataTypes() interface{}`

GetDataTypes returns the DataTypes field if non-nil, zero value otherwise.

### GetDataTypesOk

`func (o *BatchUpdateDataTypesResult) GetDataTypesOk() (*interface{}, bool)`

GetDataTypesOk returns a tuple with the DataTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypes

`func (o *BatchUpdateDataTypesResult) SetDataTypes(v interface{})`

SetDataTypes sets DataTypes field to given value.

### HasDataTypes

`func (o *BatchUpdateDataTypesResult) HasDataTypes() bool`

HasDataTypes returns a boolean if a field has been set.

### SetDataTypesNil

`func (o *BatchUpdateDataTypesResult) SetDataTypesNil(b bool)`

 SetDataTypesNil sets the value for DataTypes to be an explicit nil

### UnsetDataTypes
`func (o *BatchUpdateDataTypesResult) UnsetDataTypes()`

UnsetDataTypes ensures that no value is present for DataTypes, not even an explicit nil
### GetDataTypesVersion

`func (o *BatchUpdateDataTypesResult) GetDataTypesVersion() int64`

GetDataTypesVersion returns the DataTypesVersion field if non-nil, zero value otherwise.

### GetDataTypesVersionOk

`func (o *BatchUpdateDataTypesResult) GetDataTypesVersionOk() (*int64, bool)`

GetDataTypesVersionOk returns a tuple with the DataTypesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypesVersion

`func (o *BatchUpdateDataTypesResult) SetDataTypesVersion(v int64)`

SetDataTypesVersion sets DataTypesVersion field to given value.

### HasDataTypesVersion

`func (o *BatchUpdateDataTypesResult) HasDataTypesVersion() bool`

HasDataTypesVersion returns a boolean if a field has been set.

### GetError

`func (o *BatchUpdateDataTypesResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BatchUpdateDataTypesResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BatchUpdateDataTypesResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *BatchUpdateDataTypesResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFunctionId

`func (o *BatchUpdateDataTypesResult) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *BatchUpdateDataTypesResult) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *BatchUpdateDataTypesResult) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetStatus

`func (o *BatchUpdateDataTypesResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchUpdateDataTypesResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchUpdateDataTypesResult) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


