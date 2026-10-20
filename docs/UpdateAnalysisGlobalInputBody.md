# UpdateAnalysisGlobalInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTypeId** | Pointer to **int64** | Analysis data type to assign. Omit to leave the type as it is. | [optional] 
**Name** | Pointer to **string** | New name for the global. Omit to leave the name as it is. | [optional] 

## Methods

### NewUpdateAnalysisGlobalInputBody

`func NewUpdateAnalysisGlobalInputBody() *UpdateAnalysisGlobalInputBody`

NewUpdateAnalysisGlobalInputBody instantiates a new UpdateAnalysisGlobalInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAnalysisGlobalInputBodyWithDefaults

`func NewUpdateAnalysisGlobalInputBodyWithDefaults() *UpdateAnalysisGlobalInputBody`

NewUpdateAnalysisGlobalInputBodyWithDefaults instantiates a new UpdateAnalysisGlobalInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTypeId

`func (o *UpdateAnalysisGlobalInputBody) GetDataTypeId() int64`

GetDataTypeId returns the DataTypeId field if non-nil, zero value otherwise.

### GetDataTypeIdOk

`func (o *UpdateAnalysisGlobalInputBody) GetDataTypeIdOk() (*int64, bool)`

GetDataTypeIdOk returns a tuple with the DataTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypeId

`func (o *UpdateAnalysisGlobalInputBody) SetDataTypeId(v int64)`

SetDataTypeId sets DataTypeId field to given value.

### HasDataTypeId

`func (o *UpdateAnalysisGlobalInputBody) HasDataTypeId() bool`

HasDataTypeId returns a boolean if a field has been set.

### GetName

`func (o *UpdateAnalysisGlobalInputBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAnalysisGlobalInputBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAnalysisGlobalInputBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAnalysisGlobalInputBody) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


