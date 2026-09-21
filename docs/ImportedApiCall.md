# ImportedApiCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImportedApi** | [**ImportedApi**](ImportedApi.md) |  | 
**Subject** | [**Subject**](Subject.md) |  | 
**Reachability** | Pointer to [**NullableCallChain**](CallChain.md) |  | [optional] 

## Methods

### NewImportedApiCall

`func NewImportedApiCall(importedApi ImportedApi, subject Subject, ) *ImportedApiCall`

NewImportedApiCall instantiates a new ImportedApiCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportedApiCallWithDefaults

`func NewImportedApiCallWithDefaults() *ImportedApiCall`

NewImportedApiCallWithDefaults instantiates a new ImportedApiCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImportedApi

`func (o *ImportedApiCall) GetImportedApi() ImportedApi`

GetImportedApi returns the ImportedApi field if non-nil, zero value otherwise.

### GetImportedApiOk

`func (o *ImportedApiCall) GetImportedApiOk() (*ImportedApi, bool)`

GetImportedApiOk returns a tuple with the ImportedApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedApi

`func (o *ImportedApiCall) SetImportedApi(v ImportedApi)`

SetImportedApi sets ImportedApi field to given value.


### GetSubject

`func (o *ImportedApiCall) GetSubject() Subject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ImportedApiCall) GetSubjectOk() (*Subject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ImportedApiCall) SetSubject(v Subject)`

SetSubject sets Subject field to given value.


### GetReachability

`func (o *ImportedApiCall) GetReachability() CallChain`

GetReachability returns the Reachability field if non-nil, zero value otherwise.

### GetReachabilityOk

`func (o *ImportedApiCall) GetReachabilityOk() (*CallChain, bool)`

GetReachabilityOk returns a tuple with the Reachability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachability

`func (o *ImportedApiCall) SetReachability(v CallChain)`

SetReachability sets Reachability field to given value.

### HasReachability

`func (o *ImportedApiCall) HasReachability() bool`

HasReachability returns a boolean if a field has been set.

### SetReachabilityNil

`func (o *ImportedApiCall) SetReachabilityNil(b bool)`

 SetReachabilityNil sets the value for Reachability to be an explicit nil

### UnsetReachability
`func (o *ImportedApiCall) UnsetReachability()`

UnsetReachability ensures that no value is present for Reachability, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


