# CreateCheckoutSessionInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelUrl** | **string** | URL to redirect to on cancel. | 
**PriceId** | **string** | Price ID from /v3/billing/products. | 
**SuccessUrl** | **string** | URL to redirect to on success. | 

## Methods

### NewCreateCheckoutSessionInputBody

`func NewCreateCheckoutSessionInputBody(cancelUrl string, priceId string, successUrl string, ) *CreateCheckoutSessionInputBody`

NewCreateCheckoutSessionInputBody instantiates a new CreateCheckoutSessionInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCheckoutSessionInputBodyWithDefaults

`func NewCreateCheckoutSessionInputBodyWithDefaults() *CreateCheckoutSessionInputBody`

NewCreateCheckoutSessionInputBodyWithDefaults instantiates a new CreateCheckoutSessionInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelUrl

`func (o *CreateCheckoutSessionInputBody) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *CreateCheckoutSessionInputBody) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *CreateCheckoutSessionInputBody) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.


### GetPriceId

`func (o *CreateCheckoutSessionInputBody) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *CreateCheckoutSessionInputBody) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *CreateCheckoutSessionInputBody) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.


### GetSuccessUrl

`func (o *CreateCheckoutSessionInputBody) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *CreateCheckoutSessionInputBody) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *CreateCheckoutSessionInputBody) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


