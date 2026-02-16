# MatchedFunctionSuggestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionId** | **int64** | Unique identifier of the matched function | 
**FunctionVaddr** | **int64** | Virtual address of the matched function | 
**SuggestedName** | Pointer to **NullableString** |  | [optional] 
**SuggestedDemangledName** | **string** | De-mangled name of the function group that contains the matched functions | 

## Methods

### NewMatchedFunctionSuggestion

`func NewMatchedFunctionSuggestion(functionId int64, functionVaddr int64, suggestedDemangledName string, ) *MatchedFunctionSuggestion`

NewMatchedFunctionSuggestion instantiates a new MatchedFunctionSuggestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchedFunctionSuggestionWithDefaults

`func NewMatchedFunctionSuggestionWithDefaults() *MatchedFunctionSuggestion`

NewMatchedFunctionSuggestionWithDefaults instantiates a new MatchedFunctionSuggestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionId

`func (o *MatchedFunctionSuggestion) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *MatchedFunctionSuggestion) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *MatchedFunctionSuggestion) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetFunctionVaddr

`func (o *MatchedFunctionSuggestion) GetFunctionVaddr() int64`

GetFunctionVaddr returns the FunctionVaddr field if non-nil, zero value otherwise.

### GetFunctionVaddrOk

`func (o *MatchedFunctionSuggestion) GetFunctionVaddrOk() (*int64, bool)`

GetFunctionVaddrOk returns a tuple with the FunctionVaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionVaddr

`func (o *MatchedFunctionSuggestion) SetFunctionVaddr(v int64)`

SetFunctionVaddr sets FunctionVaddr field to given value.


### GetSuggestedName

`func (o *MatchedFunctionSuggestion) GetSuggestedName() string`

GetSuggestedName returns the SuggestedName field if non-nil, zero value otherwise.

### GetSuggestedNameOk

`func (o *MatchedFunctionSuggestion) GetSuggestedNameOk() (*string, bool)`

GetSuggestedNameOk returns a tuple with the SuggestedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedName

`func (o *MatchedFunctionSuggestion) SetSuggestedName(v string)`

SetSuggestedName sets SuggestedName field to given value.

### HasSuggestedName

`func (o *MatchedFunctionSuggestion) HasSuggestedName() bool`

HasSuggestedName returns a boolean if a field has been set.

### SetSuggestedNameNil

`func (o *MatchedFunctionSuggestion) SetSuggestedNameNil(b bool)`

 SetSuggestedNameNil sets the value for SuggestedName to be an explicit nil

### UnsetSuggestedName
`func (o *MatchedFunctionSuggestion) UnsetSuggestedName()`

UnsetSuggestedName ensures that no value is present for SuggestedName, not even an explicit nil
### GetSuggestedDemangledName

`func (o *MatchedFunctionSuggestion) GetSuggestedDemangledName() string`

GetSuggestedDemangledName returns the SuggestedDemangledName field if non-nil, zero value otherwise.

### GetSuggestedDemangledNameOk

`func (o *MatchedFunctionSuggestion) GetSuggestedDemangledNameOk() (*string, bool)`

GetSuggestedDemangledNameOk returns a tuple with the SuggestedDemangledName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedDemangledName

`func (o *MatchedFunctionSuggestion) SetSuggestedDemangledName(v string)`

SetSuggestedDemangledName sets SuggestedDemangledName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


