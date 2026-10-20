# ReportImportedAPICall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImportedApi** | [**ReportImportedAPI**](ReportImportedAPI.md) | The API called. | 
**Reachability** | Pointer to [**ReportCallChain**](ReportCallChain.md) | A call chain reaching the call, absent when none was found. | [optional] 
**Subject** | [**ReportSubject**](ReportSubject.md) | Where the call is made. | 

## Methods

### NewReportImportedAPICall

`func NewReportImportedAPICall(importedApi ReportImportedAPI, subject ReportSubject, ) *ReportImportedAPICall`

NewReportImportedAPICall instantiates a new ReportImportedAPICall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportImportedAPICallWithDefaults

`func NewReportImportedAPICallWithDefaults() *ReportImportedAPICall`

NewReportImportedAPICallWithDefaults instantiates a new ReportImportedAPICall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImportedApi

`func (o *ReportImportedAPICall) GetImportedApi() ReportImportedAPI`

GetImportedApi returns the ImportedApi field if non-nil, zero value otherwise.

### GetImportedApiOk

`func (o *ReportImportedAPICall) GetImportedApiOk() (*ReportImportedAPI, bool)`

GetImportedApiOk returns a tuple with the ImportedApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedApi

`func (o *ReportImportedAPICall) SetImportedApi(v ReportImportedAPI)`

SetImportedApi sets ImportedApi field to given value.


### GetReachability

`func (o *ReportImportedAPICall) GetReachability() ReportCallChain`

GetReachability returns the Reachability field if non-nil, zero value otherwise.

### GetReachabilityOk

`func (o *ReportImportedAPICall) GetReachabilityOk() (*ReportCallChain, bool)`

GetReachabilityOk returns a tuple with the Reachability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachability

`func (o *ReportImportedAPICall) SetReachability(v ReportCallChain)`

SetReachability sets Reachability field to given value.

### HasReachability

`func (o *ReportImportedAPICall) HasReachability() bool`

HasReachability returns a boolean if a field has been set.

### GetSubject

`func (o *ReportImportedAPICall) GetSubject() ReportSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ReportImportedAPICall) GetSubjectOk() (*ReportSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ReportImportedAPICall) SetSubject(v ReportSubject)`

SetSubject sets Subject field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


