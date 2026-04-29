# APIError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**Error** | [**ErrorBody**](ErrorBody.md) |  | 

## Methods

### NewAPIError

`func NewAPIError(error_ ErrorBody, ) *APIError`

NewAPIError instantiates a new APIError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIErrorWithDefaults

`func NewAPIErrorWithDefaults() *APIError`

NewAPIErrorWithDefaults instantiates a new APIError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *APIError) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *APIError) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *APIError) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *APIError) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetError

`func (o *APIError) GetError() ErrorBody`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *APIError) GetErrorOk() (*ErrorBody, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *APIError) SetError(v ErrorBody)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


