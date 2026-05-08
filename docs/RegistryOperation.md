# RegistryOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]ReportEvent**](ReportEvent.md) |  | [optional] 
**Key** | **string** |  | 

## Methods

### NewRegistryOperation

`func NewRegistryOperation(key string, ) *RegistryOperation`

NewRegistryOperation instantiates a new RegistryOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryOperationWithDefaults

`func NewRegistryOperationWithDefaults() *RegistryOperation`

NewRegistryOperationWithDefaults instantiates a new RegistryOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *RegistryOperation) GetEvents() []ReportEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *RegistryOperation) GetEventsOk() (*[]ReportEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *RegistryOperation) SetEvents(v []ReportEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *RegistryOperation) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *RegistryOperation) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *RegistryOperation) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetKey

`func (o *RegistryOperation) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RegistryOperation) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RegistryOperation) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


