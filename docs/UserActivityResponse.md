# UserActivityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | **string** |  | 
**ActivityScope** | **string** |  | 
**Creation** | **time.Time** |  | 
**Message** | **string** |  | 
**Sources** | **string** |  | 
**Username** | **string** |  | 

## Methods

### NewUserActivityResponse

`func NewUserActivityResponse(actions string, activityScope string, creation time.Time, message string, sources string, username string, ) *UserActivityResponse`

NewUserActivityResponse instantiates a new UserActivityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserActivityResponseWithDefaults

`func NewUserActivityResponseWithDefaults() *UserActivityResponse`

NewUserActivityResponseWithDefaults instantiates a new UserActivityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *UserActivityResponse) GetActions() string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *UserActivityResponse) GetActionsOk() (*string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *UserActivityResponse) SetActions(v string)`

SetActions sets Actions field to given value.


### GetActivityScope

`func (o *UserActivityResponse) GetActivityScope() string`

GetActivityScope returns the ActivityScope field if non-nil, zero value otherwise.

### GetActivityScopeOk

`func (o *UserActivityResponse) GetActivityScopeOk() (*string, bool)`

GetActivityScopeOk returns a tuple with the ActivityScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityScope

`func (o *UserActivityResponse) SetActivityScope(v string)`

SetActivityScope sets ActivityScope field to given value.


### GetCreation

`func (o *UserActivityResponse) GetCreation() time.Time`

GetCreation returns the Creation field if non-nil, zero value otherwise.

### GetCreationOk

`func (o *UserActivityResponse) GetCreationOk() (*time.Time, bool)`

GetCreationOk returns a tuple with the Creation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreation

`func (o *UserActivityResponse) SetCreation(v time.Time)`

SetCreation sets Creation field to given value.


### GetMessage

`func (o *UserActivityResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UserActivityResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UserActivityResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSources

`func (o *UserActivityResponse) GetSources() string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *UserActivityResponse) GetSourcesOk() (*string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *UserActivityResponse) SetSources(v string)`

SetSources sets Sources field to given value.


### GetUsername

`func (o *UserActivityResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserActivityResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserActivityResponse) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


