# ReAnalysisForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** | Tags associated with the analysis | [optional] [default to {}]
**CommandLineArgs** | Pointer to **string** | Command line arguments for dynamic execution | [optional] [default to ""]
**Priority** | Pointer to **int32** | Priority of the analysis | [optional] [default to 0]
**Essential** | Pointer to **bool** | Only runs essential parts of the analysis, skips tags/sbom/cves etc. | [optional] [default to true]
**ModelName** | Pointer to **NullableString** |  | [optional] 
**NoCache** | Pointer to **bool** | When enabled, skips using cached data within the processing. | [optional] [default to false]

## Methods

### NewReAnalysisForm

`func NewReAnalysisForm() *ReAnalysisForm`

NewReAnalysisForm instantiates a new ReAnalysisForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReAnalysisFormWithDefaults

`func NewReAnalysisFormWithDefaults() *ReAnalysisForm`

NewReAnalysisFormWithDefaults instantiates a new ReAnalysisForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *ReAnalysisForm) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ReAnalysisForm) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ReAnalysisForm) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ReAnalysisForm) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCommandLineArgs

`func (o *ReAnalysisForm) GetCommandLineArgs() string`

GetCommandLineArgs returns the CommandLineArgs field if non-nil, zero value otherwise.

### GetCommandLineArgsOk

`func (o *ReAnalysisForm) GetCommandLineArgsOk() (*string, bool)`

GetCommandLineArgsOk returns a tuple with the CommandLineArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineArgs

`func (o *ReAnalysisForm) SetCommandLineArgs(v string)`

SetCommandLineArgs sets CommandLineArgs field to given value.

### HasCommandLineArgs

`func (o *ReAnalysisForm) HasCommandLineArgs() bool`

HasCommandLineArgs returns a boolean if a field has been set.

### GetPriority

`func (o *ReAnalysisForm) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ReAnalysisForm) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ReAnalysisForm) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ReAnalysisForm) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetEssential

`func (o *ReAnalysisForm) GetEssential() bool`

GetEssential returns the Essential field if non-nil, zero value otherwise.

### GetEssentialOk

`func (o *ReAnalysisForm) GetEssentialOk() (*bool, bool)`

GetEssentialOk returns a tuple with the Essential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEssential

`func (o *ReAnalysisForm) SetEssential(v bool)`

SetEssential sets Essential field to given value.

### HasEssential

`func (o *ReAnalysisForm) HasEssential() bool`

HasEssential returns a boolean if a field has been set.

### GetModelName

`func (o *ReAnalysisForm) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *ReAnalysisForm) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *ReAnalysisForm) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *ReAnalysisForm) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### SetModelNameNil

`func (o *ReAnalysisForm) SetModelNameNil(b bool)`

 SetModelNameNil sets the value for ModelName to be an explicit nil

### UnsetModelName
`func (o *ReAnalysisForm) UnsetModelName()`

UnsetModelName ensures that no value is present for ModelName, not even an explicit nil
### GetNoCache

`func (o *ReAnalysisForm) GetNoCache() bool`

GetNoCache returns the NoCache field if non-nil, zero value otherwise.

### GetNoCacheOk

`func (o *ReAnalysisForm) GetNoCacheOk() (*bool, bool)`

GetNoCacheOk returns a tuple with the NoCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoCache

`func (o *ReAnalysisForm) SetNoCache(v bool)`

SetNoCache sets NoCache field to given value.

### HasNoCache

`func (o *ReAnalysisForm) HasNoCache() bool`

HasNoCache returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


