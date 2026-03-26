# AnalysisStringInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**StringSource**](StringSource.md) | The source of the string | 
**Vaddr** | **int32** | The virtual address of the string | 
**Value** | **string** | The string literal value | 

## Methods

### NewAnalysisStringInput

`func NewAnalysisStringInput(source StringSource, vaddr int32, value string, ) *AnalysisStringInput`

NewAnalysisStringInput instantiates a new AnalysisStringInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisStringInputWithDefaults

`func NewAnalysisStringInputWithDefaults() *AnalysisStringInput`

NewAnalysisStringInputWithDefaults instantiates a new AnalysisStringInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *AnalysisStringInput) GetSource() StringSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AnalysisStringInput) GetSourceOk() (*StringSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AnalysisStringInput) SetSource(v StringSource)`

SetSource sets Source field to given value.


### GetVaddr

`func (o *AnalysisStringInput) GetVaddr() int32`

GetVaddr returns the Vaddr field if non-nil, zero value otherwise.

### GetVaddrOk

`func (o *AnalysisStringInput) GetVaddrOk() (*int32, bool)`

GetVaddrOk returns a tuple with the Vaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaddr

`func (o *AnalysisStringInput) SetVaddr(v int32)`

SetVaddr sets Vaddr field to given value.


### GetValue

`func (o *AnalysisStringInput) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AnalysisStringInput) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AnalysisStringInput) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


