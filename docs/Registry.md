# Registry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** |  | 
**Key** | **string** |  | 
**ValueName** | **NullableString** |  | 
**Value** | **NullableString** |  | 

## Methods

### NewRegistry

`func NewRegistry(method string, key string, valueName NullableString, value NullableString, ) *Registry`

NewRegistry instantiates a new Registry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryWithDefaults

`func NewRegistryWithDefaults() *Registry`

NewRegistryWithDefaults instantiates a new Registry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *Registry) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *Registry) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *Registry) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetKey

`func (o *Registry) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Registry) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Registry) SetKey(v string)`

SetKey sets Key field to given value.


### GetValueName

`func (o *Registry) GetValueName() string`

GetValueName returns the ValueName field if non-nil, zero value otherwise.

### GetValueNameOk

`func (o *Registry) GetValueNameOk() (*string, bool)`

GetValueNameOk returns a tuple with the ValueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueName

`func (o *Registry) SetValueName(v string)`

SetValueName sets ValueName field to given value.


### SetValueNameNil

`func (o *Registry) SetValueNameNil(b bool)`

 SetValueNameNil sets the value for ValueName to be an explicit nil

### UnsetValueName
`func (o *Registry) UnsetValueName()`

UnsetValueName ensures that no value is present for ValueName, not even an explicit nil
### GetValue

`func (o *Registry) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Registry) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Registry) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *Registry) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Registry) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


