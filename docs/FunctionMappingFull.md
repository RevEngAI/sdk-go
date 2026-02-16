# FunctionMappingFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InverseStringMap** | [**map[string]InverseStringMapItem**](InverseStringMapItem.md) |  | 
**InverseFunctionMap** | [**map[string]InverseFunctionMapItem**](InverseFunctionMapItem.md) |  | 
**UnmatchedFunctions** | [**map[string]InverseValue**](InverseValue.md) |  | 
**UnmatchedCustomTypes** | [**map[string]InverseValue**](InverseValue.md) |  | 
**UnmatchedStrings** | [**map[string]InverseValue**](InverseValue.md) |  | 
**UnmatchedVars** | [**map[string]InverseValue**](InverseValue.md) |  | 
**UnmatchedGoToLabels** | [**map[string]InverseValue**](InverseValue.md) |  | 
**UnmatchedCustomFunctionPointers** | [**map[string]InverseValue**](InverseValue.md) |  | 
**UnmatchedVariadicLists** | [**map[string]InverseValue**](InverseValue.md) |  | 
**UnmatchedEnums** | [**map[string]InverseValue**](InverseValue.md) |  | 
**UnmatchedGlobalVars** | [**map[string]InverseValue**](InverseValue.md) |  | 
**Fields** | [**map[string]map[string]InverseValue**](map.md) |  | 
**UnmatchedExternalVars** | Pointer to [**map[string]InverseValue**](InverseValue.md) | No longer provided. | [optional] [default to {}]

## Methods

### NewFunctionMappingFull

`func NewFunctionMappingFull(inverseStringMap map[string]InverseStringMapItem, inverseFunctionMap map[string]InverseFunctionMapItem, unmatchedFunctions map[string]InverseValue, unmatchedCustomTypes map[string]InverseValue, unmatchedStrings map[string]InverseValue, unmatchedVars map[string]InverseValue, unmatchedGoToLabels map[string]InverseValue, unmatchedCustomFunctionPointers map[string]InverseValue, unmatchedVariadicLists map[string]InverseValue, unmatchedEnums map[string]InverseValue, unmatchedGlobalVars map[string]InverseValue, fields map[string]map[string]InverseValue, ) *FunctionMappingFull`

NewFunctionMappingFull instantiates a new FunctionMappingFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionMappingFullWithDefaults

`func NewFunctionMappingFullWithDefaults() *FunctionMappingFull`

NewFunctionMappingFullWithDefaults instantiates a new FunctionMappingFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInverseStringMap

`func (o *FunctionMappingFull) GetInverseStringMap() map[string]InverseStringMapItem`

GetInverseStringMap returns the InverseStringMap field if non-nil, zero value otherwise.

### GetInverseStringMapOk

`func (o *FunctionMappingFull) GetInverseStringMapOk() (*map[string]InverseStringMapItem, bool)`

GetInverseStringMapOk returns a tuple with the InverseStringMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInverseStringMap

`func (o *FunctionMappingFull) SetInverseStringMap(v map[string]InverseStringMapItem)`

SetInverseStringMap sets InverseStringMap field to given value.


### GetInverseFunctionMap

`func (o *FunctionMappingFull) GetInverseFunctionMap() map[string]InverseFunctionMapItem`

GetInverseFunctionMap returns the InverseFunctionMap field if non-nil, zero value otherwise.

### GetInverseFunctionMapOk

`func (o *FunctionMappingFull) GetInverseFunctionMapOk() (*map[string]InverseFunctionMapItem, bool)`

GetInverseFunctionMapOk returns a tuple with the InverseFunctionMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInverseFunctionMap

`func (o *FunctionMappingFull) SetInverseFunctionMap(v map[string]InverseFunctionMapItem)`

SetInverseFunctionMap sets InverseFunctionMap field to given value.


### GetUnmatchedFunctions

`func (o *FunctionMappingFull) GetUnmatchedFunctions() map[string]InverseValue`

GetUnmatchedFunctions returns the UnmatchedFunctions field if non-nil, zero value otherwise.

### GetUnmatchedFunctionsOk

`func (o *FunctionMappingFull) GetUnmatchedFunctionsOk() (*map[string]InverseValue, bool)`

GetUnmatchedFunctionsOk returns a tuple with the UnmatchedFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedFunctions

`func (o *FunctionMappingFull) SetUnmatchedFunctions(v map[string]InverseValue)`

SetUnmatchedFunctions sets UnmatchedFunctions field to given value.


### GetUnmatchedCustomTypes

`func (o *FunctionMappingFull) GetUnmatchedCustomTypes() map[string]InverseValue`

GetUnmatchedCustomTypes returns the UnmatchedCustomTypes field if non-nil, zero value otherwise.

### GetUnmatchedCustomTypesOk

`func (o *FunctionMappingFull) GetUnmatchedCustomTypesOk() (*map[string]InverseValue, bool)`

GetUnmatchedCustomTypesOk returns a tuple with the UnmatchedCustomTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedCustomTypes

`func (o *FunctionMappingFull) SetUnmatchedCustomTypes(v map[string]InverseValue)`

SetUnmatchedCustomTypes sets UnmatchedCustomTypes field to given value.


### GetUnmatchedStrings

`func (o *FunctionMappingFull) GetUnmatchedStrings() map[string]InverseValue`

GetUnmatchedStrings returns the UnmatchedStrings field if non-nil, zero value otherwise.

### GetUnmatchedStringsOk

`func (o *FunctionMappingFull) GetUnmatchedStringsOk() (*map[string]InverseValue, bool)`

GetUnmatchedStringsOk returns a tuple with the UnmatchedStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedStrings

