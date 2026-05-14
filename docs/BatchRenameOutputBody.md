# BatchRenameOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**RenamedCount** | **int64** | Number of functions renamed | 

## Methods

### NewBatchRenameOutputBody

`func NewBatchRenameOutputBody(renamedCount int64, ) *BatchRenameOutputBody`

NewBatchRenameOutputBody instantiates a new BatchRenameOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchRenameOutputBodyWithDefaults

`func NewBatchRenameOutputBodyWithDefaults() *BatchRenameOutputBody`

NewBatchRenameOutputBodyWithDefaults instantiates a new BatchRenameOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *BatchRenameOutputBody) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *BatchRenameOutputBody) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *BatchRenameOutputBody) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *BatchRenameOutputBody) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetRenamedCount

`func (o *BatchRenameOutputBody) GetRenamedCount() int64`

GetRenamedCount returns the RenamedCount field if non-nil, zero value otherwise.

### GetRenamedCountOk

`func (o *BatchRenameOutputBody) GetRenamedCountOk() (*int64, bool)`

GetRenamedCountOk returns a tuple with the RenamedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenamedCount

`func (o *BatchRenameOutputBody) SetRenamedCount(v int64)`

SetRenamedCount sets RenamedCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


