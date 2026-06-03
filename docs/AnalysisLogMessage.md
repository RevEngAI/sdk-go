# AnalysisLogMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **NullableString** |  | 
**Time** | **NullableString** |  | 

## Methods

### NewAnalysisLogMessage

`func NewAnalysisLogMessage(message NullableString, time NullableString, ) *AnalysisLogMessage`

NewAnalysisLogMessage instantiates a new AnalysisLogMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisLogMessageWithDefaults

`func NewAnalysisLogMessageWithDefaults() *AnalysisLogMessage`

NewAnalysisLogMessageWithDefaults instantiates a new AnalysisLogMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AnalysisLogMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AnalysisLogMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AnalysisLogMessage) SetMessage(v string)`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *AnalysisLogMessage) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *AnalysisLogMessage) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetTime

`func (o *AnalysisLogMessage) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AnalysisLogMessage) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AnalysisLogMessage) SetTime(v string)`

SetTime sets Time field to given value.


### SetTimeNil

`func (o *AnalysisLogMessage) SetTimeNil(b bool)`

 SetTimeNil sets the value for Time to be an explicit nil

### UnsetTime
`func (o *AnalysisLogMessage) UnsetTime()`

UnsetTime ensures that no value is present for Time, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


