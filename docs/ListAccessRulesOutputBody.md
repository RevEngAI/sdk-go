# ListAccessRulesOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | [**[]AccessRuleBody**](AccessRuleBody.md) | Every user and team the analysis is shared with, oldest first | 

## Methods

### NewListAccessRulesOutputBody

`func NewListAccessRulesOutputBody(rules []AccessRuleBody, ) *ListAccessRulesOutputBody`

NewListAccessRulesOutputBody instantiates a new ListAccessRulesOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAccessRulesOutputBodyWithDefaults

`func NewListAccessRulesOutputBodyWithDefaults() *ListAccessRulesOutputBody`

NewListAccessRulesOutputBodyWithDefaults instantiates a new ListAccessRulesOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *ListAccessRulesOutputBody) GetRules() []AccessRuleBody`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ListAccessRulesOutputBody) GetRulesOk() (*[]AccessRuleBody, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ListAccessRulesOutputBody) SetRules(v []AccessRuleBody)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


