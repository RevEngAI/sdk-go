# UpdateDataTypesInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTypes** | **interface{}** |  | 
**DataTypesVersion** | **int64** | Current version of the function data types. The update is rejected if the stored version has moved on. Pass 0 on the first write. | 

## Methods

### NewUpdateDataTypesInputBody

`func NewUpdateDataTypesInputBody(dataTypes interface{}, dataTypesVersion int64, ) *UpdateDataTypesInputBody`

NewUpdateDataTypesInputBody instantiates a new UpdateDataTypesInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDataTypesInputBodyWithDefaults

`func NewUpdateDataTypesInputBodyWithDefaults() *UpdateDataTypesInputBody`

NewUpdateDataTypesInputBodyWithDefaults instantiates a new UpdateDataTypesInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTypes

`func (o *UpdateDataTypesInputBody) GetDataTypes() interface{}`

GetDataTypes returns the DataTypes field if non-nil, zero value otherwise.

### GetDataTypesOk

`func (o *UpdateDataTypesInputBody) GetDataTypesOk() (*interface{}, bool)`

GetDataTypesOk returns a tuple with the DataTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypes

`func (o *UpdateDataTypesInputBody) SetDataTypes(v interface{})`

SetDataTypes sets DataTypes field to given value.


### SetDataTypesNil

`func (o *UpdateDataTypesInputBody) SetDataTypesNil(b bool)`

 SetDataTypesNil sets the value for DataTypes to be an explicit nil

### UnsetDataTypes
`func (o *UpdateDataTypesInputBody) UnsetDataTypes()`

UnsetDataTypes ensures that no value is present for DataTypes, not even an explicit nil
### GetDataTypesVersion

`func (o *UpdateDataTypesInputBody) GetDataTypesVersion() int64`

GetDataTypesVersion returns the DataTypesVersion field if non-nil, zero value otherwise.

### GetDataTypesVersionOk

`func (o *UpdateDataTypesInputBody) GetDataTypesVersionOk() (*int64, bool)`

GetDataTypesVersionOk returns a tuple with the DataTypesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypesVersion

`func (o *UpdateDataTypesInputBody) SetDataTypesVersion(v int64)`

SetDataTypesVersion sets DataTypesVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


