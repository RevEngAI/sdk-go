# DnsQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CnameChain** | Pointer to **[]string** |  | [optional] 
**Domain** | **string** |  | 
**Events** | Pointer to [**[]ReportEvent**](ReportEvent.md) |  | [optional] 
**MinTtl** | Pointer to **int64** |  | [optional] 
**ResolvedIps** | Pointer to **[]string** |  | [optional] 

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

### GetCnameChain

`func (o *DnsQuery) GetCnameChain() []string`

GetCnameChain returns the CnameChain field if non-nil, zero value otherwise.

### GetCnameChainOk

`func (o *DnsQuery) GetCnameChainOk() (*[]string, bool)`

GetCnameChainOk returns a tuple with the CnameChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnameChain

`func (o *DnsQuery) SetCnameChain(v []string)`

SetCnameChain sets CnameChain field to given value.

### HasCnameChain

`func (o *DnsQuery) HasCnameChain() bool`

HasCnameChain returns a boolean if a field has been set.

### SetCnameChainNil

`func (o *DnsQuery) SetCnameChainNil(b bool)`

 SetCnameChainNil sets the value for CnameChain to be an explicit nil

### UnsetCnameChain
`func (o *DnsQuery) UnsetCnameChain()`

UnsetCnameChain ensures that no value is present for CnameChain, not even an explicit nil
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
### GetMinTtl

`func (o *DnsQuery) GetMinTtl() int64`

GetMinTtl returns the MinTtl field if non-nil, zero value otherwise.

### GetMinTtlOk

`func (o *DnsQuery) GetMinTtlOk() (*int64, bool)`

GetMinTtlOk returns a tuple with the MinTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTtl

`func (o *DnsQuery) SetMinTtl(v int64)`

SetMinTtl sets MinTtl field to given value.

### HasMinTtl

`func (o *DnsQuery) HasMinTtl() bool`

HasMinTtl returns a boolean if a field has been set.

### GetResolvedIps

`func (o *DnsQuery) GetResolvedIps() []string`

GetResolvedIps returns the ResolvedIps field if non-nil, zero value otherwise.

### GetResolvedIpsOk

`func (o *DnsQuery) GetResolvedIpsOk() (*[]string, bool)`

GetResolvedIpsOk returns a tuple with the ResolvedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedIps

`func (o *DnsQuery) SetResolvedIps(v []string)`

SetResolvedIps sets ResolvedIps field to given value.

### HasResolvedIps

`func (o *DnsQuery) HasResolvedIps() bool`

HasResolvedIps returns a boolean if a field has been set.

### SetResolvedIpsNil

`func (o *DnsQuery) SetResolvedIpsNil(b bool)`

 SetResolvedIpsNil sets the value for ResolvedIps to be an explicit nil

### UnsetResolvedIps
`func (o *DnsQuery) UnsetResolvedIps()`

UnsetResolvedIps ensures that no value is present for ResolvedIps, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


