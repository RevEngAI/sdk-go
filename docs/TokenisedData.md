# TokenisedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]ResolvedEntity**](ResolvedEntity.md) | One entry per token in the tokenised source, with the name it resolves to and its hover metadata. | [optional] 
**LineAttribution** | Pointer to **interface{}** |  | [optional] 
**PredictedFunctionName** | Pointer to **string** | Predicted function name from the AI model | [optional] 
**Status** | **string** | Task status | 
**TokenisedDecompilation** | Pointer to **string** | Source code with placeholder tokens | [optional] 

## Methods

### NewTokenisedData

`func NewTokenisedData(status string, ) *TokenisedData`

NewTokenisedData instantiates a new TokenisedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenisedDataWithDefaults

`func NewTokenisedDataWithDefaults() *TokenisedData`

NewTokenisedDataWithDefaults instantiates a new TokenisedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *TokenisedData) GetEntities() []ResolvedEntity`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *TokenisedData) GetEntitiesOk() (*[]ResolvedEntity, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *TokenisedData) SetEntities(v []ResolvedEntity)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *TokenisedData) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### SetEntitiesNil

`func (o *TokenisedData) SetEntitiesNil(b bool)`

 SetEntitiesNil sets the value for Entities to be an explicit nil

### UnsetEntities
`func (o *TokenisedData) UnsetEntities()`

UnsetEntities ensures that no value is present for Entities, not even an explicit nil
### GetLineAttribution

`func (o *TokenisedData) GetLineAttribution() interface{}`

GetLineAttribution returns the LineAttribution field if non-nil, zero value otherwise.

### GetLineAttributionOk

`func (o *TokenisedData) GetLineAttributionOk() (*interface{}, bool)`

GetLineAttributionOk returns a tuple with the LineAttribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineAttribution

`func (o *TokenisedData) SetLineAttribution(v interface{})`

SetLineAttribution sets LineAttribution field to given value.

### HasLineAttribution

`func (o *TokenisedData) HasLineAttribution() bool`

HasLineAttribution returns a boolean if a field has been set.

### SetLineAttributionNil

`func (o *TokenisedData) SetLineAttributionNil(b bool)`

 SetLineAttributionNil sets the value for LineAttribution to be an explicit nil

### UnsetLineAttribution
`func (o *TokenisedData) UnsetLineAttribution()`

UnsetLineAttribution ensures that no value is present for LineAttribution, not even an explicit nil
### GetPredictedFunctionName

`func (o *TokenisedData) GetPredictedFunctionName() string`

GetPredictedFunctionName returns the PredictedFunctionName field if non-nil, zero value otherwise.

### GetPredictedFunctionNameOk

`func (o *TokenisedData) GetPredictedFunctionNameOk() (*string, bool)`

GetPredictedFunctionNameOk returns a tuple with the PredictedFunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedFunctionName

`func (o *TokenisedData) SetPredictedFunctionName(v string)`

SetPredictedFunctionName sets PredictedFunctionName field to given value.

### HasPredictedFunctionName

`func (o *TokenisedData) HasPredictedFunctionName() bool`

HasPredictedFunctionName returns a boolean if a field has been set.

### GetStatus

`func (o *TokenisedData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TokenisedData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TokenisedData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTokenisedDecompilation

`func (o *TokenisedData) GetTokenisedDecompilation() string`

GetTokenisedDecompilation returns the TokenisedDecompilation field if non-nil, zero value otherwise.

### GetTokenisedDecompilationOk

`func (o *TokenisedData) GetTokenisedDecompilationOk() (*string, bool)`

GetTokenisedDecompilationOk returns a tuple with the TokenisedDecompilation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenisedDecompilation

`func (o *TokenisedData) SetTokenisedDecompilation(v string)`

SetTokenisedDecompilation sets TokenisedDecompilation field to given value.

### HasTokenisedDecompilation

`func (o *TokenisedData) HasTokenisedDecompilation() bool`

HasTokenisedDecompilation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


