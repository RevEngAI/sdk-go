# FunctionMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionId** | **int64** | Unique identifier of the function | 
**MatchedFunctions** | [**[]MatchedFunction**](MatchedFunction.md) |  | 
**Confidences** | Pointer to [**[]NameConfidence**](NameConfidence.md) |  | [optional] 

## Methods

### NewFunctionMatch

`func NewFunctionMatch(functionId int64, matchedFunctions []MatchedFunction, ) *FunctionMatch`

NewFunctionMatch instantiates a new FunctionMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionMatchWithDefaults

`func NewFunctionMatchWithDefaults() *FunctionMatch`

NewFunctionMatchWithDefaults instantiates a new FunctionMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionId

`func (o *FunctionMatch) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *FunctionMatch) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *FunctionMatch) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetMatchedFunctions

`func (o *FunctionMatch) GetMatchedFunctions() []MatchedFunction`

GetMatchedFunctions returns the MatchedFunctions field if non-nil, zero value otherwise.

### GetMatchedFunctionsOk

`func (o *FunctionMatch) GetMatchedFunctionsOk() (*[]MatchedFunction, bool)`

GetMatchedFunctionsOk returns a tuple with the MatchedFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedFunctions

`func (o *FunctionMatch) SetMatchedFunctions(v []MatchedFunction)`

SetMatchedFunctions sets MatchedFunctions field to given value.


### GetConfidences

`func (o *FunctionMatch) GetConfidences() []NameConfidence`

GetConfidences returns the Confidences field if non-nil, zero value otherwise.

### GetConfidencesOk

`func (o *FunctionMatch) GetConfidencesOk() (*[]NameConfidence, bool)`

GetConfidencesOk returns a tuple with the Confidences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidences

`func (o *FunctionMatch) SetConfidences(v []NameConfidence)`

SetConfidences sets Confidences field to given value.

### HasConfidences

`func (o *FunctionMatch) HasConfidences() bool`

HasConfidences returns a boolean if a field has been set.

### SetConfidencesNil

`func (o *FunctionMatch) SetConfidencesNil(b bool)`

 SetConfidencesNil sets the value for Confidences to be an explicit nil

### UnsetConfidences
`func (o *FunctionMatch) UnsetConfidences()`

UnsetConfidences ensures that no value is present for Confidences, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


