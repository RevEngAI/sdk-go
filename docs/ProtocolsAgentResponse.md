# ProtocolsAgentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 
**Findings** | Pointer to [**[]Finding**](Finding.md) |  | [optional] 

## Methods

### NewProtocolsAgentResponse

`func NewProtocolsAgentResponse() *ProtocolsAgentResponse`

NewProtocolsAgentResponse instantiates a new ProtocolsAgentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtocolsAgentResponseWithDefaults

`func NewProtocolsAgentResponseWithDefaults() *ProtocolsAgentResponse`

NewProtocolsAgentResponseWithDefaults instantiates a new ProtocolsAgentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ProtocolsAgentResponse) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ProtocolsAgentResponse) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ProtocolsAgentResponse) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ProtocolsAgentResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetFindings

`func (o *ProtocolsAgentResponse) GetFindings() []Finding`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *ProtocolsAgentResponse) GetFindingsOk() (*[]Finding, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *ProtocolsAgentResponse) SetFindings(v []Finding)`

SetFindings sets Findings field to given value.

### HasFindings

`func (o *ProtocolsAgentResponse) HasFindings() bool`

HasFindings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


