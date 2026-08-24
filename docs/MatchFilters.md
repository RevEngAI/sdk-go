# MatchFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architectures** | Pointer to **[]string** | Restrict matches to candidates whose binary was detected as one of these architectures. Word size is part of the value, so there is no separate bits filter. Matches all architectures if omitted. | [optional] 
**BinaryIds** | Pointer to **[]int64** | Restrict the candidate pool to these binary IDs. | [optional] 
**CollectionIds** | Pointer to **[]int64** | Restrict the candidate pool to binaries in these collection IDs. | [optional] 
**Debug** | Pointer to **bool** | Restrict matches to candidates with auto/system debug symbols. Multi-platform models only; rejected for single-architecture models. | [optional] 
**DebugTypes** | Pointer to **[]string** | Restrict matches to candidates with these debug source types. Accepted: SYSTEM, USER. | [optional] 
**FunctionIds** | Pointer to **[]int64** | Restrict the candidate pool to these function IDs. | [optional] 
**IncludeUserDebug** | Pointer to **bool** | When debug is set, also match user-named functions (not only auto/system debug). No effect unless debug is true. | [optional] 
**Platforms** | Pointer to **[]string** | Restrict matches to candidates whose binary was detected as one of these platforms. Matches all platforms if omitted; a binary whose detection has not run is never matched by a non-empty filter. | [optional] 
**UserIds** | Pointer to **[]int64** | Restrict the candidate pool to functions owned by these user IDs. | [optional] 

## Methods

### NewMatchFilters

`func NewMatchFilters() *MatchFilters`

NewMatchFilters instantiates a new MatchFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchFiltersWithDefaults

`func NewMatchFiltersWithDefaults() *MatchFilters`

NewMatchFiltersWithDefaults instantiates a new MatchFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitectures

`func (o *MatchFilters) GetArchitectures() []string`

GetArchitectures returns the Architectures field if non-nil, zero value otherwise.

### GetArchitecturesOk

`func (o *MatchFilters) GetArchitecturesOk() (*[]string, bool)`

GetArchitecturesOk returns a tuple with the Architectures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitectures

`func (o *MatchFilters) SetArchitectures(v []string)`

SetArchitectures sets Architectures field to given value.

### HasArchitectures

`func (o *MatchFilters) HasArchitectures() bool`

HasArchitectures returns a boolean if a field has been set.

### SetArchitecturesNil

`func (o *MatchFilters) SetArchitecturesNil(b bool)`

 SetArchitecturesNil sets the value for Architectures to be an explicit nil

### UnsetArchitectures
`func (o *MatchFilters) UnsetArchitectures()`

UnsetArchitectures ensures that no value is present for Architectures, not even an explicit nil
### GetBinaryIds

`func (o *MatchFilters) GetBinaryIds() []int64`

GetBinaryIds returns the BinaryIds field if non-nil, zero value otherwise.

### GetBinaryIdsOk

`func (o *MatchFilters) GetBinaryIdsOk() (*[]int64, bool)`

GetBinaryIdsOk returns a tuple with the BinaryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryIds

`func (o *MatchFilters) SetBinaryIds(v []int64)`

SetBinaryIds sets BinaryIds field to given value.

### HasBinaryIds

`func (o *MatchFilters) HasBinaryIds() bool`

HasBinaryIds returns a boolean if a field has been set.

### SetBinaryIdsNil

`func (o *MatchFilters) SetBinaryIdsNil(b bool)`

 SetBinaryIdsNil sets the value for BinaryIds to be an explicit nil

### UnsetBinaryIds
`func (o *MatchFilters) UnsetBinaryIds()`

UnsetBinaryIds ensures that no value is present for BinaryIds, not even an explicit nil
### GetCollectionIds

`func (o *MatchFilters) GetCollectionIds() []int64`

GetCollectionIds returns the CollectionIds field if non-nil, zero value otherwise.

### GetCollectionIdsOk

`func (o *MatchFilters) GetCollectionIdsOk() (*[]int64, bool)`

GetCollectionIdsOk returns a tuple with the CollectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionIds

`func (o *MatchFilters) SetCollectionIds(v []int64)`

SetCollectionIds sets CollectionIds field to given value.

### HasCollectionIds

`func (o *MatchFilters) HasCollectionIds() bool`

HasCollectionIds returns a boolean if a field has been set.

