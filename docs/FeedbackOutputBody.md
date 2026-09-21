# FeedbackOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sentiment** | **NullableString** | The sentiment the caller recorded, or null when they have not left any | 

## Methods

### NewFeedbackOutputBody

`func NewFeedbackOutputBody(sentiment NullableString, ) *FeedbackOutputBody`

NewFeedbackOutputBody instantiates a new FeedbackOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackOutputBodyWithDefaults

`func NewFeedbackOutputBodyWithDefaults() *FeedbackOutputBody`

NewFeedbackOutputBodyWithDefaults instantiates a new FeedbackOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSentiment

`func (o *FeedbackOutputBody) GetSentiment() string`

GetSentiment returns the Sentiment field if non-nil, zero value otherwise.

### GetSentimentOk

`func (o *FeedbackOutputBody) GetSentimentOk() (*string, bool)`

GetSentimentOk returns a tuple with the Sentiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentiment

`func (o *FeedbackOutputBody) SetSentiment(v string)`

SetSentiment sets Sentiment field to given value.


### SetSentimentNil

`func (o *FeedbackOutputBody) SetSentimentNil(b bool)`

 SetSentimentNil sets the value for Sentiment to be an explicit nil

### UnsetSentiment
`func (o *FeedbackOutputBody) UnsetSentiment()`

UnsetSentiment ensures that no value is present for Sentiment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


