# ConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DashboardUrl** | Pointer to **string** | The domain of the RevEng.AI platform you are connected to | [optional] [default to ""]
**MaxFileSizeBytes** | **int32** | Maximum file size (in bytes) that can be uploaded for analysis | 
**AiDecompilerUnsupportedLanguages** | **[]string** | List of programming languages that are not supported for AI decompilation | 
**AiDecompilerSupportedModels** | **[]string** | List of models that support AI decompilation | 

## Methods

### NewConfigResponse

`func NewConfigResponse(maxFileSizeBytes int32, aiDecompilerUnsupportedLanguages []string, aiDecompilerSupportedModels []string, ) *ConfigResponse`

NewConfigResponse instantiates a new ConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigResponseWithDefaults

`func NewConfigResponseWithDefaults() *ConfigResponse`

NewConfigResponseWithDefaults instantiates a new ConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDashboardUrl

`func (o *ConfigResponse) GetDashboardUrl() string`

GetDashboardUrl returns the DashboardUrl field if non-nil, zero value otherwise.

### GetDashboardUrlOk

`func (o *ConfigResponse) GetDashboardUrlOk() (*string, bool)`

GetDashboardUrlOk returns a tuple with the DashboardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardUrl

`func (o *ConfigResponse) SetDashboardUrl(v string)`

SetDashboardUrl sets DashboardUrl field to given value.

### HasDashboardUrl

`func (o *ConfigResponse) HasDashboardUrl() bool`

HasDashboardUrl returns a boolean if a field has been set.

### GetMaxFileSizeBytes

`func (o *ConfigResponse) GetMaxFileSizeBytes() int32`

GetMaxFileSizeBytes returns the MaxFileSizeBytes field if non-nil, zero value otherwise.

### GetMaxFileSizeBytesOk

`func (o *ConfigResponse) GetMaxFileSizeBytesOk() (*int32, bool)`

GetMaxFileSizeBytesOk returns a tuple with the MaxFileSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFileSizeBytes

`func (o *ConfigResponse) SetMaxFileSizeBytes(v int32)`

SetMaxFileSizeBytes sets MaxFileSizeBytes field to given value.


### GetAiDecompilerUnsupportedLanguages

`func (o *ConfigResponse) GetAiDecompilerUnsupportedLanguages() []string`

GetAiDecompilerUnsupportedLanguages returns the AiDecompilerUnsupportedLanguages field if non-nil, zero value otherwise.

### GetAiDecompilerUnsupportedLanguagesOk

`func (o *ConfigResponse) GetAiDecompilerUnsupportedLanguagesOk() (*[]string, bool)`

GetAiDecompilerUnsupportedLanguagesOk returns a tuple with the AiDecompilerUnsupportedLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiDecompilerUnsupportedLanguages

`func (o *ConfigResponse) SetAiDecompilerUnsupportedLanguages(v []string)`

SetAiDecompilerUnsupportedLanguages sets AiDecompilerUnsupportedLanguages field to given value.


### GetAiDecompilerSupportedModels

`func (o *ConfigResponse) GetAiDecompilerSupportedModels() []string`

GetAiDecompilerSupportedModels returns the AiDecompilerSupportedModels field if non-nil, zero value otherwise.

### GetAiDecompilerSupportedModelsOk

`func (o *ConfigResponse) GetAiDecompilerSupportedModelsOk() (*[]string, bool)`

GetAiDecompilerSupportedModelsOk returns a tuple with the AiDecompilerSupportedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiDecompilerSupportedModels

`func (o *ConfigResponse) SetAiDecompilerSupportedModels(v []string)`

SetAiDecompilerSupportedModels sets AiDecompilerSupportedModels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


