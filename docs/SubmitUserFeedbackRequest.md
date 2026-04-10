# SubmitUserFeedbackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentRoute** | **string** | The route from where the feedback was submitted | 
**Feedback** | **string** | The user&#39;s feedback | 
**ScreenCaptureUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSubmitUserFeedbackRequest

`func NewSubmitUserFeedbackRequest(currentRoute string, feedback string, ) *SubmitUserFeedbackRequest`

NewSubmitUserFeedbackRequest instantiates a new SubmitUserFeedbackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitUserFeedbackRequestWithDefaults

`func NewSubmitUserFeedbackRequestWithDefaults() *SubmitUserFeedbackRequest`

NewSubmitUserFeedbackRequestWithDefaults instantiates a new SubmitUserFeedbackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentRoute

`func (o *SubmitUserFeedbackRequest) GetCurrentRoute() string`

GetCurrentRoute returns the CurrentRoute field if non-nil, zero value otherwise.

### GetCurrentRouteOk

`func (o *SubmitUserFeedbackRequest) GetCurrentRouteOk() (*string, bool)`

GetCurrentRouteOk returns a tuple with the CurrentRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRoute

`func (o *SubmitUserFeedbackRequest) SetCurrentRoute(v string)`

SetCurrentRoute sets CurrentRoute field to given value.


### GetFeedback

`func (o *SubmitUserFeedbackRequest) GetFeedback() string`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *SubmitUserFeedbackRequest) GetFeedbackOk() (*string, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedback

`func (o *SubmitUserFeedbackRequest) SetFeedback(v string)`

SetFeedback sets Feedback field to given value.


### GetScreenCaptureUrl

`func (o *SubmitUserFeedbackRequest) GetScreenCaptureUrl() string`

GetScreenCaptureUrl returns the ScreenCaptureUrl field if non-nil, zero value otherwise.

### GetScreenCaptureUrlOk

`func (o *SubmitUserFeedbackRequest) GetScreenCaptureUrlOk() (*string, bool)`

GetScreenCaptureUrlOk returns a tuple with the ScreenCaptureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenCaptureUrl

`func (o *SubmitUserFeedbackRequest) SetScreenCaptureUrl(v string)`

SetScreenCaptureUrl sets ScreenCaptureUrl field to given value.

### HasScreenCaptureUrl

`func (o *SubmitUserFeedbackRequest) HasScreenCaptureUrl() bool`

HasScreenCaptureUrl returns a boolean if a field has been set.

### SetScreenCaptureUrlNil

`func (o *SubmitUserFeedbackRequest) SetScreenCaptureUrlNil(b bool)`

 SetScreenCaptureUrlNil sets the value for ScreenCaptureUrl to be an explicit nil

### UnsetScreenCaptureUrl
`func (o *SubmitUserFeedbackRequest) UnsetScreenCaptureUrl()`

UnsetScreenCaptureUrl ensures that no value is present for ScreenCaptureUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


