# UpdateDataTypesOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTypes** | **interface{}** |  | 
**DataTypesVersion** | **int64** | Version of the stored function data types after the update | 
**FunctionId** | **int64** | Function ID | 

## Methods

### NewUpdateDataTypesOutputBody

`func NewUpdateDataTypesOutputBody(dataTypes interface{}, dataTypesVersion int64, functionId int64, ) *UpdateDataTypesOutputBody`

NewUpdateDataTypesOutputBody instantiates a new UpdateDataTypesOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDataTypesOutputBodyWithDefaults

`func NewUpdateDataTypesOutputBodyWithDefaults() *UpdateDataTypesOutputBody`

NewUpdateDataTypesOutputBodyWithDefaults instantiates a new UpdateDataTypesOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTypes

`func (o *UpdateDataTypesOutputBody) GetDataTypes() interface{}`

GetDataTypes returns the DataTypes field if non-nil, zero value otherwise.

### GetDataTypesOk

`func (o *UpdateDataTypesOutputBody) GetDataTypesOk() (*interface{}, bool)`

GetDataTypesOk returns a tuple with the DataTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypes

`func (o *UpdateDataTypesOutputBody) SetDataTypes(v interface{})`

SetDataTypes sets DataTypes field to given value.


### SetDataTypesNil

`func (o *UpdateDataTypesOutputBody) SetDataTypesNil(b bool)`

 SetDataTypesNil sets the value for DataTypes to be an explicit nil

### UnsetDataTypes
`func (o *UpdateDataTypesOutputBody) UnsetDataTypes()`

UnsetDataTypes ensures that no value is present for DataTypes, not even an explicit nil
### GetDataTypesVersion

`func (o *UpdateDataTypesOutputBody) GetDataTypesVersion() int64`

GetDataTypesVersion returns the DataTypesVersion field if non-nil, zero value otherwise.

### GetDataTypesVersionOk

`func (o *UpdateDataTypesOutputBody) GetDataTypesVersionOk() (*int64, bool)`

GetDataTypesVersionOk returns a tuple with the DataTypesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypesVersion

`func (o *UpdateDataTypesOutputBody) SetDataTypesVersion(v int64)`

SetDataTypesVersion sets DataTypesVersion field to given value.


### GetFunctionId

`func (o *UpdateDataTypesOutputBody) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *UpdateDataTypesOutputBody) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *UpdateDataTypesOutputBody) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