### SetCollectionIdsNil

`func (o *MatchFilters) SetCollectionIdsNil(b bool)`

 SetCollectionIdsNil sets the value for CollectionIds to be an explicit nil

### UnsetCollectionIds
`func (o *MatchFilters) UnsetCollectionIds()`

UnsetCollectionIds ensures that no value is present for CollectionIds, not even an explicit nil
### GetDebug

`func (o *MatchFilters) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *MatchFilters) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *MatchFilters) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *MatchFilters) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### GetDebugTypes

`func (o *MatchFilters) GetDebugTypes() []*string`

GetDebugTypes returns the DebugTypes field if non-nil, zero value otherwise.

### GetDebugTypesOk

`func (o *MatchFilters) GetDebugTypesOk() (*[]*string, bool)`

GetDebugTypesOk returns a tuple with the DebugTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugTypes

`func (o *MatchFilters) SetDebugTypes(v []*string)`

SetDebugTypes sets DebugTypes field to given value.

### HasDebugTypes

`func (o *MatchFilters) HasDebugTypes() bool`

HasDebugTypes returns a boolean if a field has been set.

### SetDebugTypesNil

`func (o *MatchFilters) SetDebugTypesNil(b bool)`

 SetDebugTypesNil sets the value for DebugTypes to be an explicit nil

### UnsetDebugTypes
`func (o *MatchFilters) UnsetDebugTypes()`

UnsetDebugTypes ensures that no value is present for DebugTypes, not even an explicit nil
### GetFunctionIds

`func (o *MatchFilters) GetFunctionIds() []int64`

GetFunctionIds returns the FunctionIds field if non-nil, zero value otherwise.

### GetFunctionIdsOk

`func (o *MatchFilters) GetFunctionIdsOk() (*[]int64, bool)`

GetFunctionIdsOk returns a tuple with the FunctionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionIds

`func (o *MatchFilters) SetFunctionIds(v []int64)`

SetFunctionIds sets FunctionIds field to given value.

### HasFunctionIds

`func (o *MatchFilters) HasFunctionIds() bool`

HasFunctionIds returns a boolean if a field has been set.

### SetFunctionIdsNil

`func (o *MatchFilters) SetFunctionIdsNil(b bool)`

 SetFunctionIdsNil sets the value for FunctionIds to be an explicit nil

### UnsetFunctionIds
`func (o *MatchFilters) UnsetFunctionIds()`

UnsetFunctionIds ensures that no value is present for FunctionIds, not even an explicit nil
### GetIncludeUserDebug

`func (o *MatchFilters) GetIncludeUserDebug() bool`

GetIncludeUserDebug returns the IncludeUserDebug field if non-nil, zero value otherwise.

### GetIncludeUserDebugOk

`func (o *MatchFilters) GetIncludeUserDebugOk() (*bool, bool)`

GetIncludeUserDebugOk returns a tuple with the IncludeUserDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeUserDebug

`func (o *MatchFilters) SetIncludeUserDebug(v bool)`

SetIncludeUserDebug sets IncludeUserDebug field to given value.

### HasIncludeUserDebug

`func (o *MatchFilters) HasIncludeUserDebug() bool`

HasIncludeUserDebug returns a boolean if a field has been set.

### GetPlatforms

`func (o *MatchFilters) GetPlatforms() []string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *MatchFilters) GetPlatformsOk() (*[]string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *MatchFilters) SetPlatforms(v []string)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *MatchFilters) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### SetPlatformsNil

`func (o *MatchFilters) SetPlatformsNil(b bool)`

 SetPlatformsNil sets the value for Platforms to be an explicit nil

### UnsetPlatforms
`func (o *MatchFilters) UnsetPlatforms()`

UnsetPlatforms ensures that no value is present for Platforms, not even an explicit nil
### GetUserIds

`func (o *MatchFilters) GetUserIds() []int64`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *MatchFilters) GetUserIdsOk() (*[]int64, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *MatchFilters) SetUserIds(v []int64)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *MatchFilters) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### SetUserIdsNil

`func (o *MatchFilters) SetUserIdsNil(b bool)`

 SetUserIdsNil sets the value for UserIds to be an explicit nil

### UnsetUserIds
`func (o *MatchFilters) UnsetUserIds()`

UnsetUserIds ensures that no value is present for UserIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


