# ExtractionFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Why this file failed to extract. | 
**Retryable** | **bool** | Whether re-submitting the extraction might resolve this failure. | 

## Methods

### NewExtractionFailure

`func NewExtractionFailure(message string, retryable bool, ) *ExtractionFailure`

NewExtractionFailure instantiates a new ExtractionFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtractionFailureWithDefaults

`func NewExtractionFailureWithDefaults() *ExtractionFailure`

NewExtractionFailureWithDefaults instantiates a new ExtractionFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ExtractionFailure) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ExtractionFailure) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ExtractionFailure) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRetryable

`func (o *ExtractionFailure) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *ExtractionFailure) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *ExtractionFailure) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


