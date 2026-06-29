# V2FunctionMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionId** | **int64** | Unique identifier of the function | 
**MatchedFunctions** | [**[]V2MatchedFunction**](V2MatchedFunction.md) |  | 
**Confidences** | Pointer to [**[]V2NameConfidence**](V2NameConfidence.md) |  | [optional] 

## Methods

### NewV2FunctionMatch

`func NewV2FunctionMatch(functionId int64, matchedFunctions []V2MatchedFunction, ) *V2FunctionMatch`

NewV2FunctionMatch instantiates a new V2FunctionMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2FunctionMatchWithDefaults

`func NewV2FunctionMatchWithDefaults() *V2FunctionMatch`

NewV2FunctionMatchWithDefaults instantiates a new V2FunctionMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionId

`func (o *V2FunctionMatch) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *V2FunctionMatch) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *V2FunctionMatch) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetMatchedFunctions

`func (o *V2FunctionMatch) GetMatchedFunctions() []V2MatchedFunction`

GetMatchedFunctions returns the MatchedFunctions field if non-nil, zero value otherwise.

### GetMatchedFunctionsOk

`func (o *V2FunctionMatch) GetMatchedFunctionsOk() (*[]V2MatchedFunction, bool)`

GetMatchedFunctionsOk returns a tuple with the MatchedFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedFunctions

`func (o *V2FunctionMatch) SetMatchedFunctions(v []V2MatchedFunction)`

SetMatchedFunctions sets MatchedFunctions field to given value.


### GetConfidences

`func (o *V2FunctionMatch) GetConfidences() []V2NameConfidence`

GetConfidences returns the Confidences field if non-nil, zero value otherwise.

### GetConfidencesOk

`func (o *V2FunctionMatch) GetConfidencesOk() (*[]V2NameConfidence, bool)`

GetConfidencesOk returns a tuple with the Confidences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidences

`func (o *V2FunctionMatch) SetConfidences(v []V2NameConfidence)`

SetConfidences sets Confidences field to given value.

### HasConfidences

`func (o *V2FunctionMatch) HasConfidences() bool`

HasConfidences returns a boolean if a field has been set.

### SetConfidencesNil

`func (o *V2FunctionMatch) SetConfidencesNil(b bool)`

 SetConfidencesNil sets the value for Confidences to be an explicit nil

### UnsetConfidences
`func (o *V2FunctionMatch) UnsetConfidences()`

UnsetConfidences ensures that no value is present for Confidences, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


