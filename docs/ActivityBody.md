# ActivityBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | **string** | The kind of action taken | 
**ActivityScope** | **string** | Who can see this activity entry | 
**CreatedAt** | **time.Time** | When the action happened | 
**Message** | **string** | Human-readable description of the action | 
**Sources** | **string** | The resource kind the action applied to | 
**Username** | **string** | The user who performed the action | 

## Methods

### NewActivityBody

`func NewActivityBody(actions string, activityScope string, createdAt time.Time, message string, sources string, username string, ) *ActivityBody`

NewActivityBody instantiates a new ActivityBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityBodyWithDefaults

`func NewActivityBodyWithDefaults() *ActivityBody`

NewActivityBodyWithDefaults instantiates a new ActivityBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ActivityBody) GetActions() string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ActivityBody) GetActionsOk() (*string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ActivityBody) SetActions(v string)`

SetActions sets Actions field to given value.


### GetActivityScope

`func (o *ActivityBody) GetActivityScope() string`

GetActivityScope returns the ActivityScope field if non-nil, zero value otherwise.

### GetActivityScopeOk

`func (o *ActivityBody) GetActivityScopeOk() (*string, bool)`

GetActivityScopeOk returns a tuple with the ActivityScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityScope

`func (o *ActivityBody) SetActivityScope(v string)`

SetActivityScope sets ActivityScope field to given value.


### GetCreatedAt

`func (o *ActivityBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ActivityBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ActivityBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetMessage

`func (o *ActivityBody) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ActivityBody) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ActivityBody) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSources

`func (o *ActivityBody) GetSources() string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *ActivityBody) GetSourcesOk() (*string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *ActivityBody) SetSources(v string)`

SetSources sets Sources field to given value.


### GetUsername

`func (o *ActivityBody) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ActivityBody) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ActivityBody) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


