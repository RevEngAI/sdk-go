# FeedbackSubmission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]interface{}** | Context the client captured with the feedback, stored as-is (e.g. the binary hash and the matches the caller was shown) | [optional] 
**Sentiment** | **string** | How useful the caller found the result | 

## Methods

### NewFeedbackSubmission

`func NewFeedbackSubmission(sentiment string, ) *FeedbackSubmission`

NewFeedbackSubmission instantiates a new FeedbackSubmission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackSubmissionWithDefaults

`func NewFeedbackSubmissionWithDefaults() *FeedbackSubmission`

NewFeedbackSubmissionWithDefaults instantiates a new FeedbackSubmission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *FeedbackSubmission) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FeedbackSubmission) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FeedbackSubmission) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FeedbackSubmission) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSentiment

`func (o *FeedbackSubmission) GetSentiment() string`

GetSentiment returns the Sentiment field if non-nil, zero value otherwise.

### GetSentimentOk

`func (o *FeedbackSubmission) GetSentimentOk() (*string, bool)`

GetSentimentOk returns a tuple with the Sentiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentiment

`func (o *FeedbackSubmission) SetSentiment(v string)`

SetSentiment sets Sentiment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


