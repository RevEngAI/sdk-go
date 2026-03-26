# StringFunctions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | The value of the string literal | 
**Functions** | [**[]AppApiRestV2FunctionsResponsesFunction**](AppApiRestV2FunctionsResponsesFunction.md) | The function ids the string literal was found within | 
**Source** | Pointer to [**StringSource**](StringSource.md) | The source of the string | [optional] [default to STRINGSOURCE_SYSTEM]

## Methods

### NewStringFunctions

`func NewStringFunctions(value string, functions []AppApiRestV2FunctionsResponsesFunction, ) *StringFunctions`

NewStringFunctions instantiates a new StringFunctions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStringFunctionsWithDefaults

`func NewStringFunctionsWithDefaults() *StringFunctions`

NewStringFunctionsWithDefaults instantiates a new StringFunctions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *StringFunctions) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StringFunctions) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StringFunctions) SetValue(v string)`

SetValue sets Value field to given value.


### GetFunctions

`func (o *StringFunctions) GetFunctions() []AppApiRestV2FunctionsResponsesFunction`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *StringFunctions) GetFunctionsOk() (*[]AppApiRestV2FunctionsResponsesFunction, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *StringFunctions) SetFunctions(v []AppApiRestV2FunctionsResponsesFunction)`

SetFunctions sets Functions field to given value.


### GetSource

`func (o *StringFunctions) GetSource() StringSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StringFunctions) GetSourceOk() (*StringSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StringFunctions) SetSource(v StringSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *StringFunctions) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


