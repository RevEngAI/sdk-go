# ListAnalysisFunctionsOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Functions** | [**[]AnalysisFunctionEntry**](AnalysisFunctionEntry.md) |  | 
**TotalCount** | **int64** | Total functions in the analysis, ignoring pagination. | 

## Methods

### NewListAnalysisFunctionsOutputBody

`func NewListAnalysisFunctionsOutputBody(functions []AnalysisFunctionEntry, totalCount int64, ) *ListAnalysisFunctionsOutputBody`

NewListAnalysisFunctionsOutputBody instantiates a new ListAnalysisFunctionsOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAnalysisFunctionsOutputBodyWithDefaults

`func NewListAnalysisFunctionsOutputBodyWithDefaults() *ListAnalysisFunctionsOutputBody`

NewListAnalysisFunctionsOutputBodyWithDefaults instantiates a new ListAnalysisFunctionsOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctions

`func (o *ListAnalysisFunctionsOutputBody) GetFunctions() []AnalysisFunctionEntry`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *ListAnalysisFunctionsOutputBody) GetFunctionsOk() (*[]AnalysisFunctionEntry, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *ListAnalysisFunctionsOutputBody) SetFunctions(v []AnalysisFunctionEntry)`

SetFunctions sets Functions field to given value.


### SetFunctionsNil

`func (o *ListAnalysisFunctionsOutputBody) SetFunctionsNil(b bool)`

 SetFunctionsNil sets the value for Functions to be an explicit nil

### UnsetFunctions
`func (o *ListAnalysisFunctionsOutputBody) UnsetFunctions()`

UnsetFunctions ensures that no value is present for Functions, not even an explicit nil
### GetTotalCount

`func (o *ListAnalysisFunctionsOutputBody) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListAnalysisFunctionsOutputBody) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListAnalysisFunctionsOutputBody) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


