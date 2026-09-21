# NetworkingFinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Function&#39;s virtual address, hex-encoded | 
**Categories** | **[]string** | Distinct networking categories evidenced by this function | 
**Confidence** | **string** | High when a direct name match was found, medium when the function only calls into networking APIs | 
**DirectMatches** | Pointer to [**[]NetworkingDirectMatch**](NetworkingDirectMatch.md) | Matches against the function&#39;s own name | [optional] 
**EvidenceCount** | **int64** | Total number of direct matches and network calls | 
**FunctionId** | **int64** | ID of the function the finding was reported in | 
**FunctionName** | **string** | Name of the function the finding was reported in | 
**FunctionSize** | **int64** | Size of the function in bytes | 
**NetworkCalls** | Pointer to [**[]NetworkingCall**](NetworkingCall.md) | Matches against names this function calls | [optional] 
**Remote** | **bool** | Whether this function evidences remote communication rather than only supporting it | 
**Sources** | **[]string** | Distinct networking sources evidenced by this function | 

## Methods

### NewNetworkingFinding

`func NewNetworkingFinding(address string, categories []string, confidence string, evidenceCount int64, functionId int64, functionName string, functionSize int64, remote bool, sources []string, ) *NetworkingFinding`

NewNetworkingFinding instantiates a new NetworkingFinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingFindingWithDefaults

`func NewNetworkingFindingWithDefaults() *NetworkingFinding`

NewNetworkingFindingWithDefaults instantiates a new NetworkingFinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *NetworkingFinding) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NetworkingFinding) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NetworkingFinding) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCategories

`func (o *NetworkingFinding) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *NetworkingFinding) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *NetworkingFinding) SetCategories(v []string)`

SetCategories sets Categories field to given value.


### SetCategoriesNil

`func (o *NetworkingFinding) SetCategoriesNil(b bool)`

 SetCategoriesNil sets the value for Categories to be an explicit nil

### UnsetCategories
`func (o *NetworkingFinding) UnsetCategories()`

UnsetCategories ensures that no value is present for Categories, not even an explicit nil
### GetConfidence

`func (o *NetworkingFinding) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *NetworkingFinding) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *NetworkingFinding) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.


### GetDirectMatches

`func (o *NetworkingFinding) GetDirectMatches() []NetworkingDirectMatch`

GetDirectMatches returns the DirectMatches field if non-nil, zero value otherwise.

### GetDirectMatchesOk

`func (o *NetworkingFinding) GetDirectMatchesOk() (*[]NetworkingDirectMatch, bool)`

GetDirectMatchesOk returns a tuple with the DirectMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectMatches

`func (o *NetworkingFinding) SetDirectMatches(v []NetworkingDirectMatch)`

SetDirectMatches sets DirectMatches field to given value.

### HasDirectMatches

`func (o *NetworkingFinding) HasDirectMatches() bool`

HasDirectMatches returns a boolean if a field has been set.

### SetDirectMatchesNil

`func (o *NetworkingFinding) SetDirectMatchesNil(b bool)`

 SetDirectMatchesNil sets the value for DirectMatches to be an explicit nil

### UnsetDirectMatches
`func (o *NetworkingFinding) UnsetDirectMatches()`

UnsetDirectMatches ensures that no value is present for DirectMatches, not even an explicit nil
### GetEvidenceCount

`func (o *NetworkingFinding) GetEvidenceCount() int64`

GetEvidenceCount returns the EvidenceCount field if non-nil, zero value otherwise.

### GetEvidenceCountOk

`func (o *NetworkingFinding) GetEvidenceCountOk() (*int64, bool)`

GetEvidenceCountOk returns a tuple with the EvidenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceCount

`func (o *NetworkingFinding) SetEvidenceCount(v int64)`

SetEvidenceCount sets EvidenceCount field to given value.


### GetFunctionId

`func (o *NetworkingFinding) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *NetworkingFinding) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *NetworkingFinding) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetFunctionName

`func (o *NetworkingFinding) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *NetworkingFinding) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *NetworkingFinding) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetFunctionSize

`func (o *NetworkingFinding) GetFunctionSize() int64`

GetFunctionSize returns the FunctionSize field if non-nil, zero value otherwise.

### GetFunctionSizeOk

`func (o *NetworkingFinding) GetFunctionSizeOk() (*int64, bool)`

GetFunctionSizeOk returns a tuple with the FunctionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionSize

`func (o *NetworkingFinding) SetFunctionSize(v int64)`

SetFunctionSize sets FunctionSize field to given value.


### GetNetworkCalls

`func (o *NetworkingFinding) GetNetworkCalls() []NetworkingCall`

GetNetworkCalls returns the NetworkCalls field if non-nil, zero value otherwise.

### GetNetworkCallsOk

`func (o *NetworkingFinding) GetNetworkCallsOk() (*[]NetworkingCall, bool)`

GetNetworkCallsOk returns a tuple with the NetworkCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCalls

`func (o *NetworkingFinding) SetNetworkCalls(v []NetworkingCall)`

SetNetworkCalls sets NetworkCalls field to given value.

### HasNetworkCalls

`func (o *NetworkingFinding) HasNetworkCalls() bool`

HasNetworkCalls returns a boolean if a field has been set.

### SetNetworkCallsNil

`func (o *NetworkingFinding) SetNetworkCallsNil(b bool)`

 SetNetworkCallsNil sets the value for NetworkCalls to be an explicit nil

### UnsetNetworkCalls
`func (o *NetworkingFinding) UnsetNetworkCalls()`

UnsetNetworkCalls ensures that no value is present for NetworkCalls, not even an explicit nil
### GetRemote

`func (o *NetworkingFinding) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *NetworkingFinding) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *NetworkingFinding) SetRemote(v bool)`

SetRemote sets Remote field to given value.


### GetSources

`func (o *NetworkingFinding) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *NetworkingFinding) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *NetworkingFinding) SetSources(v []string)`

SetSources sets Sources field to given value.


### SetSourcesNil

`func (o *NetworkingFinding) SetSourcesNil(b bool)`

 SetSourcesNil sets the value for Sources to be an explicit nil

### UnsetSources
`func (o *NetworkingFinding) UnsetSources()`

UnsetSources ensures that no value is present for Sources, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


