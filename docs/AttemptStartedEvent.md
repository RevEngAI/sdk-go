# AttemptStartedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempt** | **int32** |  | 
**Seq** | **int32** |  | 
**Type** | **string** |  | 

## Methods

### NewAttemptStartedEvent

`func NewAttemptStartedEvent(attempt int32, seq int32, type_ string, ) *AttemptStartedEvent`

NewAttemptStartedEvent instantiates a new AttemptStartedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttemptStartedEventWithDefaults

`func NewAttemptStartedEventWithDefaults() *AttemptStartedEvent`

NewAttemptStartedEventWithDefaults instantiates a new AttemptStartedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempt

`func (o *AttemptStartedEvent) GetAttempt() int32`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *AttemptStartedEvent) GetAttemptOk() (*int32, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *AttemptStartedEvent) SetAttempt(v int32)`

SetAttempt sets Attempt field to given value.


### GetSeq

`func (o *AttemptStartedEvent) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *AttemptStartedEvent) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *AttemptStartedEvent) SetSeq(v int32)`

SetSeq sets Seq field to given value.


### GetType

`func (o *AttemptStartedEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttemptStartedEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttemptStartedEvent) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


