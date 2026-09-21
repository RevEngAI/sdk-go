# Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Machine-readable error code, shared with synchronous API errors. | 
**Detail** | Pointer to **string** | Additional context where helpful (quota numbers, validation specifics, etc.). | [optional] 
**DocUrl** | **string** | Link to documentation explaining this error and resolution steps. | 
**Message** | **string** | Brief description of the failure. | 
**Retryable** | **bool** | Whether retrying the operation might succeed. | 

## Methods

### NewStatus

`func NewStatus(code string, docUrl string, message string, retryable bool, ) *Status`

NewStatus instantiates a new Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusWithDefaults

`func NewStatusWithDefaults() *Status`

NewStatusWithDefaults instantiates a new Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Status) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Status) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Status) SetCode(v string)`

SetCode sets Code field to given value.


### GetDetail

`func (o *Status) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *Status) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *Status) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *Status) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetDocUrl

`func (o *Status) GetDocUrl() string`

GetDocUrl returns the DocUrl field if non-nil, zero value otherwise.

### GetDocUrlOk

`func (o *Status) GetDocUrlOk() (*string, bool)`

GetDocUrlOk returns a tuple with the DocUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocUrl

`func (o *Status) SetDocUrl(v string)`

SetDocUrl sets DocUrl field to given value.


### GetMessage

`func (o *Status) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Status) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Status) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRetryable

`func (o *Status) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *Status) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *Status) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


