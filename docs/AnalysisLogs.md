# AnalysisLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageCount** | **int64** |  | 
**Messages** | [**[]AnalysisLogMessage**](AnalysisLogMessage.md) |  | 

## Methods

### NewAnalysisLogs

`func NewAnalysisLogs(messageCount int64, messages []AnalysisLogMessage, ) *AnalysisLogs`

NewAnalysisLogs instantiates a new AnalysisLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisLogsWithDefaults

`func NewAnalysisLogsWithDefaults() *AnalysisLogs`

NewAnalysisLogsWithDefaults instantiates a new AnalysisLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageCount

`func (o *AnalysisLogs) GetMessageCount() int64`

GetMessageCount returns the MessageCount field if non-nil, zero value otherwise.

### GetMessageCountOk

`func (o *AnalysisLogs) GetMessageCountOk() (*int64, bool)`

GetMessageCountOk returns a tuple with the MessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCount

`func (o *AnalysisLogs) SetMessageCount(v int64)`

SetMessageCount sets MessageCount field to given value.


### GetMessages

`func (o *AnalysisLogs) GetMessages() []AnalysisLogMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *AnalysisLogs) GetMessagesOk() (*[]AnalysisLogMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *AnalysisLogs) SetMessages(v []AnalysisLogMessage)`

SetMessages sets Messages field to given value.


### SetMessagesNil

`func (o *AnalysisLogs) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *AnalysisLogs) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


