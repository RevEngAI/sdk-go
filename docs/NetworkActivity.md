# NetworkActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connections** | Pointer to [**[]Connection**](Connection.md) |  | [optional] 
**DnsQueries** | Pointer to [**[]DnsQuery**](DnsQuery.md) |  | [optional] 
**ExtractedUrls** | Pointer to [**[]ExtractedURL**](ExtractedURL.md) |  | [optional] 
**HttpRequests** | Pointer to [**[]HttpRequest**](HttpRequest.md) |  | [optional] 

## Methods

### NewNetworkActivity

`func NewNetworkActivity() *NetworkActivity`

NewNetworkActivity instantiates a new NetworkActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkActivityWithDefaults

`func NewNetworkActivityWithDefaults() *NetworkActivity`

NewNetworkActivityWithDefaults instantiates a new NetworkActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnections

`func (o *NetworkActivity) GetConnections() []Connection`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *NetworkActivity) GetConnectionsOk() (*[]Connection, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *NetworkActivity) SetConnections(v []Connection)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *NetworkActivity) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### SetConnectionsNil

`func (o *NetworkActivity) SetConnectionsNil(b bool)`

 SetConnectionsNil sets the value for Connections to be an explicit nil

### UnsetConnections
`func (o *NetworkActivity) UnsetConnections()`

UnsetConnections ensures that no value is present for Connections, not even an explicit nil
### GetDnsQueries

`func (o *NetworkActivity) GetDnsQueries() []DnsQuery`

GetDnsQueries returns the DnsQueries field if non-nil, zero value otherwise.

### GetDnsQueriesOk

`func (o *NetworkActivity) GetDnsQueriesOk() (*[]DnsQuery, bool)`

GetDnsQueriesOk returns a tuple with the DnsQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsQueries

`func (o *NetworkActivity) SetDnsQueries(v []DnsQuery)`

SetDnsQueries sets DnsQueries field to given value.

### HasDnsQueries

`func (o *NetworkActivity) HasDnsQueries() bool`

HasDnsQueries returns a boolean if a field has been set.

### SetDnsQueriesNil

`func (o *NetworkActivity) SetDnsQueriesNil(b bool)`

 SetDnsQueriesNil sets the value for DnsQueries to be an explicit nil

### UnsetDnsQueries
`func (o *NetworkActivity) UnsetDnsQueries()`

UnsetDnsQueries ensures that no value is present for DnsQueries, not even an explicit nil
### GetExtractedUrls

`func (o *NetworkActivity) GetExtractedUrls() []ExtractedURL`

GetExtractedUrls returns the ExtractedUrls field if non-nil, zero value otherwise.

### GetExtractedUrlsOk

`func (o *NetworkActivity) GetExtractedUrlsOk() (*[]ExtractedURL, bool)`

GetExtractedUrlsOk returns a tuple with the ExtractedUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractedUrls

`func (o *NetworkActivity) SetExtractedUrls(v []ExtractedURL)`

SetExtractedUrls sets ExtractedUrls field to given value.

### HasExtractedUrls

`func (o *NetworkActivity) HasExtractedUrls() bool`

HasExtractedUrls returns a boolean if a field has been set.

### SetExtractedUrlsNil

`func (o *NetworkActivity) SetExtractedUrlsNil(b bool)`

 SetExtractedUrlsNil sets the value for ExtractedUrls to be an explicit nil

### UnsetExtractedUrls
`func (o *NetworkActivity) UnsetExtractedUrls()`

UnsetExtractedUrls ensures that no value is present for ExtractedUrls, not even an explicit nil
### GetHttpRequests

`func (o *NetworkActivity) GetHttpRequests() []HttpRequest`

GetHttpRequests returns the HttpRequests field if non-nil, zero value otherwise.

### GetHttpRequestsOk

`func (o *NetworkActivity) GetHttpRequestsOk() (*[]HttpRequest, bool)`

GetHttpRequestsOk returns a tuple with the HttpRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpRequests

`func (o *NetworkActivity) SetHttpRequests(v []HttpRequest)`

SetHttpRequests sets HttpRequests field to given value.

### HasHttpRequests

`func (o *NetworkActivity) HasHttpRequests() bool`

HasHttpRequests returns a boolean if a field has been set.

### SetHttpRequestsNil

`func (o *NetworkActivity) SetHttpRequestsNil(b bool)`

 SetHttpRequestsNil sets the value for HttpRequests to be an explicit nil

### UnsetHttpRequests
`func (o *NetworkActivity) UnsetHttpRequests()`

UnsetHttpRequests ensures that no value is present for HttpRequests, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


