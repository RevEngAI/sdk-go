# NetworkingExplainResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cancelled** | **bool** | Whether the run was cancelled | 
**DataReceived** | Pointer to **string** | Description of the data received | [optional] 
**DataSent** | Pointer to **string** | Description of the data sent | [optional] 
**DataUsage** | Pointer to **string** | How the sent or received data is used | [optional] 
**Endpoints** | Pointer to **[]string** | Concrete remote endpoints identified -- IPs, hostnames, or URLs | [optional] 
**FunctionId** | **int64** | ID of the explained function | 
**FunctionName** | Pointer to **string** | Name of the explained function | [optional] 
**FunctionsInvolved** | Pointer to [**[]NetworkingExplainedFunction**](NetworkingExplainedFunction.md) | Other functions involved in the network communication | [optional] 
**Protocol** | Pointer to **string** | Protocol used for the communication | [optional] 
**Purpose** | Pointer to **string** | Purpose of the network communication | [optional] 
**Summary** | Pointer to **string** | Explanation of the network communication the function performs | [optional] 

## Methods

### NewNetworkingExplainResult

`func NewNetworkingExplainResult(cancelled bool, functionId int64, ) *NetworkingExplainResult`

NewNetworkingExplainResult instantiates a new NetworkingExplainResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingExplainResultWithDefaults

`func NewNetworkingExplainResultWithDefaults() *NetworkingExplainResult`

NewNetworkingExplainResultWithDefaults instantiates a new NetworkingExplainResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelled

`func (o *NetworkingExplainResult) GetCancelled() bool`

GetCancelled returns the Cancelled field if non-nil, zero value otherwise.

### GetCancelledOk

`func (o *NetworkingExplainResult) GetCancelledOk() (*bool, bool)`

GetCancelledOk returns a tuple with the Cancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelled

`func (o *NetworkingExplainResult) SetCancelled(v bool)`

SetCancelled sets Cancelled field to given value.


### GetDataReceived

`func (o *NetworkingExplainResult) GetDataReceived() string`

GetDataReceived returns the DataReceived field if non-nil, zero value otherwise.

### GetDataReceivedOk

`func (o *NetworkingExplainResult) GetDataReceivedOk() (*string, bool)`

GetDataReceivedOk returns a tuple with the DataReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReceived

`func (o *NetworkingExplainResult) SetDataReceived(v string)`

SetDataReceived sets DataReceived field to given value.

### HasDataReceived

`func (o *NetworkingExplainResult) HasDataReceived() bool`

HasDataReceived returns a boolean if a field has been set.

### GetDataSent

`func (o *NetworkingExplainResult) GetDataSent() string`

GetDataSent returns the DataSent field if non-nil, zero value otherwise.

### GetDataSentOk

`func (o *NetworkingExplainResult) GetDataSentOk() (*string, bool)`

GetDataSentOk returns a tuple with the DataSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSent

`func (o *NetworkingExplainResult) SetDataSent(v string)`

SetDataSent sets DataSent field to given value.

### HasDataSent

`func (o *NetworkingExplainResult) HasDataSent() bool`

HasDataSent returns a boolean if a field has been set.

### GetDataUsage

`func (o *NetworkingExplainResult) GetDataUsage() string`

GetDataUsage returns the DataUsage field if non-nil, zero value otherwise.

### GetDataUsageOk

`func (o *NetworkingExplainResult) GetDataUsageOk() (*string, bool)`

GetDataUsageOk returns a tuple with the DataUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataUsage

`func (o *NetworkingExplainResult) SetDataUsage(v string)`

SetDataUsage sets DataUsage field to given value.

### HasDataUsage

`func (o *NetworkingExplainResult) HasDataUsage() bool`

HasDataUsage returns a boolean if a field has been set.

### GetEndpoints

`func (o *NetworkingExplainResult) GetEndpoints() []string`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *NetworkingExplainResult) GetEndpointsOk() (*[]string, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *NetworkingExplainResult) SetEndpoints(v []string)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *NetworkingExplainResult) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### SetEndpointsNil

`func (o *NetworkingExplainResult) SetEndpointsNil(b bool)`

 SetEndpointsNil sets the value for Endpoints to be an explicit nil

### UnsetEndpoints
`func (o *NetworkingExplainResult) UnsetEndpoints()`

UnsetEndpoints ensures that no value is present for Endpoints, not even an explicit nil
### GetFunctionId

`func (o *NetworkingExplainResult) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *NetworkingExplainResult) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *NetworkingExplainResult) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetFunctionName

`func (o *NetworkingExplainResult) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *NetworkingExplainResult) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *NetworkingExplainResult) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.

### HasFunctionName

`func (o *NetworkingExplainResult) HasFunctionName() bool`

HasFunctionName returns a boolean if a field has been set.

### GetFunctionsInvolved

`func (o *NetworkingExplainResult) GetFunctionsInvolved() []NetworkingExplainedFunction`

GetFunctionsInvolved returns the FunctionsInvolved field if non-nil, zero value otherwise.

### GetFunctionsInvolvedOk

`func (o *NetworkingExplainResult) GetFunctionsInvolvedOk() (*[]NetworkingExplainedFunction, bool)`

GetFunctionsInvolvedOk returns a tuple with the FunctionsInvolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionsInvolved

`func (o *NetworkingExplainResult) SetFunctionsInvolved(v []NetworkingExplainedFunction)`

SetFunctionsInvolved sets FunctionsInvolved field to given value.

### HasFunctionsInvolved

`func (o *NetworkingExplainResult) HasFunctionsInvolved() bool`

HasFunctionsInvolved returns a boolean if a field has been set.

### SetFunctionsInvolvedNil

`func (o *NetworkingExplainResult) SetFunctionsInvolvedNil(b bool)`

 SetFunctionsInvolvedNil sets the value for FunctionsInvolved to be an explicit nil

### UnsetFunctionsInvolved
`func (o *NetworkingExplainResult) UnsetFunctionsInvolved()`

UnsetFunctionsInvolved ensures that no value is present for FunctionsInvolved, not even an explicit nil
### GetProtocol

`func (o *NetworkingExplainResult) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NetworkingExplainResult) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NetworkingExplainResult) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *NetworkingExplainResult) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPurpose

`func (o *NetworkingExplainResult) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *NetworkingExplainResult) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *NetworkingExplainResult) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *NetworkingExplainResult) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetSummary

`func (o *NetworkingExplainResult) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *NetworkingExplainResult) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *NetworkingExplainResult) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *NetworkingExplainResult) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


