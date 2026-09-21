# GetConfigOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiDecompilerUnsupportedLanguages** | **[]string** | Source languages AI decompilation does not support | 
**MaxFileSizeBytes** | **int64** | Largest binary the calling user may submit for analysis, in bytes | 

## Methods

### NewGetConfigOutputBody

`func NewGetConfigOutputBody(aiDecompilerUnsupportedLanguages []string, maxFileSizeBytes int64, ) *GetConfigOutputBody`

NewGetConfigOutputBody instantiates a new GetConfigOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConfigOutputBodyWithDefaults

`func NewGetConfigOutputBodyWithDefaults() *GetConfigOutputBody`

NewGetConfigOutputBodyWithDefaults instantiates a new GetConfigOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiDecompilerUnsupportedLanguages

`func (o *GetConfigOutputBody) GetAiDecompilerUnsupportedLanguages() []string`

GetAiDecompilerUnsupportedLanguages returns the AiDecompilerUnsupportedLanguages field if non-nil, zero value otherwise.

### GetAiDecompilerUnsupportedLanguagesOk

`func (o *GetConfigOutputBody) GetAiDecompilerUnsupportedLanguagesOk() (*[]string, bool)`

GetAiDecompilerUnsupportedLanguagesOk returns a tuple with the AiDecompilerUnsupportedLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiDecompilerUnsupportedLanguages

`func (o *GetConfigOutputBody) SetAiDecompilerUnsupportedLanguages(v []string)`

SetAiDecompilerUnsupportedLanguages sets AiDecompilerUnsupportedLanguages field to given value.


### SetAiDecompilerUnsupportedLanguagesNil

`func (o *GetConfigOutputBody) SetAiDecompilerUnsupportedLanguagesNil(b bool)`

 SetAiDecompilerUnsupportedLanguagesNil sets the value for AiDecompilerUnsupportedLanguages to be an explicit nil

### UnsetAiDecompilerUnsupportedLanguages
`func (o *GetConfigOutputBody) UnsetAiDecompilerUnsupportedLanguages()`

UnsetAiDecompilerUnsupportedLanguages ensures that no value is present for AiDecompilerUnsupportedLanguages, not even an explicit nil
### GetMaxFileSizeBytes

`func (o *GetConfigOutputBody) GetMaxFileSizeBytes() int64`

GetMaxFileSizeBytes returns the MaxFileSizeBytes field if non-nil, zero value otherwise.

### GetMaxFileSizeBytesOk

`func (o *GetConfigOutputBody) GetMaxFileSizeBytesOk() (*int64, bool)`

GetMaxFileSizeBytesOk returns a tuple with the MaxFileSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFileSizeBytes

`func (o *GetConfigOutputBody) SetMaxFileSizeBytes(v int64)`

SetMaxFileSizeBytes sets MaxFileSizeBytes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


