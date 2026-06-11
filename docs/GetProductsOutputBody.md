# GetProductsOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | [**[]ProductOutput**](ProductOutput.md) | List of available products | 

## Methods

### NewGetProductsOutputBody

`func NewGetProductsOutputBody(products []ProductOutput, ) *GetProductsOutputBody`

NewGetProductsOutputBody instantiates a new GetProductsOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProductsOutputBodyWithDefaults

`func NewGetProductsOutputBodyWithDefaults() *GetProductsOutputBody`

NewGetProductsOutputBodyWithDefaults instantiates a new GetProductsOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *GetProductsOutputBody) GetProducts() []ProductOutput`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *GetProductsOutputBody) GetProductsOk() (*[]ProductOutput, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *GetProductsOutputBody) SetProducts(v []ProductOutput)`

SetProducts sets Products field to given value.


### SetProductsNil

`func (o *GetProductsOutputBody) SetProductsNil(b bool)`

 SetProductsNil sets the value for Products to be an explicit nil

### UnsetProducts
`func (o *GetProductsOutputBody) UnsetProducts()`

UnsetProducts ensures that no value is present for Products, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


