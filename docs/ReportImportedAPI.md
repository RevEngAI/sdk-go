# ReportImportedAPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Api** | **string** | Name of the API. | 
**Module** | **NullableString** | Module the API is imported from, or null when unknown. | 

## Methods

### NewReportImportedAPI

`func NewReportImportedAPI(api string, module NullableString, ) *ReportImportedAPI`

NewReportImportedAPI instantiates a new ReportImportedAPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportImportedAPIWithDefaults

`func NewReportImportedAPIWithDefaults() *ReportImportedAPI`

NewReportImportedAPIWithDefaults instantiates a new ReportImportedAPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApi

`func (o *ReportImportedAPI) GetApi() string`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *ReportImportedAPI) GetApiOk() (*string, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *ReportImportedAPI) SetApi(v string)`

SetApi sets Api field to given value.


### GetModule

`func (o *ReportImportedAPI) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ReportImportedAPI) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ReportImportedAPI) SetModule(v string)`

SetModule sets Module field to given value.


### SetModuleNil

`func (o *ReportImportedAPI) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *ReportImportedAPI) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


