# GenerateFunctionDataTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queued** | **bool** | [DEPRECATED] This value has been replaced with the &#x60;data_types_list&#x60; field | 
**Reference** | **string** | [DEPRECATED] This value has been replaced with the &#x60;data_types_list&#x60; field | 
**DataTypesList** | [**GenerationStatusList**](GenerationStatusList.md) | List of function data types information that are either already generated, or now queued for generation | 

## Methods

### NewGenerateFunctionDataTypes

`func NewGenerateFunctionDataTypes(queued bool, reference string, dataTypesList GenerationStatusList, ) *GenerateFunctionDataTypes`

NewGenerateFunctionDataTypes instantiates a new GenerateFunctionDataTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateFunctionDataTypesWithDefaults

`func NewGenerateFunctionDataTypesWithDefaults() *GenerateFunctionDataTypes`

NewGenerateFunctionDataTypesWithDefaults instantiates a new GenerateFunctionDataTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueued

`func (o *GenerateFunctionDataTypes) GetQueued() bool`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *GenerateFunctionDataTypes) GetQueuedOk() (*bool, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *GenerateFunctionDataTypes) SetQueued(v bool)`

SetQueued sets Queued field to given value.


### GetReference

`func (o *GenerateFunctionDataTypes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GenerateFunctionDataTypes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GenerateFunctionDataTypes) SetReference(v string)`

SetReference sets Reference field to given value.


### GetDataTypesList

`func (o *GenerateFunctionDataTypes) GetDataTypesList() GenerationStatusList`

GetDataTypesList returns the DataTypesList field if non-nil, zero value otherwise.

### GetDataTypesListOk

`func (o *GenerateFunctionDataTypes) GetDataTypesListOk() (*GenerationStatusList, bool)`

GetDataTypesListOk returns a tuple with the DataTypesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypesList

`func (o *GenerateFunctionDataTypes) SetDataTypesList(v GenerationStatusList)`

SetDataTypesList sets DataTypesList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


