# EnumDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | [**[]DataTypeEnumValueEntry**](DataTypeEnumValueEntry.md) | The type&#39;s constants. | 

## Methods

### NewEnumDefinition

`func NewEnumDefinition(values []DataTypeEnumValueEntry, ) *EnumDefinition`

NewEnumDefinition instantiates a new EnumDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnumDefinitionWithDefaults

`func NewEnumDefinitionWithDefaults() *EnumDefinition`

NewEnumDefinitionWithDefaults instantiates a new EnumDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *EnumDefinition) GetValues() []DataTypeEnumValueEntry`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *EnumDefinition) GetValuesOk() (*[]DataTypeEnumValueEntry, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *EnumDefinition) SetValues(v []DataTypeEnumValueEntry)`

SetValues sets Values field to given value.


### SetValuesNil

`func (o *EnumDefinition) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *EnumDefinition) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


