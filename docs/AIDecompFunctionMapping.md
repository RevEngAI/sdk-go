# AIDecompFunctionMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | [**map[string]map[string]ReplacementValue**](map.md) |  | 
**InverseFunctionMap** | [**map[string]AIDecompInverseFunctionMapItem**](AIDecompInverseFunctionMapItem.md) |  | 
**InverseStringMap** | [**map[string]AIDecompInverseStringMapItem**](AIDecompInverseStringMapItem.md) |  | 
**UnmatchedCustomFunctionPointers** | [**map[string]ReplacementValue**](ReplacementValue.md) |  | 
**UnmatchedCustomTypes** | [**map[string]ReplacementValue**](ReplacementValue.md) |  | 
**UnmatchedEnums** | [**map[string]ReplacementValue**](ReplacementValue.md) |  | 
**UnmatchedExternalVars** | [**map[string]ReplacementValue**](ReplacementValue.md) |  | 
**UnmatchedFunctions** | [**map[string]ReplacementValue**](ReplacementValue.md) |  | 
**UnmatchedGlobalVars** | [**map[string]ReplacementValue**](ReplacementValue.md) |  | 
**UnmatchedGoToLabels** | [**map[string]ReplacementValue**](ReplacementValue.md) |  | 
**UnmatchedStrings** | [**map[string]ReplacementValue**](ReplacementValue.md) |  | 
**UnmatchedVariadicLists** | [**map[string]ReplacementValue**](ReplacementValue.md) |  | 
**UnmatchedVars** | [**map[string]ReplacementValue**](ReplacementValue.md) |  | 
**UserOverrideMappings** | **map[string]string** |  | 

## Methods

### NewAIDecompFunctionMapping

`func NewAIDecompFunctionMapping(fields map[string]map[string]ReplacementValue, inverseFunctionMap map[string]AIDecompInverseFunctionMapItem, inverseStringMap map[string]AIDecompInverseStringMapItem, unmatchedCustomFunctionPointers map[string]ReplacementValue, unmatchedCustomTypes map[string]ReplacementValue, unmatchedEnums map[string]ReplacementValue, unmatchedExternalVars map[string]ReplacementValue, unmatchedFunctions map[string]ReplacementValue, unmatchedGlobalVars map[string]ReplacementValue, unmatchedGoToLabels map[string]ReplacementValue, unmatchedStrings map[string]ReplacementValue, unmatchedVariadicLists map[string]ReplacementValue, unmatchedVars map[string]ReplacementValue, userOverrideMappings map[string]string, ) *AIDecompFunctionMapping`

NewAIDecompFunctionMapping instantiates a new AIDecompFunctionMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIDecompFunctionMappingWithDefaults

`func NewAIDecompFunctionMappingWithDefaults() *AIDecompFunctionMapping`

NewAIDecompFunctionMappingWithDefaults instantiates a new AIDecompFunctionMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *AIDecompFunctionMapping) GetFields() map[string]map[string]ReplacementValue`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AIDecompFunctionMapping) GetFieldsOk() (*map[string]map[string]ReplacementValue, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AIDecompFunctionMapping) SetFields(v map[string]map[string]ReplacementValue)`

SetFields sets Fields field to given value.


### GetInverseFunctionMap

`func (o *AIDecompFunctionMapping) GetInverseFunctionMap() map[string]AIDecompInverseFunctionMapItem`

GetInverseFunctionMap returns the InverseFunctionMap field if non-nil, zero value otherwise.

### GetInverseFunctionMapOk

`func (o *AIDecompFunctionMapping) GetInverseFunctionMapOk() (*map[string]AIDecompInverseFunctionMapItem, bool)`

GetInverseFunctionMapOk returns a tuple with the InverseFunctionMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInverseFunctionMap

`func (o *AIDecompFunctionMapping) SetInverseFunctionMap(v map[string]AIDecompInverseFunctionMapItem)`

SetInverseFunctionMap sets InverseFunctionMap field to given value.


### GetInverseStringMap

`func (o *AIDecompFunctionMapping) GetInverseStringMap() map[string]AIDecompInverseStringMapItem`

GetInverseStringMap returns the InverseStringMap field if non-nil, zero value otherwise.

### GetInverseStringMapOk

