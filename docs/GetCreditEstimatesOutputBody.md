# GetCreditEstimatesOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operations** | [**[]OperationCostBody**](OperationCostBody.md) | Estimated cost of each operation | 
**SizeBuckets** | [**[]SizeBucketBody**](SizeBucketBody.md) | Size buckets in ascending order of size | 

## Methods

### NewGetCreditEstimatesOutputBody

`func NewGetCreditEstimatesOutputBody(operations []OperationCostBody, sizeBuckets []SizeBucketBody, ) *GetCreditEstimatesOutputBody`

NewGetCreditEstimatesOutputBody instantiates a new GetCreditEstimatesOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCreditEstimatesOutputBodyWithDefaults

`func NewGetCreditEstimatesOutputBodyWithDefaults() *GetCreditEstimatesOutputBody`

NewGetCreditEstimatesOutputBodyWithDefaults instantiates a new GetCreditEstimatesOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperations

`func (o *GetCreditEstimatesOutputBody) GetOperations() []OperationCostBody`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *GetCreditEstimatesOutputBody) GetOperationsOk() (*[]OperationCostBody, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *GetCreditEstimatesOutputBody) SetOperations(v []OperationCostBody)`

SetOperations sets Operations field to given value.


### GetSizeBuckets

`func (o *GetCreditEstimatesOutputBody) GetSizeBuckets() []SizeBucketBody`

GetSizeBuckets returns the SizeBuckets field if non-nil, zero value otherwise.

### GetSizeBucketsOk

`func (o *GetCreditEstimatesOutputBody) GetSizeBucketsOk() (*[]SizeBucketBody, bool)`

GetSizeBucketsOk returns a tuple with the SizeBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBuckets

`func (o *GetCreditEstimatesOutputBody) SetSizeBuckets(v []SizeBucketBody)`

SetSizeBuckets sets SizeBuckets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


