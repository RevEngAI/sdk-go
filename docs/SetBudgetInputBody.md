# SetBudgetInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditBudget** | **NullableInt64** | The most credits the investigation may spend, up to 1,000,000,000, or null to remove the limit. | 

## Methods

### NewSetBudgetInputBody

`func NewSetBudgetInputBody(creditBudget NullableInt64, ) *SetBudgetInputBody`

NewSetBudgetInputBody instantiates a new SetBudgetInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetBudgetInputBodyWithDefaults

`func NewSetBudgetInputBodyWithDefaults() *SetBudgetInputBody`

NewSetBudgetInputBodyWithDefaults instantiates a new SetBudgetInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditBudget

`func (o *SetBudgetInputBody) GetCreditBudget() int64`

GetCreditBudget returns the CreditBudget field if non-nil, zero value otherwise.

### GetCreditBudgetOk

`func (o *SetBudgetInputBody) GetCreditBudgetOk() (*int64, bool)`

GetCreditBudgetOk returns a tuple with the CreditBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditBudget

`func (o *SetBudgetInputBody) SetCreditBudget(v int64)`

SetCreditBudget sets CreditBudget field to given value.


### SetCreditBudgetNil

`func (o *SetBudgetInputBody) SetCreditBudgetNil(b bool)`

 SetCreditBudgetNil sets the value for CreditBudget to be an explicit nil

### UnsetCreditBudget
`func (o *SetBudgetInputBody) UnsetCreditBudget()`

UnsetCreditBudget ensures that no value is present for CreditBudget, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


