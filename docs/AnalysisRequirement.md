# AnalysisRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | **[]string** | CreateAnalysis field paths that must be provided or enabled to satisfy this requirement. | 
**Reason** | **string** | Why this requirement unblocks analysis. | 
**Values** | Pointer to **map[string]string** | Field paths that must carry a specific value, keyed the same way as fields. | [optional] 

## Methods

### NewAnalysisRequirement

`func NewAnalysisRequirement(fields []string, reason string, ) *AnalysisRequirement`

NewAnalysisRequirement instantiates a new AnalysisRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisRequirementWithDefaults

`func NewAnalysisRequirementWithDefaults() *AnalysisRequirement`

NewAnalysisRequirementWithDefaults instantiates a new AnalysisRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *AnalysisRequirement) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AnalysisRequirement) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AnalysisRequirement) SetFields(v []string)`

SetFields sets Fields field to given value.


### SetFieldsNil

`func (o *AnalysisRequirement) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *AnalysisRequirement) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetReason

`func (o *AnalysisRequirement) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AnalysisRequirement) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AnalysisRequirement) SetReason(v string)`

SetReason sets Reason field to given value.


### GetValues

`func (o *AnalysisRequirement) GetValues() map[string]string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *AnalysisRequirement) GetValuesOk() (*map[string]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *AnalysisRequirement) SetValues(v map[string]string)`

SetValues sets Values field to given value.

### HasValues

`func (o *AnalysisRequirement) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


