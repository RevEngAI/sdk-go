# GetSubscriptionOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndsAt** | Pointer to **time.Time** | Date access ends (CANCELING only). | [optional] 
**Price** | Pointer to [**PriceSummary**](PriceSummary.md) | Current price (ACTIVE / CANCELING / PAYMENT_ISSUE only). | [optional] 
**Product** | Pointer to [**ProductSummary**](ProductSummary.md) | Subscribed product (ACTIVE / CANCELING / PAYMENT_ISSUE only). | [optional] 
**RenewsAt** | Pointer to **time.Time** | Next billing date (ACTIVE only). | [optional] 
**Status** | **string** | Subscription state. | 
**Tier** | **string** | User&#39;s effective tier. | 

## Methods

### NewGetSubscriptionOutputBody

`func NewGetSubscriptionOutputBody(status string, tier string, ) *GetSubscriptionOutputBody`

NewGetSubscriptionOutputBody instantiates a new GetSubscriptionOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSubscriptionOutputBodyWithDefaults

`func NewGetSubscriptionOutputBodyWithDefaults() *GetSubscriptionOutputBody`

NewGetSubscriptionOutputBodyWithDefaults instantiates a new GetSubscriptionOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndsAt

`func (o *GetSubscriptionOutputBody) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *GetSubscriptionOutputBody) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *GetSubscriptionOutputBody) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *GetSubscriptionOutputBody) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.

### GetPrice

`func (o *GetSubscriptionOutputBody) GetPrice() PriceSummary`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetSubscriptionOutputBody) GetPriceOk() (*PriceSummary, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetSubscriptionOutputBody) SetPrice(v PriceSummary)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetSubscriptionOutputBody) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetProduct

`func (o *GetSubscriptionOutputBody) GetProduct() ProductSummary`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *GetSubscriptionOutputBody) GetProductOk() (*ProductSummary, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *GetSubscriptionOutputBody) SetProduct(v ProductSummary)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *GetSubscriptionOutputBody) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetRenewsAt

`func (o *GetSubscriptionOutputBody) GetRenewsAt() time.Time`

GetRenewsAt returns the RenewsAt field if non-nil, zero value otherwise.

### GetRenewsAtOk

`func (o *GetSubscriptionOutputBody) GetRenewsAtOk() (*time.Time, bool)`

GetRenewsAtOk returns a tuple with the RenewsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewsAt

`func (o *GetSubscriptionOutputBody) SetRenewsAt(v time.Time)`

SetRenewsAt sets RenewsAt field to given value.

### HasRenewsAt

`func (o *GetSubscriptionOutputBody) HasRenewsAt() bool`

HasRenewsAt returns a boolean if a field has been set.

### GetStatus

`func (o *GetSubscriptionOutputBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSubscriptionOutputBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSubscriptionOutputBody) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTier

`func (o *GetSubscriptionOutputBody) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *GetSubscriptionOutputBody) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *GetSubscriptionOutputBody) SetTier(v string)`

SetTier sets Tier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


