# AnalysisOperationBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Resource name of the operation, polled at /v3/{name} | 
**Operation** | **string** | Operation slug, the collection the operation is polled under | 
**Status** | **string** | Status of the operation&#39;s latest run | 

## Methods

### NewAnalysisOperationBody

`func NewAnalysisOperationBody(name string, operation string, status string, ) *AnalysisOperationBody`

NewAnalysisOperationBody instantiates a new AnalysisOperationBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisOperationBodyWithDefaults

`func NewAnalysisOperationBodyWithDefaults() *AnalysisOperationBody`

NewAnalysisOperationBodyWithDefaults instantiates a new AnalysisOperationBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AnalysisOperationBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AnalysisOperationBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AnalysisOperationBody) SetName(v string)`

SetName sets Name field to given value.


### GetOperation

`func (o *AnalysisOperationBody) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *AnalysisOperationBody) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *AnalysisOperationBody) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetStatus

`func (o *AnalysisOperationBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AnalysisOperationBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AnalysisOperationBody) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


