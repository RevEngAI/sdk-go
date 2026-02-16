# FunctionSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionId** | **int64** | The function ID | 
**FunctionName** | **string** | The name of the function | 
**BinaryName** | **string** | The name of the binary the function belongs to | 
**CreatedAt** | **time.Time** | The creation date of the function | 
**ModelId** | **int32** | The model ID used to analyze the binary the function belongs to | 
**ModelName** | **string** | The name of the model used to analyze the binary the function belongs to | 
**OwnedBy** | **string** | The owner of the binary the function belongs to | 

## Methods

### NewFunctionSearchResult

`func NewFunctionSearchResult(functionId int64, functionName string, binaryName string, createdAt time.Time, modelId int32, modelName string, ownedBy string, ) *FunctionSearchResult`

NewFunctionSearchResult instantiates a new FunctionSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionSearchResultWithDefaults

`func NewFunctionSearchResultWithDefaults() *FunctionSearchResult`

NewFunctionSearchResultWithDefaults instantiates a new FunctionSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionId

`func (o *FunctionSearchResult) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *FunctionSearchResult) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *FunctionSearchResult) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetFunctionName

`func (o *FunctionSearchResult) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *FunctionSearchResult) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *FunctionSearchResult) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetBinaryName

`func (o *FunctionSearchResult) GetBinaryName() string`

GetBinaryName returns the BinaryName field if non-nil, zero value otherwise.

### GetBinaryNameOk

`func (o *FunctionSearchResult) GetBinaryNameOk() (*string, bool)`

GetBinaryNameOk returns a tuple with the BinaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryName

`func (o *FunctionSearchResult) SetBinaryName(v string)`

SetBinaryName sets BinaryName field to given value.


### GetCreatedAt

`func (o *FunctionSearchResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FunctionSearchResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FunctionSearchResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModelId

`func (o *FunctionSearchResult) GetModelId() int32`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *FunctionSearchResult) GetModelIdOk() (*int32, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *FunctionSearchResult) SetModelId(v int32)`

SetModelId sets ModelId field to given value.


### GetModelName

`func (o *FunctionSearchResult) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *FunctionSearchResult) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *FunctionSearchResult) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetOwnedBy

`func (o *FunctionSearchResult) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *FunctionSearchResult) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *FunctionSearchResult) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


