# FileActivityEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]ReportEvent**](ReportEvent.md) |  | [optional] 
**Path** | **string** |  | 

## Methods

### NewFileActivityEntry

`func NewFileActivityEntry(path string, ) *FileActivityEntry`

NewFileActivityEntry instantiates a new FileActivityEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileActivityEntryWithDefaults

`func NewFileActivityEntryWithDefaults() *FileActivityEntry`

NewFileActivityEntryWithDefaults instantiates a new FileActivityEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *FileActivityEntry) GetEvents() []ReportEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *FileActivityEntry) GetEventsOk() (*[]ReportEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *FileActivityEntry) SetEvents(v []ReportEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *FileActivityEntry) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *FileActivityEntry) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *FileActivityEntry) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetPath

`func (o *FileActivityEntry) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FileActivityEntry) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FileActivityEntry) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


