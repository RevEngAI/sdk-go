# AnalysisStringsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Strings** | [**[]StringFunctions**](StringFunctions.md) | The strings associated with the analysis | 
**TotalStrings** | **int32** | The total number of strings associated with this analysis | 

## Methods

### NewAnalysisStringsResponse

`func NewAnalysisStringsResponse(strings []StringFunctions, totalStrings int32, ) *AnalysisStringsResponse`

NewAnalysisStringsResponse instantiates a new AnalysisStringsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisStringsResponseWithDefaults

`func NewAnalysisStringsResponseWithDefaults() *AnalysisStringsResponse`

NewAnalysisStringsResponseWithDefaults instantiates a new AnalysisStringsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrings

`func (o *AnalysisStringsResponse) GetStrings() []StringFunctions`

GetStrings returns the Strings field if non-nil, zero value otherwise.

### GetStringsOk

`func (o *AnalysisStringsResponse) GetStringsOk() (*[]StringFunctions, bool)`

GetStringsOk returns a tuple with the Strings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrings

`func (o *AnalysisStringsResponse) SetStrings(v []StringFunctions)`

SetStrings sets Strings field to given value.


### GetTotalStrings

`func (o *AnalysisStringsResponse) GetTotalStrings() int32`

GetTotalStrings returns the TotalStrings field if non-nil, zero value otherwise.

### GetTotalStringsOk

`func (o *AnalysisStringsResponse) GetTotalStringsOk() (*int32, bool)`

GetTotalStringsOk returns a tuple with the TotalStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStrings

`func (o *AnalysisStringsResponse) SetTotalStrings(v int32)`

SetTotalStrings sets TotalStrings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


