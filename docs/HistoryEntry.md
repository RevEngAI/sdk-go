# HistoryEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeMadeBy** | **string** | Username of the user who made the change | 
**CreatedAt** | **time.Time** | When this name change was recorded | 
**FunctionName** | **string** | Function name at this point in history | 
**HistoryId** | **int64** | History record ID | 
**IsDebug** | **bool** | Whether the function had debug info | 
**MangledName** | Pointer to **string** | Mangled function name | [optional] 
**SourceType** | **string** | Source of the rename (USER, SYSTEM, AI_UNSTRIP, etc.) | 

## Methods

### NewHistoryEntry

`func NewHistoryEntry(changeMadeBy string, createdAt time.Time, functionName string, historyId int64, isDebug bool, sourceType string, ) *HistoryEntry`

NewHistoryEntry instantiates a new HistoryEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryEntryWithDefaults

`func NewHistoryEntryWithDefaults() *HistoryEntry`

NewHistoryEntryWithDefaults instantiates a new HistoryEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeMadeBy

`func (o *HistoryEntry) GetChangeMadeBy() string`

GetChangeMadeBy returns the ChangeMadeBy field if non-nil, zero value otherwise.

### GetChangeMadeByOk

`func (o *HistoryEntry) GetChangeMadeByOk() (*string, bool)`

GetChangeMadeByOk returns a tuple with the ChangeMadeBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeMadeBy

`func (o *HistoryEntry) SetChangeMadeBy(v string)`

SetChangeMadeBy sets ChangeMadeBy field to given value.


### GetCreatedAt

`func (o *HistoryEntry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HistoryEntry) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HistoryEntry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetFunctionName

`func (o *HistoryEntry) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *HistoryEntry) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *HistoryEntry) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetHistoryId

`func (o *HistoryEntry) GetHistoryId() int64`

GetHistoryId returns the HistoryId field if non-nil, zero value otherwise.

### GetHistoryIdOk

`func (o *HistoryEntry) GetHistoryIdOk() (*int64, bool)`

GetHistoryIdOk returns a tuple with the HistoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryId

`func (o *HistoryEntry) SetHistoryId(v int64)`

SetHistoryId sets HistoryId field to given value.


### GetIsDebug

`func (o *HistoryEntry) GetIsDebug() bool`

GetIsDebug returns the IsDebug field if non-nil, zero value otherwise.

### GetIsDebugOk

`func (o *HistoryEntry) GetIsDebugOk() (*bool, bool)`

GetIsDebugOk returns a tuple with the IsDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDebug

`func (o *HistoryEntry) SetIsDebug(v bool)`

SetIsDebug sets IsDebug field to given value.


### GetMangledName

`func (o *HistoryEntry) GetMangledName() string`

GetMangledName returns the MangledName field if non-nil, zero value otherwise.

### GetMangledNameOk

`func (o *HistoryEntry) GetMangledNameOk() (*string, bool)`

GetMangledNameOk returns a tuple with the MangledName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMangledName

`func (o *HistoryEntry) SetMangledName(v string)`

SetMangledName sets MangledName field to given value.

### HasMangledName

`func (o *HistoryEntry) HasMangledName() bool`

HasMangledName returns a boolean if a field has been set.

### GetSourceType

`func (o *HistoryEntry) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *HistoryEntry) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *HistoryEntry) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


