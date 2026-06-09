# PriceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | Lowercase ISO 4217 currency code. | 
**Interval** | **string** | Billing interval at which the price recurs. | 
**UnitAmount** | **int64** | Price per billing interval, in the smallest unit of the currency. | 

## Methods

### NewPriceSummary

`func NewPriceSummary(currency string, interval string, unitAmount int64, ) *PriceSummary`

NewPriceSummary instantiates a new PriceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceSummaryWithDefaults

`func NewPriceSummaryWithDefaults() *PriceSummary`

NewPriceSummaryWithDefaults instantiates a new PriceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *PriceSummary) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PriceSummary) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PriceSummary) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetInterval

`func (o *PriceSummary) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PriceSummary) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PriceSummary) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetUnitAmount

`func (o *PriceSummary) GetUnitAmount() int64`

GetUnitAmount returns the UnitAmount field if non-nil, zero value otherwise.

### GetUnitAmountOk

`func (o *PriceSummary) GetUnitAmountOk() (*int64, bool)`

GetUnitAmountOk returns a tuple with the UnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmount

`func (o *PriceSummary) SetUnitAmount(v int64)`

SetUnitAmount sets UnitAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


