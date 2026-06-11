# ListAnalysisStringsOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Strings** | [**[]AnalysisStringItem**](AnalysisStringItem.md) |  | 
**TotalStrings** | **int64** |  | 

## Methods

### NewListAnalysisStringsOutputBody

`func NewListAnalysisStringsOutputBody(strings []AnalysisStringItem, totalStrings int64, ) *ListAnalysisStringsOutputBody`

NewListAnalysisStringsOutputBody instantiates a new ListAnalysisStringsOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAnalysisStringsOutputBodyWithDefaults

`func NewListAnalysisStringsOutputBodyWithDefaults() *ListAnalysisStringsOutputBody`

NewListAnalysisStringsOutputBodyWithDefaults instantiates a new ListAnalysisStringsOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrings

`func (o *ListAnalysisStringsOutputBody) GetStrings() []AnalysisStringItem`

GetStrings returns the Strings field if non-nil, zero value otherwise.

### GetStringsOk

`func (o *ListAnalysisStringsOutputBody) GetStringsOk() (*[]AnalysisStringItem, bool)`

GetStringsOk returns a tuple with the Strings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrings

`func (o *ListAnalysisStringsOutputBody) SetStrings(v []AnalysisStringItem)`

SetStrings sets Strings field to given value.


### SetStringsNil

`func (o *ListAnalysisStringsOutputBody) SetStringsNil(b bool)`

 SetStringsNil sets the value for Strings to be an explicit nil

### UnsetStrings
`func (o *ListAnalysisStringsOutputBody) UnsetStrings()`

UnsetStrings ensures that no value is present for Strings, not even an explicit nil
### GetTotalStrings

`func (o *ListAnalysisStringsOutputBody) GetTotalStrings() int64`

GetTotalStrings returns the TotalStrings field if non-nil, zero value otherwise.

### GetTotalStringsOk

`func (o *ListAnalysisStringsOutputBody) GetTotalStringsOk() (*int64, bool)`

GetTotalStringsOk returns a tuple with the TotalStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStrings

`func (o *ListAnalysisStringsOutputBody) SetTotalStrings(v int64)`

SetTotalStrings sets TotalStrings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


