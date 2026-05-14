# RenameOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**RenamedCount** | **int64** | Number of functions renamed | 

## Methods

### NewRenameOutputBody

`func NewRenameOutputBody(renamedCount int64, ) *RenameOutputBody`

NewRenameOutputBody instantiates a new RenameOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenameOutputBodyWithDefaults

`func NewRenameOutputBodyWithDefaults() *RenameOutputBody`

NewRenameOutputBodyWithDefaults instantiates a new RenameOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *RenameOutputBody) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *RenameOutputBody) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *RenameOutputBody) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *RenameOutputBody) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetRenamedCount

`func (o *RenameOutputBody) GetRenamedCount() int64`

GetRenamedCount returns the RenamedCount field if non-nil, zero value otherwise.

### GetRenamedCountOk

`func (o *RenameOutputBody) GetRenamedCountOk() (*int64, bool)`

GetRenamedCountOk returns a tuple with the RenamedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenamedCount

`func (o *RenameOutputBody) SetRenamedCount(v int64)`

SetRenamedCount sets RenamedCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


