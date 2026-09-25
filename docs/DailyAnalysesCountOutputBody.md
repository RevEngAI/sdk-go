# DailyAnalysesCountOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DailyCounts** | [**[]DailyCountBody**](DailyCountBody.md) | One entry per day of the last 31 days, oldest first | 

## Methods

### NewDailyAnalysesCountOutputBody

`func NewDailyAnalysesCountOutputBody(dailyCounts []DailyCountBody, ) *DailyAnalysesCountOutputBody`

NewDailyAnalysesCountOutputBody instantiates a new DailyAnalysesCountOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyAnalysesCountOutputBodyWithDefaults

`func NewDailyAnalysesCountOutputBodyWithDefaults() *DailyAnalysesCountOutputBody`

NewDailyAnalysesCountOutputBodyWithDefaults instantiates a new DailyAnalysesCountOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDailyCounts

`func (o *DailyAnalysesCountOutputBody) GetDailyCounts() []DailyCountBody`

GetDailyCounts returns the DailyCounts field if non-nil, zero value otherwise.

### GetDailyCountsOk

`func (o *DailyAnalysesCountOutputBody) GetDailyCountsOk() (*[]DailyCountBody, bool)`

GetDailyCountsOk returns a tuple with the DailyCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyCounts

`func (o *DailyAnalysesCountOutputBody) SetDailyCounts(v []DailyCountBody)`

SetDailyCounts sets DailyCounts field to given value.


### SetDailyCountsNil

`func (o *DailyAnalysesCountOutputBody) SetDailyCountsNil(b bool)`

 SetDailyCountsNil sets the value for DailyCounts to be an explicit nil

### UnsetDailyCounts
`func (o *DailyAnalysesCountOutputBody) UnsetDailyCounts()`

UnsetDailyCounts ensures that no value is present for DailyCounts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


