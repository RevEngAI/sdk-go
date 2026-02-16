# FunctionNameHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HistoryId** | **int32** | The ID of the history record | 
**ChangeMadeBy** | **string** | The user who made the change | 
**FunctionName** | **string** | The name of the function | 
**MangledName** | **string** | The mangled name of the function | 
**IsDebug** | **bool** | Whether the function is debugged | 
**SourceType** | [**FunctionSourceType**](FunctionSourceType.md) | The source type of the function | 
**CreatedAt** | **string** | The timestamp when the function name was created | 

## Methods

### NewFunctionNameHistory

`func NewFunctionNameHistory(historyId int32, changeMadeBy string, functionName string, mangledName string, isDebug bool, sourceType FunctionSourceType, createdAt string, ) *FunctionNameHistory`

NewFunctionNameHistory instantiates a new FunctionNameHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionNameHistoryWithDefaults

`func NewFunctionNameHistoryWithDefaults() *FunctionNameHistory`

NewFunctionNameHistoryWithDefaults instantiates a new FunctionNameHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistoryId

`func (o *FunctionNameHistory) GetHistoryId() int32`

GetHistoryId returns the HistoryId field if non-nil, zero value otherwise.

### GetHistoryIdOk

`func (o *FunctionNameHistory) GetHistoryIdOk() (*int32, bool)`

GetHistoryIdOk returns a tuple with the HistoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryId

`func (o *FunctionNameHistory) SetHistoryId(v int32)`

SetHistoryId sets HistoryId field to given value.


### GetChangeMadeBy

`func (o *FunctionNameHistory) GetChangeMadeBy() string`

GetChangeMadeBy returns the ChangeMadeBy field if non-nil, zero value otherwise.

### GetChangeMadeByOk

`func (o *FunctionNameHistory) GetChangeMadeByOk() (*string, bool)`

GetChangeMadeByOk returns a tuple with the ChangeMadeBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeMadeBy

`func (o *FunctionNameHistory) SetChangeMadeBy(v string)`

SetChangeMadeBy sets ChangeMadeBy field to given value.


### GetFunctionName

`func (o *FunctionNameHistory) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *FunctionNameHistory) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *FunctionNameHistory) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetMangledName

`func (o *FunctionNameHistory) GetMangledName() string`

GetMangledName returns the MangledName field if non-nil, zero value otherwise.

### GetMangledNameOk

`func (o *FunctionNameHistory) GetMangledNameOk() (*string, bool)`

GetMangledNameOk returns a tuple with the MangledName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMangledName

`func (o *FunctionNameHistory) SetMangledName(v string)`

SetMangledName sets MangledName field to given value.


### GetIsDebug

`func (o *FunctionNameHistory) GetIsDebug() bool`

GetIsDebug returns the IsDebug field if non-nil, zero value otherwise.

### GetIsDebugOk

`func (o *FunctionNameHistory) GetIsDebugOk() (*bool, bool)`

GetIsDebugOk returns a tuple with the IsDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDebug

`func (o *FunctionNameHistory) SetIsDebug(v bool)`

SetIsDebug sets IsDebug field to given value.


### GetSourceType

`func (o *FunctionNameHistory) GetSourceType() FunctionSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *FunctionNameHistory) GetSourceTypeOk() (*FunctionSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *FunctionNameHistory) SetSourceType(v FunctionSourceType)`

SetSourceType sets SourceType field to given value.


### GetCreatedAt

`func (o *FunctionNameHistory) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FunctionNameHistory) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FunctionNameHistory) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


