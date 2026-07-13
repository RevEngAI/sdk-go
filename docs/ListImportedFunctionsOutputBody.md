# ListImportedFunctionsOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImportedFunctions** | [**[]ImportedFunctionEntry**](ImportedFunctionEntry.md) |  | 
**TotalCount** | **int64** | Total imported functions for the binary, ignoring pagination. | 

## Methods

### NewListImportedFunctionsOutputBody

`func NewListImportedFunctionsOutputBody(importedFunctions []ImportedFunctionEntry, totalCount int64, ) *ListImportedFunctionsOutputBody`

NewListImportedFunctionsOutputBody instantiates a new ListImportedFunctionsOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListImportedFunctionsOutputBodyWithDefaults

`func NewListImportedFunctionsOutputBodyWithDefaults() *ListImportedFunctionsOutputBody`

NewListImportedFunctionsOutputBodyWithDefaults instantiates a new ListImportedFunctionsOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImportedFunctions

`func (o *ListImportedFunctionsOutputBody) GetImportedFunctions() []ImportedFunctionEntry`

GetImportedFunctions returns the ImportedFunctions field if non-nil, zero value otherwise.

### GetImportedFunctionsOk

`func (o *ListImportedFunctionsOutputBody) GetImportedFunctionsOk() (*[]ImportedFunctionEntry, bool)`

GetImportedFunctionsOk returns a tuple with the ImportedFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedFunctions

`func (o *ListImportedFunctionsOutputBody) SetImportedFunctions(v []ImportedFunctionEntry)`

SetImportedFunctions sets ImportedFunctions field to given value.


### SetImportedFunctionsNil

`func (o *ListImportedFunctionsOutputBody) SetImportedFunctionsNil(b bool)`

 SetImportedFunctionsNil sets the value for ImportedFunctions to be an explicit nil

### UnsetImportedFunctions
`func (o *ListImportedFunctionsOutputBody) UnsetImportedFunctions()`

UnsetImportedFunctions ensures that no value is present for ImportedFunctions, not even an explicit nil
### GetTotalCount

`func (o *ListImportedFunctionsOutputBody) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListImportedFunctionsOutputBody) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListImportedFunctionsOutputBody) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


