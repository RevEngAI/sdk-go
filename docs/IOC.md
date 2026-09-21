# IOC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | What the indicator means | 
**FunctionId** | **NullableInt64** | ID of the function it was found in. Null when the source does not resolve to one. | 
**FunctionName** | **NullableString** | Name of the function it was found in. Null when the source does not resolve to one. | 
**Source** | **NullableString** | Where in the binary it was found, usually a hex address. Null when the agent did not report one. | 
**Type** | **string** | Indicator type | 
**Value** | **string** | The indicator itself | 

## Methods

### NewIOC

`func NewIOC(description string, functionId NullableInt64, functionName NullableString, source NullableString, type_ string, value string, ) *IOC`

NewIOC instantiates a new IOC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIOCWithDefaults

`func NewIOCWithDefaults() *IOC`

NewIOCWithDefaults instantiates a new IOC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IOC) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IOC) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IOC) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFunctionId

`func (o *IOC) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *IOC) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *IOC) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### SetFunctionIdNil

`func (o *IOC) SetFunctionIdNil(b bool)`

 SetFunctionIdNil sets the value for FunctionId to be an explicit nil

### UnsetFunctionId
`func (o *IOC) UnsetFunctionId()`

UnsetFunctionId ensures that no value is present for FunctionId, not even an explicit nil
### GetFunctionName

`func (o *IOC) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *IOC) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *IOC) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### SetFunctionNameNil

`func (o *IOC) SetFunctionNameNil(b bool)`

 SetFunctionNameNil sets the value for FunctionName to be an explicit nil

### UnsetFunctionName
`func (o *IOC) UnsetFunctionName()`

UnsetFunctionName ensures that no value is present for FunctionName, not even an explicit nil
### GetSource

`func (o *IOC) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *IOC) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *IOC) SetSource(v string)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *IOC) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *IOC) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetType

`func (o *IOC) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IOC) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IOC) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *IOC) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IOC) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IOC) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


