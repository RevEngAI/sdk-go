# SubmitFeedbackBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentRoute** | **string** | The route the caller was on when they submitted feedback | 
**Feedback** | **string** | The feedback text | 
**ScreenCaptureUrl** | Pointer to **string** | Optional URL to a screen capture related to the feedback | [optional] 

## Methods

### NewSubmitFeedbackBody

`func NewSubmitFeedbackBody(currentRoute string, feedback string, ) *SubmitFeedbackBody`

NewSubmitFeedbackBody instantiates a new SubmitFeedbackBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitFeedbackBodyWithDefaults

`func NewSubmitFeedbackBodyWithDefaults() *SubmitFeedbackBody`

NewSubmitFeedbackBodyWithDefaults instantiates a new SubmitFeedbackBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentRoute

`func (o *SubmitFeedbackBody) GetCurrentRoute() string`

GetCurrentRoute returns the CurrentRoute field if non-nil, zero value otherwise.

### GetCurrentRouteOk

`func (o *SubmitFeedbackBody) GetCurrentRouteOk() (*string, bool)`

GetCurrentRouteOk returns a tuple with the CurrentRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRoute

`func (o *SubmitFeedbackBody) SetCurrentRoute(v string)`

SetCurrentRoute sets CurrentRoute field to given value.


### GetFeedback

`func (o *SubmitFeedbackBody) GetFeedback() string`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *SubmitFeedbackBody) GetFeedbackOk() (*string, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedback

`func (o *SubmitFeedbackBody) SetFeedback(v string)`

SetFeedback sets Feedback field to given value.


### GetScreenCaptureUrl

`func (o *SubmitFeedbackBody) GetScreenCaptureUrl() string`

GetScreenCaptureUrl returns the ScreenCaptureUrl field if non-nil, zero value otherwise.

### GetScreenCaptureUrlOk

`func (o *SubmitFeedbackBody) GetScreenCaptureUrlOk() (*string, bool)`

GetScreenCaptureUrlOk returns a tuple with the ScreenCaptureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenCaptureUrl

`func (o *SubmitFeedbackBody) SetScreenCaptureUrl(v string)`

SetScreenCaptureUrl sets ScreenCaptureUrl field to given value.

### HasScreenCaptureUrl

`func (o *SubmitFeedbackBody) HasScreenCaptureUrl() bool`

HasScreenCaptureUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


