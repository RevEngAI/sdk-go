# SourceDeltaEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempt** | **int32** |  | 
**Content** | **string** |  | 
**Seq** | **int32** |  | 
**Type** | **string** |  | 

## Methods

### NewSourceDeltaEvent

`func NewSourceDeltaEvent(attempt int32, content string, seq int32, type_ string, ) *SourceDeltaEvent`

NewSourceDeltaEvent instantiates a new SourceDeltaEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceDeltaEventWithDefaults

`func NewSourceDeltaEventWithDefaults() *SourceDeltaEvent`

NewSourceDeltaEventWithDefaults instantiates a new SourceDeltaEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempt

`func (o *SourceDeltaEvent) GetAttempt() int32`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *SourceDeltaEvent) GetAttemptOk() (*int32, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *SourceDeltaEvent) SetAttempt(v int32)`

SetAttempt sets Attempt field to given value.


### GetContent

`func (o *SourceDeltaEvent) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SourceDeltaEvent) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SourceDeltaEvent) SetContent(v string)`

SetContent sets Content field to given value.


### GetSeq

`func (o *SourceDeltaEvent) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *SourceDeltaEvent) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *SourceDeltaEvent) SetSeq(v int32)`

SetSeq sets Seq field to given value.


### GetType

`func (o *SourceDeltaEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SourceDeltaEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SourceDeltaEvent) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


