# AnalysisCapabilitiesOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | [**[]AnalysisCapabilityBody**](AnalysisCapabilityBody.md) | Capabilities found across the binary, ordered by function address. Empty when the binary has no capability record | 

## Methods

### NewAnalysisCapabilitiesOutputBody

`func NewAnalysisCapabilitiesOutputBody(capabilities []AnalysisCapabilityBody, ) *AnalysisCapabilitiesOutputBody`

NewAnalysisCapabilitiesOutputBody instantiates a new AnalysisCapabilitiesOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisCapabilitiesOutputBodyWithDefaults

`func NewAnalysisCapabilitiesOutputBodyWithDefaults() *AnalysisCapabilitiesOutputBody`

NewAnalysisCapabilitiesOutputBodyWithDefaults instantiates a new AnalysisCapabilitiesOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *AnalysisCapabilitiesOutputBody) GetCapabilities() []AnalysisCapabilityBody`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *AnalysisCapabilitiesOutputBody) GetCapabilitiesOk() (*[]AnalysisCapabilityBody, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *AnalysisCapabilitiesOutputBody) SetCapabilities(v []AnalysisCapabilityBody)`

SetCapabilities sets Capabilities field to given value.


### SetCapabilitiesNil

`func (o *AnalysisCapabilitiesOutputBody) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *AnalysisCapabilitiesOutputBody) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