`func (o *FunctionMappingFull) SetUnmatchedStrings(v map[string]InverseValue)`

SetUnmatchedStrings sets UnmatchedStrings field to given value.


### GetUnmatchedVars

`func (o *FunctionMappingFull) GetUnmatchedVars() map[string]InverseValue`

GetUnmatchedVars returns the UnmatchedVars field if non-nil, zero value otherwise.

### GetUnmatchedVarsOk

`func (o *FunctionMappingFull) GetUnmatchedVarsOk() (*map[string]InverseValue, bool)`

GetUnmatchedVarsOk returns a tuple with the UnmatchedVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedVars

`func (o *FunctionMappingFull) SetUnmatchedVars(v map[string]InverseValue)`

SetUnmatchedVars sets UnmatchedVars field to given value.


### GetUnmatchedGoToLabels

`func (o *FunctionMappingFull) GetUnmatchedGoToLabels() map[string]InverseValue`

GetUnmatchedGoToLabels returns the UnmatchedGoToLabels field if non-nil, zero value otherwise.

### GetUnmatchedGoToLabelsOk

`func (o *FunctionMappingFull) GetUnmatchedGoToLabelsOk() (*map[string]InverseValue, bool)`

GetUnmatchedGoToLabelsOk returns a tuple with the UnmatchedGoToLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedGoToLabels

`func (o *FunctionMappingFull) SetUnmatchedGoToLabels(v map[string]InverseValue)`

SetUnmatchedGoToLabels sets UnmatchedGoToLabels field to given value.


### GetUnmatchedCustomFunctionPointers

`func (o *FunctionMappingFull) GetUnmatchedCustomFunctionPointers() map[string]InverseValue`

GetUnmatchedCustomFunctionPointers returns the UnmatchedCustomFunctionPointers field if non-nil, zero value otherwise.

### GetUnmatchedCustomFunctionPointersOk

`func (o *FunctionMappingFull) GetUnmatchedCustomFunctionPointersOk() (*map[string]InverseValue, bool)`

GetUnmatchedCustomFunctionPointersOk returns a tuple with the UnmatchedCustomFunctionPointers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedCustomFunctionPointers

`func (o *FunctionMappingFull) SetUnmatchedCustomFunctionPointers(v map[string]InverseValue)`

SetUnmatchedCustomFunctionPointers sets UnmatchedCustomFunctionPointers field to given value.


### GetUnmatchedVariadicLists

`func (o *FunctionMappingFull) GetUnmatchedVariadicLists() map[string]InverseValue`

GetUnmatchedVariadicLists returns the UnmatchedVariadicLists field if non-nil, zero value otherwise.

### GetUnmatchedVariadicListsOk

`func (o *FunctionMappingFull) GetUnmatchedVariadicListsOk() (*map[string]InverseValue, bool)`

GetUnmatchedVariadicListsOk returns a tuple with the UnmatchedVariadicLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedVariadicLists

`func (o *FunctionMappingFull) SetUnmatchedVariadicLists(v map[string]InverseValue)`

SetUnmatchedVariadicLists sets UnmatchedVariadicLists field to given value.


### GetUnmatchedEnums

`func (o *FunctionMappingFull) GetUnmatchedEnums() map[string]InverseValue`

GetUnmatchedEnums returns the UnmatchedEnums field if non-nil, zero value otherwise.

### GetUnmatchedEnumsOk

`func (o *FunctionMappingFull) GetUnmatchedEnumsOk() (*map[string]InverseValue, bool)`

GetUnmatchedEnumsOk returns a tuple with the UnmatchedEnums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedEnums

`func (o *FunctionMappingFull) SetUnmatchedEnums(v map[string]InverseValue)`

SetUnmatchedEnums sets UnmatchedEnums field to given value.


### GetUnmatchedGlobalVars

`func (o *FunctionMappingFull) GetUnmatchedGlobalVars() map[string]InverseValue`

GetUnmatchedGlobalVars returns the UnmatchedGlobalVars field if non-nil, zero value otherwise.

### GetUnmatchedGlobalVarsOk

`func (o *FunctionMappingFull) GetUnmatchedGlobalVarsOk() (*map[string]InverseValue, bool)`

GetUnmatchedGlobalVarsOk returns a tuple with the UnmatchedGlobalVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedGlobalVars

`func (o *FunctionMappingFull) SetUnmatchedGlobalVars(v map[string]InverseValue)`

SetUnmatchedGlobalVars sets UnmatchedGlobalVars field to given value.


### GetFields

`func (o *FunctionMappingFull) GetFields() map[string]map[string]InverseValue`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *FunctionMappingFull) GetFieldsOk() (*map[string]map[string]InverseValue, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *FunctionMappingFull) SetFields(v map[string]map[string]InverseValue)`

SetFields sets Fields field to given value.


### GetUnmatchedExternalVars

`func (o *FunctionMappingFull) GetUnmatchedExternalVars() map[string]InverseValue`

GetUnmatchedExternalVars returns the UnmatchedExternalVars field if non-nil, zero value otherwise.

### GetUnmatchedExternalVarsOk

`func (o *FunctionMappingFull) GetUnmatchedExternalVarsOk() (*map[string]InverseValue, bool)`

GetUnmatchedExternalVarsOk returns a tuple with the UnmatchedExternalVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedExternalVars

`func (o *FunctionMappingFull) SetUnmatchedExternalVars(v map[string]InverseValue)`

SetUnmatchedExternalVars sets UnmatchedExternalVars field to given value.

### HasUnmatchedExternalVars

`func (o *FunctionMappingFull) HasUnmatchedExternalVars() bool`

HasUnmatchedExternalVars returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


