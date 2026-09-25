# DailyCountBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** | Analyses the user created that day | 
**Day** | **time.Time** | Day this count covers, at midnight UTC | 

## Methods

### NewDailyCountBody

`func NewDailyCountBody(count int64, day time.Time, ) *DailyCountBody`

NewDailyCountBody instantiates a new DailyCountBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyCountBodyWithDefaults

`func NewDailyCountBodyWithDefaults() *DailyCountBody`

NewDailyCountBodyWithDefaults instantiates a new DailyCountBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *DailyCountBody) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DailyCountBody) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DailyCountBody) SetCount(v int64)`

SetCount sets Count field to given value.


### GetDay

`func (o *DailyCountBody) GetDay() time.Time`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *DailyCountBody) GetDayOk() (*time.Time, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *DailyCountBody) SetDay(v time.Time)`

SetDay sets Day field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


