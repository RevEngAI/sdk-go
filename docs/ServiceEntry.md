# ServiceEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinaryPath** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Events** | Pointer to [**[]ReportEvent**](ReportEvent.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ServiceType** | Pointer to **string** |  | [optional] 
**StartType** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceEntry

`func NewServiceEntry() *ServiceEntry`

NewServiceEntry instantiates a new ServiceEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceEntryWithDefaults

`func NewServiceEntryWithDefaults() *ServiceEntry`

NewServiceEntryWithDefaults instantiates a new ServiceEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaryPath

`func (o *ServiceEntry) GetBinaryPath() string`

GetBinaryPath returns the BinaryPath field if non-nil, zero value otherwise.

### GetBinaryPathOk

`func (o *ServiceEntry) GetBinaryPathOk() (*string, bool)`

GetBinaryPathOk returns a tuple with the BinaryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryPath

`func (o *ServiceEntry) SetBinaryPath(v string)`

SetBinaryPath sets BinaryPath field to given value.

### HasBinaryPath

`func (o *ServiceEntry) HasBinaryPath() bool`

HasBinaryPath returns a boolean if a field has been set.

### GetDisplayName

`func (o *ServiceEntry) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ServiceEntry) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ServiceEntry) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ServiceEntry) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEvents

`func (o *ServiceEntry) GetEvents() []ReportEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ServiceEntry) GetEventsOk() (*[]ReportEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ServiceEntry) SetEvents(v []ReportEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ServiceEntry) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *ServiceEntry) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *ServiceEntry) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetName

`func (o *ServiceEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceType

`func (o *ServiceEntry) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *ServiceEntry) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *ServiceEntry) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *ServiceEntry) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetStartType

`func (o *ServiceEntry) GetStartType() string`

GetStartType returns the StartType field if non-nil, zero value otherwise.

### GetStartTypeOk

`func (o *ServiceEntry) GetStartTypeOk() (*string, bool)`

GetStartTypeOk returns a tuple with the StartType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartType

`func (o *ServiceEntry) SetStartType(v string)`

SetStartType sets StartType field to given value.

### HasStartType

`func (o *ServiceEntry) HasStartType() bool`

HasStartType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


