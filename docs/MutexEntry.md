# MutexEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]ReportEvent**](ReportEvent.md) |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewMutexEntry

`func NewMutexEntry(name string, ) *MutexEntry`

NewMutexEntry instantiates a new MutexEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutexEntryWithDefaults

`func NewMutexEntryWithDefaults() *MutexEntry`

NewMutexEntryWithDefaults instantiates a new MutexEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *MutexEntry) GetEvents() []ReportEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *MutexEntry) GetEventsOk() (*[]ReportEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *MutexEntry) SetEvents(v []ReportEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *MutexEntry) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *MutexEntry) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *MutexEntry) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetName

`func (o *MutexEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MutexEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MutexEntry) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


