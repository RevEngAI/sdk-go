# OperationCostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseCredits** | **float64** | Credits the operation is estimated to use regardless of the binary&#39;s size | 
**CreditsPerFunction** | **float64** | Credits the operation is estimated to use per function in the binary | 
**Operation** | **string** | The operation, as requested when creating an analysis, or binary-analysis for the analysis itself | 

## Methods

### NewOperationCostBody

`func NewOperationCostBody(baseCredits float64, creditsPerFunction float64, operation string, ) *OperationCostBody`

NewOperationCostBody instantiates a new OperationCostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationCostBodyWithDefaults

`func NewOperationCostBodyWithDefaults() *OperationCostBody`

NewOperationCostBodyWithDefaults instantiates a new OperationCostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseCredits

`func (o *OperationCostBody) GetBaseCredits() float64`

GetBaseCredits returns the BaseCredits field if non-nil, zero value otherwise.

### GetBaseCreditsOk

`func (o *OperationCostBody) GetBaseCreditsOk() (*float64, bool)`

GetBaseCreditsOk returns a tuple with the BaseCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCredits

`func (o *OperationCostBody) SetBaseCredits(v float64)`

SetBaseCredits sets BaseCredits field to given value.


### GetCreditsPerFunction

`func (o *OperationCostBody) GetCreditsPerFunction() float64`

GetCreditsPerFunction returns the CreditsPerFunction field if non-nil, zero value otherwise.

### GetCreditsPerFunctionOk

`func (o *OperationCostBody) GetCreditsPerFunctionOk() (*float64, bool)`

GetCreditsPerFunctionOk returns a tuple with the CreditsPerFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsPerFunction

`func (o *OperationCostBody) SetCreditsPerFunction(v float64)`

SetCreditsPerFunction sets CreditsPerFunction field to given value.


### GetOperation

`func (o *OperationCostBody) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *OperationCostBody) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *OperationCostBody) SetOperation(v string)`

SetOperation sets Operation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


