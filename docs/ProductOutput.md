# ProductOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditsPerMonth** | Pointer to **int64** | Credits awarded per billing month, if applicable. | [optional] 
**Description** | **string** | Human-readable product description. | 
**Features** | **[]string** | Marketing feature list for this product. | 
**Id** | **string** | Product ID. | 
**Name** | **string** | Human-readable product name. | 
**Prices** | [**[]PriceOutput**](PriceOutput.md) | All active recurring prices for this product. | 
**Tier** | Pointer to **string** | User tier associated with this product, if any. | [optional] 

## Methods

### NewProductOutput

`func NewProductOutput(description string, features []string, id string, name string, prices []PriceOutput, ) *ProductOutput`

NewProductOutput instantiates a new ProductOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductOutputWithDefaults

`func NewProductOutputWithDefaults() *ProductOutput`

NewProductOutputWithDefaults instantiates a new ProductOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditsPerMonth

`func (o *ProductOutput) GetCreditsPerMonth() int64`

GetCreditsPerMonth returns the CreditsPerMonth field if non-nil, zero value otherwise.

### GetCreditsPerMonthOk

`func (o *ProductOutput) GetCreditsPerMonthOk() (*int64, bool)`

GetCreditsPerMonthOk returns a tuple with the CreditsPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsPerMonth

`func (o *ProductOutput) SetCreditsPerMonth(v int64)`

SetCreditsPerMonth sets CreditsPerMonth field to given value.

### HasCreditsPerMonth

`func (o *ProductOutput) HasCreditsPerMonth() bool`

HasCreditsPerMonth returns a boolean if a field has been set.

### GetDescription

`func (o *ProductOutput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductOutput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductOutput) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFeatures

`func (o *ProductOutput) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ProductOutput) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ProductOutput) SetFeatures(v []string)`

SetFeatures sets Features field to given value.


### SetFeaturesNil

`func (o *ProductOutput) SetFeaturesNil(b bool)`

 SetFeaturesNil sets the value for Features to be an explicit nil

### UnsetFeatures
`func (o *ProductOutput) UnsetFeatures()`

UnsetFeatures ensures that no value is present for Features, not even an explicit nil
### GetId

`func (o *ProductOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductOutput) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ProductOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductOutput) SetName(v string)`

SetName sets Name field to given value.


### GetPrices

`func (o *ProductOutput) GetPrices() []PriceOutput`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *ProductOutput) GetPricesOk() (*[]PriceOutput, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *ProductOutput) SetPrices(v []PriceOutput)`

SetPrices sets Prices field to given value.


### SetPricesNil

`func (o *ProductOutput) SetPricesNil(b bool)`

 SetPricesNil sets the value for Prices to be an explicit nil

### UnsetPrices
`func (o *ProductOutput) UnsetPrices()`

UnsetPrices ensures that no value is present for Prices, not even an explicit nil
### GetTier

`func (o *ProductOutput) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *ProductOutput) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *ProductOutput) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *ProductOutput) HasTier() bool`

HasTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


