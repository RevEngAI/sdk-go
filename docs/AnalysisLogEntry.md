# AnalysisLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | **string** | Severity | 
**Source** | **string** | Component that emitted the line | 
**Text** | **string** | Log line text | 
**Timestamp** | **time.Time** | When the line was emitted (UTC) | 

## Methods

### NewAnalysisLogEntry

`func NewAnalysisLogEntry(level string, source string, text string, timestamp time.Time, ) *AnalysisLogEntry`

NewAnalysisLogEntry instantiates a new AnalysisLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisLogEntryWithDefaults

`func NewAnalysisLogEntryWithDefaults() *AnalysisLogEntry`

NewAnalysisLogEntryWithDefaults instantiates a new AnalysisLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *AnalysisLogEntry) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *AnalysisLogEntry) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *AnalysisLogEntry) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetSource

`func (o *AnalysisLogEntry) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AnalysisLogEntry) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AnalysisLogEntry) SetSource(v string)`

SetSource sets Source field to given value.


### GetText

`func (o *AnalysisLogEntry) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *AnalysisLogEntry) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *AnalysisLogEntry) SetText(v string)`

SetText sets Text field to given value.


### GetTimestamp

`func (o *AnalysisLogEntry) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AnalysisLogEntry) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AnalysisLogEntry) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


