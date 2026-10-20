# ListAnalysisGlobalsOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Globals** | [**[]GlobalVariable**](GlobalVariable.md) |  | 
**Total** | **int64** |  | 
**TotalUnfiltered** | **int64** |  | 

## Methods

### NewListAnalysisGlobalsOutputBody

`func NewListAnalysisGlobalsOutputBody(globals []GlobalVariable, total int64, totalUnfiltered int64, ) *ListAnalysisGlobalsOutputBody`

NewListAnalysisGlobalsOutputBody instantiates a new ListAnalysisGlobalsOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAnalysisGlobalsOutputBodyWithDefaults

`func NewListAnalysisGlobalsOutputBodyWithDefaults() *ListAnalysisGlobalsOutputBody`

NewListAnalysisGlobalsOutputBodyWithDefaults instantiates a new ListAnalysisGlobalsOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobals

`func (o *ListAnalysisGlobalsOutputBody) GetGlobals() []GlobalVariable`

GetGlobals returns the Globals field if non-nil, zero value otherwise.

### GetGlobalsOk

`func (o *ListAnalysisGlobalsOutputBody) GetGlobalsOk() (*[]GlobalVariable, bool)`

GetGlobalsOk returns a tuple with the Globals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobals

`func (o *ListAnalysisGlobalsOutputBody) SetGlobals(v []GlobalVariable)`

SetGlobals sets Globals field to given value.


### SetGlobalsNil

`func (o *ListAnalysisGlobalsOutputBody) SetGlobalsNil(b bool)`

 SetGlobalsNil sets the value for Globals to be an explicit nil

### UnsetGlobals
`func (o *ListAnalysisGlobalsOutputBody) UnsetGlobals()`

UnsetGlobals ensures that no value is present for Globals, not even an explicit nil
### GetTotal

`func (o *ListAnalysisGlobalsOutputBody) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListAnalysisGlobalsOutputBody) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListAnalysisGlobalsOutputBody) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetTotalUnfiltered

`func (o *ListAnalysisGlobalsOutputBody) GetTotalUnfiltered() int64`

GetTotalUnfiltered returns the TotalUnfiltered field if non-nil, zero value otherwise.

### GetTotalUnfilteredOk

`func (o *ListAnalysisGlobalsOutputBody) GetTotalUnfilteredOk() (*int64, bool)`

GetTotalUnfilteredOk returns a tuple with the TotalUnfiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUnfiltered

`func (o *ListAnalysisGlobalsOutputBody) SetTotalUnfiltered(v int64)`

SetTotalUnfiltered sets TotalUnfiltered field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


