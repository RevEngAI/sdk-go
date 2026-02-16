# NetworkOverviewDns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** |  | 
**Type** | **string** |  | 
**Answers** | [**[]NetworkOverviewDnsAnswer**](NetworkOverviewDnsAnswer.md) |  | 

## Methods

### NewNetworkOverviewDns

`func NewNetworkOverviewDns(host string, type_ string, answers []NetworkOverviewDnsAnswer, ) *NetworkOverviewDns`

NewNetworkOverviewDns instantiates a new NetworkOverviewDns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkOverviewDnsWithDefaults

`func NewNetworkOverviewDnsWithDefaults() *NetworkOverviewDns`

NewNetworkOverviewDnsWithDefaults instantiates a new NetworkOverviewDns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *NetworkOverviewDns) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *NetworkOverviewDns) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *NetworkOverviewDns) SetHost(v string)`

SetHost sets Host field to given value.


### GetType

`func (o *NetworkOverviewDns) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkOverviewDns) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkOverviewDns) SetType(v string)`

SetType sets Type field to given value.


### GetAnswers

`func (o *NetworkOverviewDns) GetAnswers() []NetworkOverviewDnsAnswer`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *NetworkOverviewDns) GetAnswersOk() (*[]NetworkOverviewDnsAnswer, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *NetworkOverviewDns) SetAnswers(v []NetworkOverviewDnsAnswer)`

SetAnswers sets Answers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