`func (o *AIDecompFunctionMapping) GetInverseStringMapOk() (*map[string]AIDecompInverseStringMapItem, bool)`

GetInverseStringMapOk returns a tuple with the InverseStringMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInverseStringMap

`func (o *AIDecompFunctionMapping) SetInverseStringMap(v map[string]AIDecompInverseStringMapItem)`

SetInverseStringMap sets InverseStringMap field to given value.


### GetUnmatchedCustomFunctionPointers

`func (o *AIDecompFunctionMapping) GetUnmatchedCustomFunctionPointers() map[string]ReplacementValue`

GetUnmatchedCustomFunctionPointers returns the UnmatchedCustomFunctionPointers field if non-nil, zero value otherwise.

### GetUnmatchedCustomFunctionPointersOk

`func (o *AIDecompFunctionMapping) GetUnmatchedCustomFunctionPointersOk() (*map[string]ReplacementValue, bool)`

GetUnmatchedCustomFunctionPointersOk returns a tuple with the UnmatchedCustomFunctionPointers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedCustomFunctionPointers

`func (o *AIDecompFunctionMapping) SetUnmatchedCustomFunctionPointers(v map[string]ReplacementValue)`

SetUnmatchedCustomFunctionPointers sets UnmatchedCustomFunctionPointers field to given value.


### GetUnmatchedCustomTypes

`func (o *AIDecompFunctionMapping) GetUnmatchedCustomTypes() map[string]ReplacementValue`

GetUnmatchedCustomTypes returns the UnmatchedCustomTypes field if non-nil, zero value otherwise.

### GetUnmatchedCustomTypesOk

`func (o *AIDecompFunctionMapping) GetUnmatchedCustomTypesOk() (*map[string]ReplacementValue, bool)`

GetUnmatchedCustomTypesOk returns a tuple with the UnmatchedCustomTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedCustomTypes

`func (o *AIDecompFunctionMapping) SetUnmatchedCustomTypes(v map[string]ReplacementValue)`

SetUnmatchedCustomTypes sets UnmatchedCustomTypes field to given value.


### GetUnmatchedEnums

`func (o *AIDecompFunctionMapping) GetUnmatchedEnums() map[string]ReplacementValue`

GetUnmatchedEnums returns the UnmatchedEnums field if non-nil, zero value otherwise.

### GetUnmatchedEnumsOk

`func (o *AIDecompFunctionMapping) GetUnmatchedEnumsOk() (*map[string]ReplacementValue, bool)`

GetUnmatchedEnumsOk returns a tuple with the UnmatchedEnums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedEnums

`func (o *AIDecompFunctionMapping) SetUnmatchedEnums(v map[string]ReplacementValue)`

SetUnmatchedEnums sets UnmatchedEnums field to given value.


### GetUnmatchedExternalVars

`func (o *AIDecompFunctionMapping) GetUnmatchedExternalVars() map[string]ReplacementValue`

GetUnmatchedExternalVars returns the UnmatchedExternalVars field if non-nil, zero value otherwise.

### GetUnmatchedExternalVarsOk

`func (o *AIDecompFunctionMapping) GetUnmatchedExternalVarsOk() (*map[string]ReplacementValue, bool)`

GetUnmatchedExternalVarsOk returns a tuple with the UnmatchedExternalVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedExternalVars

`func (o *AIDecompFunctionMapping) SetUnmatchedExternalVars(v map[string]ReplacementValue)`

SetUnmatchedExternalVars sets UnmatchedExternalVars field to given value.


### GetUnmatchedFunctions

`func (o *AIDecompFunctionMapping) GetUnmatchedFunctions() map[string]ReplacementValue`

GetUnmatchedFunctions returns the UnmatchedFunctions field if non-nil, zero value otherwise.

### GetUnmatchedFunctionsOk

`func (o *AIDecompFunctionMapping) GetUnmatchedFunctionsOk() (*map[string]ReplacementValue, bool)`

GetUnmatchedFunctionsOk returns a tuple with the UnmatchedFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedFunctions

`func (o *AIDecompFunctionMapping) SetUnmatchedFunctions(v map[string]ReplacementValue)`

SetUnmatchedFunctions sets UnmatchedFunctions field to given value.


### GetUnmatchedGlobalVars

