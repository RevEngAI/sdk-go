# BatchUpdateDataTypesInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Functions** | [**[]BatchUpdateDataTypesItem**](BatchUpdateDataTypesItem.md) | List of functions to update. All function IDs must belong to the analysis in the URL. | 

## Methods

### NewBatchUpdateDataTypesInputBody

`func NewBatchUpdateDataTypesInputBody(functions []BatchUpdateDataTypesItem, ) *BatchUpdateDataTypesInputBody`

NewBatchUpdateDataTypesInputBody instantiates a new BatchUpdateDataTypesInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchUpdateDataTypesInputBodyWithDefaults

`func NewBatchUpdateDataTypesInputBodyWithDefaults() *BatchUpdateDataTypesInputBody`

NewBatchUpdateDataTypesInputBodyWithDefaults instantiates a new BatchUpdateDataTypesInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctions

`func (o *BatchUpdateDataTypesInputBody) GetFunctions() []BatchUpdateDataTypesItem`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *BatchUpdateDataTypesInputBody) GetFunctionsOk() (*[]BatchUpdateDataTypesItem, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *BatchUpdateDataTypesInputBody) SetFunctions(v []BatchUpdateDataTypesItem)`

SetFunctions sets Functions field to given value.


### SetFunctionsNil

`func (o *BatchUpdateDataTypesInputBody) SetFunctionsNil(b bool)`

 SetFunctionsNil sets the value for Functions to be an explicit nil

### UnsetFunctions
`func (o *BatchUpdateDataTypesInputBody) UnsetFunctions()`

UnsetFunctions ensures that no value is present for Functions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


