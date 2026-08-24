# FunctionSignatureVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedAt** | Pointer to **time.Time** | When this version was written. Absent on a version that predates the recorded history. | [optional] 
**UpdatedBy** | Pointer to [**HistoryActor**](HistoryActor.md) | Who wrote this version. Absent on a version that predates the recorded history. | [optional] 
**Value** | [**FunctionSignatureEntry**](FunctionSignatureEntry.md) | The signature as it stood in this version. | 

## Methods

### NewFunctionSignatureVersion

`func NewFunctionSignatureVersion(value FunctionSignatureEntry, ) *FunctionSignatureVersion`

NewFunctionSignatureVersion instantiates a new FunctionSignatureVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionSignatureVersionWithDefaults

`func NewFunctionSignatureVersionWithDefaults() *FunctionSignatureVersion`

NewFunctionSignatureVersionWithDefaults instantiates a new FunctionSignatureVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdatedAt

`func (o *FunctionSignatureVersion) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FunctionSignatureVersion) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FunctionSignatureVersion) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FunctionSignatureVersion) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *FunctionSignatureVersion) GetUpdatedBy() HistoryActor`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *FunctionSignatureVersion) GetUpdatedByOk() (*HistoryActor, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *FunctionSignatureVersion) SetUpdatedBy(v HistoryActor)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *FunctionSignatureVersion) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetValue

`func (o *FunctionSignatureVersion) GetValue() FunctionSignatureEntry`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FunctionSignatureVersion) GetValueOk() (*FunctionSignatureEntry, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FunctionSignatureVersion) SetValue(v FunctionSignatureEntry)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


