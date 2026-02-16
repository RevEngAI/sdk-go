# ImportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfImports** | **int32** |  | 
**Imports** | [**[]map[string]map[string]int32**](map[string]map[string]int32.md) |  | 

## Methods

### NewImportModel

`func NewImportModel(numberOfImports int32, imports []map[string]map[string]int32, ) *ImportModel`

NewImportModel instantiates a new ImportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportModelWithDefaults

`func NewImportModelWithDefaults() *ImportModel`

NewImportModelWithDefaults instantiates a new ImportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfImports

`func (o *ImportModel) GetNumberOfImports() int32`

GetNumberOfImports returns the NumberOfImports field if non-nil, zero value otherwise.

### GetNumberOfImportsOk

`func (o *ImportModel) GetNumberOfImportsOk() (*int32, bool)`

GetNumberOfImportsOk returns a tuple with the NumberOfImports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfImports

`func (o *ImportModel) SetNumberOfImports(v int32)`

SetNumberOfImports sets NumberOfImports field to given value.


### GetImports

`func (o *ImportModel) GetImports() []map[string]map[string]int32`

GetImports returns the Imports field if non-nil, zero value otherwise.

### GetImportsOk

`func (o *ImportModel) GetImportsOk() (*[]map[string]map[string]int32, bool)`

GetImportsOk returns a tuple with the Imports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImports

`func (o *ImportModel) SetImports(v []map[string]map[string]int32)`

SetImports sets Imports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


