# AgentRunIOCBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**Source** | **string** | Where in the binary the agent found the indicator. Empty when the agent gives no location. | 
**Type** | **string** |  | 
**Value** | **string** |  | 

## Methods

### NewAgentRunIOCBody

`func NewAgentRunIOCBody(description string, source string, type_ string, value string, ) *AgentRunIOCBody`

NewAgentRunIOCBody instantiates a new AgentRunIOCBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentRunIOCBodyWithDefaults

`func NewAgentRunIOCBodyWithDefaults() *AgentRunIOCBody`

NewAgentRunIOCBodyWithDefaults instantiates a new AgentRunIOCBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AgentRunIOCBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentRunIOCBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentRunIOCBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSource

`func (o *AgentRunIOCBody) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AgentRunIOCBody) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AgentRunIOCBody) SetSource(v string)`

SetSource sets Source field to given value.


### GetType

`func (o *AgentRunIOCBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgentRunIOCBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgentRunIOCBody) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *AgentRunIOCBody) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AgentRunIOCBody) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AgentRunIOCBody) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


