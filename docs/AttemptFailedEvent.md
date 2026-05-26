# AttemptFailedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempt** | **int32** |  | 
**Error** | **string** |  | 
**Seq** | **int32** |  | 
**Type** | **string** |  | 

## Methods

### NewAttemptFailedEvent

`func NewAttemptFailedEvent(attempt int32, error_ string, seq int32, type_ string, ) *AttemptFailedEvent`

NewAttemptFailedEvent instantiates a new AttemptFailedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttemptFailedEventWithDefaults

`func NewAttemptFailedEventWithDefaults() *AttemptFailedEvent`

NewAttemptFailedEventWithDefaults instantiates a new AttemptFailedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempt

`func (o *AttemptFailedEvent) GetAttempt() int32`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *AttemptFailedEvent) GetAttemptOk() (*int32, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *AttemptFailedEvent) SetAttempt(v int32)`

SetAttempt sets Attempt field to given value.


### GetError

`func (o *AttemptFailedEvent) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AttemptFailedEvent) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AttemptFailedEvent) SetError(v string)`

SetError sets Error field to given value.


### GetSeq

`func (o *AttemptFailedEvent) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *AttemptFailedEvent) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *AttemptFailedEvent) SetSeq(v int32)`

SetSeq sets Seq field to given value.


### GetType

`func (o *AttemptFailedEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttemptFailedEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttemptFailedEvent) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


