# PriceOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | Lowercase ISO 4217 currency code (e.g. \&quot;usd\&quot;, \&quot;gbp\&quot;). | 
**Id** | **string** | Price ID. | 
**Interval** | **string** | Billing interval at which the price recurs. | 
**UnitAmount** | **int64** | Price per billing interval, expressed in the smallest unit of the currency (e.g. cents for USD, pence for GBP). | 

## Methods

### NewPriceOutput

`func NewPriceOutput(currency string, id string, interval string, unitAmount int64, ) *PriceOutput`

NewPriceOutput instantiates a new PriceOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceOutputWithDefaults

`func NewPriceOutputWithDefaults() *PriceOutput`

NewPriceOutputWithDefaults instantiates a new PriceOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *PriceOutput) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PriceOutput) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PriceOutput) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetId

`func (o *PriceOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PriceOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PriceOutput) SetId(v string)`

SetId sets Id field to given value.


### GetInterval

`func (o *PriceOutput) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PriceOutput) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PriceOutput) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetUnitAmount

`func (o *PriceOutput) GetUnitAmount() int64`

GetUnitAmount returns the UnitAmount field if non-nil, zero value otherwise.

### GetUnitAmountOk

`func (o *PriceOutput) GetUnitAmountOk() (*int64, bool)`

GetUnitAmountOk returns a tuple with the UnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmount

`func (o *PriceOutput) SetUnitAmount(v int64)`

SetUnitAmount sets UnitAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


