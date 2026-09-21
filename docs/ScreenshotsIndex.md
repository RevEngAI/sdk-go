# ScreenshotsIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** |  | 
**Screenshots** | [**[]ScreenshotEntry**](ScreenshotEntry.md) |  | 

## Methods

### NewScreenshotsIndex

`func NewScreenshotsIndex(count int64, screenshots []ScreenshotEntry, ) *ScreenshotsIndex`

NewScreenshotsIndex instantiates a new ScreenshotsIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScreenshotsIndexWithDefaults

`func NewScreenshotsIndexWithDefaults() *ScreenshotsIndex`

NewScreenshotsIndexWithDefaults instantiates a new ScreenshotsIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ScreenshotsIndex) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ScreenshotsIndex) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ScreenshotsIndex) SetCount(v int64)`

SetCount sets Count field to given value.


### GetScreenshots

`func (o *ScreenshotsIndex) GetScreenshots() []ScreenshotEntry`

GetScreenshots returns the Screenshots field if non-nil, zero value otherwise.

### GetScreenshotsOk

`func (o *ScreenshotsIndex) GetScreenshotsOk() (*[]ScreenshotEntry, bool)`

GetScreenshotsOk returns a tuple with the Screenshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshots

`func (o *ScreenshotsIndex) SetScreenshots(v []ScreenshotEntry)`

SetScreenshots sets Screenshots field to given value.


### SetScreenshotsNil

`func (o *ScreenshotsIndex) SetScreenshotsNil(b bool)`

 SetScreenshotsNil sets the value for Screenshots to be an explicit nil

### UnsetScreenshots
`func (o *ScreenshotsIndex) UnsetScreenshots()`

UnsetScreenshots ensures that no value is present for Screenshots, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


