# RequeueAnalysisRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandLineArgs** | Pointer to **string** | Command-line arguments passed to the sample when enable_dynamic_execution is set | [optional] 
**EnableDynamicExecution** | Pointer to **bool** | Run the new Analysis with dynamic execution (sandbox) enabled | [optional] 
**GenerateCapabilities** | Pointer to **bool** | Regenerate capabilities for the new Analysis instead of copying the original Binary&#39;s existing ones | [optional] 
**NoCache** | Pointer to **bool** |  | [optional] 
**ScrapeTags** | Pointer to **bool** | Re-scrape third-party tags for the new Analysis instead of copying the original Binary&#39;s existing ones | [optional] 
**Tags** | Pointer to **[]string** | Additional Tags to apply to the requeued Binary, on top of any carried over because generate_capabilities/scrape_tags are false | [optional] 

## Methods

### NewRequeueAnalysisRequest

`func NewRequeueAnalysisRequest() *RequeueAnalysisRequest`

NewRequeueAnalysisRequest instantiates a new RequeueAnalysisRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequeueAnalysisRequestWithDefaults

`func NewRequeueAnalysisRequestWithDefaults() *RequeueAnalysisRequest`

NewRequeueAnalysisRequestWithDefaults instantiates a new RequeueAnalysisRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandLineArgs

`func (o *RequeueAnalysisRequest) GetCommandLineArgs() string`

GetCommandLineArgs returns the CommandLineArgs field if non-nil, zero value otherwise.

### GetCommandLineArgsOk

`func (o *RequeueAnalysisRequest) GetCommandLineArgsOk() (*string, bool)`

GetCommandLineArgsOk returns a tuple with the CommandLineArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineArgs

`func (o *RequeueAnalysisRequest) SetCommandLineArgs(v string)`

SetCommandLineArgs sets CommandLineArgs field to given value.

### HasCommandLineArgs

`func (o *RequeueAnalysisRequest) HasCommandLineArgs() bool`

HasCommandLineArgs returns a boolean if a field has been set.

### GetEnableDynamicExecution

`func (o *RequeueAnalysisRequest) GetEnableDynamicExecution() bool`

GetEnableDynamicExecution returns the EnableDynamicExecution field if non-nil, zero value otherwise.

### GetEnableDynamicExecutionOk

`func (o *RequeueAnalysisRequest) GetEnableDynamicExecutionOk() (*bool, bool)`

GetEnableDynamicExecutionOk returns a tuple with the EnableDynamicExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDynamicExecution

`func (o *RequeueAnalysisRequest) SetEnableDynamicExecution(v bool)`

SetEnableDynamicExecution sets EnableDynamicExecution field to given value.

### HasEnableDynamicExecution

`func (o *RequeueAnalysisRequest) HasEnableDynamicExecution() bool`

HasEnableDynamicExecution returns a boolean if a field has been set.

### GetGenerateCapabilities

`func (o *RequeueAnalysisRequest) GetGenerateCapabilities() bool`

GetGenerateCapabilities returns the GenerateCapabilities field if non-nil, zero value otherwise.

### GetGenerateCapabilitiesOk

`func (o *RequeueAnalysisRequest) GetGenerateCapabilitiesOk() (*bool, bool)`

GetGenerateCapabilitiesOk returns a tuple with the GenerateCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateCapabilities

`func (o *RequeueAnalysisRequest) SetGenerateCapabilities(v bool)`

SetGenerateCapabilities sets GenerateCapabilities field to given value.

### HasGenerateCapabilities

`func (o *RequeueAnalysisRequest) HasGenerateCapabilities() bool`

HasGenerateCapabilities returns a boolean if a field has been set.

### GetNoCache

`func (o *RequeueAnalysisRequest) GetNoCache() bool`

GetNoCache returns the NoCache field if non-nil, zero value otherwise.

### GetNoCacheOk

`func (o *RequeueAnalysisRequest) GetNoCacheOk() (*bool, bool)`

GetNoCacheOk returns a tuple with the NoCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoCache

`func (o *RequeueAnalysisRequest) SetNoCache(v bool)`

SetNoCache sets NoCache field to given value.

### HasNoCache

`func (o *RequeueAnalysisRequest) HasNoCache() bool`

HasNoCache returns a boolean if a field has been set.

### GetScrapeTags

`func (o *RequeueAnalysisRequest) GetScrapeTags() bool`

GetScrapeTags returns the ScrapeTags field if non-nil, zero value otherwise.

### GetScrapeTagsOk

`func (o *RequeueAnalysisRequest) GetScrapeTagsOk() (*bool, bool)`

GetScrapeTagsOk returns a tuple with the ScrapeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrapeTags

`func (o *RequeueAnalysisRequest) SetScrapeTags(v bool)`

SetScrapeTags sets ScrapeTags field to given value.

### HasScrapeTags

`func (o *RequeueAnalysisRequest) HasScrapeTags() bool`

HasScrapeTags returns a boolean if a field has been set.

### GetTags

`func (o *RequeueAnalysisRequest) GetTags() []*string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RequeueAnalysisRequest) GetTagsOk() (*[]*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RequeueAnalysisRequest) SetTags(v []*string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RequeueAnalysisRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *RequeueAnalysisRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *RequeueAnalysisRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


