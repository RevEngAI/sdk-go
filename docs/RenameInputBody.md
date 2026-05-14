# RenameInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**NewMangledName** | Pointer to **string** | New mangled function name | [optional] 
**NewName** | **string** | New function name | 

## Methods

### NewRenameInputBody

`func NewRenameInputBody(newName string, ) *RenameInputBody`

NewRenameInputBody instantiates a new RenameInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenameInputBodyWithDefaults

`func NewRenameInputBodyWithDefaults() *RenameInputBody`

NewRenameInputBodyWithDefaults instantiates a new RenameInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *RenameInputBody) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *RenameInputBody) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *RenameInputBody) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *RenameInputBody) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetNewMangledName

`func (o *RenameInputBody) GetNewMangledName() string`

GetNewMangledName returns the NewMangledName field if non-nil, zero value otherwise.

### GetNewMangledNameOk

`func (o *RenameInputBody) GetNewMangledNameOk() (*string, bool)`

GetNewMangledNameOk returns a tuple with the NewMangledName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewMangledName

`func (o *RenameInputBody) SetNewMangledName(v string)`

SetNewMangledName sets NewMangledName field to given value.

### HasNewMangledName

`func (o *RenameInputBody) HasNewMangledName() bool`

HasNewMangledName returns a boolean if a field has been set.

### GetNewName

`func (o *RenameInputBody) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *RenameInputBody) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *RenameInputBody) SetNewName(v string)`

SetNewName sets NewName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


