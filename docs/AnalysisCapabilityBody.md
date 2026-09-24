# AnalysisCapabilityBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | **[]string** | Capabilities attributed to the function | 
**FunctionName** | **string** | Name of the function the capability was found in | 
**FunctionVaddr** | **int64** | Virtual address of that function | 

## Methods

### NewAnalysisCapabilityBody

`func NewAnalysisCapabilityBody(capabilities []string, functionName string, functionVaddr int64, ) *AnalysisCapabilityBody`

NewAnalysisCapabilityBody instantiates a new AnalysisCapabilityBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisCapabilityBodyWithDefaults

`func NewAnalysisCapabilityBodyWithDefaults() *AnalysisCapabilityBody`

NewAnalysisCapabilityBodyWithDefaults instantiates a new AnalysisCapabilityBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *AnalysisCapabilityBody) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *AnalysisCapabilityBody) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *AnalysisCapabilityBody) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.


### SetCapabilitiesNil

`func (o *AnalysisCapabilityBody) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *AnalysisCapabilityBody) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetFunctionName

`func (o *AnalysisCapabilityBody) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *AnalysisCapabilityBody) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *AnalysisCapabilityBody) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetFunctionVaddr

`func (o *AnalysisCapabilityBody) GetFunctionVaddr() int64`

GetFunctionVaddr returns the FunctionVaddr field if non-nil, zero value otherwise.

### GetFunctionVaddrOk

`func (o *AnalysisCapabilityBody) GetFunctionVaddrOk() (*int64, bool)`

GetFunctionVaddrOk returns a tuple with the FunctionVaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionVaddr

`func (o *AnalysisCapabilityBody) SetFunctionVaddr(v int64)`

SetFunctionVaddr sets FunctionVaddr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


