# DecompilationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**Decompilation** | Pointer to **string** | Source code with placeholders replaced | [optional] 
**Status** | **string** | Task status | 

## Methods

### NewDecompilationData

`func NewDecompilationData(status string, ) *DecompilationData`

NewDecompilationData instantiates a new DecompilationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecompilationDataWithDefaults

`func NewDecompilationDataWithDefaults() *DecompilationData`

NewDecompilationDataWithDefaults instantiates a new DecompilationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *DecompilationData) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *DecompilationData) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *DecompilationData) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *DecompilationData) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetDecompilation

`func (o *DecompilationData) GetDecompilation() string`

GetDecompilation returns the Decompilation field if non-nil, zero value otherwise.

### GetDecompilationOk

`func (o *DecompilationData) GetDecompilationOk() (*string, bool)`

GetDecompilationOk returns a tuple with the Decompilation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecompilation

`func (o *DecompilationData) SetDecompilation(v string)`

SetDecompilation sets Decompilation field to given value.

### HasDecompilation

`func (o *DecompilationData) HasDecompilation() bool`

HasDecompilation returns a boolean if a field has been set.

### GetStatus

`func (o *DecompilationData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DecompilationData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DecompilationData) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


