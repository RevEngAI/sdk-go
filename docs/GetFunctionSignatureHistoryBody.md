# GetFunctionSignatureHistoryBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Versions** | [**[]FunctionSignatureVersion**](FunctionSignatureVersion.md) | Every version of the signature, newest first. The first element is the current value, so the list is never empty; a signature that has never been edited has that one element only. | 

## Methods

### NewGetFunctionSignatureHistoryBody

`func NewGetFunctionSignatureHistoryBody(versions []FunctionSignatureVersion, ) *GetFunctionSignatureHistoryBody`

NewGetFunctionSignatureHistoryBody instantiates a new GetFunctionSignatureHistoryBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFunctionSignatureHistoryBodyWithDefaults

`func NewGetFunctionSignatureHistoryBodyWithDefaults() *GetFunctionSignatureHistoryBody`

NewGetFunctionSignatureHistoryBodyWithDefaults instantiates a new GetFunctionSignatureHistoryBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersions

`func (o *GetFunctionSignatureHistoryBody) GetVersions() []FunctionSignatureVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *GetFunctionSignatureHistoryBody) GetVersionsOk() (*[]FunctionSignatureVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *GetFunctionSignatureHistoryBody) SetVersions(v []FunctionSignatureVersion)`

SetVersions sets Versions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


