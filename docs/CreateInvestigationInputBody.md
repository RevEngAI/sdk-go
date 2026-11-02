# CreateInvestigationInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseType** | **string** |  | 
**Name** | **string** | Investigation name. Must contain a non-whitespace character. | 
**Scope** | Pointer to **string** | PRIVATE is visible only to you, TEAM to your whole team. Only you can modify it either way. | [optional] [default to "PRIVATE"]
**Severity** | **string** | Initial severity. | 

## Methods

### NewCreateInvestigationInputBody

`func NewCreateInvestigationInputBody(caseType string, name string, severity string, ) *CreateInvestigationInputBody`

NewCreateInvestigationInputBody instantiates a new CreateInvestigationInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInvestigationInputBodyWithDefaults

`func NewCreateInvestigationInputBodyWithDefaults() *CreateInvestigationInputBody`

NewCreateInvestigationInputBodyWithDefaults instantiates a new CreateInvestigationInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaseType

`func (o *CreateInvestigationInputBody) GetCaseType() string`

GetCaseType returns the CaseType field if non-nil, zero value otherwise.

### GetCaseTypeOk

`func (o *CreateInvestigationInputBody) GetCaseTypeOk() (*string, bool)`

GetCaseTypeOk returns a tuple with the CaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseType

`func (o *CreateInvestigationInputBody) SetCaseType(v string)`

SetCaseType sets CaseType field to given value.


### GetName

`func (o *CreateInvestigationInputBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateInvestigationInputBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateInvestigationInputBody) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *CreateInvestigationInputBody) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateInvestigationInputBody) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateInvestigationInputBody) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CreateInvestigationInputBody) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSeverity

`func (o *CreateInvestigationInputBody) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *CreateInvestigationInputBody) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *CreateInvestigationInputBody) SetSeverity(v string)`

SetSeverity sets Severity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


