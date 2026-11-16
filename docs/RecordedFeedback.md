# RecordedFeedback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | **map[string]interface{}** | The context recorded with the feedback, or null when there is none | 
**Sentiment** | **NullableString** | The sentiment the caller recorded, or null when they have not left any | 

## Methods

### NewRecordedFeedback

`func NewRecordedFeedback(metadata map[string]interface{}, sentiment NullableString, ) *RecordedFeedback`

NewRecordedFeedback instantiates a new RecordedFeedback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordedFeedbackWithDefaults

`func NewRecordedFeedbackWithDefaults() *RecordedFeedback`

NewRecordedFeedbackWithDefaults instantiates a new RecordedFeedback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *RecordedFeedback) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RecordedFeedback) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RecordedFeedback) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetSentiment

`func (o *RecordedFeedback) GetSentiment() string`

GetSentiment returns the Sentiment field if non-nil, zero value otherwise.

### GetSentimentOk

`func (o *RecordedFeedback) GetSentimentOk() (*string, bool)`

GetSentimentOk returns a tuple with the Sentiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentiment

`func (o *RecordedFeedback) SetSentiment(v string)`

SetSentiment sets Sentiment field to given value.


### SetSentimentNil

`func (o *RecordedFeedback) SetSentimentNil(b bool)`

 SetSentimentNil sets the value for Sentiment to be an explicit nil

### UnsetSentiment
`func (o *RecordedFeedback) UnsetSentiment()`

UnsetSentiment ensures that no value is present for Sentiment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


