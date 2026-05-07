# ExtractedURL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]ReportEvent**](ReportEvent.md) |  | [optional] 
**Url** | **string** |  | 

## Methods

### NewExtractedURL

`func NewExtractedURL(url string, ) *ExtractedURL`

NewExtractedURL instantiates a new ExtractedURL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtractedURLWithDefaults

`func NewExtractedURLWithDefaults() *ExtractedURL`

NewExtractedURLWithDefaults instantiates a new ExtractedURL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *ExtractedURL) GetEvents() []ReportEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ExtractedURL) GetEventsOk() (*[]ReportEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ExtractedURL) SetEvents(v []ReportEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ExtractedURL) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *ExtractedURL) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *ExtractedURL) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetUrl

`func (o *ExtractedURL) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ExtractedURL) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ExtractedURL) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


