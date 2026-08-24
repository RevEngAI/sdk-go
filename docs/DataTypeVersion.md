# DataTypeVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedAt** | Pointer to **time.Time** | When this version was written. Absent on a version that predates the recorded history. | [optional] 
**UpdatedBy** | Pointer to [**HistoryActor**](HistoryActor.md) | Who wrote this version. Absent on a version that predates the recorded history. | [optional] 
**Value** | [**DataTypeEntry**](DataTypeEntry.md) | The type as it stood in this version. | 

## Methods

### NewDataTypeVersion

`func NewDataTypeVersion(value DataTypeEntry, ) *DataTypeVersion`

NewDataTypeVersion instantiates a new DataTypeVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTypeVersionWithDefaults

`func NewDataTypeVersionWithDefaults() *DataTypeVersion`

NewDataTypeVersionWithDefaults instantiates a new DataTypeVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdatedAt

`func (o *DataTypeVersion) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DataTypeVersion) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DataTypeVersion) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DataTypeVersion) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DataTypeVersion) GetUpdatedBy() HistoryActor`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DataTypeVersion) GetUpdatedByOk() (*HistoryActor, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DataTypeVersion) SetUpdatedBy(v HistoryActor)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DataTypeVersion) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetValue

`func (o *DataTypeVersion) GetValue() DataTypeEntry`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DataTypeVersion) GetValueOk() (*DataTypeEntry, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DataTypeVersion) SetValue(v DataTypeEntry)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


