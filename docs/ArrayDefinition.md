# ArrayDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | Element count. Zero is a genuinely empty trailing array; absent means it could not be determined. | [optional] 
**ElementDataTypeId** | Pointer to **int64** | The element type. | [optional] 

## Methods

### NewArrayDefinition

`func NewArrayDefinition() *ArrayDefinition`

NewArrayDefinition instantiates a new ArrayDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArrayDefinitionWithDefaults

`func NewArrayDefinitionWithDefaults() *ArrayDefinition`

NewArrayDefinitionWithDefaults instantiates a new ArrayDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ArrayDefinition) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ArrayDefinition) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ArrayDefinition) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ArrayDefinition) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetElementDataTypeId

`func (o *ArrayDefinition) GetElementDataTypeId() int64`

GetElementDataTypeId returns the ElementDataTypeId field if non-nil, zero value otherwise.

### GetElementDataTypeIdOk

`func (o *ArrayDefinition) GetElementDataTypeIdOk() (*int64, bool)`

GetElementDataTypeIdOk returns a tuple with the ElementDataTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementDataTypeId

`func (o *ArrayDefinition) SetElementDataTypeId(v int64)`

SetElementDataTypeId sets ElementDataTypeId field to given value.

### HasElementDataTypeId

`func (o *ArrayDefinition) HasElementDataTypeId() bool`

HasElementDataTypeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