`func (o *AIDecompFunctionMapping) GetUnmatchedGlobalVars() map[string]ReplacementValue`

GetUnmatchedGlobalVars returns the UnmatchedGlobalVars field if non-nil, zero value otherwise.

### GetUnmatchedGlobalVarsOk

`func (o *AIDecompFunctionMapping) GetUnmatchedGlobalVarsOk() (*map[string]ReplacementValue, bool)`

GetUnmatchedGlobalVarsOk returns a tuple with the UnmatchedGlobalVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedGlobalVars

`func (o *AIDecompFunctionMapping) SetUnmatchedGlobalVars(v map[string]ReplacementValue)`

SetUnmatchedGlobalVars sets UnmatchedGlobalVars field to given value.


### GetUnmatchedGoToLabels

`func (o *AIDecompFunctionMapping) GetUnmatchedGoToLabels() map[string]ReplacementValue`

GetUnmatchedGoToLabels returns the UnmatchedGoToLabels field if non-nil, zero value otherwise.

### GetUnmatchedGoToLabelsOk

`func (o *AIDecompFunctionMapping) GetUnmatchedGoToLabelsOk() (*map[string]ReplacementValue, bool)`

GetUnmatchedGoToLabelsOk returns a tuple with the UnmatchedGoToLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedGoToLabels

`func (o *AIDecompFunctionMapping) SetUnmatchedGoToLabels(v map[string]ReplacementValue)`

SetUnmatchedGoToLabels sets UnmatchedGoToLabels field to given value.


### GetUnmatchedStrings

`func (o *AIDecompFunctionMapping) GetUnmatchedStrings() map[string]ReplacementValue`

GetUnmatchedStrings returns the UnmatchedStrings field if non-nil, zero value otherwise.

### GetUnmatchedStringsOk

`func (o *AIDecompFunctionMapping) GetUnmatchedStringsOk() (*map[string]ReplacementValue, bool)`

GetUnmatchedStringsOk returns a tuple with the UnmatchedStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedStrings

`func (o *AIDecompFunctionMapping) SetUnmatchedStrings(v map[string]ReplacementValue)`

SetUnmatchedStrings sets UnmatchedStrings field to given value.


### GetUnmatchedVariadicLists

`func (o *AIDecompFunctionMapping) GetUnmatchedVariadicLists() map[string]ReplacementValue`

GetUnmatchedVariadicLists returns the UnmatchedVariadicLists field if non-nil, zero value otherwise.

### GetUnmatchedVariadicListsOk

`func (o *AIDecompFunctionMapping) GetUnmatchedVariadicListsOk() (*map[string]ReplacementValue, bool)`

GetUnmatchedVariadicListsOk returns a tuple with the UnmatchedVariadicLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedVariadicLists

`func (o *AIDecompFunctionMapping) SetUnmatchedVariadicLists(v map[string]ReplacementValue)`

SetUnmatchedVariadicLists sets UnmatchedVariadicLists field to given value.


### GetUnmatchedVars

`func (o *AIDecompFunctionMapping) GetUnmatchedVars() map[string]ReplacementValue`

GetUnmatchedVars returns the UnmatchedVars field if non-nil, zero value otherwise.

### GetUnmatchedVarsOk

`func (o *AIDecompFunctionMapping) GetUnmatchedVarsOk() (*map[string]ReplacementValue, bool)`

GetUnmatchedVarsOk returns a tuple with the UnmatchedVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedVars

`func (o *AIDecompFunctionMapping) SetUnmatchedVars(v map[string]ReplacementValue)`

SetUnmatchedVars sets UnmatchedVars field to given value.


### GetUserOverrideMappings

`func (o *AIDecompFunctionMapping) GetUserOverrideMappings() map[string]string`

GetUserOverrideMappings returns the UserOverrideMappings field if non-nil, zero value otherwise.

### GetUserOverrideMappingsOk

`func (o *AIDecompFunctionMapping) GetUserOverrideMappingsOk() (*map[string]string, bool)`

GetUserOverrideMappingsOk returns a tuple with the UserOverrideMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserOverrideMappings

`func (o *AIDecompFunctionMapping) SetUserOverrideMappings(v map[string]string)`

SetUserOverrideMappings sets UserOverrideMappings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


