# ExportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfExports** | **int32** |  | 
**Exports** | **[]map[string]int32** |  | 

## Methods

### NewExportModel

`func NewExportModel(numberOfExports int32, exports []map[string]int32, ) *ExportModel`

NewExportModel instantiates a new ExportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportModelWithDefaults

`func NewExportModelWithDefaults() *ExportModel`

NewExportModelWithDefaults instantiates a new ExportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfExports

`func (o *ExportModel) GetNumberOfExports() int32`

GetNumberOfExports returns the NumberOfExports field if non-nil, zero value otherwise.

### GetNumberOfExportsOk

`func (o *ExportModel) GetNumberOfExportsOk() (*int32, bool)`

GetNumberOfExportsOk returns a tuple with the NumberOfExports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfExports

`func (o *ExportModel) SetNumberOfExports(v int32)`

SetNumberOfExports sets NumberOfExports field to given value.


### GetExports

`func (o *ExportModel) GetExports() []map[string]int32`

GetExports returns the Exports field if non-nil, zero value otherwise.

### GetExportsOk

`func (o *ExportModel) GetExportsOk() (*[]map[string]int32, bool)`

GetExportsOk returns a tuple with the Exports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExports

`func (o *ExportModel) SetExports(v []map[string]int32)`

SetExports sets Exports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


