# ProgressMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | **string** | Severity level | 
**Step** | **string** | Step name when the message was emitted | 
**Text** | **string** | Message text | 
**Timestamp** | **time.Time** | When the message was emitted | 

## Methods

### NewProgressMessage

`func NewProgressMessage(level string, step string, text string, timestamp time.Time, ) *ProgressMessage`

NewProgressMessage instantiates a new ProgressMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgressMessageWithDefaults

`func NewProgressMessageWithDefaults() *ProgressMessage`

NewProgressMessageWithDefaults instantiates a new ProgressMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *ProgressMessage) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ProgressMessage) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ProgressMessage) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetStep

`func (o *ProgressMessage) GetStep() string`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *ProgressMessage) GetStepOk() (*string, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *ProgressMessage) SetStep(v string)`

SetStep sets Step field to given value.


### GetText

`func (o *ProgressMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ProgressMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ProgressMessage) SetText(v string)`

SetText sets Text field to given value.


### GetTimestamp

`func (o *ProgressMessage) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ProgressMessage) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ProgressMessage) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


