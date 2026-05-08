# DnsQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** |  | 
**Events** | Pointer to [**[]ReportEvent**](ReportEvent.md) |  | [optional] 

## Methods

### NewDnsQuery

`func NewDnsQuery(domain string, ) *DnsQuery`

NewDnsQuery instantiates a new DnsQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsQueryWithDefaults

`func NewDnsQueryWithDefaults() *DnsQuery`

NewDnsQueryWithDefaults instantiates a new DnsQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *DnsQuery) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DnsQuery) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DnsQuery) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetEvents

`func (o *DnsQuery) GetEvents() []ReportEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *DnsQuery) GetEventsOk() (*[]ReportEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *DnsQuery) SetEvents(v []ReportEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *DnsQuery) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *DnsQuery) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *DnsQuery) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


