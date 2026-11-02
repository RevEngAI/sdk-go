# BudgetBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditBudget** | **NullableInt64** | The most credits the investigation may spend. Null when it has no limit. | 
**CreditsSpent** | **float64** | Credits spent on the investigation&#39;s conversations and the agents run from it. | 
**Spent** | **bool** | Whether the budget is spent, so nothing more that consumes credits can be started from the investigation. | 

## Methods

### NewBudgetBody

`func NewBudgetBody(creditBudget NullableInt64, creditsSpent float64, spent bool, ) *BudgetBody`

NewBudgetBody instantiates a new BudgetBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetBodyWithDefaults

`func NewBudgetBodyWithDefaults() *BudgetBody`

NewBudgetBodyWithDefaults instantiates a new BudgetBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditBudget

`func (o *BudgetBody) GetCreditBudget() int64`

GetCreditBudget returns the CreditBudget field if non-nil, zero value otherwise.

### GetCreditBudgetOk

`func (o *BudgetBody) GetCreditBudgetOk() (*int64, bool)`

GetCreditBudgetOk returns a tuple with the CreditBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditBudget

`func (o *BudgetBody) SetCreditBudget(v int64)`

SetCreditBudget sets CreditBudget field to given value.


### SetCreditBudgetNil

`func (o *BudgetBody) SetCreditBudgetNil(b bool)`

 SetCreditBudgetNil sets the value for CreditBudget to be an explicit nil

### UnsetCreditBudget
`func (o *BudgetBody) UnsetCreditBudget()`

UnsetCreditBudget ensures that no value is present for CreditBudget, not even an explicit nil
### GetCreditsSpent

`func (o *BudgetBody) GetCreditsSpent() float64`

GetCreditsSpent returns the CreditsSpent field if non-nil, zero value otherwise.

### GetCreditsSpentOk

`func (o *BudgetBody) GetCreditsSpentOk() (*float64, bool)`

GetCreditsSpentOk returns a tuple with the CreditsSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsSpent

`func (o *BudgetBody) SetCreditsSpent(v float64)`

SetCreditsSpent sets CreditsSpent field to given value.


### GetSpent

`func (o *BudgetBody) GetSpent() bool`

GetSpent returns the Spent field if non-nil, zero value otherwise.

### GetSpentOk

`func (o *BudgetBody) GetSpentOk() (*bool, bool)`

GetSpentOk returns a tuple with the Spent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpent

`func (o *BudgetBody) SetSpent(v bool)`

SetSpent sets Spent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


