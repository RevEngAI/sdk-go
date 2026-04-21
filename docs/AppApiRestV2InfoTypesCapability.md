# AppApiRestV2InfoTypesCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionName** | **string** | The name of the function with a capability | 
**FunctionVaddr** | **int64** | The virtual address of the function where the capability comes from | 
**Capabilities** | **[]string** | The list of capabilities associated with the function | 

## Methods

### NewAppApiRestV2InfoTypesCapability

`func NewAppApiRestV2InfoTypesCapability(functionName string, functionVaddr int64, capabilities []string, ) *AppApiRestV2InfoTypesCapability`

NewAppApiRestV2InfoTypesCapability instantiates a new AppApiRestV2InfoTypesCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppApiRestV2InfoTypesCapabilityWithDefaults

`func NewAppApiRestV2InfoTypesCapabilityWithDefaults() *AppApiRestV2InfoTypesCapability`

NewAppApiRestV2InfoTypesCapabilityWithDefaults instantiates a new AppApiRestV2InfoTypesCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionName

`func (o *AppApiRestV2InfoTypesCapability) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *AppApiRestV2InfoTypesCapability) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *AppApiRestV2InfoTypesCapability) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetFunctionVaddr

`func (o *AppApiRestV2InfoTypesCapability) GetFunctionVaddr() int64`

GetFunctionVaddr returns the FunctionVaddr field if non-nil, zero value otherwise.

### GetFunctionVaddrOk

`func (o *AppApiRestV2InfoTypesCapability) GetFunctionVaddrOk() (*int64, bool)`

GetFunctionVaddrOk returns a tuple with the FunctionVaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionVaddr

`func (o *AppApiRestV2InfoTypesCapability) SetFunctionVaddr(v int64)`

SetFunctionVaddr sets FunctionVaddr field to given value.


### GetCapabilities

`func (o *AppApiRestV2InfoTypesCapability) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *AppApiRestV2InfoTypesCapability) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *AppApiRestV2InfoTypesCapability) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


