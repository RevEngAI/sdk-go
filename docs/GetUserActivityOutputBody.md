# GetUserActivityOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activities** | [**[]ActivityBody**](ActivityBody.md) | Recent activity visible to the caller: their own private activity, their team&#39;s, and everyone&#39;s public activity | 

## Methods

### NewGetUserActivityOutputBody

`func NewGetUserActivityOutputBody(activities []ActivityBody, ) *GetUserActivityOutputBody`

NewGetUserActivityOutputBody instantiates a new GetUserActivityOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserActivityOutputBodyWithDefaults

`func NewGetUserActivityOutputBodyWithDefaults() *GetUserActivityOutputBody`

NewGetUserActivityOutputBodyWithDefaults instantiates a new GetUserActivityOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivities

`func (o *GetUserActivityOutputBody) GetActivities() []ActivityBody`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *GetUserActivityOutputBody) GetActivitiesOk() (*[]ActivityBody, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivities

`func (o *GetUserActivityOutputBody) SetActivities(v []ActivityBody)`

SetActivities sets Activities field to given value.


### SetActivitiesNil

`func (o *GetUserActivityOutputBody) SetActivitiesNil(b bool)`

 SetActivitiesNil sets the value for Activities to be an explicit nil

### UnsetActivities
`func (o *GetUserActivityOutputBody) UnsetActivities()`

UnsetActivities ensures that no value is present for Activities, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


