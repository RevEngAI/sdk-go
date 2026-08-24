# DataTypeEnumValueEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Constant name. | 
**Value** | **string** | Constant value, a decimal integer as a string, with no leading zeros. A string because the value may be negative or exceed 64 unsigned bits, which a JSON number cannot carry safely. | 

## Methods

### NewDataTypeEnumValueEntry

`func NewDataTypeEnumValueEntry(name string, value string, ) *DataTypeEnumValueEntry`

NewDataTypeEnumValueEntry instantiates a new DataTypeEnumValueEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTypeEnumValueEntryWithDefaults

`func NewDataTypeEnumValueEntryWithDefaults() *DataTypeEnumValueEntry`

NewDataTypeEnumValueEntryWithDefaults instantiates a new DataTypeEnumValueEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DataTypeEnumValueEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataTypeEnumValueEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataTypeEnumValueEntry) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *DataTypeEnumValueEntry) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DataTypeEnumValueEntry) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DataTypeEnumValueEntry) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


