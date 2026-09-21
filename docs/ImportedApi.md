# ImportedApi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Module** | Pointer to **NullableString** |  | [optional] 
**Api** | **string** |  | 

## Methods

### NewImportedApi

`func NewImportedApi(api string, ) *ImportedApi`

NewImportedApi instantiates a new ImportedApi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportedApiWithDefaults

`func NewImportedApiWithDefaults() *ImportedApi`

NewImportedApiWithDefaults instantiates a new ImportedApi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModule

`func (o *ImportedApi) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ImportedApi) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ImportedApi) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *ImportedApi) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *ImportedApi) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *ImportedApi) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetApi

`func (o *ImportedApi) GetApi() string`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *ImportedApi) GetApiOk() (*string, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *ImportedApi) SetApi(v string)`

SetApi sets Api field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


