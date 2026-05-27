# DecompFailedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempt** | **int32** |  | 
**Error** | **string** |  | 
**ErrorCode** | Pointer to **string** |  | [optional] 
**Seq** | **int32** |  | 
**Type** | **string** |  | 

## Methods

### NewDecompFailedEvent

`func NewDecompFailedEvent(attempt int32, error_ string, seq int32, type_ string, ) *DecompFailedEvent`

NewDecompFailedEvent instantiates a new DecompFailedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecompFailedEventWithDefaults

`func NewDecompFailedEventWithDefaults() *DecompFailedEvent`

NewDecompFailedEventWithDefaults instantiates a new DecompFailedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempt

`func (o *DecompFailedEvent) GetAttempt() int32`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *DecompFailedEvent) GetAttemptOk() (*int32, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *DecompFailedEvent) SetAttempt(v int32)`

SetAttempt sets Attempt field to given value.


### GetError

`func (o *DecompFailedEvent) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DecompFailedEvent) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DecompFailedEvent) SetError(v string)`

SetError sets Error field to given value.


### GetErrorCode

`func (o *DecompFailedEvent) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *DecompFailedEvent) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *DecompFailedEvent) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *DecompFailedEvent) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetSeq

`func (o *DecompFailedEvent) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *DecompFailedEvent) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *DecompFailedEvent) SetSeq(v int32)`

SetSeq sets Seq field to given value.


### GetType

`func (o *DecompFailedEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DecompFailedEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DecompFailedEvent) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


