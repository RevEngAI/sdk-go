# TypeSuggestionsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** | Language model that produced the suggestions. | [optional] 
**Status** | **string** | Status of the AI decompilation run that would have produced these suggestions. | 
**Types** | [**[]SuggestedTypeView**](SuggestedTypeView.md) | One entry per suggested type. Empty for a run that produced none, and for a run that predates type suggestion — the two are not distinguished. | 

## Methods

### NewTypeSuggestionsData

`func NewTypeSuggestionsData(status string, types []SuggestedTypeView, ) *TypeSuggestionsData`

NewTypeSuggestionsData instantiates a new TypeSuggestionsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypeSuggestionsDataWithDefaults

`func NewTypeSuggestionsDataWithDefaults() *TypeSuggestionsData`

NewTypeSuggestionsDataWithDefaults instantiates a new TypeSuggestionsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *TypeSuggestionsData) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *TypeSuggestionsData) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *TypeSuggestionsData) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *TypeSuggestionsData) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetStatus

`func (o *TypeSuggestionsData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TypeSuggestionsData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TypeSuggestionsData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTypes

`func (o *TypeSuggestionsData) GetTypes() []SuggestedTypeView`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *TypeSuggestionsData) GetTypesOk() (*[]SuggestedTypeView, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *TypeSuggestionsData) SetTypes(v []SuggestedTypeView)`

SetTypes sets Types field to given value.


### SetTypesNil

`func (o *TypeSuggestionsData) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *TypeSuggestionsData) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


