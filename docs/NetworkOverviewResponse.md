# NetworkOverviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dns** | [**[]NetworkOverviewDns**](NetworkOverviewDns.md) |  | 
**Metadata** | [**[]NetworkOverviewMetadata**](NetworkOverviewMetadata.md) |  | 

## Methods

### NewNetworkOverviewResponse

`func NewNetworkOverviewResponse(dns []NetworkOverviewDns, metadata []NetworkOverviewMetadata, ) *NetworkOverviewResponse`

NewNetworkOverviewResponse instantiates a new NetworkOverviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkOverviewResponseWithDefaults

`func NewNetworkOverviewResponseWithDefaults() *NetworkOverviewResponse`

NewNetworkOverviewResponseWithDefaults instantiates a new NetworkOverviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDns

`func (o *NetworkOverviewResponse) GetDns() []NetworkOverviewDns`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *NetworkOverviewResponse) GetDnsOk() (*[]NetworkOverviewDns, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *NetworkOverviewResponse) SetDns(v []NetworkOverviewDns)`

SetDns sets Dns field to given value.


### GetMetadata

`func (o *NetworkOverviewResponse) GetMetadata() []NetworkOverviewMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkOverviewResponse) GetMetadataOk() (*[]NetworkOverviewMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkOverviewResponse) SetMetadata(v []NetworkOverviewMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


