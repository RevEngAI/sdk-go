# AnalysisStringItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Functions** | [**[]AnalysisStringFunction**](AnalysisStringFunction.md) |  | 
**Source** | **string** |  | 
**Value** | **string** |  | 

## Methods

### NewAnalysisStringItem

`func NewAnalysisStringItem(functions []AnalysisStringFunction, source string, value string, ) *AnalysisStringItem`

NewAnalysisStringItem instantiates a new AnalysisStringItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisStringItemWithDefaults

`func NewAnalysisStringItemWithDefaults() *AnalysisStringItem`

NewAnalysisStringItemWithDefaults instantiates a new AnalysisStringItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctions

`func (o *AnalysisStringItem) GetFunctions() []AnalysisStringFunction`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *AnalysisStringItem) GetFunctionsOk() (*[]AnalysisStringFunction, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *AnalysisStringItem) SetFunctions(v []AnalysisStringFunction)`

SetFunctions sets Functions field to given value.


### SetFunctionsNil

`func (o *AnalysisStringItem) SetFunctionsNil(b bool)`

 SetFunctionsNil sets the value for Functions to be an explicit nil

### UnsetFunctions
`func (o *AnalysisStringItem) UnsetFunctions()`

UnsetFunctions ensures that no value is present for Functions, not even an explicit nil
### GetSource

`func (o *AnalysisStringItem) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AnalysisStringItem) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AnalysisStringItem) SetSource(v string)`

SetSource sets Source field to given value.


### GetValue

`func (o *AnalysisStringItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AnalysisStringItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AnalysisStringItem) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


