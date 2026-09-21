# CapabilitiesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | [**[]Capability**](Capability.md) | Capabilities found. A capability whose address no longer resolves within the analysis is omitted. | 

## Methods

### NewCapabilitiesResult

`func NewCapabilitiesResult(capabilities []Capability, ) *CapabilitiesResult`

NewCapabilitiesResult instantiates a new CapabilitiesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitiesResultWithDefaults

`func NewCapabilitiesResultWithDefaults() *CapabilitiesResult`

NewCapabilitiesResultWithDefaults instantiates a new CapabilitiesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *CapabilitiesResult) GetCapabilities() []Capability`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *CapabilitiesResult) GetCapabilitiesOk() (*[]Capability, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *CapabilitiesResult) SetCapabilities(v []Capability)`

SetCapabilities sets Capabilities field to given value.


### SetCapabilitiesNil

`func (o *CapabilitiesResult) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *CapabilitiesResult) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


